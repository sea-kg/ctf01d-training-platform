package postgres

import (
	"avitoStart/internal/model"
	"log"
)

func (db *Database) ExtractUsers() ([]model.User, error) {
	var user model.User
	var users []model.User
	res, err := db.db.Query("SELECT id_user, name, same_info  FROM users;")
	if err != nil {
		return nil, err
	}
	for res.Next() {
		err = res.Scan(&user.Id, &user.Name, &user.Sameinfo)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil

}

func (db *Database) DeleteUser(id string) (bool, error) {
	tx, err := db.db.Beginx()
	if err != nil {
		log.Println(err)
	}
	_, err = tx.Exec("DELETE FROM slugtraker WHERE id_user = $1;", id)
	if err != nil {
		tx.Rollback()
		log.Println(err)
	}
	_, err = tx.Exec("DELETE FROM users WHERE id_user=$1", id)
	if err != nil {
		tx.Rollback()
		log.Println(err)
	}

	err = tx.Commit()
	if err != nil {
		log.Println(err)
		return false, err
	} else {
		return true, err
	}
}

func (db *Database) AddUser(user model.User) (bool, error) {
	_, err := db.db.Exec("insert into users (name, same_info ) values ($1, $2);", user.Name, user.Sameinfo)
	if err != nil {
		log.Println(err)
		return false, err
	}
	return true, err
}
