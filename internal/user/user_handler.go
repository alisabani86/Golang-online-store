package user

import (
	"net/http"
	"server/internal/middleware"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service
	middleware.Middleware
}

func NewHandler(s Service, m middleware.Middleware) *Handler {
	return &Handler{
		Service:    s,
		Middleware: m,
	}
}

func (h *Handler) CreateUser(ctx *gin.Context) {

	var u CreateUserRequest
	if err := ctx.ShouldBindJSON(&u); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	res, err := h.Service.CreateUser(ctx.Request.Context(), &u)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, res)

}

func (h *Handler) Login(ctx *gin.Context) {

	var user LoginUserRequest
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	u, err := h.Service.Login(ctx.Request.Context(), &user)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.SetCookie("jwt", u.AccesToken, 3600, "/", "localhost", false, true)
	res := &LoginUserResponse{
		Username: u.Username,
		ID:       u.ID,
	}
	ctx.JSON(http.StatusOK, res)
}

func (h *Handler) GetCookie(ctx *gin.Context) {

	jwtCookie, err := ctx.Cookie("jwt")
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "session not found"})
		return
	}
	claim, err := h.Middleware.VerifyJWTToken(jwtCookie)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	print(claim)

	ctx.JSON(http.StatusOK, gin.H{"jwt": claim})
}

func (h *Handler) Logout(ctx *gin.Context) {
	ctx.SetCookie("jwt", "", -1, "/", "localhost", false, true)
	ctx.JSON(http.StatusOK, gin.H{"message": "Logout successfully"})

}
