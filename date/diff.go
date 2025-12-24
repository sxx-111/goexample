package main

import (
	"fmt"
	"time"
)

func main() {
	// 示例1: 相同日期
	t1 := Date(2019, 2, 26)
	t2 := Date(2019, 2, 26)
	days := t2.Sub(t1).Hours() / 24
	fmt.Printf("示例1 - 相同日期: %.0f 天\n", days) // 0

	// 示例2: 同月内不同日期
	t3 := Date(2019, 2, 1)
	t4 := Date(2019, 2, 28)
	days2 := t4.Sub(t3).Hours() / 24
	fmt.Printf("示例2 - 2019年2月1日到2月28日: %.0f 天\n", days2) // 27

	// 示例3: 跨月计算
	t5 := Date(2019, 1, 15)
	t6 := Date(2019, 3, 15)
	days3 := t6.Sub(t5).Hours() / 24
	fmt.Printf("示例3 - 2019年1月15日到3月15日: %.0f 天\n", days3) // 59

	// 示例4: 跨年计算
	t7 := Date(2018, 12, 25)
	t8 := Date(2019, 1, 5)
	days4 := t8.Sub(t7).Hours() / 24
	fmt.Printf("示例4 - 2018年12月25日到2019年1月5日: %.0f 天\n", days4) // 11

	// 示例5: 包含闰年的计算（2016年是闰年）
	t9 := Date(2016, 2, 1)
	t10 := Date(2016, 3, 1)
	days5 := t10.Sub(t9).Hours() / 24
	fmt.Printf("示例5 - 2016年2月1日到3月1日（闰年）: %.0f 天\n", days5) // 29

	// 示例6: 非闰年2月
	t11 := Date(2019, 2, 1)
	t12 := Date(2019, 3, 1)
	days6 := t12.Sub(t11).Hours() / 24
	fmt.Printf("示例6 - 2019年2月1日到3月1日（非闰年）: %.0f 天\n", days6) // 28

	// 示例7: 负数差值（t2 < t1）
	t13 := Date(2019, 5, 10)
	t14 := Date(2019, 5, 1)
	days7 := t14.Sub(t13).Hours() / 24
	fmt.Printf("示例7 - 2019年5月10日到5月1日: %.0f 天\n", days7) // -9

	// 示例8: 计算小时差
	t15 := Date(2019, 1, 1)
	t16 := Date(2019, 1, 2)
	hours := t16.Sub(t15).Hours()
	fmt.Printf("示例8 - 2019年1月1日到1月2日: %.0f 小时\n", hours) // 24

	// 示例9: 计算分钟差
	t17 := Date(2019, 1, 1)
	t18 := Date(2019, 1, 2)
	minutes := t18.Sub(t17).Minutes()
	fmt.Printf("示例9 - 2019年1月1日到1月2日: %.0f 分钟\n", minutes) // 1440

	// 示例10: 长时间跨度
	t19 := Date(2020, 1, 1)
	t20 := Date(2021, 1, 1)
	days10 := t20.Sub(t19).Hours() / 24
	fmt.Printf("示例10 - 2020年1月1日到2021年1月1日: %.0f 天\n", days10) // 366 (2020是闰年)
}

func Date(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}
