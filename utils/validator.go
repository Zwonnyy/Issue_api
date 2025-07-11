// utils/validator.go
package utils

import (
	"issueapi/data"
	"issueapi/models"
)

var ValidStatuses = map[string]bool{
	"PENDING":     true,
	"IN_PROGRESS": true,
	"COMPLETED":   true,
	"CANCELLED":   true,
}

func ValidateStatus(status string) bool {
	return ValidStatuses[status]
}

func FindUserByID(id uint) *models.User {
	for _, u := range data.Users {
		if u.ID == id {
			return &u
		}
	}
	return nil
}
