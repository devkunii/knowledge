# 3-10.`ブロック`と`Proc`

* 特定のリテラルに依存しない
* クロージャに相当

## 3-10-1.ブロックの基本

* 新たにスコープを作成する(for式、if式、while式はスコープが作成されない)
* メソッドを呼び出すときのみ記述できる
* メソッドの内部では、`yield`という式を使用することで、ブロックの内部で記述した処理を呼び出し可能

```ruby
>> def func x
>>   x + yield    # ブロックの実行結果を取得。(ここでは2)
>> end
=> :func
>> p func(1){ 2 } # ブロック付きメソッドfuncの呼び出し
3
=> 3
```

* `{}`で囲まれている部分が、ブロック
* このブロックは2を返す。メソッド`func`では **ブロックの実行結果** + **引数の合計** より、3が返される

***

### スコープが作成されるブロック

* ブロックの実行中に、ブロック内での変数xに値を代入すると、ブロックの外とは別の場所に確保される
* ブロックで初期化された変数は、ブロックの処理が終了すると消滅
* 波カッコ`{}`の代わりに、`do 〜　end`で記述可能

```ruby
>> def func y
>>   y + yield
>> end
=> :func
>> func(1) do
?>   x = 2
>> end
=> 3
>> p x
NameError: undefined local variable or method ｀x｀ for main:Object
```

***

### クロージャとしてのブロック

* スコープを生成することに加えて、ブロック生成時の変数をブロック内で参照可能
* ブロック生成時の変数を更新すると、結果が外部にも影響

```ruby
>> def func y
>>   y + yield
>> end
=> :func
>> x = 2                # ブロックの外で、変数xに値2を代入
=> 2
>> p func(1) {x += 2}   # 代入された値は、メソッド`func`にブロックを渡す
5                       # xの値を取得、更新する
=> 5
>> p x                  # ブロックの外にあるが、更新される
4
=> 4
```

* **値** ではなく、 **変数** そのものが共有される
* ブロック中でxを更新`(x+=2)`すると、xの値が更新
* このような対応付けは、`束縛`という。このような処理の生成時の環境を束縛するものを`クロージャ`という
  →メソッドの内部から外部の変数を参照できないRubyでは、重要

***

## 3-10-2.ブロックのフォーマットと判定

* ブロックは引数を受けることができ、波カッコ`{}`または`do`のあとで、引数リストを`|`で囲む
* これらの実引数は、`yield`で指定可能

```ruby
>> def func a, b
>>   a + yield(b, 3)
>> end
=> :func
>> p func(1, 2){|x,  y| x + y}
6
=> 6
```

1. `func`に1と2を渡す
2. `func`の内部では、第1引数の値1とブロックの実行結果を合計
3. `yield`はブロック引数(2と3)の値を合計して、5を返す
4. **第一引数** + **ブロックの実行結果** より、実行結果は6となる

***

### ブロックの判定

* `block_given?`メソッド：メソッド内部でブロックが指定されたかどうか判定する
* ブロックが指定された時は、それを活用する処理が記述可能(応用例)

```ruby
>> def func
>>   return 1 if block_given?
>>   2
>> end
=> :func
>> p func(){}      # ブロックが指定された場合は1を返す
1
=> 1
>> p func          # ブロック指定されていない場合は2を返す
2
=> 2
```

***

## 3-10-3.`Proc`

* ブロックをオブジェクトとして扱う際に使用
* `Proc`クラスのコンストラクタに、ブロックを指定することで生成
* 実行するには、`Proc`のインスタンスに対して`call`メソッドを使用

```ruby
>> proc = Proc.new{|x| p x}
=> #<Proc:0x007f9edf0534a0@(irb):1>
>> proc.call(1)                     # ブロックでは1を出力するので、そのブロックに1を渡している
1
=> 1
```

***

### Procオブジェクトの生成

* 何らかの初期値がプログラムの冒頭で決定、後で操作する
* Procでは、処理自体を生成して遅延評価することができる
  →初期値や現在の値の管理から解放されている

```ruby
>> def get_counter start
>>   Proc.new{|up| start += up}      # Procオブジェクト生成。startには現在の値を管理
>> end
=> :get_counter
>> count_up = get_counter(1)         # 初期値として1を設定。count_upはProcオブジェクトを参照
=> #<Proc:0x007fad7b00ea08@(irb):2>

# ...たくさんの長い処理...

>> count_up.call(1)                  # count_upの参照するブロックを実行
=> 2

# ...たくさんの長い処理...

>> count_up.call(3)                  # count_upの参照するブロックを実行
=> 5
```

***

### Proc⇄ブロック

* Procオブジェクトに、`&`を付けて最後の引数に指定することで、ブロックへ変換

```ruby
>> def func x
>>   x + yield
>> end
=> :func
>> proc = Proc.new {2}
=> #<Proc:0x007f924303fc78@(irb):4>
>> func(1, &proc)          # procを、&procとして最後の引数に指定
=> 3
```

* ブロックに、最後の引数を`&`を付けた名前を指定することで、引数としてProcオブジェクトを参照

```ruby
>> def func x, &proc
>>   x + proc.call
>> end
=> :func
>> func(1) do
?> 2
>> end
=> 3
```

***

## 3-10-4.`lambda`

`Proc.new`、`proc`とは別の書き方として`lambda`がある

`Proc`オブジェクトを生成するが、`lambda`で作成した`Proc`の方がメソッドの動きに近い

### 特徴

1. 引数のチェックが厳密になる

1. ブロックから値を返す時に`return`を使える

### 基本

```ruby
>> lmd = lambda{|x| p x}
=> #<Proc:0x007fd589850a78@(irb):1 (lambda)>
>> lmd = -> (x){ p x }                         # Ruby1.9からできた新しい書き方
=> #<Proc:0x007fd58a90a8f0@(irb):3 (lambda)>
>> lmd.call(1)
1
=> 1
```

### リターン

* `proc`：`return`を指定すると、生成元のスコープを脱出する

* `lambda`：そのブロック内で`return`すると、呼び出し元に復帰する

```ruby
# proc
>> def func
>>   proc = Proc.new{return 1}
>>   proc.call
>>   2
>> end
=> :func
>> p func
1
=> 1

# lambda
>> def func
>>   proc = lambda{return 1}
>>   proc.call
>>   2   # 実行される
>> end
=> :func
>> p func
2
=> 2
```

### 引数

* `proc`：余分な実引数を無視するか、実引数が足りない場合は`nil`を返す

* `lambda`：引数の数が一致しない場合、`ArgumentError`を発生する

```ruby
# proc
>> p1 = Proc.new{|x,y| y}
=> #<Proc:0x007feaf00d3170@(irb):13>
>> p p1.call(1)
nil
=> nil

# lambda
>> p1 = lambda{|x,y| y}
=> #<Proc:0x007feaf00c2e88@(irb):15 (lambda)>
>> p p1.call(1)
ArgumentError: wrong number of arguments (given 1, expected 2)
```

***

## 3-10-5.ブロックを受けるメソッド

* for式、while式はスコープが作成されない為、配列やハッシュの走査にはあまり使われない
* 代わりに、ブロックを受けるメソッドが使われる

***

### 配列の`each`メソッド

* ブロックの引数には、要素の値が格納される

```ruby
>> [1,2,3].each do |value|
?>   p value
>> end
1
2
3
=> [1, 2, 3]
```

***

### 配列のインデックスを指定する`each_with_index`メソッド

* ブロックで引数を2つ取ることで、第2引数にインデックスが指定される

```ruby
>> [3,4,5].each_with_index do |value, index|
?>   p value + index   # indexは0、1、2
>> end
3
5
7
=> [3, 4, 5]
```

***

### ハッシュの`each`メソッド

* キーと値の２つの引数を受ける

```ruby
>> {a:1, b:2}.each do |key, value|
?>   p "#{key}:#{value}"
>> end
"a:1"
"b:2"
=> {:a=>1, :b=>2}
```

***

### キーのみ、値のみを出力するeachメソッド

* `each_key`メソッド：キーのみ出力
* `each_value`メソッド：値のみ出力

```ruby
# each_keyメソッドの例
>> {a:1, b:2}.each_key do |key|
?> p "#{key}"
>> end
"a"
"b"
=> {:a=>1, :b=>2}

# each_valueメソッドの例
>> {a:1, b:2}.each_value do |value|
?>   p "#{value}"
>> end
"1"
"2"
=> {:a=>1, :b=>2}
```

***

### 範囲オブジェクトの`each`メソッド

* `Range`クラスでも、`each`メソッドは使用可能

```ruby
>> ("a".."e").each do |value|
?>   p value
>> end
"a"
"b"
"c"
"d"
"e"
=> "a".."e"
```

***

### 範囲を指定したループ・回数を指定したループ

* `upto`メソッド：範囲を指定した中で、値を増やしていく
* `downto`メソッド：範囲を指定した中で、値を減らしていく
* `times`メソッド：回数を指定してループを実行する

```ruby
# uptoメソッド
>> 2.upto(4) do |i|
?>   p i
>> end
2
3
4
=> 2

# downtoメソッド
>> 5.downto(1) do |i|
?>   p i
>> end
5
4
3
2
1
=> 5

# timesメソッド
>> 4.times do |i|
?>   p i
>> end
0
1
2
3
=> 4
```

***

## 3-10-6.スレッド

`Thread`クラスを用いて、マルチスレッドのプログラムを書くことができる

※マルチスレッド・・・一つのコンピュータプログラムを実行する際に、複数の処理の流れを並行して進めること。 また、そのような複数の処理の流れ。

### スレッドのサンプルコード

* `Thread`クラスをインスタンス化すると、新しいスレッドを生成することができる

* `join`メソッドによって、スレッドの終了を待つことができる

```ruby
>> t = Thread.new do
?>   p "start"
>>   sleep 5
>>   p "end"
>> end
"start"
=> #<Thread:0x007feaf08f0560@(irb):17 sleep>
>> p "wait""end"

"wait"
=> "wait"
>> t.join
=> #<Thread:0x007feaf08f0560@(irb):17 dead>
```

### スレッドの生成

* `new`、`start`、`fork`：スレッドの生成を行うことができる。これらのメソッドで指定した引数は、ブロックの引数で受けることができる

```ruby
>> 3.times do |i|
?>   Thread.start(i) do |index|
?>     p "thread-#{index} start"
>>   end
>> end
"thread-0 start"
thread-1 start""thread-2 start" "
=> 3
>>

sleep 1
=> 1
```

***

## 3-10-7.ファイバ

スレッドと同様に複数のタスクを切り替え、並行処理をする

### 特徴

* ある処理を途中まで実行して、その後任意のタイミングで前回の続きから処理を行う」ことが可能になる

* スレッドは処理しているタスクの切り替えをOSや仮想マシンが行うのに対して、ファイバは切り替えのタイミングを開発者がプログラム内で明示的に記述する

### ファイバの生成

* `Fiber`クラスをインスタンス化することで、新しいファイバを生成することができる

* ファイバへコンテキストを切り替えるには、`resume`メソッドを衣装する

* `resume`を呼び出すと、対象のファイバ内の処理が終了するか、`Fiber.yield`が呼び出されるまで、ファイバ内の処理を実行する

```ruby
>> f = Fiber.new do
?>     (1..3).each do |i|
?>       Fiber.yield i
>>     end
>>     nil
>> end
=> #<Fiber:0x007feaf00571b0>
>>
?> p f.resume
1
=> 1
>> p f.resume
2
=> 2
>> p f.resume
3
=> 3
>> p f.resume
nil
=> nil
>> p f.resume
FiberError: dead fiber called
```

* `f.resume`が呼び出されたタイミングで、ファイバ内に処理が移る。

* その後、ファイバ内で`Fiber.yield`が呼ばれると、呼び出し元の`resume`に戻る

* `resume`の戻り値は`Fiber.yield`に与えられた引数を返すため、コンソールの最初の行には`1`が出力される

* 次に`resume`が呼び出されると、ファイバ内の続きから実行されるため、コンソールには`2`、`3`と順番に出力される

* 最後には実行するものがないため、`nil`が出力され、それ以上`resume`を呼び出すと例外発生

***


|  回数   |   日付   |
|--------|----------|
|  初版   |2018/08/19|
| 第二版  |2018/10/04|

***
