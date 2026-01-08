package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now()

	fmt.Println("Today : ", now.Format("Mon, Jan 2, 2006 at 3:04pm"))

	longTimeAgo := time.Date(2010, time.May, 18, 23, 0, 0, 0, time.UTC)

	// compare time with time.Equal()

	sameTime := longTimeAgo.Equal(now)

	fmt.Println("longTimeAgo equals to now ? : ", sameTime)

	// calculate the time different between today
	// and long time ago

	diff := now.Sub(longTimeAgo)

	// convert diff to days
	days := int(diff.Hours() / 24)

	fmt.Printf("18th May 2010 was %d days ago \n", days)

	fmt.Println("\n=== time.Equal() 测试用例 ===")

	// 示例1: 相同的时间（完全相同）
	t1 := time.Date(2019, 3, 15, 14, 30, 45, 123456789, time.UTC)
	t2 := time.Date(2019, 3, 15, 14, 30, 45, 123456789, time.UTC)
	fmt.Printf("示例1 - 相同时间: t1.Equal(t2) = %v\n", t1.Equal(t2)) // true

	// 示例2: 不同的时间
	t3 := time.Date(2019, 3, 15, 14, 30, 45, 0, time.UTC)
	t4 := time.Date(2019, 3, 15, 14, 30, 46, 0, time.UTC)
	fmt.Printf("示例2 - 不同时间（差1秒）: t3.Equal(t4) = %v\n", t3.Equal(t4)) // false

	// 示例3: 相同时间但不同时区（Equal比较的是时间点，所以应该相等）
	t5 := time.Date(2019, 3, 15, 14, 30, 45, 0, time.UTC)
	t6 := time.Date(2019, 3, 15, 22, 30, 45, 0, time.FixedZone("CST", 8*3600)) // UTC+8
	fmt.Printf("示例3 - 相同时间点不同时区: t5.Equal(t6) = %v\n", t5.Equal(t6)) // true
	fmt.Printf("    t5 (UTC): %s\n", t5.Format("2006-01-02 15:04:05 MST"))
	fmt.Printf("    t6 (CST): %s\n", t6.Format("2006-01-02 15:04:05 MST"))

	// 示例4: 纳秒级别的差异
	t7 := time.Date(2019, 3, 15, 14, 30, 45, 123456789, time.UTC)
	t8 := time.Date(2019, 3, 15, 14, 30, 45, 123456790, time.UTC)
	fmt.Printf("示例4 - 纳秒差异: t7.Equal(t8) = %v\n", t7.Equal(t8)) // false

	// 示例5: 相同日期但不同时间
	t9 := time.Date(2019, 3, 15, 14, 30, 45, 0, time.UTC)
	t10 := time.Date(2019, 3, 15, 15, 30, 45, 0, time.UTC)
	fmt.Printf("示例5 - 相同日期不同时间: t9.Equal(t10) = %v\n", t9.Equal(t10)) // false

	// 示例6: 相同时间但不同日期
	t11 := time.Date(2019, 3, 15, 14, 30, 45, 0, time.UTC)
	t12 := time.Date(2019, 3, 16, 14, 30, 45, 0, time.UTC)
	fmt.Printf("示例6 - 相同时间不同日期: t11.Equal(t12) = %v\n", t11.Equal(t12)) // false

	// 示例7: 零值时间比较
	var t13 time.Time
	var t14 time.Time
	fmt.Printf("示例7 - 零值时间: t13.Equal(t14) = %v\n", t13.Equal(t14)) // true

	// 示例8: 零值时间与非零值时间
	t15 := time.Date(2019, 3, 15, 14, 30, 45, 0, time.UTC)
	var t16 time.Time
	fmt.Printf("示例8 - 零值与非零值: t15.Equal(t16) = %v\n", t15.Equal(t16)) // false

	// 示例9: 毫秒级别的差异
	t17 := time.Date(2019, 3, 15, 14, 30, 45, 123000000, time.UTC)
	t18 := time.Date(2019, 3, 15, 14, 30, 45, 123000001, time.UTC)
	fmt.Printf("示例9 - 纳秒级差异: t17.Equal(t18) = %v\n", t17.Equal(t18)) // false

	// 示例10: 相同时间但不同Location对象（相同偏移）
	t19 := time.Date(2019, 3, 15, 14, 30, 45, 0, time.UTC)
	t20 := time.Date(2019, 3, 15, 14, 30, 45, 0, time.FixedZone("UTC", 0))
	fmt.Printf("示例10 - 相同时间相同偏移不同Location: t19.Equal(t20) = %v\n", t19.Equal(t20)) // true

	// 示例11: 闰秒边界测试
	t21 := time.Date(2019, 3, 15, 14, 30, 59, 999999999, time.UTC)
	t22 := time.Date(2019, 3, 15, 14, 31, 0, 0, time.UTC)
	fmt.Printf("示例11 - 秒边界: t21.Equal(t22) = %v\n", t21.Equal(t22)) // false
	fmt.Printf("    差值: %v\n", t22.Sub(t21))

	// 示例12: 使用Before和After配合Equal
	t23 := time.Date(2019, 3, 15, 14, 30, 45, 0, time.UTC)
	t24 := time.Date(2019, 3, 15, 14, 30, 45, 0, time.UTC)
	fmt.Printf("示例12 - 使用Before/After: t23.Before(t24)=%v, t23.After(t24)=%v, t23.Equal(t24)=%v\n",
		t23.Before(t24), t23.After(t24), t23.Equal(t24)) // false, false, true

	// 示例13: 不同时区但表示同一时刻
	t25 := time.Date(2019, 3, 15, 12, 0, 0, 0, time.UTC)
	t26 := time.Date(2019, 3, 15, 20, 0, 0, 0, time.FixedZone("CST", 8*3600))
	fmt.Printf("示例13 - 不同时区同一时刻: t25.Equal(t26) = %v\n", t25.Equal(t26)) // true
	fmt.Printf("    UTC: %s\n", t25.Format("2006-01-02 15:04:05 MST"))
	fmt.Printf("    CST: %s\n", t26.Format("2006-01-02 15:04:05 MST"))

	// 示例14: 夏令时边界测试
	// 注意：Go的time包会自动处理夏令时
	loc, _ := time.LoadLocation("America/New_York")
	t27 := time.Date(2019, 3, 10, 7, 0, 0, 0, loc) // 2019年3月10日，夏令时开始
	t28 := time.Date(2019, 3, 10, 6, 0, 0, 0, time.UTC)
	fmt.Printf("示例14 - 夏令时边界: t27.Equal(t28) = %v\n", t27.Equal(t28)) // true (可能)
	fmt.Printf("    EST: %s\n", t27.Format("2006-01-02 15:04:05 MST"))
	fmt.Printf("    UTC: %s\n", t28.Format("2006-01-02 15:04:05 MST"))

	// 示例15: Unix时间戳相同
	t29 := time.Unix(1552662645, 0)
	t30 := time.Unix(1552662645, 0)
	fmt.Printf("示例15 - Unix时间戳相同: t29.Equal(t30) = %v\n", t29.Equal(t30)) // true

	// 示例16: Unix时间戳不同
	t31 := time.Unix(1552662645, 0)
	t32 := time.Unix(1552662646, 0)
	fmt.Printf("示例16 - Unix时间戳不同: t31.Equal(t32) = %v\n", t31.Equal(t32)) // false

	fmt.Println("\n=== 总结 ===")
	fmt.Println("time.Equal() 比较的是时间点（instant），而不是时间的字符串表示")
	fmt.Println("即使时区不同，只要表示的是同一时刻，Equal() 就会返回 true")
	fmt.Println("这与 == 操作符不同，== 会比较 Location 信息")

}
