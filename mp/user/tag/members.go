package tag

import (
	"errors"

	"github.com/shiguanghuxian/wechat/mp/core"
)

// 文档地址 https://developers.weixin.qq.com/doc/offiaccount/User_Management/Manage_blacklist.html

type Black struct {
	Total int `json:"total"` // 黑名单总数
	Count int `json:"count"` // 拉取的OPENID个数, 最大值为10000

	Data struct {
		OpenIdList []string `json:"openid,omitempty"`
	} `json:"data"` // 列表数据, OPENID的列表

	// 拉取列表的最后一个用户的OPENID, 如果 next_openid == "" 则表示没有了用户数据
	NextOpenId string `json:"next_openid"`
}

// GetBlackList 获取公众号的黑名单列表
func GetBlackList(clt *core.Client, beginOpenid string) (black *Black, err error) {
	var incompleteURL = "https://api.weixin.qq.com/cgi-bin/tags/members/getblacklist?access_token="
	var request struct {
		BeginOpenid string `json:"begin_openid,omitempty"`
	}
	request.BeginOpenid = beginOpenid
	var result struct {
		core.Error
		Black Black `json:"black"`
	}
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	black = &result.Black
	return
}

// BatchBlackList 拉黑用户
func BatchBlackList(clt *core.Client, openidList ...string) (err error) {
	if len(openidList) == 0 {
		err = errors.New("加入黑名单openidList为空")
		return
	}
	var incompleteURL = "https://api.weixin.qq.com/cgi-bin/tags/members/batchblacklist?access_token="
	var request struct {
		OpenidList []string `json:"openid_list"`
	}
	request.OpenidList = openidList
	var result struct {
		core.Error
	}
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	return
}

// BatchUnBlackList 取消拉黑用户
func BatchUnBlackList(clt *core.Client, openidList ...string) (err error) {
	if len(openidList) == 0 {
		err = errors.New("加入黑名单openidList为空")
		return
	}
	var incompleteURL = "https://api.weixin.qq.com/cgi-bin/tags/members/batchunblacklist?access_token="
	var request struct {
		OpenidList []string `json:"openid_list"`
	}
	request.OpenidList = openidList
	var result struct {
		core.Error
	}
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	return
}
