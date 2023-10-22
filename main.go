package main

import (

	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func printJSONStructure(data interface{}, prefix string) {
	switch v := data.(type) {
	case map[string]interface{}:
		for key, value := range v {
			fmt.Printf("%s%s (object)\n", prefix, key)
			printJSONStructure(value, prefix+"  ")
		}
	case []interface{}:
		for i, item := range v {
			fmt.Printf("%s[%d] (array)\n", prefix, i)
			printJSONStructure(item, prefix+"  ")
		}
	default:
		fmt.Printf("%s(%T)\n", prefix, v)
	}
}

func main() {

	r := gin.Default()

	r.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "hello world"})
	})

	log.Fatal(r.Run())

}
