package main

import (
	"path"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/maxproske/L44/handlers"
)

func main() {
	// Create Gin server
	r := gin.Default()

	// Routing in Gin is specific and the root path cannot be ambiguous
	// (/* would precedence over every other route in your web server)
	r.NoRoute(func(c *gin.Context) {
		// Assume that this call is asking for a file and attempt to find this file
		dir, file := path.Split(c.Request.RequestURI)
		ext := filepath.Ext(file)
		if file == "" || ext == "" {
			c.File("./ui/dist/ui/index.html")
		} else {
			c.File("./ui/dist/ui/" + path.Join(dir, file))
		}
	})

	r.GET("/todo", handlers.GetTodoListHandler)
	r.POST("/todo", handlers.AddTodoHandler)
	r.DELETE("/todo/:id", handlers.DeleteTodoHandler)
	r.PUT("/todo", handlers.CompleteTodoHandler)

	// Run web server on port 3000
	err := r.Run(":3000")
	if err != nil {
		panic(err)
	}
}
