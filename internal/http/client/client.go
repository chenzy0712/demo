package client

import (
	"os"
	"time"

	"git.kldmp.com/learning/demo/pkg/log"
	"github.com/beego/beego/v2/client/httplib"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func ClientDemo() {
	//login emqx
	req := httplib.Post("http://localhost:18083/api/v4/auth")

	req.Header("Content-Type", "application/json")
	req.SetBasicAuth("admin", "public")

	resp, err := req.Response()
	if err != nil {
		log.Warn("Login to emqx with err:%s.", err)
		return
	}
	log.Info("Login to emqx successful with resp:%+v", resp.Body)

	//nodes
	req = httplib.Post("http://localhost:18083/api/v4/nodes")
	req.Header("Content-Type", "application/json")
	req.SetBasicAuth("admin", "public")
	resp, err = req.Response()
	if err != nil {
		log.Warn("Get nodes of emqx with err:%s.", err)
		return
	}
	log.Info("Get nodes of emqx successful with resp:%+v", resp.Body)

	//mqtt test
	opts := mqtt.NewClientOptions().AddBroker("tcp://localhost:1883").SetUsername("emqx_u").SetPassword("emqx_p")
	opts.SetKeepAlive(60 * time.Second)
	// 设置消息回调处理函数
	opts.SetDefaultPublishHandler(f)
	opts.SetPingTimeout(1 * time.Second)
	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	// 订阅主题
	if token := c.Subscribe("testtopic/#", 0, nil); token.Wait() && token.Error() != nil {
		log.Info("%s", token.Error())
		os.Exit(1)
	}

	// 发布消息
	token := c.Publish("testtopic/1", 0, false, "Hello World")
	token.Wait()

	time.Sleep(6 * time.Second)
	// 取消订阅
	if token := c.Unsubscribe("testtopic/#"); token.Wait() && token.Error() != nil {
		log.Info("%s", token.Error())
		os.Exit(1)
	}

	// 断开连接
	c.Disconnect(250)
	time.Sleep(1 * time.Second)

	log.Info("Publish of emqx successful with resp:%+v", resp.Body)

	return
}

var f mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	log.Info("TOPIC: %s", msg.Topic())
	log.Info("MSG: %s", msg.Payload())
}
