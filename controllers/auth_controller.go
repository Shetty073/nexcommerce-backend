package controllers

import (
	"nexcommerce/models"
	"nexcommerce/responses"
	"nexcommerce/schemas"
	"nexcommerce/utils/jwt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// RegisterController handles user registration
// @Summary Register a new user
// @Description Creates a new user with the provided details
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body schemas.RegisterSchema true "User registration data"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /register [post]
func RegisterController(c *gin.Context) {
	var input schemas.RegisterSchema

	// Bind JSON request body to input struct
	if err := c.ShouldBindJSON(&input); err != nil {
		responses.BadRequest(c, "Validation Error", err.Error())
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		responses.InternalServerError(c, "Hashing Error", "Failed to hash password")
		return
	}

	// Create user model
	user := models.User{
		ID:           uuid.New(),
		Username:     input.Username,
		Email:        input.Email,
		Password:     string(hashedPassword),
		FirstName:    input.FirstName,
		LastName:     input.LastName,
		DateOfBirth:  input.DateOfBirth,
		Gender:       input.Gender,
		MobileNumber: input.MobileNumber,
		IsCustomer:   true,
	}

	// Save to database
	if err := user.CreateUser(); err != nil {
		responses.InternalServerError(c, err.Error(), "Failed to create user")
		return
	}

	// Respond with success
	responses.Created(c, gin.H{"message": "User registered successfully"})
}

// LoginController handles user authentication
// @Summary User login
// @Description Authenticates a user with email/username and password, and returns a JWT token
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body schemas.LoginSchema true "User login data"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /login [post]
func LoginController(c *gin.Context) {
	var input schemas.LoginSchema

	// Bind JSON request body to input struct
	if err := c.ShouldBindJSON(&input); err != nil {
		responses.BadRequest(c, "Validation Error", err.Error())
		return
	}

	// Find user by email or username
	var user models.User
	err := models.GetUserByEmailOrUsername(input.Email, input.Username, &user)
	if err != nil {
		responses.Unauthorized(c, "Authentication Error", "Invalid credentials")
		return
	}

	// Compare provided password with stored hash
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		responses.Unauthorized(c, "Authentication Error", "Invalid credentials")
		return
	}

	// Generate JWT token using your existing function
	tokenString, err := jwt.GenerateToken(user.Username, 24) // Token valid for 24 hours
	if err != nil {
		responses.InternalServerError(c, "Token Generation Error", "Failed to generate token")
		return
	}

	// Respond with token
	responses.Ok(c, gin.H{"token": tokenString})
}
