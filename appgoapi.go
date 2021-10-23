package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	)


type Test struct {
    ID     int  `json:"id"`
    Title  string  `json:"title"`
    Value  string  `json:"value"`
    Time   string `json:"time"`
}
var testsData = []Test{ 
					 {ID:1, Title:"temp air", Value:"24.1", Time:"1024"},
					 {ID:2, Title:"temp grass", Value:"27.1", Time:"2048"},
					 {ID:3, Title:"light", Value:"30", Time:"4056"},
				}
//get json format Data
func GetData(c *gin.Context){
	c.IndentedJSON(http.StatusOK, testsData)
}
func main(){
	router := gin.Default()
	router.GET("testsData/", GetData)
	router.Run("localhost:8080")
	}

