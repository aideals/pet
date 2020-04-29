package api

import (
	"Pet/models"
	"Pet/constants"
	"log"
	"net/http"
)

//收藏
func CommodityCollection(c *gin.Context) {
	openId := c.Query("openId")
	goodId := c.Query("goodsId")

	if openId != "" && goodId != "" {
	   err := models.AddToWishlist(openId,goodId)
	   if err != nil {
          log.Printf(constants.DB_ERROR)
	   }
	}
	
	c.JSON(http.StatusOK, gin.H{
		"msg" :  "success",
	})
}

//取消收藏
func DelCommodityCollection(c *gin.Context) {
	openId := c.Query("openId")
	goodId := c.Query("goodsId")

	if openId != "" && goodId != "" {
		err := models.DelCommodityCollection(openId,goodId)
		if err != nil {
			log.Printf(constants.DB_ERROR)
		}
	}
	
	c.JSON(http.StatusOK, gin.H{
		"msg" : "success",
	})
}

//加入购物车
func AddToShoppingCart(c *gin.Context) {

}