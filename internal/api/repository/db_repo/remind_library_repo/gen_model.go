package remind_library_repo

// RemindLibrary 提醒库
//go:generate gormgen -structs RemindLibrary -input .
type RemindLibrary struct {
	Id         int32  //
	Type       int32  // 提醒类型 1-单词
	UserId     int32  // 所属用户
	Body       string //
	CategoryId int32  // 分组id
	CreateTime int32  //
	UpdateTime int32  //
}
