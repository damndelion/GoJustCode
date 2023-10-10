package main

import (
	"awesomeProject/Lecture8/entity"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

const AuthorizationToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"

var users []entity.User

func main() {
	user := entity.User{
		Id:       1,
		Name:     "qwerty",
		Email:    "mail@gmail.com",
		Age:      20,
		Password: "12345678",
	}
	user2 := entity.User{
		Id:       2,
		Name:     "qwerty",
		Email:    "mail@gmail.com",
		Age:      20,
		Password: "12345678",
	}
	addUser(&user)
	addUser(&user2)
	r := gin.New()

	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	v1 := r.Group("/v1")
	{
		v1.Use(TokenVerify())
		v1.GET("users", GetUsers)
		v1.POST("user", AddUser)
		v1.DELETE("user/:id", DeleteUser)

	}

	v2 := r.Group("/v2")
	{
		v2.GET("users", GetUsers)
		v2.POST("user", AddUser)
		v2.DELETE("user/:id", DeleteUser)

	}

	err := r.Run(":8080")
	if err != nil {
		return
	}
}

func addUser(user *entity.User) {
	users = append(users, *user)
}
func GetUsers(context *gin.Context) {
	context.JSON(http.StatusOK, users)

}

func AddUser(context *gin.Context) {
	body := entity.User{}

	if err := context.BindJSON(&body); err != nil {
		log.Printf("validate err: %v", err)
		context.AbortWithError(http.StatusBadRequest, err)
		return
	}

	fmt.Println(body)
	users = append(users, body)
	context.JSON(http.StatusAccepted, &body)
}

func DeleteUser(context *gin.Context) {
	id := context.Param("id")
	for i, user := range users {
		if id == strconv.Itoa(user.Id) {
			users = append(users[:i], users[i+1:]...)
			context.JSON(http.StatusNoContent, gin.H{"message": "User deleted"})
			return
		}
	}
	context.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
}

func TokenVerify() gin.HandlerFunc {
	return func(context *gin.Context) {
		token := context.Request.Header.Get("Authorization")
		if token != AuthorizationToken {
			context.JSON(http.StatusUnauthorized, gin.H{"message": "User not authorized"})
			context.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		context.Next()
	}
}
