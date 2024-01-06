package domain

import "time"

type UserInfo struct {
	Id        int
	Name      string
	Weigh     string
	Height    string
	Age       string
	Waist     string
	UpdatedAt time.Time
}
