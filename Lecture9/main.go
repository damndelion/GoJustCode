package main

import (
	"awesomeProject/Lecture9/Entity"
	"awesomeProject/Lecture9/repo"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "user=postgres password=postgres dbname=postgres sslmode=disable")
	if err != nil {
		fmt.Println("Error opening the database connection:", err)
		return
	}

	userRepository := repo.NewUserInfoRepository(db)

	fmt.Println(userRepository.GetAll())
	fmt.Println(userRepository.GetById(1))
	user := Entity.UserInfo{UserID: 1,
		Age:     20,
		Address: "Stree 1/1",
		Phone:   "87777777777",
		Country: "Kazakhstan",
		City:    "Almaty",
	}
	userRepository.CreateUser(user)
	fmt.Println(userRepository.GetAll())

	updateUser := Entity.UserInfo{UserID: 2,
		Age:     20,
		Address: "Stree 1/1",
		Phone:   "87777777777",
		Country: "Kazakhstan",
		City:    "Astana",
	}
	userRepository.UpdateUser(2, updateUser)
	fmt.Println(userRepository.GetAll())

}
