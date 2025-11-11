package routes

import (
	"fiber-gorm/config"
	"fiber-gorm/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// SetupUserRoutes sets up all user-related routes
func SetupUserRoutes(app *fiber.App) {
	// CREATE - 사용자 생성
	app.Post("/api/users", createUser)

	// READ - 모든 사용자 조회
	app.Get("/api/users", getUsers)

	// READ - 특정 사용자 조회
	app.Get("/api/users/:id", getUser)

	// UPDATE - 사용자 수정
	app.Put("/api/users/:id", updateUser)

	// DELETE - 사용자 삭제
	app.Delete("/api/users/:id", deleteUser)
}

// createUser creates a new user
func createUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	if err := config.DB.Create(&user).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"data":    user,
	})
}

// getUsers retrieves all users
func getUsers(c *fiber.Ctx) error {
	var users []models.User

	if err := config.DB.Order("id ASC").Find(&users).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    users,
	})
}

// getUser retrieves a specific user by ID
func getUser(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   "Invalid user ID",
		})
	}

	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"error":   "User not found",
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    user,
	})
}

// updateUser updates an existing user
func updateUser(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   "Invalid user ID",
		})
	}

	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"error":   "User not found",
		})
	}

	var updateData models.User
	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
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
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	// Fetch updated user
	config.DB.First(&user, id)

	return c.JSON(fiber.Map{
		"success": true,
		"data":    user,
	})
}

// deleteUser deletes a user
func deleteUser(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   "Invalid user ID",
		})
	}

	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"error":   "User not found",
		})
	}

	if err := config.DB.Delete(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "User deleted successfully",
	})
}
