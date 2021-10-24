package main

import (
	"time"
	"fmt"
	"context"
	"log"
	"strings"
	mqtt "github.com/eclipse/paho.mqtt.golang"
//	"go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
//"go.mongodb.org/mongo-driver/mongo/readpref"
)

func AddData(data []DataPoc, collection *mongo.Collection ) {
	for _, elem := range data{
		
		insertResult, err := collection.InsertOne(context.TODO(), elem)
		if err != nil{
			log.Fatal(err)
		}
	
		fmt.Printf("value:%v", insertResult)
	}
}


var massagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Сообщение получено: %s , из топика: %s\n", msg.Payload(), msg.Topic())
	var data1 = DataPoc{name:strings.ReplaceAll(msg.Topic(),"service/weather_logger/",""), value:string(msg.Payload()), time:time.Now().Unix()}
	packet = append(packet, data1)
	
}
var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
	if len(packet) != 0{
	 	addr := "158.101.195.184"
	 	portm := 27017
	 	dbName := "hackaton"
	 	collectionName := "weather_log"	
		client, err := mongo.NewClient(options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%d",addr,portm)))
		if err != nil{
			log.Fatal(err)
		}
		err = client.Connect(context.TODO())
		if err != nil{
			log.Fatal(err)
		}
		err = client.Ping(context.TODO(), nil)
		if err != nil{
			log.Fatal(err)
		}
		fmt.Println("Connect to Mongo db is sucsses!")
	
		collection := client.Database(dbName).Collection(collectionName)
	
		AddData(packet, collection)

		packet = packet[:0]
		fmt.Println("packet is clean")
		defer client.Disconnect(context.TODO())
	}


}
var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connection lost error - %v\n", err.Error())
}


type DataPoc struct{
	name string
	value string
	time int64
}
var packet []DataPoc


func main(){
//data for db MongoDb

//data for mqtt connect

	var broker = "mqtt0.bast-dev.ru"
	var port = 1883
	var topicPrefix = "service/weather_logger/#"
	var userName = "hackathon"
	var password = "Autumn2021"
	

	
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d/", broker, port))
	opts.SetClientID(fmt.Sprintf("data-set%d", time.Now().Unix()))
	opts.SetUsername(userName)
	opts.SetPassword(password)
	opts.SetKeepAlive(time.Second * 30)
	opts.SetDefaultPublishHandler(massagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	clientMqtt := mqtt.NewClient(opts)
	
	if token := clientMqtt.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	fmt.Println("Connecting...")
	sub(clientMqtt, topicPrefix)
	for {
	publish(clientMqtt, topicPrefix)
    time.Sleep(time.Second * 30)
	}
	defer clientMqtt.Disconnect(100)

}
func publish(client mqtt.Client, topic string) {
	num := 16
	for true{
		text := fmt.Sprintf("Message %d", num)
		token := client.Publish(topic, 0, false, text)
		token.Wait()
		time.Sleep(2*time.Second)
	}
}

func sub(client mqtt.Client, topic string) {
	token := client.Subscribe(topic, 0, nil)
	token.Wait()
	fmt.Printf("Subscribed to topic: %s\n", topic)
}
