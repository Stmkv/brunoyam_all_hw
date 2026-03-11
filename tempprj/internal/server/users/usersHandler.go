package users

import (
	"net/http"
	usersDomain "temp-prj/internal/domain/users"

	"github.com/gin-gonic/gin"
)

type UserService interface {
	RegisterUser(req usersDomain.RegisterRequest) (string, error)
	LoginUser(req usersDomain.LoginRequest) (usersDomain.User, error)
}

type UsersHandler struct {
	userService UserService
}

func New(userService UserService) *UsersHandler {
	return &UsersHandler{
		userService: userService,
	}
}

// Login godoc
// @Summary Авторизация
// @Tags users
// @Accept json
// @Produce json
// @Param input body usersDomain.LoginRequest true "Login data"
// @Success 200 {object} usersDomain.User
// @Router /users/login [post]
func (h *UsersHandler) Login(ctx *gin.Context) {
	var req usersDomain.LoginRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.userService.LoginUser(req)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

// Register godoc
// @Summary Регистрация
// @Tags users
// @Accept json
// @Produce json
// @Param input body usersDomain.RegisterRequest true "Register data"
// @Success 200 {object} map[string]string
// @Router /users/register [post]
func (h *UsersHandler) Register(ctx *gin.Context) {
	var req usersDomain.RegisterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	uid, err := h.userService.RegisterUser(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"uid": uid})
}
