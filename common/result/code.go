package result

// Codes 定义状态
type Codes struct {
	SUCCESS                                 uint
	FAILED                                  uint
	NOAUTH                                  uint
	AUTHFORMATERROR                         uint
	Message                                 map[uint]string
	INVALIDTOKEN                            uint
	MissingLoginParameter                   uint
	VerificationCodeHasExpired              uint
	CAPTCHANOTTRUE                          uint
	PASSWORDNOTTRUE                         uint
	STATUSISENABLE                          uint
	ROLENAMEALREADYEXISTS                   uint
	MENUISEXIST                             uint
	DELSYSMENUFAILED                        uint
	DEPTISEXIST                             uint
	DEPTISDISTRIBUTE                        uint
	POSTALREADYEXISTS                       uint
	USERNAMEALREADYEXISTS                   uint
	MissingNewAdminParameter                uint
	FileUploadError                         uint
	MissingModificationOfPersonalParameters uint
}

// ApiCode 状态码
var ApiCode = &Codes{
	SUCCESS:                                 200,
	FAILED:                                  501,
	NOAUTH:                                  403,
	AUTHFORMATERROR:                         405,
	INVALIDTOKEN:                            406,
	MissingLoginParameter:                   407,
	VerificationCodeHasExpired:              408,
	CAPTCHANOTTRUE:                          409,
	PASSWORDNOTTRUE:                         410,
	STATUSISENABLE:                          411,
	ROLENAMEALREADYEXISTS:                   412,
	MENUISEXIST:                             413,
	DELSYSMENUFAILED:                        414,
	DEPTISEXIST:                             415,
	DEPTISDISTRIBUTE:                        416,
	POSTALREADYEXISTS:                       417,
	USERNAMEALREADYEXISTS:                   418,
	MissingNewAdminParameter:                419,
	FileUploadError:                         427,
	MissingModificationOfPersonalParameters: 428,
}

// init 初始化状态信息
func init() {
	ApiCode.Message = map[uint]string{
		ApiCode.SUCCESS:                                 "成功",
		ApiCode.FAILED:                                  "失败",
		ApiCode.NOAUTH:                                  "未授权",
		ApiCode.AUTHFORMATERROR:                         "授权格式错误",
		ApiCode.INVALIDTOKEN:                            "无效的Token",
		ApiCode.MissingLoginParameter:                   "缺少登录参数",
		ApiCode.VerificationCodeHasExpired:              "验证码已失效",
		ApiCode.CAPTCHANOTTRUE:                          "验证码不正确",
		ApiCode.PASSWORDNOTTRUE:                         "密码不正确",
		ApiCode.STATUSISENABLE:                          "您的账号被停用",
		ApiCode.ROLENAMEALREADYEXISTS:                   "角色名称已存在，重新输入",
		ApiCode.MENUISEXIST:                             "菜单已存在，重新输入",
		ApiCode.DELSYSMENUFAILED:                        "菜单已分配",
		ApiCode.DEPTISEXIST:                             "部门名称已存在",
		ApiCode.DEPTISDISTRIBUTE:                        "部门已分配，不能删除",
		ApiCode.POSTALREADYEXISTS:                       "岗位名称已存在",
		ApiCode.USERNAMEALREADYEXISTS:                   "用户名已存在",
		ApiCode.MissingNewAdminParameter:                "缺少新增参数",
		ApiCode.FileUploadError:                         "文件上传错误",
		ApiCode.MissingModificationOfPersonalParameters: "缺少必要信息",
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
