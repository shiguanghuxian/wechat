package meetingticket

import (
	"github.com/shiguanghuxian/wechat/mp/core"
)

type UpdateUserParameters struct {
	Code   string `json:"code"`              // 必须; 用户的门票唯一序列号
	CardId string `json:"card_id,omitempty"` // 可选; 要更新门票序列号所述的card_id ,  生成券时use_custom_code 填写true 时必填.

	Zone       string `json:"zone,omitempty"`        // 必须; 区域
	Entrance   string `json:"entrance,omitempty"`    // 必须; 入口
	SeatNumber string `json:"seat_number,omitempty"` // 必须; 座位号
	BeginTime  int64  `json:"begin_time,omitempty"`  // 可选; 开场时间，Unix时间戳格式。
	EndTime    int64  `json:"end_time,omitempty"`    // 可选; 结束时间，Unix时间戳格式。
}

// 更新会议门票
func UpdateUser(clt *core.Client, para *UpdateUserParameters) (err error) {
	var result core.Error

	incompleteURL := "https://api.weixin.qq.com/card/meetingticket/updateuser?access_token="
	if err = clt.PostJSON(incompleteURL, para, &result); err != nil {
		return
	}

	if result.ErrCode != core.ErrCodeOK {
		err = &result
		return
	}
	return
}
