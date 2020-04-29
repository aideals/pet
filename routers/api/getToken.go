package api 

import (
	"log"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"Pet/pkg/util"
	"Pet/constants"
	"Pet/models"
)

//从微信获取open_id和sessionKey，并生成token
func GetCodeFromTX(c *gin.Context) {
	//获取code
	code := c.Query("code")
	//根据code获取openId和session_key
	wxLoginResp,err := util.WXLogin(code)
	if err != nil {
	   c.JSON(400, gin.H {
		   "ErrCode": wxLoginResp.ErrCode,
		   "ErrMsg":  wxLoginResp.ErrMsg,
	   })
	   return 
	}
	
	data := make(map[string]interface{})
	data["openid"] = wxLoginResp.openid
	data["session_key"] = wxLoginResp.sessionKey

	isSuccess := models.AddUserAuth(data)
	if !isSuccess {
		log.Printf(constants.DB_ERROR)
	}
   
	//这里用openId和sessionKey的串接，进行MD5之后作为该用户的自定义登录态
	myToken,err := util.GenerateToken(wxLoginResp.OpenId,wxLoginResp.SessionKey)
	if err != nil {
		log.Printf("err",err)
	}
	//接下来可以将openId和sessionKey,myToken存储到数据库中
	//但这里要保证mySession唯一，以便于用mySession去索引openId 和 sessionKey
	c.JSON(http.StatusOK, gin.H{
	   "openId": wxLoginResp.OpenId,
	   "token":  myToken, 
	}) 
}

