// package main
//
// import (
// 	"fmt"
// 	"strconv"
// 	"time"
// )
//
// func main() {
// 	fmt.Println("開始")
//
// 	ch1 := make(chan int)
// 	ch2 := make(chan string)
// 	chend := make(chan struct{}) // 終了通知用のチャネル
//
// 	// チャネルを送信するゴルーチン
// 	go func(chint chan<- int, chstr chan<- string, end chan<- struct{}) {
//
// 		for i := 0; i < 10; i++ {
// 			// 偶数回はint型チャネル、奇数回はstring型チャネルを送信する
// 			if i%2 == 0 {
// 				fmt.Println("ch1へ送信")
// 				chint <- i
// 			} else {
// 				fmt.Println("ch2へ送信")
// 				chstr <- "test" + strconv.Itoa(i)
// 			}
// 		}
//
// 		time.Sleep(1 * time.Second)
// 		close(end) // クローズして通知
//
// 	}(ch1, ch2, chend)
//
// 	// 受信用の無限ループ
// 	for {
// 		select {
// 		case val := <-ch1:
// 			fmt.Println("ch1から受信：", val)
// 		case str := <-ch2:
// 			fmt.Println("ch2から受信：", str)
// 		case <-chend:
// 			fmt.Println("終了")
// 			return
// 		}
// 	}
// }

// package main
//
// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )
//
// func parallel(wg *sync.WaitGroup) {
//
// 	fmt.Println("A")
// 	time.Sleep(1 * time.Millisecond)
// 	fmt.Println("B")
// 	time.Sleep(1 * time.Millisecond)
// 	fmt.Println("C")
//
// 	// 終了を通知する
// 	wg.Done()
// }
//
// func main() {
// 	// WaitGroup構造体を初期化
// 	wg := new(sync.WaitGroup)
// 	// 3つのゴルーチンを同時に実行します
// 	for i := 0; i < 3; i++ {
// 		wg.Add(1)// WaitGroupに、ゴルーチンを1つずつ追加
// 		go parallel(wg)
// 	}
// 	// wg.Addで追加したすべてゴルーチンが、Doneで終了通知されるまで待機
// 	wg.Wait()
// }

package main

import (
	"fmt"
	"sync"
	"time"
)

func parallel(wg *sync.WaitGroup, mt *sync.Mutex) {
	// ミューテックスを使用してロックします。
	mt.Lock()
	// 関数終了後にアンロックします。
	defer mt.Unlock()

	fmt.Println("A")
	time.Sleep(1 * time.Millisecond)
	fmt.Println("B")
	time.Sleep(1 * time.Millisecond)
	fmt.Println("C")

	// 終了を通知する
	wg.Done()
}

func main() {
	// WaitGroup構造体を初期化
	wg := new(sync.WaitGroup)
	// Mutex構造体を初期化
	mt := new(sync.Mutex)
	// 3つのゴルーチンを同時に実行します
	for i := 0; i < 3; i++ {
		wg.Add(1) // WaitGroupに、ゴルーチンを1つずつ追加
		go parallel(wg, mt)
	}
	// wg.Addで追加したすべてゴルーチンが、Doneで終了通知されるまで待機
	wg.Wait()
}
