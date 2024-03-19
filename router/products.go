package router

import (
	"encoding/json"
	"net/http"
	"online_store/internal/presentation"
	"online_store/pkg"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetProduct(ctx *gin.Context) {
	h.GetCookie(ctx)
	category := ctx.Query("category")

	if category == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "category is required"})
		return
	}

	// Create a Redis client
	redisClient := pkg.NewClient()
	defer redisClient.Close() // Close the client when done

	// Check if results are cached
	cachedResults, err := redisClient.Get(ctx, "search:"+category)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if len(string(cachedResults)) == 0 {

		// Results not in cache, fetch from the database
		res, err := h.Service.GetProductBasedOnCategory(ctx.Request.Context(), category)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if len(*res) == 0 {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "no results found"})
			return
		}

		jsonResults, _ := json.Marshal(res)
		err = redisClient.Set("search:"+category, jsonResults, time.Hour)

		if err != nil {

			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error storing results in cache"})
			return
		}

		ctx.JSON(http.StatusOK, res)

		return

	}

	var res []presentation.Product
	err = json.Unmarshal(cachedResults, &res)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
	return

}

func (h *Handler) AddShopingCart(ctx *gin.Context) {
	userid := h.GetCookie(ctx)
	productId := ctx.Query("productId")
	qty := ctx.Query("quantity")
	if productId == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "productId is required"})
		return
	}

	if userid == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Login first"})
	}
	// Check if results are cached
	userid_number, err := strconv.Atoi(userid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	productId_number, err := strconv.Atoi(productId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	qty_num, err := strconv.Atoi(qty)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	psc := presentation.ShopingCart{
		UserID:    userid_number,
		ProductID: productId_number,
		Quantity:  qty_num,
	}

	//get product by product id
	prd, err := h.Service.GetProductById(ctx.Request.Context(), productId_number)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if prd == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
		return
	}

	if prd.Quantity < qty_num {
		ctx.JSON(http.StatusOK, gin.H{"Message": "product limit exceeded"})
		return
	}

	// Results not in cache, fetch from the database
	err = h.Service.AddCart(ctx, psc)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"Message": "success add cart"})

}

func (h *Handler) GetListCart(ctx *gin.Context) {
	userid := h.GetCookie(ctx)
	userid_num, err := strconv.Atoi(userid)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	res, err := h.Service.GetListCart(ctx, userid_num)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if res == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "no results found"})
		return
	}
	ctx.JSON(http.StatusOK, res)

}

func (h *Handler) DeleteCart(ctx *gin.Context) {
	uid := h.GetCookie(ctx)

	if uid == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Login first"})
		return
	}
	cartid := ctx.Query("id")

	if cartid == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}
	cartid_num, err := strconv.Atoi(cartid)

	err = h.Service.DeleteShopingCart(ctx, cartid_num)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Message": "success delete cart"})
}
