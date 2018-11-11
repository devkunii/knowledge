間違えた問題 Refinement
=====================

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



## 次のコードを実行するとどうなりますか

同じメソッドに対して`Refinement`で再定義を2つのモジュールで行っています。

もし、`using`を2行書いたとしても **1つのメソッドで有効になる再定義は1つだけ** です。

最後に書いた`using`から優先されます。

この問題では`using R2`が最後に有効化された`Refinement`です。

有効になる再定義は1つだけですので、モジュール`R2`にある`super`はクラス`C`にある`m1`を呼び出します。

よって、`super + 100`は`100 + 100`となり`200`が表示されるのが正解です。

```ruby
class C
  def m1(value)
    100 + value
  end
end

module R1
  refine C do
    def m1
      super 50
    end
  end
end

module R2
  refine C do
    def m1
      super 100
    end
  end
end

using R1
using R2

puts C.new.m1
=> 200
```

一方で、`using R1`に書いた内容はすべて無効になったかというとそういうわけではありません。

次のサンプルコードだとモジュール`R2`に`m2`が定義されていなくても呼び出すことが出来ます。

```ruby
class C
  def m1(value)
    100 + value
  end

  def m2(value)
    value + ", world"
  end
end

module R1
  refine C do
    def m1
      super 50
    end

    def m2
      super "Hello"
    end
  end
end

module R2
  refine C do
    def m1
      super 100
    end
  end
end

using R1
using R2

puts C.new.m1
puts C.new.m2
=> 200
=> Hello, world
```



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



## 次のコードを実行するとどうなりますか

`using`はメソッドの中で呼び出すことは出来ません。呼び出した場合は`RuntimeError`が発生します。

```ruby
>> class C
>> end
=> nil
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
>> class C
>>   def m1
>>     400
>>   end
>>
>>   def self.using_m
>>     using M
>>   end
>> end
=> :using_m
>>
>> C.using_m
RuntimeError: Module#using is not permitted in methods
>>
>> puts C.new.m1
400
=> nil
```



## 次のコードを実行するとどうなりますか

```ruby
class C
  def m1
    200
  end
end

module R
  refine C do
    def m1
      300
    end
  end
end

using R

class C
  def m1
    100
  end
end

puts C.new.m1
=> 300
```

Refinementで再定義したメソッドの探索ですが、prependより優先して探索が行われます。

例えば、クラスCはクラスBを継承しているとすると次のような順に探索を行います。

`Refinement` -> `prependしたモジュール` -> `クラスC` -> `includeしたモジュール` -> `クラスCの親（クラスB）`

問題では`using`の後にクラスオープンしてメソッドを再定義していますが、Refinementにある`300`が表示されます。

> 探索についてもう一度復習！
>
> 2018/11/03
