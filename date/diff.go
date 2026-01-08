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

	fmt.Println("\n=== 时间格式化示例 ===")

	// 示例11: RFC3339格式（ISO 8601）
	now := time.Date(2019, 3, 15, 14, 30, 45, 123456789, time.UTC)
	fmt.Printf("示例11 - RFC3339格式: %s\n", now.Format(time.RFC3339)) // 2019-03-15T14:30:45Z

	// 示例12: RFC3339Nano格式
	fmt.Printf("示例12 - RFC3339Nano格式: %s\n", now.Format(time.RFC3339Nano)) // 2019-03-15T14:30:45.123456789Z

	// 示例13: RFC822格式
	fmt.Printf("示例13 - RFC822格式: %s\n", now.Format(time.RFC822)) // 15 Mar 19 14:30 UTC

	// 示例14: RFC822Z格式（带时区）
	fmt.Printf("示例14 - RFC822Z格式: %s\n", now.Format(time.RFC822Z)) // 15 Mar 19 14:30 +0000

	// 示例15: RFC1123格式
	fmt.Printf("示例15 - RFC1123格式: %s\n", now.Format(time.RFC1123)) // Fri, 15 Mar 2019 14:30:45 UTC

	// 示例16: RFC1123Z格式（带时区）
	fmt.Printf("示例16 - RFC1123Z格式: %s\n", now.Format(time.RFC1123Z)) // Fri, 15 Mar 2019 14:30:45 +0000

	// 示例17: 自定义格式 - 年月日
	fmt.Printf("示例17 - 自定义格式(YYYY-MM-DD): %s\n", now.Format("2006-01-02")) // 2019-03-15

	// 示例18: 自定义格式 - 年月日时分秒
	fmt.Printf("示例18 - 自定义格式(YYYY-MM-DD HH:MM:SS): %s\n", now.Format("2006-01-02 15:04:05")) // 2019-03-15 14:30:45

	// 示例19: 自定义格式 - 中文日期
	fmt.Printf("示例19 - 自定义格式(中文): %s\n", now.Format("2006年01月02日 15:04:05")) // 2019年03月15日 14:30:45

	// 示例20: 自定义格式 - 12小时制
	fmt.Printf("示例20 - 自定义格式(12小时制): %s\n", now.Format("2006-01-02 03:04:05 PM")) // 2019-03-15 02:30:45 PM

	// 示例21: 自定义格式 - 仅时间
	fmt.Printf("示例21 - 自定义格式(仅时间): %s\n", now.Format("15:04:05")) // 14:30:45

	// 示例22: 自定义格式 - 带毫秒
	fmt.Printf("示例22 - 自定义格式(带毫秒): %s\n", now.Format("2006-01-02 15:04:05.000")) // 2019-03-15 14:30:45.123

	// 示例23: 自定义格式 - 带微秒
	fmt.Printf("示例23 - 自定义格式(带微秒): %s\n", now.Format("2006-01-02 15:04:05.000000")) // 2019-03-15 14:30:45.123456

	// 示例24: 自定义格式 - 带纳秒
	fmt.Printf("示例24 - 自定义格式(带纳秒): %s\n", now.Format("2006-01-02 15:04:05.000000000")) // 2019-03-15 14:30:45.123456789

	// 示例25: Unix时间戳格式化
	timestamp := time.Date(2019, 3, 15, 14, 30, 45, 0, time.UTC)
	fmt.Printf("示例25 - Unix时间戳: %d\n", timestamp.Unix()) // 1552662645

	// 示例26: UnixNano时间戳格式化
	fmt.Printf("示例26 - UnixNano时间戳: %d\n", timestamp.UnixNano()) // 1552662645000000000

	// 示例27: 不同时区格式化
	beijingTime := timestamp.In(time.FixedZone("CST", 8*3600))
	fmt.Printf("示例27 - 北京时间: %s\n", beijingTime.Format("2006-01-02 15:04:05 MST")) // 2019-03-15 22:30:45 CST

	// 示例28: 时区信息
	fmt.Printf("示例28 - 时区信息: %s\n", beijingTime.Format("2006-01-02 15:04:05 -0700")) // 2019-03-15 22:30:45 +0800

	// 示例29: 星期格式化
	fmt.Printf("示例29 - 星期: %s\n", now.Format("Monday"))  // Friday
	fmt.Printf("示例29 - 星期(缩写): %s\n", now.Format("Mon")) // Fri

	// 示例30: 月份格式化
	fmt.Printf("示例30 - 月份: %s\n", now.Format("January")) // March
	fmt.Printf("示例30 - 月份(缩写): %s\n", now.Format("Jan")) // Mar

	// 示例31: 组合格式 - 完整日期时间
	fmt.Printf("示例31 - 完整格式: %s\n", now.Format("2006-01-02 15:04:05 Monday")) // 2019-03-15 14:30:45 Friday

	// 示例32: Kitchen格式
	fmt.Printf("示例32 - Kitchen格式: %s\n", now.Format(time.Kitchen)) // 2:30PM

	// 示例33: Stamp格式
	fmt.Printf("示例33 - Stamp格式: %s\n", now.Format(time.Stamp)) // Mar 15 14:30:45

	// 示例34: StampMilli格式
	fmt.Printf("示例34 - StampMilli格式: %s\n", now.Format(time.StampMilli)) // Mar 15 14:30:45.123

	// 示例35: StampMicro格式
	fmt.Printf("示例35 - StampMicro格式: %s\n", now.Format(time.StampMicro)) // Mar 15 14:30:45.123456

	// 示例36: StampNano格式
	fmt.Printf("示例36 - StampNano格式: %s\n", now.Format(time.StampNano)) // Mar 15 14:30:45.123456789

	// 示例37: StampNano格式
	fmt.Printf("示例37 - StampNano格式: %s\n", now.Format(time.StampNano)) // Mar 15 14:30:45.123456789

	// 示例38: StampNano格式
	fmt.Printf("示例38 - StampNano格式: %s\n", now.Format(time.StampNano)) // Mar 15 14:30:45.123456789
}

func Date(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)

}
