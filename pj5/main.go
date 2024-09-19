package main

import (
	"fmt"
	"time"

	"github.com/go-vgo/robotgo"
)

func main() {
    fmt.Println("Starting mouse position tracking...")

    for {
        // 获取当前鼠标位置
        x, y := robotgo.GetMousePos()
        fmt.Printf("Mouse Position: x=%d, y=%d\n", x, y)

        // 延迟 500 毫秒
        time.Sleep(500 * time.Millisecond)
    }
}
