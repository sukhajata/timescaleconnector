package messagedecoder

import (
	"github.com/golang/mock/gomock"
	"github.com/sirupsen/logrus"
	"go.elastic.co/ecslogrus"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/sukhajata/devicetwin/pkg/ppmqtt"
	"github.com/sukhajata/ppmessage/ppuplink"
	"github.com/sukhajata/timescaleconnector/mocks"
	"google.golang.org/protobuf/proto"
)

func setup(mockCtrl *gomock.Controller) (*Decoder, chan ppmqtt.Message, chan interface{}, *logrus.Logger) {
	mockHelper := mocks.NewMockAuthLoggerHelper(mockCtrl)
	inChan := make(chan ppmqtt.Message)
	outChan := make(chan interface{}, 3)
	decoder := NewDecoder(mockHelper)
	// elastic logging
	log := logrus.New()
	log.SetFormatter(&ecslogrus.Formatter{
		DataKey: "labels",
	})
	log.SetReportCaller(true)

	return decoder, inChan, outChan, log
}

func TestDecoder_ProcessMessage_Alarm(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	decoder, inChan, outChan, log := setup(mockCtrl)

	// setup
	alarm := ppuplink.AlarmMessage{
		Alarmtype: "highvoltagealarm",
		Value:     255.3,
	}
	data, err := proto.Marshal(&alarm)
	require.Nil(t, err)

	msg := ppmqtt.Message{
		Topic:   "application/powerpilot/uplink/alarm/123",
		Payload: data,
	}

	// call
	decoder.ProcessMessages(inChan, outChan, log)

	// send message
	inChan <- msg

	// assert
	select {
	case decoded := <-outChan:
		a, ok := decoded.(ppuplink.AlarmMessage)
		require.True(t, ok)
		require.Equal(t, alarm.Alarmtype, a.Alarmtype)
		require.Equal(t, alarm.Value, a.Value)
	case <-time.After(time.Millisecond * 100):
		t.Fail()
	}
}

func TestDecoder_ProcessMessage_Energy(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	decoder, inChan, outChan, log := setup(mockCtrl)

	energy := ppuplink.EnergyMessage{
		Energyexportreactive: 101.12,
		Energyimportreactive: 202.23,
		Energyexportreal:     303.45,
		Energyimportreal:     404.32,
	}

	data, err := proto.Marshal(&energy)
	require.Nil(t, err)
	msg := ppmqtt.Message{
		Topic:   "application/powerpilot/uplink/energy/123",
		Payload: data,
	}

	// call
	decoder.ProcessMessages(inChan, outChan, log)

	// send message
	inChan <- msg

	// assert
	select {
	case decoded := <-outChan:
		e, ok := decoded.(ppuplink.EnergyMessage)
		require.True(t, ok)
		require.Equal(t, energy.Energyimportreactive, e.Energyimportreactive)
		require.Equal(t, energy.Energyexportreal, e.Energyexportreal)
	case <-time.After(time.Millisecond * 100):
		t.Fail()
	}

}

func TestDecoder_ProcessMessage_Inst(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	decoder, inChan, outChan, log := setup(mockCtrl)

	//publish inst message
	inst := ppuplink.InstMessage{
		Voltage:       245.6,
		Current:       10.3,
		Activepower:   20,
		Reactivepower: 10,
	}
	data, err := proto.Marshal(&inst)
	require.Nil(t, err)
	msg := ppmqtt.Message{
		Topic:   "application/powerpilot/uplink/inst/123",
		Payload: data,
	}

	// call
	decoder.ProcessMessages(inChan, outChan, log)

	// send message
	inChan <- msg

	// assert
	select {
	case decoded := <-outChan:
		i, ok := decoded.(ppuplink.InstMessage)
		require.True(t, ok, "Message should be type inst")
		require.Equal(t, inst.Voltage, i.Voltage)
		require.Equal(t, inst.Current, i.Current)
	case <-time.After(time.Millisecond * 100):
		t.Fail()
	}

}

func TestDecoder_ProcessMessage_Processed(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	decoder, inChan, outChan, log := setup(mockCtrl)

	processed := ppuplink.ProcessedMessage{
		Voltageunbalancefactor: 2.3,
		Deviceeui:              "123",
		Loopimpedance:          23.4,
		Snr:                    5,
	}
	data, err := proto.Marshal(&processed)
	require.Nil(t, err)
	msg := ppmqtt.Message{
		Topic:   "application/powerpilot/uplink/processed/123",
		Payload: data,
	}

	// call
	decoder.ProcessMessages(inChan, outChan, log)

	// send message
	inChan <- msg

	// assert
	select {
	case decoded := <-outChan:
		i, ok := decoded.(ppuplink.ProcessedMessage)
		require.True(t, ok, "Message should be type processed")
		require.Equal(t, processed.Voltageunbalancefactor, i.Voltageunbalancefactor)
		require.Equal(t, processed.Loopimpedance, i.Loopimpedance)
	case <-time.After(time.Millisecond * 100):
		t.Fail()
	}

}
