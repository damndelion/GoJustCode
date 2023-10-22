package repo

import (
	"awesomeProject/Lecture9/Entity"
	"database/sql"
	"log"
)

type UserCredentialsRepository struct {
	db *sql.DB
}

func NewUserCredentialsRepository(db *sql.DB) *UserCredentialsRepository {
	return &UserCredentialsRepository{db}
}

func (r *UserCredentialsRepository) GetAll() []Entity.UserCredentials {
	var users []Entity.UserCredentials
	res, err := r.db.Query("SELECT * FROM user_credentials")
	if err != nil {
		log.Fatal(err)
		return nil
	}
	for res.Next() {
		var user Entity.UserCredentials
		err := res.Scan(&user.ID, &user.UserID, &user.CardNum, &user.Type, &user.CVV)
		if err != nil {
			return nil
		}
		users = append(users, user)
	}

	return users

}

func (r *UserCredentialsRepository) GetById(id int) *Entity.UserCredentials {
	var user Entity.UserCredentials
	err := r.db.QueryRow("SELECT * FROM user_credentials WHERE ID = $1", id).
		Scan(&user.ID, &user.UserID, &user.CardNum, &user.Type, &user.CVV)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return &user
}

func (r *UserCredentialsRepository) CreateUser(user Entity.UserCredentials) error {
	_, err := r.db.Exec("INSERT INTO user_credentials(user_id, card_num, type, cvv) VALUES($1, $2, $3, $4)",
		user.UserID, user.CardNum, user.Type, user.CVV)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (r *UserCredentialsRepository) UpdateUser(id int, updatedUser Entity.UserCredentials) error {
	_, err := r.db.Exec("UPDATE user_credentials SET user_id = $1, card_num = $2, type = $3, cvv = $4 WHERE ID = $5",
		updatedUser.UserID, updatedUser.CardNum, updatedUser.Type, updatedUser.CVV, id)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (r *UserCredentialsRepository) DeleteUser(id int) error {
	_, err := r.db.Exec("DELETE FROM user_credentials WHERE ID = $1", id)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
