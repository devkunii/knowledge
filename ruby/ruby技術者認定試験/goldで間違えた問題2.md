goldで間違えた問題
================



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



## 次のコードを実行するとどうなりますか

> 一応正解したが、念のため
>
> 2018/11/02

```ruby
>> ^C
>> class C
>>   @val = 3
>>   attr_accessor :val
>>   class << self
>>     @val = 10
>>   end
>>   def initialize
>>     @val *= 2 if val
>>   end
>> end
=> :initialize

>> c = C.new
=> #<C:0x007fc2e593b1d8>
>> c.val += 10
NoMethodError: undefined method '+' for nil:NilClass
>> p c.val
nil
=> nil
```



### 解説

問題のコードは、13行目で`c.val`が`nil`になり、実行エラーになります。

2行目の`@val`はクラスインスタンス変数といい、特異メソッドからアクセスすることができます。

3行目の`@val`は特異クラスのクラスインスタンス変数です。

この値にアクセスするためには以下のようにアクセスします。

```ruby
class << C
  p @val
end
```

13行目の`c.val`は`attr_accessor`よりアクセスされます。

`initialize`メソッドで初期化が行われていないため、`nil`が返されます。

以下のコードは問題コードに行番号をつけています。

```ruby
1: class C
2:   @val = 3
3:   attr_accessor :val
4:   class << self
5:     @val = 10
6:   end
7:   def initialize
8:     @val *= 2 if val
9:   end
10: end
11:
12: c = C.new
13: c.val += 10
14:
15: p c.val
```



## 次のコードを実行するとどうなりますか

`include`は`Module`のインスタンスメソッドをMix-inするメソッドです。

`C.methods`は`C`の特異メソッドを表示します。

よって、`C#class_m`はインスタンスメソッドです、`C.methods`では表示されません。

```ruby
>> module M
>>   def class_m
>>     "class_m"
>>   end
>> end
=> :class_m
>>
>> class C
>>   include M
>> end
=> C

>> p C.methods.include? :class_m
false
=> false

# オブジェクトを生成した場合は、存在する
>> obj = C.new
=> #<C:0x007f91a7936cd8>
>> p obj.methods.include? :class_m
true
=> true
```



## Rubyで使用可能なオプションではないものを選択しなさい(複数)。

> かなり覚える量が多いので、一回ストップ
>
> 2018/11/02

1. `-t`：存在しない

2. `-l`：行末の自動処理を行います。各行の最後に`String#chop!`を実行。

3. `-p`：`-n`と同じだが`$_`を出力

4. `-f`：存在しない



## 次のプログラムは`Enumerator::Lazy`を使っています。先頭から5つの値を取り出すにはどのメソッドが必要ですか

```ruby
>> (1..100).each.lazy.chunk(&:even?)
=> #<Enumerator::Lazy: #<Enumerator: #<Enumerator::Generator:0x007f91a7933ba0>:each>>
```

1. `take(5)`

  ```ruby
  >> (1..100).each.lazy.chunk(&:even?).take(5)
  => #<Enumerator::Lazy: #<Enumerator::Lazy: #<Enumerator: #<Enumerator::Generator:0x007f91a799bdb8>:each>>:take(5)>
  ```

2. `take(5).force`

  ```ruby
  >> (1..100).each.lazy.chunk(&:even?).take(5).force
  => [[false, [1]], [true, [2]], [false, [3]], [true, [4]], [false, [5]]]
  ```

3. `first(5)`

  ```ruby
  >> (1..100).each.lazy.chunk(&:even?).first(5)
  => [[false, [1]], [true, [2]], [false, [3]], [true, [4]], [false, [5]]]
  ```

4. `first(5).force`

  ```ruby
  >> (1..100).each.lazy.chunk(&:even?).first(5).force
  NoMethodError: undefined method 'force' for #<Array:0x007f91a9048f70>
  ```



### 解説

値を取り出すには、

* `Enumerator::Lazy#force`

* `Enumerator::Lazy#first`

を呼び出す必要があります。

問題文には「先頭から5つ」とあるので、`first(5)`として取り出します。

また、`Enumerator::Lazy#force`で問題文の通りにするには`Enumerator::Lazy#take`も利用します。

`Enumerator::Lazy#take`は`Enumerable#take`と違い`Enumerator::Lazy`のインスタンスを戻り値にします。

そのインスタンスから`Enumerator::Lazy#force`で実際の値を取り出します。



## 次のコードを実行するとどうなりますか

```ruby

```
