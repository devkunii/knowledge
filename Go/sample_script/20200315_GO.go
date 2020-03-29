// package main
//
// import "fmt"
//
// func main() {
//     var vector struct {
//         X int
//         Y int
//     }
//
//     vector.X = 2
//     vector.Y = 5
//     fmt.Println(vector) // {2 5}
// }

// package main
//
// import "fmt"
//
// type Vector struct {
//     X int
//     Y int
// }
//
// func main() {
//     var v Vector
//
//     v.X = 2
//     v.Y = 5
//     fmt.Println(v)
// }

// package main
//
// import "fmt"
//
// type Vector struct {
//     X int
//     Y int
// }
//
// func main() {
//     v := Vector{X: 2, Y: 5}
//     fmt.Println(v) // {2 5}
// }

// package main
//
// import "fmt"
//
// func main() {
//     DisplayHello()
// }
//
// func DisplayHello() {
//     fmt.Println("Hello!")
// }

// package main
//
// import "fmt"
//
// func main() {
//     DisplaySum(2, 5) // 7が表示される
// }
//
// func DisplaySum(left int, right int) {
//     fmt.Println(left + right)
// }

// package main
//
// import "fmt"
//
// func main() {
//     DisplaySumAll(2, 5, 8, 11) // 26が表示される
// }
//
// func DisplaySumAll(values ...int) {
//     sum := 0
//     for _, value := range values {
//         sum += value
//     }
//     fmt.Println(sum)
// }

// package main
//
// import "fmt"
//
// func main() {
//     fmt.Println(Sum(2, 5)) // 7が表示される
// }
//
// func Sum(left int, right int) int {
//     return left + right
// }

// package main
//
// import "fmt"
//
// func main() {
//     result, remainder := Div(19, 4)
//     fmt.Printf("19を4で割ると%dあまり%dです。\n", result, remainder)
// }
//
// func Div(left int, right int) (int, int) { // 複数指定の場合は戻り値を丸括弧で囲む
//     return left / right, left % right
// }

// package main
//
// import "fmt"
//
// type LoopNum int
//
// func main() {
//     var three LoopNum = 3
//     three.TimesDisplay("Hello") // 「Hello」と3回表示される
// }
//
// func (n LoopNum) TimesDisplay(s string) {
//     for i := 0; i < int(n); i++ {
//         fmt.Println(s)
//     }
// }

// package main
//
// import "fmt"
//
// type SavingBox struct {
//     money int
// }
//
// func NewBox() *SavingBox {
//     return new(SavingBox)
// }
//
// func (s *SavingBox) Income(amount int) {
//     s.money += amount
// }
//
// func (s *SavingBox) Break() int {
//     lastMoney := s.money
//     s.money = 0
//     return lastMoney
// }
//
// func main() {
//     box := NewBox()
//     box.Income(100)
//     box.Income(200)
//     box.Income(500)
//
//     fmt.Printf("貯金箱を壊したら%d円出てきました。\n", box.Break())
// }

// package main
//
// import "fmt"
//
// func main() {
//     defer fmt.Println("A")
//     fmt.Println("B")
// }

// package main
//
// import "fmt"
//
// func main() {
//     defer fmt.Println("A")
//     defer fmt.Println("B")
//     defer fmt.Println("C")
//     fmt.Println("D")
// }

// package main
//
// import (
//     "fmt"
//     "os"
// )
//
// func main() {
//     defer fmt.Println("A")
//     fmt.Println("B")
//     os.Exit(0)
// }

// package main
//
// import "fmt"
//
// type Animal interface {
//     Cry()
// }
//
// type Dog struct {}
// func (d *Dog) Cry() {
//     fmt.Println("わんわん")
// }
//
// type Cat struct {}
// func (c *Cat) Cry() {
//     fmt.Println("にゃーにゃー");
// }
//
// func MakeAnimalCry(a Animal) {
//     fmt.Println("鳴け！");
//     a.Cry();
// }
//
// func main() {
//     dog := new(Dog)
//     cat := new(Cat)
//     MakeAnimalCry(dog)
//     MakeAnimalCry(cat)
// }

// package main
//
// import "fmt"
//
// type Animal interface {
//     Cry()
// }
//
// type Dog struct {}
// func (d *Dog) Cry() {
//     fmt.Println("わんわん")
// }
//
// type Cat struct {}
// func (c *Cat) Cry() {
//     fmt.Println("にゃーにゃー");
// }
//
// func MakeSomeoneCry(someone interface{}) {
//     fmt.Println("鳴け！");
//     a, ok := someone.(Animal)
//     fmt.Println(a);
//     if !ok {
//         fmt.Println("動物では無いので鳴けません。")
//         return
//     }
//     a.Cry()
// }
//
// func main() {
//     dog := new(Dog)
//     cat := new(Cat)
//     MakeSomeoneCry(dog)
//     MakeSomeoneCry(cat)
// }

// package main
//
// import (
// 	"fmt"
// )
//
// func main() {
// 	// 配列の宣言
// 	var month [12]string
// 	month[0] = "January"
// 	month[1] = "February"
// 	month[2] = "March"
// 	month[3] = "April"
// 	month[4] = "May"
// 	month[5] = "June"
// 	month[6] = "July"
// 	month[7] = "Autust"
// 	month[8] = "September"
// 	month[9] = "October"
// 	month[10] = "Nobember"
// 	month[11] = "December"
//
// 	// 配列の長さの回数、ループして値を表示します。
// 	for i := 0; i < len(month); i++ {
// 		fmt.Printf("%d月 = %s\n", i+1, month[i])
// 	}
// }

// package main
//
// import (
// 	"fmt"
// )
//
// func main() {
// 	// スライスの元となる配列を作成
// 	num := [5]int{1, 2, 3, 4, 5}
// 	// スライス型変数の宣言
// 	var slice1 []int
//
// 	// 配列全体
// 	slice1 = num[:]
// 	fmt.Println(slice1)
//
// 	// インデックス1〜4まで
// 	slice2 := num[1:4]
// 	fmt.Println(slice2)
//
// 	// インデックス4以降
// 	slice3 := num[4:]
// 	fmt.Println(slice3)
//
// 	// インデックス4以前
// 	slice4 := num[:4]
// 	fmt.Println(slice4)
// }

// package main
//
// import (
// 	"fmt"
// )
//
// func main() {
// 	// スライスの元となる配列を作成
// 	num := [...]int{1, 2, 3, 4, 5}
// 	// 配列をスライスとしてを関数に渡す
// 	plusOne(num[:])
//
// 	fmt.Println(num)
// }
//
// // 要素内の数字すべてを＋1
// func plusOne(vals []int) {
// 	for i := 0; i < len(vals); i++ {
// 		vals[i] += 1
// 	}
// }

package main

import (
	"fmt"
)

func main() {
	// スライスの元となる配列を作成
	num := [5]int{1, 2, 3, 4, 5}

	// 配列の一部をスライス
	slice1 := num[1:4]
	fmt.Println("slice1=", slice1)
	fmt.Println("len=", len(slice1))
	fmt.Println("cap=", cap(slice1))

	// スライスの一部をスライス
	slice2 := slice1[1:4] // 長さを超過した、キャパシティ最大値まで可能
	fmt.Println("slice2=", slice2)
	fmt.Println("len=", len(slice2))
	fmt.Println("cap=", cap(slice2))

}
