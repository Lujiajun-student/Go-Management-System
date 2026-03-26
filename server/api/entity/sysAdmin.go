package entity

import "Go-Management-System/common/util"

// SysAdmin 用户模型对象
type SysAdmin struct {
	ID         uint       `gorm:"column:id;comment:'主键';primaryKey;NOT NULL" json:"id"`
	PostId     int        `gorm:"column:post_id;comment:'岗位id'" json:"post_id"`
	DeptId     int        `gorm:"column:dept_id;comment:'部门id'" json:"dept_id"`
	Username   string     `gorm:"column:username;varchar(64);comment:'用户账号';NOT NULL" json:"username"`
	Password   string     `gorm:"column:password;varchar(64);comment:'密码';NOT NULL" json:"password"`
	Nickname   string     `gorm:"column:nickname;varchar(64);comment:'昵称'" json:"nickname"`
	Status     int        `gorm:"column:status;default:1;comment:'账号启用状态：1->启用；2->禁用';NOT NULL" json:"status"`
	Icon       string     `gorm:"column:icon;varchar(500);comment:'头像'" json:"icon"`
	Email      string     `gorm:"column:email;varchar(64);comment:'邮箱'" json:"email"`
	Phone      string     `gorm:"column:phone;varchar(64);comment:'电话'" json:"phone"`
	Note       string     `gorm:"column:note;varchar(500);comment:'备注'" json:"note"`
	CreateTime util.HTime `gorm:"column:create_time;comment:'创建时间';NOT NULL" json:"create_time"`
}

func (SysAdmin) TableName() string {
	return "sys_admin"
}

// JwtAdmin 鉴权用户结构体
type JwtAdmin struct {
	Id       uint   `json:"id"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Icon     string `json:"icon"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Note     string `json:"note"`
}

// LoginDto 登陆对象
type LoginDto struct {
	Username string `json:"username" validate:"required"`          // 用户名
	Password string `json:"password" validate:"required"`          // 密码
	Image    string `json:"image" validate:"required,min=4,max=6"` // 验证码
	IdKey    string `json:"id_key" validate:"required"`            // uuid
}

// AddSysAdminDto 新增用户所需参数
type AddSysAdminDto struct {
	PostId   int    `validate:"required"`
	RoleId   int    `validate:"required"`
	DeptId   int    `validate:"required"`
	Username string `validate:"required"`
	Password string `validate:"required"`
	Nickname string `validate:"required"`
	Phone    string `validate:"required"`
	Email    string `validate:"required"`
	Note     string `validate:"required"`
	Status   int    `validate:"required"`
}

// SysAdminInfo 查询用户所需参数
type SysAdminInfo struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Status   int    `json:"status"`
	PostId   int    `json:"postId"`
	DeptId   int    `json:"deptId"`
	RoleId   int    `json:"roleId"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Note     string `json:"note"`
}

// UpdateSysAdminDto 修改用户所需参数
type UpdateSysAdminDto struct {
	Id       uint
	PostId   int
	DeptId   int
	RoleId   uint
	Username string
	Nickname string
	Phone    string
	Email    string
	Note     string
	Status   int
}

// SysAdminIdDto 删除用户所需参数
type SysAdminIdDto struct {
	Id uint `json:"id"`
}

// UpdateSysAdminStatusDto 设置用户状态所需参数
type UpdateSysAdminStatusDto struct {
	Id     uint `json:"id"`
	Status int  `json:"status"`
}

// ResetSysAdminPasswordDto 重置密码
type ResetSysAdminPasswordDto struct {
	Id       uint
	Password string
}

// SysAdminVO 用户列表VO对象
type SysAdminVO struct {
	ID         uint       `json:"id"`
	Username   string     `json:"username"`
	Nickname   string     `json:"nickname"`
	Status     int        `json:"status"`
	PostId     int        `json:"postId"`
	DeptId     int        `json:"deptId"`
	RoleId     int        `json:"roleId"`
	PostName   string     `json:"postName"`
	DeptName   string     `json:"deptName"`
	RoleName   string     `json:"roleName"`
	Icon       string     `json:"icon"`
	Email      string     `json:"email"`
	Phone      string     `json:"phone"`
	Note       string     `json:"note"`
	CreateTime util.HTime `json:"createTime"`
}

// UpdatePersonalDto 修改个人信息所需参数
type UpdatePersonalDto struct {
	Id       uint
	Icon     string
	Username string `validate:"required"`
	Nickname string `validate:"required"`
	Phone    string `validate:"required"`
	Email    string `validate:"required"`
	Note     string `validate:"required"`
}

// UpdatePersonalPasswordDto 修改密码
type UpdatePersonalPasswordDto struct {
	Id            uint
	Password      string `validate:"required"`
	NewPassword   string `validate:"required"`
	ResetPassword string `validate:"required"`
}
