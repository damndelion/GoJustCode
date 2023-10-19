package repo

import (
	"awesomeProject/Lecture9/Entity"
	"database/sql"
	"log"
)

type UserInfoRepository struct {
	db *sql.DB
}

func NewUserInfoRepository(db *sql.DB) *UserInfoRepository {
	return &UserInfoRepository{db}
}

func (r *UserInfoRepository) GetAll() []Entity.UserInfo {
	var users []Entity.UserInfo
	res, err := r.db.Query("SELECT * FROM user_info")
	if err != nil {
		log.Fatal(err)
		return nil
	}
	for res.Next() {
		var user Entity.UserInfo
		err := res.Scan(&user.ID, &user.UserID, &user.Age, &user.Address, &user.Phone, &user.Country, &user.City)
		if err != nil {
			return nil
		}
		users = append(users, user)
	}

	return users

}

func (r *UserInfoRepository) GetById(id int) *Entity.UserInfo {
	var user Entity.UserInfo
	err := r.db.QueryRow("SELECT * FROM user_info WHERE ID = $1", id).
		Scan(&user.ID, &user.UserID, &user.Age, &user.Address, &user.Phone, &user.Country, &user.City)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return &user
}

func (r *UserInfoRepository) CreateUser(user Entity.UserInfo) error {
	_, err := r.db.Exec("INSERT INTO user_info(user_id, age, address, phone, country, city) VALUES($1, $2, $3, $4, $5, $6)",
		user.UserID, user.Age, user.Address, user.Phone, user.Country, user.City)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (r *UserInfoRepository) UpdateUser(id int, updatedUser Entity.UserInfo) error {
	_, err := r.db.Exec("UPDATE user_info SET user_id = $1, age = $2, phone = $3, address = $4, country = $5, city = $6 WHERE ID = $7",
		updatedUser.UserID, updatedUser.Age, updatedUser.Phone, updatedUser.Address, updatedUser.Country, updatedUser.City, id)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (r *UserInfoRepository) DeleteUser(id int) error {
	_, err := r.db.Exec("DELETE FROM user_info WHERE ID = $1", id)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
