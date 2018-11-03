## 次のプログラムを実行するとどうなりますか

```ruby
>> m = Module.new
=> #<Module:0x007f96b40472f0>

>> CONST = "Constant in Toplevel"
=> "Constant in Toplevel"
>>
>> _proc = Proc.new do
>>   CONST = "Constant in Proc"
>> end
=> #<Proc:0x007f96b38b0ac8@(irb):5>

>> m.instance_eval(<<-EOS)
  CONST = "Constant in Module instance"

  def const
    CONST
  end
EOS
=> :const

>> m.module_eval(&_proc)
(irb):6: warning: already initialized constant CONST
(irb):3: warning: previous definition of CONST was here
=> "Constant in Proc"

>> p m.const
"Constant in Module instance"
=> "Constant in Module instance"
```



### 解説

メソッド`const`は特異クラスで定義されていますので、実行することができます。

その中で参照している定数`CONST`はレキシカルに決定されますので、`"Constant in Module instance"`が表示されます。


`instance_eval`はブロックを渡す場合と、文字列を引数とする場合でネストの状態が異なります。

ブロックを渡した場合はネストは変わりませんが、文字列を引数とした場合は期待するネストの状態になります。

ネストが変わらない状態で定数の代入を行うと、再代入になり警告が表示される場合があります。

例えば、次のプログラムでは`module_eval`に文字列を引数とするとモジュールを再オープン、または定義したネストと同じです。

```ruby
>> module M
>>   p Module.nesting
>> end
[M]
=> [M]

>> M.module_eval(<<-EVAL)
  p Module.nesting
EVAL
[M]
=> [M]

>> M.instance_eval do
>>   p Module.nesting
>> end
[]
=> []

>> module M
>>   p Module.nesting
>> end
[M]
=> [M]
```



## 次のプログラムを実行すると何が表示されますか

```ruby
>> f = Fiber.new do |total|
>>   Fiber.yield total + 10
>> end
=> #<Fiber:0x007fb67d80d280>

>> p f.resume(100) + f.resume(5)
115
=> 115
```



### 解説

Fiberは軽量スレッドを提供します。

Fiber#resumeを実行するとFiber.yieldが最後に実行された行から再開するか、Fiber.newに指定したブロックの最初の評価を行います。

サンプルプログラムを実行して、処理の内容を見てみましょう。

```ruby
>> f = Fiber.new do |name|
>>   Fiber.yield "Hi, #{name}"
>> end
=> #<Fiber:0x007fa095806910>

>> p f.resume('Matz')
"Hi, Matz"
=> "Hi, Matz"
```

1. `f.resume('Matz')`を実行する。

1. `Fiber.new`のブロックを評価し、引数`name`には'Matz'をセットする。

1. 変数を展開して、'Hi, Matz'を`Fiber.yield`の引数にセットする。

1. `Fiber.yield('Hi, Matz')`を実行すると、`f.resume('Matz')`の戻り値が'Hi, Matz'になる。

1. `Fiber.yield('Hi, Matz')`の実行終了を待たずに、プログラムが終了する。

問題では、`Fiber#resume`を２回実行していますが処理の順序は同じです。

```ruby
>> f = Fiber.new do |total|
>>   Fiber.yield total + 10
>> end
=> #<Fiber:0x007f869a80d220>

>> p f.resume(100) + f.resume(5)
115
=> 115
```

1. `f.resume(100)`を実行する。

1. `Fiber.new`のブロックを評価し、引数`total`には`100`をセットする。

1. `100 + 10`を計算して`110`を`Fiber.yield`の引数にセットする。

1. `Fiber.yield(110)`を実行すると、`f.resume(100)`の戻り値が`110`になる。

1. `f.resume(5)`を実行する。

1. `Fiber.yield(110)`から処理を再開し、戻り値が`5`になる。

1. ブロックの最終行になったので、`f.resume(5)`の戻り値が`5`になる。

1. `110 + 5`を評価して、`115`が画面に表示される。


この問題のプログラムを実行すると、`115`が表示されます。

> よくわからない
>
> 2018/11/03



## 次のプログラムと同じ結果になる選択肢を選んでください

この問題ではアクセサを`attr_reader`で作成していますが、`alias`で`original_name`として別名をつけています。

新しく定義した`name`メソッドを実行すると、`Mr. Andrew`と表示されます。

```ruby
>> class Human
>>   attr_reader :name
>>
>>   alias original_name name
>>
>>   def name
>>     "Mr. " + original_name
>>   end
>>
>>   def initialize(name)
>>     @name = name
>>   end
>> end
=> :initialize

>> human = Human.new("Andrew")
=> #<Human:0x007f95588e40c8 @name="Andrew">
>> puts human.name
Mr. Andrew
=> nil
```

#### 選択肢1

`alias`と同じくメソッドの別名をつけます。オーバーライドして元のアクセサを呼び出すことができますので、問題と同じ結果になります

```ruby
>> class Human
>>   attr_reader :name
>>
>>   alias_method :original_name, :name
>>
>>   def name
>>     "Mr. " + original_name
>>   end
>>
>>   def initialize(name)
>>     @name = name
>>   end
>> end
=> :initialize

>> human = Human.new("Andrew")
=> #<Human:0x007fa7228176c8 @name="Andrew">
>> puts human.name
Mr. Andrew
=> nil
```

#### 選択肢2

`name`メソッドの中で`super`で親クラスの同名のメソッドを呼び出そうとしていますが、

親クラスのObjectにはそのようなメソッドはありませんので同じ結果になりません。

```ruby
>> class Human
>>   attr_reader :name
>>
>>   def name
>>     "Mr. " + super
>>   end
>>
>>   def initialize(name)
>>     @name = name
>>   end
>> end
=> :initialize

>> human = Human.new("Andrew")
=> #<Human:0x007fd32a0a7670 @name="Andrew">
>> puts human.name
NoMethodError: super: no superclass method 'name' for #<Human:0x007fd32a0a7670 @name="Andrew">
```

* 選択肢3

イニシャライザで初期化したインスタンス変数を`name`メソッドで参照していますので、問題と同じ結果になります。

```ruby
>> class Human
>>   attr_reader :name
>>
>>   def name
>>     "Mr. " + @name
>>   end
>>
>>   def initialize(name)
>>     @name = name
>>   end
>> end
=> :initialize

>> human = Human.new("Andrew")
=> #<Human:0x007feb948cbbf8 @name="Andrew">
>> puts human.name
Mr. Andrew
=> nil
```

* 選択肢4

`name`メソッドの中で同名のメソッドを呼び出していますので、再帰呼出し

終了せず、例外が発生しますので問題と同じ結果にはなりません。

```ruby
>> class Human
>>   attr_reader :name
>>
>>   def name
>>     "Mr. " + name
>>   end
>>
>>   def initialize(name)
>>     @name = name
>>   end
>> end
=> :initialize

>> human = Human.new("Andrew")
=> #<Human:0x007fd656213250 @name="Andrew">
>> puts human.name
SystemStackError: stack level too deep
```
