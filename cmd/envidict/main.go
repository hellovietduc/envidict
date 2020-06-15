package main

import (
	"flag"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/vietduc01100001/envidict/internal"
)

const (
	defaultServerAddr = ":5000"
	defaultClientAddr = ":3000"
	staticDir         = "/etc/envidict/static"
	dataFile          = "/etc/envidict/en-vi-dict.txt"
)

var flagClientAddr = flag.String("address", defaultClientAddr, "client address")
var flagStaticDir = flag.String("static", staticDir, "static directory")
var flagDataFile = flag.String("data", dataFile, "data file")
var db *internal.AVLTree

func main() {
	flag.Parse()
	go initDB()
	go serveClient()
	startServer()
}

func initDB() {
	db = &internal.AVLTree{}
	start := time.Now()
	internal.ParseFile(*flagDataFile, db)
	elapsed := time.Since(start)
	log.Printf("Build DB took %s", elapsed)
}

func serveClient() {
	fs := http.FileServer(http.Dir(*flagStaticDir))
	http.Handle("/", fs)

	log.Printf("App running on %s", *flagClientAddr)
	err := http.ListenAndServe(*flagClientAddr, nil)
	if err != nil {
		log.Fatal(err)
	}
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
			limit = 10
		}

		nodes := db.FuzzySearch(searchTerm, limit)
		if len(nodes) == 0 {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "Not Found",
			})
			return
		}

		response := make([]*internal.Word, limit)
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

	r.Run(defaultServerAddr)
}
