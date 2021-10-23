package main

import (
	"fmt"
	"context"
	"log"
//	"go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
//"go.mongodb.org/mongo-driver/mongo/readpref"
)

func AddData(data Test, collection *mongo.Collection ) error{
	insertResult, err := collection.InsertOne(context.TODO(), data)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Printf("value:%v",insertResult)
	return err
}



//тестовая структура
type Test struct{
	name string
	time int32
}
func main(){
	addr := "158.101.195.184"
	port := 27017
	dbName := "hackaton"
	collectionName := "weather_log"
	client, err := mongo.NewClient(options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%d",addr,port)))
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
	fmt.Printf("%T\n",collection)
	
	testData1 := Test{"test1", 42}
	err = AddData(testData1, collection)

	err = client.Disconnect(context.TODO())
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("Connection with MongoDB close")

}	
