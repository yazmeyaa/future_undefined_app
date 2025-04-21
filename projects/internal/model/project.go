package model

import "time"

type Project struct {
	ID          int64
	Title       string
	Description string
	GithubUrl   string
	ProjectUrl  string
	CreatedAt   time.Time
}
