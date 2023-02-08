package demo

import (
	"fmt"
	"time"
	my_twl "timewheel/myTimeWheel"
)
//单时间轮
//执行结果为: 1 2 5 4
func Demo1() {
	// 初始化时间轮
	// 第一个参数为tick刻度, 即时间轮多久转动一次
	// 第二个参数为时间轮槽slot数量
	// 第三个参数为回调函数
	tw := my_twl.New(1*time.Second, 5, func(data interface{}) {
		// do something
		fmt.Println("Hello!", data)
	})

	// 启动时间轮
	tw.Start()

	// 添加定时器
	// 第一个参数为延迟时间
	// 第二个参数为定时器唯一标识, 删除定时器需传递此参数
	// 第三个参数为用户自定义数据, 此参数将会传递给回调函数, 类型为interface{}
	tw.AddTimer(1*time.Second, 1, map[string]int{"uid": 1})
	tw.AddTimer(2*time.Second, 2, map[string]int{"uid": 2})
	tw.AddTimer(3*time.Second, 3, map[string]int{"uid": 3})
	tw.AddTimer(14*time.Second, 4, map[string]int{"uid": 4})
	tw.AddTimer(5*time.Second, 5, map[string]int{"uid": 5})
	// tw.AddTimer(6 * time.Second, 6, map[string]int{"uid" : 6})

	// 删除定时器, 参数为添加定时器传递的唯一标识
	tw.RemoveTimer(3)

	time.Sleep(30 * time.Second)

	// 停止时间轮
	tw.Stop()
	//select{}
}
