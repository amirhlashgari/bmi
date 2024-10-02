package main

import (
	"fmt"
	"time"
)

type Restaurant struct {
	name           string
	avgPreparation int
}

type Order struct {
	restaurantName string
	startTime      time.Time
	endTime        time.Time
}

func parseTime(t string) time.Time {
	parsed, _ := time.Parse("15:04", t)
	return parsed
}

func timeDifference(start, end time.Time) int {
	if end.Before(start) {
		end = end.Add(24 * time.Hour)
	}
	return int(end.Sub(start).Minutes())
}

func main() {
	var n int
	fmt.Scanln(&n)

	restaurants := make(map[string]int)
	for i := 0; i < n; i++ {
		var name string
		var avgTime int
		// خواندن نام رستوران و زمان میانگین آماده‌سازی
		fmt.Scanln(&name, &avgTime)
		restaurants[name] = avgTime
	}

	var m int
	fmt.Scanln(&m)

	var orders []Order
	for i := 0; i < m; i++ {
		var first, second, time1, time2, customer string
		fmt.Scanln(&first, &time1, &time2, &second)

		if _, ok := restaurants[first]; ok {
			// سفارش نوع اول: از رستوران شروع می‌شود
			start := parseTime(time1)
			end := parseTime(time2)
			orders = append(orders, Order{first, start, end})
		} else {
			// سفارش نوع دوم: از مشتری شروع می‌شود
			start := parseTime(time2)
			end := parseTime(time1)
			orders = append(orders, Order{second, start, end})
		}
	}

	var query string
	fmt.Scanln(&query)

	if query == "overlap" {
		// محاسبه بیشترین تداخل سفارش‌ها
		maxOverlap := 0
		for i := 0; i < len(orders); i++ {
			overlapCount := 1
			for j := 0; j < len(orders); j++ {
				if i != j && orders[i].startTime.Before(orders[j].endTime) && orders[j].startTime.Before(orders[i].endTime) {
					overlapCount++
				}
			}
			if overlapCount > maxOverlap {
				maxOverlap = overlapCount
			}
		}
		fmt.Println(maxOverlap)
	} else {
		// محاسبه سفارش‌های دیرکرد برای یک رستوران خاص
		lateOrders := 0
		for _, order := range orders {
			if order.restaurantName == query {
				// زمان تایید سفارش + زمان آماده‌سازی
				expectedEnd := order.startTime.Add(time.Duration(restaurants[query]+10) * time.Minute)
				actualPreparationTime := timeDifference(order.startTime, order.endTime)
				if actualPreparationTime > restaurants[query] {
					lateOrders++
				}
			}
		}
		fmt.Println(lateOrders)
	}
}

/*

برای حل این مسئله باید از قابلیت‌های مدیریت زمان و استفاده از داده‌های ساختاریافته در زبان Go بهره ببریم. در این راه‌حل، از بسته‌ی time برای مدیریت زمان استفاده می‌شود و برای خواندن و پردازش ورودی از پکیج fmt و strings استفاده خواهیم کرد.

کد زیر به ترتیب مراحل حل مسئله را شامل می‌شود:

خواندن تعداد رستوران‌ها و میانگین زمان آماده‌سازی هر کدام.
پردازش سفارش‌ها و محاسبه‌ی زمان‌های لازم بر اساس قوانین صورت‌مسئله.
محاسبه‌ی تعداد سفارش‌های دیرهنگام برای یک رستوران خاص یا بیشترین تداخل سفارش‌ها.
در نهایت با توجه به نوع درخواست (رستوران یا overlap) خروجی مناسب چاپ می‌شود.











توضیحات کلی کد:
تعریف ساختارها:
از دو ساختار Restaurant و Order برای ذخیره اطلاعات رستوران‌ها و سفارشات استفاده شده است.

تابع parseTime:
این تابع برای تبدیل رشته‌ی زمان به یک مقدار time.Time استفاده می‌شود تا بتوانیم به راحتی محاسبات مربوط به زمان را انجام دهیم.

محاسبه تداخل زمان‌ها:
برای محاسبه تداخل، سفارشات را با هم مقایسه می‌کنیم تا تعداد سفارشاتی که زمان آن‌ها با هم تداخل دارند را پیدا کنیم.

محاسبه سفارش‌های دیرهنگام:
برای محاسبه دیرکرد، زمان آماده‌سازی واقعی را با میانگین زمان آماده‌سازی مقایسه می‌کنیم و سفارش‌هایی که زمان بیشتری برای آماده‌سازی گرفته‌اند را به عنوان دیرکرد محسوب می‌کنیم.

نکات:
برای محاسبه تداخل زمان‌ها، بررسی می‌شود که آیا زمان شروع یک سفارش قبل از پایان سفارش دیگر است یا خیر.
برای مدیریت سفارشاتی که از نیمه‌شب عبور می‌کنند، زمان‌ها به‌درستی محاسبه می‌شوند.

*/
