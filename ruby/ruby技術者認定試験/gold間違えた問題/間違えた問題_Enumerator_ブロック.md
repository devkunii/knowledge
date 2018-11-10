間違えた問題 Enumerator ブロック
============================


# Enumerator


## 次のプログラムは`Enumerator::Lazy`を使っています。先頭から5つの値を取り出すにはどのメソッドが必要ですか

値を取り出すには、

* `Enumerator::Lazy#force`

* `Enumerator::Lazy#first`

を呼び出す必要があります。

問題文には「先頭から5つ」とあるので、`first(5)`として取り出します。

また、`Enumerator::Lazy#force`で問題文の通りにするには`Enumerator::Lazy#take`も利用します。

`Enumerator::Lazy#take`は`Enumerable#take`と違い`Enumerator::Lazy`のインスタンスを戻り値にします。

そのインスタンスから`Enumerator::Lazy#force`で実際の値を取り出します。

```ruby
(1..100).each.lazy.chunk(&:even?)
```



### リファレンスより

* `Enumerable#chunk`：要素を前から順にブロックで評価し、その結果によって 要素をチャンクに分けた(グループ化した)要素を持つ Enumerator を返します。

* `Enumerator::Lazy#chunk`：Enumerable#chunk_while と同じですが、Enumerator ではなく Enumerator::Lazy を返します。

```ruby
# Enumerable#chunk
>> [3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5].chunk {|n|
>>   n.even?
>> }.each {|even, ary|
>>   p [even, ary]
>> }
[false, [3, 1]]
[true, [4]]
[false, [1, 5, 9]]
[true, [2, 6]]
[false, [5, 3, 5]]
=> nil

# Enumerator::Lazy#chunk
>> 1.step.lazy.chunk{ |n| n % 3 == 0 }
=> #<Enumerator::Lazy: #<Enumerator: #<Enumerator::Generator:0x007fc2f1071348>:each>>
```

* `Enumerator::Lazy#force`：全ての要素を含む配列を返します。Lazy から実際に値を取り出すのに使います。

* `Enumerator::Lazy#first`：Enumerable オブジェクトの最初の要素、もしくは最初の n 要素を返します。

```ruby
# Enumerator::Lazy#first
>> 1.step.lazy.first
=> 1

# Enumerator::Lazy#force
>> 1.step.lazy.take(5).force
=> [1, 2, 3, 4, 5]
```

* `Enumerable#take`：Enumerable オブジェクトの先頭から n 要素を配列として返します。

* `Enumerator::Lazy#take`：Enumerable#take と同じですが、配列ではなくEnumerator::Lazy を返します。

```ruby
# Enumerable#take
>> a = [1, 2, 3, 4, 5, 0]
=> [1, 2, 3, 4, 5, 0]
>> a.take(3)
=> [1, 2, 3]

# Enumerator::Lazy#take
>> 1.step.lazy.take(5)
=> #<Enumerator::Lazy: #<Enumerator::Lazy: #<Enumerator: 1:step>>:take(5)>
>> 1.step.lazy.take(5).force
=> [1, 2, 3, 4, 5]
```



### 解答

```ruby
# 解答1
>> (1..100).each.lazy.chunk(&:even?).take(5)
=> #<Enumerator::Lazy: #<Enumerator::Lazy: #<Enumerator: #<Enumerator::Generator:0x007fe94b841a20>:each>>:take(5)>

# 解答2
>> (1..100).each.lazy.chunk(&:even?).take(5).force
=> [[false, [1]], [true, [2]], [false, [3]], [true, [4]], [false, [5]]]

# 解答3
>> (1..100).each.lazy.chunk(&:even?).first(5)
=> [[false, [1]], [true, [2]], [false, [3]], [true, [4]], [false, [5]]]

# 解答4
>> (1..100).each.lazy.chunk(&:even?).first(5).force
NoMethodError: undefined method 'force' for #<Array:0x007fe94d058a50>
```



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

# 実行結果
Awesome 5
Awesome 4
Awesome 3
Awesome 2
Awesome 1
```

```ruby
# 選択肢1
:with_prefix

# 選択肢2
:reverse_each

# 選択肢3
__method__

# 選択肢4
:each
```


### 解説

ブロックを渡さない場合は、Enumeratorオブジェクトを作成してメソッドをチェーン出来るようにします。

Enumeratorオブジェクトを作成するためには、`to_enum`または、`enum_for`を呼びます。

これらの引数にメソッド名をシンボルで指定することでチェーンした先でブロックを渡されたときにどのメソッドを評価すればよいかが分かります。

この問題では、`with_prefix`を再び評価する必要がありますので、`__method__`または:with_prefixを引数に指定します。`__method__`はメソッドの中で呼び出すと、そのメソッド名になります。

```ruby
>> def awesome_method
>>   __method__
>> end
=> :awesome_method

>> p awesome_method
:awesome_method
=> :awesome_method # :awesome_methodとシンボルでメソッド名が分かります
```



### 文法

* `Enumerator#each`：生成時のパラメータに従ってブロックを繰り返します。

  →`each do |f| ~ end`のインスタンスメソッド

* `to_enum`・`enum_for`：レシーバのオブジェクトと列挙用のメソッドを元にしてEnumeratorオブジェクトを作成します。

  →Enumeratorクラスの`new`メソッドを使って、`Enumerator.new(obj, method, arg)`とするのと同じです。

  →引数`method`には、メソッドの名前をシンボルか文字列で渡します。`method`を指定しないと`each`メソッドが使われます。

  ```ruby
  >> enum = "hello".enum_for(:each_byte)
  => #<Enumerator: "hello":each_byte>
  >> p enum.collect {|byte| byte.to_s(16) }
  ["68", "65", "6c", "6c", "6f"]
  => ["68", "65", "6c", "6c", "6f"]
  ```



  ## 次のプログラムと同じ実行結果が得られる実装を選択肢から選んでください。

  ```ruby
  >> class Array
  >>   def succ_each(step = 1)
  >>     return enum_for(:succ_each, step) unless block_given?
  >>
  >>     each do |int|
  >>       yield int + step
  >>     end
  >>   end
  >> end
  => :succ_each
  >>
  >> p [98, 99, 100].succ_each(2).map {|succ_chr| succ_chr.chr}
  ["d", "e", "f"]
  => ["d", "e", "f"]
  >>
  >> [101, 102, 103].succ_each(5) do |succ_chr|
  >>   p succ_chr.chr
  >> end
  "j"
  "k"
  "l"
  => [101, 102, 103]
  ```


  ```ruby
  # 選択肢1
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
  >>
  >> p [98, 99, 100].succ_each(2).map {|succ_chr| succ_chr.chr}     # to_enumを使用していないので、エラー
  ArgumentError: wrong number of arguments (given 1, expected 0)
  >>
  >> [101, 102, 103].succ_each(5) do |succ_chr|
  >>   p succ_chr.chr
  >> end
  "j"
  "k"
  "l"
  => [101, 102, 103]
  ```

  ```ruby
  # 選択肢2
  >> class Array
  >>   def succ_each(step = 1)
  >>     return to_enum(:succ_each) unless block_given?
  >>
  >>     each do |int|
  >>       yield int + step
  >>     end
  >>   end
  >> end
  => :succ_each
  >>
  >> p [98, 99, 100].succ_each(2).map {|succ_chr| succ_chr.chr}
  ["c", "d", "e"]
  => ["c", "d", "e"]
  >>
  >> [101, 102, 103].succ_each(5) do |succ_chr|
  >>   p succ_chr.chr
  >> end
  "j"
  "k"
  "l"
  => [101, 102, 103]
  ```

  ```ruby
  # 3
  >> class Array
  >>   def succ_each(step = 1)
  >>     return to_enum(:succ_each, step) unless block_given?
  >>
  >>     each do |int|
  >>       yield int + step
  >>     end
  >>   end
  >> end
  => :succ_each
  >>
  >> p [98, 99, 100].succ_each(2).map {|succ_chr| succ_chr.chr}
  ["d", "e", "f"]
  => ["d", "e", "f"]
  >>
  >> [101, 102, 103].succ_each(5) do |succ_chr|
  >>   p succ_chr.chr
  >> end
  "j"
  "k"
  "l"
  => [101, 102, 103]
  ```

  ```ruby
  # 4
  >> class Array
  >>   def succ_each(step = 1)
  >>     unless block_given?
  >>       Enumerator.new do |yielder|
  >>         each do |int|
  >>           yielder << int + step
  >>         end
  >>       end
  >>     else
  >>       each do |int|
  >>         yield int + step
  >>       end
  >>     end
  >>   end
  >> end
  => :succ_each
  ```



  ### 解説

  ブロックを渡す場合と、チェーンを行う場合の両方を考慮する必要があります。

  チェーンを行う場合はEnumeratorオブジェクトを作成する必要があります。

  →作成に必要なメソッド：`enum_for`・`to_enum`


  問題では、`enum_for`を使っていますので選択肢のうち`to_enum`を使っている選択肢が答えのひとつです。

  ただし、`to_enum`は引数にメソッド名と、そのメソッドに必要な引数を指定する必要があります。

  問題では`succ_each`メソッドに引数2を渡していますので、Enumeratorオブジェクトを作成するときに必要になります。

  また、Enumeratorオブジェクトは`new`メソッドで作成することが出来ます。この問題ですと少し冗長ではありますが、全体的には次のとおりです。

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

  なお、チェーンした先で渡されたブロックを評価するためには`Enumerator::Yielder`のオブジェクトを利用します。

  オブジェクトに対して、`<<`を実行することでブロック内で評価した結果を受け取ることが出来ます。



  ## 次のプログラムの`__(1)__`に適切な内容を選択して実行すると、[97, 112, 112, 108, 101]と表示されます。期待した結果を得られるように正しい選択肢を選んでください

  ```ruby
  enum_char = Enumerator.new do |yielder|
    "apple".each_char do |chr|
      __(1)__
    end
  end

  array = enum_char.map do |chr|
    chr.ord
  end

  p array
  ```

  1. `yielder.call chr`

  2. `yielder(chr)`

  3. `yielder << chr`

  4. `yielder.inject chr`



  ### 解説

  `map`メソッドのブロックはEnumeratorオブジェクトをレシーバーとした場合にEnumerator::Yielderオブジェクトとなります。

  この問題のプログラム上では変数`yielder`を指します。

  Enumerator::Yielderを評価するには、`<<`を呼び出します。

  選択肢にある他のメソッドは実装されていません。

  ```ruby
  >> enum_char = Enumerator.new do |yielder|
  >>   "apple".each_char do |chr|
  >>     yielder << chr
  >>   end
  >> end
  => #<Enumerator: #<Enumerator::Generator:0x007fc1bf8dd168>:each>

  >> array = enum_char.map do |chr|
  >>   chr.ord
  >> end
  => [97, 112, 112, 108, 101]

  >> p array
  [97, 112, 112, 108, 101]
  => [97, 112, 112, 108, 101]
  ```



  ## 次のコードを実行するとどうなりますか

  > 一応正解したが、念のため
  >
  > 2018/11/02

  `lazy`は`Enumerator::Lazy`クラスを返します。

  `Enumerator::Lazy`クラスは`map`や`select`メソッドに遅延評価を提供します。


  `take(3)`が実行されると`1`から`3`まで`map`に渡されたものと判断され、`inject`に渡されます。

  よって、答えは`12`になります。

  この時、`4`から`10`までの評価は発生しません。

  ```ruby
  >> p (1..10).lazy.map{|num|
  >>   num * 2
  >> }.take(3).inject(0, &:+)
  12
  => 12
  ```



  ## 以下の実行結果になるように、`__X__`に記述する適切なコードを全て選びなさい

  ```ruby
  p __X__

  # 実行結果
  [1, 4, 9]
  ```

  1. `[1, 2, 3].map{ |x| x ** 2 }`

  1. `[1, 2, 3].collect{ |x| x ** 2 }`

  1. `[1, 2, 3].inject{ |x, y| x + y ** 2 }`

  1. `[1, 2, 3].inject([]){ |x, y| x << y ** 2 }`



  ### 解説

  * `map`・`collect`：対象の配列の各要素をブロック内で評価した結果を配列で返す

  * `inject`：2番目のブロック内の引数に配列の各要素が、1番目のブロック内の引数にブロック内での評価結果が渡される。

  * 最初の要素を評価する時、`inject`の引数を省略した場合、ブロック引数はそれぞれ

    →「配列の1番目の要素、配列の2番目の要素」

  * `inject`の引数を与えた場合はブロック引数はそれぞれ

    →「`inject`に与えた引数、配列の1番目の要素」

  ```ruby
  #
  # 選択肢1
  #
  >> p [1, 2, 3].map{ |x| x ** 2 }
  [1, 4, 9]
  => [1, 4, 9]

  #
  # 選択肢2
  #
  >> p [1, 2, 3].collect{ |x| x ** 2 }
  [1, 4, 9]
  => [1, 4, 9]

  #
  # 選択肢3
  #
  >> p [1, 2, 3].inject{ |x, y| x + y ** 2 }
  14
  => 14

  #
  # 選択肢4
  #
  >> p [1, 2, 3].inject([]){ |x, y| x << y ** 2 }
  [1, 4, 9]
  => [1, 4, 9]
  ```































# ブロック

## 次のコードを実行するとどうなりますか

```ruby
>> def m1(*)
>>   str = yield if block_given?
>>   p "m1 #{str}"
>> end
=> :m1
>>
?> def m2(*)
>>   str = yield if block_given?
>>   p "m2 #{str}"
>> end
=> :m2
>>
?> m1 m2 do
?>   "hello"
>> end
"m2 "
"m1 hello"
=> "m1 hello"
```



### 解説

問題のコードで使用されているメソッド類は以下の通りです。

  * `Kernel#block_given?`はブロックが渡された場合は、真になります。

  * `yield`はブロックの内容を評価します。

  * `{ }`は`do end`よりも結合度が高い為、実行結果に差が出ます。

問題のコードは以下のように解釈されます。

* `m1`の引数と解釈されるため、`m2`の戻り値は`m2`が表示されます。

* `m1`へ`do .. end`のブロックが渡されます。よって、`m1 hello`が表示されます。

```ruby
m1(m2) do
  "hello"
end

# 実行結果
# "m2 "
# "m1 hello"
```

問題のコードを`do ... end`で置き換えた場合は以下の実行結果になります。

> こちらの方だと思って解答してしまった
>
> 2018/10/26

```ruby
m1 m2 {  # m1 (m2 { .. } ) と解釈される
  "hello"
}

# 実行結果
# m2 hello
# m1
```



## 次のコードを実行するとどうなりますか

`lambda`を`call`する際の引数は省略できません。

`lambda`に似た機能に`Proc`があります。

似ていますが、異なる部分もあります。

次の表が`lambda`と`Proc`の違いになります。

|        特徴        |           Proc           |      lambda     |
|:------------------|:-------------------------|:----------------|
|      引数の数      |            曖昧           |       厳密       |
|     引数の渡し方    |       Proc.new { \       |      x, y\      |
|return, brake, next|    call以降が実行されない   |call以降も実行される|



```ruby
>> local = 0
=> 0

>> p1 = lambda { |arg1, arg2|
>>   arg1, arg2 = arg1.to_i, arg2.to_i
>>   local += [arg1, arg2].max
>> }
=> #<Proc:0x007ffdf1084840@(irb):3 (lambda)>

>> p1.call("1", "2")
=> 2
>> p1.call("7", "5")
=> 9
>> p1.call("9")      # こちらは実行されない
ArgumentError: wrong number of arguments (given 1, expected 2)

>> p local
9
=> 9
```



## 次のコードを実行するとどうなりますか

問題のコードで使用されているメソッド類は以下の通りです。

* `Kernel#block_given?`はブロックが渡された場合は、真になります。

* `yield`はブロックの内容を評価します。

* `{ }`は`do end`よりも結合度が高い為、実行結果に差が出ます。

問題のコードは以下のように解釈されます。

1. `m2`へブロックが渡され、`m2 hello`が表示されます。

1. `m1`へは引数が渡され、ブロックは渡されません。よって、`m1`が表示されます。

```ruby
m1 (m2 {
      "hello"
    }
)

# 実行結果
# "m2 hello"
# "m1 "
```

問題のコードを`do end`で置き換えた場合は以下の実行結果になります。

```ruby
m1 m2 do  # m1(m2) do と解釈されます。
  "hello"
end

# 実行結果
# "m2 "
# "m1 hello"
```


```ruby
>> def m1(*)
>>   str = yield if block_given?
>>   p "m1 #{str}"
>> end
=> :m1
>>
?> def m2(*)
>>   str = yield if block_given?
>>   p "m2 #{str}"
>> end
=> :m2
>>
?> m1 m2 {
?>   "hello"
>> }
"m2 hello"
"m1 "
=> "m1 "
```



## 以下のコードの中で文法として正しいものを全て選びなさい

1. `1.upto 5 do |x| puts x end`

1. `1.upto(5) do |x| puts x end`

1. `1.upto 5 { |x| puts x }`

1. `1.upto(5) { |x| puts x }`



### 解説

ブロック引数を`{...}`で囲む場合は、引数の`()`を省略することができない

`do...end`で囲む場合は、引数の`()`を省略することができる

> 誤植かもしれないので、パス
>
> 2018/11/07

```ruby
#
# 選択肢1
#
>> 1.upto 5 do |x| puts x end
1
2
3
4
5
=> 1

#
# 選択肢2
#
>> 1.upto(5) do |x| puts x end
1
2
3
4
5
=> 1

#
# 選択肢3
#
>> 1.upto 5 { |x| puts x }
1
2
3
4
5
=> 1

#
# 選択肢4
#
>> 1.upto(5) { |x| puts x }
1
2
3
4
5
=> 1
```



## 次のプログラムの実行結果を得るために`__(1)__`に適切なメソッドをすべて選んでください。

```ruby
class Array
  def succ_each(step = 1)
    return __(1)__(__method__, step) unless block_given?

    each do |int|
      yield int + step
    end
  end
end

[97, 98, 99].succ_each.map {|int|
  p int.chr
}

# 実行結果
"b"
"c"
"d"
```

```ruby
# 選択肢1
enum

# 選択肢2
enum_chr

# 選択肢3
to_enum

# 選択肢4
enum_for
```



### 解説

ブロックを渡さない場合は、Enumeratorオブジェクトを作成してメソッドをチェーン出来るようにします。

また、ブロックを渡す場合は`yield`で評価してブロックを評価します。

内部イテレータを実装する場合は次のような構造になることが多いです。

レシーバーの配列の要素に対して、`step`だけ数値を進めるようなイテレータを作ります。

簡単化の為、整数値しかない配列だけを想定します。

```ruby
class Array
  def succ_each(step = 1)
    # Enumeratorオブジェクトを作成して、メソッドチェーンできるようにします。
    # ブロックの有無は、block_given?メソッドで判定します。
    return to_enum(:succ_each, step) unless block_given?

    # ブロックを渡された場合の実装です。
    # ブロックはyieldで評価します。
    each do |int|
      yield int + step
    end
  end
end
```

`to_enum`または、`enum_for`でEnumeratorオブジェクトを作成しますが、引数に実行対象のメソッド名をシンボルで指定します。

先程のサンプルコードでは`to_enum(:succ_each, step)`と書いていますが、これは`to_enum(__method__, step)`と書くのと同じです

`__method__`は実行中のメソッドをシンボルで返します。



## 次のコードを実行するとどうなりますか

```ruby
>> local = 0
=> 0

>> p1 = Proc.new { |arg1, arg2|
>>   arg1, arg2 = arg1.to_i, arg2.to_i
>>   local += [arg1, arg2].max
>> }
=> #<Proc:0x007fe1b6883638@(irb):3>

>> p1.call("1", "2")
=> 2
>> p1.call("7", "5")
=> 9
>> p1.call("9")
=> 18

>> p local
18
=> 18
```



### 解説

`Proc`は`call`の際に引数の数を省略することができます。

その際、不足の引数へは`nil`が代入されます。

`Proc`に似た機能に`lambda`があります。

似ていますが、異なる部分もあります。

次の表が`Proc`と`lambda`の違いになります。

|    特徴   |    Proc    |lambda|
|----------|------------|------|
|  引数の数  |     曖昧    |  厳密 |
|引数の渡し方|Proc.new { \| x, y\|

`return`, `brake`, `next`	`call`以降が実行されない	call以降も実行される

そのほか、`lambda`はアロー演算子で定義することができます。



## 以下の実行結果になるように、`__X__`に記述する適切なコードを選びなさい

```ruby
hi = __X__
hi.call("World")

# 実行結果
"Hello, World."
```

1. `->{ |x| puts "Hello, #{x}." }`

1. `->{(x) puts "Hello, #{x}." }`

1. `->(x){ puts "Hello, #{x}." }`

1. `\(x) -> { puts "Hello, #{x}." }`



### 解説

Ruby1.9から新たに追加された記法となります

```ruby
>> hi = ->(x){ puts "Hello, #{x}." }
=> #<Proc:0x007fe4e7825ab8@(irb):1 (lambda)>
>> hi.call("World")
Hello, World.
=> nil
```



## 以下の実行結果になるように、`__X__`に記述する適切なコードを全て選びなさい

> 正解したが、念のため
>
> 2018/11/07

```ruby
__X__

tag(:p) { "Hello, World." }

# 実行結果
<p>Hello, World.</p>
```

* 選択肢1

  ```ruby
  def tag(name)
    puts "<#{name}>#{yield}</#{name}>"
  end
  ```

* 選択肢2

  ```ruby
  def tag(name)
    puts "<#{name}>#{yield.call}</#{name}>"
  end
  ```

* 選択肢3

  ```ruby
  def tag(name, &block)
    puts "<#{name}>#{block}</#{name}>"
  end
  ```

* 選択肢4

  ```ruby
  def tag(name, &block)
    puts "<#{name}>#{block.call}</#{name}>"
  end
  ```



### 解説

ブロック付きメソッドから呼び出し元ブロックを事項するには、

* `yield`を使う

* 引数に`&`をつけた変数を定義し、ブロックにProcオブジェクトとして取得してから、`Proc#call`を呼び出す

```ruby
#
# 選択肢1
#
>> def tag(name)
>>   puts "<#{name}>#{yield}</#{name}>"
>> end
=> :tag
>> tag(:p) { "Hello, World." }
<p>Hello, World.</p>
=> nil

#
# 選択肢2
#
>> def tag(name)
>>   puts "<#{name}>#{yield.call}</#{name}>"
>> end
=> :tag
>> tag(:p) { "Hello, World." }
NoMethodError: undefined method 'call' for "Hello, World.":String

#
# 選択肢3
#
>>   def tag(name, &block)
>>     puts "<#{name}>#{block}</#{name}>"
>>   end
=> :tag
>> tag(:p) { "Hello, World." }
<p>#<Proc:0x007ff2e406a5c0@(irb):4></p>
=> nil

#
# 選択肢4
#
>> def tag(name, &block)
>>   puts "<#{name}>#{block.call}</#{name}>"
>> end
=> :tag
>> tag(:p) { "Hello, World." }
<p>Hello, World.</p>
=> nil
```



## 次のプログラムを実行するとどうなりますか

```ruby
>> val = 100
=> 100

>> def method(val)
>>   yield(15 + val)       # 15 + 100 = 115
>> end
=> :method

>> _proc = Proc.new{|arg| val + arg }   # 115 + 100 = 215
=> #<Proc:0x007fce069018d0@(irb):7>

>> p method(val, &_proc)       # トップレベルのvalを指定
215
=> 215
```

ブロックにあるローカル変数`val`はトップレベルにあるものと同じです。

`100 + (15 + 100)`を計算して、この問題の答えは215です。



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

引数名に`&`を付与することでブロック引数になります。

ブロック引数は他の引数より後に記述します。

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

2. 選択肢2

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

3. 選択肢3

```ruby
>> def bar(&block, n)
>>   block.call
>> end
SyntaxError:

>> bar(5) {
>>   puts "hello, world"
>> }
NoMethodError:
```

4. 選択肢4

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
