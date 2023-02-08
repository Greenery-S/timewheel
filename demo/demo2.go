package demo

import (
	"fmt"
	"time"
	my_twl "timewheel/myTimeWheel"
)
// 失败,有问题
// 尝试构建多级时间轮,暂时不考虑circle参数
func Demo2(){
	//10秒轮
	tw2:=my_twl.New(1*time.Second,10,func(i interface{}) {
		fmt.Println("<---10 r--->",i)
		fmt.Println(i)
	})
	//100秒轮
	tw1:=my_twl.New(10*time.Second,10,func(i interface{}) {
		fmt.Println("<---100r--->",i)
		go tw2.AddTimer(5*time.Second,2,i)
	})
	go tw1.AddTimer(25*time.Second,1,"uid:1")

	time.Sleep(30 * time.Second)

	// 停止时间轮
	tw2.Stop()
	tw1.Stop()
}