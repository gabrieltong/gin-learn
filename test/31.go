package test;

import (
	"fmt"
	"sync"
)
// 等待所有 goroutine 执行完毕
// 进入死锁
func main() {
	var wg sync.WaitGroup
	var a uint8;
	var b byte;
	a = 255;
	b = a;
	fmt.Printf("%v", a)
	fmt.Printf("%v", b)
	workerCount := 10000 *10
 	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go doIt(i, &wg)
	}

	wg.Wait()
 	fmt.Println("all done!")
}

func doIt(workerID int, wg *sync.WaitGroup) {
	w := wg
	fmt.Printf("[%v] is running\n", workerID)
	defer w.Done()
	fmt.Printf("[%v] is done\n", workerID)
}