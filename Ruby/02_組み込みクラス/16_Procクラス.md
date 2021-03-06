16　Procクラス
=============

## 目次

* [Procクラスとは](#0Procクラスとは)

* [ブロック付きメソッドの引数として利用する](#1ブロック付きメソッドの引数として利用する)

* [手続きオブジェクトの中での処理の中断](#2手続きオブジェクトの中での処理の中断)

* [Proc.new以外の手続きオブジェクト生成](#3Procnew以外の手続きオブジェクト生成)



## 0.Procクラスとは

* ブロックを実行時のローカル変数のスコープなどのコンテキストと共にオブジェクト化した、手続きオブジェクトを扱うクラス

* この手続きオブジェクトは、名前のない関数(無名関数)のように使うことができる



### メソッド

* `call`：この手続きを実行する

```ruby
>> f = Proc.new { puts 'OK' }
=> #<Proc:0x007ff5ce926b40@(irb):58>
>> f.call
OK
=> nil
```

* `arity`：生成された手続きオブジェクトの引数の数を取得

```ruby
>> f = Proc.new {|str| puts str }
=> #<Proc:0x007ff5ce85e118@(irb):60>
>> f.arity
=> 1
>> f.call('NG')
NG
=> nil
```

* オブジェクト生成時のコンテキストを保持しているため、ローカル変数の値などは実行時の状態に応じて変化する

```ruby
>> i = 30
=> 30
>> j = 40
=> 40
>> f = Proc.new { puts i + j }
=> #<Proc:0x007ff5cf09dc48@(irb):66>
>> f.call
70
=> nil
>> i = 100
=> 100
>> f.call
140
=> nil
```



## 1.ブロック付きメソッドの引数として利用する

* ブロック付きメソッドに手続きオブジェクトを渡すこともできる

* 変数の前に`&`を指定して渡す

```ruby
>> f = Proc.new {|i| puts i}
=> #<Proc:0x007ff5cf07f130@(irb):71>
>> 3.times(&f)
0
1
2
=> 3
```



## 2.手続きオブジェクトの中での処理の中断

* `next`：手続きオブジェクトの中で処理を中断して、呼び出し元へ値を戻す

```ruby
>> f = Proc.new {
?>   next "next"   # 中断
>>   "last"
>> }
=> #<Proc:0x007ff5cf06c670@(irb):73>
>> f.call
=> "next"
```



## 3.Proc.new以外の手続きオブジェクト生成

* `lambda`、`proc`：Kernelモジュールのメソッド。手続きオブジェクトを生成する。



### 手続きオブジェクトにおける引数の数

* `lambda`などでは、`proc`メソッドで生成した手続きオブジェクトでは、引数の数が異なると`ArgumentError`が発生

* `Proc.new`で生成した手続きオブジェクトでは、引数への多重代入のようになるので、エラーが発生しない

```ruby
# Proc.newメソッド
>> f = Proc.new {|a, b, c| p a, b, c}
=> #<Proc:0x007ff5cf0554c0@(irb):78>
>> f.call(1, 9)
1
9
nil
=> [1, 9, nil]

# lambdaメソッド
>> g = lambda {|a, b, c| p a, b, c}
=> #<Proc:0x007ff5cf03e2c0@(irb):80 (lambda)>
>> g.call(1, 9)
ArgumentError: wrong number of arguments (given 2, expected 3)
```



### 手続きオブジェクトの中でのジャンプ構文

* `break`では、

  * `lambda`メソッドで生成した手続きオブジェクトでは、その手続きオブジェクトを抜ける

  * `Proc.new`、`proc`メソッドでは、`LocalJumpError`が発生する

```ruby
# Proc.newメソッド
>> f = Proc.new { break }
=> #<Proc:0x007ff5ce9349c0@(irb):82>
>> f.call
LocalJumpError: break from proc-closure

# lambdaメソッド
>> g = lambda { break }
=> #<Proc:0x007ff5cea78ef8@(irb):90 (lambda)>
>> g.call
=> nil
```

* `return`では、

  * `lambda`メソッドで生成した手続きオブジェクトでは、その手続きオブジェクトを抜ける

  * `Proc.new`、`proc`メソッドでは、その手続きオブジェクトの外側を抜けようとするので、`LocalJumpError`が発生する

```ruby
# Proc.newメソッド
>> f = Proc.new { return }
=> #<Proc:0x007ff5cf02eac8@(irb):92>
>> f.call
LocalJumpError: unexpected return
>> def foo
>>   Proc.new {
?>     return 1
>>   }.call
>>   return 2
>> end
=> :foo
>> foo
=> 1

# lambdaメソッド
>> g = lambda { return }
=> #<Proc:0x007ff5cea24268@(irb):101 (lambda)>
>> g.call
=> nil
```



| 版     | 年/月/日   |
| ------ | ---------- |
| 第二版 | 2019/05/11 | 
