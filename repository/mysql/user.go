package mysql

import (
	"database/sql"
	"fmt"
	"go_cast/S11P01-game/entity"
)

func (d *MySQLDB) IsPhoneNumberUnique(phoneNumber string) (bool, error) {
	var count int
	err := d.db.QueryRow("SELECT COUNT(*) FROM users WHERE phone_number = ?", phoneNumber).Scan(&count)
	if err != nil {
		if err == sql.ErrNoRows {
			return true, nil
		}
		return false, fmt.Errorf("failed to check phone number uniqueness: %w", err)
	}

	return count == 0, nil
}

func (d *MySQLDB) Register(user entity.User) (entity.User, error) {
	res, err := d.db.Exec("INSERT INTO users (name, phone_number, password) VALUES (?, ?, ?)", user.Name, user.PhoneNumber, user.Password)
	if err != nil {
		return entity.User{}, fmt.Errorf("can not insert user: %w", err)
	}

	// error is always nil
	id, _ := res.LastInsertId()
	user.ID = uint(id)

	return user, nil
}

func (d *MySQLDB) GetUserByPhoneNumber(phoneNumber string) (entity.User, bool, error) {
	var user entity.User
	err := d.db.QueryRow("SELECT id, name, phone_number, password FROM users WHERE phone_number = ?", phoneNumber).Scan(&user.ID, &user.Name, &user.PhoneNumber, &user.Password)
	if err != nil {
		return entity.User{}, false, fmt.Errorf("can not get user by phone number: %w", err)
	}
	return user, true, nil
}
