package routers

import (
	"github.com/gin-gonic/gin"

	"Pet/routers/api"
	// "Pet/routers/api/v1"
	"Pet/setting"
	"Pet/middleware/jwt"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)
	
	// r.GET("/auth",api.GetAuth)

	r.GET("/account/login",api.GetCodeFromTX)               //从前端获取code
	r.GET("/home/page",api.HomePage)                        //首页
	r.GET("/home/dog",api.Dog)                              //获取狗狗的spu
	r.GET("/home/cat",api.Cat)                              //获取猫的spu
	r.GET("/home/petSupplies",api.PetSupplies)              //获取宠物用品的spu
	r.GET("/home/goodsDetail",api.GoodsDetail)              //获取商品详情
	r.GET("/home/goodsSearch",api.GoodsSearch)              //搜索
	r.GET("/goodsDetail/collect",api.CommodityCollection)   //商品收藏
	r.GET("/goodsDetail/cancel",api.CancelCollection)       //取消收藏
	r.GET("/goodsDetail/shoppingCart",api.AddToShoppingCart)    //加入购物车

	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{

	}

	return r
}
