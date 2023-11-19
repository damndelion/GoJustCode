package repo

import (
	"awesomeProject/Lecture9/Entity"
	"database/sql"
	"log"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) GetAll() []Entity.User {
	var users []Entity.User
	res, err := r.db.Query("SELECT * FROM users")
	if err != nil {
		log.Fatal(err)
		return nil
	}
	for res.Next() {
		var user Entity.User
		err := res.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.WalletHash)
		if err != nil {
			return nil
		}
		users = append(users, user)
	}

	return users

}

func (r *UserRepository) GetById(id int) *Entity.User {
	var user Entity.User
	err := r.db.QueryRow("SELECT * FROM users WHERE ID = $1", id).
		Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.WalletHash)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return &user
}

func (r *UserRepository) CreateUser(user Entity.User) error {
	_, err := r.db.Exec("INSERT INTO users(name, email, password, wallet_hash) VALUES($1, $2, $3, $4)",
		user.Name, user.Email, user.Password, user.WalletHash)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (r *UserRepository) UpdateUser(id int, updatedUser Entity.User) error {
	_, err := r.db.Exec("UPDATE users SET name = $1, email = $2, password = $3, wallet_hash = $4 WHERE ID = $5",
		updatedUser.Name, updatedUser.Email, updatedUser.Password, updatedUser.WalletHash, id)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (r *UserRepository) DeleteUser(id int) error {
	_, err := r.db.Exec("DELETE FROM users WHERE ID = $1", id)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
