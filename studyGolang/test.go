// You can edit this code!
// Click here and start typing.
package main

import "fmt"
import "github.com/gin-gonic/gin"
import "github.com/gorilla/mux"

func main() {
	fmt.Println("Hello, 世界")
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
