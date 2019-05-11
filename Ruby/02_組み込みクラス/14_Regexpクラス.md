14 Regexpクラス
==============

## 目次

* [Regexpクラスとは](#0Regexpクラスとは)

* [正規表現オブジェクトを生成](#1正規表現オブジェクトを生成)

* [正規表現オブジェクトでマッチングする](#2正規表現オブジェクトでマッチングする)

* [正規表現の特殊文字をエスケープする](#3正規表現の特殊文字をエスケープする)

* [マッチした結果を取得する](#4マッチした結果を取得する)

* [正規表現の論理和を求める](#5正規表現の論理和を求める)

* [正規表現オブジェクトのオプションや属性を取得する](#6正規表現オブジェクトのオプションや属性を取得する)



## 0.Regexpクラスとは

* 正規表現オブジェクトを扱うクラス

* 正規表現を使って文字列やデータのマッチングを行うときに、使用



## 1.正規表現オブジェクトを生成

* 正規表現は、正規表現リテラルを使って表現する

* リテラルの末尾には、オプションが指定できる。

* オプションは、以下のものなどがある。また、オプションの複数指定もできる

  * `i`：大文字小文字の違いを無視

  * `m`：正規表現の`.`で改行にマッチさせる

  * `x`：空白や`#`から始まるコメントを無視する

```ruby
>> a = /abcdefg/i  # 大文字小文字の違いを無視する
=> /abcdefg/i
>> a.class
=> Regexp
```

* `Regexp.new`、`Regexp.compile`：正規表現オブジェクトを生成する。2つ目の引数に、マッチングのオプションを指定できる。

  * `Regexp::IGNORECASE`：大文字小文字の違いを無視する

  * `Regexp::MULTILINE`：正規表現の`.`が改行にマッチするようになる

  * `Regexp::EXTENDED`：バックスラッシュでエスケープされていない空白と、`#`から改行までを無視する

* 論理和を使って複数指定も可能

* マッチングするときの文字コードを、3番目の引数で指定することも可能

```ruby
>> a = Regexp.new("abcdefg", Regexp::MULTILINE | Regexp::IGNORECASE)
=> /abcdefg/mi # オプションのmとi
```



## 2.正規表現オブジェクトでマッチングする

* `match`：正規表現オブジェクトで文字列とマッチングさせる

  * マッチした場合には`MatchData`オブジェクトを、しなかった場合には`nil`を返す

  * 最初の部分しかマッチしない

* `=~`：正規表現オブジェクトで文字列とマッチングさせる

  * マッチすればマッチした位置のインデックスが、しなかった場合は`nil`を返す

* `===`：正規表現オブジェクトで文字列とマッチングさせる

  * マッチすれば`true`、しなかった場合は`false`が返る

* `~`：特殊変数`$_`とマッチングする

```ruby
# matchメソッド
>> a = Regexp.new("abc")
=> /abc/
>> a.match("abcdef")
=> #<MatchData "abc">

# =~メソッド
>> a = Regexp.new("abc")
=> /abc/
>> a =~ "abcdef"          # 0番目の文字がマッチした
=> 0
>> "abcdefg" =~ a         # 0番目の文字がマッチした
=> 0

# ===メソッド
>> a = Regexp.new("abc")
=> /abc/
>> a === "abcdef"
=> true

# ~メソッド
>> $_ = "abcdefg"
=> "abcdefg"
>> a = Regexp.new("abc")
=> /abc/
>> ~ a
=> 0
```



## 3.正規表現の特殊文字をエスケープする

* `Regexp.escape`、`Regexp.quote`：ピリオド`.`、カッコ`[]`などでマッチングする際に、これらの文字を自動的にエスケープする

```ruby
>> Regexp.escape("array.push(hash[key])")
=> "array\\.push\\(hash\\[key\\]\\)"
```



## 4.マッチした結果を取得する

* `Regexp.last_match`：正規表現でマッチした結果を取得

  * `MatchData`オブジェクト(現在のスコープ(トップレベルやクラス・モジュール・メソッド定義)の中で最後に行った正規表現のマッチ結果)を返す

  * 特殊変数`$_`でも取得できる

```ruby
>> /abcdefg/ =~ "abcdefghijklmnopqrstuvwxyz"
=> 0
>> Regexp.last_match
=> #<MatchData "abcdefg">
>> $~
=> #<MatchData "abcdefg">
```

* `Regexp.last_match`メソッドに整数値を与えると、該当のマッチ文字列が得られる。

  * `0`であれば正規表現にマッチした文字列

  * それ以降の整数では、カッコにマッチした部分文字列が得られる

  * これらの文字列はそれぞれ特殊変数`$&`、`$1`、`$2`などでも取得可能

```ruby
>> /(abc)d(efg)/ =~ "abcdefghijklmnopqrstuvwxyz"
=> 0
>> Regexp.last_match(0)
=> "abcdefg"
>> $&
=> "abcdefg"
>> Regexp.last_match(1)
=> "abc"
>> $1
=> "abc"
>> $2
=> "efg"
```



## 5.正規表現の論理和を求める

* `Regexp.union`：複数の正規表現を結合し、そのどれかにマッチするような新しい正規表現を求める

```ruby
>> a = Regexp.new("abc")
=> /abc/
>> b = Regexp.new("ABC")
=> /ABC/
>> c = Regexp.union(a, b)
=> /(?-mix:abc)|(?-mix:ABC)/
>> c =~ "abc"
=> 0
>> Regexp.last_match
=> #<MatchData "abc">
```



## 6.正規表現オブジェクトのオプションや属性を取得する

* `options`：正規表現オブジェクトを生成する時に設定したオプションの論理和を返す

> ### オプション
>
> * `Regexp::IGNORECASE`
>
> * `Regexp::MULTILINE`
>
> * `Regexp::EXTENDED`

* `casefold?`：オプション`Regexp::IGNORECASE`が設定してあるかどうかを返す

* `encoding`：正規表現オブジェクトがコンパイルされている文字コードを`Encoding`オブジェクトとして返す

* `source`：正規表現の元となった文字列表現を返す

  * `to_s`：他の正規表現に埋め込んでも元の意味が保たれるような形式

  * `inspect`：`to_s`メソッドよりも自然な形式な文字列になるが、元の意味は保たれない

```ruby
# optionsメソッド
>> a = Regexp.new("abcdefg", Regexp::MULTILINE | Regexp::IGNORECASE)
=> /abcdefg/mi
>> a.options
=> 5
>> a.options & Regexp::IGNORECASE
=> 1
>> a.options & Regexp::EXTENDED
=> 0

# casefold?メソッド
>> a = Regexp.new("abcdefg")
=> /abcdefg/
>> a.casefold?
=> false
>> a = Regexp.new("abcdefg", Regexp::MULTILINE | Regexp::IGNORECASE)
=> /abcdefg/mi
>> a.casefold?
=> true

# encodingメソッド
>> a = Regexp.new("ルビー")
=> /ルビー/
>> a.encoding
=> #<Encoding:UTF-8>
>> a = Regexp.new("ルビー".encode("EUC-JP"))
=> /\x{A5EB}\x{A5D3}\x{A1BC}/
>> a.encoding
=> #<Encoding:EUC-JP>

# sourceメソッド
>> a = Regexp.new("abcdefg", Regexp::MULTILINE | Regexp::IGNORECASE)
=> /abcdefg/mi
>> a.source
=> "abcdefg"
>> a.to_s
=> "(?mi-x:abcdefg)"
>> a.inspect
=> "/abcdefg/mi"
```



| 版     | 年/月/日   |
| ------ | ---------- |
| 第二版 | 2019/05/11 | 
