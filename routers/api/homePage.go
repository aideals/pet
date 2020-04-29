package api

import (
	"log"
	"Pet/models"
	"Pet/constants"
)

//首页
func HomePage(c *gin.Context) {
	
   viewPagers,err := models.QueryViewPager()
   if err != nil {
      log.Printf(constants.DB_ERROR)
   }

   goodsSpu,err := models.QueryGoodsSpu()
   if err != nil {
	   log.Printf(constants.DB_ERROR)
   }

   c.JSON(http.StatusOK, gin.H{
	  "viewPagers" : viewPagers,
	  "goodsSpu"   : goodsSpu,
   })
}

//商品详情
func GoodsDetail(c *gin.Context) {
	spuId := c.Query("id")
	
	if spuId != "" {
		goodsDetail,err := models.QueryGoodsSkuByGoodsId(spuId)
		if err != nil {
			log.Printf(constants.DB_ERROR)	
		}

		c.JSON(http.StatusOK, gin.H{
			"id"             : goodsDetail.Id,
			"name"           : goodsDetail.Name,
			"price"          : goodsDetail.Price,
			"weight"         : goodsDetail.Weight,
			"imageUrl"       : goodsDetail.ImageUrl,
			"age"            : goodsDetail.Age,
			"description"    : goodsDetail.Description,
			"class"          : goodsDetail.Class,
			"state"          : goodsDetail.State,
			"stock"          : goodsDetail.Stock,
			"details"        : goodsDetail.Details,
			"logo"           : goodsDetail.Logo,
			"spuId"          : goodsDetail.SupId,
		})
	}
}

//搜索
func GoodsSearch(c *gin.Context) {
	searchTitle := c.Query("searchTitle")

	if searchTitle != "" {
		goodsSku,err := models.QueryGoodsSkuBySearchLabel(searchTitle)
		if err != nil {
			log.Printf(constants.DB_ERROR)
		}

		c.JSON(http.StatusOK, gin.H{
			"goodsSku" : goodsSku,
		})
	}
}