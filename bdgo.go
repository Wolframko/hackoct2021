package dbgo

import (
	"fmt"
	"context"
	"log"
	"go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "go.mongodb.org/mongo-driver/mongo/readpref"
)

//тестовая структура
type Test struct{
	int Age
	string Name
	string City
}
func main(){
	addr := "158.101.195.184"
	port := 2707
	client, err := mongo.NewClient(options.Client()).ApplyURI(fmt.Sprintf("mongodb://%s:%d",addr,port))
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
	
