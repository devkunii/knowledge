19 Enumeratorクラス
==================

## 目次

* [Enumeratorクラス](#0Enumeratorクラス)

* [Enumeratorオブジェクト](#1Enumeratorオブジェクト)

* [with_index](#2with_index)

* [外部イテレータ](#3外部イテレータ)

* [Enumerator・Lazy](#4EnumeratorLazy)

* [ジェネレータ](#5ジェネレータ)



## 0.Enumeratorクラス

* 繰り返し処理を行う様々なメソッドに対して、`Enumerable`の機能を提供する

  * Stringクラス(Enumerable無)に対して、`each_char`、`each_line`などの繰り返しのメソッドを適用

* 外部イテレータ(繰り返しのタイミングを制御できる方式)と呼ばれる柔軟な処理を提供する



## 1.Enumeratorオブジェクト

* 繰り返し処理を行うメソッドを、ブロック無しで呼び出した場合、`Enumeratorオブジェクト`が返される

```ruby
>> [].each                     # Array#each
=> #<Enumerator: []:each>
>> {}.each                     # Hash#each
=> #<Enumerator: {}:each>
>> (1..10).each                # Range#each
=> #<Enumerator: 1..10:each>
>> ''.each                     # Stringにはeachメソッドは無い(Enumerableがincludeされていない)
NoMethodError: undefined method 'each' for "":String
>> ''.each_char                # String#each_char
=> #<Enumerator: "":each_char>
>> 10.times                    # Integer#times
=> #<Enumerator: 10:times>
>> loop                        # Kernel.#loop
=> #<Enumerator: main:loop>
```

* Enumeratorオブジェクトは、以下の方法で生成される

  * `Object#to_emun`

  * `Object#enum_for`

  * `Enumerator.new`

* 引数には、`each`の代わりに使うメソッド名を渡すことができる

```ruby
# to_enum・enum_forメソッド
>> [1, 2, 3].to_enum
=> #<Enumerator: [1, 2, 3]:each>
>> [1, 2, 3].enum_for
=> #<Enumerator: [1, 2, 3]:each>

# Enumerator.newメソッド
>> str = "xyz"
=> "xyz"
>> enum = Enumerator.new(str, :each_byte)
(irb):18: warning: Enumerator.new without a block is deprecated; use Object#to_enum
=> #<Enumerator: "xyz":each_byte>

# enum_forを用いた例
>> 'Alice'.enum_for(:each_char)
=> #<Enumerator: "Alice":each_char>
```

> 実際の例
>
> String#each_lineを元にEnumeratorオブジェクトを生成し、Enumerable#mapにより行列の配列を生成

```ruby
>> lines = <<EOM
Alice
Bob
Charlie
EOM
=> "Alice\nBob\nCharlie\n"
>>
>> enum = lines.each_line
=> #<Enumerator: "Alice\nBob\nCharlie\n":each_line>
>> enum.map {|line| line.length}
=> [6, 4, 8]
```



## 2.with_index

* 生成時のパラメータに従って、要素にインデックスを添えて繰り返す

* 繰り返しのたびに、インクリメントされる値と一緒に繰り返し処理を行うことができる

* インデックスは offset から始まる

* ブロックを指定した場合の戻り値は生成時に指定したレシーバ自身

> 例
>
> 1つ目：単に繰り返しを行う例
>
> 2つ目：インデックスと一緒に繰り返したい場合の例
>
> Enumerable#selectの条件に、配列の添字を使うことができる

```ruby
# 1つ目(Enumerable#each_with_index)
>> %w(Alice Bob Charlie).each_with_index do |name, index|
>>   puts "#{index}: #{name}"
>> end
0: Alice
1: Bob
2: Charlie
=> ["Alice", "Bob", "Charlie"]

# 1つ目(Enumerator#with_index)
>> %w(Alice Bob Charlie).each.with_index do |name, index|
>>   puts "#{index}: #{name}"
>> end
0: Alice
1: Bob
2: Charlie
=> ["Alice", "Bob", "Charlie"]

# 2つ目
>> %w(Alice Bob Charlie).select.with_index do |name, index|
>>   index > 0
>> end
=> ["Bob", "Charlie"]
```



## 3.外部イテレータ

* 内部イテレータ：繰り返し処理は、これらのメソッドに渡したブロックの中で行われる

  * 繰り返しを簡潔に記述でき、本質的な処理に集中できる

  * 一方で、2つの配列を同時に繰り返したい場合には、Array#eachではうまくいかなくなる(繰り返しの処理は、each内部で行われている)



### 外部イテレータ

* `外部イテレータ`：繰り返しの処理のタイミンングを制御できる

* Enumeratorオブジェクトには、以下の方法で`外部イテレータ`を制御できる

  * `next`：次の要素を返し、内部で指している要素の位置を1つ先に進める

  * `rewind`：初めから繰り返しを行う

```ruby
>> enum = [4, 4, 2, 3].to_enum
=> #<Enumerator: [4, 4, 2, 3]:each>
>>
>> enum.next
=> 4
>> enum.next
=> 4
>> enum.next
=> 2
>>
>> enum.rewind            # 始めから繰り返しをやり直す
=> #<Enumerator: [4, 4, 2, 3]:each>
>>
>> enum.next
=> 4
>> enum.next
=> 4
>> enum.next
=> 2
>> enum.next
=> 3
>> enum.next
StopIteration: iteration reached an end
```

* Enumerator#nextは、次の要素が無い場合は例外`StopIteration`を発生させる

* 自分で`rescue`することもできるが、Kernel#loopを使って制御することもできる

  * ブロック内で発生した例外`StopIteration`を補足して、ループを終了させる

```ruby
>> enum = [4, 4, 2, 3].to_enum
=> #<Enumerator: [4, 4, 2, 3]:each>
>>
>> loop do
>>   puts enum.next
>> end
4
4
2
3
=> [4, 4, 2, 3]
```

* 繰り返しを途中で止めて、他の処理を挟むことができたり、複数の繰り返し処理を同時に行うことができる

> 例
>
> Enumeratorオブジェクトを使って、人物の名前の配列と年齢の配列を同時に繰り返して出力

```ruby
>> people = %w(Alice Bob Charlie).to_enum
=> #<Enumerator: ["Alice", "Bob", "Charlie"]:each>
>> ages   = [14, 32, 28].to_enum
=> #<Enumerator: [14, 32, 28]:each>

>> loop do
>>   person = people.next
>>   age    = ages.next
>>
>>   puts "#{person} (#{age})"
>> end
Alice (14)
Bob (32)
Charlie (28)
=> ["Alice", "Bob", "Charlie"]
```

* `feed`：map、selectのようなブロックの戻り値を使用するタイプのメソッドを呼び出す。

  * 戻り値に当たる値を渡す

  * メソッド自体の戻り値は、StopIteration#resultで得ることができる

  * Enumerator 内部の yield が返す値を設定

  * これで値を設定しなかった場合は yield は nil

  * この値は内部で yield された時点でクリア

> 例
>
> Enumeratorを用いて、正規表現にマッチする要素を配列から取得

```ruby
>> enum = %w(Alice Bob Charlie).select
=> #<Enumerator: ["Alice", "Bob", "Charlie"]:select>

# /li/にマッチする要素だけを得る
>> loop do
>>   begin
>>     person = enum.next
>>
>>     enum.feed /li/ === person     # ブロックの戻り値に当たる値
>>   rescue StopIteration => e
>>     p e.result                    # selectの戻り値を表示させる
>>     break
>>   end
>> end
["Alice", "Charlie"]                 # selectの戻り値
=> nil
```



## 4.Enumerator・Lazy

* 繰り返し処理の実行を遅延させることができる

  * 大きな配列や無限の要素を持つオブジェクトの集まりを手軽に扱える

```ruby
# 無限にある数値をEnumerable#mapが全て繰り返す
# 結果は帰ってこない
#
>> (0..Float::INFINITY).map {|n| n.succ}.select {|n| n.odd?}.take(3) # 返ってこない
=> IRB::Abort: abort then interrupt!                                 # 強制終了
```

```ruby
# 全ての要素を処理することが不可能であったり、時間や空間を要する場合には、Enumerator::Lazyを用いる
#
>> odd_numbers = (0..Float::INFINITY).lazy.map {|n| n.succ}.select {|n| n.odd?}.take(3)
=> #<Enumerator::Lazy: #<Enumerator::Lazy: #<Enumerator::Lazy: #<Enumerator::Lazy: 0..Infinity>:map>:select>:take(3)>
>> odd_numbers.force
=> [1, 3, 5]
```

* `Enumerator::Lazy`に対してメソッドを呼び出すと、戻り値として新しい`Enumerator::Lazy`オブジェクトが得られる

  * この時点では、実際の処理は行われない

* 値を取り出すには、

  * `Enumerator::Lazy#force`：必要な回数だけ処理が実行され、値が得られる

  * `Enumerator::Lazy#first`：最初の要素を返す。また、回数の指定があった場合は、先頭からn番目までの値が得られる

を呼び出す必要があります。

* `Enumerable#lazy`では、交互で値を出力することができない。また、余分な中間データが生成する

  * `Enumerator::Lazy`では、それらの問題を解決する

> 例
>
> 繰り返しの順番を確かめる

```ruby
# Enumerator::Lazy
>> (0..Float::INFINITY).lazy.map {|n|
>>   puts "map: #{n}"
>>   n.succ
>> }.select {|n|
>>   puts "select: #{n}"
>>   n.odd?
>> }.take(3).force
map: 0
select: 1
map: 1
select: 2
map: 2
select: 3
map: 3
select: 4
map: 4
select: 5
=> [1, 3, 5]
```



## 5.ジェネレータ

* ジェネレータ：部分的な結果だけを計算して返し、次回また続きから計算できるサブルーチンのこと

* 実際の計算を最小限で済ませ、大きなデータを少ないリソースで扱うことができる

### Enumerator::Yielder

* Enumerator.new にブロックを渡すと、ブロック引数として Enumerator::Yielder のインスタンスが渡される。

* 生成された Enumerator オブジェクトに対して`each`を呼ぶと、この生成時に指定されたブロックを実行

* Yielderオブジェクトに対して`<<`メソッドが呼ばれるたびに、 `each`に渡されたブロックが繰り返される(値が返される)

* `new`に渡されたブロックが終了した時点で`each`の繰り返しが終わります。 このときのブロックの返り値が`each`の返り値となります。

```ruby
>> enum = Enumerator.new do |yielder|
>>   yielder.yield 1
>>   yielder.yield 2
>>   yielder.yield 3
>> end
=> #<Enumerator: #<Enumerator::Generator:0x007fc1bf036848>:each>
>> enum.next()
=> 1
>> enum.next()
=> 2
>> enum.next()
=> 3
>> enum.next()
StopIteration: iteration reached an end
```



| 版     | 年/月/日   |
| ------ | ---------- |
| 初版   | 2018/10/13 |
| 第二版 | 2018/11/02 |
| 第三版 | 2019/05/11 |
