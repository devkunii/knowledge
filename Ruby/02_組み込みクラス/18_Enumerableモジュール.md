18 Enumerableモジュール
======================

## 目次

* [Enumerableモジュールとは](#0Enumerableモジュールとは)

* [map・collectメソッド](#1mapcollectメソッド)

* [each_with_indexメソッド](#2each_with_indexメソッド)

* [inject・reduceメソッド](#3injectreduceメソッド)

* [each_slice・each_consメソッド](#4each_sliceeach_consメソッド)

* [reverse_eachメソッド](#5reverse_eachメソッド)

* [all?・any?・none?・one?・member?・include?メソッド](#6allanynoneonememberincludeメソッド)

* [find・find_index・selectなどのメソッド](#7findfind_indexselectなどのメソッド)

* [sort・sort_byメソッド](#8sortsort_byメソッド)

* [max・minメソッド](#9maxminメソッド)

* [countメソッド](#10countメソッド)

* [cycleメソッド](#11cycleメソッド)

* [group_byメソッド](#12group_byメソッド)

* [zipメソッド](#13zipメソッド)

* [first・takeメソッド](#14firsttakeメソッド)

* [take_while・dropメソッド](#15take_whiledropメソッド)

* [drop_whileメソッド](#16drop_whileメソッド)

* [select・rejectメソッド](#17selectrejectメソッド)

* [lazyメソッド](#18lazyメソッド)

* [partitionメソッド](#19partitionメソッド)



## 0.Enumerableモジュールとは

* `Array`、`Hash`クラスにインクルードされている

* 全てのメソッドが`each`メソッドを元に定義されているため、`each`メソッドが定義されているクラスであれば、そのクラスでも利用可能



## 1.map・collectメソッド

* `map`、`collect`：与えられたブロックを評価した結果の配列を返す

```ruby
>> [1, 2, 3, 4, 5].map{|i| i ** 2}
=> [1, 4, 9, 16, 25]
```

## 2.each_with_indexメソッド

* `each_with_index`：要素とそのインデックスをブロックに渡して繰り返す

```ruby
>> [:a, :b, :c, :d, :e].each_with_index{|v, i| puts "#{v} => #{i}"}
a => 0
b => 1
c => 2
d => 3
e => 4
=> [:a, :b, :c, :d, :e]
```

## 3.inject・reduceメソッド

* `Enumerable#inject`：ブロックを使用して繰り返し計算を行う

  * 自身のたたみこみ演算を行う(初期値と自身の要素を順に組み合わせて結果を返す)

  * 引数を省略した場合は、配列の先頭がブロック引数の1番目に渡されます。

  * 引数を指定した場合は、その値が初期値になります。

  * ブロック引数の1番目は前回の戻り値が渡されます。初回は、初期値が渡されます。

  * ブロック引数の2番目は要素が順番に渡されます

```ruby
>> [1, 2, 3, 4, 5].inject(0) {|result, v| result + v ** 2}  # 1から5までの数値の2乗の和を求める(1+4+9+16+25=55)
=> 55
```



## 4.each_slice・each_consメソッド

* `each_slice`：要素を指定された数で区切ってブロックに渡す

  * 要素数が指定された数で割きれない場合は、最後だけ渡される数が少なくなる

* `each_cons`：先頭から要素を1つずつ選び、さらに余分に指定された数に合うように要素を選び、それらをブロックに渡していく

```ruby
# each_consメソッド
>> (1..10).each_cons(3) {|items| p items}
[1, 2, 3]
[2, 3, 4]
[3, 4, 5]
[4, 5, 6]
[5, 6, 7]
[6, 7, 8]
[7, 8, 9]
[8, 9, 10]
=> nil

# each_sliceメソッド
>> (1..10).each_slice(3) {|items| p items}
[1, 2, 3]
[4, 5, 6]
[7, 8, 9]
[10]
=> nil
```



## 5.reverse_eachメソッド

* `reverse_each`：`each`メソッドとは逆順にブロックに要素を渡して繰り返す

```ruby
>> [1, 2, 3, 4, 5].reverse_each {|i| puts i}
5
4
3
2
1
=> [1, 2, 3, 4, 5]
```



## 6.all?・any?・none?・one?・member?・include?メソッド

* `all?`：全ての要素が真であれば`true`を返す

* `any?`：真である要素が1つでもあれば`true`を返す

* `none?`：全ての要素が偽であれば`true`を返す

* `one?`：1つの要素だけが真であれば`true`を返す

* `member?`、`include?`：指定された値と`==`メソッドが`true`となる要素がある場合に`true`を返す

```ruby
>> [1, nil, 3].all?
=> false
>> [1, nil, 3].any?
=> true
>> [].all?
=> true
>> [].any?
=> false

# include?メソッド
>> [1, 2, 3, 4, 5].include?(3)
=> true
```



## 7.find・find_index・selectなどのメソッド

* `find`、`detect`：ブロックを評価して最初に真となる要素を返す

* `find_index`：要素の代わりにインデックスを返す

* `find_all`、`select`：ブロックの評価が真となる全ての要素を返す

* `reject`：偽になった全ての要素を返す

* `grep`：指定したパターンとマッチする(`==`メソッドが`true`となる)要素を全て含んだ配列を返す

```ruby
>> [1, 2, 3, 4, 5].find {|i| i % 2 == 0}
=> 2
>> [1, 2, 3, 4, 5].find_index {|i| i % 2 == 0}
=> 1
>> [1, 2, 3, 4, 5].select {|i| i % 2 == 0}
=> [2, 4]
```



## 8.sort・sort_byメソッド

* `sort`：要素を`<=>`メソッドで比較して昇順にソートした配列を、新たに生成して返す

  * ブロックをとる場合は、ブロックの評価結果を元にソートする

* `sort_by`：ブロックの評価結果を`<=>`メソッドで比較して昇順にソートした配列を使って、元の配列をソートした新しい配列を生成して返す

```ruby
>> ["aaa", "b", "cc"].sort{|a, b| a.length <=> b.length} # 1番目の引数の方が大きい場合は、正の数・・・徐々に大きくなる
=> ["b", "cc", "aaa"]
>> ["aaa", "b", "cc"].sort_by{|a| a.length}
=> ["b", "cc", "aaa"]
```



## 9.max・minメソッド

* `max`：要素の最大値を返す

* `min`：要素の最小値を返す

  * `<=>`メソッドで比較するため、全ての要素がそれに対応する必要がある

  * ブロックを渡すと、ブロックの評価結果を元に大小判定を行う

* `max_by`、`min_by`：ブロックの評価結果が最大であった要素を返す

```ruby
>> (1..10).map{|v| v % 5 + v}
=> [2, 4, 6, 8, 5, 7, 9, 11, 13, 10]
>> (1..10).max{|a, b| (a % 5 + a) <=> (b % 5 + b)}
=> 9
>> (1..10).max_by{|v| v % 5 + v}
=> 9
```



## 10.countメソッド

* `count`：要素数を返す

```ruby
>> [1, 2, 3, 4, 5].count
=> 5
```



## 11.cycleメソッド

* `cycle`：要素を先頭から順に取り出し、末尾まで到達したら再度先頭に戻り、それを繰り返す

```ruby
>> [:a, :b, :c].cycle{|v| p v}
:a
:b
:c
:a
# 省略
```



## 12.group_byメソッド

* `group_by`：ブロックの評価結果をキーとし、同じキーを持つ要素を配列としたハッシュを返す

```ruby
>> (1..10).group_by{|v| v % 2}
=> {1=>[1, 3, 5, 7, 9], 0=>[2, 4, 6, 8, 10]} # 余りは1か0のどれか
```



## 13.zipメソッド

* `zip`：自身と引数に指定した配列から、1つずつ要素を取り出して配列を作り、それを要素とする配列を返す

```ruby
>> [:a, :b, :c].zip([1, 2, 3], ["a", "b", "c"])
=> [[:a, 1, "a"], [:b, 2, "b"], [:c, 3, "c"]]
```



## 14.first・takeメソッド

* `take`：先頭から指定した数の要素を配列として返す

* `first`：`take`メソッドと同じだが、数を指定しない場合に先頭の要素のみを返す

```ruby
>> [:a, :b, :c].take(2)
=> [:a, :b]
>> [:a, :b, :c].first
=> :a
```



## 15.take_while・dropメソッド

* `take_while`：先頭からブロックを評価し、最初に偽になった要素の直前までを返す

* `drop`：`take`メソッドとは逆に、先頭から指定した数の要素を取り除いた残りの要素を配列として返す

```ruby
# take_whileメソッド
>> [:a, :b, :c, :d, :e].take_while { |e| e != :d }
=> [:a, :b, :c]

# dropメソッド
>> [:a, :b, :c, :d, :e].drop(3)
=> [:d, :e]
```



## 16.drop_whileメソッド

* `drop_while`：先頭からブロックを評価し、最初に偽になった要素の手前までを切り捨て、残りの要素を配列として返す

```ruby
>> [:a, :b, :c, :d, :e].drop_while { |e| e != :c}
=> [:c, :d, :e]
```



## 17.select・rejectメソッド

* `select`：各要素に対してブロックの評価結果が真であった要素を含む配列を返す

* `reject`：ブロックの評価結果が偽であった要素を含む配列を返す

```ruby
>> [1, 2, 3, 4, 5].select { |e| e % 2 == 0 }
=> [2, 4]
>> [1, 2, 3, 4, 5].reject { |e| e % 2 == 0 }
=> [1, 3, 5]
```



## 18.lazyメソッド

* `lazy`：`map`、`select`メソッドなどのメソッドが、遅延評価を行うように再定義される

  * 遅延評価になるとそれぞれのメソッドが配列でなく`Enumerator::Lazy`を返すようになる

  * そのため、メソッドを評価するタイミングを文字通り遅らせることができる

```ruby
>> a = [1, 2, 3, 4, 5].lazy.select { |e| e % 2 == 0 }
=> #<Enumerator::Lazy: #<Enumerator::Lazy: [1, 2, 3, 4, 5]>:select>
>> b = a.map { |e| e * 2}
=> #<Enumerator::Lazy: #<Enumerator::Lazy: #<Enumerator::Lazy: [1, 2, 3, 4, 5]>:select>:map>
>> c = a.take(3)
=> #<Enumerator::Lazy: #<Enumerator::Lazy: #<Enumerator::Lazy: [1, 2, 3, 4, 5]>:select>:take(3)>
>> c.to_a  # ここで評価される
=> [2, 4]
```



## 19.partitionメソッド

* `partition`：各要素を、ブロックの条件を満たす要素と満たさない要素に分割する

* 各要素に対してブロックを評価して、その値が真であった要素の配列と、 偽であった要素の配列の2つを配列に入れて返します。

ブロックを省略した場合は、各要素に対しブロックを評価し

  * 上のようにその値が真であった要素の配列と、 偽であった要素の配列のペアを返すような Enumerator を 返します。

```ruby
>> a = (1..5).partition(&:odd?)
=> [[1, 3, 5], [2, 4]]
>> p a
[[1, 3, 5], [2, 4]]
=> [[1, 3, 5], [2, 4]]
```



| 版     | 年/月/日   |
| ------ | ---------- |
| 初版   | 2018/10/13 |
| 第二版 | 2018/11/02 |
| 第三版 | 2019/05/11 |
