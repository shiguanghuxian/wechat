package page

import (
	"github.com/shiguanghuxian/wechat/mp/core"
)

type UpdateParameters struct {
	PageId      int64  `json:"page_id"`               // 必须, 摇周边页面唯一ID
	Title       string `json:"title,omitempty"`       // 必须, 在摇一摇页面展示的主标题，不超过6个字
	Description string `json:"description,omitempty"` // 必须, 在摇一摇页面展示的副标题，不超过7个字
	PageURL     string `json:"page_url,omitempty"`    // 必须, 跳转链接
	IconURL     string `json:"icon_url,omitempty"`    // 必须, 在摇一摇页面展示的图片。图片需先上传至微信侧服务器，用“素材管理-上传图片素材”接口上传图片，返回的图片URL再配置在此处
	Comment     string `json:"comment,omitempty"`     // 可选, 页面的备注信息，不超过15个字
}

// 编辑页面信息
func Update(clt *core.Client, para *UpdateParameters) (err error) {
	var result core.Error

	incompleteURL := "https://api.weixin.qq.com/shakearound/page/update?access_token="
	if err = clt.PostJSON(incompleteURL, para, &result); err != nil {
		return
	}

	if result.ErrCode != core.ErrCodeOK {
		err = &result
		return
	}
	return
}
