package mqttconn

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var (
	mqttClient   mqtt.Client
	mqttErr      error
	mqttTopic    string
	mqttServer   string
	mqttClientId string
)

func GetCapsuleMqttConn() (mqtt.Client, error) {
	return mqttClient, mqttErr
}

func InitMqttConn(mqttSrv, mqttClientId string, messageHandler mqtt.MessageHandler) (mqtt.Client, error) {
	if mqttClient == nil {
		opts := mqtt.NewClientOptions()
		opts.AddBroker(fmt.Sprintf("tcp://%s", mqttSrv))
		opts.SetClientID(mqttClientId)
		// TODO: add credential management
		// opts.SetUsername("username")
		// opts.SetPassword("password")
		opts.SetDefaultPublishHandler(messageHandler)
		// TODO: add connect and connectLost handlers management
		// opts.OnConnect = connectHandler
		// opts.OnConnectionLost = connectLostHandler
		mqttClient = mqtt.NewClient(opts)
		token := mqttClient.Connect()
		token.Wait()
		mqttErr = token.Error()
		return mqttClient, mqttErr
	}
	return mqttClient, mqttErr
}

func SetCapsuleMqttTopic(topic string) {
	mqttTopic = topic
}
func GetCapsuleMqttTopic() string {
	return mqttTopic
}

func SetCapsuleMqttServer(mqttSrv string) {
	mqttServer = mqttSrv
}
func GetCapsuleMqttServer() string {
	return mqttServer
}

func SetCapsuleMqttClientId(clientId string) {
	mqttClientId = clientId
}
func GetCapsuleMqttClientId() string {
	return mqttClientId
}
