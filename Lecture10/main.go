package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"log"
	"net/http"
	"time"
)

func main() {
	r := gin.New()
	v1 := r.Group("/v1")
	{
		v1.GET("/address", listAddresses)
	}
	r.Run(":8080")
}

// this function is related to final project
func listAddresses(ctx *gin.Context) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	defer redisClient.Close()

	addresses, err := redisClient.Get(context.Background(), "addresses").Result()
	if err == redis.Nil {
		//logic from final project to get all the wallet address from blockchain

		//wallets, err := blockchain.NewWallets()
		//if err != nil {
		//	log.Panic(err)
		//}
		//
		//addresses := wallets.GetAddresses()

		err = redisClient.Set(context.Background(), "address", addresses, 5*time.Minute).Err()
		if err != nil {
			log.Println("Error caching addresses:", err)
		}
	}

	ctx.JSON(http.StatusOK, addresses)
}
