gold 模擬試験
============

## 以下のコードを実行するとどうなりますか。該当するものを全て選びなさい

```ruby
>> a = 1.0 + 1
=> 2.0
>> a = a + (1/2r)
=> 2.5
>> a = a + (1 + 2i)
=> (3.5+2i)
```



### 解説

* FloatインスタンスとRationalインスタンスの演算は、Floatインスタンスとなる

* FloatインスタンスとRationalインスタンスの演算は、Complexインスタンスとなる


|         演算         |戻り値クラス|
|---------------------|-----------|
|FixnumとRationalの演算 | Rational |
| FloatとRationalの演算 |   Float  |
| FixnumとComplexの演算 |  Complex |
|  FloatとComplexの演算 |  Complex |
|     Date同士の減算    |  Rational |
|     Time同士の減算    |   Float   |
|   DateTime同士の減算  |  Rational |



## 以下のコードを実行するとどうなりますか

> 正解しているけど、念のため
>
> 2018/11/11

```ruby
>> class A
>>   private
>>   def hoge
>>     puts "A"
>>   end
>> end
=> :hoge
>> class B < A
>>   public :hoge
>> end
=> B
>> B.new.hoge
A
=> nil
```



### 解説

クラスAの`private`メソッド`hoge`は、サブクラスBで`public`に **再定義** されている



## 以下のコードを実行するとどうなりますか

> 正解しているけど、念のため
>
> 2018/11/11

```ruby
>> ary = Array.new(3, "a")
=> ["a", "a", "a"]
>> ary[0].next!
=> "b"
>> p ary
["b", "b", "b"]
=> ["b", "b", "b"]
```



### 解説

`Array.new(3, "a")`は、サイズ3の配列を生成し、文字列オブジェクト"a"を全ての要素に設定する

文字列オブジェクト"a"はコピーされるのではなく、全て同一のオブジェクト

-> 配列の要素は、全て同一の文字列オブジェクト"a"を参照する

従って、１つの要素を変更すると、全ての要素が変更される



## `__X__`に記述する適切なコードを全て選びなさい。(複数選択)

```ruby
Thread.__X__ do
end
```

1. `start`

1. `run`

1. `kick`

1. `new`

1. `fork`



### 解説

Threadクラスで、スレッドを生成し実行するメソッドは、

* `start`

* `new`

* `fork`

です



## 以下のコードで、case文の比較に使用されている演算子はどれですか

> 正解だったが、念のため
>
> 2018/11/11

```ruby
a = [1, "Hello", false]
a.each do |x|
  puts case x
       when String then "string"
       when Numeric then "number"
       when TrueClass, FalseClass then "boolean"
       end
end
```

1. `=`

1. `==`

1. `===`

1. `class`



### 解説

case文の比較には、`===`を使用する



## 以下の実行結果になるように、`__X__`に記述する適切なコードを選びなさい

```ruby
class Hoge
  def fuga(o=nil)
    __X__
  end
  private
  def hoge
    puts "Hoge"
  end
end
Hoge.new.fuga(Hoge.new)

# 実行結果
Hoge
```

1. `hoge`

1. `self.hoge`

1. `o.hoge`

1. `Hoge.hoge`



### 解説

`private`メソッドは、レシーバを指定した呼び出しはできない



## 4行目で生成(raise)される例外オブジェクトのクラスは何ですか

```ruby
>> begin
>>   "cat".narrow        # String#narrowメソッドは存在しない
>> rescue NameError
>>   raise
>> end
NoMethodError: undefined method `narrow` for "cat":String
```



### 解説

rescue節の`raise`は、rescue節が処理中の例外オブジェクトを再生成する

オブジェクトに存在しないメソッドを実行した場合、NoMethodErrorオブジェクトが発生する

2行目でNoMethodErrorオブジェクトが発生するので、3行目以降のrescue節でキャッチしたNoMethodErrorオブジェクトが

そのままraiseされる



## 以下のコードを実行するとどうなりますか

```ruby
>> class A
>>   @@x = 0
>>   class << self
>>     @@x = 1
>>     def x
>>       @@x
>>     end
>>   end
>>   def x
>>     @@x = 2
>>   end
>> end
=> :x

>> class B < A
>>   @@x = 3
>> end
=> 3

>> p A.x
3
=> 3
```



### 解説

クラス変数は、スーパークラス~サブクラス間で共有される

クラスBの`@@x = 3`が変数`x`の最終の定義になるので、`3`が表示される



## 以下のコードを実行するとどうなりますか

```ruby
>> module Mod
>>   def foo
>>     puts "Mod"
>>   end
>> end
=> :foo
>> class Cls1
>>   def foo
>>     puts "Cls1"
>>   end
>> end
=> :foo
>> class Cls2 < Cls1
>>   include Mod
>>   undef foo
>> end
=> nil
>> Cls2.new.foo
NoMethodError: undefined method `foo` for #<Cls2:0x007f91ad083c98>
```



### 解説

クラスCls1で定義された`foo`は、クラスCls2で`undef`により定義が取り消されている



## 以下のコードと同じ実行結果になるコードはどれですか

```ruby
>> a, b = [1, 2]
=> [1, 2]

# 解説
>> a
=> 1
>> b
=> 2
```

1. 選択肢1

  ```ruby
  a=[1,2]
  b=[]
  ```

1. 選択肢2

  ```ruby
  a=1
  b=2
  ```

1. 選択肢3

  ```ruby
  a=[1,2]
  b=nil
  ```

1. 選択肢4

  ```ruby
  a=nil
  b=[1,2]
  ```



### 解説

配列の多重代入の問題

左辺には、右辺の同じ位置にある値が代入される



## 以下の実行結果になるように、`__X__`に記述する適切なコードを全て選びなさい。(複数選択)

```ruby
x = ["abc","defgk","lopq"]
p x.sort{ |a, b| __X__ }

# 実行結果
["abc","lopq","defgk"]
```

1. `a<=>b`

1. `b<=>a`

1. `a.size <=> b.size`

1. `b.size <=> a.size`

1. `a.size - b.size`

1. `b.size - a.size`



### 解説

`sort`は、ブロック引数内に比較のアルゴリズムを記述する

演算子`<=>`は、両オペランドの大小を比較する

* 左オペランドが右オペランドよりも小さい場合は、負の値

* 等しい場合は、`0`

* 左オペランドが右オペランドよりも大きい場合は、正の値

を返す



## 以下のコードを実行すると何が表示されますか。なお、ファイル名はconstant_1.rbです

```ruby
>> class C1
>>   MSG = "msg1"
>>   MSG2 = "msg2"
>>   class C2
>>     MSG = "C2:msg1"
>>     puts MSG
>>     puts MSG2
>>   end
>>   puts MSG
>>   puts MSG2
>> end
C2:msg1
msg2
msg1
msg2
=> nil
```



### 解説

クラスがネストしている場合、定数の探索は

`自分のクラス -> 外側のクラス`

の順で行われる



## 以下のコードを実行すると何が表示されますか

```ruby
begin
  exit
rescue StandardError
  puts "StandardError"
rescue SystemExit
  puts "SystemExit"
end
puts "End"
=> SystemExit
=> End
```



### 解説

組み込み関数`exit`は、例外SystemExitを発生させる

これを`rescue`すれば、実行を継続する

`rescue`しなければ、プログラムを終了する



## stringioライブラリの説明として適切な記述を全て選びなさい

1. 文字列をIOオブジェクトと同じように扱える

1. ファイル入出力専用の文字列である

1. IOクラスのサブクラスである

1. 文字列をファイルに読み書きできる



### 解説

> 解答：1

stringioライブラリは、文字列にIOオブジェクトと同じインターフェースを持たせるStringIOクラスを含んでいる

また、StringIOクラスはIOクラスを継承したクラスではない



## 以下のコードと同じ意味のコードを選びなさい

```ruby
def foo
  puts "Hello"
end
```

1. 選択肢1

  ```ruby
  class Object
    def foo
      puts "Hello"
    end
  end
  ```

1. 選択肢2

  ```ruby
  class Object
    private
    def foo
      puts "Hello"
    end
  end
  ```

1. 選択肢3

  ```ruby
  class Module
    def foo
      puts "Hello"
    end
  end
  ```

1. 選択肢4

  ```ruby
  class Module
    private
    def foo
      puts "Hello"
    end
  end
  ```



### 解説

> 解答：2

トップレベルはObjectクラスのオブジェクト

トップレベルで定義されたメソッドの可視性は`private`



## 以下のコードを実行した結果はどうなりますか。

```ruby
a, b, c = [1, 2]
p a
p b
p c

# 実行結果
1
2
nil
```



### 解説

多重代入において、左辺の要素数が右辺よりも多い場合、余った左辺の要素にはnilが代入される



## "B"と出力するコードを全て選びなさい

1. 選択肢1

  ```ruby
  >> class Object
  >>   def self.const_missing a
  >>     p "#{a}"
  >>   end
  >> end
  => :const_missing
  >> B
  "B"
  => "B"
  ```

1. 選択肢2

  ```ruby
  >> class Module
  >>   def self.const_missing a
  >>     p "#{a}"
  >>   end
  >> end
  => :const_missing
  >> B
  NameError: uninitialized constant B
  ```

1. 選択肢3

  ```ruby
  >> class Hoge
  >>   def self.const_missing a
  >>     p "#{a}"
  >>   end
  >> end
  => :const_missing
  >> Hoge::B
  "B"
  => "B"
  ```

1. 選択肢4

  ```ruby
  >> class Hoge
  >>   def self.const_missing a
  >>     p "#{a}"
  >>   end
  >> end
  => :const_missing
  >> B
  NameError: uninitialized constant B
  ```



### 解説

`const_missing`は、クラスに定数が定義されていない場合に呼び出されるメソッドです

`const_missing`を定義することで、存在しない定数にアクセスした場合の任意の処理を実行できる

選択肢2はModuleクラスに`const_missing`を定義している。選択肢4はHogeクラスに`const_missing`を定義している

呼び出している定数`B`は、トップレベル(Objectクラス)の定数Bを呼び出しているので、

デフォルトの`const_missing`、すなわち例外NameErrorが発生する



## 以下の実行結果になるように、`__X__`に記述する適切なコードを全て選びなさい

```ruby
class Err1 < StandardError; end
class Err2 < Err1; end
begin
  __X__
rescue Err1 => e
  puts "Error"
end

# 実行結果
Error
```

1. `raise StandardError`

1. `raise Err1`

1. `raise Err2`

1. `raise`



### 解説

> 解答：2, 3

出題コードの`rescue Err1`は、

* Err1

* Err1のサブクラス(Err2)

の例外を捕捉する



## 以下のコードを実行するとどうなりますか

```ruby
>> class Root
>>   def m
>>     puts "Root"
>>   end
>> end
=> :m
>> class A < Root
>>   def m
>>     puts "A"
>>   end
>> end
=> :m
>> class B < A
>>   def m
>>     puts "B"
>>   end
>>   undef m
>> end
=> nil
>> B.new.m
NoMethodError: undefined method `m` for #<B:0x007ffb0289dcf8>
```



### 解説

`m`メソッドは、クラスBで`undef`により定義が取り消されている

このため、`B.new.m`を実行すると、例外NoMethodErrorが発生する



## 以下のコードを実行するとどうなりますか

```ruby
>> module M1
>> end
=> nil
>> module M2
>> end
=> nil
>> class Cls1
>>   include M1
>> end
=> Cls1
>> class Cls2 < Cls1
>>   def foo
>>     p self.ancestors
>>   end
>>   include M2
>> end
=> Cls2
>> Cls2.new.foo
NoMethodError: undefined method `ancestors` for #<Cls2:0x007fb3b01e5a58>
```



### 解説

`ancestors`は、Moduleクラスのインスタンスメソッド

レシーバがModuleクラスのインスタンス(つまりモジュール)、またはClassクラスのインスタンス(つまりクラス)の場合に有効

出題コードの`self.ancestors`は、`self`がCls2クラスのインスタンス(Classクラスのインスタンスではない)を示すため、

例外NoMethodErrorが発生する



## 可変長引数について適切な記述を選びなさい

1. 可変長引数にデフォルト値付引数を指定できる

1. 引数はハッシュとして扱える

1. 可変長引数には`*`を付ける

1. １つのメソッドに可変長引数を複数指定することができる



### 解説

> 解答：3

可変長引数にデフォルト値は指定できない

引数は配列として扱われる

可変長引数は1つのメソッドに1つしか指定できない



## DateTimeクラスのオブジェクトに`1`を加算するとどうなりますか

> 一応正解したが、念のため
>
> 2018/11/11

1. 1秒後の時刻を表す

1. 1時間後の時刻を表す

1. 1日後の時刻を表す

1. 1年後の時刻を表す



### 解説

DateTimeクラスのオブジェクトに`1`を加算すると、1日後の時刻を示す。Dateクラスも同様

Timeクラスの場合は、1秒後の時刻を示す



## 以下のコードを実行するとどうなりますか。全て選びなさい

```ruby
>> module Mod
>>   def foo
>>     puts "Mod"
>>     super
>>   end
>> end
=> :foo
>> class Cls1
>>   def foo
>>     puts "Cls1"
>>   end
>> end
=> :foo
>> class Cls2 < Cls1
>>   prepend Mod
>> end
=> Cls2
>> Cls2.new.foo
Mod
Cls1
=> nil
```



### 解説

`prepend`で取り込まれたメソッドは、元から定義されていたメソッドよりも先に呼び出される

また、`prepend`で呼び出されたモジュール内で`super`を呼び出すと、元から定義されていたメソッドが呼び出される



## rdocコメントのマークアップとして適切な記述を全て選びなさい

1. `*word*`で太字

1. `_word_`で等幅フォント

1. `+word+`で斜体

1. `*`で番号なしリスト



### 解説

> 解答：1, 4

`_word_`はイタリック体を表す

`+word+`はタイプライタ(等幅)を示す



##  以下のコードを実行すると何が表示されますか

```ruby
>> d1 = Time.new
=> 2018-11-11 13:54:10 +0900
>> d2 = Time.new
=> 2018-11-11 13:54:10 +0900
>> p (d2 - d1).class
Float
=> Float
```



### 解説

Time同士の減算は、Float型になる



## 以下のコードを実行した時、正しい出力結果を選びなさい

```ruby
f = Fiber.new {
      print "A "
      Fiber.yield "B "
      print "C "
    }
print "D "
print f.resume
print "E "
=> D A B E
```



### 解説

Fiberのブロック内は、`f.resume`が呼び出された時に評価され、`Fiber.yield`まで実行する

`Fiber.yield`が呼ばれると、引数の"B"と共に元のコンテキストに処理を戻す

再び`f.resume`が呼ばれると`Fiber.yield`の次の行から実行しますが、

この問題では`f.resume`を1回しか呼んでいないため、`print "C"`は実行されない

※参考

```ruby
f = Fiber.new {
      print "A "
      Fiber.yield "B "
      print "C "
    }
print "D "
print f.resume
print "E "
print f.resume  # 2回目の`f.resume`(`print "C "`が実行される)
=> D A B E C
```



## Kernelモジュールのfreezeメソッドについて適切な記述を全て選びなさい。(複数選択)

> 正解だったが、念のため
>
> 2018/11/11

1. `clone`は`freeze`、`taint`、特異メソッドなどの情報も含めた完全な複製を作成する

1. `dup`は`freeze`、`taint`、特異メソッドなどの情報も含めた完全な複製を作成する

1. クラスだけでなくモジュールも`freeze`可能である

1. モジュールをインクルードしたクラスを`freeze`することはできない



### 解説

> 解答：1, 3

`clone`と`dup`は、いずれもオブジェクトを複製する

`clone`は、`freeze`、特異メソッドなどの情報も含めた完全な複製を返す

`dup`は、オブジェクトおよび`taint`情報を含めた複製を返す

モジュールは、`freeze`可能

モジュールをインクルードしたクラスも`freeze`可能



## Kernelモジュールのcloneメソッドについて、適切な記述を全て選びなさい。(複数選択)

1. `freeze`、特異メソッドなどの情報も含めてコピーする

1. `freeze`、特異メソッドなどの情報はコピーしない

1. 参照先のオブジェクトもコピーされる

1. 参照先のオブジェクトはコピーされない



### 解説

> 解答：1, 4

Kernelモジュールの`clone`メソッドは、`freeze`、`taint`、特異メソッドなどの情報も含めた完全なコピーを作成する

参照先のオブジェクトはコピーしません(シャローコピー)



## 以下のコードを実行すると何が表示されますか

> 正解したが、念のため
>
> 2018/11/11

```ruby
>> class A
>>   @@a = 0
>>   def initialize
>>     @@a = @@a + 1
>>   end
>>   def A.a
>>     @@a
>>   end
>> end
=> :a
>> class B < A
>> end
=> nil
>> A.new
=> #<A:0x007facc91c8b30>
>> A.new
=> #<A:0x007facca063168>
>> B.new
=> #<B:0x007facca061818>
>> B.new
=> #<B:0x007facc91c3e28>
>> p A.a
4
=> 4
```



### 解説

クラス変数はスーパークラス~サブクラス間で共有される

クラスAの`initialize`は、A、Bのインスタンス化により4回呼び出され、

`@@a`が更新される



## オブジェクトのマーシャリングについて、適切な記述を全て選びなさい

1. IOクラスのオブジェクトは、マーシャリングできない

1. 特異メソッドを持つオブジェクトは、マーシャリングできない

1. ユーザーが作成したオブジェクトは、マーシャリングできない

1. 無名のクラスやモジュールは、マーシャリングできない



### 解説

> 解答：1, 2, 4

システムの状態を保持するオブジェクト(IO、File、Dir、Socket)や特異メソッドを定義したオブジェクトは、マーシャリングできない

また、無名のクラスやモジュールもマーシャリングできない



## 以下の実行結果になるように、`__X__`に記述する適切なコードを全て選びなさい

```ruby
require "json"
h = { "a" => 1, "b" => 2 }
puts __X__

# 実行結果
{ "a" => 1, "b" => 2 }
```

1. `h.to_json`

1. `JSON.dump(h)`

1. `JSON.new(h)`

1. `JSON.to_json(h)`



### 解説

> 解答：1, 2

標準添付ライブラリのjsonは、HashクラスにJSON文字列を生成する`to_json`メソッドを追加する

また、`JSON.dump`メソッドでも同様のJSON文字列を生成する



## 以下のコードを実行するとどうなりますか。全て選びなさい

```ruby
t1 = Thread.start do
       raise ThreadError
     end
```

1. 例外が発生するが、メッセージを出力せずに停止状態になる

1. 例外は発生せず、停止状態になる

1. `-d`オプションを付けた場合には、例外は発生せずに停止状態になる

1. `-d`オプションを付けた場合には、例外発生のメッセージを出力して終了する



### 解答

> 解答：1, 4

マルチスレッドで例外が発生し、その例外が`rescue`で処理されなかった場合、

例外が発生したスレッドだけが警告なしで終了し、他のスレッドは実行を続ける

`-d`オプションを付けるとデバックモードでの実行となり、

いずれかのスレッドで例外が発生した時点dねインタプリタ全体が中断する



| 版 |  年/月/日 |
|----|---------|
|初版|2018/11/11|
