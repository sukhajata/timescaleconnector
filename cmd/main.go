package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/sukhajata/devicetwin/pkg/db"
	"github.com/sukhajata/devicetwin/pkg/errorhelper"
	"github.com/sukhajata/devicetwin/pkg/grpchelper"
	"github.com/sukhajata/devicetwin/pkg/ppmqtt"
	"github.com/sukhajata/timescaleconnector/api"
	"go.elastic.co/ecslogrus"
	"math/rand"
	"os"
	"time"

	"github.com/sukhajata/timescaleconnector/internal/messagedecoder"
	"github.com/sukhajata/timescaleconnector/internal/timescale"

	//helpers "powerpilot.visualstudio.com/PowerPilot/_git/pphelpers.git"
	pbLogger "github.com/sukhajata/pplogger"
	//ppmqtt "powerpilot.visualstudio.com/PowerPilot/_git/ppmqtt.git"
	"github.com/sukhajata/devicetwin/pkg/loggerhelper"
)

var (
	grpcLoggerServerAddress = getenv("loggerServiceAddress", "logger-service:9031")
	mqttBroker              = getenv("mqttBroker", "")
	mqttUsername            = getenv("mqttUsername", "superAdmin")
	mqttPassword            = getenv("mqttPassword", "powerpilot")
	mqttTopic               = getenv("mqttTopic", "$share/fast-path-group/application/powerpilot/uplink/#")
	psqlURL                 = getenv("postgresUrl", "postgres://postgres:tea@dev-timescale.dev.svc/devpower?sslmode=require")

	log          *logrus.Logger
	decoder      *messagedecoder.Decoder
	mqttClient   *ppmqtt.PPClient
	dbClient     *timescale.TimescaleClient
	loggerHelper loggerhelper.Helper

	//mqttInputChannel chan ppmqtt.Message
	decodedChannel chan interface{}
)

func getenv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func connectMQTT() {
	var err error
	mqttClient, err = ppmqtt.NewClient(mqttBroker, mqttUsername, mqttPassword, "timescale-connector")
	errorhelper.PanicOnError(err)

	err = mqttClient.Subscribe(mqttTopic)
	errorhelper.PanicOnError(err)

	loggerhelper.WriteToLog("Connected to mqtt broker")

	// process mqtt uplinks on a couple of new go routines
	// mqtt client will close ReceiveChan on error, ending any previous running of this function
	decoder.ProcessMessages(mqttClient.ReceiveChan, decodedChannel, log)
	decoder.ProcessMessages(mqttClient.ReceiveChan, decodedChannel, log)

	// listen for mqtt errors and reconnect
	go func() {
		err := <-mqttClient.ErrorChan
		loggerhelper.WriteToLog(err.Error())
		connectMQTT()
	}()

}

func connectTimescale() {
	// db engine
	dbEngine, err := db.NewTimescaleEngine(psqlURL)
	errorhelper.PanicOnError(err)

	dbClient = timescale.NewTimescaleClient(dbEngine)

	// listen for db errors and reconnect
	go func() {
		for err := range dbClient.ErrorChan {
			switch v := err.(type) {
			case *db.FatalError:
				loggerhelper.WriteToLog(fmt.Sprintf("Fatal error, reconnecting: %v", v.Error()))
				dbEngine.Close()
				connectTimescale()
			default:
				loggerhelper.WriteToLog(v.Error())
			}
		}
	}()

	loggerhelper.WriteToLog("Connected to timescale")

	// listen for decoded messages on a couple of new go routines
	dbClient.Listen(decodedChannel)
	dbClient.Listen(decodedChannel)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// elastic logging
	log = logrus.New()
	log.SetFormatter(&ecslogrus.Formatter{
		DataKey: "labels",
	})
	log.SetReportCaller(true)

	// logger
	conn, err := grpchelper.ConnectGRPC(grpcLoggerServerAddress)
	errorhelper.PanicOnError(err)
	defer func() {
		err := conn.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()
	grpcLogger := pbLogger.NewLoggerServiceClient(conn)

	loggerHelper = loggerhelper.NewHelper(grpcLogger, "timescale-connector")

	// channel for messages from mqtt client
	//mqttInputChannel = make(chan ppmqtt.Message, 2)

	// channel for decoded messages
	decodedChannel = make(chan interface{})

	// message decoder
	decoder = messagedecoder.NewDecoder(loggerHelper)

	// mqtt
	connectMQTT()

	// timescale
	connectTimescale()

	fmt.Println("started everything")

	// start health check server
	api.NewHealthCheckServer()

}
