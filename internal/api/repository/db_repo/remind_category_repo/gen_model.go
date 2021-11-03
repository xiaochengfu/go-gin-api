package remind_category_repo

// RemindCategory 提醒组
//go:generate gormgen -structs RemindCategory -input .
type RemindCategory struct {
	Id         int32  //
	Name       string // 分组名称
	CreateTime int32  //
	UpdateTime int32  //
}
