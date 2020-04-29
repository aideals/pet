package models

import (
	"log"
)

type HomePageResp struct {
	ViewPager []ViewPager 
	GoodsSpu  []GoodsSpu
}

type ViewPager struct {
	Id   int     `gorm:"column:id" json:"id"`
	Url  string  `gorm:"column:url" json:"url"`
}

type GoodsSpu struct {
	ID int64 `gorm:"column:id" json:"id" form:"id"`
	Name string `gorm:"column:name" json:"name" form:"name"`
	Price int64 `gorm:"column:price" json:"price" form:"price"`
	Weight int64 `gorm:"column:weight" json:"weight" form:"weight"`
	ImageUrl string `gorm:"column:image_url" json:"image_url" form:"image_url"`
	Age int64 `gorm:"column:age" json:"age" form:"age"`
}

func QueryViewPager() ([]ViewPager, err) {
  var viewPager []ViewPager

  if err := db.Limit(3).Find(&viewPager).Error; err != nil {
	log.Printf(constants.DB_ERROR)
	return nil,err	
  } 
  
  return viewPager,nil
}

func QueryGoodsSpu() ([]GoodsSpu, err) {
  var goodsSpu []GoodsSpu

  if err :=	db.Limit(12).Offset(12).Find(&goodsSpu).Error; err != nil {
	log.Printf(constants.DB_ERROR)
	return nil,err
  }

  return goodsSpu,nil
}





	
    
