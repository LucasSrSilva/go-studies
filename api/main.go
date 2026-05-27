package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Pizza struct {
	ID    int
	sabor string
	preco float64
}

func main() {
	router := gin.Default()
	router.GET("/pizzas", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"pizzas": "To",
		})
	})
	var pizzas = []Pizza{
		{ID: 1, sabor: "Baiana", preco: 15},
		{ID: 2, sabor: "Palmito", preco: 20},
		{ID: 3, sabor: "Escarola", preco: 25},
	}
	fmt.Println(pizzas)
	router.Run()
}
