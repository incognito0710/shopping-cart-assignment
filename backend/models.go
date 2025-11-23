package main

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Username string `json:"username" gorm:"uniqueIndex;not null"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

type Item struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type Cart struct {
	ID     uint `json:"id" gorm:"primaryKey"`
	UserID uint `json:"user_id"`
	ItemID uint `json:"item_id"`
}

type Order struct {
	ID     uint `json:"id" gorm:"primaryKey"`
	UserID uint `json:"user_id"`
	CartID uint `json:"cart_id"`
}
