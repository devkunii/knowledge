// package main
//
// import "fmt"
//
// func main() {
//   fmt.Println(2 + 1); // 3
//   fmt.Println(2 - 1); // 1
//   fmt.Println(2 * 1); // 2
//   fmt.Println(2 / 1); // 2
//   fmt.Println(2 % 1); // 0
//   fmt.Println(2 & 1); // 0
//   fmt.Println(2 | 1); // 3
//   fmt.Println(2 ^ 1);  // 3
//   fmt.Println(2 &^ 1); // 2
//   fmt.Println(2 << 1); // 4
//   fmt.Println(2 >> 1); // 1
// }

// package main
//
// import "fmt"
//
// func main() {
//   fmt.Println(+5); // 5
//   fmt.Println(-5); // -5
//   fmt.Println(^5); // -6
//   a := 5
//   a++
//   fmt.Println(a); // 6
//   a--
//   fmt.Println(a); // 5
// }

// package main
//
// import "fmt"
//
// func main() {
//   a := 2
//   b := 1
//   fmt.Println(a==b); // false
//   fmt.Println(a!=b); // true
//   fmt.Println(a<b); // false
//   fmt.Println(a<=b); // false
//   fmt.Println(a>b); // true
//   fmt.Println(a>=b); // true
// }

// package main
//
// import "fmt"
//
// func main() {
//   a := 2
//   b := 1
//   fmt.Println(a==b&&a!=b); // false
//   fmt.Println(a==b||a!=b); // true
//   fmt.Println(!true); // false
// }

// package main
//
// import "fmt"
//
// func main() {
//   a := 2
//   fmt.Println(&a);
//   p := &a
//   fmt.Println(*p);
// }

// package main
//
// import "fmt"
//
// func main() {
//   a := 2
//   b := 1
//   b = a
//   fmt.Println(b);
// }

// package main
//
// import (
//     "fmt"
//     "time"
// )
//
// func main() {
//     hour := time.Now().Hour()
//     if hour >= 6 && hour < 12 {
//         fmt.Println("朝です。")
//     } else if hour < 19 {
//         fmt.Println("昼です。")
//     } else {
//         fmt.Println("夜です。")
//     }
// }

// package main
//
// import (
//     "fmt"
//     "time"
// )
//
// func main() {
//   if hour := time.Now().Hour(); hour >= 6 && hour < 12 {
//       fmt.Println("朝です。")
//   } else if hour < 19 {
//       fmt.Println("昼です。")
//   } else {
//       fmt.Println("夜です。")
//   }
// }

// package main
//
// import "fmt"
//
// func main() {
//   dayOfWeek := "月"
//   switch dayOfWeek {
//   case "土":
//       fmt.Println("大概は休みです。");
//   case "日":
//       fmt.Println("ほぼ間違いなく休みです。")
//   default:
//       fmt.Println("仕事です・・・。")
//   }
// }

// package main
//
// import "fmt"
//
// func main() {
//   dayOfWeek := "土"
//   switch dayOfWeek {
//   case "土":
//       fallthrough
//   case "日":
//       fmt.Println("休みです。")
//   default:
//       fmt.Println("仕事です・・・。")
//   }
// }

// package main
//
// import "fmt"
//
// func main() {
//   dayOfWeek := "月"
//   switch dayOfWeek {
//   case "土", "日":
//       fmt.Println("休みです。")
//   default:
//       fmt.Println("仕事です・・・。")
//   }
// }

// package main
//
// import (
//   "fmt"
//   "time"
// )
//
// func main() {
//   hour := time.Now().Hour()
//   switch {
//   case hour >= 6 && hour < 12:
//       fmt.Println("朝です。")
//   case  hour < 19:
//       fmt.Println("昼です。")
//   default:
//       fmt.Println("夜です。")
//   }
// }

// package main
//
// import "fmt"
//
// func main() {
//   for i := 1; i < 100; i++ {
//       if i / 2 != 0 {
//           fmt.Println(i)
//       }
//   }
// }

// package main
//
// import "fmt"
//
// func main() {
//   i := 0
//   for {
//       i++
//       if i >= 100 {
//           break
//       } else if i / 2 == 0 {
//           continue
//       }
//       fmt.Println(i)
//   }
// }

// package main
//
// import "fmt"
//
// func main() {
//   dayOfWeeks := [...]string{"月", "火", "水", "木", "金", "土", "日"}
//   for arrayIndex, dayOfWeek := range dayOfWeeks {
//       fmt.Printf("%d番目の曜日は%s曜日です。\n", arrayIndex + 1, dayOfWeek)
//   }
// }

// package main
//
// import (
// 	"fmt"
// )
//
// func main() {
// 	// int型のポインタ変数
// 	var pointer *int
// 	// int型変数
// 	var n int = 100
//
// 	// &（アドレス演算子）を使って、nのアドレスを代入
// 	pointer = &n
//
// 	fmt.Println("nのアドレス：", &n)
// 	fmt.Println("pointerの値：", pointer)
//
// 	fmt.Println("nの値：", n)
// 	// *(間接参照演算子）を利用して、ポインタの中身を取得
// 	fmt.Println("pointerの中身：", *pointer)
// }

// package main
//
// import (
// 	"fmt"
// )
//
// func main() {
// 	a, b := 10, 10
//
// 	// aはそのまま、bはアドレス演算子をつけて呼び出す
// 	called(a, &b)
//
// 	fmt.Println("値渡し：", a)
// 	fmt.Println("ポインタ渡し：", b)
// }
//
// func called(a int, b *int) {
// 	// 変数をそのまま変更
// 	a = a + 1
// 	// 変数の中身を変更
// 	*b = *b + 1
// }

// package main
//
// import (
// 	"fmt"
// )
//
// func main() {
//   // int型のメモリ割り当て
//   var p *int = new(int)
//   fmt.Println(p);
//
//   // 構造体myStruct型のメモリ割り当て
//   type myStruct struct {
//   	a int
//   	b int
//   }
//   var my *myStruct = new(myStruct)
//   fmt.Println(my);
// }

// package main
//
// import (
// 	"fmt"
// )
//
// func main() {
// 	var b bool
// 	var i int
// 	var r rune
// 	var f float64
// 	var c complex64
// 	var s string
//
// 	fmt.Println("bool =    ", b)
// 	fmt.Println("int =     ", i)
// 	fmt.Println("rune =    ", r)
// 	fmt.Println("float =   ", f)
// 	fmt.Println("complex = ", c)
// 	fmt.Println("string =  ", s)
// }

// package main
//
// import "fmt"
//
// func main() {
// FOR_LABEL:
// 	for i := 0; i < 10; i++ {
// 		switch {
// 		case i == 3:
// 			// for文からの脱出
// 			break FOR_LABEL
//
// 		default:
// 			fmt.Println(i)
// 		}
// 	}
// }

// package main
//
// func main() {
//   LABEL1:
//   for i := 0; i < 10; i++ {
//   	for j := 0; j < 10; j++ {
//   		if i == 0 && j == 5 {
//   			// 1番目のforへcontinue
//   			continue LABEL1
//   		} else if i == 1 && j == 1 {
//   			// 2番目のforへcontinue
//   			continue
//   		}
//   	}
//   }
// }

package main

func main() {
  for i := 0; i < 10; i++ {
  	if i == 2 {
  		// for文の外にあるLABELへ移動
  		goto LABEL
  	}
  }
  LABEL:
}
