package timescale

import (
	"fmt"
	"github.com/sukhajata/devicetwin/pkg/db"
	"github.com/sukhajata/devicetwin/pkg/loggerhelper"
	"time"

	pb "github.com/sukhajata/ppmessage/ppuplink"
)

func timestampToString(msec int64) string {
	tm := time.Unix(msec/1000, 0)
	layout := "2006-01-02 15:04:05+00"
	return tm.UTC().Format(layout)
}

// TimescaleClient implements DBClient
type TimescaleClient struct {
	dbEngine  db.SQLEngine
	ErrorChan chan error
}

func NewTimescaleClient(dbEngine db.SQLEngine) *TimescaleClient {
	return &TimescaleClient{
		dbEngine:  dbEngine,
		ErrorChan: make(chan error, 2),
	}
}

// Listen for messages
func (c *TimescaleClient) Listen(decodedChan <-chan interface{}) {
	go func() {
		for msg := range decodedChan {
			var err error
			switch t := msg.(type) {
			case pb.AlarmMessage:
				err = c.InsertAlarmMsg(t)
			case pb.CircuitEnergyMessage:
				err = c.InsertCircuitEnergyMsg(t)
			case pb.CircuitLoadMessage:
				err = c.InsertCircuitLoadMessage(t)
			case pb.EnergyMessage:
				err = c.InsertEnergyMsg(t)
			case pb.GatewayMessage:
				err = c.InsertGatewayMsg(t)
			case pb.GeoscanMessage:
				err = c.InsertGeoscanMsg(t)
			case pb.HarmonicsLowerMessage:
				err = c.InsertHarmonicsLowerMsg(t)
			case pb.HarmonicsUpperMessage:
				err = c.InsertHarmonicsUpperMsg(t)
			case pb.HVAlarmMessage:
				err = c.InsertHVAlarmMsg(t)
			case pb.InstMessage:
				err = c.InsertInstMsg(t)
			case pb.InstP2PMessage:
				err = c.InsertInstP2PMsg(t)
			case pb.MeterStatusMessage:
				err = c.InsertMeterStatus(t)
			case pb.PQMessage:
				err = c.InsertPQMsg(t)
			case pb.PQEventsMessage:
				err = c.InsertPQEventsMsg(t)
			case pb.ProcessedMessage:
				err = c.InsertProcessedMsg(t)
			case pb.ResendResponseMessage:
				err = c.InsertResendResponseMsg(t)
			case pb.S11PQMessage:
				err = c.InsertS11PQ(t)
			case pb.UplinkMessage:
				err = c.InsertUplinkMsg(t)
				loggerhelper.WriteToLog(fmt.Sprintf("%v FCnt %v %v", t.Deviceeui, t.Fctn, t.Rawdata))
			case pb.VoltageStatsMessage:
				err = c.InsertVoltageStats(t)
			default:
				loggerhelper.WriteToLog(fmt.Sprintf("Unhandled message of type %T. Value %v", t, t))
			}
			if err != nil {
				c.ErrorChan <- err
				// c.loggerHelper.LogError("insertMessage", err.Error(), pbLogger.ErrorMessage_WARNING)
			}
		}
	}()
}

// InsertAlarmMsg inserts alarm msg
func (c *TimescaleClient) InsertAlarmMsg(msg pb.AlarmMessage) error {
	sql := `insert into "ALARM"(
		"DEVICEEUI",
		"TIMESTAMP",
		"PHASEID",
		"ALARMTYPE",
		"VALUE"
		) values($1, $2, $3, $4, $5)`
	dt := timestampToString(int64(msg.Timesent))

	return c.dbEngine.Exec(sql,
		msg.Deviceeui,
		dt,
		msg.Phaseid,
		msg.Alarmtype,
		msg.Value,
	)

}

// InsertCircuitEnergyMsg inserts circuit energy msg
func (c *TimescaleClient) InsertCircuitEnergyMsg(msg pb.CircuitEnergyMessage) error {
	sql := `insert into "CIRCUIT_ENERGY"(
		"DEVICEEUI",
		"TIMESTAMP",
		"CIRCUITNUMBER",
		"REALNETENERGY_PHASE_0",
		"REACTIVENETENERGY_PHASE_0",
		"REALNETENERGY_PHASE_1",
		"REACTIVENETENERGY_PHASE_1",
		"REALNETENERGY_PHASE_2",
		"REACTIVENETENERGY_PHASE_2"
		) values($1, $2, $3, $4, $5, $6, $7, $8, $9)`
	dt := timestampToString(int64(msg.Timesent))

	return c.dbEngine.Exec(sql,
		msg.Deviceeui,
		dt,
		msg.Circuitnumber,
		msg.RealnetenergyPhase_0,
		msg.ReactivenetenergyPhase_0,
		msg.RealnetenergyPhase_1,
		msg.ReactivenetenergyPhase_1,
		msg.RealnetenergyPhase_2,
		msg.ReactivenetenergyPhase_2,
	)

}

// InsertCircuitLoadMessage inserts circuit load msg
func (c *TimescaleClient) InsertCircuitLoadMessage(msg pb.CircuitLoadMessage) error {
	sql := `insert into "CIRCUIT_LOAD"(
		"DEVICEEUI",
    "TIMESTAMP",
		"CIRCUITNUMBER",
		"CURRENTSMA_PHASE_0",
   	"CURRENTSMA_PHASE_1",
    "CURRENTSMA_PHASE_2",
    "POWERACTIVESMA_PHASE_0",
    "POWERACTIVESMA_PHASE_1",
    "POWERACTIVESMA_PHASE_2",
    "POWERREACTIVESMA_PHASE_0",
    "POWERREACTIVESMA_PHASE_1",
    "POWERREACTIVESMA_PHASE_2",
    "POWERAPPARENTSMA_PHASE_0",
    "POWERAPPARENTSMA_PHASE_1",
    "POWERAPPARENTSMA_PHASE_2",
    "POWERACTIVEMAX_PHASE_0",
    "POWERACTIVEMAX_PHASE_1",
    "POWERACTIVEMAX_PHASE_2",
    "POWERREACTIVEMAX_PHASE_0",
    "POWERREACTIVEMAX_PHASE_1",
    "POWERREACTIVEMAX_PHASE_2",
    "POWERAPPARENTMAX_PHASE_0",
    "POWERAPPARENTMAX_PHASE_1",
    "POWERAPPARENTMAX_PHASE_2",
    "POWERACTIVEMIN_PHASE_0",
    "POWERACTIVEMIN_PHASE_1",
    "POWERACTIVEMIN_PHASE_2",
    "POWERREACTIVEMIN_PHASE_0",
    "POWERREACTIVEMIN_PHASE_1",
    "POWERREACTIVEMIN_PHASE_2",
    "POWERAPPARENTMIN_PHASE_0",
    "POWERAPPARENTMIN_PHASE_1",
		"POWERAPPARENTMIN_PHASE_2"
		) values($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33)`
	dt := timestampToString(int64(msg.Timesent))

	return c.dbEngine.Exec(sql,
		msg.Deviceeui,
		dt,
		msg.Circuitnumber,
		msg.CurrentsmaPhase_0,
		msg.CurrentsmaPhase_1,
		msg.CurrentsmaPhase_2,
		msg.PoweractivesmaPhase_0,
		msg.PoweractivesmaPhase_1,
		msg.PoweractivesmaPhase_2,
		msg.PowerreactivesmaPhase_0,
		msg.PowerreactivesmaPhase_1,
		msg.PowerreactivesmaPhase_2,
		msg.PowerapparentsmaPhase_0,
		msg.PowerapparentsmaPhase_1,
		msg.PowerapparentsmaPhase_2,
		msg.PoweractivemaxPhase_0,
		msg.PoweractivemaxPhase_1,
		msg.PoweractivemaxPhase_2,
		msg.PowerreactivemaxPhase_0,
		msg.PowerreactivemaxPhase_1,
		msg.PowerreactivemaxPhase_2,
		msg.PowerapparentmaxPhase_0,
		msg.PowerapparentmaxPhase_1,
		msg.PowerapparentmaxPhase_2,
		msg.PoweractiveminPhase_0,
		msg.PoweractiveminPhase_1,
		msg.PoweractiveminPhase_2,
		msg.PowerreactiveminPhase_0,
		msg.PowerreactiveminPhase_1,
		msg.PowerreactiveminPhase_2,
		msg.PowerapparentminPhase_0,
		msg.PowerapparentminPhase_1,
		msg.PowerapparentminPhase_2,
	)
}

// InsertEnergyMsg inserts energy msg
func (c *TimescaleClient) InsertEnergyMsg(msg pb.EnergyMessage) error {
	sql := `insert into "ENERGY"(
		"DEVICEEUI",
		"TIMESTAMP",
		"PHASEID",
		"ENERGYEXPORTREACTIVE",
		"ENERGYIMPORTREACTIVE",
		"ENERGYEXPORTREAL",
		"ENERGYIMPORTREAL"
		) values($1, $2, $3, $4, $5, $6, $7)`
	dt := timestampToString(int64(msg.Timesent))

	return c.dbEngine.Exec(sql,
		msg.Deviceeui,
		dt,
		msg.Phaseid,
		msg.Energyexportreactive,
		msg.Energyimportreactive,
		msg.Energyexportreal,
		msg.Energyimportreal,
	)
}

// InsertGatewayMsg inserts gateway msg
func (c *TimescaleClient) InsertGatewayMsg(msg pb.GatewayMessage) error {
	sql := `insert into "GATEWAYRX"(
		"DEVICEEUI",
    "TIMESTAMP",
    "GATEWAYID",
    "RSSI",
    "SNR",
    "LATITUDE",
    "LONGITUDE",
    "ALTITUDE",
    "GATEWAYTIMESTAMP",
    "FREQUENCY",
    "RAWDATA"
	) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`
	dt := timestampToString(int64(msg.Timesent))

	return c.dbEngine.Exec(sql,
		msg.Deviceeui,
		dt,
		msg.Gatewayid,
		msg.Rssi,
		msg.Snr,
		msg.Latitude,
		msg.Longitude,
		msg.Altitude,
		msg.Gatewaytimesent,
		msg.Frequency,
		msg.Rawdata,
	)
}

// InsertGeoscanMsg inserts geoscan msg
func (c *TimescaleClient) InsertGeoscanMsg(msg pb.GeoscanMessage) error {
	sql := `insert into "GEOSCAN"(
		"DEVICEEUI",
		"TIMESTAMP",
		"BSSID",
		"RSSID"
		) values($1, $2, $3, $4)`
	dt := timestampToString(int64(msg.Timesent))

	return c.dbEngine.Exec(sql,
		msg.Deviceeui,
		dt,
		msg.Bssid,
		msg.Rssid,
	)
}

// InsertHarmonicsLowerMsg insert
func (c *TimescaleClient) InsertHarmonicsLowerMsg(msg pb.HarmonicsLowerMessage) error {
	sql := `insert into "HARMONICS_LOWER"(
		"DEVICEEUI",
		"TIMESTAMP",
		"VOLTAGER",
		"CURRENTR",
		"VOLTAGEW",
		"CURRENTW",
		"VOLTAGEB",
		"CURRENTB",
		"ACTIVEPOWER",
		"REACTIVEPOWER",
		"THDVR",
		"THDIR",
		"THIRDHARMONICR",
		"FIFTHHARMONICR",
		"SEVENTHHARMONICR",
		"THDVW",
		"THDIW",
		"THIRDHARMONICW",
		"FIFTHHARMONICW",
		"SEVENTHHARMONICW",
		"THDVB",
		"THDIB",
		"THIRDHARMONICB",
		"FIFTHHARMONICB",
		"SEVENTHHARMONICB"
		) values($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25)`
	dt := timestampToString(int64(msg.Timesent))

	return c.dbEngine.Exec(sql,
		msg.Deviceeui,
		dt,
		msg.VoltageR,
		msg.CurrentR,
		msg.VoltageW,
		msg.CurrentW,
		msg.VoltageB,
		msg.CurrentB,
		msg.Activepower,
		msg.Reactivepower,
		msg.ThdvR,
		msg.ThdiR,
		msg.ThirdharmonicR,
		msg.FifthharmonicR,
		msg.SeventhharmonicR,
		msg.ThdvW,
		msg.ThdiW,
		msg.ThirdharmonicW,
		msg.FifthharmonicW,
		msg.SeventhharmonicW,
		msg.ThdvB,
		msg.ThdiB,
		msg.ThirdharmonicB,
		msg.FifthharmonicB,
		msg.SeventhharmonicB,
	)

}

// InsertHarmonicsUpperMsg insert
func (c *TimescaleClient) InsertHarmonicsUpperMsg(msg pb.HarmonicsUpperMessage) error {
	sql := `insert into "HARMONICS_UPPER"(
		"DEVICEEUI",
		"TIMESTAMP",
		"VOLTAGER",
		"CURRENTR",
		"VOLTAGEW",
		"CURRENTW",
		"VOLTAGEB",
		"CURRENTB",
		"ACTIVEPOWER",
		"REACTIVEPOWER",
		"THDVR",
		"THDIR",
		"NINTHHARMONICR",
		"ELEVENTHHARMONICR",
		"THIRTEENTHHARMONICR",
		"THDVW",
		"THDIW",
		"NINTHHARMONICW",
		"ELEVENTHHARMONICW",
		"THIRTEENTHHARMONICW",
		"THDVB",
		"THDIB",
		"NINTHHARMONICB",
		"ELEVENTHHARMONICB",
		"THIRTEENTHHARMONICB"
		) values($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25)`
	dt := timestampToString(int64(msg.Timesent))

	return c.dbEngine.Exec(sql,
		msg.Deviceeui,
		dt,
		msg.VoltageR,
		msg.CurrentR,
		msg.VoltageW,
		msg.CurrentW,
		msg.VoltageB,
		msg.CurrentB,
		msg.Activepower,
		msg.Reactivepower,
		msg.ThdvR,
		msg.ThdiR,
		msg.NinthharmonicR,
		msg.EleventhharmonicR,
		msg.ThirteenthharmonicR,
		msg.ThdvW,
		msg.ThdiW,
		msg.NinthharmonicW,
		msg.EleventhharmonicW,
		msg.ThirteenthharmonicW,
		msg.ThdvB,
		msg.ThdiB,
		msg.NinthharmonicB,
		msg.EleventhharmonicB,
		msg.ThirteenthharmonicB,
	)

}

// nsertHVAlarmMsg inserts hv alarm
func (c *TimescaleClient) InsertHVAlarmMsg(msg pb.HVAlarmMessage) error {
	sql := `insert into "HVALARM"(
		"DEVICEEUI",
		"TIMESTAMP",
		"ALARMTYPE",
		"IMBALANCE",
		"VOLTAGER",
		"VOLTAGEW",
		"VOLTAGEB",
		"VOLTAGEAB",
		"VOLTAGEBC",
		"VOLTAGECA",
		"ANGLEA",
		"ANGLEB",
		"ANGLEC"
		) values($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)`
	dt := timestampToString(int64(msg.Timesent))

	return c.dbEngine.Exec(sql,
		msg.Deviceeui,
		dt,
		msg.Alarmtype,
		msg.Imbalance,
		msg.VoltageR,
		msg.VoltageW,
		msg.VoltageB,
		msg.VoltageAb,
		msg.VoltageBc,
		msg.VoltageCa,
		msg.AngleA,
		msg.AngleB,
		msg.AngleC,
	)

}

// InsertInstMsg inserts inst msg
func (c *TimescaleClient) InsertInstMsg(msg pb.InstMessage) error {
	sql := `insert into "INST"(
		"DEVICEEUI",
		"TIMESTAMP",
		"PHASEID",
		"VOLTAGE",
		"CURRENT",
		"ACTIVEPOWER",
		"REACTIVEPOWER"
		) values($1, $2, $3, $4, $5, $6, $7)`
	dt := timestampToString(int64(msg.Timesent))

	return c.dbEngine.Exec(sql,
		msg.Deviceeui,
		dt,
		msg.Phaseid,
		msg.Voltage,
		msg.Current,
		msg.Activepower,
		msg.Reactivepower,
	)
}

// InsertInstP2PMsg insert msg
func (c *TimescaleClient) InsertInstP2PMsg(msg pb.InstP2PMessage) error {
	sql := `insert into "INSTP2P"(
		"DEVICEEUI",
		"TIMESTAMP",
		"VOLTAGER",
		"CURRENTR",
		"VOLTAGEW",
		"CURRENTW",
		"VOLTAGEB",
		"CURRENTB",
		"ACTIVEPOWER",
		"REACTIVEPOWER",
		"THDVR",
		"THDIR",
		"THDVW",
		"THDIW",
		"THDVB",
		"THDIB",
		"VOLTAGEAB",
		"VOLTAGEBC",
		"VOLTAGECA",
		"CURRENTNEUTRAL"
		) values($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20)`

	dt := timestampToString(int64(msg.Timesent))

	return c.dbEngine.Exec(sql,
		msg.Deviceeui,
		dt,
		msg.VoltageR,
		msg.CurrentR,
		msg.VoltageW,
		msg.CurrentW,
		msg.VoltageB,
		msg.CurrentB,
		msg.Activepower,
		msg.Reactivepower,
		msg.ThdvR,
		msg.ThdiR,
		msg.ThdvW,
		msg.ThdiW,
		msg.ThdvB,
		msg.ThdiB,
		msg.VoltageAb,
		msg.VoltageBc,
		msg.VoltageCa,
		msg.Currentneutral,
	)

}

// InsertPQMsg inserts pq
func (c *TimescaleClient) InsertPQMsg(msg pb.PQMessage) error {
	sql := `insert into "PQ"(
		"DEVICEEUI",
		"TIMESTAMP",
		"PHASEID",
		"VOLTAGEMAX",
		"CURRENTMAX",
		"POWERACTIVEMAX",
		"POWERREACTIVEMAX",
		"THDVMAX",
		"VOLTAGESMA",
		"CURRENTSMA",
		"POWERACTIVESMA",
		"POWERREACTIVESMA",
		"THDVSMA",
		"VOLTAGEMIN",
		"CURRENTMIN",
		"POWERACTIVEMIN",
		"POWERREACTIVEMIN",
		"THDVMIN",
		"MOMENTARYSAG",
		"MOMENTARYSWELL",
		"TEMPORARYSAG",
		"TEMPORARYSWELL",
		"SUSTAINEDUNDERVOLTAGE",
		"SUSTAINEDOVERVOLTAGE",
		"PROLONGEDUNDERVOLTAGE",
		"PROLONGEDOVERVOLTAGE",
		"THDISMA",
		"NEUTRALCURRENTSMA",
		"NEUTRALCURRENTMAX",
		"POWERAPPARENTSMA",
		"POWERAPPARENTMAX",
		"POWERAPPARENTMIN"
		) values($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32)`

	dt := timestampToString(int64(msg.Timesent))

	return c.dbEngine.Exec(sql,
		msg.Deviceeui,
		dt,
		msg.Phaseid,
		msg.Voltagemax,
		msg.Currentmax,
		msg.Poweractivemax,
		msg.Powerreactivemax,
		msg.Thdvmax,
		msg.Voltagesma,
		msg.Currentsma,
		msg.Poweractivesma,
		msg.Powerreactivesma,
		msg.Thdvsma,
		msg.Voltagemin,
		msg.Currentmin,
		msg.Poweractivemin,
		msg.Powerreactivemin,
		msg.Thdvmin,
		msg.Momentarysag,
		msg.Momentaryswell,
		msg.Temporarysag,
		msg.Temporaryswell,
		msg.Sustainedundervoltage,
		msg.Sustainedovervoltage,
		msg.Prolongedundervoltage,
		msg.Prolongedovervoltage,
		msg.Thdisma,
		msg.Neutralcurrentsma,
		msg.Neutralcurrentmax,
		msg.Powerapparentsma,
		msg.Powerapparentmax,
		msg.Powerapparentmin,
	)
}

// InsertS11PQ insert msg
func (c *TimescaleClient) InsertS11PQ(msg pb.S11PQMessage) error {
	sql := `insert into "S11PQ"(
		"DEVICEEUI",
		"TIMESTAMP",
		"PHASEID",
		"ADDRESS",
		"VOLTAGEMAX",
		"VOLTAGESMA",
		"VOLTAGEMIN",
		"CURRENTSMA",
		"THDVSMA",
		"POWERACTIVEMAX",
		"POWERREACTIVEMAX",
		"POWERACTIVESMA",
		"POWERREACTIVESMA",
		"POWERACTIVEMIN",
		"POWERREACTIVEMIN",
		"MOMENTARYSAG",
		"MOMENTARYSWELL",
		"TEMPORARYSAG",
		"TEMPORARYSWELL",
		"SUSTAINEDUNDERVOLTAGE",
		"SUSTAINEDOVERVOLTAGE",
		"PROLONGEDUNDERVOLTAGE",
		"PROLONGEDOVERVOLTAGE"
		) values($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23)`

	dt := timestampToString(int64(msg.Timesent))

	return c.dbEngine.Exec(sql,
		msg.Deviceeui,
		dt,
		msg.Phaseid,
		msg.Address,
		msg.Voltagemax,
		msg.Voltagesma,
		msg.Voltagemin,
		msg.Currentsma,
		msg.Thdvsma,
		msg.Poweractivemax,
		msg.Powerreactivemax,
		msg.Poweractivesma,
		msg.Powerreactivesma,
		msg.Poweractivemin,
		msg.Powerreactivemin,
		msg.Momentarysag,
		msg.Momentaryswell,
		msg.Temporarysag,
		msg.Temporaryswell,
		msg.Sustainedundervoltage,
		msg.Sustainedovervoltage,
		msg.Prolongedundervoltage,
		msg.Prolongedovervoltage,
	)
}

// InsertPQEventsMsg inserts msg
func (c *TimescaleClient) InsertPQEventsMsg(msg pb.PQEventsMessage) error {
	sql := `insert into "PQEVENTS"(
		"DEVICEEUI",
		"TIMESTAMP",
		"PHASEID",
		"RETAINEDVOLTAGE",
		"PQEVENTTYPE",
		"PQEVENTDURATION"
		) values($1, $2, $3, $4, $5, $6)`
	dt := timestampToString(int64(msg.Timesent))

	return c.dbEngine.Exec(sql,
		msg.Deviceeui,
		dt,
		msg.Phaseid,
		msg.Retainedvoltage,
		msg.Pqeventtype,
		msg.Pqeventduration,
	)
}

// InsertProcessedMsg inserts msg
func (c *TimescaleClient) InsertProcessedMsg(msg pb.ProcessedMessage) error {
	sql := `insert into "PROCESSED"(
		"DEVICEEUI",
		"TIMESTAMP",
		"PHASEID",
		"LOOPIMPEDANCE",
		"RSSI",
		"SNR",
		"ALIVECNT",
		"SAIFI",
		"MEMFREE",
		"TEMPCNT",
		"SUSTAINEDCNT",
		"PROLONGEDCNT",
		"MDI",
		"MDITIMESTAMP",
		"VOLTAGEUNBALANCEFACTOR"
		) values($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)`
	dt := timestampToString(int64(msg.Timesent))

	return c.dbEngine.Exec(sql,
		msg.Deviceeui,
		dt,
		msg.Phaseid,
		msg.Loopimpedance,
		msg.Rssi,
		msg.Snr,
		msg.Alivecnt,
		msg.Saifi,
		msg.Memfree,
		msg.Tempcnt,
		msg.Sustainedcnt,
		msg.Prolongedcnt,
		msg.Mdi,
		msg.Mditimestamp,
		msg.Voltageunbalancefactor,
	)

}

// InsertResendResponseMsg inserts msg
func (c *TimescaleClient) InsertResendResponseMsg(msg pb.ResendResponseMessage) error {
	sql := `insert into "RESEND_RESPONSE"(
		"DEVICEEUI",
		"TIMESTAMP",
		"MESSAGEID",
		"MESSAGETYPE",
		"SPARE",
		"STATE",
		"ERROR"
		) values($1, $2, $3, $4, $5, $6, $7)`
	dt := timestampToString(int64(msg.Timesent))

	return c.dbEngine.Exec(sql,
		msg.Deviceeui,
		dt,
		msg.Messageid,
		msg.Messagetype,
		msg.Spare,
		msg.State,
		msg.Error,
	)

}

// InsertUplinkMsg inserts uplink
func (c *TimescaleClient) InsertUplinkMsg(msg pb.UplinkMessage) error {
	sql := `insert into "UPLINK"(
		"DEVICEEUI",
		"TIMESTAMP",
		"RSSI",
		"SNR",
		"FREQUENCY",
		"DR",
		"FCNT",
		"RAWDATA",
		"MESSAGEID",
		"MESSAGETYPE",
		"RESENT"
		) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`
	dt := timestampToString(int64(msg.Timesent))

	return c.dbEngine.Exec(sql,
		msg.Deviceeui,
		dt,
		msg.Rssi,
		msg.Snr,
		msg.Frequency,
		msg.Dr,
		msg.Fctn,
		msg.Rawdata,
		msg.Messageid,
		msg.Messagetype,
		msg.Resent,
	)
}

// InsertVoltageStats insert vs
func (c *TimescaleClient) InsertVoltageStats(msg pb.VoltageStatsMessage) error {
	sql := `insert into "VOLTAGE_STATS"(
		"DEVICEEUI",
		"TIMESTAMP",
		"PHASEID",
		"MEAN",
		"VARIANCE",
		"STARTTIME",
		"STOPTIME",
		"H0_213",
		"H213_215",
		"H215_217",
		"H217_219",
		"H219_221",
		"H221_223",
		"H223_225",
		"H225_227",
		"H227_229",
		"H229_231",
		"H231_233",
		"H233_235",
		"H235_237",
		"H237_239",
		"H239_241",
		"H241_243",
		"H243_245",
		"H245_247",
		"H247_249",
		"H249_300"
		) values($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27)`

	dt := timestampToString(int64(msg.Timesent))
	return c.dbEngine.Exec(sql,
		msg.Deviceeui,
		dt,
		msg.Phaseid,
		msg.Mean,
		msg.Variance,
		msg.Starttime,
		msg.Stoptime,
		msg.H0_213,
		msg.H213_215,
		msg.H215_217,
		msg.H217_219,
		msg.H219_221,
		msg.H221_223,
		msg.H223_225,
		msg.H225_227,
		msg.H227_229,
		msg.H229_231,
		msg.H231_233,
		msg.H233_235,
		msg.H235_237,
		msg.H237_239,
		msg.H239_241,
		msg.H241_243,
		msg.H243_245,
		msg.H245_247,
		msg.H247_249,
		msg.H249_300,
	)

}

// InsertMeterStatus inserts meter status
func (c *TimescaleClient) InsertMeterStatus(msg pb.MeterStatusMessage) error {
	sql := `insert into "METER_STATUS" (
		"DEVICEEUI",
		"TIMESTAMP",
		"MODBUSADDRESS",
		"DIGITALINPUTS",
		"DIGITALOUTPUTS",
		"ALARMSTATE1",
		"ALARMSTATE2",
		"COMMSSTATE1",
		"COMMSSTATE2",
		"SPARE1"
		) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`

	dt := timestampToString(int64(msg.Timesent))
	return c.dbEngine.Exec(sql,
		msg.Deviceeui,
		dt,
		msg.Modbusaddress,
		msg.Digitalinputs,
		msg.Digitaloutputs,
		msg.Alarmstate1,
		msg.Alarmstate2,
		msg.Commsstate1,
		msg.Commsstate2,
		msg.Spare1,
	)

}
