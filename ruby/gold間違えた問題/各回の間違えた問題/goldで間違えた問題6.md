goldで間違えた問題
================

## 次のコードを実行するとどうなりますか

```ruby
>> class C
>> end
=> nil
>>
?> module M
>>   refine C do
?>     def m1
>>       100
>>     end
>>   end
>> end
=> #<refinement:C@M>
>>
?> class C
>>   def m1
>>     400
>>   end
>>
?>   def self.using_m
>>     using M
>>   end
>> end
=> :using_m
>>
?> C.using_m
RuntimeError: Module#using is not permitted in methods
>>
?> puts C.new.m1
400
=> nil
```



### 解説

`using`はメソッドの中で呼び出すことは出来ません。呼び出した場合はRuntimeErrorが発生します。



## 次のプログラムを実行するとどうなりますか

```ruby
>> class Object
>>   CONST = "1"
>>   def const_succ
>>     CONST.succ!
>>   end
>> end
=> :const_succ
>>
?> class Child1
>>   const_succ
>>   class << self
>>     const_succ
>>   end
>> end
=> "3"
>>
?> class Child2
>>   const_succ
>>   def initialize
>>     const_succ
>>   end
>> end
=> :initialize
>>
?> Child1.new
=> #<Child1:0x007fe51b0bf538>
>> Child2.new
=> #<Child2:0x007fe51b0bd4e0>
>>
?> p Object::CONST
"5"
=> "5"
```



### 解説

> 正解したが、念のため
>
> 2018/11/16

クラスObjectにメソッドを定義すると特異クラスでもそのメソッドを利用することが出来ます。

問題のプログラムを順に実行すると、答えは"5"になります。

> 補足　Object#const_succについて
> 内部で`String#succ!`を実行しています。このメソッドはレシーバーの文字列を次の文字列へ進めます。
> この問題ですと、"1"→"2"・・・と1ずつ繰り上がります。
> また、定数に対して行っていますが破壊的メソッドの呼び出しですので再代入にはならず警告は表示されません。

```ruby
class Object
  CONST = "1"
  def const_succ
    CONST.succ!
  end
end

class Child1
  const_succ # "2"になる
  class << self
    const_succ # "3"になる
  end
end

class Child2
  const_succ # "4になる"
  def initialize
    const_succ
  end
end

Child1.new # "4"のまま
Child2.new # "5"になる

p Object::CONST
```



##次のプログラムを実行すると`215`が表示されます。`__(1)__`に適切な選択肢を選んでください。

```ruby
val = 100

def method(val)
  yield(15 + val)
end

_proc = Proc.new{|arg| val + arg }

p method(val, __(1)__)
```

1. `&_proc.to_proc`

1. `_proc.to_proc`

1. `_proc`

1. `&_proc`



### 解説

Procオブジェクトをメソッドで実行するにはブロックに変換する必要があります。

`&`をProcオブジェクトのプレフィックスにすることでブロックに変換することが出来ます。

また、`to_proc`メソッドはProcオブジェクトを生成して返します。Procオブジェクトをレシーバに`to_proc`を実行するとレシーバ自身が返ってきます。

この問題の答えは、`&_proc.to_proc`と`&_proc`です。



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



| 版 |  年/月/日 |
|---|----------|
|初版|2018/11/16|
