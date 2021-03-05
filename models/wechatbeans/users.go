package wechatbeans

type Users struct {
	Subscribe       int64   `json:"subscribe"`
	Openid          string  `json:"openid"`
	Nickname        string  `json:"nickname"`
	Sex             int64   `json:"sex"`
	Language        string  `json:"language"`
	City            string  `json:"city"`
	Province        string  `json:"province"`
	Country         string  `json:"country"`
	Headimgurl      string  `json:"headimgurl"`
	Subscribe_time  int64   `json:"subscribe_time"`
	Unionid         string  `json:"unionid"`
	Remark          string  `json:"remark"`
	Groupid         int     `json:"groupid"`
	Tagid_list      []int64 `json:"tagid_list"`
	Subscribe_scene string  `json:"subscribe_scene"`
	Qr_scene        int64   `json:"qr_scene"`
	Qr_scene_str    string  `json:"qr_scene_str"`
}
