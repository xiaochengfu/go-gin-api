package sys_user_repo

// SysUser 用户表
//go:generate gormgen -structs SysUser -input .
type SysUser struct {
	Id            int32  // 用户id
	Username      string // 用户名（邮箱前缀）
	Password      string // 用户登录密码
	Nickname      string // 用户真实姓名
	Phone         string // 用户手机号
	Email         string // 用户邮箱
	Avatar        string // 用户头像 json数据 {"url":"","path":""}
	Sex           int32  // 性别 1-男 2-女 0-保密
	CreateBy      string // 创建人
	Status        int32  // 用户状态 1-正常 2-停用
	LastLoginTime int32  // 最近登录时间
	UpdateTime    int32  // 修改时间
	CreateTime    int32  // 用户创建时间
	DeleteTime    int32  // 删除时间
}
