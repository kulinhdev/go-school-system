package controllers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/kulinhdev/serentyspringsedu/helpers"
	"github.com/kulinhdev/serentyspringsedu/initializers"
	"github.com/kulinhdev/serentyspringsedu/models"
	"gorm.io/gorm"
)

type AuthController struct {
	DB *gorm.DB
}

func NewAuthController(DB *gorm.DB) AuthController {
	return AuthController{DB: DB}
}

// RegisterUser registers a new user.
// @Summary Register a new user
// @ID register-user
// @Tags User - Authentication
// @Accept json
// @Produce json
// @Param user body models.UserRegister true "User data"
// @Success 201 {object} models.UserResponse
// @Router /api/auth/register [post]
func (ctl *AuthController) RegisterUser(ctx *gin.Context) {
	var payload *models.UserRegister

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
		return
	}

	if payload.Password != payload.PasswordConfirm {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Passwords do not match"})
		return
	}

	hashedPassword, err := helpers.HashPassword(payload.Password)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": err.Error()})
		return
	}

	newUser := models.User{
		Name:     payload.Name,
		Username: payload.Username,
		Email:    strings.ToLower(payload.Email),
		Password: hashedPassword,
		Status:   1,
		Role:     1,
		Gender:   1,
		Photo:    payload.Photo,
	}

	result := ctl.DB.Create(&newUser)

	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "duplicate key value violates unique") {
			ctx.JSON(http.StatusConflict, gin.H{"status": "error", "message": "User with that email already exists"})
		} else {
			ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": result.Error.Error()})
		}
		return
	}

	userResponse := &models.UserResponse{
		ID:        newUser.ID,
		Name:      newUser.Name,
		Username:  newUser.Username,
		Email:     newUser.Email,
		Phone:     newUser.Phone,
		Photo:     newUser.Photo,
		Gender:    newUser.Gender,
		CreatedAt: newUser.CreatedAt,
		UpdatedAt: newUser.UpdatedAt,
	}
	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": gin.H{"user": userResponse}})
}

// LoginUser logs in a user.
// @Summary Log in a user
// @ID login-user
// @Tags User - Authentication
// @Accept json
// @Produce json
// @Param user body models.UserLogin true "User login data"
// @Success 200 {object} models.UserResponse
// @Router /api/auth/login [post]
func (ctl *AuthController) LoginUser(ctx *gin.Context) {
	var payload *models.UserLogin

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
		return
	}

	var user models.User
	result := ctl.DB.First(&user, "email = ?", strings.ToLower(payload.Email))
	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid email!"})
		return
	}

	if err := helpers.CheckPassword(payload.Password, user.Password); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Password is not correct!"})
		return
	}

	config := initializers.Config

	accessToken, err := helpers.CreateToken(config.AccessTokenExpiresIn, user.ID, config.AccessTokenPrivateKey)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
		return
	}

	refreshToken, err := helpers.CreateToken(config.RefreshTokenExpiresIn, user.ID, config.RefreshTokenPrivateKey)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
		return
	}

	ctx.SetCookie("access_token", accessToken, config.AccessTokenMaxAge*60, "/", "localhost", false, true)
	ctx.SetCookie("refresh_token", refreshToken, config.RefreshTokenMaxAge*60, "/", "localhost", false, true)
	ctx.SetCookie("logged_in", "true", config.AccessTokenMaxAge*60, "/", "localhost", false, false)

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "access_token": accessToken})
}

// RefreshAccessToken refreshes the access token.
// @Summary Refresh the access token
// @ID refresh-access-token
// @Tags User - Authentication
// @Accept json
// @Produce json
// @Success 200 {object} models.UserResponse
// @Router /api/auth/refresh [get]
func (ctl *AuthController) RefreshAccessToken(ctx *gin.Context) {
	message := "Could not refresh access token: %s!"

	cookie, err := ctx.Cookie("refresh_token")

	if err != nil {
		errorMessage := fmt.Sprintf(message, err)
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "error", "message": errorMessage})
		return
	}

	config := initializers.Config

	sub, err := helpers.ValidateToken(cookie, config.RefreshTokenPublicKey)
	if err != nil {
		errorMessage := fmt.Sprintf(message, err)
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "error", "message": errorMessage})
		return
	}

	var user models.User
	result := ctl.DB.First(&user, "id = ?", fmt.Sprint(sub))
	if result.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "error", "message": "the user belonging to this token no longer exists"})
		return
	}

	accessToken, err := helpers.CreateToken(config.AccessTokenExpiresIn, user.ID, config.AccessTokenPrivateKey)
	if err != nil {
		errorMessage := fmt.Sprintf(message, err)
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "error", "message": errorMessage})
		return
	}

	ctx.SetCookie("access_token", accessToken, config.AccessTokenMaxAge*60, "/", "localhost", false, true)
	ctx.SetCookie("logged_in", "true", config.AccessTokenMaxAge*60, "/", "localhost", false, false)

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "access_token": accessToken})
}

// LogoutUser logs out a user.
// @Summary Log out a user
// @ID logout-user
// @Tags User - Authentication
// @Accept json
// @Produce json
// @Router /api/auth/logout [get]
func (ctl *AuthController) LogoutUser(ctx *gin.Context) {
	ctx.SetCookie("access_token", "", -1, "/", "localhost", false, true)
	ctx.SetCookie("refresh_token", "", -1, "/", "localhost", false, true)
	ctx.SetCookie("logged_in", "", -1, "/", "localhost", false, false)

	ctx.JSON(http.StatusOK, gin.H{"status": "Logged out success!"})
}
