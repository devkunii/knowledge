goldで間違えた問題3
=================

## 次のコードを実行するとどうなりますか

```ruby
>> module M1
>> end
=> nil

>> module M2
>> end
=> nil

>> class C
>>   include M1
>>   include M2
>> end
=> C

>> p C.ancestors
[C, M2, M1, Object, Kernel, BasicObject]
=> [C, M2, M1, Object, Kernel, BasicObject]
```


### 解説

`include`はモジュールのメソッドをインスタンスメソッドとして追加します。

メソッド探索順は`self`の後に追加されます。

複数回`include`された場合は、後に宣言された方からメソッド探索されます。

> 後に`include`した方が、先に`include`したものを覆う感覚
>
> `include`は、覆う感覚



## 次のコードを実行するとどうなりますか

```ruby
class C
  def self.m1
    200
  end
end

module R
  refine C.singleton_class do
    def m1
      100
    end
  end
end

using R

puts C.m1
=> 100
```



### 解説

`Module#refine`は無名のモジュールを作成します。ブロック内の`self`は無名モジュールになります。

```ruby
>> class C
>> end
=> nil

>> module M
>>   refine C do
>>     self # 無名モジュールを指します
>>   end
>> end
=> #<refinement:C@M>
```

Refinementでクラスメソッドを再定義する場合は次のように`singleton_class`を使います。

ブロックの中で`self.m1`としないのがポイントです。

```ruby
class C
  def self.m1
    'C.m1'
  end
end

module M
  refine C.singleton_class do
    def m1
      'C.m1 in M'
    end
  end
end

using M

puts C.m1
=> C.m1 in M
```



## 次のプログラムを実行するとどうなりますか

```ruby
>> module M
>>   def refer_const
>>     CONST
>>   end
>> end
=> :refer_const

>> module E
>>   CONST = '010'
>> end
=> "010"

>> class D
>>   CONST = "001"
>> end
=> "001"

>> class C < D
>>   include E
>>   include M
>>   CONST = '100'
>> end
=> "100"

>> c = C.new
=> #<C:0x007f9c1704f7f8>
>> p c.refer_const
NameError: uninitialized constant M::CONST
```



### 解説

`refer_const`はモジュールMにありますが、`CONST`はレキシカルに決定されるためモジュールMのスコープを探索します。

この問題では`CONST`が見つからないため例外が発生します。



## 次のプログラムは"Hello, world"と表示します。同じ結果になる選択肢はどれですか（複数選択）

* 問題

```ruby
>> module M
>>   CONST = "Hello, world"
>>   def self.say
>>     CONST
>>   end
>> end
=> :say

>> p M::say
"Hello, world"
=> "Hello, world"
```

1. 選択肢1

```ruby
>> module M
>>   CONST = "Hello, world"
>> end
=> "Hello, world"

>> module M
>>   def self.say
>>     CONST
>>   end
>> end
=> :say

>> p M::say
"Hello, world"
=> "Hello, world"
```

1. 選択肢2

```ruby
>> module M
>>   CONST = "Hello, world"
>> end
=> "Hello, world"

>> M.instance_eval(<<-CODE)
  def say
    CONST
  end
CODE
=> :say

>> p M::say
NameError: uninitialized constant #<Class:M>::CONST
```

1. 選択肢3

```ruby
>> module M
>>   CONST = "Hello, world"
>> end
=> "Hello, world"

>> class << M
>>   def say
>>     CONST
>>   end
>> end
=> :say

>> p M::say
NameError: uninitialized constant #<Class:M>::CONST
```

1. 選択肢4

```ruby
>> module M
>>   CONST = "Hello, world"
>> end
=> "Hello, world"

>> M.module_eval(<<-CODE)
  def self.say
    CONST
  end
CODE
=> :say

>> p M::say
"Hello, world"
=> "Hello, world"
```



### 解説

1. 選択肢1

  定数の定義はメモリ上にあるテーブルに管理されます。

  モジュールMを別々に書いたとしてもテーブルを参照して値を取得できます

  ```ruby
  >> module M
  >>   CONST = "Hello, world"
  >> end
  => "Hello, world"

  >> module M
  >>   def self.say
  >>     CONST
  >>   end
  >> end
  => :say

  >> p M::say
  "Hello, world"
  => "Hello, world"
  ```

1. 選択肢2

  `instance_eval`の引数に文字列を指定するとネストの状態はモジュールMの特異クラスになります。

  `CONST`はモジュールMにのみありますので、例外が発生します。

  ```ruby
  >> module M
  >>   CONST = "Hello, world"
  >> end
  => "Hello, world"

  >> M.instance_eval(<<-CODE)
    def say
      CONST
    end
  CODE
  => :say

  >> p M::say
  NameError: uninitialized constant #<Class:M>::CONST
  ```

1. 選択肢3

  特異クラス定義のコンテキストでは、ネストの状態はモジュールMの特異クラスになります。

  `CONST`はモジュールMにのみありますので、例外が発生します

  ```ruby
  >> module M
  >>   CONST = "Hello, world"
  >> end
  => "Hello, world"

  >> class << M
  >>   def say
  >>     CONST
  >>   end
  >> end
  => :say

  >> p M::say
  NameError: uninitialized constant #<Class:M>::CONST
  ```

1. 選択肢4

  `module_eval`の引数に文字列を指定するとネストの状態はモジュールMになります。

  `CONST`はモジュールMにありますので値を取得できます。

  ```ruby
  >> module M
  >>   CONST = "Hello, world"
  >> end
  => "Hello, world"

  >> M.module_eval(<<-CODE)
    def self.say
      CONST
    end
  CODE
  => :say

  >> p M::say
  "Hello, world"
  => "Hello, world"
  ```



## 以下のコードを実行するとどうなりますか

```ruby
>> def hoge(*args, &block)
>>   block.call(*args)
>> end
=> :hoge

>> hoge(1,2,3,4) do |*args|
>>   p args.length > 0 ? "hello" : args
>> end
"hello"
=> "hello"
```



### 解説

問題のソースコード

```ruby
1: def hoge(*args, &block)
2:   block.call(*args)
3: end
4:
5: hoge(1,2,3,4) do |*args|
6:   p args.length > 0 ? "hello" : args
7: end
```

1行目で引数の値を配列として受け取り、ブロックに配列を渡しています。

2行目で`*`を付けて引数を渡しているので、配列が展開されます`(1, 2, 3, 4)`。

5行目でブロック変数を渡していますが、`*args`と宣言されているので、`[1, 2, 3, 4]`が渡されます。

6行目で`args.length > 0`の結果は真となり、`hello`が出力されます。



### 参考

```ruby
# 配列展開なし
>> hoge(1,2,3,4) do |*args|
?> p args
>> end
[1, 2, 3, 4]
=> [1, 2, 3, 4]

# 配列展開あり
>> hoge(1,2,3,4) do |*args|
?> p *args
>> end
1
2
3
4
=> [1, 2, 3, 4]
```



## 次のコードを実行するとどうなりますか

```ruby
>> p [1,2,3,4].map(&self.method(:*))
NameError: undefined method `*' for class `#<Class:#<Object:0x007fea718ba4f8>>'
```



### 解説

問題の`self`はObjectクラスのインスタンスになります。

Objectクラスには`*`メソッドが定義されていないためエラーになります

```ruby
>> Object.instance_methods.grep(/[*]/)
=> []
>> Object.methods.grep(/[*]/)
=> []
```



## 以下のコードを実行するとどうなりますか

```ruby
>> class S
>>   @@val = 0
>>   def initialize
>>     @@val += 1
>>   end
>> end
=> :initialize

>> class C < S
>>   class << C
>>     @@val += 1
>>   end
>> end
=> 1

>> C.new
=> #<C:0x007ffa358d4050>
>> C.new
=> #<C:0x007ffa358ce628>
>> S.new
=> #<S:0x007ffa358cccb0>
>> S.new
=> #<S:0x007ffa358c7300>

>> p C.class_variable_get(:@@val)
5
=> 5
```



### 解説

`@@val`に`1`加算しているタイミングは以下です。

* Cクラスの特異クラスを定義

* `C.new`の呼び出し

* `S.new`の呼び出し

```ruby
class S
  @@val = 0
  def initialize  # 初期化による加算
    @@val += 1
  end
end

class C < S
  class << C      # クラスCにおける、クラスメソッドの定義(`C.new`もクラスメソッドの呼び出し)
    @@val += 1    # クラスCのクラスメソッドを呼び出すと、加算される
  end
end

C.new
C.new
S.new
S.new

p C.class_variable_get(:@@val)
```



## 次のコードを実行するとどうなりますか

```ruby
require 'date'

d = Date.today - Date.new(2015,10,1)
p d.class
=> Rational
```



### 解説

Dateクラス同士の減算はRationalになります。

その他、似たクラスの演算を以下にまとめます。

|       演算      |戻り値クラス|
|----------------|----------|
|  Date同士の減算  | Rational |
|  Time同士の減算  |   Float  |
|DateTime同士の減算| Rational |



## 次のプログラムを実行するとどうなりますか

> 正解していたが、念のため
>
> 2018/11/10

```ruby
>> class C
>>   @@val = 10
>> end
=> 10

>> module B
>>   @@val = 30
>> end
=> 30

>> module M
>>   include B
>>   @@val = 20
>>
>>   class << C
>>     p @@val
>>   end
>> end
20
=> 20
```



### 解説

クラス変数はクラスに所属するあらゆるもので情報を共有する為にあり、

特異クラス定義の中でクラス変数を定義してもレキシカルに決定されます。

次のプログラムではクラス変数は共有されます。

```ruby
>> class C
>>   class << self
>>     @@val = 10
>>   end
>> end
=> 10

>> p C.class_variable_get(:@@val)
10
=> 10
```

この問題ではクラスCの特異クラス定義をモジュールMで行っています。

**クラス変数はレキシカルに決定** されますので答えは`20`です。



## 次のコードを実行するとどうなりますか

```ruby
>> begin
>>   print "liberty" + :fish.to_s
>> rescue TypeError
>>   print "TypeError."
>> rescue
>>   print "Error."
>> else
>>   print "Else."
>> ensure
>>   print "Ensure."
>> end
libertyfishElse.Ensure.=> nil
```



### 解説

* `:fish`はSymbolクラスのオブジェクトです。

  `Symbol#to_s`でStringオブジェクトが返されます。

* `String#+`の引数はStringクラスを期待します。

  Stringクラス以外のオブジェクトが渡された場合は、`TypeError`を発生させます。

* エラーを受け取るためには`rescue`で、例外を受け取った際の処理を記述します。

* エラーが発生しなかった場合の処理を行うには`else`を用います。

* エラー発生有無に関わらず、必ず実行される、後処理を行うには`ensure`を用います。

> `print`は、改行をせずに出力する！！



## 次のプログラムを実行するとどうなりますか

```ruby
>> mod = Module.new
=> #<Module:0x007fc676047bf0>

>> mod.module_eval do
>>   EVAL_CONST = 100
>> end
=> 100

>> puts "EVAL_CONST is defined? #{mod.const_defined?(:EVAL_CONST)}"
EVAL_CONST is defined? true
=> nil
>> puts "EVAL_CONST is defined? #{Object.const_defined?(:EVAL_CONST)}"
EVAL_CONST is defined? true
=> nil
```



### 解説

定数のスコープはレキシカルに決定されます。

**ブロックはネストの状態を変更しない** ので、`module_eval`のブロックで定義した定数はこの問題ではトップレベルで定義したことになります。

定数`EVAL_CONST`はトップレベルで定義していることになりますので、Objectクラスに定数あることが確認することが出来ます。

また、Moduleクラスのインスタンスには直接、定数は定義されていませんが継承関係を探索して参照することが出来ます。

`const_defined?`メソッドは第2引数に継承関係を探索するか指定出来るため、この問題では探索を行うかによって結果が変わります。

```ruby
>> mod = Module.new
=> #<Module:0x007f7f0504b600>

>> mod.module_eval do
>>   EVAL_CONST = 100
>> end
=> 100

>> puts Object.const_defined? :EVAL_CONST
true
=> nil

>> puts mod.const_defined? :EVAL_CONST
true
=> nil

# 第2引数にfalseを指定すると継承関係まで探索しない
>> puts mod.const_defined? :EVAL_CONST, false
false
=> nil
```

この問題では指定してない（デフォルト値`true`）ため探索を行い、定数をどちらも見つけることが出来ます。



## 期待した出力結果になるように`XXXX`に適切なコードを選べ

```ruby
class String
  XXXX
end

p "12345".hoge

#
# 実行結果
#
54321
```

1. `alias :hoge, :reverse`

1. `alias :reverse, :hoge`

1. `alias hoge reverse`

1. `alias reverse hoge`



### 解説

`alias`式はメソッドやグローバル変数に別名を付けることができます。

定義は以下のようにします。

```ruby
alias new_method old_method
alias :new_method :old_method
alias $new_global_val $old_global_val
```

メソッド内でメソッドに別名をつける必要がある場合は、`Module#alias_method`を使います。

```ruby
alias_method "new_method", "old_method"
alias_method :new_method, :old_method
```

```ruby
>> class String
>>   alias hoge reverse
>> end
=> nil
>> p "12345".hoge
"54321"
=> "54321"
```



## 次のコードを実行するとどうなりますか。

```ruby
>> 10.times{|d| print d < 2...d > 5 ? "O" : "X" }
OOOOOOOXXX=> 10
```



### 解説

`Integer#times`は`0`から`self -1`までの数値を順番にブロックに渡すメソッドです。

`d < 2...d > 5`の部分は条件式に範囲式を記述しています。

この式は、フリップフロップ回路のように一時的に真偽を保持するような挙動をとります。

> わからないので、パス
>
> 2018/11/10



## 次のプログラムの実行結果を得るために`__(1)__`に適切なメソッドをすべて選んでください。

```ruby
module Enumerable
  def with_prefix(prefix)
    return to_enum(__(1)__, prefix) { size } unless block_given?

    each do |char|
      yield "#{prefix} #{char}"
    end
  end
end

[1,2,3,4,5].with_prefix("Awesome").reverse_each {|char|
  puts char
}

#
# 実行結果
#
Awesome 5
Awesome 4
Awesome 3
Awesome 2
Awesome 1
```

1. `:with_prefix`

1. `:reverse_each`

1. `__method__`

1. `:each`



### 解説

ブロックを渡さない場合は、Enumeratorオブジェクトを作成してメソッドをチェーン出来るようにします。

Enumeratorオブジェクトを作成するためには、`to_enum`または、`enum_for`を呼びます。

これらの引数にメソッド名をシンボルで指定することでチェーンした先でブロックを渡されたときにどのメソッドを評価すればよいかが分かります。

この問題では、`with_prefix`を再び評価する必要がありますので、`__method__`または`:with_prefix`を引数に指定します。

`__method__`はメソッドの中で呼び出すと、そのメソッド名になります。

```ruby
>> def awesome_method
>>   __method__
>> end
=> :awesome_method

>> p awesome_method
:awesome_method
=> :awesome_method
```



## 次のコードを実行するとどうなりますか

```ruby
>> module M
>>   def self.class_m
>>     "M.class_m"
>>   end
>> end
=> :class_m

>> class C
>>   include M
>> end
=> C

>> p C.methods.include? :class_m
false
=> false
```



### 解説

問題コードの注意すべき点は以下の通りです。

* `include`はModuleの **インスタンスメソッド** をMix-inするメソッドです。

* `def self.class_m`と宣言すると、特異クラスのメソッドになります。

* `C.methods`はCの特異メソッドを表示します。

よって、Cには`class_m`が追加されません。

```ruby
#
# includeをextendに変更
#
>> module M
>>   def self.class_m
>>     "M.class_m"
>>   end
>> end
=> :class_m

>> class C
>>   extend M
>> end
=> C

>> p C.methods.include? :class_m
false
=> false

#
# `self`を削除
#
>> module M
>>   def class_m
>>     "M.class_m"
>>   end
>> end
=> :class_m

>> class C
>>   extend M
>> end
=> C

>> p C.methods.include? :class_m
true
=> true
```



## 次のコードを実行するとどうなりますか

```ruby
>> begin
>>   raise
>> rescue => e
>>   puts e.class
>> end
RuntimeError
=> nil
```



### 解説

引数なしで`raise`を呼び出すと、`RuntimeError`例外が発生します。



## 次のコードを実行するとどうなりますか

```ruby
>> class S
>>   def initialize
>>     puts "S#initialize"
>>   end
>> end
=> :initialize

>> class C < S
>>   def initialize(*args)
>>     super
>>     puts "C#initialize"
>>   end
>> end
=> :initialize

>> C.new(1,2,3,4,5)
ArgumentError: wrong number of arguments (given 5, expected 0)
```



### 解説

問題のコードは`ArgumentError: wrong number of arguments (5 for 0)`が発生します。

`super`と呼び出した場合は、現在のメソッドと同じ引数が引き継がれます。

引数を渡さずにオーバーライドしたメソッドを呼び出す際は`super()`とします。

問題のコードは次のように修正します。

修正後

```ruby
>> class S
>>   def initialize
>>     puts "S#initialize"
>>   end
>> end
=> :initialize

>> class C < S
>>   def initialize(*args)
>>     super() # 引数なしを明示的に指定する
>>     puts "C#initialize"
>>   end
>> end
=> :initialize

>> C.new(1,2,3,4,5)
S#initialize
C#initialize
=> #<C:0x007fe18a8ead78>
```



## 次の2つのプログラムを実行するとどうなりますか

* `lib.rb`の内容

```ruby
module Lib
  $num += 1
end
```

* `program.rb`の内容

```ruby
$num = 0
1..10.times do |n|
  require './lib.rb'
end
puts $num
```



### 解説

`require`はRubyライブラリをロードします。

`require`と`load`の違い

  * `require`は同じファイルは1度のみロードする、`load`は無条件にロードする。

  * `require`は`.rb`や`.so`を自動補完する、`load`は補完は行わない。

  * r`equire`はライブラリのロード、`load`は設定ファイルの読み込みに用いる。

```ruby
# requireの場合
$num = 0
1..10.times do |n|
  require './lib.rb'
end
puts $num
=> 1

# loadの場合
$num = 0
1..10.times do |n|
  load './lib.rb'
end
puts $num
=> 10
```



## 次のプログラムの期待値を得られるように正しいメソッドを選んでください。

```ruby
require 'json'

json = <<JSON
{
  "price":100,
  "order_code":200,
  "order_date":"2018/09/20",
  "tax":0.8
}
JSON

hash = __(1)__
p hash

#
# 期待値
#
{"price"=>100, "order_code"=>200, "order_date"=>"2018/09/20", "tax"=>0.8}
```

1. `JSON.load json`

1. `JSON.save json`

1. `JSON.parse json`

1. `JSON.read json`



### 解説

JSON形式の文字列をHashオブジェクトにするメソッドを選ぶ必要があります。

`JSON.load`または、`JSON.parse`は引数にJSON形式の文字列を指定するとHashオブジェクトに変換します。

```ruby
require 'json'

json = <<JSON
{
  "price":100,
  "order_code":200,
  "order_date":"2018/09/20",
  "tax":0.8
}
JSON

using_parse = JSON.parse json
p using_parse
=> {"price"=>100, "order_code"=>200, "order_date"=>"2018/09/20", "tax"=>0.8}

using_load = JSON.load json
p using_load
=> {"price"=>100, "order_code"=>200, "order_date"=>"2018/09/20", "tax"=>0.8}
```



## 次のプログラムを実行するとどうなりますか

> 正解していたが、念のため
>
> 2018/11/10

```ruby
>> characters = ["a", "b", "c"]
=> ["a", "b", "c"]

>> characters.each do |chr|
>>   chr.freeze
>> end
=> ["a", "b", "c"]

>> upcased = characters.map do |chr|
>>   chr.upcase
>> end
=> ["A", "B", "C"]

>> p upcased
["A", "B", "C"]
=> ["A", "B", "C"]
```



### 解説

`freeze`はオブジェクトの破壊的な変更を禁止します。

```ruby
>> char = "a"
=> "a"
>> char.freeze
=> "a"
>> p char.upcase!
RuntimeError: cannot modify frozen String
```

問題では配列の各要素を破壊的な変更を禁止しています。

さらにその要素を`upcase`で大文字に変換していますが、破壊的な変更ではないため例外は発生しません。

`["A", "B", "C"]`がこの問題の答えです。



| 版 |  年/月/日 |
|----|----------|
|初版|2018/11/10|
