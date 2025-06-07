package main

import (
	"fmt"
	NewUserType "myGoApp/types"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {

	//This gin.Default set up a router with logger abd recovery middleware attached
	router := gin.Default()
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "hello from gin framework",
			"Status":  "success",
		})
	})
	router.GET("/users/data", func(c *gin.Context) {
		users := []gin.H{
			{"name": "sangam Silwal", "id": 123},
			{"name": "Ram lal chaudhary", "id": 123},
		}

		c.JSON(http.StatusOK, gin.H{
			"Data":    users,
			"message": "User retrieved Succesfully",
		})
	})

	router.POST("/users", func(c *gin.Context) {
		var user NewUserType.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "ValidationError",
				"Details": err.Error(),
			})
			return
		}
		c.JSON(http.StatusCreated, gin.H{
			"message": "User Created Successfully",
			"User":    user,
		})
	})

	router.GET("/users/:id", func(ctx *gin.Context) {
		userID := ctx.Param("id")
		id, err := strconv.Atoi(userID)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid user ID format",
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"Message": "Param taken Successfully",
			"ParamId": id,
		})
	})

	//Multiple parameters in URL
	router.GET("/filter", func(ctx *gin.Context) {
		tags := ctx.QueryArray("tags")
		fmt.Println(tags)
		categories := ctx.Query("tags")
		fmt.Println(categories)
		ctx.JSON(http.StatusOK, gin.H{
			"tags":       tags,
			"categories": categories,
			"message":    "Filter applied",
		})
	})

	router.Run(":8080")
}
