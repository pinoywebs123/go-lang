package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type todo struct {
	Date string
}

var todos = []todo{
	{Date: time.Now().Format("02")},
	{Date: time.Now().AddDate(0, 0, +1).Format("02")},
	{Date: time.Now().AddDate(0, 0, +2).Format("02")},
	{Date: time.Now().AddDate(0, 0, +3).Format("02")},
	{Date: time.Now().AddDate(0, 0, +4).Format("02")},
	{Date: time.Now().AddDate(0, 0, +5).Format("02")},
	{Date: time.Now().AddDate(0, 0, +6).Format("02")},
}

func getTodos(context *gin.Context) {

	//context.IndentedJSON(http.StatusOK, todos)

	for _, element := range todos {
		intVar, err := strconv.Atoi(element.Date)
		if intVar%2 != 0 {
			fmt.Println(intVar, err)
		}
		//fmt.Println(element)
	}
}

func main() {
	router := gin.Default()
	router.GET("/morley", getTodos)
	router.Run("localhost:9090")
}
