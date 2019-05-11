03 `Object`クラス
================

## 目次

* [Objectクラスとは](#0Objectクラスとは)

* [オブジェクトのID](#1オブジェクトのID)

* [オブジェクトのクラス](#2オブジェクトのクラス)

* [オブジェクトの比較](#3オブジェクトの比較)

* [オブジェクトのメソッド一覧](#4オブジェクトのメソッド一覧)

* [オブジェクトの複製](#5オブジェクトの複製)

* [オブジェクトの凍結](#6オブジェクトの凍結)

* [オブジェクトの汚染](#7オブジェクトの汚染)

* [インスタンス変数にアクセスする](#8インスタンス変数にアクセスする)

* [未定義メソッドの呼び出し](#9未定義メソッドの呼び出し)

* [オブジェクトの文字列表現](#10オブジェクトの文字列表現)

* [特別なオブジェクト](#11特別なオブジェクト)



## 0.Objectクラスとは

* 全てのクラスのスーパークラス

* `Kernel`モジュールをインクルードしているため、全てのオブジェクトで`Kernel`モジュールのメソッドを使用可能



## 1.オブジェクトのID

* `object_id`：オブジェクトに割り当てられる、重複しない整数を取得する

```ruby
>> a = "foo"
=> "foo"
>> a.object_id
=> 70175072383100
>> a.__id__
=> 70175072383100
```

* Rubyでは **全てがオブジェクト** なので、同じリテラルでも **基本的にオブジェクトIDは異なる**

* しかし、以下のクラスのインスタンスは同じオブジェクトid

  * `TrueClass`

  * `FalseClass`

  * `NilClass`

  * `Symbol`

  * `Fixnum`

```ruby
>> a.object_id
=> 70175072383100
>> "foo".object_id    # 異なるオブジェクトid
=> 70175072330980
>> :foo.object_id
=> 1159388
>> :foo.object_id     # 同じオブジェクトid(Hash)
=> 1159388
```



## 2.オブジェクトのクラス

* `class`：オブジェクトのクラスを調べるメソッド

```ruby
>> "foo".class
=> String
>> :foo.class
=> Symbol
```



## 3.オブジェクトの比較

* `equal?`：オブジェクトID(object_id)を比較する

  * 同じ文字列だったら、`false`(オブジェクトIDが異なる為)

* `eql?`：オブジェクトのハッシュ値が同じかどうかを比較

  * 同じ文字列だったら、`true`

```ruby
>> a = "foo"
=> "foo"
>> a.hash
=> -3839024608656029219
>> a.object_id
=> 70175068384400
>> b = "foo"
=> "foo"
>> b.hash
=> -3839024608656029219     # hashの値は`a`と同じ
>> b.object_id
=> 70175068354380           # object_idの値は`a`と異なる
>> a.eql?(b)
=> true
>> a.equal?(b)
=> false
```

* `===`：`case`式で利用される

  * オブジェクトを比較する

  * 新しいクラスを作った際に定義すると良い

* `==`：オブジェクトの内容が同じかどうかを比較

* 比較メソッドはクラスによって再定義されるため、結果はクラスに依存することに注意する

  * 例)`String`クラスでは、 **`==`メソッドは`eql?`メソッドと同じ結果を返す**

```ruby
>> a = "foo"
=> "foo"
>> b = "foo"
=> "foo"
>> a.eql?(b)
=> true
>> a == b
=> true
```



## 4.オブジェクトのメソッド一覧

### オブジェクトに定義されているメソッドを取得するためのメソッド

* `methods`：全ての呼び出し可能なメソッド

* `private_methods`：プライベートメソッド

* `protected_methods`：プロテクテッドメソッド

* `public_methods`：パブリックメソッド

* `singleton_methods`：特異メソッド

> これらのメソッドは、一覧を配列で返す。配列の要素は、メソッド名のシンボル

```ruby
>> a = "foo"
=> "foo"
>> a.methods
=> [:include?, :%, ] # 省略
```



## 5.オブジェクトの複製

* `clone`：`dup`に加えて、凍結状態(freeze)、特異メソッドも複製

* `dup`：汚染状態(taint)、インスタンス変数、ファイナライザを複製

* ここでの複製は、シャローコピー(浅いコピー)であり、自分自身の複製しかできない

  > 例)配列の要素の参照先は、複製できない

```ruby
>> a = "hoge"
=> "hoge"
>> a.object_id
=> 70175072632680
>> b = a.dup
=> "hoge"
>> b.object_id
=> 70175072614940
```


* 例)foo,barは同じ配列オブジェクトを参照しているとする

  * 複製したものに追加したら、元々の方も追加される

  > * foo.object_id # barと同じ整数値
  >
  > * bar.object_id # fooと同じ整数値
  >
  > * baz.object_id # 上記2つとは違う整数値

```ruby
foo = [1,2,3]
bar = foo
baz = foo.dup

bar[3] = 4
p foo
p bar
p baz

# 解答
=> [1, 2, 3, 4]
=> [1, 2, 3, 4]
=> [1, 2, 3]
```



## 6.オブジェクトの凍結

* `freeze`：オブジェクトの内容の変更を禁止する。

  * 凍結されたオブジェクトを変更しようとすると、`RuntimeError`が発生

```ruby
>> a = "abc"
=> "abc"
>> a.freeze
=> "abc"
>> a[0] = 'z'
RuntimeError: can｀t modify frozen String
```



## 7.オブジェクトの汚染

* オブジェクトの *汚染マーク* がセットされている場合、`tained?`メソッドが`true`を返す

* *汚染マーク* をつけるには`taint`メソッドを、外すには`untaint`メソッドを使用

* オブジェクトの汚染については[こちらを参照](https://docs.ruby-lang.org/ja/2.1.0/doc/spec=2fsafelevel.html)

```ruby
>> a = "string"
=> "string"
>> a.tainted?
=> false
>> a.taint      # 汚染マークをつける
=> "string"
>> a.tainted?
=> true
>> a.untaint    # 汚染マークを外す
=> "string"
>> a.tainted?
=> false
```



## 8.インスタンス変数にアクセスする

* `instance_variable_get`：インスタンス変数の取得

* `instance_variable_set`：インスタンス変数の設定

* `instance_variables`：インスタンス変数の一覧

* インスタンス変数は、`:@hoge`(シンボル)か、`"@hoge"`(文字列)で指定

```ruby
>> class Foo
>>   def initialize
>>     @hoge = 1
>>   end
>> end
=> :initialize
>> f = Foo.new
=> #<Foo:0x007fa5d186d908 @hoge=1>
>> f.instance_variables                # インスタンス変数の一覧を返す
=> [:@hoge]
>> f.instance_variable_get(:@hoge)     # インスタンス変数の取得
=> 1
>> f.instance_variable_set(:@hoge,2)   # インスタンス変数の設定
=> 2
>> f.instance_variable_get(:@hoge)     # インスタンス変数を設定後の取得
=> 2
```



## 9.未定義メソッドの呼び出し

* オブジェクトに定義されていないメソッドが呼び出されたとき、Rubyは`method_missing`メソッドを呼びだす

* `method_missing`メソッドが定義されていない場合は、`NoMethodError`になる

```ruby
>> class Bar
>>   def method_missing(name, *args)
>>     puts name
>>   end
>> end
=> :method_missing
>> b = Bar.new
=> #<Bar:0x007fe42880e210>
>> b.hoge                  # method_missingが呼び出され、`hoge`が出力
hoge
=> nil
>> "string".hoge           # Stringクラスには、method_missingが未定義のため、例外
NoMethodError: undefined method ｀hoge｀ for "string":String
```



## 10.オブジェクトの文字列表現

* `to_s`メソッド：オブジェクトの内容や値の文字列表現を返す

  * 例)オブジェクトのクラス名を表示

* `inspect`メソッド：オブジェクトを人間が読める形式に変換。主にデバッグ

  * 例)インスタンス変数とその値まで表示。標準出力の`p`で使用

```ruby
>> a = 1.2
=> 1.2
>> a.to_s            # 浮動小数点数1.2を文字列に変換
=> "1.2"
>> class Hoge
>>   def initialize
>>     @foo = "bar"
>>   end
>> end
=> :initialize
>> hoge = Hoge.new
=> #<Hoge:0x007fe428987ba0 @foo="bar">
>> hoge.to_s         # インスタンスhogeを、文字列に変換
=> "#<Hoge:0x007fe428987ba0>"
>> hoge.inspect      # インスタンスhogeを、文字列に変換(inspect)
=> "#<Hoge:0x007fe428987ba0 @foo=\"bar\">"
```



## 11.特別なオブジェクト

* `TrueClass`：`true`

* `FalseClass`：`false`

* `NilClass`：`nil`

* これらのクラスのオブジェクトは、唯一のインスタンス

  > 変更できない



| 版     | 年/月/日   |
| ------ | ---------- |
| 第二版 | 2019/05/11 |
