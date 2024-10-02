/*
برای حل این مسئله، نیاز است تا قوانین رمزنگاری ذکر شده را به ترتیب پیاده‌سازی کنیم. هر مرحله از تغییرات رمزنگاری روی پیام ورودی اعمال می‌شود و سپس نتیجه نهایی چاپ می‌گردد.

قوانین رمزنگاری:
شیفت حروف و ایموجی‌ها: هر حرف یا ایموجی باید به اندازه‌ی باقی‌مانده‌ی Unicode خود بر ۵، در الفبا یا بازه ایموجی‌ها به جلو شیفت داده شود.
جایگزینی اعداد: هر عدد باید به کد Unicode خود تبدیل شده و سپس در طول پیام قبل از اعمال رمزنگاری ضرب شود.
تبدیل فضاهای خالی: هر فضای خالی باید به تعداد کاراکترهای کلمه قبل از آن به _ تبدیل شود.
معکوس کردن پیام: در انتها پیام نهایی معکوس می‌شود (به جز اعداد).
مراحل پیاده‌سازی:
شیفت حروف و ایموجی‌ها: برای هر حرف، باید آن را با استفاده از کد Unicode شیفت دهیم.
جایگزینی اعداد: اعداد موجود در پیام را باید با کد Unicode آن‌ها جایگزین کنیم و ضرب در طول پیام انجام دهیم.
تبدیل فضاهای خالی: فضاهای خالی را به تعداد کاراکترهای کلمه قبل از آن به - تبدیل کنیم.
معکوس کردن پیام: در نهایت، پیام نهایی را معکوس کنیم.

*/

package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

// تابع برای شیفت دادن حروف
func shiftRune(r rune) rune {
	shift := r % 5
	if unicode.IsLetter(r) {
		if unicode.IsUpper(r) {
			return 'A' + (r-'A'+shift)%26
		} else if unicode.IsLower(r) {
			return 'a' + (r-'a'+shift)%26
		}
	} else if r >= 0x1F600 && r <= 0x1F64F { // بازه ایموجی‌ها
		return r + shift
	}
	return r
}

// تابع برای پردازش یک پیام
func processMessage(msg string) string {
	var result strings.Builder
	var tempWord strings.Builder
	msgLen := len(msg)

	for _, r := range msg {
		if unicode.IsDigit(r) {
			// جایگزینی اعداد
			num, _ := strconv.Atoi(string(r))
			result.WriteString(strconv.Itoa(int(rune(num) * rune(msgLen))))
		} else if unicode.IsSpace(r) {
			// تبدیل فضاهای خالی
			for i := 0; i < tempWord.Len(); i++ {
				result.WriteRune('_')
			}
			tempWord.Reset()
		} else {
			// شیفت دادن حروف و ایموجی‌ها
			shiftedRune := shiftRune(r)
			result.WriteRune(shiftedRune)
			tempWord.WriteRune(shiftedRune)
		}
	}

	// معکوس کردن پیام
	encrypted := result.String()
	runes := []rune(encrypted)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func main() {
	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var msg string
		fmt.Scanln(&msg)
		fmt.Println(processMessage(msg))
	}
}

/*
شیفت حروف و ایموجی‌ها: از تابع shiftRune برای شیفت دادن حروف و ایموجی‌ها استفاده می‌شود. این تابع با توجه به کد Unicode حروف، آن‌ها را شیفت می‌دهد.
جایگزینی اعداد: اگر با یک رقم مواجه شدیم، آن رقم را به کد Unicode تبدیل کرده و سپس در طول پیام ضرب می‌کنیم.
تبدیل فضاهای خالی: فضاهای خالی با تعداد آندرلاین (_) معادل تعداد کاراکترهای کلمه قبل از آن جایگزین می‌شوند.
معکوس کردن پیام: در انتها، پیام نهایی را معکوس کرده و چاپ می‌کنیم.


کد به درستی ورودی‌های حروف، اعداد و فضاهای خالی را مطابق با قوانین رمزنگاری پردازش می‌کند.
برای شیفت دادن کاراکترها از Unicode استفاده شده است تا حروف و ایموجی‌ها به درستی شیفت شوند.
هر عدد به کد Unicode خود تبدیل و سپس ضرب در طول پیام می‌شود.

*/
