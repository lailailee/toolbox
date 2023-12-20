package toolbox

import (
	"time"

	"go.uber.org/zap"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type MqttHandler interface {
	ProcessPublish(msg mqtt.Message) (e error)
}

type Message struct {
	Cmd      int
	Data     interface{}
	ModuleID int
}

const (
	LocalSubMsg = iota + 1
	LocalSysExit
	LocalRelayRsp
)

// LClient 的配置
type LClient struct {
	Client     mqtt.Client
	Opts       *mqtt.ClientOptions
	log        *zap.SugaredLogger
	subscribe  []string
	hander     MqttHandler
	LocalMsg   chan Message
	MqttConfig MqttConfig
}

type MqttConfig struct {
	Address     string
	UserName    string
	PassWord    string
	ClientID    string
	Logger      *zap.SugaredLogger
	Description string
	Subscribe   []string
	Handler     MqttHandler
}

// NewLocalClient create local client
func NewMqttClient(cfg MqttConfig) *LClient {
	local := &LClient{}
	local.Opts = mqtt.NewClientOptions()
	local.subscribe = cfg.Subscribe
	local.hander = cfg.Handler
	local.LocalMsg = make(chan Message, 1000)
	local.log = cfg.Logger
	local.Opts.Password = cfg.PassWord
	local.Opts.Username = cfg.UserName
	local.Opts.ClientID = cfg.ClientID
	local.Opts.AddBroker(cfg.Address)
	local.MqttConfig = cfg
	return local
}

// Connect connect
func (local *LClient) Connect() {
	var localmessageSubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
		mmsg := Message{
			Cmd:  LocalSubMsg,
			Data: msg,
		}
		local.LocalMsg <- mmsg
	}
	local.Opts.SetDefaultPublishHandler(localmessageSubHandler) // 如果不是默认类型的数据就执行回调函数
	local.Opts.SetOnConnectHandler(func(c mqtt.Client) {
		for _, v := range local.subscribe {
			token := c.Subscribe(v, 0, localmessageSubHandler)
			if token.Wait() && token.Error() != nil {
				local.log.Errorf("mqtt Connect Error: %v", token.Error())
			}
		}
	})
	// 设置连接超时
	local.Opts.SetConnectTimeout(time.Duration(60) * time.Second)
	// 创建客户端连接
	local.Client = mqtt.NewClient(local.Opts)
	token := local.Client.Connect()
	if token.Wait() && token.Error() != nil {
		local.log.Infof("mqtt connect error: %v", token.Error())
	} else {
		local.log.Infof("connetct ok")
	}
}

func (local *LClient) Publish(topic string, qos byte, retained bool, payload interface{}) error {
	token := local.Client.Publish(topic, qos, retained, payload)
	if token.Wait() && token.Error() != nil {
		local.log.Errorf("send [%v][%v] error", topic, token.Error())
		return token.Error()
	}
	return nil
}

// Run
func (local *LClient) Run() {
	local.log.Debugf("-------------- %v mqtt client run -----------", local.MqttConfig.Description)
	for {
		select {
		case msg := <-local.LocalMsg:
			local.cmnHandler(&msg)
			break
		}
	}
}

func (local *LClient) cmnHandler(msg *Message) {
	switch msg.Cmd {
	case LocalSubMsg:
		{
			mmsg := msg.Data.(mqtt.Message)
			local.hander.ProcessPublish(mmsg)
			break
		}
	case LocalSysExit:
		{
			break
		}
	case LocalRelayRsp:
		{
			break
		}
	}
}
