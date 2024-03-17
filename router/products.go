package router

import (
	"encoding/json"
	"net/http"
	"server/internal/presentation"
	"server/pkg"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetProduct(ctx *gin.Context) {
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

		// Store results in cache
		jsonResults, _ := json.Marshal(res)
		err = redisClient.Set("search:"+category, jsonResults, time.Hour)

		if err != nil {
			// Handle cache set error (log, etc.)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error storing results in cache"})
			return
		}

		ctx.JSON(http.StatusOK, res)

	}
	var res []presentation.Product
	err = json.Unmarshal(cachedResults, &res)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, res)

}
