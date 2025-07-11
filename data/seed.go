// data/seed.go
package data

import (
	"issueapi/models"
	"time"
)

var Users = []models.User{
	{ID: 1, Name: "김개발"},
	{ID: 2, Name: "이디자인"},
	{ID: 3, Name: "박기획"},
}

var Issues = []models.Issue{
	{
		ID:          1,
		Title:       "테스트 이슈",
		Description: "서버 실행 확인용 초기 이슈",
		Status:      "PENDING",
		User:        nil,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	},
}
