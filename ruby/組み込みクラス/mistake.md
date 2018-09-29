
### 次のコードを実行するとどうなりますか

```ruby
arr = (1..30).to_a
container = []

arr.each_slice(7) do |i|
  container << i
end

p container.length
5
=> 5
>> container
=> [[1, 2, 3, 4, 5, 6, 7], [8, 9, 10, 11, 12, 13, 14], [15, 16, 17, 18, 19, 20, 21], [22, 23, 24, 25, 26, 27, 28], [29, 30]]
# 配列containerに5個の配列
```

***

### 次のプログラムを実行するとどうなりますか

```ruby
>> raise ['Error Message']
TypeError: exception class/object expected
```

`StandardError`を継承しないクラスのインスタンスを`raise`メソッドの引数に指定すると、
`TypeError`が発生し、メッセージが表示されます。

* `TypeError`：メソッドの引数に期待される型ではないオブジェクトや、期待される振る舞いを持たないオブジェクトが渡された時に発生します。
→`StandardError`の下のクラス

***

### 次のコードを実行するとどうなりますか

```ruby
>> a, = (1..5).partition(&:odd?)
=> [[1, 3, 5], [2, 4]]
>> p a
[1, 3, 5]
=> [1, 3, 5]

# 例
>> a,b = (1..5).partition(&:odd?)
=> [[1, 3, 5], [2, 4]]
>> p b
[2, 4]
=> [2, 4]
```

* `Enumerable#partition`：ブロックの条件を満たす要素と満たさない要素に分割します。

`a,=`と指定すると、`Enumerable#partition`の条件を満たす要素がaに入ります。この場合は、`a`のみに値が入る
→`a,b=`だと、`b`にも値が入る

#### 構文の解説

* `&`：ブロックとして展開することを意味する。配列で定義します、という`*`と同じ役割

* `&:odd?`：`Integer#odd?`メソッド(数値が奇数か判定)を、ブロックとして渡す。

* `(1..5).partition(&:odd?)`：1から5の間で、奇数の配列を分割

***

### 実行してもエラーにならないコードを選べ

```ruby
>> def bar(*n1, n2, *n3)
>>   puts n1
>>   puts n2
>> end

>> def bar(*n1, n2)
>>   puts n1
>>   puts n2
>> end
=> :bar
>> bar 5, 6, 7, 8
5
6
7
8
```

* 可変長引数は、2つ定義することはできません。

```ruby
>> def bar(n1, *n2, n3)
>>   puts "n1: #{n1}, n2: #{n2}, n3: #{n3}" # n1: 5, n2: [6, 7], n3: 8
>> end
=> :bar
>> bar 5, 6, 7, 8
n1: 5, n2: [6, 7], n3: 8
=> nil
```

***

### 次のコードを実行するとどうなりますか

```ruby
>> str = "Liberty Fish   \r\n"
=> "Liberty Fish   \r\n"
>> str.strip!
=> "Liberty Fish"
>> p str
"Liberty Fish"
=> "Liberty Fish"
```

***

#### メソッド名	説明

* `strip`：文字列の先頭と末尾の空白文字(\t\r\n\f\v)を取り除きます。

* `chomp`：末尾から改行コードを取り除きます。

* `chop`：末尾の文字を取り除きます。ただし、文字列の末尾が"\r\n"であれば、2文字とも取り除きます。

```ruby
>> str = "Liberty Fish   \r\n"
=> "Liberty Fish   \r\n"

# chopメソッド
>> str.chop
=> "Liberty Fish   "
>> str.chop.chop.chop.chop.chop
=> "Liberty Fis"

# chompメソッド
>> str.chomp
=> "Liberty Fish   "

# stripメソッド
>> str.strip
=> "Liberty Fish"
```

***

### 次のコードを実行するとどうなりますか

```ruby
>> (10..15).to_a.map.with_index(1) do |elem, i|
?>   puts i
>> end
1
2
3
4
5
6
=> [nil, nil, nil, nil, nil, nil]
```

#### 解説

* `Enumerator#with_index(offset)`：は要素にインデックスを添えてを繰り返します。インデックスはoffsetから開始します。

```ruby
# Enumerable#each_with_indexメソッド
>> (10..15).each_with_index.map { |user, i| "#{i+1}: #{user}" }
=> ["1: 10", "2: 11", "3: 12", "4: 13", "5: 14", "6: 15"]

# Enumerator#with_index(offset)メソッド
>> (10..15).map.with_index { |user, i| "#{i+1}: #{user}" }
=> ["1: 10", "2: 11", "3: 12", "4: 13", "5: 14", "6: 15"]

# offset指定時
>> (10..15).map.with_index(1) { |user, i| "#{i}: #{user}" }
=> ["1: 10", "2: 11", "3: 12", "4: 13", "5: 14", "6: 15"]
```

***

### 次のコードを実行するとどうなりますか

* `Hash#each`のブロックパラメータは`Array`で渡されます。

```ruby
>> h = {a: 100, b: 200}
=> {:a=>100, :b=>200}
>> h.each {|p|
?>   p p.class
>> }
Array
Array
=> {:a=>100, :b=>200}
```

***

### 次のコードを実行するとどうなりますか

* メソッドと変数の探索順位は変数が先です。

```ruby
>> hoge = 0
=> 0
>> def hoge
>>   x = 0
>>   5.times do |i|
?>     x += 1
>>   end
>>   x
>> end
=> :hoge
>> puts hoge
0
=> nil
```

***

### 以下のコードを実行するとどうなりますか

Rubyではメソッド内で定数を定義することができません。
複数回メソッドを呼び出した場合に、定数が不定となるため定義できません。
宣言された場合は、`SyntaxError`が発生します。

```ruby
>> def hoge
>>   x = 10
>>   Y = x < 10 ? "C" : "D"
>>   puts Y
>> end
SyntaxError: (irb):153: dynamic constant assignment
```

***

### 以下のコードを実行するとどうなりますか

大文字アルファベットから始まる識別子は定数です。
Rubyの定数は警告が表示された上で、上書きが可能です。

```ruby
>> X = 10
=> 10
>> X = X < 10 ? "C" : "D"
(irb):159: warning: already initialized constant X
(irb):158: warning: previous definition of X was here
=> "D"
>> puts X
D
=> nil
```

***

### 次のコードを実行するとどうなりますか

```ruby
# 問題
s = <<'EOF'
Hello,
Ruby
EOF
'EOF'

p s

# 解答
>> s = <<'EOF'
Hello,
Ruby
EOF
=> "Hello,\nRuby\n"
>> 'EOF'
=> "EOF"
>>
?> p s
"Hello,\nRuby\n"
=> "Hello,\nRuby\n"
```

#### 識別子の開始ラベルによってヒアドキュメントの解釈の方法が異なります。

* 開始ラベル：説明
* `"識別子"`：式展開が有効
* `識別子`：式展開が有効。ダブルクオートと同じ結果
* `'識別子'`：式展開できない(シングルクオート)
* ``識別子``：コマンド出力(バッククオート)

***

### 次のプログラムを実行するとどうなりますか

* `*hash`：`*`演算子(splat演算子)で、ハッシュを展開することができる

```ruby
>> hash = {a: 100, b: 200}
=> {:a=>100, :b=>200}
>> def splat_hash(a, b)
>>   p a
>>   p b
>> end
=> :splat_hash
>> splat_hash(*hash)
[:a, 100]
[:b, 200]
=> [:b, 200]
```

***

### 次のプログラムの期待値を得られるように正しいメソッドを選択肢から選んでください。

* `key`：引数を１つ取り、キーを返す

* `values`：ハッシュの全ての値の配列を返す

* `values_at`：指定された **キー** に対応する **値** を配列で返す。可変長なので、引数を何個も取れる

```ruby
>> hash = {price: 100, order_code: 200, order_date: "2018/09/20", tax: 0.8}
=> {:price=>100, :order_code=>200, :order_date=>"2018/09/20", :tax=>0.8}

# keyメソッド
>> p hash.key(:price, :tax)
ArgumentError: wrong number of arguments (given 2, expected 1)

# valuesメソッド
>> p hash.values(:price, :tax)
ArgumentError: wrong number of arguments (given 2, expected 0)

# values_atメソッド
>> p hash.values_at(:price, :tax)
[100, 0.8]
=> [100, 0.8]

# []メソッド
>> p hash.[](:price, :tax)
ArgumentError: wrong number of arguments (given 2, expected 1)
```

***

### 次のプログラムを実行するとどうなりますか

* `Enumerable#select`：ブロックの戻り値がtrueになる要素を配列にして返します。レシーバーをすべて走査して繰り返しを終了します。

配列の長さは10ですので、ブロックの戻り値がtrueかを問わず10が表示されます。

```ruby
>> $val = 0
=> 0
>>
?> class Count
>>   def self.up
>>     $val = $val + 1
>>     $val == 3 ? true : false
>>   end
>> end
=> :up
>>
?> [*1..10].select do   # *1で、配列として展開
?>   Count.up
>> end
=> [3]
>>
?> p $val
10
=> 10
```

***

### `KeyError`と`StopIteration`を捕捉するプログラムを選択肢から選んでください。

複数の例外クラスを捕捉するには代表的なものは3つです。

1. rescue節を捕捉したい例外クラスだけ記述する
2. rescue節で例外クラスをすべて記述する
3. 先の選択肢と同じですが、配列で指定してsplat演算子(`*`)で展開する
  →`*`は、配列の内容を並び順に展開する

```ruby
# 1.rescue節を捕捉したい例外クラスだけ記述する
begin
  # `KeyError`と`StopIteration`が発生する処理
rescue KeyError

rescue StopIteration

end

# 2.rescue節で例外クラスをすべて記述する
begin
  # `KeyError`と`StopIteration`が発生する処理
rescue KeyError, StopIteration

end

# 3.配列で指定してsplat演算子(`*`)で展開する
begin
  # `KeyError`と`StopIteration`が発生する処理
rescue *[KeyError, StopIteration]

end
```

***

### 次のプログラムを実行するとどうなりますか

`Enumerable#any?`はブロックの戻り値が`true`になると繰り返しをその時点で止めます。
繰り返しが止まるのは3回目の繰り返し、つまり`$val`が3になった時点です。

```ruby
# 問題
$val = 0

class Count
  def self.up
    $val = $val + 1
    $val == 3 ? true : false
  end
end

[*1..10].any? do
  Count.up
end

p $val

# 解答
>> $val = 0
=> 0
>>
?> class Count
>>   def self.up
>>     $val = $val + 1
>>     $val == 3 ? true : false
>>   end
>> end
=> :up
>>
?> [*1..10].any? do
?>   Count.up
>> end
=> true
>>
?> p $val
3
=> 3
```

***

### 次のコードを実行するとどうなりますか

* `"%d"`：10進数表現で数値を出力します。

```ruby
>> p "Hello%d" % 5
"Hello5"
=> "Hello5"
```

***

### 次のコードを実行するとどうなりますか

```ruby
# 問題
arr = [1,2].product([3,4]).transpose
p arr

# 解答
>> arr = [1,2].product([3,4]).transpose
=> [[1, 1, 2, 2], [3, 4, 3, 4]]
>> p arr
[[1, 1, 2, 2], [3, 4, 3, 4]]
=> [[1, 1, 2, 2], [3, 4, 3, 4]]
```

#### 解説

* `product`：レシーバーの配列と引数の配列からそれぞれ1つ要素を取り出し新しい配列を作成し、全ての配列を要素とする配列を返します。

```ruby
>> [1, 2].product([3, 4])
=> [[1, 3], [1, 4], [2, 3], [2, 4]]
```

* `transpose`：レシーバーの配列から行と列を入れ替えた配列を作成し返します。

```ruby
>> [[1, 3],
?>  [1, 4],
?>  [2, 3],
?>  [2, 4]
>> ].transpose
=> [[1, 1, 2, 2], [3, 4, 3, 4]]
```

***

### 以下の実行結果として正しいものを選択しなさい。

```ruby
# 問題
p [1, 2, 3].inject{|x, y| x + y ** 2} rescue p $!
p [1, 2, 3].inject(0){|x, y| x + y ** 2} rescue p $!
p [1, 2, 3].inject([]){|x, y| x << y ** 2} rescue p $!
p [1, 2, 3].inject do|x, y| x + y ** 2 end rescue p $!

# 解答
>> p [1, 2, 3].inject{|x, y| x + y ** 2} rescue p $!
14
=> 14
>> p [1, 2, 3].inject(0){|x, y| x + y ** 2} rescue p $!
14
=> 14
>> p [1, 2, 3].inject([]){|x, y| x << y ** 2} rescue p $!
[1, 4, 9]
=> [1, 4, 9]
>> p [1, 2, 3].inject do|x, y| x + y ** 2 end rescue p $!
#<LocalJumpError: no block given>
=> #<LocalJumpError: no block given>
```

#### 解説

* `Enumerable#inject`：ブロックを使用して繰り返し計算を行います。自身のたたみこみ演算を行う(初期値と自身の要素を順に組み合わせて結果を返す)

  * 引数を省略した場合は、要素1がブロック引数の1番目に渡されます。

  * 引数を指定した場合は、その値が初期値になります。

  * ブロック引数の1番目は前回の戻り値が渡されます。初回は、初期値が渡されます。

  * ブロック引数の2番目は要素が順番に渡されます

```ruby
# injectメソッドを使用しない場合
>> sum = 0
=> 0
>> (1..10).each {|i| sum += i }
=> 1..10
>> puts sum
55
=> nil

# injectメソッド使用時
>> (1..10).inject(0) {|sum, i| sum + i }
=> 55

# injectメソッド使用時(引数なし)
>> (1..10).inject {|sum, i| sum + i }
=> 55
```

#### 一行目

```ruby
>> p [1, 2, 3].inject{|x, y| x + y ** 2} rescue p $!
14
=> 14

# 解説
1 = 0 + 1 ** 2  # 0番目の1の二乗
5 = 1 + 2 ** 2  # 1(+1)番目の2の二乗
14 = 5 + 3 ** 2
```

#### 二行目

```ruby
>> p [1, 2, 3].inject(0){|x, y| x + y ** 2} rescue p $!
14
=> 14

# 解説
5 = 1 + 2 ** 2   # 一行目の2列目から始まる(初期値指定済みなので)
14 = 5 + 3 ** 2
```

#### 三行目

```ruby
>> p [1, 2, 3].inject([]){|x, y| x << y ** 2} rescue p $!
[1, 4, 9]
=> [1, 4, 9]

# 解説
[1] = [] << 1 ** 2
[1, 4] = [1] << 2 ** 2
[1, 4, 9] = [1, 4] << 3 ** 2
```

#### 四行目

`[1, 2, 3].inject`までが`p`メソッドの引数となるため、`p`メソッドへブロックが不正に渡されるため、エラーとなります。

```ruby
p([1, 2, 3].inject) do|x, y|
  x + y ** 2
end rescue p $!

# 解説
>> p([1, 2, 3].inject) do|x, y|
?>   x + y ** 2
>> end rescue p $!
#<LocalJumpError: no block given>
=> #<LocalJumpError: no block given>
```

***

### 次のコードを実行するとどうなりますか

`p1`と`p2`は別の`Proc`オブジェクトのため、`hoge`メソッド内の`current`変数は共有されません。
よって、`p2`の結果は6になります。

```ruby
>> def hoge(step = 1)
>>   current = 0
>>   Proc.new {
?>     current += step
>>   }
>> end
=> :hoge
>>
?> p1 = hoge
=> #<Proc:0x007ffd8f882f18@(irb):64>
>> p2 = hoge(2)
=> #<Proc:0x007ffd8f8829a0@(irb):64>
>>
?> p1.call
=> 1
>> p1.call
=> 2
>> p1.call
=> 3
>> p2.call
=> 2
>> p2.call
=> 4
>>
?> p p2.call
6
=> 6
```

***

### 次のコードを実行するとどうなりますか

`foo (2) * 2`はメソッド名と引数の間に空白があるため、`foo((2) * 2)`が呼ばれたと解釈されます。
よって、4の4乗の256になります。

```ruby
# 問題
def foo(n)
  n ** n
end

puts foo (2) * 2

# 解答
>> def foo(n)
>>   n ** n
>> end
=> :foo
>>
?> puts foo (2) * 2
256
=> nil
```

***

### 次のコードを実行するとどうなりますか

`foo(2) * 2`は`foo`メソッドによって`4`となるので、`4 * 2`=8となる

```ruby
def foo(n)
  n ** n
end

puts foo(2) * 2

# 解答
>> def foo(n)
>>   n ** n
>> end
=> :foo
>>
?> puts foo(2) * 2
8
```

### 次のコードを実行するとどうなりますか

今回の問題では文字列`"Hello"`にフォーマットに必要な指示子が無いためそのまま出力されます。

```ruby
# 問題
p "Hello" % 5

# 解答
?> p "Hello" % 5
"Hello"
=> "Hello"
```

***

### 次のコードを実行するとどうなりますか

`unless`は条件が成立しない場合に中の処理が実行されます。
`else`を用いることはできますが、`elsif`を用いることはできません。

```ruby
def hoge(n)
  unless n != 3
    "hello"
  elsif n == 5
    "world"
  end
end

str = ''
str.concat hoge(3)
str.concat hoge(5)

puts str

# エラーが発生する
```

***

### 実行してもエラーにならないコードを選べ

```ruby
# 1問目
(1..10).each
.reverse_each
.each do |i|
  puts i
end
# irbやpryではSyntaxErrorとなりますが、通常はSyntaxErrorとはなりません。

# 2問目
>> (1..10).each.
?> reverse_each.
?> each do |i|
?>   puts i
>> end
10
9
8
7
6
5
4
3
2
1
=> #<Enumerator: 1..10:each>

# 3問目
>> (1..10).each \  # 1, 2行目の行末でバックスラッシュ()を記述することで、1行のコードとみなされます。
?> .reverse_each \
?> .each do |i|
?>   puts i
>> end
10
9
8
7
6
5
4
3
2
1
=> #<Enumerator: 1..10:each>

# 4問目
>> (1..10).to_a.each.
?> reverse_each.
?> each do |i|
?>   puts i
>> end
10
9
8
7
6
5
4
3
2
1
=> #<Enumerator: [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]:each>
```

***

### 次のコードを実行するとどうなりますか

ヒアドキュメントにインデントを加える場合は、`<<-`識別子とします。
インデントを取り除くためには、Ruby 2.3から追加された`<<~`識別子を用います。

```ruby
>> s = <<-EOF
      Hello,
      Ruby
      EOF
=> "      Hello,\n      Ruby\n"
>>
?> p s
"      Hello,\n      Ruby\n"
=> "      Hello,\n      Ruby\n"

# 空白の対処法
>> s = <<~EOF
      Hello,
      Ruby
      EOF
=> "Hello,\nRuby\n"
```

***

### 次のプログラムの説明で正しい選択肢を選んでください。

ヒアドキュメントでは改行、スペースは削除されません。
`EOS`と`"EOS"`はヒアドキュメント内で式展開を行います。
識別子をシングルクォートで`'EOS'`囲う場合は式展開は行われず、全て文字列として扱われます。

```ruby
# 問題
str = <<EOS
よりニッチに。よりユニークに。
  IT市場はもちろん、ヘルスケア・医療・介護など
    次世代市場における企業や生活者のユーザビリティを向上させる
      サービス、ソフトウェアを開発しています。
    #{1 + 1}
EOS
puts str

# 解答
>> str = <<EOS
よりニッチに。よりユニークに。
  IT市場はもちろん、ヘルスケア・医療・介護など
    次世代市場における企業や生活者のユーザビリティを向上させる
      サービス、ソフトウェアを開発しています。
    #{1 + 1}
EOS
=> "よりニッチに。よりユニークに。\n  IT市場はもちろん、ヘルスケア・医療・介護など\n    次世代市場における企業や生活者のユーザビリティを向上させる\n      サービス、ソフトウェアを開発しています。\n    2\n"
>> puts str
よりニッチに。よりユニークに。
  IT市場はもちろん、ヘルスケア・医療・介護など
    次世代市場における企業や生活者のユーザビリティを向上させる
      サービス、ソフトウェアを開発しています。
    2
=> nil
```

***

### 次のプログラムを実行するとどうなりますか

`sub`は第一引数の検索対象のパターンと、第二引数の置換後の文字列を使って１回だけレシーバーの文字列を置換します。

問題では置換を５回していますので、`"I love @pple, b@n@n@ @nd grape"`が正解です

```ruby
# 問題
chars = "I love apple, banana and grape"

5.times do
  chars = chars.sub("a", "@")
end

p chars

# 解答
>> chars = "I love apple, banana and grape"
=> "I love apple, banana and grape"
>>
?> 5.times do
?>   chars = chars.sub("a", "@")
>> end
=> 5
>>
?> p chars
"I love @pple, b@n@n@ @nd grape"
=> "I love @pple, b@n@n@ @nd grape"
```

***

### 実行後の textfile.txt 内容になるようにXXXXに適切なコードを選べ

```ruby
open('textfile.txt', XXXX) do |f|
  data = f.read.upcase
  f.rewind
  f.puts data
end

# 実行前
recode 1
recode 2
recode 3

# 実行後
RECODE 1
RECODE 2
RECODE 3
```

#### 解説

* `w`：書き込みモードで開くため、`f.read`でエラーに

* `a+`：読み込みモード + 追記書き込みモード
  ファイルの読み込みは、ファイルの先頭から行いますが、書き込みは、ファイルの末尾に行います。
  `f.rewind`でファイルポインタをファイルの先頭に移動したとしても、ファイルの末尾に書き込まれます。

* `w+`：新規作成・読み込み + 書き込みモードで開きます。既にファイルが存在する場合は、空になります。

* `r+`：読み込み + 書き込みモードで開きます。

***

### 次のコードを実行するとどうなりますか。

```ruby
v1 = 1 - 1 == 0
v2 = v1 || raise RuntimeError
puts v2 && false

# 解答
>> v1 = 1 - 1 == 0
=> true
>> v2 = v1 || raise RuntimeError
SyntaxError: (irb):10: syntax error, unexpected tCONSTANT, expecting keyword_do or '{' or '('
v2 = v1 || raise RuntimeError
>> puts v2 && false
NameError: undefined local variable or method 'v2' for main:Object
```

#### 解説

2行目では、`||`の優先順位が高いため、`v1 || raise`が評価されます。
しかし、後ろに`RuntimeError`があるため、シンタックスエラーになっています。
この挙動を回避するには、以下の3つの方法があります。



```ruby
# 方法1
v2 = v1 or raise RuntimeError

# 方法2
v2 = v1 || raise(RuntimeError)

# 方法3
v2 = v1 || (raise RuntimeError)
```

***

### 次のコードを実行するとどうなりますか

`rescue`に処理対象の例外クラスの指定がない場合は、`StandardError`のサブクラス全てを捕捉します。
Rubyの組み込み例外は全て`StandardError`のサブクラスです。

```ruby
>> begin
?>   raise StandardError.new
>> rescue => e
>>   puts e.class
>> end
StandardError
=> nil
```

***

### 次のコードを実行するとどうなりますか

`do ... end`と`{ ... }`を比べた場合、`{ ... }`の方が結合度が強いです。

問題の式の場合、`do ... end`の結合度が弱いため、`p([1, 2, 3, 4].map)`が評価されます。
問題のように式の内容を直接使用する際は、`{ ... }`を使用します。

```ruby
p [1,2,3,4].map do |e| e * e end

# 解答
>> p [1,2,3,4].map do |e| e * e end
#<Enumerator: [1, 2, 3, 4]:map>
=> #<Enumerator: [1, 2, 3, 4]:map>

# 例1({}内で記述)
>> p [1,2,3,4].map{ |e| e * e}
[1, 4, 9, 16]
=> [1, 4, 9, 16]

# 例2(変数aに代入して表示)
>> a = [1, 2, 3, 4].map do |e|
?>   e ** e
>> end
=> [1, 4, 27, 256]
>> p a
[1, 4, 27, 256]
=> [1, 4, 27, 256]
```

***

### 次のコードを実行するとどうなりますか

```textfile
# list.txt
1
2
3
4
```

```ruby
io = File.open('list.txt')

while not io.eof?
  io.readlines
  io.seek(0, IO::SEEK_CUR)
  p io.readlines
end

# 解答
>> io = File.open('list.txt')
=> #<File:list.txt>
>>
?> while not io.eof?
>>   io.readlines
>>   io.seek(0, IO::SEEK_CUR)
>>   p io.readlines
>> end
[]
=> nil
```

#### 解説

* `IO#eof?`：ファイルポインタが終端にある場合、`true`を返します。

* `IO.readlines`：ファイルから全てを読み込みます。

* `IO#seek(offset, whence)`：ファイルポインタを`whence`から`offset`まで移動します。

  * `IO::SEEK_CUR`：現在のファイルポインタから

  * `IO::SEEK_SET`：ファイルの先頭からの位置を表す(デフォルト)

  * `IO::SEEK_END`：ファイルの末尾からを表す

4行目で、ファイルから全て読み込んだ時点で、ファイルポインタはファイル終端にあります。

5行目は、ファイル終端から0文字移動するため、

6行目では[]のみ表示されます。

***

### ソースコードの文字コードを`US-ASCII`に設定するマジックコメントの書き方として正しいものを全て選択してください

```ruby
# 1
# coding: us-ascii

# 2
# encoding: us-ascii

# 3
# -*- charset: us-ascii -*-

# 4
# CODING: US-ASCII
```

マジックコメントは、`coding: エンコーディング`が正しければ、その前後には任意の文字を並べることができる

また、大文字・小文字の区別はない。

***

### 以下のコードを実行するとどうなりますか

```ruby
x = 0
def hoge
  (1...5).each do |i|
    x += 1
  end
  puts x
end
hoge

# 解答
>> x = 0
=> 0
>> def hoge
>>   (1...5).each do |i|
?>     x += 1             # 変数xは、hogeメソッドの中で定義されていない。未定義の変数に代入しようとしているので、エラー。
>>   end
>>   puts x
>> end
=> :hoge
>> hoge
NoMethodError: undefined method '+' for nil:NilClass
```

***

### 以下のコードを実行するとどうなりますか。正しいものを全て選択してください。

```ruby
begin
  puts 1 + "2"
rescue
  puts "Error."
rescue TypeError
  puts "Type Error"
ensure
  puts "Ensure."
end

# 解答
>> begin
?>   puts 1 + "2"
>> rescue
>>   puts "Error."
>> rescue TypeError     # この部分は、上の`rescue`節で処理されている
>>   puts "Type Error"
>> ensure
?>   puts "Ensure."
>> end
Error.
Ensure.
=> nil
```

#### 解説

`rescue`節において例外型を省略すると、`StandardError`とそのサブクラス例外を捕捉する。

`TypeError`は`StandardError`のサブクラスなので、前の行の`rescue`節で処理される

```ruby
>> err = TypeError.new
=> #<TypeError: TypeError>
>> err.class
=> TypeError
>> err.class.superclass
=> StandardError
>> err.class.superclass.superclass
=> Exception
```

***

### 以下のコードを実行するとどうなりますか

```ruby
class Hoge
  attr_reader :message
  def initialize
    @message = "Hello"
  end
end

class Piyo < Hoge
  def initialize
    @message = "Hi"
    super
  end
end

puts Piyo.new.message

# 解答
>> class Hoge
>>   attr_reader :message
>>   def initialize
>>     @message = "Hello"
>>   end
>> end
=> :initialize
>>
?> class Piyo < Hoge
>>   def initialize
>>     @message = "Hi"
>>     super
>>   end
>> end
=> :initialize
>>
?> puts Piyo.new.message
Hello
=> nil
```

#### 解説

1. `Piyo`クラス：インスタンス化される(`initialize`メソッド)。`@message`に`"Hi"`の文字を格納した後に、`super`メソッドを呼び出す。

2. `Hoge`クラス：`Piyo`のスーパークラス。`initialize`メソッドで`@message`に`"Hello"`を代入している。

* `super`：メソッドが受け取った引数を、 **そのまま** スーパークラスの同名メソッドに渡して実行する。(`()`と引数を付けない場合)

***

### 以下のメソッドは円の面積を求めるコードですが、このままでは動きません。このコードを動かすための修正方法を選びなさい

```ruby
def area r
  return r * r * PI
end

# 解答
# １つ目
>> def area r
>>   return r * r * Math::PI
>> end
=> :area
>> area 3
=> 28.274333882308138

# ２つ目
>> def area r
>>   include Math
>>   return r * r * PI
>> end
=> :area
>> area 3
=> 28.274333882308138
```

#### 解説

定数`PI`は`Math`モジュールに定義されている。
このモジュールを使用するには、

* `Math`モジュールを、`include`する

* `Math::PI`と記述して、`PI`のモジュールを明示的に宣言する

***

### オーバーライド不可の演算子を選びなさい

以下の演算子以外は全てオーバーライド(再定義)可能

#### オーバーライド不可の演算子

* `::`：クラスやモジュールのネスト

* `&&`：論理積(優先度高)

* `||`：論理和(優先度高)

* `?:`：条件演算子

* `..`：`Range`オブジェクト(以下)

* `...`：`Range`オブジェクト(未満)

* `=`：代入

* `not`：否定(優先度低)

* `and`：論理積(優先度低)

* `or`：論理和(優先度低)

加えて、自己代入演算子


***

### 以下のコードを実行するとどうなりますか

```ruby
s = "Hello!"
def s.greet
  puts "Hi!"
end

class String
  def greet
    puts "Hello!"
  end
end

s.greet

# 解答
>> s = "Hello!"
=> "Hello!"
>> def s.greet
>>   puts "Hi!"
>> end
=> :greet
>>
?> class String
>>   def greet
>>     puts "Hello!"
>>   end
>> end
=> :greet
>>
?> s.greet
Hi!
=> nil
```

#### 解説

クラスを拡張して定義したメソッドよりも、特異メソッドが優先される。

変数`s`が参照する`String`オブジェクトは、`String`クラスを拡張して定義した`greet`よりも、

特異メソッドとして定義した`greet`を優先して実行する

***

### 以下の実行結果になるように`XXXX`に記述する適切なコードを選びなさい

```ruby
class Employee
  attr_reader :id
  attr_accessor :name
  def initialize id, name
    @id = id
    @name = name
  end
  def to_s
    return "#{@id}:#{name}"
  end
  XXXX
end

employees = []
employees << Employee.new("3","Tanaka")
employees << Employee.new("1","Suzuki")
employees << Employee.new("2","Sato")
employees.sort!
employees.each do |employee| puts employee end

# 実行結果
1:Suzuki
2:Sato
3:Tanaka
```

```ruby
# 解答
>> class Employee
>>   attr_reader :id
>>   attr_accessor :name
>>   def initialize id, name
>>     @id = id
>>     @name = name
>>   end
>>   def to_s
>>     return "#{@id}:#{name}"
>>   end
>>   def  <=> other
>>     return self.id <=> other.id
>>   end
>> end
=> :<=>
>>
?> employees = []
=> []
>> employees << Employee.new("3","Tanaka")
=> [#<Employee:0x007ff822989078 @id="3", @name="Tanaka">]
>> employees << Employee.new("1","Suzuki")
=> [#<Employee:0x007ff822989078 @id="3", @name="Tanaka">, #<Employee:0x007ff82291ca68 @id="1", @name="Suzuki">]
>> employees << Employee.new("2","Sato")
=> [#<Employee:0x007ff822989078 @id="3", @name="Tanaka">, #<Employee:0x007ff82291ca68 @id="1", @name="Suzuki">, #<Employee:0x007ff822961460 @id="2", @name="Sato">]
>> employees.sort!
=> [#<Employee:0x007ff82291ca68 @id="1", @name="Suzuki">, #<Employee:0x007ff822961460 @id="2", @name="Sato">, #<Employee:0x007ff822989078 @id="3", @name="Tanaka">]
>> employees.each do |employee| puts employee end
1:Suzuki
2:Sato
3:Tanaka
=> [#<Employee:0x007ff82291ca68 @id="1", @name="Suzuki">, #<Employee:0x007ff822961460 @id="2", @name="Sato">, #<Employee:0x007ff822989078 @id="3", @name="Tanaka">]
```

#### 解説

配列を`sort`するには、配列の要素クラスで演算子`<=>`を定義する

***

### 以下のコードを実行するとどうなりますか

```ruby
s = <<"EOB"
Hello,
Ruby
World.
EOB
"EOB"
p s

# 解答
>> s = <<"EOB"
Hello,
Ruby
World.
EOB
=> "Hello,\nRuby\nWorld.\n"
>> "EOB"                      # "EOB"は別に出力される。変数sとは関係ない
=> "EOB"
>> p s
"Hello,\nRuby\nWorld.\n"
=> "Hello,\nRuby\nWorld.\n"
```

***

### 以下のコードを実行するとどうなりますか

```
# data
abcdef
```

```ruby
File.open("data") do |io|
  while not io.eof?
    print io.read(1)
    io.seek(0, IO::SEEK_SET)
  end
end

# 解答
>> File.open("data") do |io|
?>   while not io.eof?
>>     print io.read(1)
>>     io.seek(0, IO::SEEK_SET)
>>   end
>> end
aaaaaaaaaaaaa
# 省略
```

#### 解説

`while not io.eof? ... end`は、ファイル`data`のファイルポイントが最後出ない時の条件を指定している

`io.seek(0, IO::SEEK_SET)`は、ファイルポインタを先頭に移すので、ファイル`data`の先頭文字`a`を読んで表示する処理を繰り返す

***

### 次のコードを実行するとどうなりますか

一行目の`x`と、二行目の`x`は異なるスコープなので、別の変数として扱われる

スコープを抜けた後の`x`は、`0`となる

```ruby
x = 0
[1,2,3].each do |x|
  print x.to_s + " "
end
puts x

# 解答
>> x = 0
=> 0
>> [1,2,3].each do |x|
?>   print x.to_s + " "
>> end
1 2 3 => [1, 2, 3]
>> puts x
0
=> nil
```

***

### 以下の実行結果になるように`XXXX`に記述する適切なコードを選びなさい

```ruby
y = false
y XXXX (raise "failed")
puts ("succeeded!")

# 出力結果
succeeded!
```

```ruby
# 解答
>> y = false
=> false
>> y && (raise "failed")   # `&&`メソッドは、左辺がfalseだと次の処理にいく
=> false
>> puts ("succeeded!")
succeeded!
=> nil

# 間違いの例
>> y = false
=> false
>> y || (raise "failed")   # `||`メソッドは、左辺がfalseだと右辺を探しに行く
RuntimeError: failed
>> puts ("succeeded!")
succeeded!
=> nil
```

***

### 以下のコードでは何が出力されますか

```ruby
>> p ?A
"A"
=> "A"
```

***

### 以下のコードを実行するとどうなりますか

```ruby
class Parent
  attr_reader :name
  def initialize name
    @name = name
  end
end

class Child < Parent
  def initialize name
    @name = "Child" + name
  end
end

puts Child.new("Hoge").name
```

```ruby
# 結果
>> class Parent
>>   attr_reader :name
>>   def initialize name
>>     @name = name
>>   end
>> end
=> :initialize
>>
?> class Child < Parent
>>   def initialize name
>>     @name = "Child" + name
>>   end
>> end
=> :initialize
>>
?> puts Child.new("Hoge").name
ChildHoge
=> nil
```

#### 解説

`Child`クラスのオブジェクトが生成されると、初期化時に実行される`initialize`メソッドは`Child`メソッドとなる

従って、インスタンス変数`name`には`ChildHoge`が出力される

***

### 以下のコードの実行結果として適切なものを選びなさい

```ruby
a = [1,2,3]
b = [1,3,5]
c = [2,3,4]

p a + b - c

# 解答
>> a = [1,2,3]
=> [1, 2, 3]
>> b = [1,3,5]
=> [1, 3, 5]
>> c = [2,3,4]
=> [2, 3, 4]
>>
?> p a + b - c
[1, 1, 5]
=> [1, 1, 5]
```

#### 解説

配列に演算子`+`を適用すると、両オペランドの配列を連結した配列を生成して返す

```ruby
>> p a + b
[1, 2, 3, 1, 3, 5]
=> [1, 2, 3, 1, 3, 5]
```

演算子`-`は、左オペランドの配列を削除した配列を返す

```ruby
>> p a - b
[2]
=> [2]
```

***

### 以下のコードの実行結果として適切なものを選びなさい

```ruby
p "find!find!find!find!find!find".index("!", 5)

# 解答
>> p "find!find!find!find!find!find".index("!", 5)
9
=> 9
```

#### 解説

`String#index`は、部分文字列を探索し見つかった位置を返す。

第2引数には、探索の開始位置を指定する。

出題コードの探索の開始位置は、5(6番目)。

***

### １つ以上の数字のみで構成される行にマッチする正規表現を選びなさい

```ruby
# 1問目
/^[0-9].$/       # 1文字ではマッチしない

# 2問目
/^[0-9]*$/       # 1文字ではマッチしない

# 3問目
/^[0-9][0-9]*$/  # 正解

# 4問目
/^[0-9][0-9].*$/ # 1文字ではマッチしない
```

#### 解説

「１つ以上の数字のみで構成される行」とは、「数字で始まり、数字の繰り返しで終わる行」という意味になる
→この条件にマッチするのは、選択肢3

1問目：数字で始まる2文字の行

2問目：数字の0文字以上の繰り返しの行(空行でもマッチ)

4問目：数字2文字で始まる行。3文字目以降が数字以外でもマッチ

***

### 以下のコードを実行すると何が表示されますか

```ruby
p "abc def 123 ghi 456".scan(/\d+/).length

# 解答
>> p "abc def 123 ghi 456".scan(/\d+/).length
2
=> 2
```

#### 解説

`String#scan`は、正規表現にマッチする部分文字列を配列で返す

正規表現`\d+`は、数字の繰り返しとなる

出題コードの文字列で数字のみの繰り返しパターンは、`123`、`456`の2つ

```ruby
>> p "abc def 123 ghi 456".scan(/\d+/)
["123", "456"]
=> ["123", "456"]
```

***

### 以下のコードを実行すると何が表示されますか

```ruby
p "HogeHOGEhoge"[/[A-Z][^A-Z]+/]

# 解答
>> p "HogeHOGEhoge"[/[A-Z][^A-Z]+/]
"Hoge"
=> "Hoge"
```

#### 解説

正規表現`/[A-Z][^A-Z]+/`は、最初が英大文字、2文字目以降が英大文字以外の文字列にマッチする

***

### 以下のコードを実行すると何が表示されますか

```ruby
h = {1 => "Hoge", 2 => "Piyo", 3 => "fuga"}
h.reject {|x, y| x < 2}
p h

# 解答
>> h = {1 => "Hoge", 2 => "Piyo", 3 => "fuga"}
=> {1=>"Hoge", 2=>"Piyo", 3=>"fuga"}
>> h.reject {|x, y| x < 2}             # 非破壊メソッド
=> {2=>"Piyo", 3=>"fuga"}
>> p h
{1=>"Hoge", 2=>"Piyo", 3=>"fuga"}
=> {1=>"Hoge", 2=>"Piyo", 3=>"fuga"}
```

#### 解説

`Hash#reject`は、ブロックの実行結果が`true`になる要素を削除したハッシュを返す

ただし、非破壊メソッドなので元のハッシュは変更しない

***

### 以下の実行結果になるように、`XXXX`に記述する適切なコードを選びなさい

```ruby
puts File.XXXX ("/", "user", "bin")

# 解答
>> puts File.join("/", "user", "bin")
/user/bin
=> nil
```

#### 解答

`File`クラスの`join`は、`File::SEPARATOR`(ファイルパスのセパレータを保持する定数。値は環境によらず`/`)

を間に入れて引数の文字列を連結したファイルパスを作成する

ただし、前の文字列が`File::SEPARATOR`で終わる場合は、`File::SEPARATOR`を間に入れない

***

### 期待した出力結果になるようにXXXXに適切なコードを選べ

```ruby
d = Time.new(2015, 1, 5)
puts d.strftime(XXXX)

# 出力結果
01/05/15

# 解答
>> d = Time.new(2015, 1, 5)
=> 2015-01-05 00:00:00 +0900
>> puts d.strftime("%x")
01/05/15
=> nil
```

* `%x`：日付(%m/%d/%y)

* `%m`：月を表す数字(01-12)

* `%M`：分(00-59)

* `%d`：日(01-31)

* `%D`：日付(%m/%d/%y)

* `%y`：西暦の下2桁(00-99)

* `%Y`：西暦を表す数(9999)


```ruby
>> d = Time.new(2015, 1, 5)
=> 2015-01-05 00:00:00 +0900

# "x"
>> puts d.strftime("%x")
01/05/15
=> nil

# "%m/%d/%Y"
>> puts d.strftime("%m/%d/%Y")
01/05/2015
=> nil

# "%m/%D/%y"
>> puts d.strftime("%m/%D/%y")
01/01/05/15/15
=> nil

# "%M/%d/%y"
>> puts d.strftime("%M/%d/%y")
00/05/15
=> nil
```

***

### 学習マラソンで間違えた問題

```ruby
>> p (1..10).lazy.map{|num|
?>   num * 2
>> }.take(3).inject(0, &:+)
12
=> 12

# 解説
p (1..10).lazy.map{|num| num * 2 }.take(3).inject(0, &:+) # 2+4+6=12

# 例1
>> p (1..10).map{|num| num * 2 }.inject(0, &:+)
110
=> 110

# 例2
>> p (1..10).map{|num| num * 2 }
[2, 4, 6, 8, 10, 12, 14, 16, 18, 20]
=> [2, 4, 6, 8, 10, 12, 14, 16, 18, 20]

# 例3
>> p (1..10).lazy.map{|num| num * 2 }.take(3)
#<Enumerator::Lazy: #<Enumerator::Lazy: #<Enumerator::Lazy: 1..10>:map>:take(3)>
=> #<Enumerator::Lazy: #<Enumerator::Lazy: #<Enumerator::Lazy: 1..10>:map>:take(3)>
```

```ruby
while not DATA.eof?
  lines = DATA.readlines
  lines.map(&:chomp!)
  lines.reverse
  p lines
end

__END__
1
2
3
4

# 解答
["1", "2", "3", "4"]
```

* `DATA`：`__END__`以降をアクセスする`File`オブジェクト

```ruby
class C
  CONST = "Good, night"
end

module M
  CONST = "Good, evening"
end

module M
  class C
    CONST = "Hello, world"
  end
end

module M
  class C
    p CONST
  end
end

# 解答
>> class C
>>   CONST = "Good, night"
>> end
=> "Good, night"
>>
?> module M
>>   CONST = "Good, evening"
>> end
=> "Good, evening"
>>
?> module M
>>   class C
>>     CONST = "Hello, world"
>>   end
>> end
=> "Hello, world"
>>
?> module M
>>   class C
>>     p CONST
>>   end
>> end
"Hello, world"
=> "Hello, world"
```

```ruby
class C
  @val = 3
  attr_accessor :val
  class << self
    @val = 10
  end
  def initialize
    @val *= 2 if val
  end
end

c = C.new
c.val += 10

p c.val
```

### Rubyで使用可能なオプションではないものを選択しなさい(複数)。

***

```ruby

module M
  def class_m
    "class_m"
  end
end

class C
  include M
end

p C.methods.include? :class_m


# 解答
?> module M
>>   def class_m
>>     "class_m"
>>   end
>> end
=> :class_m
>>
?> class C
>>   include M
>> end
=> C
>>
?> p C.methods.include? :class_m
false
=> false
```

***


### 次のプログラムは"Hello, world"と表示します。同じ結果になる選択肢はどれですか（複数選択

```ruby
module M
  CONST = "Hello, world"

  class C
    def awesome_method
      CONST
    end
  end
end

p M::C.new.awesome_method

# 1問目
module M
  CONST = "Hello, world"
end

class M::C
  def awesome_method
    CONST
  end
end

p M::C.new.awesome_method

# 2問目

```

***

### 実行後の textfile.txt 内容になるようにXXXXに適切なコードを選べ。ただし、空ファイルは作成済みである。

```ruby
File.open('testfile.txt', XXXX) do |f|
  f.write("recode 1\n")
  f.seek(0, IO::SEEK_SET)
  f.write("recode 2\n")
end

# 実行後の textfile.txt 内容
recode 1
recode 2
```

```ruby
File.open('testfile.txt', "w+") do |f|
  f.write("recode 1\n")
  f.seek(0, IO::SEEK_SET)
  f.write("recode 2\n")
end
=> recode 2

File.open('testfile.txt', "a+") do |f|
  f.write("recode 1\n")
  f.seek(0, IO::SEEK_SET)
  f.write("recode 2\n")
end
=> recode 1
=> recode 2
```

* `w+`
`w+`は新規作成・読み込み + 書き込みモードで開きます。
既にファイルが存在する場合は、空になります。
`IO#seek`はファイルポインタを指定の位置に移動します。`IO:SEEK_SET`がファイルの先頭からの位置を指定する識別子です。

よって、`recode 1`を書き込み後にファイルの先頭にファイルポインタを移動し、`recode 2`で上書きしています。

※`w`も同様

* `a+`
`a+`はファイルを読み込みモード + 追記書き込みモードで開きます。
ファイルの読み込みは、ファイルの先頭から行いますが、書き込みは、ファイルの末尾に行います。

***

### 以下のコードを実行した時の出力として正しいものを1つ選択してください。

Timeクラスのオブジェクトに対して「+」メソッドや「-」メソッドを使って数値を加算減算できます。数値の単位は秒です。

```ruby
>> t = Time.now + (60*60*24)
=> 2018-09-29 00:07:02 +0900

# 実行時の日時から24時間後(86400秒後)の日時が表示される
>> p t
2018-09-29 00:07:02 +0900
=> 2018-09-29 00:07:02 +0900
```

***

### ローカル変数の名前として正しいものをすべて選択してください。

１文字目は`アルファベット小文字`か`_`で始める必要があります。

また、２文字目以降はアルファベットもしくは数字を使用します。

予約語はローカル変数として使用できません

***

### 以下のコードを実行した出力として正しいものを１つ選択してください。

`Integer#downto(min)`は、引数`min`まで数を１ずつ減らしながら実行されます。

`select`：要素に対してブロックの評価が真であった要素をすべて含む配列が返されます。

```ruby
p 100.downto(90).select{|x| x%2==0}

# 解答
>> p 100.downto(90).select{|x| x%2==0}
[100, 98, 96, 94, 92, 90]
=> [100, 98, 96, 94, 92, 90]
```

***

### 以下のコードにおける4行目`io.rewind`について正しい説明を１つ選択してください。

* `rewind`：ファイルポインタを先頭に移動し、`lineno`の値を`0`にする

```ruby
File.open("foo.txt","r") do |io|
puts io.gets
 puts io.read
  io.rewind
  p lines = io.readlines
end

# 解答
>> File.open("foo.txt","r") do |io|
?> puts io.gets
>>  puts io.read
>>   io.rewind
>>   p lines = io.readlines
>> end
[]
=> []
```

***

### 以下のコードを実行したときの出力として適切な物を1つ選択してください。

* `File#join`：定数`FILE::SEPARATOR "/"`を使って文字列を連結します。

```ruby
p File.join("ruby", "exam","silver")

# 解答
>> p File.join("ruby", "exam","silver")
"ruby/exam/silver"
=> "ruby/exam/silver"
```

***

### エラーの種類

* `NameError`：未定義のローカル変数や定数を参照した場合

* `IndexError`：添字が範囲外のときに発生します。

```ruby
>> s = "foo"
=> "foo"
>> begin
?>   s[4] = ?b
>> rescue IndexError
>>   puts "error"
>> end
error
=> nil
```

***

### 以下のコードを実行したときの出力として適切な物を1つ選択してください。

論理演算子の評価方法について、

* `&&`演算子：左辺が真と評価されたときのみ右辺も評価されます。

* `||`演算子：右辺が評価されるのは、左辺が偽と評価された場合です。

```ruby
ary = []
ary << 1 && false
true || ary << 2
false && ary << 3
false || ary << 4
p ary

# 解答
>> ary = []
=> []
>> ary << 1 && false
=> false
>> true || ary << 2  # 左辺がtrueなので、右辺は処理しない
=> true
>> false && ary << 3 # 左辺がfalseなので、処理しない
=> false
>> false || ary << 4
=> [1, 4]
>> p ary
[1, 4]
=> [1, 4]
```

***

### 以下のコードを実行したときの出力として適切な物を1つ選択してください。

`String#**`は定義されていません。

演算子の優先順位は`*`よりも`**`が高いため,

高い `2**2回(４回)"foo"`が繰り返される新しい文字列を返します。

```ruby
p "foo" * 2 **2

# 解答
>> p "foo" * 2 **2
"foofoofoofoo"
=> "foofoofoofoo"
```

***

### テキストファイルを読み込んだファイルオブジェクトから一行ずつ読み込み表示したい。目的に一致するIOクラスのメソッドを２つ選択してください。

`IO#gets`と`IO#readline`はファイルオブジェクトから一行読み込んで、読み込みに成功した時にはその文字列を返します。

`IO#gets`と`IO#readline`の違いはEOFに到達した時の振る舞いのみです。

* `IO#gets`は`nil`

* `IO#readline`は`EOFError`を返します。

***

# ここからまだ復習できていない！！

### 以下のコードを実行した時にIOErrorが発生した。理由として考えられるものはどれか１つ選択してください。

* 読み込みモードでファイルが開かれているため

モードを明示的に指定しない場合、読み込みモード`"r"`でファイルを開きます。

そのため、エラーが発生しています。選択肢Bの場合は、`IOError`ではなく`Errno::ENOENT`（システムコールに依存したエラー）を発生させます。

例題では、読み込みモードを`"w"`として開いた場合のファイルの内容は、`Time.now`で得た現在時刻を指定した`format`文字列で出力したものになります。
(例：2015/06/07)

```ruby
File.open("foo.txt") do |io|
  io.write(Time.now.strftime("%Y/%m/%d"))
end
```

***

### 次のメソッドでDirクラスのクラスメソッドではないものをすべて選択してください。(2つ選択)

* `Dir.basename`

* `Dir.extname`

***

### 以下のコードを実行した時の正しい出力結果を1つ選択してください。


inject (Enumerable)は前回のブロックの戻り値をブロックに渡して繰り返し実行します。

例題では条件演算子を使い、２つの整数値を比較しより大きな値を取り出しそれをブロックに渡しています。

条件演算子(文法） 式1 ? 式1が真だった場合の値 : 式1が偽だった場合の値

```ruby
numbers = [3,89,40,39,29,10,50,59,69]
num = numbers.inject do |i,j|
  i > j ? i : j
end
p num

# 解答
>> numbers = [3,89,40,39,29,10,50,59,69]
=> [3, 89, 40, 39, 29, 10, 50, 59, 69]
>> num = numbers.inject do |i,j|
?>   i > j ? i : j
>> end
=> 89
>> p num
89
=> 89
```

***

### 以下の出力になる時の ___(1)___ に入るものとして適切なものを1つ選択してください。

`uri`はRubyに標準で添付されているライブラリーの1つです。

`require`することによって、プログラム内で呼び出すことができます。

`include`、`extend`はモジュールを読み込む際に使用します。

`import`はRuby以外の言語で同様の目的の際に利用されることがありますが、Rubyでは使えません。

```ruby
___(1)___ 'uri'
uri = URI::HTTP.build({host:'www.ruby.or.jp', path:'/ja/certification/examination/'})
puts uri

# 解答
>> require 'uri'
=> true
>> uri = URI::HTTP.build({host:'www.ruby.or.jp', path:'/ja/certification/examination/'})
=> #<URI::HTTP http://www.ruby.or.jp/ja/certification/examination/>
>> puts uri
http://www.ruby.or.jp/ja/certification/examination/
=> nil
```

***

### 以下のコードを実行した時の正しい出力結果を1つ選択してください。

Enumerable#detectメソッドは要素に対してブロックを評価した値が真になった最初の要素を返します。

例題では1~20までの数字に対して一番最初の5の倍数を求めていますので正解はCです。

Enumerable#findを使っても同様の結果になります。

```ruby
numbers = (1..20).to_a
p numbers.detect{|x| x % 5 == 0}

# 解答
>> numbers = (1..20).to_a
=> [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20]
>> p numbers.detect{|x| x % 5 == 0}
5
=> 5
```

***

### 以下のコードを実行した時の出力結果として正しいものを１つ選択してください。

raise関数によって明示的に例外を発生させることができます。

例題ではString#ascii_only?を使いテキストにASCII文字以外が使われている場合には例外を発生させています。

例外処理の基本形は問題28の解説を参照してください。

```ruby
class NonasciiError < StandardError
end

File.open("sample.txt") do |io|
  io.each_line do |str|
    begin
      raise(NonasciiError, "non ascii character detected") unless str.ascii_only?
    rescue => ex
      puts "#{ex.message} : #{str}"
    end
  end
end

[sample.txtの内容]
Ruby Association
ルビーアソシエーション
るびー
Ruby on Rails

# 解答
>> class NonasciiError < StandardError
>> end
=> nil
>>
?> File.open("sample.txt") do |io|
?>   io.each_line do |str|
?>     begin
?>       raise(NonasciiError, "non ascii character detected") unless str.ascii_only?
>>     rescue => ex
>>       puts "#{ex.message} : #{str}"
>>     end
>>   end
>> end
non ascii character detected : ルビーアソシエーション
non ascii character detected : るびー
=> #<File:sample.txt (closed)>
```

***

### 10進数で10を表すものを2つ選択してください。

0xAは16進数、012は8進数です。いずれも10進数では10となります。

puts "#{0xA}" => 10

puts '#{012}' => #{012}   ダブルクォーテーション内では式展開されますが、シングルクォーテーションでは式展開が行われません。

```ruby
A. 0xA
B. 0xFF
C. 012
D. 077
E. 0x10
```

***
