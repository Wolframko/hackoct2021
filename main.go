package main

import(
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"time"
)

var massagePubHandler mqtt.MessageHandler= func(client mqtt.Client, msg mqtt.Message){
												fmt.Printf("Сообщение отправлено: %s , из топика: %s\n", msg.PayLoad(), msg.Topic())
												}
var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client){
												fmt.Println("Connected")
											}
var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error){
														fmt.Printf("Connection lost error - %v\n", err)
													}

	


func main(){
	var broker = "mqtt0.bast-dev.ru"
	var port = 1883
	var topic_prefix = "service/weather_logger"
	var user_name = "hackathon"
	var password = "Autumn2021"
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d"))
	opts.SetClientID("go_mqtt_client")
	opts.SetUsername(user_name)
	opts.SetPassword(password)
	opts.SetDefaultPublishHandler(massagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	client := mqtt.NewClient(opts)
	
	if token := client.Connect(); token.Wait() && token.Error != nil{
		panic(token.Error())
	}

	sub(client)
	publish(client)

}
func publish(client mqtt.Client) {
    num := 10
    for i := 0; i < num; i++ {
        text := fmt.Sprintf("Message %d", i)
        token := client.Publish(topic_prefix, 0, false, text)
        token.Wait()
        time.Sleep(time.Second)
    }
}

func sub(client mqtt.Client) {
    topic := topic_prefix
    token := client.Subscribe(topic, 1, nil)
    token.Wait()
  fmt.Printf("Subscribed to topic: %s", topic)
}

