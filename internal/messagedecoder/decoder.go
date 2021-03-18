package messagedecoder

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"strings"

	"github.com/sukhajata/devicetwin/pkg/loggerhelper"
	"github.com/sukhajata/devicetwin/pkg/ppmqtt"
	pbLogger "github.com/sukhajata/pplogger"
	pb "github.com/sukhajata/ppmessage/ppuplink"
	"google.golang.org/protobuf/proto"
)

// Decoder - decodes mqtt messages
type Decoder struct {
	loggerHelper loggerhelper.Helper
}

func NewDecoder(loggerHelper loggerhelper.Helper) *Decoder {
	return &Decoder{
		loggerHelper: loggerHelper,
	}
}

func IsClosed(ch <-chan interface{}) bool {
	select {
	case <-ch:
		return true
	default:
	}

	return false
}

// ProcessMessages - processes uplink messages on inChan, sends result on outChan
func (p *Decoder) ProcessMessages(inChan <-chan ppmqtt.Message, outChan chan<- interface{}, log *logrus.Logger) {
	go func() {
		for msg := range inChan {
			parts := strings.Split(msg.Topic, "/")
			if len(parts) < 4 {
				loggerhelper.WriteToLog("Topic not valid, skipping")
				return
			}
			topic := parts[3]

			if topic == "alarm" {
				var alarmMsg pb.AlarmMessage
				err := proto.Unmarshal(msg.Payload, &alarmMsg)
				if err != nil {
					p.loggerHelper.LogError("alarm", err.Error(), pbLogger.ErrorMessage_FATAL)
					log.WithError(err).Fatal("unmarshal AlarmMessage failed")
					return
				}
				outChan <- alarmMsg
			} else if topic == "circuitenergy" {
				var circuitEnergyMsg pb.CircuitEnergyMessage
				err := proto.Unmarshal(msg.Payload, &circuitEnergyMsg)
				if err != nil {
					p.loggerHelper.LogError("circuitenergy", err.Error(), pbLogger.ErrorMessage_FATAL)
					log.WithError(err).Fatal("unmarshal CircuitEnergyMessage failed")
					return
				}
				outChan <- circuitEnergyMsg
			} else if topic == "circuitload" {
				var circuitLoadMsg pb.CircuitLoadMessage
				err := proto.Unmarshal(msg.Payload, &circuitLoadMsg)
				if err != nil {
					p.loggerHelper.LogError("circuitload", err.Error(), pbLogger.ErrorMessage_FATAL)
					log.WithError(err).Fatal("unmarshal CircuitLoadMessage failed")
					return
				}
				outChan <- circuitLoadMsg
			} else if topic == "energy" {
				var energyMsg pb.EnergyMessage
				err := proto.Unmarshal(msg.Payload, &energyMsg)
				if err != nil {
					p.loggerHelper.LogError("energy", err.Error(), pbLogger.ErrorMessage_FATAL)
					log.WithError(err).Fatal("unmarshal EnergyMessage failed")
					return
				}
				outChan <- energyMsg
			} else if topic == "gateway" {
				var gatewayMsg pb.GatewayMessage
				err := proto.Unmarshal(msg.Payload, &gatewayMsg)
				if err != nil {
					p.loggerHelper.LogError("gateway", err.Error(), pbLogger.ErrorMessage_FATAL)
					log.WithError(err).Fatal("unmarshal GatewayMessage failed")
					return
				}
				outChan <- gatewayMsg
			} else if topic == "geoscan" {
				var geoscanMsg pb.GeoscanMessage
				err := proto.Unmarshal(msg.Payload, &geoscanMsg)
				if err != nil {
					p.loggerHelper.LogError("geoscan", err.Error(), pbLogger.ErrorMessage_FATAL)
					log.WithError(err).Fatal("unmarshal GeoscanMessage failed")
					return
				}
				outChan <- geoscanMsg
			} else if topic == "harmonicslower" {
				var harmonicsLowerMsg pb.HarmonicsLowerMessage
				err := proto.Unmarshal(msg.Payload, &harmonicsLowerMsg)
				if err != nil {
					p.loggerHelper.LogError("harmonicslower", err.Error(), pbLogger.ErrorMessage_FATAL)
					log.WithError(err).Fatal("unmarshal HarmonicsLowerMessage failed")
					return
				}
				outChan <- harmonicsLowerMsg
			} else if topic == "harmonicsupper" {
				var harmonicsUpperMsg pb.HarmonicsUpperMessage
				err := proto.Unmarshal(msg.Payload, &harmonicsUpperMsg)
				if err != nil {
					p.loggerHelper.LogError("harmonicsupper", err.Error(), pbLogger.ErrorMessage_FATAL)
					log.WithError(err).Fatal("unmarshal HarmonicsUpperMessage failed")
					return
				}
				outChan <- harmonicsUpperMsg
			} else if topic == "hvalarm" {
				var hvAlarmMsg pb.HVAlarmMessage
				err := proto.Unmarshal(msg.Payload, &hvAlarmMsg)
				if err != nil {
					p.loggerHelper.LogError("hvalarm", err.Error(), pbLogger.ErrorMessage_FATAL)
					log.WithError(err).Fatal("unmarshal HVAlarmMessage failed")
					return
				}
				outChan <- hvAlarmMsg
			} else if topic == "inst" {
				var instmsg pb.InstMessage
				err := proto.Unmarshal(msg.Payload, &instmsg)
				if err != nil {
					p.loggerHelper.LogError("inst", err.Error(), pbLogger.ErrorMessage_FATAL)
					log.WithError(err).Fatal("unmarshal InstMessage failed")
					return
				}
				outChan <- instmsg
			} else if topic == "instp2p" {
				var instp2pMsg pb.InstP2PMessage
				err := proto.Unmarshal(msg.Payload, &instp2pMsg)
				if err != nil {
					p.loggerHelper.LogError("instp2p", err.Error(), pbLogger.ErrorMessage_FATAL)
					log.WithError(err).Fatal("unmarshal InstP2PMessage failed")
					return
				}
				outChan <- instp2pMsg
			} else if topic == "meterstatus" {
				var meterstatus pb.MeterStatusMessage
				err := proto.Unmarshal(msg.Payload, &meterstatus)
				if err != nil {
					p.loggerHelper.LogError("meterstatus", err.Error(), pbLogger.ErrorMessage_FATAL)
					log.WithError(err).Fatal("unmarshal MeterStatusMessage failed")
					return
				}
				outChan <- meterstatus
			} else if topic == "pq" {
				var pqmsg pb.PQMessage
				err := proto.Unmarshal(msg.Payload, &pqmsg)
				if err != nil {
					p.loggerHelper.LogError("pq", err.Error(), pbLogger.ErrorMessage_FATAL)
					log.WithError(err).Fatal("unmarshal PQMessage failed")
					return
				}
				outChan <- pqmsg
			} else if topic == "pqevents" {
				var pqEventsMsg pb.PQEventsMessage
				err := proto.Unmarshal(msg.Payload, &pqEventsMsg)
				if err != nil {
					p.loggerHelper.LogError("pqevents", err.Error(), pbLogger.ErrorMessage_FATAL)
					log.WithError(err).Fatal("unmarshal PQEventsMessage failed")
					return
				}
				outChan <- pqEventsMsg
			} else if topic == "processed" {
				var processedMsg pb.ProcessedMessage
				err := proto.Unmarshal(msg.Payload, &processedMsg)
				if err != nil {
					p.loggerHelper.LogError("processed", err.Error(), pbLogger.ErrorMessage_FATAL)
					log.WithError(err).Fatal("unmarshal ProcessedMessage failed")
					return
				}
				outChan <- processedMsg
			} else if topic == "resendresponse" {
				var resendResponse pb.ResendResponseMessage
				err := proto.Unmarshal(msg.Payload, &resendResponse)
				if err != nil {
					p.loggerHelper.LogError("resendresponse", err.Error(), pbLogger.ErrorMessage_FATAL)
					log.WithError(err).Fatal("unmarshal ResendResponseMessage failed")
					return
				}
				outChan <- resendResponse
			} else if topic == "s11pq" {
				var s11pq pb.S11PQMessage
				err := proto.Unmarshal(msg.Payload, &s11pq)
				if err != nil {
					p.loggerHelper.LogError("s11pq", err.Error(), pbLogger.ErrorMessage_FATAL)
					log.WithError(err).Fatal("unmarshal S11PQMessage failed")
					return
				}
				outChan <- s11pq
			} else if topic == "uplink" {
				var uplink pb.UplinkMessage
				err := proto.Unmarshal(msg.Payload, &uplink)
				if err != nil {
					p.loggerHelper.LogError("uplink", err.Error(), pbLogger.ErrorMessage_FATAL)
					log.WithError(err).Fatal("unmarshal UplinkMessage failed")
					return
				}
				outChan <- uplink
			} else if topic == "voltagestats" {
				var voltagestats pb.VoltageStatsMessage
				err := proto.Unmarshal(msg.Payload, &voltagestats)
				if err != nil {
					p.loggerHelper.LogError("voltagestats", err.Error(), pbLogger.ErrorMessage_FATAL)
					log.WithError(err).Fatal("unmarshal VoltageStatsMessage failed")
					return
				}
				outChan <- voltagestats
			} else {
				fmt.Println("unknown message type")
			}
		}
	}()

}
