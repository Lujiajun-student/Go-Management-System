package result

// Codes 定义状态
type Codes struct {
	SUCCESS                  uint
	FAILED                   uint
	NOAUTHORIZATION          uint
	AUTHORIZATIONFORMATERROR uint
	Message                  map[uint]string
}

// ApiCode 状态码
var ApiCode = &Codes{
	SUCCESS:                  200,
	FAILED:                   501,
	NOAUTHORIZATION:          403,
	AUTHORIZATIONFORMATERROR: 405,
}

// init 初始化状态信息
func init() {
	ApiCode.Message = map[uint]string{
		ApiCode.SUCCESS:                  "成功",
		ApiCode.FAILED:                   "失败",
		ApiCode.NOAUTHORIZATION:          "未授权",
		ApiCode.AUTHORIZATIONFORMATERROR: "授权格式错误",
	}
}

// GetMessage 供外部使用的获取数据方法，提供状态码到状态信息的映射
func (c *Codes) GetMessage(code uint) string {
	message, ok := c.Message[code]
	if !ok {
		return ""
	}
	return message
}
