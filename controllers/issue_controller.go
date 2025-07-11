// controllers/issue_controller.go
package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"issueapi/data"
	"issueapi/models"
	"issueapi/utils"
)

type CreateIssueRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	UserID      *uint  `json:"userId"`
}

type UpdateIssueRequest struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Status      *string `json:"status"`
	UserID      *uint   `json:"userId"`
}

func CreateIssue(c *gin.Context) {
	var req CreateIssueRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "필수 파라미터 누락", "code": 400})
		return
	}

	var user *models.User
	status := "PENDING"

	if req.UserID != nil {
		user = utils.FindUserByID(*req.UserID)
		if user == nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "존재하지 않는 사용자", "code": 400})
			return
		}
		status = "IN_PROGRESS"
	}

	id := uint(len(data.Issues) + 1)
	now := time.Now()
	issue := models.Issue{
		ID:          id,
		Title:       req.Title,
		Description: req.Description,
		Status:      status,
		User:        user,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	data.Issues = append(data.Issues, issue)
	c.JSON(http.StatusCreated, issue)
}

func GetIssues(c *gin.Context) {
	status := c.Query("status")
	var result []models.Issue

	if status == "" {
		result = data.Issues
	} else if !utils.ValidateStatus(status) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "잘못된 상태값", "code": 400})
		return
	} else {
		for _, i := range data.Issues {
			if i.Status == status {
				result = append(result, i)
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{"issues": result})
}

func GetIssueByID(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	for _, i := range data.Issues {
		if int(i.ID) == id {
			c.JSON(http.StatusOK, i)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "이슈를 찾을 수 없음", "code": 404})
}

func UpdateIssue(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	for i := range data.Issues {
		if int(data.Issues[i].ID) != id {
			continue
		}

		issue := &data.Issues[i]
		if issue.Status == "COMPLETED" || issue.Status == "CANCELLED" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "완료되었거나 취소된 이슈는 수정할 수 없습니다.", "code": 400})
			return
		}

		var req UpdateIssueRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "잘못된 요청 형식", "code": 400})
			return
		}

		if req.Title != nil {
			issue.Title = *req.Title
		}
		if req.Description != nil {
			issue.Description = *req.Description
		}
		if req.UserID != nil {
			user := utils.FindUserByID(*req.UserID)
			if user == nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "존재하지 않는 사용자", "code": 400})
				return
			}
			issue.User = user
			if issue.Status == "PENDING" && req.Status == nil {
				issue.Status = "IN_PROGRESS"
			}
		} else if req.UserID == nil && req.Status != nil && *req.Status != "PENDING" && *req.Status != "CANCELLED" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "담당자 없이 상태 변경 불가", "code": 400})
			return
		}
		if req.UserID == nil && issue.User != nil {
			issue.User = nil
			issue.Status = "PENDING"
		}
		if req.Status != nil {
			if !utils.ValidateStatus(*req.Status) {
				c.JSON(http.StatusBadRequest, gin.H{"error": "잘못된 상태값", "code": 400})
				return
			}
			issue.Status = *req.Status
		}

		issue.UpdatedAt = time.Now()
		c.JSON(http.StatusOK, issue)
		return
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "이슈를 찾을 수 없음", "code": 404})
}