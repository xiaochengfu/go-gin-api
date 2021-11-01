package remind_plan_repo

// RemindPlan 提醒计划
//go:generate gormgen -structs RemindPlan -input .
type RemindPlan struct {
	Id         int32  //
	Name       string //
	UserId     int32  //
	CategoryId int32  // 组id
	LibraryId  string //
	Time       string // 时间时分秒
	Status     int32  // 状态 1-开启 2-关闭 3-删除
	Type       int32  // 提醒策略类型 1-指定时间 2-每x几分提醒一次
	CircleType int32  // 多个提醒项时，周期内 1-随机 2-顺序
	CreateTime int32  //
	UpdateTime int32  //
}
