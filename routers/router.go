package routers

import (
	"github.com/gin-gonic/gin"

	"Pet/routers/api"
	"Pet/setting"
	"Pet/middleware/jwt"
	"Pet/handler"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)
	
	// r.GET("/auth",api.GetAuth)

	r.GET("/account/login",handler.GetCodeFromTX)               //从前端获取code
	r.GET("/home/page",handler.HomePage)                        //首页
	r.GET("/home/dog",handler.Dog)                              //获取狗狗的spu
	r.GET("/home/cat",handler.Cat)                              //获取猫的spu
	r.GET("/home/petSupplies",handler.PetSupplies)              //获取宠物用品的spu
	r.GET("/home/goodsDetail",handler.GoodsDetail)              //获取商品详情
	r.GET("/home/goodsSearch",handler.GoodsSearch)              //搜索
	r.GET("/goodsDetail/collect",handler.CommodityCollection)   //商品收藏
	r.GET("/goodsDetail/cancel",handler.CancelCollection)       //取消收藏
	r.GET("/goodsDetail/shoppingCart",handler.AddToShoppingCart)    //加入购物车

	return r
}
