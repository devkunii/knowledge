間違えた問題 Enumerator
=====================

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
