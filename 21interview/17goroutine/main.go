package main

import (
	"context"
	"fmt"
	"time"
)

/**
假设有一个超长的切片，切片的元素类型为int，切片中的元素为乱序排列。
限时5秒，使用多个goroutine查找切片中是否存在给定值，在找到目标值或者超时后立刻结束所有goroutine的执行。

比如切片为：[23, 32, 78, 43, 76, 65, 345, 762, …… 915, 86]，
查找的目标值为345，如果切片中存在目标值程序输出:"Found it!"并且立即取消仍在执行查找任务的goroutine。
如果在超时时间未找到目标值程序输出:"Timeout! Not Found"，同时立即取消仍在执行查找任务的goroutine。
*/

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	target := 10
	count := 3

	resultChan := make(chan bool, 0)
	c, cancel := context.WithTimeout(context.Background(), time.Second*5)

	each := len(nums) / count
	for i := 0; i < count; i++ {
		scope := nums[each*i : each*(i+1)]
		go searchTarget(c, scope, target, resultChan)
	}

	select {
	case <-resultChan:
		fmt.Println("Found it!")
		cancel()
	case <-c.Done():
	}
}

func searchTarget(c context.Context, nums []int, target int, resultChan chan bool) {
	var i int
	for {
		select {
		case <-c.Done():
			fmt.Println("Timeout! Not Found")
			return
		default:
		}

		if i >= len(nums) {
			return
		}

		if nums[i] == target {
			resultChan <- true
			return
		} else {
			fmt.Println(nums[i])
		}
		i++
		time.Sleep(time.Second)
	}
}
