package chan_test

import (
	"fmt"
	"time"
)

func ChanTest() {
	EndChan := 3

	// 构建了一个通道
	ch := make(chan int, 1)

	// 开启一个并发匿名函数
	go func() {

		// 从0循环到EndChan
		for idx := 0; idx <= EndChan; idx++ {

			// 发送值到通道
			ch <- idx

			time.Sleep(time.Second)
			fmt.Println("after sleep", idx)
		}
	}()

	// 遍历通道中的数据
	for data := range ch {
		// 打印通道中获取到的数据
		fmt.Println(data)
		data = 100
		fmt.Println("after changed", data)

		// 数据是EndChan的时候退出循环
		if data == EndChan {
			break
		}
	}

	for i := 0; i < 10; i++ {
		select {
		case x := <-ch:
			fmt.Println(x)
		case ch <- i:
		}
	}

	fmt.Println(ch)
}
