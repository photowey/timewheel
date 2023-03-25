package timewheel

import (
	"time"
)

// Task 任务模型
type Task struct {
	Delay     time.Duration // 延迟时间
	Key       any           // 任务标识
	Args      any           // 回调函数参数
	HandleFun Handler       // 指定回调函数 HandleFun(Args)
	round     int           // 圈速
}
