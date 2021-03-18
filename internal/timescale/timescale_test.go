package timescale_test

import (
	"github.com/golang/mock/gomock"
	"github.com/sukhajata/ppmessage/ppuplink"
	"github.com/sukhajata/timescaleconnector/internal/timescale"
	"github.com/sukhajata/timescaleconnector/mocks"
	"testing"
	"time"
)

func setupTimescaleTest(mockCtrl *gomock.Controller) (*timescale.TimescaleClient, *mocks.MockSQLEngine, chan interface{}) {
	mockDBEngine := mocks.NewMockSQLEngine(mockCtrl)
	client := timescale.NewTimescaleClient(mockDBEngine)
	decodedChan := make(chan interface{}, 2)
	return client, mockDBEngine, decodedChan
}

func TestTimescaleClient_InsertAlarmMsg(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	client, mockDBEngine, decodedChan := setupTimescaleTest(mockCtrl)

	go client.Listen(decodedChan)

	alarm := ppuplink.AlarmMessage{
		Deviceeui: "1234",
		Alarmtype: "highvoltagealarm",
		Timesent:  1613693724000,
		Value:     255.3,
		Phaseid:   1,
	}

	// send mock alarm on channel
	decodedChan <- alarm

	// expect
	mockDBEngine.EXPECT().Exec(gomock.Any(), alarm.Deviceeui, gomock.Any(), alarm.Phaseid, alarm.Alarmtype, alarm.Value).Return(nil).Times(1)

	// give time to receive message
	time.Sleep(time.Millisecond * 3)

}

func TestTimescaleClient_InsertEnergyMsg(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	client, mockDBEngine, decodedChan := setupTimescaleTest(mockCtrl)

	go client.Listen(decodedChan)

	energy := ppuplink.EnergyMessage{
		Deviceeui:            "123",
		Timesent:             1613693724000,
		Energyexportreactive: 101.12,
		Energyimportreactive: 202.23,
		Energyexportreal:     303.45,
		Energyimportreal:     404.32,
		Phaseid:              1,
	}

	// send energy msg
	decodedChan <- energy

	// expect
	mockDBEngine.EXPECT().Exec(gomock.Any(), energy.Deviceeui, gomock.Any(), energy.Phaseid, energy.Energyexportreactive, energy.Energyimportreactive, energy.Energyexportreal, energy.Energyimportreal).Return(nil).Times(1)

	// give time to receive message
	time.Sleep(time.Millisecond * 3)

}

func TestTimescaleClient_InsertPQMsg(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	client, mockDBEngine, decodedChan := setupTimescaleTest(mockCtrl)

	go client.Listen(decodedChan)

	pq := ppuplink.PQMessage{
		Deviceeui:             "1234",
		Phaseid:               1,
		Timesent:              1613693724000,
		Voltagemax:            253.45,
		Currentmax:            3.67,
		Poweractivemax:        334.67,
		Powerreactivemax:      12.4,
		Thdvmax:               23.4,
		Voltagesma:            245.23,
		Currentsma:            3.56,
		Poweractivesma:        123.3,
		Powerreactivesma:      5.7,
		Thdvsma:               3.45,
		Voltagemin:            236.3,
		Currentmin:            0.2,
		Poweractivemin:        2.4,
		Powerreactivemin:      0.23,
		Thdvmin:               2.78,
		Momentarysag:          23,
		Momentaryswell:        12,
		Temporarysag:          0,
		Temporaryswell:        2,
		Sustainedundervoltage: 12,
		Sustainedovervoltage:  3,
		Prolongedundervoltage: 2,
		Prolongedovervoltage:  7,
		Thdisma:               2.34,
		Neutralcurrentsma:     5.78,
		Neutralcurrentmax:     7.98,
		Powerapparentsma:      234,
		Powerapparentmax:      354,
		Powerapparentmin:      231,
	}

	decodedChan <- pq

	// expect
	mockDBEngine.EXPECT().Exec(
		gomock.Any(),
		pq.Deviceeui,
		gomock.Any(),
		pq.Phaseid,
		pq.Voltagemax,
		pq.Currentmax,
		pq.Poweractivemax,
		pq.Powerreactivemax,
		pq.Thdvmax,
		pq.Voltagesma,
		pq.Currentsma,
		pq.Poweractivesma,
		pq.Powerreactivesma,
		pq.Thdvsma,
		pq.Voltagemin,
		pq.Currentmin,
		pq.Poweractivemin,
		pq.Powerreactivemin,
		pq.Thdvmin,
		pq.Momentarysag,
		pq.Momentaryswell,
		pq.Temporarysag,
		pq.Temporaryswell,
		pq.Sustainedundervoltage,
		pq.Sustainedovervoltage,
		pq.Prolongedundervoltage,
		pq.Prolongedovervoltage,
		pq.Thdisma,
		pq.Neutralcurrentsma,
		pq.Neutralcurrentmax,
		pq.Powerapparentsma,
		pq.Powerapparentmax,
		pq.Powerapparentmin,
	).Return(nil).Times(1)

	time.Sleep(time.Millisecond * 4)

}

func TestTimescaleClient_InsertInstMsg(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	client, mockDBEngine, decodedChan := setupTimescaleTest(mockCtrl)

	go client.Listen(decodedChan)

	inst := ppuplink.InstMessage{
		Deviceeui:     "1234",
		Phaseid:       1,
		Timesent:      1613693724000,
		Voltage:       235.67,
		Current:       3.4,
		Activepower:   342,
		Reactivepower: 3.4,
	}

	decodedChan <- inst

	// expect
	mockDBEngine.EXPECT().Exec(gomock.Any(), inst.Deviceeui, gomock.Any(), inst.Phaseid, inst.Voltage, inst.Current, inst.Activepower, inst.Reactivepower).Return(nil).Times(1)

	// give time to receive message
	time.Sleep(time.Millisecond * 3)
}
