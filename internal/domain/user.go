package domain

import "time"

type User struct {
	id           int
	phone_number string
	password     string
	role         string
	created_at   time.Time
	updated_at   time.Time
	deleted_at   *time.Time
}

func (u User) Id() int {
	return u.id
}

func (u User) Phone_number() string {
	return u.phone_number
}

func (u User) Password() string {
	return u.password
}

func (u User) Role() string {
	return u.role
}

func (u User) Created_at() time.Time {
	return u.created_at
}

func (u User) Updated_at() time.Time {
	return u.updated_at
}

func (u User) Deleted_at() *time.Time {
	return u.deleted_at
}

type NewUser struct {
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}
