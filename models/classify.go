package models 

import (
)

type ClassifyGoodsSpu struct {
	ID int64 `gorm:"column:id" json:"id" form:"id"`
	Name string `gorm:"column:name" json:"name" form:"name"`
	Price int64 `gorm:"column:price" json:"price" form:"price"`
	Weight int64 `gorm:"column:weight" json:"weight" form:"weight"`
	ImageUrl string `gorm:"column:image_url" json:"image_url" form:"image_url"`
	Age int64 `gorm:"column:age" json:"age" form:"age"`
	Classify int64 `gorm:"column:classify" json:"classify" form:"classify"`
}

func QueryDogByClassify(classify int)([]ClassifyGoodsSpu,error) {
	var goodsSpu []ClassifyGoodsSpu
	
	if err := db.Limit(12).Offset(12).Where("classify = ?",classify).Find(&goodsSpu); err != nil {
		return nil,nil
	}

	return goodsSpu,nil
}

func QueryCatByClassify(classify int)([]ClassifyGoodsSpu,error) {
	var goodsSpu []ClassifyGoodsSpu

	if err := db.Limit(12).Offset(12).Where("classify = ?",classify).Find(&goodsSpu); err != nil {
		return nil,nil
	}

	return goodsSpu,nil
}

func QueryPetSuppliesByClassify(classify int)([]ClassifyGoodsSpu,error) {
	var goodsSpu []ClassifyGoodsSpu

	if err := db.Limit(12).Offset(12).Where("classify = ?",classify).Find(&goodsSpu); err != nil {
		return nil,nil
	}

	return goodsSpu,nil
}

