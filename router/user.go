package router

import (
	"net/http"
	"online_store/internal/presentation"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateUser(ctx *gin.Context) {

	var u presentation.CreateUserRequest
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

	var userreq presentation.LoginUserRequest
	if err := ctx.ShouldBindJSON(&userreq); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	u, err := h.Service.Login(ctx.Request.Context(), &userreq)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.SetCookie("jwt", u.AccesToken, 3600, "/", "localhost", false, true)
	res := &presentation.LoginUserResponse{
		Username: u.Username,
		ID:       u.ID,
	}
	ctx.JSON(http.StatusOK, res)
}

func (h *Handler) GetCookie(ctx *gin.Context) string {

	jwtCookie, err := ctx.Cookie("jwt")
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "session not found"})
		return ""
	}
	claim, err := h.Middleware.VerifyJWTToken(jwtCookie)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return ""
	}

	ctx.JSON(http.StatusOK, gin.H{"jwt": claim})
	return claim.ID
}

func (h *Handler) Logout(ctx *gin.Context) {
	ctx.SetCookie("jwt", "", -1, "/", "localhost", false, true)
	ctx.JSON(http.StatusOK, gin.H{"message": "Logout successfully"})

}
