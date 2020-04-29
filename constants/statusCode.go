package constants

const (
	SUCCESS        = 200 //成功
	ERROR          = 500 //错误
	DB_ERROR       = 501 //数据库错误
	INVALID_PARAMS = 400 //请求参数错误

	ERROR_AUTH_CHECK_TOKEN_FAIL    = 20001 //token鉴权失败
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT = 20002 //token已超时
	ERROR_AUTH_TOKEN               = 20003 //token生成失败
	ERROR_AUTH                     = 20004 //token错误
	ERROR_OPENID_SESSIONKEY_FAIL   = 20005 //获取openid和session_key错误
)
