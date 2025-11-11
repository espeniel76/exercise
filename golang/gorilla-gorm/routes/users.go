package routes

import (
	"encoding/json"
	"gorilla-gorm/config"
	"gorilla-gorm/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// SetupUserRoutes sets up all user-related routes
func SetupUserRoutes(router *mux.Router) {
	// CREATE - 사용자 생성
	router.HandleFunc("/api/users", createUser).Methods("POST")

	// READ - 모든 사용자 조회
	router.HandleFunc("/api/users", getUsers).Methods("GET")

	// READ - 특정 사용자 조회
	router.HandleFunc("/api/users/{id}", getUser).Methods("GET")

	// UPDATE - 사용자 수정
	router.HandleFunc("/api/users/{id}", updateUser).Methods("PUT")

	// DELETE - 사용자 삭제
	router.HandleFunc("/api/users/{id}", deleteUser).Methods("DELETE")
}

// sendJSON sends a JSON response
func sendJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

// createUser creates a new user
func createUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		sendJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	if err := config.DB.Create(&user).Error; err != nil {
		sendJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	sendJSON(w, http.StatusCreated, map[string]interface{}{
		"success": true,
		"data":    user,
	})
}

// getUsers retrieves all users
func getUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User

	if err := config.DB.Order("id ASC").Find(&users).Error; err != nil {
		sendJSON(w, http.StatusInternalServerError, map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	sendJSON(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"data":    users,
	})
}

// getUser retrieves a specific user by ID
func getUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		sendJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "Invalid user ID",
		})
		return
	}

	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		sendJSON(w, http.StatusNotFound, map[string]interface{}{
			"success": false,
			"error":   "User not found",
		})
		return
	}

	sendJSON(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"data":    user,
	})
}

// updateUser updates an existing user
func updateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		sendJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "Invalid user ID",
		})
		return
	}

	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		sendJSON(w, http.StatusNotFound, map[string]interface{}{
			"success": false,
			"error":   "User not found",
		})
		return
	}

	var updateData models.User
	if err := json.NewDecoder(r.Body).Decode(&updateData); err != nil {
		sendJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	// Update user fields
	if err := config.DB.Model(&user).Updates(map[string]interface{}{
		"email":    updateData.Email,
		"username": updateData.Username,
		"password": updateData.Password,
	}).Error; err != nil {
		sendJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	// Fetch updated user
	config.DB.First(&user, id)

	sendJSON(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"data":    user,
	})
}

// deleteUser deletes a user
func deleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		sendJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"error":   "Invalid user ID",
		})
		return
	}

	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		sendJSON(w, http.StatusNotFound, map[string]interface{}{
			"success": false,
			"error":   "User not found",
		})
		return
	}

	if err := config.DB.Delete(&user).Error; err != nil {
		sendJSON(w, http.StatusInternalServerError, map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	sendJSON(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"message": "User deleted successfully",
	})
}
