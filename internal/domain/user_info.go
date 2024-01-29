package domain

import "time"

type UserInfo struct {
	Id        int
	Name      string
	Gender    string
	Age       int
	UpdatedAt time.Time
}
