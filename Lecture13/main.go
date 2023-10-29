package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.Static("/static", "Lecture13/files")
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
