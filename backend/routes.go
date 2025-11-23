package main

import "github.com/gin-gonic/gin"

func SetupRoutes(r *gin.Engine) {
	// USERS
	r.POST("/users", CreateUser)
	r.GET("/users", ListUsers)
	r.POST("/users/login", Login)

	// ITEMS
	r.POST("/items", CreateItem)
	r.GET("/items", ListItems)

	// CARTS
	r.POST("/carts", AddToCart)
	r.GET("/carts", ListCarts)

	// ORDERS
	r.POST("/orders", CreateOrder)
	r.GET("/orders", ListOrders)
}
