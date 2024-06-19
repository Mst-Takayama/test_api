package services

import (
	"fmt"
	"strings"
	"test_api/database"
	"test_api/models"
)

func GetUsers(email string, id, page, limit int) ([]models.User, error) {
	var conditions []string
	var args []interface{}

	if email != "" {
		conditions = append(conditions, "email LIKE $"+fmt.Sprint(len(args)+1))
		args = append(args, "%"+email+"%")
	}

	if id > 0 {
		conditions = append(conditions, "id = $"+fmt.Sprint(len(args)+1))
		args = append(args, id)
	}

	offset := (page - 1) * limit
	query := "SELECT id, email FROM users"

	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	query += fmt.Sprintf(" LIMIT %d OFFSET %d", limit, offset)

	rows, err := database.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, rows.Err()
}
