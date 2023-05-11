package model

import (
	"errors"
	"time"
)

const UserTable = "user"

type User struct {
	ID        int64     `gorm:"type:bigint(20) NOT NULL auto_increment;primary_key;" json:"id,omitempty"`
	Name      string    `gorm:"unique_index:name" json:"name,omitempty"`
	CreatedAt time.Time `gorm:"type:timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"type:timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP" json:"updated_at,omitempty"`
}

func CreateUser(u *User) error {
	res := UserM.Table(UserTable).Create(&u)
	return res.Error
}

func GetUser(id int64) ([]User, error) {
	var data []User
	res := UserM.Table(UserTable).
		Where("id = ?", id).
		Scan(&data)

	return data, res.Error
}

func DeleteUser(id int64) error {
	res := UserM.Table(UserTable).
		Where("id = ?", id).
		Delete(&User{})

	return res.Error
}

func UpdateUser(u *User) error {
	updateDate := make(map[string]interface{})
	if u.ID == 0 {
		return errors.New("ID 有誤")
	}

	if u.Name != "" {
		updateDate["name"] = u.Name
	}

	res := UserM.Model(&User{}).Where("id = ?", u.ID).Update(updateDate)

	return res.Error
}

func Close() error {
	if err := UserM.Close(); err != nil {
		return err
	}

	return nil
}
