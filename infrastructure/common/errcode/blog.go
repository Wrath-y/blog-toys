package errcode

var (
	BlogNetworkBusy  = &ErrCode{40000, "网络繁忙，请稍后重试", ""}
	BlogInvalidParam = &ErrCode{40001, "无效的参数", ""}
	BlogInvalidSign  = &ErrCode{40002, "无效的签名", ""}
	BlogBodyTooLarge = &ErrCode{40003, "请求消息体过大", ""}
)

var (
	FindFriendsFailed = &ErrCode{40200, "友链获取失败", ""}
	GetPixivsFailed   = &ErrCode{40201, "pixiv图片获取失败", ""}
)
