package util 

import (
   
	"net/http"
	"fmt"
	"encoding/json"
	"errors"
)

type WXLoginResp struct {
	OpenId   string    `json:"openid"`
	SessionKey string  `json:"session_key"`
	UnionId   string   `json:"unionid"`
	ErrCode   int      `json:"errcode"`
	ErrMsg    string   `json:"errmsg"`
	ExpireIn    int    `json:"expire_in"`
}

//通过code获取openid和session_key
func WXLogin(code string) (*WXLoginResp,error) {
	url := "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"
	
	//合成url,这里的appId和secret是在微信公众平台上获取的
	//url = fmt.Sprintf(url,appId,secret,code)

	//创建http get请求
	resp,err := http.Get(url)
	if err != nil {
		return nil, err
	}
	
	defer resp.Body.Close()

	//解析http请求中的body的数据到定义的结构体中
	wxResp := WXLoginResp{}
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&wxResp); err != nil {
		return nil,err
	}

	//判断微信接口返回的是否是一个异常情况
	if wxResp.ErrCode != 0 {
		return nil, errors.New(fmt.Sprintf("ErrCode:%s ErrMsg:%s",wxResp.ErrCode,wxResp.ErrMsg))
	}

	return &wxResp,nil
}

