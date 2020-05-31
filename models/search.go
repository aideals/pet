package models 

import (
	"fmt"
	"Pet/constants"
)

type GoodsSku struct {
	ID int64 `gorm:"column:id" json:"id" form:"id"`
	Name string `gorm:"column:name" json:"name" form:"name"`
	Price int64 `gorm:"column:price" json:"price" form:"price"`
	Weight int64 `gorm:"column:weight" json:"weight" form:"weight"`
	ImageUrl string `gorm:"column:image_url" json:"image_url" form:"image_url"`
	Age int64 `gorm:"column:age" json:"age" form:"age"`
	Description string `gorm:"column:description" json:"description" form:"description"`
	Class int64 `gorm:"column:class" json:"class" form:"class"`
	State int64 `gorm:"column:state" json:"state" form:"state"`
	Stock int64 `gorm:"column:stock" json:"stock" form:"stock"`
	Details string `gorm:"column:details" json:"details" form:"details"`
	Logo string `gorm:"column:logo" json:"logo" form:"logo"`
	SpuId int64 `gorm:"column:spu_id" json:"spu_id" form:"spu_id"`
}

func QueryGoodsSkuBySearchLabel(searchLable string) ([]GoodsSku, error) {
	var goodsSku []GoodsSku

	if err := db.Where("searchLable LIKE ?","%jin%").Error; err != nil {
		fmt.Print(constants.DB_ERROR)
		return nil,err
	}

	return goodsSku,nil
}