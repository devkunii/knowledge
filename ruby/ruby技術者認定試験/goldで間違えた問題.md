goldで間違えた問題
================

## 次のプログラムは"Hello, world"と表示します。同じ結果になる選択肢はどれですか（複数選択）

```ruby
>> module M
>>   CONST = "Hello, world"
>>
>>   class C
>>     def awesome_method
>>       CONST
>>     end
>>   end
>> end
(irb):21: warning: already initialized constant M::CONST
(irb):6: warning: previous definition of CONST was here
=> :awesome_method

>> p M::C.new.awesome_method
"Hello, world"
=> "Hello, world"
```



### 選択肢1

定数の参照はレキシカルに行われます。

`M::C#awesome_method`のコンテキストに`CONST`がないため例外が発生します。

```ruby
>> module M
>>   CONST = "Hello, world"
>> end
=> "Hello, world"
>>
?> class M::C
>>   def awesome_method
>>     CONST
>>   end
>> end
=> :awesome_method
>>
>> p M::C.new.awesome_method
NameError: uninitialized constant M::C::CONST
```



### 選択肢2

`class_eval`にブロックを渡した場合は、ブロック内のネストはモジュール`M`になります。

そのコンテキストから定数を探しますので`"Hello, world"`が表示されます。

```ruby
>> class C
>> end
=> nil
>>
?> module M
>>   CONST = "Hello, world"
>>
?>   C.class_eval do
?>     def awesome_method
>>       CONST
>>     end
>>   end
>> end
=> :awesome_method
>>
>> p C.new.awesome_method
"Hello, world"
=> "Hello, world"
```



### 選択肢3

`class_eval`に文字列を渡した場合のネストの状態はクラス`C`です。

`CONST`はクラス`C`にありますので`"Hello, world"`が表示されます。

```ruby
>> class C
>>   CONST = "Hello, world"
>> end
=> "Hello, world"
>>
?> module M
>>   C.class_eval(<<-CODE)
    def awesome_method
      CONST
    end
  CODE
>> end
=> :awesome_method
>>
>> p C.new.awesome_method
"Hello, world"
=> "Hello, world"
```



### 選択肢4

`class_eval`にブロックを渡した場合は、ブロック内のネストはモジュール`M`になります。

そのコンテキストから定数を探しますがないため例外が発生します。

```ruby
>> class C
>>   CONST = "Hello, world"
>> end
=> "Hello, world"
>>
>> module M
>>   C.class_eval do
>>     def awesome_method
>>       CONST
>>     end
>>   end
>> end
=> :awesome_method
>>
>> p C.new.awesome_method
NameError: uninitialized constant M::CONST
```



## 以下のコードを実行するとどうなりますか

`C#initialize`が`S#initialize`をオーバーライドされているため、`@@val += 1`は実行されません。

`class << C ~ end`の処理はクラスを定義した時点で、実行されます。

```ruby
>> class S
>>   @@val = 0
>>   def initialize
>>     @@val += 1
>>   end
>> end
=> :initialize
>>
>> class C < S
>>   class << C
>>     @@val += 1
>>   end
>>
>>   def initialize
>>   end
>> end
=> :initialize
>>
>> C.new                   # initializeで実行されていない(オーバーライドされている)
=> #<C:0x007f826e078d00>
>> C.new                   # initializeで実行されていない(オーバーライドされている)
=> #<C:0x007f826e073328>
>> S.new
=> #<S:0x007f826e0719b0>
>> S.new
=> #<S:0x007f826e070038>
>>
>> p C.class_variable_get(:@@val)
3
=> 3
```



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



## 次のコードを実行するとどうなりますか

`include`はモジュールのメソッドをインスタンスメソッドとして追加します。

メソッド探索順は`self`の後に追加されます。

複数回`include`された場合は、後に宣言された方からメソッド探索されます。

![多重インクルード(1)](./images/gold/多重インクルード(1).png)

```ruby
>> module M1
>> end
=> nil
>>
>> module M2
>> end
=> nil
>>
>> class C
>>   include M1
>>   include M2
>> end
=> C
>>
>> p C.ancestors
[C, M2, M1, Object, Kernel, BasicObject]
=> [C, M2, M1, Object, Kernel, BasicObject]
```



## 次のプログラムを実行するとどうなりますか

```ruby
>> module M1
>>   class C1
>>     CONST = "001"
>>   end
>>
>>   class C2 < C1
>>     CONST = "010"
>>
>>     module M2
>>       CONST = "011"
>>
>>       class Ca
>>         CONST = "100"
>>       end
>>
>>       class Cb < Ca
>>         p CONST
>>       end
>>     end
>>   end
>> end
"011"
=> "011"
```



### 解説

Rubyは定数の参照はレキシカルに決定されます。

名前空間ではなく、プログラム上の定義された場所と使われている場所の静的な位置づけが重要です。

例えば、次のプログラムでは期待した結果が得られません。`CONST`がモジュールMのスコープにあるためです。

```ruby
>> module M
>>   CONST = "Hello, world"
>> end
=> "Hello, world"
>>
>> class M::C
>>   def awesome_method
>>     CONST
>>   end
>> end
=> :awesome_method
>>
>> p M::C.new.awesome_method
NameError: uninitialized constant M::C::CONST
```

一方で同じレキシカルスコープにある場合は例外は発生しません。

```ruby
>> module M
>>   CONST = "Hello, world"
>>
>>   class C
>>     def awesome_method
>>       CONST
>>     end
>>   end
>> end
=> :awesome_method
>>
>> p M::C.new.awesome_method
"Hello, world"
=> "Hello, world"
```

また、使われている定数の場所がネストされている場合は内側から順に定数の探索が始まります。

レキシカルスコープに定数がない場合は、スーパークラスの探索を行います。

クラス`Cb`から最も物理的に近いのは`M2::CONST`であるため答えは`"011"`になります。

スーパークラスの探索はこの場合には行われません。



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

→こちらの方だと思って解答してしまった(2018/10/26)

```ruby
m1 m2 {  # m1 (m2 { .. } ) と解釈される
  "hello"
}

# 実行結果
# m2 hello
# m1
```



## 次のコードを実行するとどうなりますか

`Refinement`は有効化したスコープのみに影響を与えることが出来ます。

この問題ではクラスオープンした際に`using`で`Refinement`を有効化していますが、

スコープ外は無効になります。

よって、`puts C.new.m1`とした結果は`400`になります。

```ruby
>> class C
>>   def m1
>>     400
>>   end
>> end
=> :m1
>>
>> module M
>>   refine C do
>>     def m1
>>       100
>>     end
>>   end
>> end
=> #<refinement:C@M>
>>
>> class C       # クラスの再オープン時に、refinentを定義している
>>   using M
>> end
=> C
>>
>> puts C.new.m1
400
=> nil
```



## 以下のコードを実行するとどうなりますか

`initialize`の可視性は`private`に設定されています。

`initialize`の可視性を`public`に設定したとしても、必ずprivateになります。

```ruby
>> class C
>> private
>>   def initialize
>>   end
>> end
=> :initialize
>>
>> p C.new.public_methods.include? :initialize
false
=> false
```



## 次のコードを実行するとどうなりますか。

`Class#name`はクラス名を文字列で返します。

`Human#name`クラスは`Class#name`をオーバーライドしているので、`const_get`が呼ばれます。

`const_get`は、`self`に定義された定数を探索します。自クラスに定義がない場合は、メソッドと同様に探索を行います。

問題コードの5行目時点のインスタンスは`Fukuzawa`クラスです。

よって、`Human#name`は`Fukuzawa`クラスの`Yukichi`を返します。

```ruby
>> class Human
>>   NAME = "Unknown"
>>
>>   def self.name
>>     const_get(:NAME)
>>   end
>> end
=> :name
>>
>> class Fukuzawa < Human
>>   NAME = "Yukichi"
>> end
=> "Yukichi"
>>
>> puts Fukuzawa.name
Yukichi
=> nil
```



## 次のコードを実行するとどうなりますか

`method_missing`は、継承チェーンを辿った末にメソッドが見つからなかった場合に、呼び出されます。

`method_missing`も継承チェーンを辿ります。

よって、`B#method_missing`が出力されます。

```Ruby
>> module M
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
>>   def method_missing(id, *args)
>>     puts "B#method_missing"
>>   end
>> end
=> :method_missing
>>
>> obj = B.new
=> #<B:0x007f876f01e188>
>> obj.dummy_method
B#method_missing
=> nil
```
