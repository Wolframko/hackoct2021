package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"log"
	)


type Test struct {
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Value  string  `json:"value"`
    Time   string `json:"time"`
}
var testsData = []Test{ 
					 {ID:"1", Title:"temp air", Value:"24.1", Time:"1024"},
					 {ID:"2", Title:"temp grass", Value:"27.1", Time:"2048"},
					 {ID:"3", Title:"light", Value:"30", Time:"4056"},
				}

func main(){
	router := gin.Default()
	router.GET("testsData/", GetData)
	router.POST("testsData/", PostData)
	router.GET("testsData/:id", GetDataById)

	router.Run("localhost:8080")
	}



//get json format Data
func GetData(c *gin.Context){
	c.IndentedJSON(http.StatusOK, testsData)
}
//post json format Data
func PostData(c *gin.Context){
	var newDataTest= Test{ID:"4", Title:"fuel", Value:"42", Time:"434"}
	if err:= c.BindJSON(&newDataTest); err!=nil{
		log.Printf("JSON dont get, error: %v\n", err)
		return
	}
	testsData = append(testsData, newDataTest)
	c.IndentedJSON(http.StatusCreated, newDataTest)


}
func GetDataById(c *gin.Context){
	id := c.Param("id")
	
	for _,elem := range testsData{
		if elem.ID == id{
			c.IndentedJSON(http.StatusOK,elem)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message":"data not found"})
}


