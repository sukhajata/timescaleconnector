package timescale

import pb "github.com/sukhajata/ppmessage/ppuplink"

type DBClient interface {
	Listen(decodedChan <-chan interface{})
	InsertAlarmMsg(msg pb.AlarmMessage) error
	InsertCircuitEnergyMsg(msg pb.CircuitEnergyMessage) error
	InsertCircuitLoadMessage(msg pb.CircuitLoadMessage) error
	InsertEnergyMsg(msg pb.EnergyMessage) error
	InsertGatewayMsg(msg pb.GatewayMessage) error
	InsertGeoscanMsg(msg pb.GeoscanMessage) error
	InsertHarmonicsLowerMsg(msg pb.HarmonicsLowerMessage) error
	InsertHarmonicsUpperMsg(msg pb.HarmonicsUpperMessage) error
	InsertHVAlarmMsg(msg pb.HVAlarmMessage) error
	InsertInstMsg(msg pb.InstMessage) error
	InsertInstP2PMsg(msg pb.InstP2PMessage) error
	InsertPQMsg(msg pb.PQMessage) error
	InsertS11PQ(msg pb.S11PQMessage) error
	InsertPQEventsMsg(msg pb.PQEventsMessage) error
	InsertProcessedMsg(msg pb.ProcessedMessage) error
	InsertResendResponseMsg(msg pb.ResendResponseMessage) error
	InsertUplinkMsg(msg pb.UplinkMessage) error
	InsertVoltageStats(msg pb.VoltageStatsMessage) error
}
