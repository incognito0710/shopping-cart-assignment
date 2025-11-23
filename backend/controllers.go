package main

import (
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ---------------- USER ------------------

func CreateUser(c *gin.Context) {
	var user User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body"})
		return
	}

	if err := DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "create user failed"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func ListUsers(c *gin.Context) {
	var users []User
	DB.Find(&users)
	c.JSON(http.StatusOK, users)
}

func Login(c *gin.Context) {
	var loginData User
	if err := c.BindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body"})
		return
	}

	var user User
	if err := DB.Where("username = ? AND password = ?", loginData.Username, loginData.Password).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username/password"})
		return
	}

	// Generate simple token (for assignment only)
	token := strconv.Itoa(rand.Int())
	user.Token = token
	if err := DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "save token failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

// ---------------- ITEMS ------------------

func CreateItem(c *gin.Context) {
	var item Item
	if err := c.BindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body"})
		return
	}
	if err := DB.Create(&item).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "create item failed"})
		return
	}
	c.JSON(http.StatusOK, item)
}

func ListItems(c *gin.Context) {
	var items []Item
	DB.Find(&items)
	c.JSON(http.StatusOK, items)
}

// ---------------- CARTS ------------------

type addCartBody struct {
	ItemID uint `json:"item_id"`
}

func AddToCart(c *gin.Context) {
	token := c.GetHeader("token")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token required"})
		return
	}

	var user User
	if err := DB.Where("token = ?", token).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	var body addCartBody
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body"})
		return
	}

	cart := Cart{
		UserID: user.ID,
		ItemID: body.ItemID,
	}
	if err := DB.Create(&cart).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "add to cart failed"})
		return
	}

	c.JSON(http.StatusOK, cart)
}

func ListCarts(c *gin.Context) {
	var carts []Cart
	DB.Find(&carts)
	c.JSON(http.StatusOK, carts)
}

// ---------------- ORDERS ------------------

type createOrderBody struct {
	CartID uint `json:"cart_id"`
	UserID uint `json:"user_id"`
}

func CreateOrder(c *gin.Context) {
	var body createOrderBody
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body"})
		return
	}

	order := Order{
		UserID: body.UserID,
		CartID: body.CartID,
	}
	if err := DB.Create(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "create order failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order created!", "order": order})
}

func ListOrders(c *gin.Context) {
	var orders []Order
	DB.Find(&orders)
	c.JSON(http.StatusOK, orders)
}
