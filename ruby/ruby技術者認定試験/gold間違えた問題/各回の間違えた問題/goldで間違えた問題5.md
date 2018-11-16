goldで間違えた問題5
=================

## 次のコードを実行するとどうなりますか

```ruby
>>   def method_missing(id, *args)
>>     puts "M#method_missing"
>>   end
>> end
=> :method_missing
>> class A
>>   include M
>>   def method_missing(id, *args)
>>     puts "A#method_missing"
>>   end
>> end
=> :method_missing
>> class B < A
>>   class << self
>>     def method_missing(id, *args)
>>       puts "B.method_missing"
>>     end
>>   end
>> end
=> :method_missing
>>
?> B.new.dummy_method
A#method_missing
=> nil
```



### 解説

`method_missing`は、継承チェーンを辿った末にメソッドが見つからなかった場合に、呼び出されます。

`method_missing`も継承チェーンを辿ります。

`class << self; end`で定義されたメソッドは、特異クラスのインスタンスメソッドになります。

よって、`B.method_missing`ではなく、`A#method_missing`が出力されます。



## 次のプログラムを実行するとどうなりますか

```ruby
>> class C
>>   CONST = "Hello, world"
>> end
=> "Hello, world"
>>
?> $c = C.new
=> #<C:0x007fa7dd80c5c8>
>>
?> class D
>>   class << $c
>>     def say
>>       CONST
>>     end
>>   end
>> end
=> :say
>>
?> p $c.say
"Hello, world"
=> "Hello, world"
```



### 解説

レキシカルスコープには定数はありません。その場合はスーパークラスを探索します。

特異クラスの継承関係にクラスCがありますので定数を見つけることができます。

参考：特異クラスの継承関係

```ruby
>> $c.class
=> C
>> $c.class.superclass
=> Object

# 継承関係
[#<Class:#<C:0x007fa4741607e0>>, C, Object, Kernel, BasicObject]
```



## 次のプログラムを実行するとどうなりますか

```ruby
>> m = Module.new
=> #<Module:0x007ffed884cfe0>

>> CONST = "Constant in Toplevel"
=> "Constant in Toplevel"

>> _proc = Proc.new do
>>   CONST = "Constant in Proc"
>> end
=> #<Proc:0x007ffed98fd038@(irb):5>

>> m.module_eval(<<-EOS)
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
NoMethodError: undefined method `const` for #<Module:0x007ffed884cfe0>
```



### 解説

メソッド`const`は特異クラスで定義されていないので、例外が発生します。

`const`メソッドを実行したい場合は次のように`module_function`または`instance_eval`を使う必要があります。

```ruby
m.module_eval(<<-EOS) # module_eval のまま
  CONST = "Constant in Module instance"

  def const
    CONST
  end

  module_function :const # module_function にシンボルでメソッドを指定する
EOS
```

```ruby
m.instance_eval(<<-EOS) # instance_eval で特異クラスにメソッドを定義する
  CONST = "Constant in Module instance"

  def const
    CONST
  end
EOS
```



## 次のコードを実行するとどうなりますか

```ruby
>> class S
>>   def initialize
>>     puts "S#initialize"
>>   end
>> end
=> :initialize
>>
?> class C < S
>>   def initialize(*args)
>>     super()
>>     puts "C#initialize"
>>   end
>> end
=> :initialize
>>
?> C.new(1,2,3,4,5)
S#initialize
C#initialize
=> #<C:0x007f815f86b8d8>
```



### 解説

> スーパークラスには、引数が対応していない

`super`はスーパークラスと同名のメソッドが呼ばれます。

引数ありのメソッドで`super`を呼び出すと、引数ありのメソッドが呼ばれますが、そのメソッドが存在しない場合は、ArgumentErrorが発生します。

引数ありのメソッドで引数なしのスーパークラスを呼び出すには、`super()`と明示的に呼び出す必要があります。



## 次のコードを実行するとどうなりますか

```ruby
>> class S
>>   def initialize(*)
>>     puts "S#initialize"
>>   end
>> end
=> :initialize
>>
?> class C < S
>>   def initialize(*args)
>>     super
>>     puts "C#initialize"
>>   end
>> end
=> :initialize
>>
?> C.new(1,2,3,4,5)
S#initialize
C#initialize
=> #<C:0x007ff8b302f590>
```



### 解説

> 正解していたが、念のため
>
> 2018/11/13

`def initialize(*)`は無名の可変長引数を表します。

`super`はスーパークラスにある現在のメソッドと同じメソッドを呼び出します。

`super`は引数指定なしで呼び出した場合は、現在のメソッドと同じ引数が引き渡されます。

スーパークラスで引数を受け取る必要がない場合は、`initialize(*)`とすることで、サブクラスで引数を意識する必要が無くなります。



## 次のプログラムと同じ実行結果が得られる実装を選択肢から選んでください。

```ruby
class Array
  def succ_each(step = 1)
    return enum_for(:succ_each, step) unless block_given?

    each do |int|
      yield int + step
    end
  end
end

p [98, 99, 100].succ_each(2).map {|succ_chr| succ_chr.chr}

[101, 102, 103].succ_each(5) do |succ_chr|
  p succ_chr.chr
end

# 実行結果
["d", "e", "f"]
"j"
"k"
"l"
```

1. 選択肢1

```ruby
>> class Array
>>   def succ_each(step = 1)
>>     return each(:succ_each) unless block_given?
>>
>>     each do |int|
>>       yield int + step
>>     end
>>   end
>> end
=> :succ_each

>> p [98, 99, 100].succ_each(2).map {|succ_chr| succ_chr.chr}
ArgumentError: wrong number of arguments (given 1, expected 0)

>> [101, 102, 103].succ_each(5) do |succ_chr|
>>   p succ_chr.chr
>> end
"j"
"k"
"l"
=> [101, 102, 103]
```

1. 選択肢2

```ruby
>> class Array
>>   def succ_each(step = 1)
>>     return to_enum(:succ_each) unless block_given?
>>
?>     each do |int|
?>       yield int + step
>>     end
>>   end
>> end
=> :succ_each
>>
?> p [98, 99, 100].succ_each(2).map {|succ_chr| succ_chr.chr}
["c", "d", "e"]
=> ["c", "d", "e"]
>>
?> [101, 102, 103].succ_each(5) do |succ_chr|
?>   p succ_chr.chr
>> end
"j"
"k"
"l"
=> [101, 102, 103]
```

1. 選択肢3

```ruby
>> class Array
>>   def succ_each(step = 1)
>>     return to_enum(:succ_each, step) unless block_given?
>>
?>     each do |int|
?>       yield int + step
>>     end
>>   end
>> end
=> :succ_each
>>
?> p [98, 99, 100].succ_each(2).map {|succ_chr| succ_chr.chr}
["d", "e", "f"]
=> ["d", "e", "f"]
>>
?> [101, 102, 103].succ_each(5) do |succ_chr|
?>   p succ_chr.chr
>> end
"j"
"k"
"l"
=> [101, 102, 103]
```

1. 選択肢4

```ruby
>> class Array
>>   def succ_each(step = 1)
>>     unless block_given?
>>       Enumerator.new do |yielder|
?>         each do |int|
?>           yielder << int + step
>>         end
>>       end
>>     else
?>       each do |int|
?>         yield int + step
>>       end
>>     end
>>   end
>> end
=> :succ_each
```



### 解答

ブロックを渡す場合と、チェーンを行う場合の両方を考慮する必要があります。

チェーンを行う場合はEnumeratorオブジェクトを作成する必要があります。作成に必要なメソッドは`enum_for`と`to_enum`です。


問題では、`enum_for`を使っていますので選択肢のうち`to_enum`を使っている選択肢が答えのひとつです。

ただし、`to_enum`は引数にメソッド名とそのメソッドに必要な引数を指定する必要があります。

問題では`succ_each`メソッドに`引数2`を渡していますのでEnumeratorオブジェクトを作成するときに必要になります。


また、Enumeratorオブジェクトは`new`メソッドで作成することが出来ます。

この問題ですと少し冗長ではありますが、全体的には次のとおりです。

```ruby
class Array
  def succ_each(step = 1)
    unless block_given? # ブロックが無い場合は、オブジェクトを作成
      Enumerator.new do |yielder|
        each do |int|
          yielder << int + step
        end
      end
    else # ブロックがある場合の実装
      each do |int|
        yield int + step
      end
    end
  end
end
```

これも答えのひとつで、この問題では`to_enum(:succ_each, step)`とEnumeratorオブジェクトを作成する選択肢が答えになります。

なお、チェーンした先で渡されたブロックを評価するためにはEnumerator::Yielderのオブジェクトを利用します。

オブジェクトに対して、`<<`を実行することでブロック内で評価した結果を受け取ることが出来ます。



## 次のプログラムと同じ結果になる選択肢を選んでください。

```ruby
>> module M
>>   def self.a
>>     100
>>   end
>> end
=> :a
>>
?> p M.a
100
=> 100
```

1. 選択肢1

```ruby
>> module M
>>   include self
>>   def a
>>     100
>>   end
>> end
ArgumentError: cyclic include detected
>>
?> p M.a
NoMethodError: undefined method `a` for M:Module
```

1. 選択肢2

```ruby
>> module M
>>   extend self
>>   def a
>>     100
>>   end
>> end
=> :a
>>
?> p M.a
100
=> 100
```

1. 選択肢3

```ruby
>> module M
>>   def a
>>     100
>>   end
>>
?>   module_function :a
>> end
=> M
>>
?> p M.a
100
=> 100
```

1. 選択肢4

```ruby
>> module M
>>   class << self
>>     def a
>>       100
>>     end
>>   end
>> end
=> :a
>>
?> p M.a
100
=> 100
```



### 解答

モジュールにクラスメソッドを定義するには３つ方法があります。

この問題の答えは次のとおりです。

* 方法1

```ruby
module M
  extend self
  def a
    100
  end
end

p M.a
```

* 方法2

```ruby
module M
  def a
    100
  end

  module_function :a
end

p M.a
```

* 方法3

```ruby
module M
  class << self
    def a
      100
    end
  end
end

p M.a
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

# 期待値
{"price"=>100, "order_code"=>200, "order_date"=>"2018/09/20", "tax"=>0.8}
```

1. `JSON.load json`

1. `JSON.save json`

1. `JSON.parse json`

1. `JSON.read json`



### 解説

> 解答：1

JSON形式の文字列をHashオブジェクトにするメソッドを選ぶ必要があります。

* `JSON.load`

* `JSON.parse`は引数にJSON形式の文字列を指定

するとHashオブジェクトに変換します。



| 版|  年/月/日  |
|---|-----------|
|初版|2018/11/13|
