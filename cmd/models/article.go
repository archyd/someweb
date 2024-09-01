package models

import (
	"time"
)

type Article struct {
	Id         int
	PostedDate time.Time
	Title      string
	Content    string
}
