package my_twl

import (
	"time"
)

// AddTimer 添加定时器 key为定时器唯一标识
func (tw *TimeWheel) AddTimer(delay time.Duration, key interface{}, data interface{}) {
	if delay < 0 {
		return
	}
	tw.addTaskChannel <- Task{delay: delay, key: key, data: data}
}

// 新增任务到链表中
func (tw *TimeWheel) addTask(task *Task) {
	pos, circle := tw.getPositionAndCircle(task.delay)
	task.circle = circle

	tw.slots[pos].PushBack(task)

	if task.key != nil { //把任务注册都tw上,value为所在时间槽
		tw.timer[task.key] = pos
	}
}

// 获取定时器在槽中的位置, 时间轮需要转动的圈数
func (tw *TimeWheel) getPositionAndCircle(d time.Duration) (pos int, circle int) {
	delaySeconds := int(d.Seconds())
	intervalSeconds := int(tw.interval.Seconds())
	circle = int(delaySeconds / intervalSeconds / tw.slotNum)
	pos = int(tw.currentPos+delaySeconds/intervalSeconds) % tw.slotNum

	return
}