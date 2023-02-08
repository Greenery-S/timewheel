package my_twl

import (
	"container/list"
	"time"
)

// Job 延时任务回调函数
type Job func(interface{})
// TaskData 回调函数参数类型

// TimeWheel 时间轮
type TimeWheel struct {
	interval time.Duration // 指针每隔多久往前移动一格
	ticker   *time.Ticker
	slots    []*list.List // 时间轮槽
	// key: 定时器唯一标识 value: 定时器所在的槽, 主要用于删除定时器, 不会出现并发读写，不加锁直接访问
	timer             map[interface{}]int
	currentPos        int              // 当前指针指向哪一个槽
	slotNum           int              // 槽数量
	job               Job              // 定时器回调函数
	addTaskChannel    chan Task        // 新增任务channel
	removeTaskChannel chan interface{} // 删除任务channel
	stopChannel       chan bool        // 停止定时器channel
}

// Task 延时任务
type Task struct {
	delay  time.Duration // 延迟时间
	circle int           // 时间轮需要转动几圈
	key    interface{}   // 定时器唯一标识, 用于删除定时器
	data   interface{}      // 回调函数参数
}