package homepage

import (
	"github.com/shiguanghuxian/wechat/mp/core"
)

// 默认模板
func NewSetParameters1(shopId int64) interface{} {
	return &struct {
		ShopId     int64 `json:"shop_id"`
		TemplateId int64 `json:"template_id"`
	}{
		ShopId:     shopId,
		TemplateId: 0,
	}
}

// 自定义url
func NewSetParameters2(shopId int64, url string) interface{} {
	para := struct {
		ShopId     int64 `json:"shop_id"`
		TemplateId int64 `json:"template_id"`
		Struct     struct {
			URL string `json:"url"`
		} `json:"struct"`
	}{
		ShopId:     shopId,
		TemplateId: 1,
	}

	para.Struct.URL = url
	return &para
}

// 设置商家主页
//  要求 para 经过 encoding/json 后满足指定的格式要求
func Set(clt *core.Client, para interface{}) (err error) {
	var result core.Error

	incompleteURL := "https://api.weixin.qq.com/bizwifi/homepage/set?access_token="
	if err = clt.PostJSON(incompleteURL, para, &result); err != nil {
		return
	}

	if result.ErrCode != core.ErrCodeOK {
		err = &result
		return
	}
	return
}
