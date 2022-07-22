package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.GET("/", roothandler)
	r.GET("/hello", hellohandler)
	r.GET("/character/:name", characterhandler)

	r.Run()

}
func roothandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Name": "Calvin Jefen",
		"bio":  "A Software Engineer",
	})
}
func hellohandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"content":  "hello world",
		"subtitle": "belajar golang api",
	})
}

func characterhandler(c *gin.Context) {
	name := c.Param("name")
	c.JSON(http.StatusOK, gin.H{
		"character name": name,
	})
}
