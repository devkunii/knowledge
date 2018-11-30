合格教本基礎力問題gold
===================



## Rubyの`-l`オプションの説明として正しいものを全て選びなさい

1. 引数で指定したディレクトリを`$LOAD_PATH`変数に追加する

1. 引数で指定したファイルを読み込む

1. 引数で指定したディレクトリを環境変数`LUBYLIB`に追加する

1. 引数で指定したディレクトリは`require`や`load`メソッドが呼ばれた時に検索される



### 解説

`-l`は、ファイルをロードするパスを指定するオプション

指定したディレクトリは`$LOAD_PATH`変数(`$:`)に追加される

環境変数の`RUBYLIB`に指定したパスも同様に`$LOAD_PATH`変数にパスを追加するが、`-l`オプションで指定したパスが

直接`LUBYLIB`に追加されることはない

2の引数で指定したファイルを読み込むのは、`-r`オプション



## 以下のコードを実行すると何が表示されますか

> 正解したけど、念のため
>
> 2018/11/07

```ruby
>> x, *y = *[0, 1, 2]
=> [0, 1, 2]
>> p x, y
0
[1, 2]
=> [0, [1, 2]]
```



### 解説

多重代入の問題

右辺の「`*`」は無視して良い

左辺の`x`には、配列の最初の要素が格納される

`y`には「`*`」がついているので、残りの要素が配列として代入される



## `__X__`に記述すると、以下の実行結果にならないコードを全て選びなさい

> 正解したけど、念のため
>
> 2018/11/07

```ruby
puts __X__

#
# 実行結果
#
0.8
```

1. `4 / 5`

1. `4.0 / 5`

1. `4/5r`

1. `4 / 5.0`



### 解説

整数(Integer)同士の演算は整数となり、小数点以下は丸められる

選択肢1の結果は`0`

選択肢3の`4/5r`はRationalクラスのオブジェクトを生成し、`(4/5r).to_f`とすれば`0.8`となるが、

そのままでは`(4/5)`と出力される



## 以下のコードを実行するとどうなりますか

> 正解したけど、念のため
>
> 2018/11/07

```ruby
>> class Err1 < StandardError; end
=> nil
>> class Err2 < Err1; end
=> nil

>> begin
>>   raise Err2
>> rescue => e
>>   puts "StandardError"
>> rescue Err2 => ex
>>   puts ex.class
>> end
StandardError
=> nil
```



### 解説

begin節の`raise`で発生する例外オブジェクトのクラスはErr2

2個目のrescue節ではErr2、もしくはErr2を継承した例外オブジェクトを対象とするように書かれているが、

1個目のrescue節でStandardErrorから派生する全ての例外を対象とするため、

2個目のrescue節では処理されず、1個目のrescue節で処理される



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



## 以下の実行結果になるように、`__X__`に記述する適切なコードを選びなさい

```ruby
def hoge(*args)
  p __X__
end
hoge [1, 2, 3]

# 実行結果
[1, 2, 3]
```

1. `args`

1. `*args`

1. `&args`

1. `$args`



### 解説

「`*`」の付いたメソッド引数は、可長変引数となる

`hoge`の引数`args`は配列として扱われる

`hoge`に配列`[1, 2, 3]`を渡すと、`args[0]`に渡される

したがって、実行結果を`[1, 2, 3]`にするには、

* `args[0]`

* `*args`

にする

```ruby
>> def hoge(*args)
>>   p *args
>> end
=> :hoge
>> hoge [1, 2, 3]
[1, 2, 3]
=> [1, 2, 3]
```



## 以下の実行結果になるように、`__X__`に記述する適切なコードを選びなさい

> 正解しているが、念のため
>
> 2018/11/07

```ruby
def hoge __X__
  puts "#{x}, #{y}, #{params[:z]}"
end

hoge x: 1, z: 3

# 実行結果
1, 2, 3
```

1. `(x:, y: 2, params: *)`

1. `(x:, y: 2, *params)`

1. `(x:, y: 2, **params)`

1. `(x:, y: 2, params: **)`



### 解説

キーワード引数の中で任意の値を受け取れるようにしたい場合は、`**params`のように引数の前に`**`を繋げることで、

明示的に定義したキーワード以外の引数をHashオブジェクトで受け取ることができる

```ruby
>> def hoge (x:, y: 2, **params)
>>   puts "#{x}, #{y}, #{params[:z]}"
>> end
=> :hoge

>> hoge x: 1, z: 3
1, 2, 3
=> nil
```



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



## 以下の実行結果になるように、`__X__`と`__Y__`に記述する適切なコードを選びなさい

```ruby
a, b = __X__ do
  for x in 1..10
    for y in 1..10
      __Y__ if x + y == 10
    end
  end
end

puts a, b

# 実行結果
1
9
```

* 選択肢1

  ```ruby
  __X__: loop
  __Y__: break
  ```

* 選択肢2

  ```ruby
  __X__: catch
  __Y__: throw [x, y]
  ```

* 選択肢3

  ```ruby
  __X__: catch
  __Y__: break [x, y]
  ```

* 選択肢4

  ```ruby
  __X__: catch :exit
  __Y__: throw :exit, [x, y]
  ```



### 解説

`throw`メソッドを呼ぶと、第1引数で指定した`tag`のあるcatchブロックの終わりまでジャンプします。

この時、`throw`メソッドの第2引数に渡した値が`catch`メソッドの戻り値になる

```ruby
>> a, b = catch :exit do
>>   for x in 1..10
>>     for y in 1..10
>>       throw :exit, [x, y]if x + y == 10
>>     end
>>   end
>> end
=> [1, 9]

>> puts a, b
1
9
=> nil
```



## 以下のコードを実行するとどうなりますか

```ruby
>> module Mod
>>   def foo
>>     puts "Mod"
>>   end
>> end
=> :foo
>> class Cls1
>>   include Mod
>>   def foo
>>     puts "Cls1"
>>     super
>>   end
>> end
=> :foo
>> class Cls2 < Cls1
>>   def foo
>>     puts "Cls2"
>>     super
>>   end
>> end
=> :foo

>> Cls2.new.foo
Cls2
Cls1
Mod
=> nil
```



### 解説

メソッドは、

* 自分のクラス->includeしているモジュール->スーパークラス->スーパークラスのincludeしているモジュール

の順に検索される

Cls2クラスのオブジェクトで`foo`を呼び出すと、Cl2クラスの`foo`が実行され、Cls2を画面に出力したあと、

`super`でCls1クラスの`foo`を呼び出します。

Cls1クラスの`foo`の中でCls1を画面に出力した後、`super`でModモジュールの`foo`を呼び出し、Modを画面に出力します

したがって、Cls2、Cls1、Modの順番で出力されます



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
>>   p self.ancestors
>>   include M2
>> end
[Cls2, Cls1, M1, Object, Kernel, BasicObject]
=> Cls2
```



### 解説

`ancestors`は、クラス、モジュールの優先順で配列に格納して返す

Cls2クラスの`include M2`は`ancestors`実行後のため、`ancestors`の対象外です



## 以下の実行結果になるように、`__X__`に記述する適切なコードを全て選びなさい(複数選択)

```ruby
class Example
  def hoge
    self.piyo
  end
  __X__
  def piyo
    puts "piyo"
  end
end

Example.new.hoge

# 実行結果
piyo
```

1. `private`

1. `protected`

1. `public`

1. 何も記述しない



### 解説

`self.piyo`と`piyo`メソッドに`self`のレシーバをつけているので、`piyo`メソッドが`private`メソッドだとエラーになる

それ以外の`protected`、`public`、何も記述しない場合は`self.piyo`で呼び出すことができる



## Kernelモジュールのメソッドについて、正しい記述を選びなさい(2つ選択)

> 正解しているけど、念のため
>
> 2018/11/09

1. `load`は同じファイルを複数回ロードしない

1. `require`は拡張子なしでファイルをロードできる

1. `freeze`を実行すると、そのクラスはインスタンス化できなくなる

1. `dup`は特異メソッドなどの情報をコピーしない



### 解説

* `load`は、指定したファイルを複数回ロードする

* `freeze`は、指定したオブジェクトを変更不可にする



## 以下のコードについて適切な記述はどれですか

```ruby
%r|(http://www(\.)(.*)/)| =~ "http://www.abc.com/"
```

1. `$0`の値は、nilである

1. `$1`の値は、`http://www`である

1. `$2`の値は、`.abc`である

1. `$3`の値は、`abc.com`である



### 解説

`$0`はスクリプトのファイル名

`$1`、`$2`...はそれぞれグループ化された正規表現にマッチした文字列になる

```ruby
>> %r|(http://www(\.)(.*)/)| =~ "http://www.abc.com/"
=> 0

#
# 選択肢1
#
>> $0
=> "irb"

#
# 選択肢2
#
>> $1
=> "http://www.abc.com/"

#
# 選択肢3
#
>> $2
=> "."

#
# 選択肢4
#
>> $3
=> "abc.com"
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



## 以下のコードを実行したらどうなりますか。テストされるファイル`foo.rb`は存在していると仮定します。

```ruby
require 'test/unit'
require 'foo'
class TC_Foo < Test::Unit::TestCase
  def foo_test
    # テストコード
  end
end
```

1. `foo_test`が規約に沿っていないので実行されない

1. `test/unit`の記述が誤りのため、例外が発生する

1. `Test::Unit::TestCase`の記述が誤りのため、例外が発生する

1. テストが実行される



### 解説

`Test::Unit`のテストメソッド名は、`test_`で始める



## `socket`ライブラリにあるクラスを全て選びなさい

1. BasicServer

1. UDPSocket

1. UDPServer

1. BasicSocket

1. TCPSocket



### 解説

`socket`ライブラリに、

* BasicServerクラス

* UDPServerクラス

は存在しない



## 以下の実行結果になるように、`__X__`に記述する適切なコードを選びなさい

```ruby
require "date"

date = Date.new(2000, 10, 10)
puts date __X__ 1

# 実行結果
2000-11-10
```

1. `+`

1. `-`

1. `<<`

1. `>>`



### 解説

* `date + 1`：`date`の1日後の日付を返す

* `date - 1`：1日前の日付を返す

* `date << 1`：一ヶ月前の日付を返す

* `date >> 1`：一ヶ月後の日付を返す



## 以下の記述で適切なものを全て選びなさい

1. StringIOは、IOと同じインターフェースを持つ

1. FloatオブジェクトとRationalオブジェクトの演算結果は、Floatオブジェクトとなる

1. rdocにおいて、`*`は番号付きリストを作成する

1. Threadクラスの`start`、`fork`、`run`は、いずれも新しいスレッドを実行する

1. DateTimeオブジェクトに`1`加算すると、翌日のデータとなる



### 解説

* rdocの「`*`」は、番号なしリストを作る

* Threadクラスには`run`は存在しない



## 以下のコードはHTTPデダウンロードしたデータを表示するコードです。`__X__`に記述する適切なコードを選びなさい

```ruby
require 'open-uri'

__X__ ("http://docs.ruby-lang.org/ja/2.1.0/doc/index.html") do |f|
  print f.read
end
```

1. `open`

1. `OpenURI.open`

1. `open_uri`

1. `OpenURI.get`



### 解説

`open_uri`はKernelモジュールの`open`を再定義し、HTTPで指定したファイルを取得する



| 版 | 年/月/日  |
|---|----------|
|初版|2018/11/07|
