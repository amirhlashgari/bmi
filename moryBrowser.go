/**
برای شبیه‌سازی تاریخچه مرورگر با استفاده از لیست پیوندی دوطرفه، باید به ساختارهای داده‌ای BrowserHistory و متدهای مربوط به آن توجه کنیم. در اینجا، یک لیست پیوندی دوطرفه برای مدیریت تاریخچه مرورگر پیاده‌سازی می‌شود، که شامل عملیات بازدید از صفحات جدید، رفتن به صفحه قبلی یا بعدی، پاک کردن تاریخچه و گرفتن URL صفحه فعلی است.

کد زیر این مفاهیم را پیاده‌سازی می‌کند:

*/

package main

import (
	"errors"
	"fmt"
)

// Node structure for doubly linked list
type Node struct {
	url  string
	prev *Node
	next *Node
}

// BrowserHistory structure that holds the browsing history
type BrowserHistory struct {
	current *Node
	head    *Node
	tail    *Node
}

// NewBrowserHistory initializes a new BrowserHistory
func NewBrowserHistory() *BrowserHistory {
	return &BrowserHistory{
		current: nil,
		head:    nil,
		tail:    nil,
	}
}

// VisitNewPage adds a new page to the browser history
func (bh *BrowserHistory) VisitNewPage(url string) {
	newNode := &Node{url: url}

	// If there is no current page, set this as the first page
	if bh.current == nil {
		bh.head = newNode
		bh.tail = newNode
		bh.current = newNode
		return
	}

	// If we're in the middle of the history, remove all forward pages
	if bh.current.next != nil {
		bh.current.next = nil
		bh.tail = bh.current
	}

	// Add new node to the end of the history
	newNode.prev = bh.current
	bh.current.next = newNode
	bh.tail = newNode
	bh.current = newNode
}

// Back moves to the previous page in the history
func (bh *BrowserHistory) Back() error {
	if bh.current == nil || bh.current.prev == nil {
		return errors.New("no previous page")
	}
	bh.current = bh.current.prev
	return nil
}

// Forward moves to the next page in the history
func (bh *BrowserHistory) Forward() error {
	if bh.current == nil || bh.current.next == nil {
		return errors.New("no next page")
	}
	bh.current = bh.current.next
	return nil
}

// ClearHistory clears the browsing history and resets to initial state
func (bh *BrowserHistory) ClearHistory() {
	bh.current = nil
	bh.head = nil
	bh.tail = nil
}

// GetCurrentURL returns the URL of the current page
func (bh *BrowserHistory) GetCurrentURL() (string, error) {
	if bh.current == nil {
		return "", errors.New("no page visited")
	}
	return bh.current.url, nil
}

func main() {
	browserHistory := NewBrowserHistory()

	// Visiting new pages
	browserHistory.VisitNewPage("https://golang.org")
	browserHistory.VisitNewPage("https://example.com")
	browserHistory.VisitNewPage("https://github.com")

	// Print current URL
	url, err := browserHistory.GetCurrentURL()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Current URL:", url)
	}

	// Move back in history
	err = browserHistory.Back()
	if err != nil {
		fmt.Println(err)
	}
	url, _ = browserHistory.GetCurrentURL()
	fmt.Println("After Back, Current URL:", url)

	// Move forward in history
	err = browserHistory.Forward()
	if err != nil {
		fmt.Println(err)
	}
	url, _ = browserHistory.GetCurrentURL()
	fmt.Println("After Forward, Current URL:", url)

	// Clear history
	browserHistory.ClearHistory()
	url, err = browserHistory.GetCurrentURL()
	if err != nil {
		fmt.Println("History cleared. No current page.")
	}
}

/**

توضیح ساختار و متدها:
ساختار Node:

هر گره در لیست پیوندی دوطرفه شامل یک URL و اشاره‌گرهایی به گره قبلی و بعدی است.
ساختار BrowserHistory:

شامل سه فیلد است:
current: گره‌ای که نشان‌دهنده صفحه فعلی است.
head: اولین گره در لیست.
tail: آخرین گره در لیست.
تابع NewBrowserHistory:

برای مقداردهی اولیه ساختار BrowserHistory.
متد VisitNewPage:

صفحه جدیدی را به انتهای لیست اضافه می‌کند.
اگر در وسط تاریخچه باشیم، تمام صفحات جلویی حذف می‌شوند.
متد Back:

به صفحه قبلی در تاریخچه می‌رود.
اگر در اولین صفحه باشیم، ارور no previous page برمی‌گرداند.
متد Forward:

به صفحه بعدی در تاریخچه می‌رود.
اگر در آخرین صفحه باشیم، ارور no next page برمی‌گرداند.
متد ClearHistory:

تمام تاریخچه را پاک کرده و به حالت اولیه بازنشانی می‌کند.
متد GetCurrentURL:

URL صفحه فعلی را برمی‌گرداند.
اگر هیچ صفحه‌ای بازدید نشده باشد، ارور no page visited برمی‌گرداند.
نحوه اجرا:
این کد را در یک فایل main.go کپی کنید و با اجرای آن می‌توانید تاریخچه مرورگر را شبیه‌سازی کنید.

*/
