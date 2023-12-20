package test

import (
	"fmt"
	"testing"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/lailailee/toolbox"
	"go.uber.org/zap"
)

var _log *zap.SugaredLogger

func checkPort(ip string, port string, t time.Duration) {
	for {
		if e0 := toolbox.CheckTcpPort(ip, port); e0 != nil {
			_log.Warnf("probe local mqtt broker error, [%v]", e0)
		} else {
			_log.Infof("probe local mqtt broker ok [%v:%v]", ip, port)
			break
		}
		time.Sleep(t)
	}
}

func TestFile_NewMqttClient(t *testing.T) {
	t.Log("TestFile_LoadIniFile")
	var loggerConfig = toolbox.LoggerConfig{
		Filename:   "./mqtt.log",
		MaxSize:    5,
		MaxBackups: 3,
		MaxAge:     3,
		Compress:   false,
		LocalTime:  true,
	}
	logger := toolbox.InitLogger(loggerConfig, toolbox.DebugLevel)
	logger.Infoln("111")
	_log = logger

	checkPort("192.168.8.236", "1883", time.Second*5)
	var subscribe = []string{
		"/",
	}
	handler := MppHandler{
		log: logger.With("module", "local_mqtt"),
	}
	var mqttConfig = toolbox.MqttConfig{
		Address:     "tcp://192.168.8.236:1883",
		PassWord:    "hahaha",
		ClientID:    "master" + fmt.Sprintf("%v", time.Now().Unix()),
		Logger:      logger,
		Description: "mqtt_test",
		Subscribe:   subscribe,
		Handler:     &handler,
	}
	client := toolbox.NewMqttClient(mqttConfig)
	client.Connect()
	client.Run()
}

// MppHandler mqtt client handler
type MppHandler struct {
	log *zap.SugaredLogger
}

// ProcessPublish handles publish packet
func (h *MppHandler) ProcessPublish(pkt mqtt.Message) (e error) {
	switch {
	default:
		e = h.processUnkownTopic(pkt)
		break
	}
	return e
}

func (h *MppHandler) processResponse(pkt mqtt.Message, cmd int) (e error) {
	var msg toolbox.Message
	msg.Cmd = cmd
	msg.Data = pkt
	return e
}

//	func (h *MppHandler) processCommand(pkt mqtt.Message) (e error) {
//		var msg kernel.Message
//		msg.Cmd = event.EventCommand
//		msg.Data = pkt
//		kernel.ReciveMsgCh <- msg
//		return e
//	}
func (h *MppHandler) processUnkownTopic(pkt mqtt.Message) (e error) {
	// var msg kernel.Message
	// msg.Cmd = kernel.RemotePassthroughMsg
	// msg.Data = pkt
	// kernel.RemoteMgMsgCh <- &msg
	h.log.Warnf("unkown topic[%v]", pkt.Topic())
	return e
}
