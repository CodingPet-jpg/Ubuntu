package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()                                         // 获取当前时间
	beforeDay := t.AddDate(0, 0, -1).Format("2006-01-02")   // 三个参数分别是年月日，此处获取的是前一天的日期
	beforeMonth := t.AddDate(0, -1, 0).Format("2006-01-02") // 前一个月的日期
	beforeYear := t.AddDate(-1, 0, 0).Format("2006-01-02")  // 去年的当天日期
	fmt.Println(beforeYear)
	fmt.Println(beforeDay)
	fmt.Println(beforeMonth)
}
