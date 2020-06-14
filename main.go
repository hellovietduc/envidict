package main

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const (
	dataFilePath    = "./en-vi-dict.txt"
	serverAddr      = ":8080"
	suggestionLimit = 10
)

var db *AVLTree

func main() {
	initDB()
	startServer()
}

func initDB() {
	db = &AVLTree{}
	start := time.Now()
	ParseFile(db)
	elapsed := time.Since(start)
	log.Printf("Build DB took %s", elapsed)
}

func startServer() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET"},
	}))

	r.GET("/status", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
	})

	r.GET("/search", func(c *gin.Context) {
		searchTerm := c.Query("t")
		limitRequested := c.Query("l")
		exactMatch := c.Query("e")

		node := db.Search(searchTerm)
		if exactMatch == "true" && node != nil {
			c.JSON(http.StatusOK, gin.H{
				"match": node.Value,
			})
			return
		}

		limit, err := strconv.Atoi(limitRequested)
		if err != nil {
			limit = suggestionLimit
		}

		nodes := db.FuzzySearch(searchTerm, limit)
		if len(nodes) == 0 {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "Not Found",
			})
			return
		}

		response := make([]*Word, limit)
		for i, node := range nodes {
			response[i] = node.Value
		}

		if node == nil {
			c.JSON(http.StatusOK, gin.H{
				"related": response,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"match":   node.Value,
				"related": response,
			})
		}
	})

	r.Run(serverAddr)
}
