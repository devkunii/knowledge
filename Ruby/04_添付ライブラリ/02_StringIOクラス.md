02 StringIOクラス
================

## 目次

* [テキストライブラリの種類](#0テキストライブラリの種類)

* [StringIO](#1StringIO)

* [StringIOのインスタンスの生成](#2StringIOのインスタンスの生成)

* [バッファへ書き込むメソッド](#3バッファへ書き込むメソッド)

* [バッファから読み込むメソッド](#4バッファから読み込むメソッド)

* [ポインタを移動するメソッド](#5ポインタを移動するメソッド)



## 0.テキストライブラリの種類

テキストライブラリには、以下のものがある

* `stringio`：文字列をIOオブジェクトのように取り扱うことができる

* `digest`：文字列のハッシュ化を行う

* `erb`：テンプレートに基づいて出力する

ここでは、`StringIO`について扱う



## 1.StringIOとは

* 文字列をIOクラスと同じインターフェースで取り扱うためのクラス

* メリットとしては、ログ出力のケースが挙げられている

  * ファイルに書き込まずデータベースやメモリ上など 多様な出力先に文字列を出し入れする際に、

  * DuckTyping(IOオブジェクトのようにファイルのように読み書きできればそれで良いじゃないか)の観点で メリットがある

* IOクラスと同じインターフェースをもつが、直接の継承関係はない

![StringIOの継承ツリー](./images/6-2/StringIOの継承ツリー.png)



## 2.StringIOのインスタンスの生成

* インスタンスを生成するには、ファイルディスクリプタの代わりに文字列を渡す

* 第1引数：stringは、文字列を指定する(省略すると、空文字列を指定したものとされる)

* 第2引数：modeは、読み書きに関するオプション(デフォルトでは`r+`)

  * `r`：読み込みモードでインスタンスを作成。ポインタはバッファの先頭を示す

  * `w`：書き込みモードでインスタンスを作成。バッファは空文字列となる

  * `a`：書き込みモードでインスタンスを作成。ポインタはバッファの末尾を示す

  * `r+`：読み書きモードでインスタンスを作成。ポインタはバッファの先頭を示す

  * `w+`：読み書きモードでインスタンスを作成。バッファは空文字列となる

  * `a+`：読み書きモードでインスタンスを作成。ポインタはバッファの末尾を示す

```ruby
>> StringIO.new(string = '', mode = 'r+')
=> #<StringIO:0x007fc6d7138580>
>> StringIO.open(string = '', mode = 'r+')
=> #<StringIO:0x007fc6d7138120>
>> StringIO.open(string = '', mode = 'r+'){|io| p io}
=> #<StringIO:0x007fc6d713b848>
#<StringIO:0x007fc6d713b848>
>> sio = StringIO.new             # 空文字列によるStringIOインスタンスを生成
=> #<StringIO:0x007fd5060a7b10>
```

* `open`：`new`と同様の処理を行う。ブロック内で直接StringIOインスタンスを操作することができる。戻り値は、

  * ブロックを使用しないと、StringIOクラスのインスタンスが返る

  * ブロックを使用すると、ブロック内の評価結果が返る

```ruby
>> require 'stringio'

>> sio = StringIO.open "Hello, StringIO."
>> p sio
=> #<StringIO:0x007fe87d8da1b8>

>> sio = StringIO.open "Hello, StringIO." do |io|
>>         p io.read #=> "Hello, StringIO."
>>         nil
>>       end

>> p sio
=> nil
```



## 3.バッファへ書き込むメソッド

* バッファに書き込むメソッドの中で代表的なものは、`putc`、`puts`、`print`、`printf`

* いずれも、IOクラスで実装されている同名のメソッドと同じ動きをStringIOインスタンス内のバッファに対して行う

* 書き込み後は、バッファが指し示すポインタが入力文字分、移動する



### putc

```ruby
putc(ch)
```

* 1文字だけバッファに書き込むメソッド

* 引数に2文字以上の文字列を与えた場合は、先頭の1文字のみをバッファに書き込み、残りの文字は無視する

* 入力に成功すると、戻り値には引数に渡した`ch`が返る

```ruby
>> p sio = StringIO.new
=> #<StringIO:0x007fab488f3dd8>

>> p sio.putc "a"
=> "a"
>> p sio.string
=> "a"

>> p sio.string = ""   # バッファを空にして、ポインタを0にする
=> ""

>> p sio.putc "bcd"
=> "bcd"
>> p sio.string        # 2文字以上の文字列を与えた場合は、先頭の1文字以外は無視される
=> "b"

>> p sio.putc          # 引数を省略すると、エラー
=> wrong number of arguments (given 0, expected 1) (ArgumentError)
```



### puts

* 引数に与えた文字列に改行を付加して、バッファに書き込む

* 引数を省略すると、改行のみをバッファに書き込む

* 複数の文字列が与えられた場合は、それぞれの文字列に対して改行を付加してバッファに書き込む(戻り値は`nil`)

* 引数には配列を使用可能で、配列内の値をバッファに書き込む

* 引数に配列以外を指定した場合は、`to_ary`メソッドで配列への変換を試み、次に`to_s`メソッドで文字列への変換を試みる

* `print`メソッドでは、`puts`メソッドとは異なり、末尾に改行を付加しない

```ruby
>> sio = StringIO.new
=> #<StringIO:0x007ff5c90c26e8>
>> sio.puts "abc"
=> nil
>> sio.string
=> "abc\n"

>> sio.string = ""
=> ""
>> sio.puts "abc", "def", "hij"
=> nil
>> sio.string
=> "abc\ndef\nhij\n"

# putsメソッドに配列を渡した場合は、配列内それぞれの値がバッファに書き込まれる
>> sio.string = ""
=> ""
>> sio.puts ["abcd", "efgh", "ijkl"]
=> nil
>> sio.string
=> "abcd\nefgh\nijkl\n"

# putsメソッドにnilを渡した場合は、改行のみがバッファに書き込まれる
>> sio.string = ""
=> ""
>> sio.puts nil
=> nil
>> sio.string
=> "\n"

>> sio.string = ""
=> ""
>> sio.print "abc", "def", "hij"
=> nil
>> sio.string
=> "abcdefhij"
```



### printf

* 組み込み関数の`printf`と同じ記述方法で、フォーマットに従ってバッファに書き込む

```ruby
printf(format_string) [, obj, ...]
```



## 4.バッファから読み込むメソッド

* `read`, `readchar`, `readline`などが用意されている

* 読み込み後は、バッファを指し示すポインタが出力文字数分移動する



### read

* バッファ内のポインタが現在指し示している位置から引数に指定した文字数分、文字を取り出す

* 取り出す前に終端になった場合は、終端までの文字列を取り出す

* すでにポインタが文字列の終端を指している場合は、戻り値に`nil`が返る

* 第2引数に変数が指定されている場合は、その変数に取り出した文字列を格納する

* 引数のintegerを省略すると、ポインタが現在指し示している位置から文字列の終端までの文字を取り出す

* ポインタが文字列の終端を指しているときに引数を省略すると、空文字列`""`が返る

```ruby
>> sio = StringIO.new
=> #<StringIO:0x007faf8203bab8>
>> sio.string = "Hello World."
=> "Hello World."
>> sio.read 5
=> "Hello"

>> sio.read
=> " World."

# ポインタが文字列の終端に達している場合、
# 引数を省略すると空文字列が、文字列を与えるとnilが返る
>> sio.read
=> ""
>> sio.read 1
=> nil

# ポインタを0に戻して、outputに対して文字列を出力する
>> sio.pos = 0
=> 0
>> output = ""
=> ""
>> sio.read 5, output
=> "Hello"
>> output
=> "Hello"
```



### getc/readchar

* `getc`：バッファから1文字読み込み、Stringオブジェクトで返す。ポインタが文字列の終端を指している場合は`nil`を返す

* `readchar`：同様に1文字を読み込んでStringオブジェクトを返すが、ポインタが文字列の終端を指している場合は`EOFError`が発生

```ruby
>> sio = StringIO.new
=> #<StringIO:0x007f9ea911eaa0>
>> sio.string
=> ""
>> sio.getc
=> nil

>> sio.pos = sio.string.length
=> 0
>> sio.getc
=> nil
>> sio.readchar
=> in 'readchar': end of file reached (EOFError)
```



### gets/readline

* `gets`：バッファから1行単位で読み込み、文字列を返す。ポインタが文字列の終端を指している場合は、`nil`が返る

* `readline`：バッファから1行単位で読み込み、文字列を返す。ポインタが文字列の終端を指している場合は、例外`EOFError`が発生

```ruby
>> sio = StringIO.new
=> #<StringIO:0x007fb91309e948>
>> sio.string = "Hello World.\nHello IOString World."
=> "Hello World.\nHello IOString World."
>> sio.gets
=>"Hello World.\n"
>> sio.readline
=> "Hello IOString World."

>> sio.gets
=> nil
>> sire.readline
=> undefined local variable or method 'sire' for main:Object (NameError)
```



## 5.ポインタを移動するメソッド

### pos=/seek

* `pos=`：ポインタの指し示す場所を変更するメソッド。先頭を0とした絶対値を指定する

* `seek`：whenceで指定したオプションを基準とした相対値をoffsetに指定する。whenceを省略した場合は、`IO::SEEK_SET`が使用される

  * `IO::SEEK_SET`：ファイルの先頭から(デフォルト)

  * `IO::SEEK_CUR`：現在のファイルポインタから

  * `IO::SEEK_END`：ファイルの末尾から

```ruby
>> sio = StringIO.new
=> #<StringIO:0x007fb2da103d28>
>> sio.string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
=> "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
>> sio.pos = 10
=> 10
>> sio.readline
=> "KLMNOPQRSTUVWXYZ"

>> sio.seek -10, IO::SEEK_END
=> 0
>> sio.readline
=> "QRSTUVWXYZ"
```



| 版     | 年/月/日   |
| ------ | ---------- |
| 初版   | 2018/10/17 |
| 第二版 | 2019/05/13 |
