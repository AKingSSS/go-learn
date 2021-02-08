package main // 程序的包

/*import "fmt"
import "time"*/

// 合并
import (
	"fmt"
	"time"
)

// main 函数
func main() {
	time.Sleep(3 * time.Second)
	// go 建议不加分号
	fmt.Println("hello world")
}
