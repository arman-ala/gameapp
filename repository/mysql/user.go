package mysql

import (
	"fmt"
	"go_cast/S11P01-game/entity"
)

func (d *MySQLDB) IsPhoneNumberUnique(phoneNumber string) (bool, error) {
	var count int
	err := d.db.QueryRow("SELECT COUNT(*) FROM users WHERE phone_number = ?", phoneNumber).Scan(&count)
	if err != nil {
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
