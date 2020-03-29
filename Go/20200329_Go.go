// package main
//
// import (
//     "fmt"
//     "net"
// )
//
// func main() {
//     listener, err := net.Listen("tcp", "0.0.0.0:3000")
//     if err != nil {
//         fmt.Printf("Listen error: %s\n", err)
//         return
//     }
//     defer listener.Close()
//
//     conn, err := listener.Accept()
//     if err != nil {
//         fmt.Printf("Accept error: %s\n", err)
//         return
//     }
//     defer conn.Close()
//
//     fmt.Println("クライアントからの受信メッセージ:")
//     buf := make([]byte, 1024)
//     for {
//         n, err := conn.Read(buf)
//         if n == 0 {
//             break
//         }
//         if err != nil {
//             fmt.Printf("Read error: %s\n", err)
//         }
//         fmt.Print(string(buf[:n]))
//     }
// }

// package main
//
// import (
//     "fmt"
//     "net"
// )
//
// func main() {
//     conn, err := net.Dial("tcp", "0.0.0.0:3000")
//     if err != nil {
//         fmt.Printf("Dial error: %s\n", err)
//         return
//     }
//     defer conn.Close()
//
//     sendMsg := "Test Message.\n"
//     conn.Write([]byte(sendMsg))
// }

// package main
//
// import (
//     "fmt"
// 	"io"
//     "net/http"
// )
//
// func main() {
// 	res, err := http.Get("https://golang.org")
// 	if err != nil {
// 		fmt.Println("Request error:", err)
// 		return
// 	}
// 	defer res.Body.Close()
//
// 	buf := make([]byte, 256)
// 	for {
// 		n, err := res.Body.Read(buf)
// 		if n == 0 || err == io.EOF {
// 			break;
// 		} else if err != nil {
// 			fmt.Println("Read response body error:", err)
// 			return
// 		}
// 		fmt.Println(string(buf[:n]))
// 	}
// }

// package main
//
// import (
//     "fmt"
//     "os"
// )
//
// const BUFSIZE = 1024 // 読み込みバッファのサイズ
//
// func main() {
//     file, err := os.Open(`/Users/kunii.sotaro/work/knowledge/Go/20200329_Go.go`)
//     if err != nil {
//         // Openエラー処理
//     }
//     defer file.Close()
//
//     buf := make([]byte, BUFSIZE)
//     for {
//         n, err := file.Read(buf)
//         if n == 0 {
//             break
//         }
//         if err != nil {
//             // Readエラー処理
//             break
//         }
//
//         fmt.Print(string(buf[:n]))
//     }
// }

// package main
//
// import "os"
//
// func main() {
//     file, err := os.Create(`/Users/kunii.sotaro/work/knowledge/Go/tmp.txt`)
//     if err != nil {
//         // Openエラー処理
//     }
//     defer file.Close()
//
//     output := "testmessage"
//     file.Write(([]byte)(output))
// }

// package main
//
// import (
//     "fmt"
//     "time"
// )
//
// func main() {
//     fmt.Println(time.Now())
// }

// package main
//
// import "fmt"
// import "time"
//
// func main() {
// 	t := time.Date(2015, 9, 13, 12, 35, 42, 123456789, time.Local)
// 	fmt.Println(t)
// }

// package main
//
// import (
//     "fmt"
//     "time"
// )
//
// func main() {
//     now := time.Now()
//     fmt.Printf("フォーマット指定なし：%s\n", now.String())
//     fmt.Printf("フォーマット指定あり：%s\n", now.Format("2006/01/02 Mon 15:04:05"))
// }

// package main
//
// import (
//     "fmt"
//     "time"
// )
//
// func main() {
//     now := time.Now()
//     fmt.Printf("文字列表現：%s\n", now.String())
//     fmt.Printf("年：%d\n", now.Year())
//     fmt.Printf("月：%d\n", now.Month())
//     fmt.Printf("日：%d\n", now.Day())
//     fmt.Printf("曜日：%s\n", now.Weekday().String())
//     fmt.Printf("時：%d\n", now.Hour())
//     fmt.Printf("分：%d\n", now.Minute())
//     fmt.Printf("秒：%d\n", now.Second())
//     fmt.Printf("ナノ秒：%d\n", now.Nanosecond())
// }

// package main
//
// import "fmt"
// import "time"
//
// func main() {
// 	base := time.Date(2015, 11, 1, 0, 0, 0, 0, time.Local)
//     same := time.Date(2015, 11, 1, 0, 0, 0, 0, time.Local)
//     before := time.Date(2015, 10, 31, 23, 59, 59, 0, time.Local)
//     after := time.Date(2015, 11, 1, 0, 0, 1, 0, time.Local)
//
// 	fmt.Println(base.Equal(same))  // baseとsameが等しければtrue
// 	fmt.Println(base.Before(after)) // baseがafterよりも過去であればtrue
// 	fmt.Println(base.After(before)) // baseがbeforeよりも未来であればtrue
// }

// package main
//
// import "fmt"
// import "time"
//
// func main() {
// 	base := time.Date(2015, 11, 1, 0, 0, 0, 0, time.Local)
//     fmt.Println(base)
//     fmt.Println(base.Add(7 * time.Hour))
//     fmt.Println(base.Add(30 * time.Minute))
//     fmt.Println(base.Add(-5 * time.Second))
// }

// package main
//
// import "fmt"
// import "time"
//
// func main() {
// 	base := time.Date(2015, 11, 1, 0, 0, 0, 0, time.Local)
//     other := time.Date(2015, 10, 30, 20, 15, 32, 0, time.Local)
//     fmt.Println(base)
//     fmt.Println(other)
//     fmt.Printf("二つの時刻の差は%s秒です。\n", base.Sub(other))
// }

// package main
//
// import "fmt"
// import "strings"
//
// func main() {
//     fmt.Println(strings.Contains("abcdefg", "cde"))
//     fmt.Println(strings.Contains("abcdefg", "hij"))
// }

// package main
//
// import "fmt"
// import "strings"
//
// func main() {
//     fmt.Println(strings.HasPrefix("abcdefg", "abc"))
//     fmt.Println(strings.HasPrefix("abcdefg", "bcd"))
//     fmt.Println(strings.HasSuffix("abcdefg", "def"))
//     fmt.Println(strings.HasSuffix("abcdefg", "efg"))
// }

// package main
//
// import "fmt"
// import "strings"
//
// func main() {
//     base := "aBcDeF"
//     fmt.Println(strings.ToUpper(base))
//     fmt.Println(strings.ToLower(base))
// }

// package main
//
// import "fmt"
// import "strings"
//
// func main() {
//     base := "!!!?!??   abcdef???!!!"
//     fmt.Println(strings.TrimLeft(base, "!"))
//     fmt.Println(strings.TrimRight(base, "!?"))
//     fmt.Println(strings.Trim(base, "!? "))
// }

// package main
//
// import "fmt"
// import "strings"
//
// func main() {
//     base := "abcabcabcabc"
//     fmt.Println(strings.Replace(base, "bc", "yz", 2))
//     fmt.Println(strings.Replace(base, "abc", "xyz", -1))
// }

// package main
//
// import "fmt"
// import "strings"
//
// func main() {
//     base := "ab::cd::efg::hij"
//     fmt.Println(strings.Split(base, "::"))
//     fmt.Println(strings.SplitN(base, "::", 3))
// }

// package main
//
// import "fmt"
// import "strings"
//
// func main() {
//     array := []string{"C:", "work", "abc.txt"}
//     fmt.Println(strings.Join(array, "/"))
// }

// package main
//
// import (
// 	"fmt"
// 	"os"
// )
//
// func main() {
// 	fmt.Println(os.Args)
//
// 	if len(os.Args) != 4 {
// 		fmt.Println("指定された引数の数が間違っています。")
// 		os.Exit(1)
// 	}
//
// 	fmt.Printf("実行ファイル名: %s\n", os.Args[0])
// 	fmt.Printf("引数1: %s\n", os.Args[1])
// 	fmt.Printf("引数2: %s\n", os.Args[2])
// 	fmt.Printf("引数3: %s\n", os.Args[3])
// }

// package main
//
// import (
// 	"flag"
// 	"fmt"
// )
//
// var (
// 	intOpt  = flag.Int("i", 1234, "help message for i option")
// 	boolOpt = flag.Bool("b", false, "help message for b option")
// 	strOpt  = flag.String("s", "default", "help message for s option")
// )
//
// func main() {
//
// 	flag.Parse()
//
// 	fmt.Println(*intOpt)
// 	fmt.Println(*boolOpt)
// 	fmt.Println(*strOpt)
// }

// package main
//
// import (
//     "encoding/csv"
//     "fmt"
//     "io"
//     "strings"
// )
//
// func main() {
//     lines := []string{
//         "りんご,Apple,バラ科",
//         "みかん,Orange,ミカン科",
//         "すいか,Watermelon,ウリ科",
//     }
//     csvStr := strings.Join(lines, "\n")
//
//     r := csv.NewReader(strings.NewReader(csvStr))
//     for {
//         record, err := r.Read()
//         if err == io.EOF {
//             break
//         }
//         if err != nil {
//             // 読み込みエラー発生
//             fmt.Println("Read error: ", err)
//             break
//         }
//
//         fmt.Printf("日本語名[%s] 英語名[%s] 科分類[%s]\n", record[0], record[1], record[2])
//     }
// }

// package main
//
// import (
//     "encoding/csv"
//     "fmt"
//     "strings"
// )
//
// func main() {
//     lines := []string{
//         "りんご,Apple,バラ科",
//         "みかん,Orange,ミカン科",
//         "すいか,Watermelon,ウリ科",
//     }
//     csvStr := strings.Join(lines, "\n")
//
//     r := csv.NewReader(strings.NewReader(csvStr))
//     records, err := r.ReadAll()
//     if err != nil {
//         // 読み込みエラー発生
//         fmt.Println("Read error: ", err)
//         return
//     }
//     for _, record := range records {
//         fmt.Printf("日本語名[%s] 英語名[%s] 科分類[%s]\n", record[0], record[1], record[2])
//     }
// }
//
// package main
//
// import (
//     "bytes"
//     "encoding/csv"
//     "fmt"
// )
//
// func main() {
//     records := [][]string{
//         []string{"りんご", "Apple", "バラ科"},
//         []string{"みかん", "Orange", "ミカン科"},
//         []string{"すいか", "Watermelon", "ウリ科"},
//     }
//
//     buf := new(bytes.Buffer)
//     w := csv.NewWriter(buf)
//     for _, record := range records {
//         if err := w.Write(record); err != nil {
//             // 書き込みエラー発生
//             fmt.Println("Write error: ", err)
//             return
//         }
//         w.Flush() // Flush関数を呼び出したタイミングで実際の出力が行われる
//     }
//     fmt.Println(buf.String())
// }

// package main
//
// import (
//     "bytes"
//     "encoding/csv"
//     "fmt"
// )
//
// func main() {
//     records := [][]string{
//         []string{"りんご", "Apple", "バラ科"},
//         []string{"みかん", "Orange", "ミカン科"},
//         []string{"すいか", "Watermelon", "ウリ科"},
//     }
//
//     buf := new(bytes.Buffer)
//     w := csv.NewWriter(buf)
//     if err := w.WriteAll(records); err != nil {
//         // 書き込みエラー発生
//         fmt.Println("Write error: ", err)
//         return
//     }
//     // WriteAll関数は内部でFlushを行っているため、Flush関数の呼び出しは不要
//     fmt.Println(buf.String())
// }
//
// package main
//
// import (
//     "encoding/json"
//     "fmt"
// )
//
// type Country struct {
//     Name string              `json:"name"`
//     Prefectures []Prefecture `json:"prefectures"`
// }
//
// type Prefecture struct {
//     Name string    `json:"name"`
//     Capital string `json:"capital"`
//     Population int `json:"population"`
// }
//
// func main() {
//     jsonStr := `
// {
//   "name": "日本",
//   "prefectures": [
//     {
//       "name": "東京都",
//       "capital": "東京",
//       "population": 13482040
//     },
//     {
//       "name": "埼玉県",
//       "capital": "さいたま市",
//       "population": 7249287
//     },
//     {
//       "name": "神奈川県",
//       "capital": "横浜市",
//       "population": 9116252
//     }
//   ]
// }
// `
//     jsonBytes := ([]byte)(jsonStr)
//     data := new(Country)
//
//     if err := json.Unmarshal(jsonBytes, data); err != nil {
//         fmt.Println("JSON Unmarshal error:", err)
//         return
//     }
//
//     fmt.Println(data.Name)
//     fmt.Println(data.Prefectures[0].Name)
//     fmt.Println(data.Prefectures[1].Capital)
//     fmt.Println(data.Prefectures[2].Population)
// }

// package main
//
// import (
//     "encoding/json"
//     "fmt"
// )
//
// type Country struct {
//     Name string              `json:"name"`
//     Prefectures []Prefecture `json:"prefectures"`
// }
//
// type Prefecture struct {
//     Name string    `json:"name"`
//     Capital string `json:"capital"`
//     Population int `json:"population"`
// }
//
// func main() {
//     tokyo := Prefecture{Name:"東京都", Capital:"東京", Population:13482040}
//     saitama := Prefecture{Name:"埼玉県", Capital:"さいたま市", Population:7249287}
//     kanagawa := Prefecture{Name:"神奈川県", Capital:"横浜市", Population:9116252}
//     japan := Country{
//         Name:"日本",
//         Prefectures:[]Prefecture{tokyo, saitama, kanagawa},
//     }
//
//     jsonBytes, err := json.Marshal(japan)
//     if err != nil {
//         fmt.Println("JSON Marshal error:", err)
//         return
//     }
//
//     fmt.Println(string(jsonBytes))
// }

// package main
//
// import (
//     "bytes"
//     "encoding/json"
//     "fmt"
// )
//
// type Country struct {
//     Name string              `json:"name"`
//     Prefectures []Prefecture `json:"prefectures"`
// }
//
// type Prefecture struct {
//     Name string    `json:"name"`
//     Capital string `json:"capital"`
//     Population int `json:"population"`
// }
//
// func main() {
//     tokyo := Prefecture{Name:"東京都", Capital:"東京", Population:13482040}
//     saitama := Prefecture{Name:"埼玉県", Capital:"さいたま市", Population:7249287}
//     kanagawa := Prefecture{Name:"神奈川県", Capital:"横浜市", Population:9116252}
//     japan := Country{
//         Name:"日本",
//         Prefectures:[]Prefecture{tokyo, saitama, kanagawa},
//     }
//
//     jsonBytes, err := json.Marshal(japan)
//     if err != nil {
//         fmt.Println("JSON Marshal error:", err)
//         return
//     }
//
//     out := new(bytes.Buffer)
//     // プリフィックスなし、スペース4つでインデント
//     json.Indent(out, jsonBytes, "", "    ")
//     fmt.Println(out.String())
// }

// package main
//
// import (
//     "encoding/xml"
//     "fmt"
// )
//
// type Group struct {
//     Name      string    `xml:"name"`
//     Companies []Company `xml:"company"`
// }
//
// type Company struct {
//     Name    string  `xml:"name"`
//     Website Website `xml:"website"`
// }
//
// type Website struct {
//     Name string `xml:",chardata"`
//     URL  string `xml:"url,attr"`
// }
//
// func main() {
//     xmlStr := `
// <?xml version="1.0" encoding="UTF-8"?>
// <group>
//   <name>ABCグループ</name>
//   <company>
//     <name>ABC株式会社</name>
//     <website url="http://abc.com">ABC公式ウェブサイト</website>
//   </company>
//   <company>
//     <name>ABCソリューション株式会社</name>
//     <website url="http://abc.com/sol">ソリューション事業について</website>
//   </company>
// </group>
// `
//     data := new(Group)
//     if err := xml.Unmarshal([]byte(xmlStr), data); err != nil {
//         fmt.Println("XML Unmarshal error:", err)
//         return
//     }
//     fmt.Println(data.Name)
//     fmt.Println(data.Companies[0].Name)
//     fmt.Println(data.Companies[1].Website.Name)
//     fmt.Println(data.Companies[1].Website.URL)
// }

package main

import (
    "encoding/xml"
    "fmt"
)

type Group struct {
    Name      string    `xml:"name"`
    Companies []Company `xml:"company"`
}

type Company struct {
    Name    string  `xml:"name"`
    Website Website `xml:"website"`
}

type Website struct {
    Name string `xml:",chardata"`
    URL  string `xml:"url,attr"`
}

func main() {
    head := Company{
        Name: "ABC株式会社",
        Website: Website{Name: "ABC公式ウェブサイト", URL: "http://abc.com"},
    }
    sol := Company{
        Name: "ABCソリューション株式会社",
        Website: Website{Name: "ソリューション事業について", URL: "http://abc.com/sol"},
    }
    data := new(Group)
    data.Name = "ABCグループ"
    data.Companies = []Company{head, sol}

    // インデントなし
    noIndent, err := xml.Marshal(data)
    if err != nil {
        fmt.Println("XML Marshal error:", err)
        return
    }
    fmt.Println(string(noIndent))

    fmt.Println("----------------")

    // インデントあり
    withIndent, err := xml.MarshalIndent(data, "", "    ")
    if err != nil {
        fmt.Println("XML Marshal error:", err)
        return
    }
    fmt.Println(string(withIndent))
}
