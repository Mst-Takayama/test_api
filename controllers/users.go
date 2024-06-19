package controllers

import (
	"net/http"
	"strconv"
	"test_api/models"
	"test_api/services"

	"github.com/gin-gonic/gin"
)

// UserResponse はAPIレスポンスの構造体です
type UserResponse struct {
	Status  int           `json:"status"`
	Message string        `json:"message"`
	Result  []models.User `json:"result"`
}

// GetUsers は、ユーザー一覧を取得するハンドラ関数です
func GetUsers(c *gin.Context) {
	// クエリパラメータを取得
	email := c.Query("email")
	idStr := c.Query("id")
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")

	id, err := strconv.Atoi(idStr)
	if err != nil || id < 0 {
		id = 0
	}

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		limit = 10
	}

	users, err := services.GetUsers(email, id, page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := UserResponse{
		Status:  http.StatusOK,
		Message: "ユーザー一覧を取得しました",
		Result:  users,
	}

	c.JSON(http.StatusOK, response)
}
