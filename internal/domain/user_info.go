package domain

import "time"

type UserInfo struct {
	Id        int
	Name      string
	Gender    string
	Weigh     string
	Height    string
	Age       int
	Waist     string
	UpdatedAt time.Time
}
