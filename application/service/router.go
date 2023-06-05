package application

import (
	Auth "shortTermStrategy/domain/auth_aggregate"
	Short "shortTermStrategy/domain/short_game"
	Task "shortTermStrategy/domain/timed_task_aggregate"
)

type ServiceApplication struct {
	//登录及菜单权限
	AuthService Auth.AuthService
	//定时任务
	TimedTask Task.TimedTaskService
	//涨停板数据
	ShortService Short.ShortService
}
