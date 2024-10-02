package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ساختار برای ذخیره مختصات
type Coordinates struct {
	x int
	y int
}

func main() {
	locations := make(map[string]Coordinates)

	scanner := bufio.NewScanner(os.Stdin)

	// خواندن مختصات start از ورودی
	scanner.Scan()
	startLine := scanner.Text()

	// استخراج اطلاعات start
	startParts := strings.Fields(startLine)
	startX, _ := strconv.Atoi(strings.TrimPrefix(startParts[1], "x="))
	startY, _ := strconv.Atoi(strings.TrimPrefix(startParts[2], "y="))

	// ذخیره مختصات start
	locations[startParts[0]] = Coordinates{startX, startY}

	// خواندن مکان‌های جدید
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)

		locationName := parts[0]
		fromLocation := parts[2]
		deltaXStr := parts[3][2:]
		deltaYStr := parts[4][2:]

		// دریافت مختصات مکان fromLocation
		fromCoords := locations[fromLocation]

		// تبدیل تغییرات x و y به عدد
		deltaX, _ := strconv.Atoi(deltaXStr)
		deltaY, _ := strconv.Atoi(deltaYStr)

		// محاسبه مختصات جدید
		newCoords := Coordinates{
			x: fromCoords.x + deltaX,
			y: fromCoords.y + deltaY,
		}

		// ذخیره مختصات جدید
		locations[locationName] = newCoords

		// چاپ مختصات جدید
		fmt.Printf("%s x=%d y=%d\n", locationName, newCoords.x, newCoords.y)
	}
}
