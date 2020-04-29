package models 

import (
	"error"
	"log"
	"Pet/constants"
)

type WishlistGoodsSpu struct {
	ID int64 `gorm:"column:id" json:"id" form:"id"`
	Name string `gorm:"column:name" json:"name" form:"name"`
	Price int64 `gorm:"column:price" json:"price" form:"price"`
	Weight int64 `gorm:"column:weight" json:"weight" form:"weight"`
	ImageUrl string `gorm:"column:image_url" json:"image_url" form:"image_url"`
	Age int64 `gorm:"column:age" json:"age" form:"age"`
	Classify int64 `gorm:"column:classify" json:"classify" form:"classify"`
	Users []User `gorm:"many2many:commodityCollection"`
}

type User struct {
	ID int64 `gorm:"column:id" json:"id" form:"id"`
	UserName string `gorm:"column:user_name" json:"user_name" form:"user_name"`
	PhoneNumber int64 `gorm:"column:phone_number" json:"phone_number" form:"phone_number"`
	Age string `gorm:"column:age" json:"age" form:"age"`
	Sex string `gorm:"column:sex" json:"sex" form:"sex"`
	HeadPortrait string `gorm:"column:head_portrait" json:"head_portrait" form:"head_portrait"`
	Password int64 `gorm:"column:password" json:"password" form:"password"`
	Openid string `gorm:"column:openid" json:"openid" form:"openid"`
	SessionKey string `gorm:"column:session_key" json:"session_key" form:"session_key"`
	WishlistGoodsSpu []WishlistGoodsSpu `gorm:"many2many:commodityCollection"`
}

type CommodityCollection struct {
	Userid int64 `gorm:"column:userid" json:"userid" form:"userid"`
	Goodsspuid int64 `gorm:"column:goodsspuid" json:"goodsspuid" form:"goodsspuid"`
}

//添加到购物车
func AddToWishlist(openId,goodsId int) {
	commodityCollection := CommodityCollection{Userid: int64(openId), Goodsspuid: int64(goodsId)}
	
	if err := db.Create(&commodityCollection).Error; err != nil {
       log.Printf(constants.DB_ERROR)
	}
}

//从购物车删除
func DelWishlist(goodsId,openId int) int {
	commodityCollection := CommodityCollection{Userid: int64(openId), Goodsspuid: int64(goodsId)}

	if err := db.Model(&User{}).Association("WishlistGoodsSpu").Delete(&commodityCollection).Error; err != nil {
		log.Printf(constants.DB_ERROR)
	}

	return constants.SUCCESS
}

