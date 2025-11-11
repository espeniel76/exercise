package routes

import (
	"echo-gorm/config"
	"echo-gorm/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// SetupUserRoutes sets up all user-related routes
func SetupUserRoutes(e *echo.Echo) {
	// CREATE - 사용자 생성
	e.POST("/api/users", createUser)

	// READ - 모든 사용자 조회
	e.GET("/api/users", getUsers)

	// READ - 특정 사용자 조회
	e.GET("/api/users/:id", getUser)

	// UPDATE - 사용자 수정
	e.PUT("/api/users/:id", updateUser)

	// DELETE - 사용자 삭제
	e.DELETE("/api/users/:id", deleteUser)
}

// createUser creates a new user
func createUser(c echo.Context) error {
	var user models.User

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
	}

	if err := config.DB.Create(&user).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"success": true,
		"data":    user,
	})
}

// getUsers retrieves all users
func getUsers(c echo.Context) error {
	var users []models.User

	if err := config.DB.Order("id ASC").Find(&users).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"data":    users,
	})
}

// getUser retrieves a specific user by ID
func getUser(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "Invalid user ID",
		})
	}

	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"success": false,
			"error":   "User not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"data":    user,
	})
}

// updateUser updates an existing user
func updateUser(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "Invalid user ID",
		})
	}

	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"success": false,
			"error":   "User not found",
		})
	}

	var updateData models.User
	if err := c.Bind(&updateData); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
	}

	// Update user fields
	if err := config.DB.Model(&user).Updates(map[string]interface{}{
		"email":    updateData.Email,
		"username": updateData.Username,
		"password": updateData.Password,
	}).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
	}

	// Fetch updated user
	config.DB.First(&user, id)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"data":    user,
	})
}

// deleteUser deletes a user
func deleteUser(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "Invalid user ID",
		})
	}

	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"success": false,
			"error":   "User not found",
		})
	}

	if err := config.DB.Delete(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"message": "User deleted successfully",
	})
}
