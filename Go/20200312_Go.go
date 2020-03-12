// package main
//
// import "fmt"
//
// type Score int
//
// func main() {
//     var myScore Score = 100
//     fmt.Printf("私の点数は%d点です。\n", myScore)
// }

// package main
//
// import "fmt"
//
// func main() {
//     var readFunc func(struct{name string; meaning string}) string
//     var dict struct{name string; meaning string}
//     readFunc = readOut
//     dict.name = "コーヒー"
//     dict.meaning = "コーヒー豆から作られる黒色の飲み物"
//     fmt.Println(readFunc(dict))
// }
//
// func readOut(s struct{name string; meaning string}) string {
//     return fmt.Sprintf("「%s」 は 「%s」 という意味です", s.name, s.meaning)
// }

// package main
//
// import "fmt"

// type Dictionary struct {
//     name string
//     meaning string
// }
//
// type ReadFunc func(Dictionary) string
//
// func main() {
//     var readFunc ReadFunc
//     var dict Dictionary
//     readFunc = readOut
//     dict.name = "コーヒー"
//     dict.meaning = "コーヒー豆から作られる黒色の飲み物"
//     fmt.Println(readFunc(dict))
// }
//
// func readOut(d Dictionary) string {
//     return fmt.Sprintf("「%s」 は 「%s」 という意味です", d.name, d.meaning)
// }

// type Score int
// func (s Score) Show() { fmt.Printf("点数は%d点です\n", s) }
// func main() {
//     var myScore Score = 100
//     myScore.Show()
// }

// package main
//
// import "fmt"
//
// type Score int
//
// func main() {
//     var myScore Score = 100
//     // showInt(myScore) /* この記述方法は型が異なるので不可 */
//     showInt(int(myScore))
// }
//
// func showInt(i int) {
//     fmt.Printf("value: %d\n", i)
// }

// package main
//
// import "fmt"
//
// func main() {
//   var name string
//   name = "Mr. Go"
//   fmt.Println("Hello,", name)
// }

// package main
//
// import "fmt"
//
// func main() {
//   var name = "Mr. Go"
//   fmt.Println("Hello,", name)
// }

// package main
//
// import "fmt"
//
// func main() {
//   name := "Mr. Go"
//   fmt.Println("Hello,", name)
// }

// package main
//
// import "fmt"
//
// func main() {
//   const title = "Go言語入門"
//   fmt.Println(title);
// }

// package main
//
// import "fmt"
//
// func main() {
//   fmt.Println(1234);
//   fmt.Println(053);
//   fmt.Println(0xA3);
//   fmt.Println(0XA3);
// }

// package main
//
// import "fmt"
//
// func main() {
//   fmt.Println(3.1415);
//   fmt.Println(.25);
//   fmt.Println(12.);
//   fmt.Println(1.25e-3);
// }

// package main
//
// import "fmt"
//
// func main() {
//   fmt.Println(2i);
//   fmt.Println(3.1415i);
//   fmt.Println(1.25e-3i);
// }

// package main
//
// import "fmt"
//
// func main() {
//   fmt.Println('a');
//   fmt.Println('あ');
//   fmt.Println('\n');
//   fmt.Println('\u12AB');
// }

// package main
//
// import "fmt"
//
// func main() {
//   fmt.Println(`abc`);
//   fmt.Println(`\n`); // 改行ではなく\とnの二文字として扱われる。
//   fmt.Println(`ab
// cd`); // 前の行と合わせて、改行を含む1つの文字列として扱われる。
// }

package main

import "fmt"

func main() {
  fmt.Println("abc");
  fmt.Println("ab\ncd"); // abとcdの間に改行が挿入される
  fmt.Println("\u3042\u3044\u3046"); // 「あいう」のコードポイント表記
}
