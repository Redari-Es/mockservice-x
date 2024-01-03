package main

import (
	"fmt"
	"time"

	"github.com/brianvoe/gofakeit"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"info": time.Now(),
		})
	})
	_ = r.Run(":3030")
}

// Generta a list of random mock data

func generateData() []gin.H {
	var dataList []gin.H
	for i := 0; i < 5; i++ {
		data := gin.H{
			"name":    gofakeit.Name(),
			"title":   gofakeit.JobTitle(),
			"country": gofakeit.Country(),
			"car":     fmt.Sprintf("%s - %s", gofakeit.CarMaker(), gofakeit.CarModel()),
			"address": gofakeit.Address().Address,
		}
		dataList = append(dataList, data)
	}
	return dataList
}
