06 Stringクラス
==============

## 目次

* [文字列の文字コード情報](#1文字列の文字コード情報)

* [文字列の比較](#2文字列の比較)

* [文字列の切り出し](#3文字列の切り出し)

* [文字列の変更](#4文字列の変更)

* [文字列の置換](#5文字列の置換)

* [文字列の連結](#6文字列の連結)

* [文字列の大文字・小文字への変換](#7文字列の大文字小文字への変換)

* [文字列の末尾や先頭にある空白や改行を削除](#8文字列の末尾や先頭にある空白や改行を削除)

* [文字列を逆順にする](#9文字列を逆順にする)

* [文字列の長さ](#10文字列の長さ)

* [文字列の割り付け](#11文字列の割り付け)

* [批評文字列を変換する](#12批評文字列を変換する)

* [文字列をアンパックする](#13文字列をアンパックする)

* [文字列内での検索](#14文字列内での検索)

* [次の文字列を求める](#15次の文字列を求める)

* [文字列に対する繰り返し](#16文字列に対する繰り返し)

* [他のクラスへの変換](#17他のクラスへの変換)



## 1.文字列の文字コード情報

| 文字コード    | 特徴                                                                             |
| ------------- | -------------------------------------------------------------------------------- |
| `UTF-8`       | 主に利用されている文字コード                                                     |
| `EUC-JP`      | 古いUNIX系システムで利用されていた文字コード                                     |
| `JIS`         | JISで策定された7bit文字のみを使った文字コード                                    |
| `Shift_JIS`   | マルチバイト文字とASCII文字を切り替えることなく利用できるようにした文字コード    |
| `Windows-31J` | 主にWindowsで利用されいるShift_JISの亜種。機種依存文字などが利用できる文字コード |
| `US-ASCII`    | ASCII文字だけで構成されている文字コード                                          |



### 文字列のエンコーディングの取得・変更

* `encoding`メソッド：`String`オブジェクトの文字コード情報(エンコーディング)を取得

* `encode`：引数で指定した文字コードに変換した新しいインスタンスを返す

* `encode!`：オブジェクトのエンコーディングを変更(破壊的メソッド)

* 文字列の操作は、このエンコーディング情報を元に行われるため、エンコーディングの異なる文字列を結合したり比較する際には注意が必要

```ruby
# エンコーディングの取得
>> a = "abc"
=> "abc"
>> a.encoding
=> #<Encoding:UTF-8>

# エンコーディングの変更(encode)
>> b = a.encode("EUC-JP")
=> "abc"
>> b.encoding
=> #<Encoding:EUC-JP>

# エンコーディングの変更(encode!)
>> a = "ルビー"
=> "ルビー"
>> a.encode!("EUC-JP")
=> "\x{A5EB}\x{A5D3}\x{A1BC}"
>> a.encoding
=> #<Encoding:EUC-JP>
```



## 2.文字列の比較

* `==`：大文字小文字、全角半角を含めて同じかどうかを比較

* `>`、`>=`、`<`、`<=`：文字列同士をアスキーコードで比較(`true`or`false`)

* `<=>`、`casecmp`：比較した結果を-1、0、1の整数値で返す(UFO演算子と同じ働き)
  →`<=>`は大文字小文字を区別するが、`casecmp`は区別しない

```ruby
>> "abc" == "abc"
=> true
>> "abc" == "ABC"  # 大文字小文字は区別
=> false
>> "a" < "b"       # 辞書順では、aの方が小さい
=> true
>> "A" > "a"       # なぜかわからないので、あとで調べる
=> false
>> "aa" < "b"      # 辞書順では、aaの方が小さい
=> true
>> "a" <=> "b"     # aの方が小さい(-1)
=> -1
```



### 文字列と数値の比較

* `==`メソッド：異なるクラスのオブジェクトと比較できる。

> 型の自動変換は行われないので、`false`

```ruby
>> '100' == 100
=> false
>> '100' >= 100
ArgumentError: comparison of String with 100 failed
```



### 文字列比較のときのエンコーディング

* `==`、`eql?`メソッド：両者のエンコーディングが等しく、文字列自身のバイト列表現が等しい場合のみ`true`を返す

* ただし、両者のエンコーディングがASCII互換でASCII文字しか含まない場合は、エンコーディングが異なる場合もバイト表現が一致すれば`true`を返す

```ruby
# ASCII互換
>> a = "abc"
=> "abc"
>> b = a.encode("EUC-JP")
=> "abc"
>> b.encoding
=> #<Encoding:EUC-JP>
>> a == b
=> true

>> a = "ルビー"
=> "ルビー"
>> b = a.encode("EUC-JP")
=> "\x{A5EB}\x{A5D3}\x{A1BC}"
>> a == b
=> false
```



## 3.文字列の切り出し

* `[]`、`slice`：文字列から指定された一部分を切り出すメソッド

* `slice!`：文字列から指定された一部分を切り出し、返した文字を元の文字列から取り除く

* `split`：文字列や正規表現を使って文字列を分割

```ruby
# []
>> a = 'abcdef'
=> "abcdef"
>> puts a[2]
c
=> nil

# slice
>> a.slice(2)
=> "c"

# slice!
>> a.slice!(2)
=> "c"
>> puts a
abdef
=> nil

# split
>> 'abcdefg'.split('d')
=> ["abc", "efg"]
>> 'abcdef'.split(/d/)
=> ["abc", "ef"]
>> "abcde\nfghij".split(/\n/)  # ダブルクオートに注意
=> ["abcde", "fghij"]
```

* `String#slice(nth, len)`：文字列の`nth`目から`len`文字の文字列を作って返す

```ruby
# 解答
>> string = "test code"
=> "test code"
>> string.slice(0,4)
=> "test"
>> p string
"test code"
=> "test code"

# 破壊的メソッドの場合
>> string = "test code"
=> "test code"
>> string.slice!(0,4)
=> "test"
>> p string    # 0から4番目(test )が削除
" code"
=> " code"
```

* `";|:"`は文字列中にないので、そのまま(正規表現で指定すれば、分割される)

```ruby
>> str = "1;2:3;4"
=> "1;2:3;4"
>> p str.split(";|:")
["1;2:3;4"]
=> ["1;2:3;4"]

# 例
>> p str.split(/;|:/)
["1", "2", "3", "4"]
=> ["1", "2", "3", "4"]
```

* `String#split`メソッドは引数で指定した特定の文字列を区切り文字として、文字列から配列を生成する

  * また、第二引数で生成される配列の要素数を指定することもできます。

```ruby
>> str = "a,b,c,d"
=> "a,b,c,d"
>> p str.split(/,/, 2)
["a", "b,c,d"]
=> ["a", "b,c,d"]
```



### 数値を指定した場合

* 数値を指定した場合は、数値の位置にある文字を返す

* 負の数値の場合は、末尾から数えてた位置の文字を返す

```ruby
>> 'abcdefg'[2]       # 0,1,2・・・"c"
=> "c"
>> 'abcdefg'.slice(2) # 0,1,2・・・"c"
=> "c"
>> 'abcdefg'[-2]      # 6,5・・・"f"
=> "f"
>> a = 'abcdefg'
=> "abcdefg"
>> a.slice!(2)        # 0,1,2・・・"c"でカット
=> "c"
>> puts a
abdefg
=> nil
```



### 範囲指定の場合

* 範囲として`Range`オブジェクトを指定した場合は、該当する範囲の文字列を返す

* 範囲指定は、開始位置と長さで指定可能。

  * 開始位置が範囲外の場合：`nil`

  * 開始位置が負の場合：末尾から数えた位置となる

  * 長さが文字列より長い場合には、可能な部分までを返す

```ruby
# Rangeオブジェクトで指定
>> puts a
abdefg
=> nil
>> 'abcdefg'[1..3]
=> "bcd"
>> 'abcdefg'.slice(1..3)
=> "bcd"
>> a = 'abcdefg'
=> "abcdefg"
>> a.slice!(1..3)
=> "bcd"
>> puts a
aefg
=> nil

# 開始位置と長さで指定
>> 'abcdefg'[1,3]
=> "bcd"
>> 'abcdefg'.slice(1,3)
=> "bcd"
>> 'abcdefg'[-2,3]  # 最後から2文字目から、3文字(どう頑張っても最後から2文字分しか出力されない)
=> "fg"
>> a = 'abcdefg'
=> "abcdefg"
>> a.slice!(1,3)
=> "bcd"
>> puts a
aefg
=> nil
```



### 文字列で指定

* 元の文字列に含まれていればその部分を、含まれていなければ`nil`を返す

```ruby
>> 'abcdefg'["bc"]
=> "bc"
>> 'abcdefg'.slice("bc")
=> "bc"
>> 'abcdefg'["bd"]  # "bd"で表される文字は含まれない
=> nil
>> a = 'abcdefg'
=> "abcdefg"
>> a.slice!("bc")
=> "bc"
>> puts a
adefg
=> nil
```



### 正規表現で指定

* マッチした部分があればその部分を、マッチしなければ`nil`を返す

```ruby
>> 'abcdefg'[/bc/]
=> "bc"
>> 'abcdefg'.slice(/bc/)
=> "bc"
>> a = 'abcdefg'
=> "abcdefg"
>> a.slice!(/bc/)
=> "bc"
>> puts a
adefg
=> nil
```

* `slice`は、最初にマッチしたものを返す

```ruby
>> p "hogepiyohogehoge".slice(/o../)
"oge"
=> "oge"
```



## 4.文字列の変更

* `[]=`、`insert`メソッド：文字列の一部分を変更する

  * 範囲や位置の指定ができ、該当する部分を新しい文字列で置換することができる

```ruby
>> a = 'abcdefg'
=> "abcdefg"
>> a[1..3] = 'xyz'
=> "xyz"
>> puts a
axyzefg
=> nil

# インデックスが足りなくても、挿入する文字自体はそのまま挿入される。指定していない文字も消される。
>> a = 'abcdefgh'
=> "abcdefgh"
>> a[1..2] = 'xyz'
=> "xyz"
>> a
=> "axyzdefgh"
```



## 5.文字列の置換

* 定数を置換する場合、警告が発生しない。オブジェクトIDが変更されない為



### `sub`メソッド・`gsub`メソッド

* `sub`メソッド：指定したパターンにマッチした最初の部分を、特定の文字列に置換する

* `gsub`メソッド：マッチした全ての部分を置換する

```ruby
>> a = 'abcdefg-abcdefg'
=> "abcdefg-abcdefg"
>> a.sub(/abc/, 'xyz')
=> "xyzdefg-abcdefg"
>> a.sub(/abc/, 'xyz')
=> "xyzdefg-abcdefg"
```

* ブロックも取ることができ、その場合にはブロックにマッチした部分が渡され、ブロックの実行結果と置換される

```ruby
>> a = 'abcdefg-abcdefg'
=> "abcdefg-abcdefg"
>> a.sub(/abc/) {|str| 'xyz'}
=> "xyzdefg-abcdefg"
>> a.gsub(/abc/) {|str| 'xyz'}
=> "xyzdefg-xyzdefg"
```



### `tr`メソッド・`tr_s`メソッド

* `tr`メソッド：指定したパターンに含まれる文字を検索し、それを特定の文字列やパターンに合わせて置換する

* `tr_s`メソッド：`tr`メソッドの機能に加えて、重複する文字を1文字に圧縮

```ruby
>> a = 'aabbccddeeffgg'
=> "aabbccddeeffgg"
>> a.tr('a-c', 'A-C')
=> "AABBCCddeeffgg"
>> a.tr_s('a-c', 'A-C')
=> "ABCddeeffgg"
```



### `delete`メソッド

* 指定したパターンに含まれる文字を、元の文字列から削除

* パターンを複数指定すると、全てのパターンに含まれる文字列のみ削除する

```ruby
>> a = 'aabbccddeeffgg'
=> "aabbccddeeffgg"
>> a.delete('a-f', 'd-g')
=> "aabbccgg"
```

### 例題

```ruby
>> puts "0123456789-".delete("^13-56-")
13456-
=> nil

>> puts "0123456789".delete("0-58-")
679
=> nil

>> puts "Ruby on Rails".delete("Rails")  # R、a、i、l、sの5文字を指定して取り除く
uby on
=> nil
```

* `String#delete`は、引数に含まれる文字を文字列から取り除く

* `^`で始まる文字列は、その文字列以外を削除する。(例：`[^1]`の場合、`1`以外を削除する)

* `[-]`は文字の範囲を示す。(例：`[1-3]`の場合、1,2,3を意味する)

* `-`の両端に文字列がある場合は範囲指定をしていることになる

* `"0-5"`で0から5までの数字を取り除きますが、続く`"8-"`では範囲指定とは見なされず、8と-を削除する

* 正規表現とは別物として考える！！



### `squeeze`メソッド

* 指定した文字が複数並んでいた場合に、一文字に圧縮

```ruby
>> a = 'aabbccddeeffgg'
=> "aabbccddeeffgg"
>> a.squeeze('a-e')
=> "abcdeffgg"
```



### `replace`メソッド

* 引数の文字列で自分自身の内容を置き換える

```ruby
>> a = 'abc'
=> "abc"
>> a.object_id
=> 70159960401620
>> a.replace('xyz')
=> "xyz"
>> a.object_id
=> 70159960401620    # 置換されただけなので、オブジェクトIDは変更されない
>> puts a
xyz
=> nil
```



## 6.文字列の連結

* `+`：文字列を結合した新しいオブジェクトを生成

* `*`：文字列の内容を指定した数値の数だけ、繰り返した文字列を返す

* `<<`、`concat`：元のオブジェクトの内容に文字列を追記

  > 破壊的メソッド

* 異なるエンコーディングの文字列を結合するとき、互換性がない場合はエラー

* 互換性がなくても、文字列がASCII文字列のみからなるASCII互換であれば結合可能

```ruby
>> a = 'abc'
=> "abc"
>> a.object_id
=> 70263014985560   # 元々のオブジェクト
>> a << 'def'
=> "abcdef"
>> a.object_id
=> 70263014985560   # 同じオブジェクト
>> a = a + 'ghi'
=> "abcdefghi"
>> a.object_id
=> 70263018562880   # 新しいオブジェクト
>> 'abc' * 2
=> "abcabc"
```


* `String`に`append`メソッドはない

  * 文字列を結合するには、`String<<`を用います。

```ruby
>> a = "Ruby"
=> "Ruby"
>> b = " on Rails"
=> " on Rails"
>> a.append b
NoMethodError: undefined method ｀append｀ for "Ruby":String

# 文法を変えた場合
>> a << b
=> "Ruby on Rails"
>> a.reverse       # 破壊的メソッドではない
=> "sliaR no ybuR"
>> p a
"Ruby on Rails"
=> "Ruby on Rails"
>> p a.index("R", 1)   # 右から1番目(u)から、最初にRが見つかる場所(8番目)
8
=> 8
```



## 7.文字列の大文字・小文字への変換

* `capitalize`：文字列の先頭にある半角英字を大文字に、残りの半角英字を小文字にして返す

* `downcase`、`upcase`：それぞれ半角英字を小文字化、大文字化して返す

* `swapcase`：半角英字の小文字を大文字に、大文字を小文字に変更して返す

```ruby
>> a = 'abcDEF'
=> "abcDEF"
>> a.capitalize
=> "Abcdef"
>> a.upcase
=> "ABCDEF"
>> a.downcase
=> "abcdef"
>> a.swapcase
=> "ABCdef"
```



## 8.文字列の末尾や先頭にある空白や改行を削除

* `chomp`：末尾から引数で指定する改行コードを取り除いた文字列を返す。

  * 指定がない場合は、`"\r"`、`"\r\n"`、`"\n"`の全てを改行コードと見なして取り除く

* `strip`、`lstrip`、`rstrip`：先頭と末尾、先頭、末尾にある空白文字を取り除いた文字列を返す

  * `"\t"`、`"\r"`、`"\n"`、`"\f"`、`"\v"`など

* `chop`：末尾の文字を取り除いた文字列を返す

```ruby
>> a = "\nabcdef\n"
=> "abcdef\n"
>> a.chomp         # 末尾の改行コードを取り除く
=> "\nabcdef"
>> a.strip         # 改行コードを取り除く
=> "abcdef"
>> a.lstrip        # 左の改行コードを取り除く
=> "abcdef\n"
>> a.rstrip        # 右の改行コードを取り除く
=> "\nabcdef"
>> a.chop          # 末尾の文字を取り除く
=> "\nabcdef"
>> a.chop.chop     # 末尾から2文字取り除く
=> "\nabcde"
>> a.chomp.chomp   # 改行コードのみ取り除く
=> "\nabcdef"
```


* `str.chop`は末尾の文字を取り除く

  * ただし、文字列の末尾が`"\r\n"`であれば、2文字とも取り除く

  * 破壊的メソッドではないので、`self`は影響を受けない

```ruby
>> str = "Liberty Fish   \r\n"
=> "Liberty Fish   \r\n"
>> str.chop
=> "Liberty Fish   "
>> p str
"Liberty Fish   \r\n"
=> "Liberty Fish   \r\n"
>> str.chop!
=> "Liberty Fish   "
>> p str
"Liberty Fish   "
=> "Liberty Fish   "
```



## 9.文字列を逆順にする

* `reverse`：並び順をバイト単位で逆にする

```ruby
>> a = "abcdef"
=> "abcdef"
>> a.reverse
=> "fedcba"
```



## 10.文字列の長さ

* `length`、`size`：文字数を返す

* `count`：指定されたパターンに該当する文字がいくつあるかを返す

* `empty?`：文字列が空かどうかを返す

* `bytesize`：バイト数が返る

```ruby
# 文字列の長さ
>> a = "abcdef"
=> "abcdef"
>> a.length
=> 6
>> a.count('a-c')
=> 3
>> a.empty?
=> false
>> "".empty?
=> true

# 文字列のバイト数
>> a = "ルビー"
=> "ルビー"
>> a.length
=> 3
>> a.bytesize
=> 9
```



## 11.文字列の割り付け

* 文字列をある長さの文字列の中に割り付ける

* メソッドを呼び出すときに、以下のものを指定できる

  * 返す文字列の長さ

  * 割り付ける余白に使用する文字

* `center`：中央に割り付ける

* `ljust`：左側に割り付ける

* `rjust`：右側に割り付ける

```ruby
>> a = "abc"
=> "abc"
>> a.center(20)
=> "        abc         "
>> a.center(20, "*")
=> "********abc*********"
>> a.ljust(20)
=> "abc                 "
>> a.rjust(20, "-")
=> "-----------------abc"
```



## 12.批評文字列を変換する

* `dump`：文字列の中にある改行コード(`"\n"`)や、タブ文字(`"\t"`)の非表示文字列をバックスラッシュ記法に置き換えた文字列を返す

  * 例)改行コード(`"\n"`)、タブ文字(`"\t"`)

```ruby
>> a = "abc\tdef\tghi\n"
=> "abc\tdef\tghi\n"
>> puts a
abc	def	ghi
=> nil
>> puts a.dump
"abc\tdef\tghi\n"
=> nil
```



## 13.文字列をアンパックする

* `unpack`： **Array#pack** メソッドでパックされた文字列を、指定したテンプレートにし違ってアンパックする

  * 例)MINEエンコードされた文字列のデコードや、バイナリデータから数値への変換などに利用

```ruby
# MINEエンコードされた文字列をアンパックする
>> '440r440T4408'.unpack('m')
=> ["\xE3\x8D+\xE3\x8D\x13\xE3\x8D<"]
```

* アンパックされた文字列のエンコーディングは、ASCII-8BITになる

```ruby
>> a = '440r440T4408'.unpack('m').first
=> "\xE3\x8D+\xE3\x8D\x13\xE3\x8D<"
>> a.force_encoding('UTF-8')
=> "\xE3\x8D+\xE3\x8D\u0013\xE3\x8D<"
```



## 14.文字列内での検索

* `include?`：指定した文字列が含まれている場合に`true`を返す

* `index`：指定された文字や文字コードの数値・正規表現などのパターンを、指定された一から右方向に検索して、最初に見つかった位置を返す

* `rindex`：`index`メソッドとは逆に左方向に検索する。同じ位置を参照する

* `match`：指定された正規表現によるマッチングを行い、マッチした場合には`MatchData`オブジェクトを返す

  * 1個だけマッチしたオブジェクトを返す

* `scan`：指定された正規表現にマッチした部分文字列の配列を返す。ブロックを渡すことも可能

  * 複数個マッチした配列を返す

```ruby
>> "abcdefg".include?("abc")
=> true
>> "abcdefg".index("bc")
=> 1
>> "abcdefg".match(/[c-f]/)
=> #<MatchData "c">
>> "abcdefg".scan(/[c-d]/)
=> ["c", "d"]
```

* `match`が一度しか正規表現によるマッチを行わないのに対し、`scan`は繰り返しマッチを行う

```ruby
>> s = "To be or not to be, that is the question."
=> "To be or not to be, that is the question."
>> hash = Hash.new(0)
=> {}
>> s.scan(/\w+/) {|i| hash[i] += 1}
=> "To be or not to be, that is the question."
>> p hash["be"]
2
=> 2
```



## 15.次の文字列を求める

* `succ`、`next`：次の文字列を求める

> ここでの **次の文字列** とは、、、

> 対象となる文字列の右端がアルファベットならアルファベット順に、数字であれば10進数の数値と見なして計算

> 計算時に桁上がりが発生した場合は、それに応じて１つ左の文字が変化するか、文字列が伸長する

```ruby
>> "abc123".succ
=> "abc124"
>> "123abc".succ
=> "123abd"
>> "123xyz".succ
=> "123xza"
>> "99999".succ
=> "100000"
>> "zzzzz".succ
=> "aaaaaa"
>> "ZZZZZ".succ
=> "AAAAAA"
```



## 16.文字列に対する繰り返し

* `each_line`、`lines`：文字列の各行に対して繰り返す。オプションとして行の区切りを指定可能

* `each_byte`、`bytes`：文字列をバイト単位で繰り返す

* `each_char`、`chars`：文字単位で繰り返す

* `upto`：自分自身から指定された文字列まで、`succ`メソッドで生成される次の文字列を使って繰り返す

```ruby
>> "abc\ndef\nghi".each_line {|c| puts c}
abc
def
ghi
=> "abc\ndef\nghi"
>> "abc".each_byte {|c| puts c
>> }
97
98
99
=> "abc"
>> "ルビー".each_char {|c| puts c}
ル
ビ
ー
=> "ルビー"
>> "a".upto("c") {|c| puts c}
a
b
c
=> "a"
```



## 17.他のクラスへの変換

### `hex`メソッド

* 文字列が16進数であるとして、数値に変換

* 接頭辞`0x`、`0X`とアンダースコア`_`は無視され、16進数以外の文字がある場合にはそれ以降も無視

```ruby
>> "abc".hex
=> 2748
>> "azc".hex
=> 10
>> "0xazc".hex
=> 10
```



### `oct`メソッド

* 文字列が8進数であるとして、数値に変換

* `hex`メソッドとは異なり、接頭辞に応じて8進数以外の変換も行う

```ruby
>> "010".oct
=> 8
>> "0b010".oct   # 0bで2進数指定
=> 2
>> "0x010".oct   # 0xで16進数指定
=> 16
```



### `to_i`メソッド

* 文字列を10進数の整数であるとして、数値に変換する

* 整数として見なせない文字があれば、そこまでを変換して以降を無視する

* 空文字であれば、0を返す

```ruby
>> "123".to_i
=> 123
>> "0123".to_i
=> 123
>> "0x123".to_i
=> 0
>> "".to_i
=> 0
>> "110".to_i(2)   # 110(2進数)を、10進数に変換
=> 6
```



### `to_f`メソッド

* 文字列を10進数の小数として、`Float`オブジェクトに変換する

* 小数と見なせない文字があれば、そこまでを変換して以降を無視する

* 空文字であれば、0.0を返す

```ruby
>> "1.23".to_f
=> 1.23
>> "01.23".to_f
=> 1.23
>> "0x1.23".to_f
=> 0.0
>> "".to_f
=> 0.0
```



### `to_s`、`to_str`メソッド

* 自分自身を返す

```ruby
>> "ルビー".to_s
=> "ルビー"
>> "ルビー".to_str
=> "ルビー"

# IntegerクラスからStringクラスへの変換
>> 7.to_s(3)   # 7を3進数の文字列に変換
=> "21"
```



### `to_sym`、`intern`メソッド

* 文字列に対応するシンボル値を返す

```ruby
=> :abc
>> a.object_id
=> 70263022623280    # オブジェクトid
>> b = :abc
=> :abc
>> b.object_id
=> 70263022623280    # 同じオブジェクトid
```

>以下のクラスは、`String`クラスに用意されていない
>
> * `to_h`
>
> * `to_a`



| 版     | 年/月/日   |
| ------ | ---------- |
| 第二版 | 2019/05/11 |
