goldで間違えた問題3
=================



## 次のコードを実行するとどうなりますか

`include`はModuleのインスタンスメソッドをMix-inするメソッドです。

`C.methods`はCの特異メソッドを表示します。

よって、`C#class_m`はインスタンスメソッドです、C.methodsでは表示されません。

```ruby
>> module M
>>   def class_m
>>     "class_m"
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



## 次のコードを実行するとどうなりますか

```ruby
>> module M
>>   def class_m
>>     "class_m"
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

`extend`は引数に指定したモジュールのメソッドを特異メソッドとして追加します。

問題の`C.methods...`は特異メソッドの一覧を取得します。



## 以下のコードを実行するとどうなりますか

> 正解だったが、念のため
>
> 2018/11/03

```ruby
>> class C
>> private
>>   def initialize
>>   end
>> end
=> :initialize

>> p C.new.public_methods.include? :initialize
false
=> false
```

`initialize`の可視性は`private`に設定されています。

`initialize`の可視性を`public`に設定したとしても、必ず`private`になります。

```ruby
>> class C
>> public
>>   def initialize
>>   end
>> end
=> :initialize

>> p C.new.public_methods.include? :initialize
false
=> false
```



## 次のコードを実行するとどうなりますか。

```ruby
>> class Human
>>   NAME = "Unknown"
>>
>>   def self.name
>>     const_get(:NAME)
>>   end
>> end
=> :name

>> class Fukuzawa < Human
>>   NAME = "Yukichi"
>> end
=> "Yukichi"

>> puts Fukuzawa.name
Yukichi
=> nil
```

`Class#name`はクラス名を文字列で返します。

`Human#name`クラスは`Class#name`をオーバーライドしているので、`const_get`が呼ばれます。


`const_get`は、`self`に定義された定数を探索します。自クラスに定義がない場合は、メソッドと同様に探索を行います。

問題コードの5行目時点のインスタンスはFukuzawaクラスです。

よって、`Human#name`はFukuzawaクラスのYukichiを返します。

```ruby
1: class Human
2:   NAME = "Unknown"
3:
4:   def self.name
5:     const_get(:NAME)
6:   end
7: end
8:
9: class Fukuzawa < Human
10:   NAME = "Yukichi"
11: end
12:
13: puts Fukuzawa.name
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
=> #<C:0x007fbba68f29d8>
>> p c.refer_const
NameError: uninitialized constant M::CONST
```

`refer_const`はモジュールMにありますが、`CONST`はレキシカルに決定されるためモジュールMのスコープを探索します。

この問題では`CONST`が見つからないため例外が発生します。



## 期待した出力結果になるようにXXXXに適切なコードを選べ

```ruby
class String
  XXXX
end

p "12345".hoge
```

```ruby
# 実行結果
54321

# 解答
alias_method :hoge, :reverse
```

`alias_method`は既に存在するメソッドの別名を付けます。

宣言は`alias 新メソッド名 旧メソッド名`形式で行います。


よく似たメソッドに`alias`があります。異なる点は下記です。

`alias`のメソッド名は識別子かSymbolを受け取ります。

`alias_method`のメソッド名はStringかSymbolを受け取ります。



## 次のコードを実行するとどうなりますか

```ruby
require 'date'

d = Date.today - Date.new(2015,10,1)
p d.class
=> Rational
```

Dateクラス同士の減算はRationalになります。

その他、似たクラスの演算を以下にまとめます。

|        演算       | 戻り値クラス |
|------------------|------------|
|  `Date`同士の減算  | `Rational` |
|  `Time`同士の減算  |   `Float`  |
|`DateTime`同士の減算| `Rational` |



## 次のコードを実行するとどうなりますか

```ruby
>> p "Matz is my tEacher"[/[a-z][A-Z].*/]
"tEacher"
=> "tEacher"
```

スラッシュ(/)で囲まれた文字列は正規表現と扱われます。

問題では、文字列からString#[]で正規表現を用いて部分文字列を抜き取ります。

問題の正規表現`/[a-z][A-Z].*/`を分解すると以下の意味になります。

  * `[a-z]`：1文字目が小文字英字

  * `[A-Z]`：2文字目が大文字英字

  * `.*`：任意の1文字が0回以上繰り返す

以上に該当する部分文字列が表示されます。



## 次のプログラムを実行するとどうなりますか

```ruby
>> m = Module.new
=> #<Module:0x007fb8b5846f50>

>> CONST = "Constant in Toplevel"
=> "Constant in Toplevel"

>> _proc = Proc.new do
>>   CONST = "Constant in Proc"
>> end
=> #<Proc:0x007fb8b48b0a28@(irb):5>

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
NoMethodError: undefined method 'const' for #<Module:0x007fb8b5846f50>
```

メソッド`const`は特異クラスで定義されていないので、例外が発生します。

`const`メソッドを実行したい場合は次のように`module_function`または`instance_eval`を使う必要があります。

```ruby
>> m.module_eval(<<-EOS)  # module_eval のまま
  CONST = "Constant in Module instance"

  def const
    CONST
  end

  module_function :const  # module_function にシンボルでメソッドを指定する
EOS
(eval):1: warning: already initialized constant #<Module:0x007fb8b5846f50>::CONST
(eval):1: warning: previous definition of CONST was here
=> #<Module:0x007fb8b5846f50>
```

```ruby
>> m.instance_eval(<<-EOS)  # instance_eval で特異クラスにメソッドを定義する
  CONST = "Constant in Module instance"

  def const
    CONST
  end
EOS
=> :const
```



## 次のコードを実行するとどうなりますか

```ruby
>> module M
>>   def foo
>>     super
>>     puts "M#foo"
>>   end
>> end
=> :foo

>> class C2
>>   def foo
>>     puts "C2#foo"
>>   end
>> end
=> :foo

>> class C < C2
>>   def foo
>>     super
>>     puts "C#foo"
>>   end
>>   include M
>> end
=> C

>> C.new.foo
C2#foo
M#foo
C#foo
=> nil
```

`include`はモジュールのメソッドをインスタンスメソッドとして追加します。

メソッド探索順は`self`の後に追加されます。

→`C2` -> `M` -> `C` の順番



## 次のコードを実行するとどうなりますか

> 正解だったが、念のため
>
> 2018/11/03

```ruby
>> def foo(n)
>>   n ** n
>> end
=> :foo

>> foo = Proc.new { |n|
>>   n * 3
>> }
=> #<Proc:0x007fb69283b6e8@(irb):5>

>> puts foo[2] * 2
12
=> nil
```

メソッドと変数の探索順位は変数が先です。

Procは`call`または`[]`で呼び出すことができます。

問題では、`foo[2]`と宣言されているため、探索順の早いProcオブジェクトが呼び出されます。

もし、`foo 2`と宣言された場合は、メソッドが選択されます。



## 実行してもエラーにならないコードを選べ

1. 選択肢1

```ruby
>> def bar(&block)
>>   block.yield
>> end
=> :bar

>> bar do
>>   puts "hello, world"
>> end
hello, world
=> nil
```

1. 選択肢2

```ruby
>> def bar(&block)
>>   block.call
>> end
=> :bar

>> bar do
>>   puts "hello, world"
>> end
hello, world
=> nil
```

1. 選択肢3

```ruby
>> def bar(&block, n)
>>   block.call
>> end
SyntaxError: (irb):1: syntax error, unexpected ',', expecting ')'
>>
>> bar(5) {
>>   puts "hello, world"
>> }
NoMethodError: undefined method 'bar' for main:Object
```

1. 選択肢4

```ruby
>> def bar(n, &block)
>>   block.call
>> end
=> :bar

>> bar(5) {
>>   puts "hello, world"
>> }
hello, world
=> nil
```

引数名に&を付与することでブロック引数になります。

ブロック引数は他の引数より後に記述します。



## 次のプログラムを実行するとどうなりますか

```ruby
>> module SuperMod
>>   module BaseMod
>>     p Module.nesting
>>   end
>> end
[SuperMod::BaseMod, SuperMod]
=> [SuperMod::BaseMod, SuperMod]
```



### 解説

`Module.nesting`はネストの状態を表示します。

次のプログラムを実行すると、`[SuperMod]`と表示されます。

```ruby
>> module SuperMod
>>   p Module.nesting
>> end
[SuperMod]
=> [SuperMod]
```

モジュールがネストされた場合は、ネストの状態をすべて表示します。

ネストされたモジュールはプレフィックスに外側にあるモジュールが付与されます。

また、ネスト状態はすべて表示されますがネストが内側から順に表示されます。

```ruby
module SuperMod
  p Module.nesting #=> [SuperMod]

  module BaseMod
    p Module.nesting #=> [SuperMod::BaseMod, SuperMod]

    module BaseBaseMod
      p Module.nesting #=> [SuperMod::BaseMod::BaseBaseMod, SuperMod::BaseMod, SuperMod]
    end
  end
end
```



## 次のプログラムを実行するとどうなりますか

```ruby
class C
end

module M
  refine C do
    def m1(value)
      super value - 100
    end
  end
end

class C
  def m1(value)
    value - 100
  end
end

using M

class K < C
  def m1(value)
    super value - 100
  end
end

puts K.new.m1 400
=> 100
```



### 解説

`super`を実行した場合にもRefinementが影響します。

```ruby
class C
end

module M
  refine C do
    def m1(value)
      p "define m1 using Refinement"
      super value - 100 # 300 - 100
    end
  end
end

class C
  def m1(value)
    p "define m1 in C"
    value - 100 # 200 - 100
  end
end

using M # ここからRefinementが有効になる

class K < C
  def m1(value)
    p "define m1 in K"
    super value - 100 # 400 - 100
    # Refinementが有効なのでsuperはモジュールMにあるm1を参照する
  end
end

puts K.new.m1 400
```

プログラムを実行するとコメントは次の順に表示されます。

1. "define m1 in K"

1. "define m1 using Refinement"

1. "define m1 in C"

`super`を実行したクラスの親クラスにRefinemnetがあれば同名のメソッドを探索して実行します。

さらに、Refinementのなかで`super`を実行するとRefinementの対象クラスのメソッドを探索します。



## 次のコードの実行結果が`falsetrue`になるように`XXXX`,`YYYY`に適切なコードを選択せよ

```ruby
class Company
  XXXX
  attr_reader :id
  attr_accessor :name
  def initialize id, name
    @id = id
    @name = name
  end
  def to_s
    "#{id}:#{name}"
  end
  YYYY
end

c1 = Company.new(3, 'Liberyfish')
c2 = Company.new(2, 'Freefish')
c3 = Company.new(1, 'Freedomfish')

print c1.between?(c2, c3)
print c2.between?(c3, c1)
```

```ruby
# 解答
class Company
  include Comparable
  attr_reader :id
  attr_accessor :name
  def initialize id, name
    @id = id
    @name = name
  end
  def to_s
    "#{id}:#{name}"
  end
  def <=> other
    self.id <=> other.id
  end
end

c1 = Company.new(3, 'Liberyfish')
c2 = Company.new(2, 'Freefish')
c3 = Company.new(1, 'Freedomfish')

print c1.between?(c2, c3)
print c2.between?(c3, c1)
```



### 解説

`between?`で値を比較するためには、Comparableを`include`する必要があります。

Comparableは比較に`<=>`を使用しています。

自作クラスの場合はオブジェクトIDが比較対象となります。

通常は、`Comparable#<=>`をオーバーライドします。

`Fixnum#<=>(other)`は以下の結果を返します。

* `self`が`other`より大きい場合は、`1`を返します。

* `self`が`other`と等しい場合は、`0`を返します。

* `self`が`other`より小さい場合は、`-1`を返します。

`extend`はモジュールのインスタンスメソッドを特異メソッドとして追加します。

インスタンス変数からメソッドを参照することができなくなるので、エラーになります。

Sortableモジュールは存在しません。



## 次のプログラムを実行するとどうなりますか

```ruby
>> mod = Module.new
=> #<Module:0x007ff489046ca0>

>> mod.module_eval do
>>   EVAL_CONST = 100
>> end
=> 100

>> puts "EVAL_CONST is defined? #{mod.const_defined?(:EVAL_CONST, false)}"
EVAL_CONST is defined? false
=> nil
>> puts "EVAL_CONST is defined? #{Object.const_defined?(:EVAL_CONST, false)}"
EVAL_CONST is defined? true
=> nil
```

定数のスコープはレキシカルに決定されます。

ブロックはネストの状態を変更しないので、`module_eval`のブロックで定義した定数は

この問題ではトップレベルで定義したことになります。

また、文字列を引数とした場合はネストの状態を変更します。ネストの状態が変更されるので、

この問題ではモジュールの中でプログラムを書いたことと同じことになります。



## 次のプログラムを実行するとどうなりますか

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

`freeze`はオブジェクトの破壊的な変更を禁止します

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



## 次のコードを実行するとどうなりますか

```ruby
>> CONST_LIST_A = ['001', '002', '003']
=> ["001", "002", "003"]
>> begin
>>   CONST_LIST_A.map{|id| id << 'hoge'}
>> rescue
>> end
=> ["001hoge", "002hoge", "003hoge"]

>> CONST_LIST_B = ['001', '002', '003'].freeze
=> ["001", "002", "003"]
>> begin
>>   CONST_LIST_B.map{|id| id << 'hoge'}
>> rescue
>> end
=> ["001hoge", "002hoge", "003hoge"]

>> CONST_LIST_C = ['001', '002', '003'].freeze
=> ["001", "002", "003"]
>> begin
>>   CONST_LIST_C.map!{|id| id << 'hoge'}
>> rescue
>> end
=> nil

>> CONST_LIST_D = ['001', '002', '003'].freeze
=> ["001", "002", "003"]
>> begin
>>   CONST_LIST_D.push('add')
>> rescue
>> end
=> nil

>> p CONST_LIST_A
["001hoge", "002hoge", "003hoge"]
=> ["001hoge", "002hoge", "003hoge"]
>> p CONST_LIST_B
["001hoge", "002hoge", "003hoge"]
=> ["001hoge", "002hoge", "003hoge"]
>> p CONST_LIST_C
["001", "002", "003"]
=> ["001", "002", "003"]
>> p CONST_LIST_D
["001", "002", "003"]
=> ["001", "002", "003"]
```



### 解説

変数は1文字目を大文字にすると定数になります。定数には次の特徴があります。

1. 代入を行うと警告が発生しますが、値は変更されます。

1. 中身を直接変更した場合は値が変わります。ただし、警告は発生しません。

特徴1の例

```ruby
>> CONST = ["001", "002", "003"]
=> ["001", "002", "003"]
>> CONST = ["A", "B", "C"]
(irb):2: warning: already initialized constant CONST
(irb):1: warning: previous definition of CONST was here
=> ["A", "B", "C"]
>> p CONST
["A", "B", "C"]
=> ["A", "B", "C"]
```

特徴2の例

```ruby
>> CONST = ["001", "002", "003"]
=> ["001", "002", "003"]
>> CONST[0] = "A"
=> "A"
>> p CONST
["A", "002", "003"]
=> ["A", "002", "003"]
```

freezeはオブジェクトを凍結します。凍結されたオブジェクトは次の特徴があります。

1. 破壊的な操作ができません。

1. オブジェクトの代入ができます。

1. 自作クラスのインスタンス変数をfreezeしない限り、変更できます。

特徴1の実行結果

```ruby
>> hoge = "hoge".freeze
=> "hoge"
>> hoge.upcase!
RuntimeError: cannot modify frozen String
>> p hoge
"hoge"
=> "hoge"
```

特徴2の実行結果

```ruby
>> hoge = "hoge".freeze
=> "hoge"
>> hoge = "foo".freeze
=> "foo"
>> p hoge
"foo"
=> "foo"
```

特徴3の実行結果

```ruby
>> class Fish
>>   attr_accessor :name
>>   def initialize(name)
>>     @name = name
>>   end
>> end
=> :initialize

>> liberty = Fish.new("liberty")
=> #<Fish:0x007faf438d8818 @name="liberty">
>> liberty.name.upcase!
=> "LIBERTY"
>> p liberty
#<Fish:0x007faf438d8818 @name="LIBERTY">
=> #<Fish:0x007faf438d8818 @name="LIBERTY">
```



|  版  | 年/月/日 |
|-----|----------|
| 初版 |2018/11/03|
