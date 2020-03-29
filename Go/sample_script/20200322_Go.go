// package main
//
// import (
// 	"fmt"
//   "os"
// )
//
// func main() {
//   // ファイルを開く
//   file, err := os.Open("test.txt")
//   // エラー判定
//   if err != nil {
//     // 失敗
//     fmt.Println(err.Error())
//     } else {
//     // 成功
//     fmt.Println("Successful!")
//     file.Close()
// 	}
// }

// package main
//
// import (
// 	"fmt"
// 	"os"
// )
//
// type myAppError struct {
// 	msg string //エラーメッセージ
// }
//
// func (e myAppError) Error() string {
// 	// エラーメッセージを返す
// 	return e.msg
// }
//
// func testOpen() (file os.File, err error) {
// 	file, err := os.Open("test.txt")
// 	if err != nil {
// 		// エラーが発生したので、myAppError構造体の値を作成し、エラー情報として返す
// 		return nil, myAppError{"ファイルが開けません - " + err.Error()}
// 	}
// 	// 成功した場合は、error情報をnilとして返す。
// 	return file, nil
// }
//
// func main() {
// 	file, err := testOpen()
// 	if err != nil {
// 		// エラー情報が戻ってきたので、エラーとして処理する。
// 		fmt.Println("失敗しました。 - " + err.Error())
// 		os.Exit(1)
// 	}
// 	// 後略
// }

// package main
//
// func func1() {
// 	panic("Occured panic!")
// }
//
// func main() {
// 	func1()
// }

// package main
//
// import "fmt"
//
// func func1() {
// 	defer func() {
// 		fmt.Println("defer 2")
// 	}()
// 	panic("Occured panic!")
// }
//
// func main() {
// 	defer func() {
// 		fmt.Println("defer 1")
// 	}()
// 	func1()
// }

// package main
//
// func main() {
// 	arr := [...]int{1, 2, 3}
//
// 	index := 3
//
// 	arr[index] = 0
// }

// package main
//
// import "fmt"
//
// func func1(b bool) {
// 	defer func() {
// 		fmt.Println("defer start.")
//
// 		if err := recover(); err != nil {
// 		    // パニック中だった
// 			fmt.Println("recovery!")
// 		}
// 		fmt.Println("defer end.")
// 	}()
// 	if b {
// 		panic("Occure panic!")
// 	}
// }
//
// func main() {
// 	func1(false)
// 	func1(true)
// }

// package main
//
// import (
// 	"fmt"
// 	"time"
// )
//
// func main() {
// 	fmt.Println("main start.")
//
// 	fmt.Println("普通に関数を呼び出す")
// 	serialno()
//
// 	fmt.Println("ゴルーチンとして呼び出す")
// 	go serialno()
//
// 	// ゴルーチン呼び出し後、sleepする
// 	time.Sleep(1 * time.Second)
//
// 	fmt.Println("main end.")
// }
//
// func serialno() {
// 	for i := 0; i < 5; i++ {
// 		fmt.Println(i)
// 		// 1秒間sleepする
// 		time.Sleep(1 * time.Second)
// 	}
// }

// package main
//
// import (
// 	"fmt"
// )
//
// func main() {
// 	// int型チャネルの作成
// 	c := make(chan int)
//
// 	// 送信専用チャネルを受け取り、1〜10までの数値を送信する
// 	go func(s chan<- int) {
// 		for i := 0; i < 10; i++ {
// 			s <- i
// 		}
// 		close(s)
// 	}(c)
//
// 	for {
// 		// チャネルからの受信を待機
// 		val, ok := <-c
// 		if !ok {
// 			// チャネルがクローズしたので、終了する
// 			break
// 		}
// 		// 受信したデータを表示
// 		fmt.Println(val)
// 	}
// }

// package main
//
// import (
// 	"fmt"
// 	"time"
// )
//
// func main() {
// 	// キャパシティ0で、int型チャネルの作成
// 	c := make(chan int)
//
// 	// 負荷のかかる作業（5秒待機）を3回繰り返した後、通知する
// 	go func(s chan<- int) {
// 		for i := 0; i < 3; i++ {
// 			time.Sleep(5 * time.Second)
// 			fmt.Println(i+1, "回完了")
// 		}
// 		// 適当な数値を送信
// 		s <- 0
// 	}(c)
//
// 	// 受信を待機
// 	<-c
//
// 	fmt.Println("終了")
// }

package main

import (
	"fmt"
)

// 全ゴルーチン数
const goroutines = 5

func main() {
	// 共有データを保持するチャネル
	counter := make(chan int)
	// 全ゴルーチン終了通知用のチャネル
	end := make(chan bool)

	// 5個のゴルーチンを実行する
	for i := 0; i < goroutines; i++ {
		// 共有データ(counter)を受信し、インクリメントする
		go func(counter chan int) {
			// チャネルから共有データの受信
			val := <-counter
			// +1する
			val++
			fmt.Println("counter = ", val)

			if val == goroutines {
				// 最後のゴルーチンの場合は、終了通知用のチャネルへ送信
				end <- true
			}
			// +1したデータを、他のゴルーチンへ送信
			counter <- val
		}(counter)
	}
	// 初期値をチャネルに送信
	counter <- 0
	// 全ゴルーチンの終了を待機
	<-end
	fmt.Println("終了")
}
