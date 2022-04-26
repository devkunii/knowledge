07 string型
===========

## 1. string型の定義

以下の特徴を持つ

1. 可変長の文字列処理が可能

2. 基本データ型と同一の演算子を利用可能



### basic_stringクラスとstring型

* string型は、`basic_string`テンプレートクラスをchar型で具現化したもの

* `typedef`で宣言される

```cpp
#include <string>
```

* `charT`：どのデータ型を処理するか指定

* `traits`：文字の特性処理(文字の型、大小比較、配列操作など)を行うためのテンプレート構造体

* `Allocator`：記憶割り当てと解放を行うクラス

```cpp
template<class charT, class traits = char_traits<charT>,
class Allocator = allocator<charT> > clas basic_string { };

typedef basic_string<char> string;
typedef basic_string<wchar_t> wstring;
```



## 2. string型の初期化と演算子

### string型の初期化

```cpp
char cc[] = "abcde";
string s1;
string s2 = "";
string s3 = cc;
string s4 = s3;
string s5(5, 'Z');
string s6 = "LITERAL";
string s7(s6, 2, 3);
string s8(cc+2, 3);
```



### string型の配列

* 初期化要素数が配列長よりも少ないとき、足りない部分は`""`で初期化される

* 配列長は初期化要素で自動判定される

```cpp
string s1[2];
string s2[4] = { "aaa", "bbb", "ccc", "ddd" };
string s3[4] = { "eee", "fff", "ggg" };
string s4[] = { "iii", "jjj", "kkk", "lll" };
```



### string型対応の演算子

* `+`：連結

* `[]`：添字処理

* `<<`、`>>`：入出力処理

```cpp
=, +, +=, ==, !=, <, <=, >=, [], <<, >>
```



### 演算の作用対象

* string型、char[]型、Cスタイル文字列に対応している

```cpp
char cc[80] = "aaaa";
string s1, s2 = "bbbb";

s1 = "cccc";
s1 = cc;
s1 = s2;

s1 = s2 + "dddd";
s1 = s2 + cc;
s1 = s2 + s2;
```

* 非string要素だけの演算はできない

* 最初の演算の結果が新たにstring型になるので、その後に非string要素が連続しても正しく処理される

```cpp
// s1 = "aaa" + "bbb";
// s1 = "aaa" + "bbb" + s2;
s1 = s2 + "aaa" + "bbb"
```



### char型配列への代入

* string型オブジェクトにchar[]型文字列を代入可能

* char[]型配列に、string型オブジェクトを直接代入することはできない

  * ただし、`c_str()`メンバ関数を用いることで、string文字列からCスタイル文字列に変換すれば正しく処理が可能

```cpp
s1 = cc;
// strcpy(cc, s1);
strcpy(cc, s1.c_str());
```



### 添字処理

* 特定の文字だけを取り出し・設定することが可能

* 添字が実際の文字列のサイズを超えた位置を指定した場合、動作は未定義となる

```cpp
string s = "ABCD";
char ch;
ch = s[2];
s[2] = 'Z';
```



### 文字の代入

```cpp
string s;
s = 'A';
s += 'B';
s  = s + 'C';

// string t = 'A';
```



## 3. stringクラスのメンバ関数1

### size_typeとnpos定数

* stringクラスのメンバ関数は、文字サイズや文字位置を指定するために、以下の記法を用いる

  * size_type型は`typedef`で宣言される

  * この型で表現できる最大値は、`npos`：「文字列<`npos`」という関係を持つ

```cpp
size_type n;
```

```cpp
s1.assign(s2, 4, 6);
s1.assign(s2, 4, string::npos);
```



### string文字列とCスタイル文字列

* Cスタイル文字列は、終端に`\0`を持つ

* string型文字列は、終端マークはない

* string文字列をCスタイル文字列として使用するためのサービス関数

  * `c_str()`：末尾に`\0`を追加したchar配列を内部に作成、その先頭アドレスを返す

  * `data()`：char配列を内部に作成し、その先頭アドレスを返す(`\0`の追加はない)

  * `copy()`：指定文字数をchar配列にコピー(`\0`の追加はない)

```cpp
#include <iostream>
#include <string>
#include <cstring>
using namespace std;

int main()
{
  char cc[80];
  string ss = "ABCD";
  const char* p;

  p = ss.c_str();
  strcpy(cc, ss.c_str());
  cout << p << " " << cc << endl;

  p = ss.data();
  cout << *p << endl;*

  ss.copy(cc, 3);
  cc[3] = '\0';
  cout << cc << endl;

  ss.copy(cc, string::npos);
  cc[ss.length()] = '\0';
  cout << cc << endl;
  return 0;
}
```



## 4. stringクラスのメンバ関数2

### stringオブジェクトのサイズ

* `length()`、`size()`：現在の文字列の長さを返す

```cpp
string s = "ABCDEF";
cout << a.length()  << " " << s.size() << endl;
```

* `max_size()`：stringオブジェクトとして格納可能な最大サイズ

```cpp
cout << s.max_size() << endl;
```

* `capacity()`：実際に確保されているメモリサイズ

```cpp
cout << s.capacity() << endl;
```

* `reserve()`：指定値を保証する領域サイズを設定する

```cpp
s.reserve(80);
cout << s.capacity() << endl;
```

* `resize()`：文字列長を指定する。現在よりも長くなる場合には、後ろに空白文字を連結する

```cpp
s.resize(10, ' ');
cout << s.size() << endl;
cout << "[" << s << "]\n";
```



### stringクラスのメンバ関数の用法

> 前準備

```cpp
string str_s = "ABCDE";
char chr_s[] = "abcde";
string s = "jklmno", s2 = "pqrstu";
char ch, cc[80];
string::size_type siz;
int n;
```

* `at()`：添字計算と同じ処理をする

  * 有効範囲外を指定すると、out_of_range例外がスローされる

  * 添字処理では例外スローはないので注意

```cpp
ch = s.at(2);
s.at(2) = 'X';
```

* `empty()`：〜が空なら

```cpp
if (s.empty())
```

* `assign()`：stringオブジェクトに、char*文字列、string文字列のそれらの部分文字列、同一連続文字列を設定

  * 文字数指定の時、実体以上の大きな数値を指定すると、「すべての文字列」という意味になる

```cpp
s.assign(chr_s);
s.assign(str_s);
s.assign(chr_s+1, 3);
s.assign(str_s, 1, 3);
s.assign(str_s, 1, string::npos);
s.assign(4, 'Y');
```

* `append()`：stringオブジェクトに、char*文字列、string文字列、それらの部分文字列、同一連続文字を追加

```cpp
s.append(chr_s);
s.append(str_s);
s.append(chr_s+2, 4);
s.append(str_s, 2, 4);
s.append(str_s, 2, string::npos);
s.append(4, 'Y');
```

* `insert()`：stringオブジェクトに、char*文字列、string文字列、それらの部分文字列、同一連続文字を挿入する

```cpp
s.insert(3, chr_s);
s.insert(3, str_s);
s.insert(3, chr_s+2, 4);
s.insert(3, str_s, 2, 4);
s.insert(3, str_s, 2, string::npos);
s.insert(3, 4, 'Y');
```

* `replace()`：stringオブジェクトの指定位置から指定文字数を、指定された文字列と置換する

```cpp
s.replace(2, 3, "12345");
s.replace(2, 3, s2, 1, 2);
```

* `substr()`：部分文字列を取り出す

```cpp
s = str_s.substr(3, 4);
```

* `find()`：文字列を探索し、見つかった位置を返す。見つからないときは`npos`を返す。検索開始位置を指定可能

```cpp
n = s.find("ab");
n = s.find("ab", 5);
```

* `rfind()`：後ろから文字列を探索、見つかった位置を返す。見つからないときは`npos`を返す。

```cpp
n = s.rfind("ab");
n = s.rfind("ab", 5);
```

* `compare()`：文字列を比較する

  * 戻り値

    * 負：左が小さい

    * 0：等しい

    * 正：右が大きい

```cpp
if (s.compare(s2) < 0)
if (s.compare(cc) < 0)
if (s.compare("DDD") < 0)
if (s.compare(1, 3, s2) > 0)
if (s.compare(1, 3, s2, 3, 3) > 0)
```

* `erase()`、`clear()`：部分文字列または全文字列を削除する

```cpp
s.erase(3, 4);
s.erase();
s.clear();
```



| 版  | 年月日     |
| --- | ---------- |
| 1st | 2020/08/12 |
