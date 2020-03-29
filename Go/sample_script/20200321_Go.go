// package main
//
// import (
// 	"fmt"
// )
//
// func main() {
// 	// スライスを作成
// 	s1 := []int{1, 2, 3, 4, 5}// スライスの内容を指定し、配列と同時にスライス作成
// 	fmt.Println("s1=", s1)
//
// 	// スライスに要素を追加
// 	s2 := append(s1, 6, 7)
// 	fmt.Println("s2=", s2)
//
// 	// スライスにスライスを追加
// 	s3 := append(s1, s2...)
// 	fmt.Println("s3=", s3)
// }

// package main
//
// import (
// 	"fmt"
// )
//
// func main() {
// 	// スライスを作成
// 	src1 := []int{1, 2}
// 	dest := []int{97, 98, 99}
//
// 	// destへsrc1の内容をすべてコピー
// 	count := copy(dest, src1)
// 	fmt.Println("copy count=", count)
// 	fmt.Println(dest)
//
// 	fmt.Println()
//
// 	src2 := []int{3}
// 	// destの3つめのインデックスに、src2をコピー
// 	count = copy(dest[2:], src2)
// 	fmt.Println("copy count=", count)
// 	fmt.Println(dest)
// }

// package main
//
// import (
// 	"fmt"
// )
//
// func main() {
// 	// string型スライスを作成
// 	s1 := make([]string, 5, 10)
// 	fmt.Println("len=", len(s1))
// 	fmt.Println("cap=", cap(s1))
//
// 	fmt.Println()
// 	// キャパシティを省略
// 	s2 := make([]string, 5)
// 	fmt.Println("len=", len(s2))
// 	fmt.Println("cap=", cap(s2))
// }

// package main
//
// import (
// 	"fmt"
// )
//
// func main() {
// 	// マップの作成
// 	currencies := make(map[string]string)
//
// 	// キーを指定して値を格納
// 	currencies["日本"] = "JPY"
// 	currencies["USA"] = "USD"
// 	currencies["EU"] = "EUR"
// 	currencies["中国"] = "CNY"
//
// 	// キーを指定して値を取得
// 	fmt.Println(currencies["日本"])
// 	fmt.Println(currencies["中国"])
//
// 	fmt.Println("==すべて取得=======")
// 	// for rangeを使用して、マップ内の値すべてを取得
// 	for country, currency := range currencies {
// 		// 1つめの戻り値にキー、2つめの戻り値に値が返ります
// 		fmt.Println(country, currency)
// 	}
// 	fmt.Println("==ここまで========")
// }

package main

import "fmt"

var testVar int = showAndReturn("Declare", 1)

func init() {
    testVar = showAndReturn("Init", 2)
}

func main() {
    testVar = showAndReturn("Main", 3)
}

func showAndReturn(timing string, i int) int {
    fmt.Println(timing, ":", i)
    return i
}
