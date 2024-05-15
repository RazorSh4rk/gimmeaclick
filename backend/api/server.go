package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"razorsh4rk.github.io/urlshorten/db"
)

var server = gin.Default()

func Setup() {
	server.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "PUT"},
	}))

	server.GET("/ping", getPing)
	server.GET("/g", getValue)
	server.PUT("/p", putValue)
	server.GET("/all", getAll)
}

func Run() {
	server.Run(":9001")
}

func getPing(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func getValue(c *gin.Context) {
	key := c.Query("k")
	fmt.Println(key)

	if key == "" {
		c.JSON(401, gin.H{
			"message": "NOKEY",
		})
		return
	}

	value, err := db.Get(key)

	if err != nil {
		c.JSON(404, gin.H{
			"message": "NOEXIST",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": value,
	})
}

func putValue(c *gin.Context) {
	value := c.Query("v")

	if value == "" {
		c.JSON(401, gin.H{
			"message": "NOVALUE",
		})
		return
	}

	if !strings.HasPrefix(value, "www.") {
		value = "www." + value
	}

	if !strings.HasPrefix(value, "http") {
		value = "http://" + value
	}

	var id string

	res, err := http.Get("https://random-word-api.herokuapp.com/word")
	if err != nil {
		fmt.Println(err)
		id = uuid.New().String()
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		id = uuid.New().String()
	}

	var words []string
	if err = json.Unmarshal(body, &words); err != nil {
		fmt.Println(err)
		id = uuid.New().String()
	}

	if id == "" {
		id = words[0]
	}

	if err = db.Put(id, value); err != nil {
		fmt.Println(err)
		c.JSON(500, gin.H{
			"message": "SAVEERR",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": id,
	})
}

func getAll(c *gin.Context) {
	result, err := db.GetAll()

	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, result)
}
