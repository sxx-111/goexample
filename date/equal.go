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

	fmt.Println("\n=== 更多 time.Equal() 测试用例 ===")

	// Case 1: 相同时间点（完全相同）
	t1 := time.Date(2023, time.January, 15, 10, 30, 45, 0, time.UTC)
	t2 := time.Date(2023, time.January, 15, 10, 30, 45, 0, time.UTC)
	fmt.Printf("Case 1 - 相同时间点: %v\n", t1.Equal(t2)) // true

	// Case 2: 不同时区但相同时间点
	t3 := time.Date(2023, time.January, 15, 10, 30, 45, 0, time.UTC)
	t4 := time.Date(2023, time.January, 15, 18, 30, 45, 0, time.FixedZone("CST", 8*3600)) // UTC+8
	fmt.Printf("Case 2 - 不同时区但相同时间点: %v\n", t3.Equal(t4)) // true

	// Case 3: 不同纳秒精度
	t5 := time.Date(2023, time.January, 15, 10, 30, 45, 0, time.UTC)
	t6 := time.Date(2023, time.January, 15, 10, 30, 45, 1000000, time.UTC) // 1毫秒
	fmt.Printf("Case 3 - 不同纳秒精度(相差1毫秒): %v\n", t5.Equal(t6)) // false

	// Case 4: 相同纳秒精度
	t7 := time.Date(2023, time.January, 15, 10, 30, 45, 123456789, time.UTC)
	t8 := time.Date(2023, time.January, 15, 10, 30, 45, 123456789, time.UTC)
	fmt.Printf("Case 4 - 相同纳秒精度: %v\n", t7.Equal(t8)) // true

	// Case 5: 不同日期
	t9 := time.Date(2023, time.January, 15, 10, 30, 45, 0, time.UTC)
	t10 := time.Date(2023, time.January, 16, 10, 30, 45, 0, time.UTC)
	fmt.Printf("Case 5 - 不同日期: %v\n", t9.Equal(t10)) // false

	// Case 6: 不同时间（同一天）
	t11 := time.Date(2023, time.January, 15, 10, 30, 45, 0, time.UTC)
	t12 := time.Date(2023, time.January, 15, 11, 30, 45, 0, time.UTC)
	fmt.Printf("Case 6 - 不同时间(同一天): %v\n", t11.Equal(t12)) // false

	// Case 7: 零值时间比较
	var zeroTime time.Time
	t13 := time.Date(2023, time.January, 15, 10, 30, 45, 0, time.UTC)
	fmt.Printf("Case 7 - 零值时间比较: %v\n", zeroTime.Equal(t13)) // false
	fmt.Printf("Case 7 - 零值时间与零值比较: %v\n", zeroTime.Equal(time.Time{})) // true

	// Case 8: 未来时间与现在比较
	futureTime := time.Date(2030, time.December, 31, 23, 59, 59, 0, time.UTC)
	fmt.Printf("Case 8 - 未来时间与现在比较: %v\n", futureTime.Equal(now)) // false

	// Case 9: 相同时间但不同时区（不同时间点）
	t14 := time.Date(2023, time.January, 15, 10, 30, 45, 0, time.UTC)
	t15 := time.Date(2023, time.January, 15, 10, 30, 45, 0, time.FixedZone("EST", -5*3600)) // UTC-5
	fmt.Printf("Case 9 - 相同时间但不同时区(不同时间点): %v\n", t14.Equal(t15)) // false

	// Case 10: 使用 time.Unix 创建的时间比较
	unixTime1 := time.Unix(1673782245, 0)
	unixTime2 := time.Date(2023, time.January, 15, 10, 30, 45, 0, time.UTC)
	fmt.Printf("Case 10 - Unix时间戳与Date创建的时间比较: %v\n", unixTime1.Equal(unixTime2)) // 取决于时间戳值

	// Case 11: 毫秒级精度比较
	t16 := time.Date(2023, time.January, 15, 10, 30, 45, 0, time.UTC)
	t17 := time.Date(2023, time.January, 15, 10, 30, 45, 999999, time.UTC) // 接近1毫秒
	fmt.Printf("Case 11 - 毫秒级精度比较(相差999999纳秒): %v\n", t16.Equal(t17)) // false

	// Case 12: 完全相同的时间（包括纳秒）
	t18 := time.Date(2023, time.January, 15, 10, 30, 45, 123456789, time.UTC)
	t19 := t18 // 直接赋值
	fmt.Printf("Case 12 - 直接赋值的时间比较: %v\n", t18.Equal(t19)) // true

	// Case 13: 使用 time.Parse 创建的时间比较
	parsedTime1, _ := time.Parse(time.RFC3339, "2023-01-15T10:30:45Z")
	parsedTime2 := time.Date(2023, time.January, 15, 10, 30, 45, 0, time.UTC)
	fmt.Printf("Case 13 - Parse创建的时间与Date创建的时间比较: %v\n", parsedTime1.Equal(parsedTime2)) // true

	// Case 14: 不同月份但相同日期时间
	t20 := time.Date(2023, time.January, 15, 10, 30, 45, 0, time.UTC)
	t21 := time.Date(2023, time.February, 15, 10, 30, 45, 0, time.UTC)
	fmt.Printf("Case 14 - 不同月份但相同日期时间: %v\n", t20.Equal(t21)) // false

	// Case 15: 闰年2月29日
	leapDay1 := time.Date(2020, time.February, 29, 12, 0, 0, 0, time.UTC)
	leapDay2 := time.Date(2020, time.February, 29, 12, 0, 0, 0, time.UTC)
	fmt.Printf("Case 15 - 闰年2月29日比较: %v\n", leapDay1.Equal(leapDay2)) // true

	// Case 16: 跨时区比较（北京时间 vs UTC）
	beijingTime := time.Date(2023, time.January, 15, 18, 30, 45, 0, time.FixedZone("CST", 8*3600))
	utcTime := time.Date(2023, time.January, 15, 10, 30, 45, 0, time.UTC)
	fmt.Printf("Case 16 - 北京时间(18:30)与UTC时间(10:30)比较: %v\n", beijingTime.Equal(utcTime)) // true

	// Case 17: 夏令时影响（如果适用）
	// 注意：这个例子展示夏令时，但实际结果取决于具体时区规则
	t22 := time.Date(2023, time.July, 15, 10, 30, 45, 0, time.UTC)
	t23 := time.Date(2023, time.July, 15, 10, 30, 45, 0, time.UTC)
	fmt.Printf("Case 17 - 夏令时期间的时间比较: %v\n", t22.Equal(t23)) // true

	// Case 18: 时间截断比较（忽略纳秒）
	t24 := time.Date(2023, time.January, 15, 10, 30, 45, 123456789, time.UTC)
	t25 := time.Date(2023, time.January, 15, 10, 30, 45, 0, time.UTC)
	t24Truncated := t24.Truncate(time.Second)
	fmt.Printf("Case 18 - 截断纳秒后的时间比较: %v\n", t24Truncated.Equal(t25)) // true

	// Case 19: 使用 Add 方法后的时间比较
	baseTime := time.Date(2023, time.January, 15, 10, 30, 45, 0, time.UTC)
	addedTime := baseTime.Add(0) // 添加0时间
	fmt.Printf("Case 19 - Add(0)后的时间比较: %v\n", baseTime.Equal(addedTime)) // true

	// Case 20: 使用 Round 方法后的时间比较
	t26 := time.Date(2023, time.January, 15, 10, 30, 45, 500000000, time.UTC)
	t27 := t26.Round(time.Second)
	t28 := time.Date(2023, time.January, 15, 10, 30, 46, 0, time.UTC) // 四舍五入后应该是46秒
	fmt.Printf("Case 20 - Round后的时间比较: %v\n", t27.Equal(t28)) // true

}
