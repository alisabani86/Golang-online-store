package router

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) checkout(ctx *gin.Context) {
	userid := h.GetCookie(ctx)
	if userid == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Login first"})
		return
	}
	cartId := ctx.Query("cartId")
	id, err := strconv.Atoi(cartId)
	// get cart
	totalprice, err := h.Service.GetTotalPrice(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if totalprice == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "cart not found"})
		return
	}

	userid_numb, err := strconv.Atoi(userid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// cek balance
	account, err := h.Service.GetBalance(ctx.Request.Context(), userid_numb)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	if account.Balance < totalprice {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Insufficient balance"})
		return
	}

	//if enough insert trancation

	// create transaction
	err = h.Service.InsertTrancation(ctx.Request.Context(), account.ID, totalprice, account.AccountNumber)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//upadte balance
	err = h.Service.UpdateBalance(ctx.Request.Context(), userid_numb, totalprice)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})

	//insert order
	err = h.Service.InsertOrder(ctx.Request.Context(), userid_numb, float64(totalprice), id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//update deleted_at
	err = h.Service.UpdateShoppingCart(ctx, id)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "order purchased"})
}
