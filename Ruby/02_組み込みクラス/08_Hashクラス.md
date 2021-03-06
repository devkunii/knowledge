08 Hashクラス
============

## 目次

* [ハッシュとは](#0ハッシュとは)

* [ハッシュの生成](#1ハッシュの生成)

* [ハッシュのキーや値を取得する](#2ハッシュのキーや値を取得する)

* [ハッシュを変更する](#3ハッシュを変更する)

* [ハッシュを調べる](#4ハッシュを調べる)

* [ハッシュを使った繰り返し](#5ハッシュを使った繰り返し)

* [ハッシュをソートする](#6ハッシュをソートする)

* [ハッシュを変換する](#7ハッシュを変換する)

* [ハッシュに変換する](#8ハッシュに変換する)



## 0.ハッシュとは

* ハッシュは連想配列とも呼ばれ、配列でのインデックスにあたるキーとして、数値以外のRubyオブジェクトを利用可能

* Rubyでは、ハッシュはHashクラスのオブジェクトとして生成される



## 1.ハッシュの生成

* ハッシュは、ハッシュ式と呼ばれる記法を使用して生成できる

  * `[]`

  * `Hash.new`

  * `{}`

  * `Hash({})`

* ハッシュ式では、 **キー** と要素である **値** とを、`=>`を使った組み合わせで表現する

```ruby
# ハッシュ式
>> a = {"apple" => "fruit", "coffee" => "drink"}
=> {"apple"=>"fruit", "coffee"=>"drink"}
>> a.class
=> Hash

# []メソッド
>> Hash["apple", "fruit", "coffee", "drink"]   # キーと値を、順番にカンマ`,`で列挙する
=> {"apple"=>"fruit", "coffee"=>"drink"}

# newメソッド
>> a = Hash.new
=> {}
>> a["apple"]
=> nil
>> a = Hash.new("NONE")   # キーが存在しない場合の初期値を設定できる
=> {}
>> a["apple"]
=> "NONE"

# newメソッド(ブロック)
>> a = Hash.new{|hash, key| hash[key] = nil}      # 初期値の設定
=> {}
>> a["apple"]
=> nil
>> a = Hash.new{|hash, key| hash[key] = "NONE"}
=> {}
>> a["apple"]
=> "NONE"

# {}メソッド
>> {}
=> {}

# Hash{()}メソッド
>> Hash({})
=> {}
```

* 初期値とブロックの参照は、以下のもので参照可能

  * 初期値：`default`メソッド

  * ブロック：`default_proc`メソッド

* 初期値は、あとで`default=`メソッドで指定可能

```ruby
>> a = Hash.new("NONE")
=> {}
>> a.default
=> "NONE"
>> a["apple"]
=> "NONE"
>> a.default = "Not exists"
=> "Not exists"
>> a["apple"]
=> "Not exists"
```



## 2.ハッシュのキーや値を取得する

* `[]`メソッド：指定されたキーに対応する値を返す

* `keys`、`values`メソッド：ハッシュの全てのキーと値の配列を生成する

* `values_at`メソッド：指定されたキーに対応する値を、配列で返す

  * 可変長なので、引数を何個も取れる

* `fetch`メソッド：与えられたキーに対する **値** を返す

  * キーが存在しない場合には、2番目の引数が与えられた場合にはその値を、ブロックが与えられていた場合はそのブロックを評価した結果を返す

  * 2番目の引数がない場合は、`IndexError`となる

* `select`メソッド：キーと値の組み合わせについてブロックを評価して、結果が真となる組み合わせのみを含むハッシュを返す

* `find_all`メソッド：キーと値の組み合わせについてブロックを評価するが、返り値はキーと値の配列

  * `select`と機能はほとんど変わらない。

```ruby
# []メソッド
>> a = {"apple" => "fruit", "coffee" => "drink"}
=> {"apple"=>"fruit", "coffee"=>"drink"}
>> a["apple"]
=> "fruit"

# keys、valuesメソッド
>> a = {"apple" => "fruit", "coffee" => "drink"}
=> {"apple"=>"fruit", "coffee"=>"drink"}
>> a.keys
=> ["apple", "coffee"]
>> a.values
=> ["fruit", "drink"]

# values_atメソッド
>> a = {1 => "a", 2 => "b", 3 => "c", 4 => "d"}
=> {1=>"a", 2=>"b", 3=>"c", 4=>"d"}
>> a.values_at(1, 3)
=> ["a", "c"]

# fetchメソッド
>> a = {1 => "a", 2 => "b", 3 => "c", 4 => "d"}
=> {1=>"a", 2=>"b", 3=>"c", 4=>"d"}
>> a.fetch(5, "NONE")
=> "NONE"
>> a.fetch(5){|key| % 2 == 0}
=> false

# selectメソッド
>> a = {1 => "a", 2 => "b", 3 => "c", 4 => "d"}
=> {1=>"a", 2=>"b", 3=>"c", 4=>"d"}
>> a.select{|key, value| key % 2 == 0}
=> {2=>"b", 4=>"d"}
>> a.find_all{|key, value| key % 2 == 0}
=> [[2, "b"], [4, "d"]]
```



## 3.ハッシュを変更する

### `[]=`メソッド

* 配列の場合と同様に、指定されたキーに対応する値を変更する

* キーが存在しない場合には、そのキーと値を登録する。

```ruby
# 破壊的メソッド
>> a = {"apple" => "fruit", "coffee" => "drink"}
=> {"apple"=>"fruit", "coffee"=>"drink"}
>> a["apple"] = "red"
=> "red"
>> a
=> {"apple"=>"red", "coffee"=>"drink"}
>> a["orange"] = "orange"
=> "orange"
>> a
=> {"apple"=>"red", "coffee"=>"drink", "orange"=>"orange"}
```

### `delete`メソッド

* 指定されたキーに対応する値を取り除く

* キーが存在していれば対応する値を、そうでなければ`nil`を返す

* ブロックが与えられた場合には、キーが存在しない場合にブロックの評価結果を返す

```ruby
# 破壊的メソッド
>> a = {"apple" => "fruit", "coffee" => "drink"}
=> {"apple"=>"fruit", "coffee"=>"drink"}
>> a.delete("apple")
=> "fruit"
>> a
=> {"coffee"=>"drink"}
```

### `reject`メソッド

* ブロックを評価した結果が真になる値を取り除いたハッシュを生成して返す

* 元のオブジェクトは変更されない。

> `reject!`とは異なるメソッド
>
> 1要素ずつブロックを要素に渡し、その評価結果が真になった要素を全て取り除いた自分自身を返す

```ruby
>> a = {"apple" => "fruit", "coffee" => "drink"}
=> {"apple"=>"fruit", "coffee"=>"drink"}
>> a.reject{|key, value| value == "drink"}         # 値が"drink"であるものを取り除く
=> {"apple"=>"fruit"}
>> a
=> {"apple"=>"fruit", "coffee"=>"drink"}
```

### `delete_if`、`reject!`メソッド

* ブロックを評価した結果が真になる値を取り除く

* 元のオブジェクトが変更される(破壊的メソッド)

```ruby
>> a = {"apple" => "fruit", "coffee" => "drink"}
=> {"apple"=>"fruit", "coffee"=>"drink"}
>> a.reject!{|key, value| value == "drink"}
=> {"apple"=>"fruit"}
>> a
=> {"apple"=>"fruit"}
```

### `replace`メソッド

* 引数で与えられたハッシュで自分自身を置き換える。

```ruby
>> a = {"apple" => "fruit", "coffee" => "drink"}
=> {"apple"=>"fruit", "coffee"=>"drink"}
>> a.object_id
=> 70364050765140
>> a.replace({"orange" => "fruit", "tea" => "drink"})
=> {"orange"=>"fruit", "tea"=>"drink"}
>> a.object_id
=> 70364050765140       # 同じオブジェクトID・・・自分自身の置き換え
```

### `shift`メソッド

* ハッシュから先頭のキーと値の組み合わせを1つ取り除き、その組み合わせを配列として返す

```ruby
# 破壊的メソッド
>> a = {"apple" => "fruit", "coffee" => "drink"}
=> {"apple"=>"fruit", "coffee"=>"drink"}
>> a.shift
=> ["apple", "fruit"]
>> a
=> {"coffee"=>"drink"}
```

### `merge`メソッド

* 自分自身と引数で指定されたハッシュを統合した、新しいハッシュオブジェクトを返す

* デフォルト値は自分自身の設定が引き継がれる。

* ブロックが与えられた場合は、キーと自分自身の値、指定されたハッシュの値が渡され、ブロックの評価結果が新しいハッシュの値となる

```ruby
>> a = {"apple" => "foods", "coffee" => "drink"}
=> {"apple"=>"foods", "coffee"=>"drink"}
>> a.merge({"orange" => "fruit", "tea" => "drink", "apple" => "fruit"})
=> {"apple"=>"fruit", "coffee"=>"drink", "orange"=>"fruit", "tea"=>"drink"}
>> a
=> {"apple"=>"foods", "coffee"=>"drink"}
>> a.merge({"orange" => "fruit", "tea" => "drink"}){|key, self_val, other_val| self_val}
=> {"apple"=>"foods", "coffee"=>"drink", "orange"=>"fruit", "tea"=>"drink"}
```

### `merge!`、`update`メソッド

* 自分自身と引数で指定されたハッシュを統合する。

* `merge`メソッドとは異なり、元のオブジェクトが変更される。(破壊的メソッド)

```ruby
>> a = {"apple" => "foods", "coffee" => "drink"}
=> {"apple"=>"foods", "coffee"=>"drink"}
>> a.merge!({"orange" => "fruit", "tea" => "drink", "apple" => "fruit"})
=> {"apple"=>"fruit", "coffee"=>"drink", "orange"=>"fruit", "tea"=>"drink"}
>> a
=> {"apple"=>"fruit", "coffee"=>"drink", "orange"=>"fruit", "tea"=>"drink"}
```

### `invert`メソッド

* **キー** と **値** を逆にしたハッシュを返す。

* ただし、値が重複している場合には、結果は不定になる

```ruby
>> a = {"apple" => "foods", "coffee" => "drink"}
=> {"apple"=>"foods", "coffee"=>"drink"}
>> a.invert
=> {"foods"=>"apple", "drink"=>"coffee"}
>> {"orange" => "fruit", "coffee" => "drink", "apple" => "fruit", "tea" => "drink"}.invert
=> {"fruit"=>"apple", "drink"=>"tea"}
```

* 入れ替えの結果キーが重複した場合は、後に定義された方が優先される

* `Hash#revert`メソッドは存在しない。

```ruby
>> h = {a: 100, b: 100}
=> {:a=>100, :b=>100}
>> puts h.invert
{100=>:b}
=> nil
```



### `clear`メソッド

* ハッシュを空にする

```ruby
>> a = {"apple" => "foods", "coffee" => "drink"}
=> {"apple"=>"foods", "coffee"=>"drink"}
>> a.clear
=> {}
```



## 4.ハッシュを調べる

* `length`、`size`メソッド：ハッシュの組み合わせの数を返す

* `empty?`メソッド：ハッシュが空かどうかを調べる

* `has_key?`、`include?`、`key?`、`member?`メソッド：ハッシュに **キー** が存在する場合に真を返す

* `has_value?`、`valid?`メソッド：ハッシュに **値** が存在する場合に真を返す

```ruby
# size、empty?メソッド
>> a = {"apple" => "fruit", "coffee" => "drink"}
=> {"apple"=>"fruit", "coffee"=>"drink"}
>> a.size
=> 2
>> a.empty?
=> false

# key?メソッド
>> a = {"apple" => "fruit", "coffee" => "drink"}
=> {"apple"=>"fruit", "coffee"=>"drink"}
>> a.key?("apple")
=> true
>> a.key?("orange")
=> false

# value?メソッド
>> a = {"apple" => "fruit", "coffee" => "drink"}
=> {"apple"=>"fruit", "coffee"=>"drink"}
>> a.value?("fruit")
=> true
>> a.key?("foods")
=> false
```



## 5.ハッシュを使った繰り返し

* `each`、`each_pair`メソッド：与えられたブロックに **キー** と **値** を渡して評価する

* `each_key`、`each_value`メソッド： **キー** と **値** を与えられたブロックに渡して評価する

```ruby
# eachメソッド
>> a = {"apple" => "fruit", "coffee" => "drink"}
=> {"apple"=>"fruit", "coffee"=>"drink"}
>> a.each{|key, value| puts "#{key} => #{value}\n"}
apple => fruit
coffee => drink
=> {"apple"=>"fruit", "coffee"=>"drink"}

# each_key、each_valueメソッド
>> a = {"apple" => "fruit", "coffee" => "drink"}
=> {"apple"=>"fruit", "coffee"=>"drink"}
>> a.each_key{|key| puts "key: #{key}\n"}
key: apple
key: coffee
=> {"apple"=>"fruit", "coffee"=>"drink"}
```



## 6.ハッシュをソートする

* `sort`メソッド：ハッシュとキーと値の組み合わせの配列に変換し、それをソートした結果を返す。

* ハッシュ自身はソートされない。ブロックが与えられた場合には、キーと値の組み合わせの配列が渡される。

```ruby
>> a = {4 => "a", 3 => "b", 2 => "c", 1 => "d"}
=> {4=>"a", 3=>"b", 2=>"c", 1=>"d"}
>> a.sort
=> [[1, "d"], [2, "c"], [3, "b"], [4, "a"]]
>> a.sort{|a, b| a[1] <=> b[1]}              # 0番目("a")と1番目("b")を比較して、1番目の方が大きいので、評価は-1・・・1番目の方からソート
=> [[4, "a"], [3, "b"], [2, "c"], [1, "d"]]
```



## 7.ハッシュを変換する

* `to_a`：ハッシュを配列に変換する

  * キーと値の組み合わせを配列の配列として生成する

  * 二次元配列の場合に適用可能

```ruby
>> a = {4 => "a", 3 => "b", 2 => "c", 1 => "d"}
=> {4=>"a", 3=>"b", 2=>"c", 1=>"d"}
>> a.to_a
=> [[4, "a"], [3, "b"], [2, "c"], [1, "d"]]
```



## 8.ハッシュに変換する

* `to_h`：2次元配列からハッシュを生成する

  * ただし、このメソッドは`Array`、`Hash`クラスのみ定義されている

```ruby
>> [[1, "data 1"], [2, "data 2"]].to_h
=> {1=>"data 1", 2=>"data 2"}
```



| 版     | 年/月/日   |
| ------ | ---------- |
| 第二版 | 2019/05/11 |
