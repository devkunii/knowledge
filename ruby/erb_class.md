# 第5章 組み込みクラス

## Silverの範囲

* String
* Array
* Hash
* IO
* Dir
* File
* Time
* Regexp

5章では、5-18〜5-20はやらない

***

## 5-1.`BasicObject`クラス

* `Object`クラスのスーパークラス

* `Object`クラスでは定義されているメソッドが多すぎる場合など、
  特殊な用途のために用意されているクラス

***

## 5-2.`Kernel`モジュール

* 全てのクラスから参照できるメソッドを定義しているモジュール

* [参照](https://docs.ruby-lang.org/ja/2.1.0/class/Kernel.html)

***

## 5-3.`Object`クラス

* 全てのクラスのスーパークラス

* `Kernel`モジュールをインクルードしているため、全てのオブジェクトで`Kernel`モジュールのメソッドを使用可能

***

### 5-3-1.オブジェクトのID

* `object_id`：オブジェクトに割り当てられる、重複しない整数を取得する

```ruby
>> a = "foo"
=> "foo"
>> a.object_id
=> 70175072383100
>> a.__id__
=> 70175072383100
```

* Rubyでは **全てがオブジェクト** なので、同じリテラルでも **基本的にオブジェクトIDは異なる**
  →しかし、`TrueClass`、`FalseClass`、`NilClass`、`Symbol`、`Fixnum`クラスのインスタンスは同じオブジェクトid

```ruby
>> a.object_id
=> 70175072383100
>> "foo".object_id    # 異なるオブジェクトid
=> 70175072330980
>> :foo.object_id
=> 1159388
>> :foo.object_id     # 同じオブジェクトid(Hash)
=> 1159388
```

***

### 5-3-2.オブジェクトのクラス

* `class`：オブジェクトのクラスを調べるメソッド

```ruby
>> "foo".class
=> String
>> :foo.class
=> Symbol
```

***

### 5-3-3.オブジェクトの比較

* `equal?`：オブジェクトID(object_id)を比較する

* `eql?`：オブジェクトのハッシュ値が同じかどうかを比較

```ruby
>> a = "foo"
=> "foo"
>> a.hash
=> -3839024608656029219
>> a.object_id
=> 70175068384400
>> b = "foo"
=> "foo"
>> b.hash
=> -3839024608656029219     # hashの値は`a`と同じ
>> b.object_id
=> 70175068354380           # object_idの値は`a`と異なる
>> a.eql?(b)
=> true
>> a.equal?(b)
=> false
```

* `===`：`case`式で利用される。オブジェクトを比較する。新しいクラスを作った際に定義すると良い

* `==`：オブジェクトの内容が同じかどうかを比較

* 比較メソッドはクラスによって再定義されるため、結果はクラスに依存することに注意する
  →例)`String`クラスでは、`==`メソッドは`eql?`メソッドと同じ結果を返す

```ruby
>> a = "foo"
=> "foo"
>> b = "foo"
=> "foo"
>> a.eql?(b)
=> true
>> a == b
=> true
```

***

### 5-3-4.オブジェクトのメソッド一覧

#### オブジェクトに定義されているメソッドを取得するためのメソッド

* `methods`：全ての呼び出し可能なメソッド

* `private_methods`：プライベートメソッド

* `protected_methods`：プロテクテッドメソッド

* `public_methods`：パブリックメソッド

* `singleton_methods`：特異メソッド

* これらのメソッドは、一覧を配列で返す。配列の要素は、メソッド名のシンボル

```ruby
>> a = "foo"
=> "foo"
>> a.methods
=> [:include?, :%, :*, :+, :to_c, :unicode_normalize, :unicode_normalize!, :unicode_normalized?, :count, :partition, :unpack, :unpack1, :sum, :next, :casecmp, :casecmp?, :insert, :bytesize, :match, :match?, :succ!, :+@, :-@, :index, :rindex, :<=>, :replace, :clear, :upto, :getbyte, :==, :===, :setbyte, :=~, :scrub, :[], :[]=, :chr, :scrub!, :dump, :byteslice, :upcase, :next!, :empty?, :eql?, :downcase, :capitalize, :swapcase, :upcase!, :downcase!, :capitalize!, :swapcase!, :hex, :oct, :split, :lines, :reverse, :chars, :codepoints, :prepend, :bytes, :concat, :<<, :freeze, :inspect, :intern, :end_with?, :crypt, :ljust, :reverse!, :chop, :scan, :gsub, :ord, :start_with?, :length, :size, :rstrip, :succ, :center, :sub, :chomp!, :sub!, :chomp, :rjust, :lstrip!, :gsub!, :chop!, :strip, :to_str, :to_sym, :rstrip!, :tr, :tr_s, :delete, :to_s, :to_i, :tr_s!, :delete!, :squeeze!, :each_line, :squeeze, :strip!, :each_codepoint, :lstrip, :slice!, :rpartition, :each_byte, :each_char, :to_f, :slice, :ascii_only?, :encoding, :force_encoding, :b, :valid_encoding?, :tr!, :encode, :encode!, :hash, :to_r, :<, :>, :<=, :>=, :between?, :clamp, :instance_of?, :kind_of?, :is_a?, :tap, :public_send, :public_method, :singleton_method, :remove_instance_variable, :define_singleton_method, :method, :instance_variable_set, :extend, :to_enum, :enum_for, :!~, :respond_to?, :object_id, :send, :display, :nil?, :class, :singleton_class, :clone, :dup, :itself, :taint, :tainted?, :untaint, :untrust, :untrusted?, :trust, :frozen?, :methods, :singleton_methods, :protected_methods, :private_methods, :public_methods, :instance_variable_get, :instance_variables, :instance_variable_defined?, :!, :!=, :__send__, :equal?, :instance_eval, :instance_exec, :__id__]
```

***

### 5-3-5.オブジェクトの複製

* `clone`：`dup`に加えて、凍結状態(freeze)、特異メソッドも複製

* `dup`：汚染状態(taint)、インスタンス変数、ファイナライザを複製

* ここでの複製は、シャローコピー(浅いコピー)であり、自分自身の複製しかできない
  →例)配列の要素の参照先は、複製できない

```ruby
>> a = "hoge"
=> "hoge"
>> a.object_id
=> 70175072632680
>> b = a.dup
=> "hoge"
>> b.object_id
=> 70175072614940
```

***

### 5-3-6.オブジェクトの凍結

* `freeze`：オブジェクトの内容の変更を禁止する。

* 凍結されたオブジェクトを変更しようとすると、`RuntimeError`が発生

```ruby
>> a = "abc"
=> "abc"
>> a.freeze
=> "abc"
>> a[0] = 'z'
RuntimeError: can｀t modify frozen String
```

***

### 5-3-7.オブジェクトの汚染

* オブジェクトの *汚染マーク* がセットされている場合、`tained?`メソッドが`true`を返す

* *汚染マーク* をつけるには`taint`メソッドを、外すには`untaint`メソッドを使用

* オブジェクトの汚染については[参照](https://docs.ruby-lang.org/ja/2.1.0/doc/spec=2fsafelevel.html)

```ruby
>> a = "string"
=> "string"
>> a.tainted?
=> false
>> a.taint      # 汚染マークをつける
=> "string"
>> a.tainted?
=> true
>> a.untaint    # 汚染マークを外す
=> "string"
>> a.tainted?
=> false
```

***

### 5-3-8.インスタンス変数にアクセスする

* `instance_variable_get`：インスタンス変数の取得

* `instance_variable_set`：インスタンス変数の設定

* `instance_variables`：インスタンス変数の一覧

* インスタンス変数は、`:@hoge`(シンボル)か、`"@hoge"`(文字列)で指定

```ruby
>> class Foo
>>   def initialize
>>     @hoge = 1
>>   end
>> end
=> :initialize
>> f = Foo.new
=> #<Foo:0x007fa5d186d908 @hoge=1>
>> f.instance_variables                # インスタンス変数の一覧を返す
=> [:@hoge]
>> f.instance_variable_get(:@hoge)     # インスタンス変数の取得
=> 1
>> f.instance_variable_set(:@hoge,2)   # インスタンス変数の設定
=> 2
>> f.instance_variable_get(:@hoge)     # インスタンス変数を設定後の取得
=> 2
```

***

### 5-3-9.未定義メソッドの呼び出し

* オブジェクトに定義されていないメソッドが呼び出されたとき、Rubyは`method_missing`メソッドを呼びだす

* `method_missing`メソッドが定義されていない場合は、`NoMethodError`になる

```ruby
>> class Bar
>>   def method_missing(name, *args)
>>     puts name
>>   end
>> end
=> :method_missing
>> b = Bar.new
=> #<Bar:0x007fe42880e210>
>> b.hoge                  # method_missingが呼び出され、`hoge`が出力
hoge
=> nil
>> "string".hoge           # method_missingが未定義のため、例外
NoMethodError: undefined method ｀hoge｀ for "string":String
```

***

### 5-3-10.オブジェクトの文字列表現

* `to_s`メソッド：オブジェクトの内容や値の文字列表現を返す
  例)オブジェクトのクラス名を表示

* `inspect`メソッド：オブジェクトを人間が読める形式に変換。主にデバッグ
  例)インスタンス変数とその値まで表示

```ruby
>> a = 1.2
=> 1.2
>> a.to_s            # 浮動小数点数1.2を文字列に変換
=> "1.2"
>> class Hoge
>>   def initialize
>>     @foo = "bar"
>>   end
>> end
=> :initialize
>> hoge = Hoge.new
=> #<Hoge:0x007fe428987ba0 @foo="bar">
>> hoge.to_s         # インスタンスhogeを、文字列に変換
=> "#<Hoge:0x007fe428987ba0>"
>> hoge.inspect      # インスタンスhogeを、文字列に変換(inspect)
=> "#<Hoge:0x007fe428987ba0 @foo=\"bar\">"
```

***

### 5-3-11.特別なオブジェクト

* `TrueClass`：`true`

* `FalseClass`：`false`

* `NilClass`：`nil`

* これらのクラスのオブジェクトは、唯一のインスタンス
  →変更できない

***

## 5-4.数値型クラス

* `Numeric`クラス：数値を表すクラス。他のクラスはこのクラスを継承
  * `Integer`クラス：整数を表すクラス
  * `Float`クラス：浮動小数点数を表すクラス
  * `Complex`クラス：複素数を扱うクラス
  * `Rational`クラス：有理数を扱うクラス

***

### 5-4-1.`Numeric`クラス

* 小数点の切り捨て、切り上げには以下のメソッドを使用
  * `ceil`：それ自身と同じかそれ自身より大きな整数のうち、最小のものを返す
  * `floor`：それ自身より小さな整数のうち、最大のものを返す
  * `round`：もっとも近い整数を返す
  * `truncate`：それ自身と`0`の間にある整数で、もっとも近いものを返す

```ruby
>> 1.9.ceil       # 1.9より大きな整数・・・2
=> 2
>> 1.9.floor      # 1.9より小さな整数・・・1
=> 1
>> 1.9.round      # 一番近い整数・・・2
=> 2
>> 1.9.truncate   # 1.9と0の間で一番近い整数・・・1
=> 1
>> -1.1.ceil      # -1.1より大きな整数・・・-1
=> -1
>> -1.1.floor     # -1.1より小さな整数・・・-2
=> -2
>> -1.1.round     # -1.1に一番近い整数・・・-1
=> -1
>> -1.1.truncate  # -1.1と0の間で一番近い整数・・・-1
=> -1
```

***

#### 数値の絶対値

* `abs`メソッド：数値の絶対値を返す

```ruby
>> -3.abs
=> 3
```

***

#### 数値を使った繰り返し

* `step`メソッド：1から100までの繰り返しなどを返す

* `downto`メソッド：30から15までの繰り返しなどを返す

```ruby
>> 1.step(10,2){|n| puts n}   # 1から10までの奇数を表示
1
3
5
7
9
=> 1
```

***

### 5-4-2.`Float`クラス

* 浮動小数点数として保存される
  →計算過程で誤差が発生することがある

* 小数で処理できない処理(シフト演算、ビット演算など)は実装されていない

```ruby
>> 1.23 & 2
NoMethodError: undefined method ｀&｀ for 1.23:Float
```

***

### 5-4-3.`Integer`クラス

* `**`メソッド：整数のべき乗を求める

* `/`メソッド：整数の除算を求める。ただし、整数同士の除算では、整数が結果として返る

```ruby
# べき乗
>> 2.**(4)
=> 16
>> 2 ** 4
=> 16

# 除算
>> 10 / 3
=> 3
>> 10 / 3.0
=> 3.3333333333333335
```

***

#### 整数に対応する文字を求める

* `chr`メソッド：アスキーコードに対応する文字を求める
  →対応する文字がない場合は、`RangeError`になる

```ruby
>> 65.chr
=> "A"
>> -1.chr
RangeError: -1 out of char range
```

***

#### 次の整数、前の整数を求める

* `next`メソッド：(整数に対して)次の整数を返す。`succ`メソッドでもある

* `pred`メソッド：自身の数から-1した整数を返す

```ruby
>> 10.next
=> 11
>> 10.succ
=> 11
>> 10.pred
=> 9
```

***

#### 整数を使った繰り返し

* `times`メソッド：その整数の数だけ、与えられたブロックを実行

* `upto`メソッド：ある整数からある整数まで、与えられたブロックを実行

* `downto`メソッド：ある整数からある整数まで、与えられたブロックを実行

```ruby
>> sum = 0
=> 0
>> 10.times {|i| sum += i}
=> 10
>> puts sum                    # なぜ45になるのかわからない
45
=> nil
```

***

### 5-4-4.`Fixnum`クラスと`Bignum`クラス

* `Fixnum`クラス：123のような数値リテラルのクラス

* `Bignum`クラス：数値が大きい数値リテラルのクラス

* どちらも基本的に同じなので、`Fixnum`クラスをここでは扱う

***

#### 数値の剰余

* `%`メソッド：剰余に使われる

```ruby
>> 10.modulo(3)
=> 1
>> 10 % 3
=> 1
```

***

#### ビット演算

* `|`：論理和
* `&`：論理積
* `^`：排他的論理和
* `~`：否定
* `<<`：左シフト演算
* `>>`：右シフト演算

例)10(0b1010)と3(0b0011)のビット演算

```ruby
>> 10 | 3
=> 11            # 0b1011
>> 10 & 3
=> 2             # 0b0010
>> 10 ^ 3
=> 9             # 0b1001
>> ~ 10
=> -11           # 0b..10101
>> 10 << 1
=> 20            # 0b10110
>> 10 >> 1
=> 5             # 0b0101
```

※わからないので、もう一度学習
2018/08/25

***

#### 数値の小数化

* `to_f`メソッド：`Float`クラスに変換するメソッド

```ruby
>> 10.to_f
=> 10.0
```

***

#### 小数を使った繰り返し

* `step`メソッド：小数で繰り返しを行うメソッド

```ruby
>> 1.5.step(21.5, 2.5){|f| puts f}
1.5
4.0
6.5
9.0
11.5
14.0
16.5
19.0
21.5
=> 1.5
```

***

### 5-4-5.`Rational`クラス

有理数を扱う`Numeric`クラスのサブクラス

* `Rational`クラスによる有理数のインスタンスを作成するには、
  `Rational(a,b=1)`のように表す(`a`：分子、`b`：分母)

* `denominator`メソッド：有理数の分母を求める

* `numerator`メソッド：有理数の分子を求める

* `Numeric`クラスでも、これらのメソッドは使用可能
  →返り値は、自身を`Rational`に変換した場合に対応する値

```ruby
>> a = Rational(1,2)   # Rationalクラスによるインスタンスの作成(1/2)
=> (1/2)
>> a.denominator       # 1/2の、分母(2)
=> 2
>> a.numerator         # 1/2の、分子(1)
=> 1
>> 0.25.denominator    # 0.25は1/4。これの分母(4)
=> 4
>> 0.25.numerator      # 0.25は1/4。これの分母(1)
=> 1
```

***

#### `divmod`メソッド

* 引数`other`で除算した結果と、その剰余を配列で返す

* `abs`メソッド：絶対値を返す

```ruby
>> Rational(1,2).divmod Rational(1,3)
=> [1, (1/6)]               # 1 * (1/3) + (1/6) = 1/2
>> Rational(-4,13).abs
=> (4/13)
```

***

#### 整数を丸める結果を返すメソッド

* `floor`メソッド：小さい方の整数に丸める

* `ceil`メソッド：大きい方の整数に丸める

* `truncate`メソッド：0に近い方の整数に丸める

* `round`メソッド：絶対値を四捨五入する

```ruby
# 定義
>> a = Rational(1,3)
=> (1/3)
>> b = Rational(2,3)
=> (2/3)
>> c = Rational(-1,3)
=> (-1/3)
>> d = Rational(-2,3)
=> (-2/3)

# floorメソッド
>> a.floor
=> 0
>> b.floor
=> 0
>> c.floor
=> -1
>> d.floor
=> -1

# ceilメソッド
>> a.ceil
=> 1
>> b.ceil
=> 1
>> c.ceil
=> 0
>> d.ceil
=> 0

# truncateメソッド
>> a.truncate
=> 0
>> b.truncate
=> 0
>> c.truncate
=> 0
>> d.truncate
=> 0

# roundメソッド
>> a.round
=> 0
>> b.round
=> 1
>> c.round
=> 0
>> d.round
=> -1
```

***

#### `Rational`クラスと他の`Numeric`クラスとの四則演算

* `Rational`クラスは、四則演算の中で演算対象のクラスを判別して、返す結果のクラスを決定

* 演算の対象が`Rational`または`Integer`の場合・・・`Rational`が演算結果

* 演算の対象が`Float`の場合・・・`Float`が演算結果

```ruby
>> Rational(1,2) + Rational(1,3)
=> (5/6)
>> Rational(1,2) + 1
=> (3/2)
>> Rational(1,2) + 0.25
=> 0.75
```

***

### 5-4-6.`Complex`クラス

* 複素数を扱う`Numeric`クラスのサブクラス

* 複素数のインスタンスを作成するには、
  `Complex(a,b)`と記述(`a`：実数、`b`：虚数)

* 複素数リテラルでは、虚部の数値の値の後に`i`をつけて表すことができる

* 複素数の実部と虚部は、

  * `real`メソッド：     実部

  * `imaginary`メソッド：虚部

* `Numeric`クラスでもこれらのメソッドを使用できるが、

  * 実部：`self`

  * 虚部：`0`

  を返す

```ruby
>> a = (1+2i)
=> (1+2i)
>> a.real
=> 1
>> a.imaginary
=> 2
>> 3.real
=> 3
>> 3.imaginary
=> 0
```

***

#### 複素数の極座標表現

* (x+iy)=r(cosθ+sinθ)

* `abs`メソッド：`r`(絶対値)を返す

* `arg`メソッド：`θ`(偏角)を返す

* `polar`メソッド：絶対値と偏角を、配列で返す

```ruby
>> Complex(1,3).abs     # 絶対値
=> 3.1622776601683795
>> Complex(1,3).arg     # 偏角
=> 1.2490457723982544
>> Complex(1,3).polar   # 絶対値&偏角
=> [3.1622776601683795, 1.2490457723982544]
```

***

#### `Complex`クラスと他の`Numeric`クラスの四則演算

* `Numeric`クラスのサブクラスも`real`メソッド、`imaginary`メソッドが使用可能
  →`Bignum`、`Fixnum`、`Float`などのインスタンスを跨いでも四則演算が可能

* 四則演算の結果は、`Complex`クラスのインスタンスで返る

```ruby
>> Complex(1,1) + 0.5
=> (1.5+1i)
```

***

## 5-5.`Encoding`クラス

### 主なエンコーディング

* `Encoding::UTF_8`：`UTF-8`を表すエンコーディング

* `Encoding::EUC_JP`：`EUC_JP`を表すエンコーディング

* `Encoding::ISO_2022_JP`：`JIS`を表すエンコーディングで、Rubyではダミーエンコーディング

* `Encoding::Shift_JIS`：`Shift_JIS`を表すエンコーディング

* `Encoding::Windows_31J`：Windowsで用いられる`Shift_JIS`の亜種(`CP932`とも言う)
  →`Encoding::CP932`でも参照可能

* `Encoding::ASCII`：`US-ASCII`を表すエンコーディング
  →`Encoding::US_ASCII`でも参照可能

* `Encoding::ASCII_8BIT`：`ASCII`互換のエンコーディングで、文字コードを持たないデータや、文字列を単なるバイト列とσ知恵扱いたい場合に利用

***

#### 規定の外部エンコーディング

* エンコーディングが指定されていないときは、規定の外部エンコーディングは各システムに依存

* Linuxであれば、localeに`UTF-8`が指定されている

* `default_external`メソッド：規定の外部エンコーディングを取得

```ruby
>> Encoding.default_external
=> #<Encoding:UTF-8>
```

***

#### エンコーディングの互換性

* `compatible?`メソッド：異なるエンコーディングの間の互換性を調べる

* 互換性がある場合には、エンコーディングを、ない場合には`nil`を返す

```ruby
>> Encoding.compatible?(Encoding::UTF_8, Encoding::US_ASCII)
=> #<Encoding:UTF-8>
>> Encoding.compatible?(Encoding::UTF_8, Encoding::Shift_JIS)
=> nil
```

* 互換性のあるエンコーディングでは文字列を結合できるが、互換性のない場合はエラーになり結合できない

```ruby
>> a = "ルビー"
=> "ルビー"
>> b = a.encode("EUC-JP")
=> "\x{A5EB}\x{A5D3}\x{A1BC}"
>> a + b
Encoding::CompatibilityError: incompatible character encodings: UTF-8 and EUC-JP
```

* ただし、互換性のないエンコーディングでもどちらか一方の文字列がASCII文字しか含まない場合は結合可能

```ruby
>> a = "abc"
=> "abc"
>> b = "あいう".encode("EUC-JP")
=> "\x{A4A2}\x{A4A4}\x{A4A6}"
>> b.encoding
=> #<Encoding:EUC-JP>
>> (a + b)
=> "abc\x{A4A2}\x{A4A4}\x{A4A6}"
>> (a + b).encoding
=> #<Encoding:EUC-JP>
```

***

## `String`クラス

### 5-6-1.文字列の文字コード情報

* `UTF-8`：主に利用されている文字コード

* `EUC-JP`：古いUNIX系システムで利用されていた文字コード

* `JIS`：JISで策定された7bit文字のみを使った文字コード

* `Shift_JIS`：マルチバイト文字とASCII文字を切り替えることなく利用できるようにした文字コード

* `Windows-31J`：主にWindowsで利用されいるShift_JISの亜種。機種依存文字などが利用できる文字コード

* `US-ASCII`：ASCII文字だけで構成されている文字コード

***

#### 文字列のエンコーディングの取得・変更

* `encoding`メソッド：`String`オブジェクトの文字コード情報(エンコーディング)を取得

* `encode`：引数で指定した文字コードに変換した新しいインスタンスを返す

* `encode!`：オブジェクトのエンコーディングを変更

* 文字列の操作は、このエンコーディング情報を元に行われるため、エンコーディングの異なる文字列を結合したり比較する際には注意が必要

```ruby
# エンコーディングの取得
>> a = "abc"
=> "abc"
>> a.encoding
=> #<Encoding:UTF-8>

# エンコーディングの変更(encode)
>> b = a.encode("EUC-JP")
=> "abc"
>> b.encoding
=> #<Encoding:EUC-JP>

# エンコーディングの変更(encode!)
>> a = "ルビー"
=> "ルビー"
>> a.encode!("EUC-JP")
=> "\x{A5EB}\x{A5D3}\x{A1BC}"
>> a.encoding
=> #<Encoding:EUC-JP>
```

***

### 5-6-2.文字列の比較

* `==`：大文字小文字、全角半角を含めて同じかどうかを比較

* `>`、`>=`、`<`、`<=`：文字列同士をアスキーコードで比較(`true`or`false`)

* `<=>`、`casecmp`：比較した結果を-1、0、1の整数値で返す(UFO演算子と同じ働き)
  →`<=>`は大文字小文字を区別するが、`casecmp`は区別しない

```ruby
>> "abc" == "abc"
=> true
>> "abc" == "ABC"  # 大文字小文字は区別
=> false
>> "a" < "b"       # 辞書順では、aの方が小さい
=> true
>> "A" > "a"       # なぜかわからないので、あとで調べる
=> false
>> "aa" < "b"      # 辞書順では、aaの方が小さい
=> true
>> "a" <=> "b"     # aの方が小さい(-1)
=> -1
```

***

#### 文字列と数値の比較

* `==`メソッド：異なるクラスのオブジェクトと比較できる。(型の自動変換は行われないので、`false`)

```ruby
>> '100' == 100
=> false
>> '100' >= 100
ArgumentError: comparison of String with 100 failed
```

***

#### 文字列比較のときのエンコーディング

* `==`、`eql?`メソッド：両者のエンコーディングが等しく、文字列自身のバイト列表現が等しい場合のみ`true`を返す

* ただし、両者のエンコーディングがASCII互換でASCII文字しか含まない場合は、エンコーディングが異なる場合もバイト表現が一致すれば`true`を返す

```ruby
# ASCII互換
>> a = "abc"
=> "abc"
>> b = a.encode("EUC-JP")
=> "abc"
>> b.encoding
=> #<Encoding:EUC-JP>
>> a == b
=> true

>> a = "ルビー"
=> "ルビー"
>> b = a.encode("EUC-JP")
=> "\x{A5EB}\x{A5D3}\x{A1BC}"
>> a == b
=> false
```

***

### 5-6-3.文字列の切り出し

* `[]`、`slice`：文字列から指定された一部分を切り出すメソッド

* `slice!`：文字列から指定された一部分を切り出し、返した文字を元の文字列から取り除く

* `split`：文字列や正規表現を使って文字列を分割

```ruby
# []
>> a = 'abcdef'
=> "abcdef"
>> puts a[2]
c
=> nil

# slice
>> a.slice(2)
=> "c"

# slice!
>> a.slice!(2)
=> "c"
>> puts a
abdef
=> nil

# split
>> 'abcdefg'.split('d')
=> ["abc", "efg"]
>> 'abcdef'.split(/d/)
=> ["abc", "ef"]
>> "abcde\nfghij".split(/\n/)  # ダブルクオートに注意
=> ["abcde", "fghij"]
```

***

#### 数値を指定した場合

* 数値を指定した場合は、数値の位置にある文字を返す

* 負の数値の場合は、末尾から数えてた位置の文字を返す

```ruby
>> 'abcdefg'[2]       # 0,1,2・・・"c"
=> "c"
>> 'abcdefg'.slice(2) # 0,1,2・・・"c"
=> "c"
>> 'abcdefg'[-2]      # 6,5・・・"f"
=> "f"
>> a = 'abcdefg'
=> "abcdefg"
>> a.slice!(2)        # 0,1,2・・・"c"でカット
=> "c"
>> puts a
abdefg
=> nil
```

***

#### 範囲指定の場合

* 範囲として`Range`オブジェクトを指定した場合は、該当する範囲の文字列を返す

* 範囲指定は、開始位置と長さで指定可能。

  * 開始位置が範囲外の場合：`nil`

  * 開始位置が負の場合：末尾から数えた位置となる

  * 長さが文字列より長い場合には、可能な部分までを返す

```ruby
# Rangeオブジェクトで指定
>> puts a
abdefg
=> nil
>> 'abcdefg'[1..3]
=> "bcd"
>> 'abcdefg'.slice(1..3)
=> "bcd"
>> a = 'abcdefg'
=> "abcdefg"
>> a.slice!(1..3)
=> "bcd"
>> puts a
aefg
=> nil

# 開始位置と長さで指定
>> 'abcdefg'[1,3]
=> "bcd"
>> 'abcdefg'.slice(1,3)
=> "bcd"
>> 'abcdefg'[-2,3]
=> "fg"
>> a = 'abcdefg'
=> "abcdefg"
>> a.slice!(1,3)
=> "bcd"
>> puts a
aefg
=> nil
```

***

#### 文字列で指定

* 元の文字列に含まれていればその部分を、含まれていなければ`nil`を返す

```ruby
>> 'abcdefg'["bc"]
=> "bc"
>> 'abcdefg'.slice("bc")
=> "bc"
>> 'abcdefg'["bd"]  # "bd"で表される文字は含まれない
=> nil
>> a = 'abcdefg'
=> "abcdefg"
>> a.slice!("bc")
=> "bc"
>> puts a
adefg
=> nil
```

***

#### 正規表現で指定

* マッチした部分があればその部分を、マッチしなければ`nil`を返す

```ruby
>> 'abcdefg'[/bc/]
=> "bc"
>> 'abcdefg'.slice(/bc/)
=> "bc"
>> a = 'abcdefg'
=> "abcdefg"
>> a.slice!(/bc/)
=> "bc"
>> puts a
adefg
=> nil
```

***

### 5-6-4.文字列の変更

* `[]=`、`insert`メソッド：文字列の一部分を変更する。範囲や位置の指定ができ、該当する部分を新しい文字列で置換することができる

```ruby
>> a = 'abcdefg'
=> "abcdefg"
>> a[1..3] = 'xyz'
=> "xyz"
>> puts a
axyzefg
=> nil
```

***

### 5-6-5.文字列の置換

#### `sub`メソッド・`gsub`メソッド

* `sub`メソッド：指定したパターンにマッチした最初の部分を、特定の文字列に置換する

* `gsub`メソッド：マッチした全ての部分を置換する

```ruby
>> a = 'abcdefg-abcdefg'
=> "abcdefg-abcdefg"
>> a.sub(/abc/, 'xyz')
=> "xyzdefg-abcdefg"
>> a.sub(/abc/, 'xyz')
=> "xyzdefg-abcdefg"
```

* ブロックも取ることができ、その場合にはブロックにマッチした部分が渡され、ブロックの実行結果と置換される

```ruby
>> a = 'abcdefg-abcdefg'
=> "abcdefg-abcdefg"
>> a.sub(/abc/) {|str| 'xyz'}
=> "xyzdefg-abcdefg"
>> a.gsub(/abc/) {|str| 'xyz'}
=> "xyzdefg-xyzdefg"
```

***

#### `tr`メソッド・`tr_s`メソッド

* `tr`メソッド：指定したパターンに含まれる文字を検索し、それを特定の文字列やパターンに合わせて置換する

* `tr_s`メソッド：`tr`メソッドの機能に加えて、重複する文字を1文字に圧縮

```ruby
>> a = 'aabbccddeeffgg'
=> "aabbccddeeffgg"
>> a.tr('a-c', 'A-C')
=> "AABBCCddeeffgg"
>> a.tr_s('a-c', 'A-C')
=> "ABCddeeffgg"
```

***

#### `delete`メソッド

* 指定したパターンに含まれる文字を、元の文字列から削除

* パターンを複数指定すると、全てのパターンに含まれる文字列のみ削除する

```ruby
>> a = 'aabbccddeeffgg'
=> "aabbccddeeffgg"
>> a.delete('a-f', 'd-g')
=> "aabbccgg"
```

***

#### `squeeze`メソッド

* 指定した文字が複数並んでいた場合に、一文字に圧縮

```ruby
>> a = 'aabbccddeeffgg'
=> "aabbccddeeffgg"
>> a.squeeze('a-e')
=> "abcdeffgg"
```

***

#### `replace`メソッド

* 引数の文字列で自分自身の内容を置き換える

```ruby
>> a = 'abc'
=> "abc"
>> a.replace('xyz')
=> "xyz"
>> puts a
xyz
=> nil
```

***

### 5-6-6.文字列の連結

* `+`：文字列を結合した新しいオブジェクトを生成

* `<<`、`concat`：元のオブジェクトの内容に文字列を追記

* `*`：文字列の内容を指定した数値の数だけ、繰り返した文字列を返す

* 異なるエンコーディングの文字列を結合するとき、互換性がない場合はエラー

* 互換性がなくても、文字列がASCII文字列のみからなるASCII互換であれば結合可能

```ruby
>> a = 'abc'
=> "abc"
>> a.object_id
=> 70263014985560   # 元々のオブジェクト
>> a << 'def'
=> "abcdef"
>> a.object_id
=> 70263014985560   # 同じオブジェクト
>> a = a + 'ghi'
=> "abcdefghi"
>> a.object_id
=> 70263018562880   # 新しいオブジェクト
>> 'abc' * 2
=> "abcabc"
```

***

### 5-6-7.文字列の大文字・小文字への変換

* `capitalize`：文字列の先頭にある半角英字を大文字に、残りの半角英字を小文字にして返す

* `downcase`、`upcase`：それぞれ半角英字を小文字化、大文字化して返す

* `swapcase`：半角英字の小文字を大文字に、大文字を小文字に変更して返す

```ruby
>> a = 'abcDEF'
=> "abcDEF"
>> a.capitalize
=> "Abcdef"
>> a.upcase
=> "ABCDEF"
>> a.downcase
=> "abcdef"
>> a.swapcase
=> "ABCdef"
```

***

### 5-6-8.文字列の末尾や先頭にある空白や改行を削除

* `chomp`：末尾から引数で指定する改行コードを取り除いた文字列を返す。
  →指定がない場合は、`"\r"`、`"\r\n"`、`"\n"`の全てを改行コードと見なして取り除く

* `strip`、`lstrip`、`rstrip`：先頭と末尾、先頭、末尾にある空白文字を取り除いた文字列を返す
  →`"\t"`、`"\r"`、`"\n"`、`"\f"`、`"\v"`など

* `chop`：末尾の文字を取り除いた文字列を返す

```ruby
>> a = "abcdef\n"
=> "abcdef\n"
>> a.chomp         # 改行コードを取り除く
=> "abcdef"
>> a.strip         # 改行コードを取り除く
=> "abcdef"
>> a.lstrip
=> "abcdef\n"
>> a.rstrip
=> "abcdef"
>> a.chop          # 末尾の文字を取り除く
=> "abcdef"
>> a.chop.chop
=> "abcde"
>> a.chomp.chomp   # 改行コードのみ取り除く
=> "abcdef"
```

***

### 5-6-9.文字列を逆順にする

* `reverse`：並び順をバイト単位で逆にする

```ruby
>> a = "abcdef"
=> "abcdef"
>> a.reverse
=> "fedcba"
```

***

### 5-6-10.文字列の長さ

* `length`、`size`：文字数を返す

* `count`：指定されたパターンに該当する文字がいくつあるかを返す

* `empty?`：文字列が空かどうかを返す

* `bytesize`：バイト数が返る

```ruby
# 文字列の長さ
>> a = "abcdef"
=> "abcdef"
>> a.length
=> 6
>> a.count('a-c')
=> 3
>> a.empty?
=> false
>> "".empty?
=> true

# 文字列のバイト数
>> a = "ルビー"
=> "ルビー"
>> a.length
=> 3
>> a.bytesize
=> 9
```

***

### 5-6-11.文字列の割り付け

文字列をある長さの文字列の中に割り付ける

メソッドを呼び出すときに、

  * 返す文字列の長さ

  * 割り付ける余白に使用する文字

を指定できる

* `center`：中央に割り付ける

* `ljust`：左側に割り付ける

* `rjust`：右側に割り付ける

```ruby
>> a = "abc"
=> "abc"
>> a.center(20)
=> "        abc         "
>> a.center(20, "*")
=> "********abc*********"
>> a.ljust(20)
=> "abc                 "
>> a.rjust(20, "-")
=> "-----------------abc"
```

***

### 5-6-12.批評文字列を変換する

* `dump`：文字列の中にある改行コード(`"\n"`)や、タブ文字(`"\t"`)の非表示文字列を
  バックスラッシュ記法に置き換えた文字列を返す
  →例)改行コード(`"\n"`)、タブ文字(`"\t"`)

```ruby
>> a = "abc\tdef\tghi\n"
=> "abc\tdef\tghi\n"
>> puts a
abc	def	ghi
=> nil
>> puts a.dump
"abc\tdef\tghi\n"
=> nil
```

***

### 5-6-13.文字列をアンパックする

* `unpack`： **Array#pack** メソッドでパックされた文字列を、指定したテンプレートにし違ってアンパックする

例)MINEエンコードされた文字列のデコードや、バイナリデータから数値への変換などに利用

```ruby
# MINEエンコードされた文字列をアンパックする
>> '440r440T4408'.unpack('m')
=> ["\xE3\x8D+\xE3\x8D\x13\xE3\x8D<"]
```

* アンパックされた文字列のエンコーディングは、ASCII-8BITになる

```ruby
>> a = '440r440T4408'.unpack('m').first
=> "\xE3\x8D+\xE3\x8D\x13\xE3\x8D<"
>> a.force_encoding('UTF-8')
=> "\xE3\x8D+\xE3\x8D\u0013\xE3\x8D<"
```

***

### 5-6-14.文字列内での検索

* `include?`：指定した文字列が含まれている場合に`true`を返す

* `index`：指定された文字や文字コードの数値・正規表現などのパターンを、指定された一から右方向に検索して、
  最初に見つかった位置を返す

* `rindex`：`index`メソッドとは逆に左方向に検索する

* `match`：指定された正規表現によるマッチングを行い、マッチした場合には`MatchData`オブジェクトを返す

* `scan`：指定された正規表現にマッチした部分文字列の配列を返す。ブロックを渡すことも可能

```ruby
>> "abcdefg".include?("abc")
=> true
>> "abcdefg".index("bc")
=> 1
>> "abcdefg".match(/[c-f]/)
=> #<MatchData "c">
>> "abcdefg".scan(/[c-d]/)
=> ["c", "d"]
```

***

### 5-6-15.次の文字列を求める

* `succ`、`next`：次の文字列を求める

ここでの **次の文字列** とは、、、

対象となる文字列の右端がアルファベットならアルファベット順に、数字であれば10進数の数値と見なして計算

計算時に桁上がりが発生した場合は、それに応じて１つ左の文字が変化するか、文字列が伸長する

```ruby
>> "abc123".succ
=> "abc124"
>> "123abc".succ
=> "123abd"
>> "123xyz".succ
=> "123xza"
>> "99999".succ
=> "100000"
>> "zzzzz".succ
=> "aaaaaa"
>> "ZZZZZ".succ
=> "AAAAAA"
```

***

### 5-6-16.文字列に対する繰り返し

* `each_line`、`lines`：文字列の各行に対して繰り返す。オプションとして行の区切りを指定可能

* `each_byte`、`bytes`：文字列をバイト単位で繰り返す

* `each_char`、`chars`：文字単位で繰り返す

* `upto`：自分自身から指定された文字列まで、`succ`メソッドで生成される次の文字列を使って繰り返す

```ruby
>> "abc\ndef\nghi".each_line {|c| puts c}
abc
def
ghi
=> "abc\ndef\nghi"
>> "abc".each_byte {|c| puts c
>> }
97
98
99
=> "abc"
>> "ルビー".each_char {|c| puts c}
ル
ビ
ー
=> "ルビー"
>> "a".upto("c") {|c| puts c}
a
b
c
=> "a"
```

***

### 5-6-17.他のクラスへの変換

#### `hex`メソッド

* 文字列が16進数であるとして、数値に変換

* 接頭辞`0x`、`0X`とアンダースコア`_`は無視され、
  16進数以外の文字がある場合にはそれ以降も無視

```ruby
>> "abc".hex
=> 2748
>> "azc".hex
=> 10
>> "0xazc".hex
=> 10
```

***

#### `oct`メソッド

* 文字列が8進数であるとして、数値に変換

* `hex`メソッドとは異なり、接頭辞に応じて8進数以外の変換も行う

```ruby
>> "010".oct
=> 8
>> "0b010".oct
=> 2
>> "0x010".oct
=> 16
```

***

#### `to_i`メソッド

* 文字列を10進数の整数であるとして、数値に変換する

* 整数として見なせない文字があれば、そこまでを変換して以降を無視する

* 空文字であれば、0を返す

```ruby
>> "123".to_i
=> 123
>> "0123".to_i
=> 123
>> "0x123".to_i
=> 0
>> "".to_i
=> 0
```

***

#### `to_f`メソッド

* 文字列を10進数の小数として、`Float`オブジェクトに変換する

* 小数と見なせない文字があれば、そこまでを変換して以降を無視する

* 空文字であれば、0.0を返す

```ruby
>> "1.23".to_f
=> 1.23
>> "01.23".to_f
=> 1.23
>> "0x1.23".to_f
=> 0.0
>> "".to_f
=> 0.0
```

***

#### `to_s`、`to_str`メソッド

* 自分自身を返す

```ruby
>> "ルビー".to_s
=> "ルビー"
>> "ルビー".to_str
=> "ルビー"
```

***

#### `to_sym`、`intern`メソッド

* 文字列に対応するシンボル値を返す

```ruby
=> :abc
>> a.object_id
=> 70263022623280    # オブジェクトid
>> b = :abc
=> :abc
>> b.object_id
=> 70263022623280    # 同じオブジェクトid
```

***

## 5-7.`Array`クラス

### 5-7-1.配列の生成

* `配列式`と呼ばれる記法で配列を生成

* `Array`クラスのクラスメソッドである、 **`[]`メソッド** や **`new`メソッド** を使って配列を生成

```ruby
# 配列の生成(1)
>> a = [1, 2, 3]
=> [1, 2, 3]
>> a.class
=> Array

# 配列の生成(2)
>> Array[1, 2, 3]            # `[]`メソッドは、引数に要素を指定
=> [1, 2, 3]

# 配列の生成(3)
>> Array.new(3, "str")       # `new`メソッドでの、配列の長さと初期値を指定
=> ["str", "str", "str"]

# 配列の生成(4)
>> Array.new([1, 2, 3])      # `new`メソッドでの、引数に配列を指定する方法(配列の複製)
=> [1, 2, 3]

# 配列の生成(5)
>> Array.new(3) {|i| i * 3}  # `new`メソッドに、配列の長さとブロックを渡す方法(ブロックには、配列のインデックスを渡す)
=> [0, 3, 6]
```

***

### 5-7-2.配列に要素を追加する

* `<<`メソッド・`push`メソッド：指定された引数にあるオブジェクトを、自分自身の末尾に追加

* `concat`メソッド：指定された配列を、自分自身の末尾に連結

* `insert`メソッド：1番目の引数で指定された場所に、それ以降で指定されたオブジェクトを挿入

* `unshift`メソッド：指定されたオブジェクトを、配列の先頭に追加

* これらの5つのメソッドは、破壊的メソッド
  →元のオブジェクトの内容を書き換える

* `+`メソッド：自分自身と引数で与えられた配列を連結した配列を新たに生成して返す

```ruby
>> a = [1, 2, 3]
=> [1, 2, 3]
>> a << 4                      # 末尾に[4]を追加
=> [1, 2, 3, 4]
>> a.concat [5, 6]             # 末尾に[5,6]を追加
=> [1, 2, 3, 4, 5, 6]
>> a.insert(3, 9)              # 配列の3番目に、[9]を追加
=> [1, 2, 3, 9, 4, 5, 6]
>> a.object_id
=> 70235495510920
>> b = a + [10]                # 配列aの末尾に、10を追加
=> [1, 2, 3, 9, 4, 5, 6, 10]
>> b.object_id
=> 70235499700020
>> a = [1, 2, 3]
=> [1, 2, 3]
>> a.unshift(10)               # 配列aの先頭に、[10]を追加
=> [10, 1, 2, 3]
```

***

### 5-7-3.配列の要素を変更する

* `[]=`メソッド：指定したインデックスにある要素を書き換える

※インデックスは、整数や`Range`オブジェクト、始点と終点を指定可能。
配列の要素数よりも大きな数が指定された場合は、自動的に配列の長さが伸長され、その部分は`nil`で初期化される。

* `fill`メソッド：配列の全ての要素を指定したオブジェクトに変更

※引数を2つ以上取る場合は、配列の始点と終点や、`Range`オブジェクトを取ることもでき、その場合は該当する部分のみ変更する。
ブロックを取ることもでき、その場合はブロックの評価結果で要素を変更する。

* `replace`メソッド：引数で指定された配列で、自分自身の内容を置き換える。オブジェクトIDが変化しない。

```ruby
# `[]=`メソッド
>> a = [1, 2, 3]
=> [1, 2, 3]
>> a[1] = 10          # 配列の1番目に、10を追加
=> 10
>> a
=> [1, 10, 3]
>> a[1..2] = [11, 12] # 配列の1〜2番目に、11と12を追加
=> [11, 12]
>> a
=> [1, 11, 12]
>> a[8] = 8           # 配列の8番目に、8を追加(3〜7番目は`nil`)
=> 8
>> a
=> [1, 11, 12, nil, nil, nil, nil, nil, 8]

# `fill`メソッド
>> a = [1, 2, 3]
=> [1, 2, 3]
>> a.fill("s")                   # aの配列の内容を、全て"s"に変更
=> ["s", "s", "s"]
>> a.fill("t", 1..2)             # aの配列のうち、1〜2番目を"t"に変更
=> ["s", "t", "t"]
>> a
=> ["s", "t", "t"]
>> a.fill(1..2){|index| index}   # aの配列のうち、1〜2番目を1、2に変更
=> ["s", 1, 2]

# `replace`メソッド
>> a = [1, 2, 3]
=> [1, 2, 3]
>> a.object_id
=> 70273659485340
>> a = [1, 2, 3]
=> [1, 2, 3]
>> a.object_id          # 異なるオブジェクトID
=> 70273659433960
>> a.replace([4, 5, 6])
=> [4, 5, 6]
>> a.object_id
=> 70273659433960       # 置換する前と同じオブジェクトID
```

***

### 5-7-4.配列の要素を参照する

* `[]`メソッド：変更の場合と同様に、 **整数** や **`Range`オブジェクト** 、始点と終点で指定したインデックスに対応する要素を返す

* `slice`メソッド：`[]`メソッドと同様

* `at`メソッド：`[]`メソッドと同様だが、インデックスが整数の場合に利用する。要素数よりも大きい場合は、`nil`を返す

* `values_at`メソッド：`[]`メソッドと同様の動作をするが、結果を配列で返す

```ruby
>> a = [1, 2, 3]
=> [1, 2, 3]
>> a[1]            # 配列aの1番目の値を取り出す
=> 2
>> a.at(1)         # 配列aの1番目の値を取り出す
=> 2
>> a[1..2]         # 配列aの1〜2番目の配列を取り出す
=> [2, 3]
>> a.values_at(1)  # 配列aの1番目の配列を取り出す
=> [2]
```

***

* `fetch`メソッド：`at`メソッドと同様。引数がインデックスのみの場合は、`IndexError`が発生する。
  2番目の引数がある場合にはその値を、ブロックを取っている場合はその評価結果を返す

* `first`メソッド：配列の先頭を返す。引数が指定された場合は、先頭から指定した数だけ要素を返す

* `last`メソッド：配列の末尾の要素を返す。引数が指定された場合は、末尾から指定した数だけ要素を返す

* `assoc`メソッド：配列の配列を検索し、その配列の最初の要素が指定された値と`==`で等しければ、その配列を返す
  該当する要素がなければ、`nil`を返す

* `rassoc`メソッド：配列の配列を検索し、その配列のインデックス1の要素が指定された値と`==`で等しければ、その要素を返す

```ruby
# `fetch`メソッド
>> a = [1, 2, 3]
=> [1, 2, 3]
>> a.fetch(4)
IndexError: index 4 outside of array bounds: -3...3
>> a.fetch 4, "ERROR"
=> "ERROR"
>> a.fetch(4){|n| "ERROR #{n}"}
=> "ERROR 4"

# `first`、`last`メソッド
>> a = [1, 2, 3, 4, 5]
=> [1, 2, 3, 4, 5]
>> a.first
=> 1
>> a.last
=> 5
>> a.first(3)   # 最初から3個まで出力
=> [1, 2, 3]

# `assoc`メソッド
>> a = [[1, 2], [3, 4], [5, 6], [7, 8]]
=> [[1, 2], [3, 4], [5, 6], [7, 8]]
>> a.assoc(3)     # 配列の最初の要素が3の配列は、[3, 4]
=> [3, 4]

# `rassoc`メソッド
>> a = [[1, 2], [3, 4], [5, 6], [7, 8]]
=> [[1, 2], [3, 4], [5, 6], [7, 8]]
>> a.rassoc(4)    # 配列の要素に4が含まれる配列は、[3, 4]
=> [3, 4]
```

***

### 5-7-5.配列の要素を調べる

* `include?`メソッド：指定された値が要素の中に存在する場合に、 **真** を返す

* `index`、`rindex`メソッド：それぞれの配列の先頭と末尾から指定された値と`==`で等しい要素の位置を返す
  見つからない場合は、`nil`を返す

```ruby
>> a = [1, 2, 3, 4, 5]
=> [1, 2, 3, 4, 5]
>> a.include?(3)
=> true
>> a.include?(10)
=> false
>> a.index(4)
=> 3
>> a.rindex(4)
=> 3
```

***

### 5-7-6.配列の要素を削除する

* `delete_at`メソッド：指定されたインデックスに対応する要素を取り除き、その要素を返す

* `delte_if`・`reject!`メソッド：ブロックに要素を渡し、その評価結果が真になった要素を全て取り除いた自分自身を返す

* `delete`メソッド：指定された値と`==`メソッドで等しい要素があれば、取り除いてその値を、なければ`nil`を返す

* `clear`メソッド：要素を全て削除する

* `slice!`メソッド：指定されたインデックスに対応する要素を取り除き、その取り除いた要素を返す。(インデックスには整数やRangeオブジェクト、視点と長さを指定可能)

* `shift`メソッド：先頭から指定された数だけ要素を取り除いて返す。指定がなければ、1が指定されたとして先頭の要素を返す

* `pop`メソッド：末尾から指定された数だけ要素を取り除いて返す。指定がなければ、1が指定されたとして末尾の要素を返す(`shift`メソッドの逆)

* `-`メソッド：指定された配列にある要素を、自分自身から取り除いた配列を返す

```ruby
# delete_atメソッド
>> a = [1, 2, 3, 4, 5]
=> [1, 2, 3, 4, 5]
>> a.delete_at(2)
=> 3
>> a
=> [1, 2, 4, 5]

# delete_ifメソッド
>> a = [1, 2, 3, 4, 5]
=> [1, 2, 3, 4, 5]
>> a.delete_if{|n| n % 2 == 0}
=> [1, 3, 5]

# deleteメソッド
>> a = [1, 2, 3, 4, 5]
=> [1, 2, 3, 4, 5]
>> a.delete(3)
=> 3
>> a
=> [1, 2, 4, 5]
>> a.delete(10)
=> nil
>> a
=> [1, 2, 4, 5]

# clearメソッド
>> a = [1, 2, 3, 4, 5]
=> [1, 2, 3, 4, 5]
>> a.clear
=> []

# slice!メソッド
>> a = [1, 2, 3, 4, 5]
=> [1, 2, 3, 4, 5]
>> a.slice!(2,2)        # 2番目から2つの値を削除(2、3番目の値を削除)
=> [3, 4]
>> a
=> [1, 2, 5]

# shiftメソッド
>> a = [1, 2, 3, 4, 5]
=> [1, 2, 3, 4, 5]
>> a.shift(2)           # 先頭から2個削除
=> [1, 2]
>> a
=> [3, 4, 5]
>> a.shift              # 先頭から1個削除
=> 3
>> a
=> [4, 5]

# popメソッド
>> a = [1, 2, 3, 4, 5]
=> [1, 2, 3, 4, 5]
>> a.pop(2)             # 末尾から2個削除
=> [4, 5]
>> a
=> [1, 2, 3]
>> a.pop                # 末尾から1個削除
=> 3
>> a
=> [1, 2]

# -メソッド
>> a = [1, 2, 3, 4, 5]
=> [1, 2, 3, 4, 5]
>> a - [1, 2]
=> [3, 4, 5]
>> a - [1, 3, 5, 7]
=> [2, 4]
```

***

### 5-7-7.配列の演算

Rubyでは、配列を集合と見なした演算ができる

* `|`メソッド：和集合を求める。両方のいずれかに含まれる要素を含む配列を返す

* `&`メソッド：積集合を求める。両方に含まれる要素を含む配列を返す

```ruby
# |メソッド
>> [1, 2, 3] | [1, 3, 5]
=> [1, 2, 3, 5]

# &メソッド
>> [1, 2, 3] & [1, 3, 5]
=> [1, 3]
```

***

### 5-7-8.配列の比較

* `==`メソッド：配列同士を比較する。先頭から要素を比較し、全ての要素が等しければその結果を返す

* `<=>`メソッド：先頭から比較し、同じであれば`0`、左辺が大きければ`1`、右辺が大きければ`-1`を返す

```ruby
>> [1, 2, 3] == [1, 3, 5]
=> false
>> [1, 2, 3] <=> [1, 3, 5]
=> -1
```

***

### 5-7-9.配列の要素での繰り返し

* `each`メソッド：ブロックに各要素が渡る

* `each_index`メソッド：配列のインデックスが渡る

* `reverse_each`メソッド：要素の逆順に繰り返す

* `cycle`メソッド：配列の要素を順に繰り返し、末尾まで来たらまた先頭に戻って繰り返し続ける

```ruby
# eachメソッド
>> [1, 3, 5, 7, 9].each{|n| puts n * 2}
2
6
10
14
18
=> [1, 3, 5, 7, 9]

# each_indexメソッド
>> [1, 3, 5, 7, 9].each_index{|n| puts n * 2}
0
2
4
6
8
=> [1, 3, 5, 7, 9]

# reverse_eachメソッド
>> [1, 3, 5, 7, 9].reverse_each{|n| puts n * 2}
18
14
10
6
2
=> [1, 3, 5, 7, 9]

# cycleメソッド
>> [1, 2, 3].cycle{|n| puts n }
1
2
3
1
2
3
# ...(省略)
```

***

### 5-7-10.配列の要素を連結する

* `join`、`*`メソッド：配列の要素を指定された区切り文字で連結した文字列を返す

```ruby
>> [1, 2, 3].join(",")
=> "1,2,3"
```

***

### 5-7-11.配列の長さを求める

* `length`、`size`メソッド：配列の長さを求める。空の場合は0を返す

* `empty?`メソッド：配列が空の場合に真を返す

```ruby
>> [1, 2, 3].length
=> 3
>> [].length
=> 0
>> [].empty?
=> true
```

***

### 5-7-12.配列をソートする

* `sort`メソッド：配列をソートする。要素同士の比較は、`<=>`メソッドが使用される。

* ブロックが与えられた場合、2つの要素が渡され、その評価結果に応じてソートされる。

  * 正の整数：1つ目の要素が大きい

  * 0：同じ

  * 負の整数：2つ目の要素が大きい


```ruby
>> a = [1, 3, 5, 2, 4, 6]
=> [1, 3, 5, 2, 4, 6]
>> a.sort
=> [1, 2, 3, 4, 5, 6]
>> a
=> [1, 3, 5, 2, 4, 6]
>> a.sort!                 # 破壊的メソッド(自分自身をソート結果で書き換える)
=> [1, 2, 3, 4, 5, 6]
>> a
=> [1, 2, 3, 4, 5, 6]

# ブロックが与えられた場合
>> [1, 3, 5, 2, 4, 6].sort{|a, b| a <=> b}    # 0番目が1・・・bの方が大きいので負の整数
=> [1, 2, 3, 4, 5, 6]
>> [1, 3, 5, 2, 4, 6].sort{|a, b| b <=> a}    # 0番目が1・・・bの方が小さいので正の整数
=> [6, 5, 4, 3, 2, 1]
```

このブロックはわからないので、省略
2018/9/8

***

### 5-7-13.配列を変換する

* `uniq`メソッド：配列から重複した要素を取り除いた配列を返す

* `uniq!`メソッド：`uniq`メソッドと同じだが、取り除かれなかった場合に`nil`を返す

* `compact`メソッド：要素の`nil`を取り除いた配列を返す。

* `compact!`メソッド：`compact`メソッドと同じだが、取り除かれなかった場合に`nil`を返す

* `reverse`メソッド：配列の要素を逆順に並べ替えた配列を返す

* `flatten`メソッド：配列を再帰的に平滑化した配列を返す。引数が指定された場合は、その深さまで再帰的に平滑化する。

* `flatten!`メソッド：`flatten`メソッドと同じだが、平滑化されなかった場合には`nil`を返す

* `map`、`collect`メソッド：要素ごとにブロックを評価し、その結果で要素を書き換えた配列を返す

* `shuffle`メソッド：配列の要素をシャッフルして返す

```ruby
# uniqメソッド
>> [1, 1, 2, 3, 3].uniq
=> [1, 2, 3]

# compactメソッド
>> [1, nil, 2, nil, 3].compact
=> [1, 2, 3]

# reverseメソッド
>> [1, 2, 3, 4, 5].reverse
=> [5, 4, 3, 2, 1]

# flattenメソッド
>> [[[1, 2], 3], [[4, 5], 6]].flatten
=> [1, 2, 3, 4, 5, 6]
>> [[[1, 2], 3], [[4, 5], 6]].flatten(1)
=> [[1, 2], 3, [4, 5], 6]

# mapメソッド
>> [1, 2, 3, 4, 5].map{|n| n * 2}
=> [2, 4, 6, 8, 10]

# shuffleメソッド
>> [1, 2, 3, 4, 5].shuffle
=> [5, 1, 2, 3, 4]
```

***

### 5-7-14.配列を組み合わせて生成する

* `product`メソッド：自身と与えられた配列から1つずつ要素を取って組み合わせた配列を作り、その全ての組み合わせを要素とする配列を返す

* `zip`メソッド：自身と与えられた配列から1つずつ要素を取って配列を作り、それを要素とする配列を返す。`product`メソッドとは異なり、
  組み合わせは前から順に同じインデックスのもののみとなる。ブロックが与えられた場合は、自分自身と引数に指定された配列を順にブロックに渡す。

* `*`メソッド：与えられた数値の分だけ、繰り返した配列を返す

```ruby
# productメソッド
>> [1, 2].product(["a", "b", "c"])
=> [[1, "a"], [1, "b"], [1, "c"], [2, "a"], [2, "b"], [2, "c"]]

# zipメソッド
>> [1, 2].zip(["a", "b"])
=> [[1, "a"], [2, "b"]]
>> [1, 2].zip(["a", "b"], ["x", "y"])
=> [[1, "a", "x"], [2, "b", "y"]]

# *メソッド
>> [1, 2] * 4
=> [1, 2, 1, 2, 1, 2, 1, 2]
```

***

### 5-7-15.配列をパックする

* `pack`メソッド：自身を指定されたテンプレートに従ってパックする。

```ruby
>> ["ルビー"].pack('m')
=> "44Or44OT44O8\n"
```

***

## 5-8.`Hash`クラス

ハッシュは連想配列とも呼ばれ、配列でのインデックスにあたるキーとして、数値以外のRubyオブジェクトを利用可能

Rubyでは、ハッシュはHashクラスのオブジェクトとして生成される

### 5-8-1.ハッシュの生成

* ハッシュは、ハッシュ式と呼ばれる記法、`[]`メソッド、`new`メソッドを使用して生成できる

* ハッシュ式では、 **キー** と要素である **値** とを`=>`を使った組み合わせで表現する

```ruby
# ハッシュ式
>> a = {"apple" => "fruit", "coffee" => "drink"}
=> {"apple"=>"fruit", "coffee"=>"drink"}
>> a.class
=> Hash

# []メソッド
>> Hash["apple", "fruit", "coffee", "drink"]   # キーと値を、順番にカンマ`,`で列挙する
=> {"apple"=>"fruit", "coffee"=>"drink"}

# newメソッド
>> a = Hash.new
=> {}
>> a["apple"]
=> nil
>> a = Hash.new("NONE")   # キーが存在しない場合の初期値を設定できる
=> {}
>> a["apple"]
=> "NONE"

# newメソッド(ブロック)
>> a = Hash.new{|hash, key| hash[key] = nil}      # 初期値の設定
=> {}
>> a["apple"]
=> nil
>> a = Hash.new{|hash, key| hash[key] = "NONE"}
=> {}
>> a["apple"]
=> "NONE"
```

* 初期値とブロックの参照は、

  * 初期値：`default`メソッド

  * ブロック：`default_proc`メソッド

  で参照可能

* 初期値は、あとで`default=`メソッドで指定可能

```ruby
>> a = Hash.new("NONE")
=> {}
>> a.default
=> "NONE"
>> a["apple"]
=> "NONE"
>> a.default = "Not exists"
=> "Not exists"
>> a["apple"]
=> "Not exists"
```

***

### 5-8-2.ハッシュの **キー** や **値** を取得する

* `[]`メソッド：指定されたキーに対応する値を返す

* `keys`、`values`メソッド：ハッシュの全てのキーと値の配列を生成する

* `values_at`メソッド：指定されたキーに対応する値を、配列で返す

* `fetch`メソッド：与えられたキーに対する値を返す。キーが存在しない場合には、2番目の引数が与えられた場合にはその値を、ブロックが与えられていた場合はそのブロックを評価した結果を返す

* `select`メソッド：キーと値の組み合わせについてブロックを評価して、結果が真となる組み合わせのみを含むハッシュを返す

* `find_all`メソッド：キーと値の組み合わせについてブロックを評価するが、返り値はキーと値の配列

```ruby
# []メソッド
>> a = {"apple" => "fruit", "coffee" => "drink"}
=> {"apple"=>"fruit", "coffee"=>"drink"}
>> a["apple"]
=> "fruit"

# keys、valuesメソッド
>> a = {"apple" => "fruit", "coffee" => "drink"}
=> {"apple"=>"fruit", "coffee"=>"drink"}
>> a.keys
=> ["apple", "coffee"]
>> a.values
=> ["fruit", "drink"]

# values_atメソッド
>> a = {1 => "a", 2 => "b", 3 => "c", 4 => "d"}
=> {1=>"a", 2=>"b", 3=>"c", 4=>"d"}
>> a.values_at(1, 3)
=> ["a", "c"]

# fetchメソッド
>> a = {1 => "a", 2 => "b", 3 => "c", 4 => "d"}
=> {1=>"a", 2=>"b", 3=>"c", 4=>"d"}
>> a.fetch(5, "NONE")
=> "NONE"
>> a.fetch(5){|key| % 2 == 0}
=> false

# selectメソッド
>> a = {1 => "a", 2 => "b", 3 => "c", 4 => "d"}
=> {1=>"a", 2=>"b", 3=>"c", 4=>"d"}
>> a.select{|key, value| key % 2 == 0}
=> {2=>"b", 4=>"d"}
>> a.find_all{|key, value| key % 2 == 0}
=> [[2, "b"], [4, "d"]]
```

***

### 5-8-3.ハッシュを変更する

#### `[]=`メソッド

配列の場合と同様に、指定されたキーに対応する値を変更する。キーが存在しない場合には、そのキーと値を登録する。

```ruby
>> a = {"apple" => "fruit", "coffee" => "drink"}
=> {"apple"=>"fruit", "coffee"=>"drink"}
>> a["apple"] = "red"
=> "red"
>> a
=> {"apple"=>"red", "coffee"=>"drink"}
>> a["orange"] = "orange"
=> "orange"
>> a
=> {"apple"=>"red", "coffee"=>"drink", "orange"=>"orange"}
```

#### `delete`メソッド

指定されたキーに対応する値を取り除く。キーが存在していれば対応する値を、そうでなければ`nil`を返す。ブロックが与えられた場合には、キーが存在しない場合にブロックの評価結果を返す

```ruby
>> a = {"apple" => "fruit", "coffee" => "drink"}
=> {"apple"=>"fruit", "coffee"=>"drink"}
>> a.delete("apple")
=> "fruit"
>> a
=> {"coffee"=>"drink"}
```

#### `reject`メソッド

ブロックを評価した結果が真になる値を取り除いたハッシュを生成して返す。元のオブジェクトは変更されない。
→`reject!`とは異なるメソッド(1要素ずつブロックを要素に渡し、その評価結果が真になった要素を全て取り除いた自分自身を返す)

```ruby
>> a = {"apple" => "fruit", "coffee" => "drink"}
=> {"apple"=>"fruit", "coffee"=>"drink"}
>> a.reject{|key, value| value == "drink"}         # 値が"drink"であるものを取り除く
=> {"apple"=>"fruit"}
>> a
=> {"apple"=>"fruit", "coffee"=>"drink"}
```

#### `delete_if`、`reject!`メソッド

ブロックを評価した結果が真になる値を取り除く。元のオブジェクトが変更される

```ruby
>> a = {"apple" => "fruit", "coffee" => "drink"}
=> {"apple"=>"fruit", "coffee"=>"drink"}
>> a.reject!{|key, value| value == "drink"}
=> {"apple"=>"fruit"}
>> a
=> {"apple"=>"fruit"}
```

#### `replace`メソッド

引数で与えられたハッシュで自分自身を置き換える。

```ruby
>> a = {"apple" => "fruit", "coffee" => "drink"}
=> {"apple"=>"fruit", "coffee"=>"drink"}
>> a.object_id
=> 70364050765140
>> a.replace({"orange" => "fruit", "tea" => "drink"})
=> {"orange"=>"fruit", "tea"=>"drink"}
>> a.object_id
=> 70364050765140       # 同じオブジェクトID・・・自分自身の置き換え
```

#### `shift`メソッド

ハッシュから先頭のキーと値の組み合わせを1つ取り除き、その組み合わせを配列として返す

```ruby
>> a = {"apple" => "fruit", "coffee" => "drink"}
=> {"apple"=>"fruit", "coffee"=>"drink"}
>> a.shift
=> ["apple", "fruit"]
>> a
=> {"coffee"=>"drink"}
```

#### `merge`メソッド

自分自身と引数で指定されたハッシュを統合した、新しいハッシュオブジェクトを返す。

デフォルト値は自分自身の設定が引き継がれる。

ブロックが与えられた場合は、キーと自分自身の値、指定されたハッシュの値が渡され、ブロックの評価結果が新しいハッシュの値となる

```ruby
>> a = {"apple" => "foods", "coffee" => "drink"}
=> {"apple"=>"foods", "coffee"=>"drink"}
>> a.merge({"orange" => "fruit", "tea" => "drink", "apple" => "fruit"})
=> {"apple"=>"fruit", "coffee"=>"drink", "orange"=>"fruit", "tea"=>"drink"}
>> a
=> {"apple"=>"foods", "coffee"=>"drink"}
>> a.merge({"orange" => "fruit", "tea" => "drink"}){|key, self_val, other_val| self_val}
=> {"apple"=>"foods", "coffee"=>"drink", "orange"=>"fruit", "tea"=>"drink"}
```

#### `merge!`、`update`メソッド

自分自身と引数で指定されたハッシュを統合する。

`merge`メソッドとは異なり、元のオブジェクトが変更される。

```ruby
>> a = {"apple" => "foods", "coffee" => "drink"}
=> {"apple"=>"foods", "coffee"=>"drink"}
>> a.merge!({"orange" => "fruit", "tea" => "drink", "apple" => "fruit"})
=> {"apple"=>"fruit", "coffee"=>"drink", "orange"=>"fruit", "tea"=>"drink"}
>> a
=> {"apple"=>"fruit", "coffee"=>"drink", "orange"=>"fruit", "tea"=>"drink"}
```

#### `invert`メソッド

キーと値を逆にしたハッシュを返す。

ただし、値が重複している場合には、結果は不定になる

```ruby
>> a = {"apple" => "foods", "coffee" => "drink"}
=> {"apple"=>"foods", "coffee"=>"drink"}
>> a.invert
=> {"foods"=>"apple", "drink"=>"coffee"}
>> {"orange" => "fruit", "coffee" => "drink", "apple" => "fruit", "tea" => "drink"}.invert
=> {"fruit"=>"apple", "drink"=>"tea"}
```

#### `clear`メソッド

ハッシュを空にする

```ruby
>> a = {"apple" => "foods", "coffee" => "drink"}
=> {"apple"=>"foods", "coffee"=>"drink"}
>> a.clear
=> {}
```

***

### 5-8-4.ハッシュを調べる

* `length`、`size`メソッド：ハッシュの組み合わせの数を返す

* `empty?`メソッド：ハッシュが空かどうかを調べる

* `has_key?`、`include?`、`key?`、`member?`メソッド：ハッシュに **キー** が存在する場合に真を返す

* `has_value?`、`valid?`メソッド：ハッシュに **値** が存在する場合に真を返す

```ruby
# size、empty?メソッド
>> a = {"apple" => "fruit", "coffee" => "drink"}
=> {"apple"=>"fruit", "coffee"=>"drink"}
>> a.size
=> 2
>> a.empty?
=> false

# key?メソッド
>> a = {"apple" => "fruit", "coffee" => "drink"}
=> {"apple"=>"fruit", "coffee"=>"drink"}
>> a.key?("apple")
=> true
>> a.key?("orange")
=> false

# value?メソッド
>> a = {"apple" => "fruit", "coffee" => "drink"}
=> {"apple"=>"fruit", "coffee"=>"drink"}
>> a.value?("fruit")
=> true
>> a.key?("foods")
=> false
```

***

### 5-8-5.ハッシュを使った繰り返し

* `each`、`each_pair`メソッド：与えられたブロックに **キー** と **値** を渡して評価する

* `each_key`、`each_value`メソッド： **キー** と **値** を与えられたブロックに渡して評価する

```ruby
# eachメソッド
>> a = {"apple" => "fruit", "coffee" => "drink"}
=> {"apple"=>"fruit", "coffee"=>"drink"}
>> a.each{|key, value| puts "#{key} => #{value}\n"}
apple => fruit
coffee => drink
=> {"apple"=>"fruit", "coffee"=>"drink"}

# each_key、each_valueメソッド
>> a = {"apple" => "fruit", "coffee" => "drink"}
=> {"apple"=>"fruit", "coffee"=>"drink"}
>> a.each_key{|key| puts "key: #{key}\n"}
key: apple
key: coffee
=> {"apple"=>"fruit", "coffee"=>"drink"}
```

***

### 5-8-6.ハッシュをソートする

* `sort`メソッド：ハッシュとキーと値の組み合わせの配列に変換し、それをソートした結果を返す。

* ハッシュ自身はソートされない。ブロックが与えられた場合には、キーと値の組み合わせの配列が渡される。

```ruby
>> a = {4 => "a", 3 => "b", 2 => "c", 1 => "d"}
=> {4=>"a", 3=>"b", 2=>"c", 1=>"d"}
>> a.sort
=> [[1, "d"], [2, "c"], [3, "b"], [4, "a"]]
>> a.sort{|a, b| a[1] <=> b[1]}
=> [[4, "a"], [3, "b"], [2, "c"], [1, "d"]]
```

***

### 5-8-7.ハッシュを変換する

* `to_a`：ハッシュを配列に変換する。キーと値の組み合わせを配列の配列として生成する。

```ruby
>> a = {4 => "a", 3 => "b", 2 => "c", 1 => "d"}
=> {4=>"a", 3=>"b", 2=>"c", 1=>"d"}
>> a.to_a
=> [[4, "a"], [3, "b"], [2, "c"], [1, "d"]]
```

***

## 5-9.`Symbol`クラス

コロン`:`で始まる文字列で表す

変更不可であり、また同じ表記であれば必ず同じオブジェクトIDであるため、必ず同値である。

### `Symbol`オブジェクト

```ruby
>> a = :foo
=> :foo
>> a.object_id
=> 1159388
>> b = :foo
=> :foo
>> b.object_id
=> 1159388      # 変数は異なるが、同じオブジェクトID
```

* Ruby内部では、メソッド名や変数名などの **名前** は整数値で管理されている。この整数をRubyのコード上で表現したものがSymbol

* 文字列そのものが必要でない場合に利用すると良い。例えば、ハッシュのキーなど、名前が重要な場合によく利用する。

***

### 5-9-1.定義済みの`Symbol`オブジェクトを取得する

* `Symbol.all_symbols`メソッド：定義済みのSymbolオブジェクトを取得できる

```ruby
>> :foo
=> :foo
>> Symbol.all_symbols
=> [:!, :"\"", :"#", :"$", :%, :&, :"'", #・・・省略
```

***

### 5-9-2.`Symbol`オブジェクトに対応する文字列を取得する

* `id2name`、`to_s`メソッド：`Symbol`オブジェクトに対応する文字列を取得する

```ruby
>> :foo.to_s
=> "foo"
>> :foo.id2name
=> "foo"
```

***

## 5-10.`Dir`クラス、`File`クラス、`IO`クラス

* `Dir`クラス：ディレクトリの移動や作成、ディレクトリ内のファイル一覧の取得など、ディレクトリを扱うクラス

* `File`クラス：ファイルの読み取りや書き込み、新規作成や削除など、ファイルを扱うクラス

* `IO`クラス：`File`クラスのスーパークラスで、ファイルやプロセスなどとの入出力を扱うクラス

***

### 5-10-1.`Dir`クラス

#### ディレクトリを開く、閉じる

* `open`メソッド：ディレクトリを開く。返り値は`Dir`クラスのオブジェクトで、例えば`each`メソッドでファイル一覧を取得できる

* `close`メソッド：開いたディレクトリを閉じる

```ruby
>> dir = Dir.open("/usr/local/bin")
=> #<Dir:/usr/local/bin>
>> dir.each{|file| puts file}
.
..
pg_standby
pg_rewind
# ・・・(省略)・・・
convert
pg_dump
pydoc2
=> #<Dir:/usr/local/bin>
>> dir.close
=> nil
```

***

#### 開いているディレクトリのパスの取得

* `path`メソッド：開いているディレクトリのパスを取得

* `Dir.open`メソッドはブロックを取ることができ、この場合はブロックを出るときに自動的に閉じられる

```ruby
>> Dir.open("/usr/local/bin"){|dir| puts dir.path}
/usr/local/bin
=> nil          # 自動的に閉じられる
```

***

#### カレントディレクトリの取得

* `Dir.pwd`、`Dir.getwd`メソッド：カレントディレクトリを取得する

```ruby
>> Dir.pwd
=> "/Users/MacUser/work/rails/shared_hobby"
```

***

#### カレントディレクトリの移動

* `chdir`メソッド：カレントディレクトリを指定されたディレクトリに変更する。

* 指定がない場合、環境変数 **HOME** や **LOGDIR** が設定されていれば、そのディレクトリに移動する

* ブロックが与えられた場合には、そのブロック内でのみディレクトリを移動し、ブロックを出るときに元に戻る。
  ディレクトリの移動に成功すれば0を返す

```ruby
>> Dir.chdir("/usr/local")
=> 0
>> Dir.pwd
=> "/usr/local"
>> Dir.chdir("/usr/local/bin"){|dir| puts Dir.pwd}
/usr/local/bin
=> nil
>> Dir.pwd
=> "/usr/local"
```

***

#### ディレクトリの作成

* `mkdir`メソッド：指定したパスのディレクトリを作成する。2つ目の引数にパーミッション(mode)を指定可能

* 通常、パーミッションは3桁の8進数で指定。実際のパーミッションは、指定された値と`unmask`をかけた値(`mode & ~unmask`)となる

パーミッションがよくわからないので、省略

2018/9/1o

```ruby
>> Dir.mkdir("/tmp/foo")
=> 0
>> Dir.mkdir("/tmp/bar", 0755)
=> 0
```

***

#### ディレクトリの削除

* `rmdir`メソッド：ディレクトリを削除する

```ruby
>> Dir.mkdir("/tmp/foo")
=> 0
>> Dir.rmdir("/tmp/foo")
=> 0
```

***

### 5-10-2.`File`クラス

#### ファイルを開いて読み込む

* `File.open`、`File.new`メソッド：ファイルを開く。

* 引数としてファイル名だけを与えると、読み取りモードで開く。

* ファイルが存在しない場合は、エラーが発生する。

* ファイいるを開くとファイルオブジェクトが返り、

  * `read`メソッド：ファイルの内容を取得

  * `close`メソッド：ファイルを閉じる

* ファイルの入出力時には、エンコーディングが有効になる

* `File.open`メソッドにブロックを与えると、ブロック終了時に自動的にファイルを閉じることができる
  →ファイルの閉じ忘れを防ぐ為にも、通常はこの形式で使う

```ruby
>> file = File.open("README.md")
=> #<File:README.md>
>> file.read
=> "# README\n\nThis README would normally document whatever steps are necessary to get the\napplication up and running.\n\nThings you may want to cover:\n\n* Ruby version\n\n* System dependencies\n\n* Configuration\n\n* Database creation\n\n* Database initialization\n\n* How to run the test suite\n\n* Services (job queues, cache servers, search engines, etc.)\n\n* Deployment instructions\n\n* ...\n"
>> file.close
=> nil

# 入出力時のエンコーディング
>> Encoding.default_external
=> #<Encoding:UTF-8>
>> file = File.open("README.md")
=> #<File:README.md>
>> file.read
=> "# README\n\nThis README would normally document whatever steps are necessary to get the\napplication up and running.\n\nThings you may want to cover:\n\n* Ruby version\n\n* System dependencies\n\n* Configuration\n\n* Database creation\n\n* Database initialization\n\n* How to run the test suite\n\n* Services (job queues, cache servers, search engines, etc.)\n\n* Deployment instructions\n\n* ...\n"
>> file.read.encoding
=> #<Encoding:UTF-8>

# ファイルをブロックで開く
>> File.open("README.md"){|file| file.read}
=> "# README\n\nThis README would normally document whatever steps are necessary to get the\napplication up and running.\n\nThings you may want to cover:\n\n* Ruby version\n\n* System dependencies\n\n* Configuration\n\n* Database creation\n\n* Database initialization\n\n* How to run the test suite\n\n* Services (job queues, cache servers, search engines, etc.)\n\n* Deployment instructions\n\n* ...\n"
```

***

#### ファイルのモード

* `File.open`メソッドの2番目の引数は、ファイルを開くモードを指定できる

  * `"r"`：読み込みモード

  * `"w"`：書き込みモード。既存ファイルの場合は、ファイルの内容を空にする

  * `"a"`：追記モード。常にファイルの末尾に追加される

  * `"r+"`：読み書きモード。ファイルの読み書き位置が先頭になる

  * `"w+"`：読み書きモード。`"r+"`と同じだが、既存ファイルの場合はファイルの内容が空になる

  * `"a+"`：読み書きモード。ファイルの読み込み位置は先頭に、書き込み位置は常に末尾になる

***

#### ファイルのエンコーディング

* モードの後ろに、 **ファイルのエンコーディング(外部エンコーディング)** と **読み込んだ時のエンコーディング(内部エンコーディング)** を指定可能

* 書き込むことも同時に指定可能

```ruby
# ファイルのエンコーディング指定(読み込み)
>> f = File.open('shift_jis.txt', 'r:shift_jis:utf-8')
=> #<File:shift_jis.txt>
>> f.read
=> "ルビー" # utf-8に変換されている

# ファイルのエンコーディング指定(読み込み・書き込み)
>> f = File.open('shift_jis.txt', 'w+:shift_jis:utf-8')  # ファイルのエンコーディングをShift_JISにする
=> #<File:shift_jis.txt>
>> f.write 'ルビー'.encode('euc-jp')                      # 書き込む内容をEUC-JPにする
=> 6
>> f.rewind
=> 0
>> f.read(4)
=> "\x83\x8B\x83r"
```

***

#### ファイルに書き込む

* `write`メソッド：ファイルに文字を書き込む。ファイルオブジェクトのメソッド

* `IO`クラスに、他のメソッドの記述あり

```ruby
>> File.open('new-file', "w") {|file| file.write "This is new file."}
=> 17
```

***

#### ファイルの属性を取得する

**ファイルの属性を取得するメソッド**

* `File.basename`：指定されたパスからファイル名を取得する

* `File.dirname`：指定されたパスからディレクトリ名を取得する

* `File.extname`：指定されたパスから拡張子を取得する

* `File.split`：指定されたパスからディレクトリ名とファイル名の配列を取得する

* `File.stat`、`File.lstat`：ファイルの属性を示す`File::Stat`クラスのオブジェクトを返す

* `File.atime`、`File.ctime`、`File.mtime`：それぞれのファイルの最終アクセス時刻、状態が変更された時刻、最終更新時刻を取得する

**ファイルオブジェクトのメソッドによる取得**

* `path`：ファイルを開くときに使用したパスを返す

* `lstat`：ファイルの属性を示す`File::Stat`クラスのオブジェクトを返す

* `actime`、`ctime`、`mtime`：それぞれのファイルの最終アクセス時刻、状態が変更された時刻、最終更新時刻を取得する

```ruby
>> File.mtime('README.md')
=> 2018-04-22 22:19:21 +0900
>> File.open('README.md') {|file| file.mtime}
=> 2018-04-22 22:19:21 +0900
```

***

#### ファイルをテストする(`Filetest`モジュール)

ファイルの存在確認や、ディレクトリかどうかの判定など、ファイルをテストするメソッド
→`FileTest`モジュールのメソッド

* `File.exist?`：指定されたパスが存在しているかを調べる

* `File.file?`、`File.directory?`、`File.symlink?`：それぞれ指定されたパスがファイルか、ディレクトリか、シンボリックリンクかを調べる

* `File.executable?`、`File.readable?`、`File.writable?`：それぞれ指定されたファイルが実行可能か、読み取り可能か、書き込み可能かを調べる

* `File.size`：指定されたファイルのサイズを返す

```ruby
>> File.directory?('/usr/local')
=> true
>> File.directory?('/usr/local/bin/zsh')
=> false
```

***

#### ファイルの属性を設定する

* `File.chmod`メソッド：ファイルの属性を変更する

* `File.chown`メソッド：ファイルの所有者を変更する

* `File.utime`メソッド：ファイルの最終アクセス時刻や更新時刻を設定する

```ruby
# File.chmod、File.chownメソッド
>> File.chmod(0644, 'README.md')
=> 1
>> File.chown(501, 20, 'README.md')
=> 1

# File.utimeメソッド
>> File.utime(Time.now, Time.now, 'README.md')
=> 1
```

***

#### ファイルのパスを絶対パスに展開する

* `File.expand_path`メソッド：指定されたパスを絶対パスに展開する

```ruby
>> File.expand_path('README.md')
=> "/Users/MacUser/work/rails/shared_hobby/README.md"
```

***

#### ファイルを削除する、リネームする

* `delete`、`unlink`メソッド：指定されたファイルを削除する。削除に失敗した場合は、エラーが発生する

* `truncate`メソッド：ファイルを指定したバイト数に切り詰める。

* `rename`メソッド：1つ目の引数で指定したファイル名を、2つ目の引数で指定したファイル名に変更する。リネーム先のファイルが存在する場合は、ファイルを上書きする

```ruby
# deleteメソッド
>> File.delete('README.md')
=> 1

# truncateメソッド
>> File.truncate('README.md', 0)
=> 0
>> File.open('README.md', "w") {|file| file.truncate(0)}
=> 0

# renameメソッド
>> File.rename('README.md', 'READ_ME.md')
=> 0
```

***

#### ファイルをロックする

* `flock`メソッド：ファイルをロックする。引数にはロック方法を指定する。

```ruby
>> File.open('README.md', "w") {|file| file.flock(File::LOCK_EX)}
=> 0
```

***

### 5-10-3.`IO`クラス

* `File`クラスのスーパークラスであり、基本的な入出力機能を備えたクラス

* 多くのメソッドは`File`クラスでも利用できる

* 標準出力(`STDOUT`)、標準入力(`STDIN`)、標準エラー出力(`STDERR`)は`IO`クラスのオブジェクト

***

#### `IO`を開く

* ファイルを開くには、`Kernel`モジュールの`open`メソッドを使用

* ファイル名とファイルを開く時のモードを指定して`open`メソッドを実行すると、`File`オブジェクトが返る

```ruby
>> io = open('README.md')
=> #<File:README.md>

# エンコーディングを指定してファイルを開く
>> io = open('README.md', 'w+:shift_jis:euc-jp')  # w+：読み書きモード。外部エンコーディング：shift_jis、内部エンコーディング：euc-jp
=> #<File:README.md>
```

* `open`メソッドで、ファイル名の代わりに、`|`に続いてコマンドを指定すると、コマンドの出力結果を得ることができる
  →`IO`オブジェクトが返る

```ruby
>> io = open('| ls -la')
=> #<IO:fd 11>
```

* `open`メソッドで、開いたファイルの内容を読み込む。エンコーディングが未指定の場合は、`Encoding.default_external`で指定されたものになる

```ruby
>>io = open('| ls -la README.md')
=> #<IO:fd 13>
>> puts io.read
-rw-r--r--  1 MacUser  staff  0  9 13 22:35 README.md
=> nil
>> io.read.encoding
=> #<Encoding:UTF-8>
```

* `write`メソッドで、開いたファイルに書き込む

```ruby
>> STDOUT.write('There is new technology.')
There is new technology.=> 24
```

* `close`メソッドで、ファイルを閉じる。ただし、ファイルを開く`open`メソッドでブロックを渡している場合は、ブロック終了時に自動的にファイルが閉じられる。

```ruby
>> open('README.md'){|io| puts io.read}

=> nil
```

* `IO.popen`メソッドで、コマンドをサブプロセスとして実行し、そのプロセスと入出力のパイプを開くことができる

* `close_write`メソッドは、`IO`オブジェクトの書き込み用の`IO`を閉じるメソッド
  読み込み用の`IO`を閉じるメソッドは、`close_read`メソッドとなる

```ruby
>> IO.popen('grep -i ruby', 'r+') do |io|
?> io.write('This is Ruby program')
>> io.close_write
>> puts io.read
>> end
This is Ruby program
=> nil
```

***

#### `IO`からの入力


* `IO.read`、`read`：`IO`から内容を読み込む。長さが指定されていれば、その長さだけ読み込む。長さを指定した場合のみ、バイナリ読み込みとなり、エンコーディングが **ASCII-8BIT** となる

* `IO.foreach`、`each`、`each_lines`：指定されたファイルを開き、各行をブロックに渡して実行する

* `readlines`：ファイルを全て読み込んで、その各行の配列を返す

* `readline`、`gets`：`IO`オブジェクトから1行読み込む時に用いる

* `each_byte`：与えられたブロックに`IO`オブジェクトから1バイトずつ整数として読み込んで渡していく

* `getbyte`、`readbyte`：`IO`オブジェクトから1バイト読み込んで整数として返す

* `each_char`：与えられたブロックに`IO`オブジェクトから1文字ずつ読み込んで渡していく

* `getc`、`readchar`：`IO`オブジェクトから1文字読み込む。その文字に対応する文字列を返す

```ruby
# IO.readメソッド
>> IO.read("README.md", 5)
=> "# REA"
>> IO.read("README.md", 5).encoding
=> #<Encoding:ASCII-8BIT>

# IO.foreachメソッド
>> IO.foreach("README.md"){|line| puts line}
# README

This README would normally document whatever steps are necessary to get the
application up and running.
# 中略
=> nil

# readlinesメソッド
>> open("README.md").readlines
=> ["# README\n", "\n", # ・・・中略
]

# getsメソッド
>> io = open("README.md")
=> #<File:README.md>
>> io.gets
=> "# README\n"
>> io.gets
=> "\n"

# each_byteメソッド
>> io = open("README.md")
=> #<File:README.md>
>> io.each_byte{|i| puts i}
35
32
82
69
# ・・・省略
=> #<File:README.md>

# getbyteメソッド
>> io = open("README.md")
=> #<File:README.md>
>> io.getbyte
=> 35
>> io.getbyte
=> 32

# each_charメソッド
>> io.each_char{|c| puts c }
#

R
E
# ・・・省略
=> #<File:README.md>

# getcメソッド
>> io = open("README.md")
=> #<File:README.md>
>> io.getc
=> "#"
>> io.getc
=> " "
```

***

#### 空ファイルや`EOF`になった時の振る舞い

* `IO.read`：空ファイルの場合は`""`が返る。読み込む長さが指定されている場合には`nil`が返る

* `IO.readlines`：空ファイルの場合は、空配列`[]`が返る

* `IO.foreach`：ブロックが実行されない

* `each`、`each_byte`：`EOF`であれば何もしない

* `getc`、`gets`：`nil`が返る

* `read`：長さが指定されていない場合は`""`、指定されている場合は`nil`が返る

* `readchar`、`readline`：`EOF`Errorエラーが発生する

* `readlines`：空配列`[]`が返る

* `getbyte`：`nil`が返る

* `readbyte`：`EOF`Errorエラーが発生する

***

#### `IO`への出力

* `write`：`IO`に対して引数の文字列を出力する。引数が文字列以外の場合は、`to_s`メソッドで文字列化して出力
  →出力が成功すると、出力した文字列のバイト数を返す

* `puts`：`IO`に対して複数のオブジェクトを出力する。引数が文字列や配列でない場合、`to_ary`メソッドにより配列化し、次に各要素を`to_s`メソッドにより文字列化して出力する

* `print`：`IO`に対して複数のオブジェクトを出力する。`puts`メソッドと異なり、複数のオブジェクトが指定されると、各オブジェクトの間に`$,`の値を出力する。`$\`に値が設定されていれば最後に出力する。引数が文字列でない場合には、`to_s`メソッドで文字列化して出力する

* `printf`：指定されたフォーマットに従って引数の値を出力する。

* `putc`：`IO`に引数の文字を出力する。

  * 引数が整数の場合は、その最下位バイトを文字コードとする文字

  * 引数が文字列の場合は、先頭の1文字を出力する

  * どちらでもない場合は、`to_int`メソッドで整数化して出力する

* `<<`：`IO`に指定されたオブジェクトを出力する。返り値が`IO`オブジェクト自身となるため、メソッドチェーンを用いることができる

```ruby
# writeメソッド
>> STDOUT.write('There is new technology.')
There is new technology.=> 24

# putsメソッド
>> STDOUT.puts('Abcdefg', 'Hijklmn')
Abcdefg
Hijklmn
=> nil

# printメソッド
>> $, = "\n"
=> "\n"
>> STDOUT.print('This is first line.', 'This is second line.')
This is first line.
This is second line.=> nil

# printfメソッド
>> STDOUT.printf('%010d', 123456)
0000123456=> nil

# <<メソッド
>> STDOUT << "This" << " " << "is" << " " << "README" << "."
This is README.=> #<IO:<STDOUT>>
```

* `flush`：`IO`の内部バッファをフラッシュして出力する

* Rubyでは、通常`IO`への出力は一旦内部バッファに蓄積されるため、`write`メソッドや`puts`メソッドを実行してもすぐにはファイルに書き込まれない

```ruby
>> io = open('README.md', 'w+')
=> #<File:README.md>
>> io.write('This is new README.md')
=> 21
>> `cat README.md`
=> ""
>> io.flush                # この時に初めて出力される
=> #<File:README.md>
>> `cat README.md`
=> "This is new README.md"
```

***

#### `IO`オブジェクトの状態を調べる

* `stat`：`IO`オブジェクトの状態を表す`File::Stat`オブジェクトを返す

* `closed?`：`IO`オブジェクトが閉じられたかどうかを調べる

* `eof?`：ファイルの終端に到達したかどうかを調べる

* `lineno`：現在の行番号(getsメソッドが呼び出された回数)を調べる
  `lineno=`メソッドで設定することも可能

* `sync`：出力する際のバッファのモードを調べる。返り値が`true`の場合には、出力メソッドの実行毎にバッファがフラッシュされる

```ruby
# statメソッド
>> io = open('README.md', 'w+')
=> #<File:README.md>
>> io.stat
=> #<File::Stat dev=0x1000004, ino=8606215164, mode=0100644, nlink=1, uid=501, gid=20, rdev=0x0, size=0, blksize=4194304, blocks=0, atime=2018-09-15 11:41:45 +0900, mtime=2018-09-15 11:41:44 +0900, ctime=2018-09-15 11:41:44 +0900, birthtime=2018-09-11 21:43:26 +0900>

# eof?、closed?メソッド
>> io = open('README.md', 'r+')
=> #<File:README.md>
>> io.read
=> ""
>> io.eof?
=> true
>> io.close
=> nil
>> io.closed?
=> true

# linenoメソッド
>> io = open('README.md')
=> #<File:README.md>
>> io.read
=> "# README\n\nThis README would normally" #省略済み
>> io.rewind
=> 0
>> io.gets
=> "# README\n"
>> io.lineno
=> 1
>> io.lineno = 10
=> 10
>> io.gets
=> "\n"
>> io.lineno
=> 11

# syncメソッド
>> io = open('README.md')
=> #<File:README.md>
>> io.sync                   # openされているだけなので、false
=> false
```

***

#### ファイルポインタの移動や設定

* `rewind`：ファイルポインタを先頭に移動し、`lineno`の値を`0`に設定

* `pos`：ファイルポインタの位置の取得、設定をする

* `seek`：指定した数だけファイルポインタを、2番目の引数の位置から移動する

  * `IO::SEEK_SET`：ファイルの先頭からの位置を表す定数(デフォルト)

  * `IO::SEEK_CUR`：現在のファイルのポインタの位置からを表す

  * `IO::SEEK_END`：ファイルの末尾からを表す

  を指定できる

```ruby
# 共通
>> io = open('README.md')
=> #<File:README.md>

# rewindメソッド
>> io.read
=> "# README\n\nThis README would normally document " # 省略済み
>> io.read
=> ""
>> io.rewind # 先頭に戻る
=> 0
>> io.read
=> "# README\n\nThis README would normally document " # 省略済み

# posメソッド
>> io.pos
=> 374
>> io.pos = 15
=> 15
>> io.read
=> "README would normally document " # 省略済み

# seekメソッド
>> io.seek(10)
=> 0
>> io.read
=> "This README would normally document " # 省略済み
>> io.seek(-10, IO::SEEK_END)
=> 0
>> io.read
=> "ns\n\n* ...\n"
```

***

## 5-11.`Time`クラス

* 時刻を表すクラス

* 時刻は、世界標準時の1970年1月1日午前0時(起算時)からの経過秒数で保持される

* タイムゾーンとして、 **UTC** か **地方時刻(ローカルタイム)** を指定することができる

### 5-11-1.`Time`オブジェクトの生成

* `Time.new`、`Time.now`：現在時刻の`Time`オブジェクトを生成して返す。タイムゾーンはローカルタイム(日本なら、`+0900`)

* `Time.at`：引数で指定した起算時からの秒数に対応する`Time`オブジェクトを生成して返す。さらに精度が必要な場合には、2番目の引数にマイクロ秒を指定

* `Time.mktime`、`Time.local`：与えられた引数に対応する`Time`オブジェクトを生成して返す

  * 引数が7個までの場合、先頭から`年`、`月`、`日`、`時`、`分`、`秒`、`マイクロ秒`を指定できる

  * 年のみ省略できない。他の引数を省略した場合は、`1`、`1`、`0`、`0`、`0`、`0`が指定されたとみなす

  * 月に関しては、`"Jan"`や`"Feb"`などの英語の月名の省略形も指定できる

  * 引数が10個の場合は、先頭から`秒`、`分`、`時`、`日`、`月`、`年`、`曜日に対応する数値`、`年日`、`夏時間かどうかの真理値`、`タイムゾーン`を指定できる

  * ただし、`曜日に対応する数値`、`年日`、`夏時間かどうかの真理値`、`タイムゾーン`に関しては無視される

* `Time.gm`、`Time.utc`：引数の数やその順序は`Time.mktime`と同じだが、生成される`Time`オブジェクトのタイムゾーンが **UTC** になる

```ruby
# Time.nowメソッド
>> Time.now
=> 2018-09-15 12:08:25 +0900

# Time.atメソッド
>> Time.at(1234567890)
=> 2009-02-14 08:31:30 +0900
>> Time.at(1234567890, 1234567890)  # マイクロ秒指定
=> 2009-02-14 08:52:04 +0900

# Time.mktimeメソッド(引数7個)
>> Time.mktime(2017)          # 年のみ指定
=> 2017-01-01 00:00:00 +0900
>> Time.mktime(2017, 7, 7)    # 年、月、日のみ指定
=> 2017-07-07 00:00:00 +0900

# Time.mktimeメソッド(引数10個)
>> Time.mktime(0, 0, 0, 7, 7, 2017, 4, 188, false, "JST")
=> 2017-07-07 00:00:00 +0900

# Time.gmメソッド
>> Time.gm(2017)
=> 2017-01-01 00:00:00 UTC
>> Time.gm(2017, 7, 7)
=> 2017-07-07 00:00:00 UTC
```

***

### 5-11-2.`Time`オブジェクトの属性を取得する

* `year`：年を取得

* `mon`、`month`：月を取得

* `day`、`mday`：日を取得

* `hour`：時を取得

* `min`：分を取得

* `sec`：秒を取得

* `usec`、`tv_usec`：マイクロ秒を取得

```ruby
>> t = Time.mktime(2017, 1, 2, 3, 4, 5, 6)
=> 2017-01-02 03:04:05 +0900
>> t.year     # 年を取得
=> 2017
>> t.mday     # 日を取得
=> 2
>> t.sec      # 秒を取得
=> 5
```

* `wday`：曜日に対応する数値を返す。日曜(0)〜土曜(6)が対応

* `yday`：1月1日からの日数を返す

```ruby
>> t = Time.mktime(2017, 1, 2, 3, 4, 5, 6)
=> 2017-01-02 03:04:05 +0900
>> t.wday      # 月曜日
=> 1
>> t.yday      # 1月1日から2日目
=> 2
```

* `isdst`、`dst?`：夏時間かどうかを返す

* `gmt?`、`utc?`：タイムゾーンが **UTC** かどうかを返す

```ruby
>> t = Time.mktime(2017, 1, 2, 3, 4, 5, 6)
=> 2017-01-02 03:04:05 +0900
>> t.dst?
=> false
>> t.gmt?
=> false
```

* `gmtoff`、`gmt_offset`：UTC時刻との差を秒単位の数値として返す

```ruby
>> t = Time.mktime(2017, 1, 2, 3, 4, 5, 6)
=> 2017-01-02 03:04:05 +0900
>> t.gmtoff
=> 32400
>> t.gmtoff / 3600  # 時に変更
=> 9
```

***

### 5-11-3.タイムゾーンを変更する

* `localtime`：タイムゾーンをローカルタイムに変更

* `gmtime`、`utc`：タイムゾーンをUTCに変更

```ruby
>> t = Time.mktime(2017, 1, 2, 3, 4, 5, 6)
=> 2017-01-02 03:04:05 +0900
>> t.localtime
=> 2017-01-02 03:04:05 +0900
>> t.gmtime
=> 2017-01-01 18:04:05 UTC
```

* `getlocal`：タイムゾーンをローカルタイムに変更した新しい`Time`オブジェクトを返す

* `getgm`、`getutc`：タイムゾーンをUTCに変更した新しい`Time`オブジェクトを返す

```ruby
>> t = Time.mktime(2017, 1, 2, 3, 4, 5, 6)
=> 2017-01-02 03:04:05 +0900
>> t.object_id
=> 70329225673020
>> t1 = t.getlocal
=> 2017-01-02 03:04:05 +0900
>> t1.object_id
=> 70329225651440
```

***

### 5-11-4.`Time`オブジェクトの演算

* `+`：指定した秒数後の`Time`オブジェクトを返す

* `-`：指定した秒数前の`Time`オブジェクトを返す

```ruby
>> t = Time.mktime(2017, 1, 2, 3, 4, 5, 6)
=> 2017-01-02 03:04:05 +0900
>> t + 60 * 60 * 3
=> 2017-01-02 06:04:05 +0900
>> t - 60 * 60 * 3
=> 2017-01-02 00:04:05 +0900

# Timeオブジェクト同士の差
>> t1 = Time.mktime(2011, 1, 2, 3, 4, 5, 6)
=> 2011-01-02 03:04:05 +0900
>> t2 = Time.mktime(2011, 2, 3, 4, 5, 6, 7)
=> 2011-02-03 04:05:06 +0900
>> t2 - t1
=> 2768461.000001
```

***

### 5-11-5.`Time`オブジェクトの変換

* `to_i`、`tv_sec`：起算時からの秒数を整数で返す

* `to_f`：起算時からの秒数を浮動小数点数で返す

```ruby
>> t = Time.mktime(2017, 1, 2, 3, 4, 5, 6)
=> 2017-01-02 03:04:05 +0900
>> t.to_i
=> 1483293845
>> t.to_f
=> 1483293845.000006
```

* `to_a`：先頭から、秒、分、時、日、月、年、曜日に対応する数値、年日、夏時間かどうかの真理値、タイムゾーン順の配列を返す

```ruby
>> t = Time.mktime(2017, 1, 2, 3, 4, 5, 6)
=> 2017-01-02 03:04:05 +0900
>> t.to_a
=> [5, 4, 3, 2, 1, 2017, 1, 2, false, "JST"]
```

* `to_s`：UNIXコマンドのdateコマンドのような形式の文字列を返す

```ruby
>> t = Time.mktime(2017, 1, 2, 3, 4, 5, 6)
=> 2017-01-02 03:04:05 +0900
>> t.to_s
=> "2017-01-02 03:04:05 +0900"
```

* `strftime`：指定したフォーマットに従って文字列に変換した結果を返す

```ruby
>> t = Time.mktime(2017, 1, 2, 3, 4, 5, 6)
=> 2017-01-02 03:04:05 +0900
>> t.strftime("%Y年%m月%d日 %H時%M分%S秒")
=> "2017年01月02日 03時04分05秒"
```

この表は、後に画像として貼る
2018/9/15

***

## 5-12.`Regexp`クラス

* 正規表現オブジェクトを扱うクラス

* 正規表現を使って文字列やデータのマッチングを行うときに、使用

### 5-12-1.正規表現オブジェクトを生成

* 正規表現は、正規表現リテラルを使って表現する。リテラルの末尾には、オプションが指定できる。

* オプションは、

  * `i`：大文字小文字の違いを無視

  * `m`：正規表現の`.`で改行にマッチさせる

  * `x`：空白や`#`から始まるコメントを無視する

  などがある。また、オプションの複数指定もできる

```ruby
>> a = /abcdefg/i  # 大文字小文字の違いを無視する
=> /abcdefg/i
>> a.class
=> Regexp
```

* `Regexp.new`、`Regexp.compile`：正規表現オブジェクトを生成する。2つ目の引数に、マッチングのオプションを指定できる。

  * `Regexp::IGNORECASE`：大文字小文字の違いを無視する

  * `Regexp::MULTILINE`：正規表現の`.`が改行にマッチするようになる

  * `Regexp::EXTENDED`：バックスラッシュでエスケープされていない空白と、`#`から改行までを無視する

  * 論理和を使って複数指定も可能

  * マッチングするときの文字コードを、3番目の引数で指定することも可能

```ruby
>> a = Regexp.new("abcdefg", Regexp::MULTILINE | Regexp::IGNORECASE)
=> /abcdefg/mi
```

***

### 5-12-2.正規表現オブジェクトでマッチングする

* `match`：正規表現オブジェクトで文字列とマッチングさせる。マッチした場合には`MatchData`オブジェクトを、しなかった場合には`nil`を返す

* `=~`：正規表現オブジェクトで文字列とマッチングさせる。マッチすればマッチした位置のインデックスが、しなかった場合は`nil`を返す

* `===`：正規表現オブジェクトで文字列とマッチングさせる。マッチすれば`true`、しなかった場合は`false`が返る

* `~`：特殊変数`$_`とマッチングする

```ruby
# matchメソッド
>> a = Regexp.new("abc")
=> /abc/
>> a.match("abcdef")
=> #<MatchData "abc">

# =~メソッド
>> a = Regexp.new("abc")
=> /abc/
>> a =~ "abcdef"          # 0番目の文字がマッチした
=> 0
>> "abcdefg" =~ a         # 0番目の文字がマッチした
=> 0

# ===メソッド
>> a = Regexp.new("abc")
=> /abc/
>> a === "abcdef"
=> true

# ~メソッド
>> $_ = "abcdefg"
=> "abcdefg"
>> a = Regexp.new("abc")
=> /abc/
>> ~ a
=> 0
```

***

### 5-12-3.正規表現の特殊文字をエスケープする

* `Regexp.escape`、`Regexp.quote`：ピリオド`.`、カッコ`[]`などでマッチングする際に、これらの文字を自動的にエスケープする

```ruby
>> Regexp.escape("array.push(hash[key])")
=> "array\\.push\\(hash\\[key\\]\\)"
```

***

### 5-12-4.マッチした結果を取得する

* `Regexp.last_match`：正規表現でマッチした結果を取得。`MatchData`オブジェクト(現在のスコープ(トップレベルやクラス・モジュール・メソッド定義)の中で最後に行った正規表現のマッチ結果)を返す。特殊変数`$_`でも取得できる

```ruby
>> /abcdefg/ =~ "abcdefghijklmnopqrstuvwxyz"
=> 0
>> Regexp.last_match
=> #<MatchData "abcdefg">
>> $~
=> #<MatchData "abcdefg">
```

* `Regexp.last_match`メソッドに整数値を与えると、該当のマッチ文字列が得られる。

  * `0`であれば正規表現にマッチした文字列

  * それ以降の整数では、カッコにマッチした部分文字列が得られる。これらの文字列はそれぞれ特殊変数`$&`、`$1`、`$2`などでも取得可能

```ruby
>> /(abc)d(efg)/ =~ "abcdefghijklmnopqrstuvwxyz"
=> 0
>> Regexp.last_match(0)
=> "abcdefg"
>> $&
=> "abcdefg"
>> Regexp.last_match(1)
=> "abc"
>> $1
=> "abc"
>> $2
=> "efg"
```

***

### 5-12-5.正規表現の論理和を求める

* `Regexp.union`：複数の正規表現を結合し、そのどれかにマッチするような新しい正規表現を求める

```ruby
>> a = Regexp.new("abc")
=> /abc/
>> b = Regexp.new("ABC")
=> /ABC/
>> c = Regexp.union(a, b)
=> /(?-mix:abc)|(?-mix:ABC)/
>> c =~ "abc"
=> 0
>> Regexp.last_match
=> #<MatchData "abc">
```

***

### 5-12-6.正規表現オブジェクトのオプションや属性を取得する

* `options`：正規表現オブジェクトを生成する時に設定したオプションである`Regexp::IGNORECASE`、`Regexp::MULTILINE`、`Regexp::EXTENDED`の論理和を返す

* `casefold?`：オプション`Regexp::IGNORECASE`が設定してあるかどうかを返す

* `encoding`：正規表現オブジェクトがコンパイルされている文字コードを`Encoding`オブジェクトとして返す

* `source`：正規表現の元となった文字列表現を返す

  * `to_s`：他の正規表現に埋め込んでも元の意味が保たれるような形式

  * `inspect`：`to_s`メソッドよりも自然な形式な文字列になるが、元の意味は保たれない

```ruby
# optionsメソッド
>> a = Regexp.new("abcdefg", Regexp::MULTILINE | Regexp::IGNORECASE)
=> /abcdefg/mi
>> a.options
=> 5
>> a.options & Regexp::IGNORECASE
=> 1
>> a.options & Regexp::EXTENDED
=> 0

# casefold?メソッド
>> a = Regexp.new("abcdefg")
=> /abcdefg/
>> a.casefold?
=> false
>> a = Regexp.new("abcdefg", Regexp::MULTILINE | Regexp::IGNORECASE)
=> /abcdefg/mi
>> a.casefold?
=> true

# encodingメソッド
>> a = Regexp.new("ルビー")
=> /ルビー/
>> a.encoding
=> #<Encoding:UTF-8>
>> a = Regexp.new("ルビー".encode("EUC-JP"))
=> /\x{A5EB}\x{A5D3}\x{A1BC}/
>> a.encoding
=> #<Encoding:EUC-JP>

# sourceメソッド
>> a = Regexp.new("abcdefg", Regexp::MULTILINE | Regexp::IGNORECASE)
=> /abcdefg/mi
>> a.source
=> "abcdefg"
>> a.to_s
=> "(?mi-x:abcdefg)"
>> a.inspect
=> "/abcdefg/mi"
```

***

## 5-13.`Exception`クラス

* 全ての例外クラスのスーパークラス。エラーが発生した場合や、`raise`メソッドで例外を発生した時に、このクラスのオブジェクトが生成される

#### 例外クラスの自作

```ruby
>> class MyError < RuntimeError; end
=> nil
>> begin
?>   raise MyError
>> rescue => ex
>>   p ex
>> end
#<MyError: MyError>
=> #<MyError: MyError>
```

***

#### エラーメッセージを指定する

```ruby
>> class MyError < RuntimeError; end
=> nil
>> begin
>>   raise MyError.exception('エラーが発生しました。')
>> rescue => ex
>>   p ex
>> end
#<MyError: エラーが発生しました。>
=> #<MyError: エラーが発生しました。>
```

***

#### エラーメッセージを取得する

* `message`、`to_s`、`to_str`：例外オブジェクトに設定されているエラーメッセージを取得

* オブジェクトにエラーメッセージを設定するには、`new`、`exception`メソッドの呼び出し時に引数で指定

```ruby
>> class MyError < RuntimeError; end
=> nil
>> begin
?>   raise MyError.exception('エラーが発生しました。')
>> rescue => ex
>>   p ex.message
>> end
"エラーが発生しました。"
=> "エラーが発生しました。"
```

***

#### バックトレースを取得

* `backtrace`：発生した例外のバックトレース情報を取得。これらは、配列で返る。

```ruby
>> class MyError < RuntimeError; end
=> nil
>> begin
?> raise MyError.exception('エラーが発生しました。')
>> rescue => ex
>> p ex.backtrace
>> end
["(irb):41:in `irb_binding'", "/Users/MacUser/.rbenv/versions/2.4.1/lib/ruby/2.4.0/irb/workspace.rb:87:in `eval'",] # 省略
=> ["(irb):41:in `irb_binding'", "/Users/MacUser/.rbenv/versions/2.4.1/lib/ruby/2.4.0/irb/workspace.rb:87:in `eval'",] # 省略
>>
```

***

#### バックトレースを取得(オリジナルの情報追加)

* `set_backtrace`：バックトレース情報にオリジナルの情報を設定する。ただし、それまでのバックトレース情報は上書きされることに注意。

```ruby
>> class MyError < RuntimeError; end
=> nil
>> begin
?>   raise MyError.exception('エラーが発生しました。')
>> rescue => ex
>>   ex.set_backtrace("This is new backtrace.")
>>   p ex.backtrace
>> end
["This is new backtrace."]
=> ["This is new backtrace."]
```

***

## 5-14.`Proc`クラス

* ブロックを実行時のローカル変数のスコープなどのコンテキストと共にオブジェクト化した、手続きオブジェクトを扱うクラス

* この手続きオブジェクトは、名前のない関数(無名関数)のように使うことができる

* `call`：この手続きを実行する

```ruby
>> f = Proc.new { puts 'OK' }
=> #<Proc:0x007ff5ce926b40@(irb):58>
>> f.call
OK
=> nil
```

* `arity`：生成された手続きオブジェクトの引数の数を取得

```ruby
>> f = Proc.new {|str| puts str }
=> #<Proc:0x007ff5ce85e118@(irb):60>
>> f.arity
=> 1
>> f.call('NG')
NG
=> nil
```

* オブジェクト生成時のコンテキストを保持しているため、ローカル変数の値などは実行時の状態に応じて変化する

```ruby
>> i = 30
=> 30
>> j = 40
=> 40
>> f = Proc.new { puts i + j }
=> #<Proc:0x007ff5cf09dc48@(irb):66>
>> f.call
70
=> nil
>> i = 100
=> 100
>> f.call
140
=> nil
```

***

### 5-14-1.ブロック付きメソッドの引数として利用する

* ブロック付きメソッドに手続きオブジェクトを渡すこともできる

* 変数の前に`&`を指定して渡す

```ruby
>> f = Proc.new {|i| puts i}
=> #<Proc:0x007ff5cf07f130@(irb):71>
>> 3.times(&f)
0
1
2
=> 3
```

***

### 5-14-2.手続きオブジェクトの中での処理の中断

* `next`：手続きオブジェクトの中で処理を中断して、呼び出し元へ値を戻す

```ruby
>> f = Proc.new {
?>   next "next"   # 中断
>>   "last"
>> }
=> #<Proc:0x007ff5cf06c670@(irb):73>
>> f.call
=> "next"
```

***

### 5-14-3.`Proc.new`以外の手続きオブジェクト生成

* `lambda`、`proc`：Kernelモジュールのメソッド。手続きオブジェクトを生成する。

* いくつかの場面で振る舞いが異なる

#### 手続きオブジェクトにおける引数の数

* `lambda`などでは、`proc`メソッドで生成した手続きオブジェクトでは、引数の数が異なると`ArgumentError`が発生

* `Proc.new`で生成した手続きオブジェクトでは、引数への多重代入のようになるので、エラーが発生しない

```ruby
# Proc.newメソッド
>> f = Proc.new {|a, b, c| p a, b, c}
=> #<Proc:0x007ff5cf0554c0@(irb):78>
>> f.call(1, 9)
1
9
nil
=> [1, 9, nil]

# lambdaメソッド
>> g = lambda {|a, b, c| p a, b, c}
=> #<Proc:0x007ff5cf03e2c0@(irb):80 (lambda)>
>> g.call(1, 9)
ArgumentError: wrong number of arguments (given 2, expected 3)
```

***

#### 手続きオブジェクトの中でのジャンプ構文

* `break`では、

  * `lambda`メソッドで生成した手続きオブジェクトでは、その手続きオブジェクトを抜ける

  * `Proc.new`、`proc`メソッドでは、`LocalJumpError`が発生する

```ruby
# Proc.newメソッド
>> f = Proc.new { break }
=> #<Proc:0x007ff5ce9349c0@(irb):82>
>> f.call
LocalJumpError: break from proc-closure

# lambdaメソッド
>> g = lambda { break }
=> #<Proc:0x007ff5cea78ef8@(irb):90 (lambda)>
>> g.call
=> nil
```

* `return`では、

  * `lambda`メソッドで生成した手続きオブジェクトでは、その手続きオブジェクトを抜ける

  * `Proc.new`、`proc`メソッドでは、その手続きオブジェクトの外側を抜けようとするので、`LocalJumpError`が発生する

```ruby
# Proc.newメソッド
>> f = Proc.new { return }
=> #<Proc:0x007ff5cf02eac8@(irb):92>
>> f.call
LocalJumpError: unexpected return
>> def foo
>>   Proc.new {
?>     return 1
>>   }.call
>>   return 2
>> end
=> :foo
>> foo
=> 1

# lambdaメソッド
>> g = lambda { return }
=> #<Proc:0x007ff5cea24268@(irb):101 (lambda)>
>> g.call
=> nil
```

***

## 5-15.`Module`クラス

* ある機能をひとまとめにしたモジュールのためのクラス

* クラスのクラスである`Class`クラスでは、この`Module`クラスを継承しているため、全てのクラスでこれらの有用なメソッドが利用できる

* モジュールは、`include`を使用して、任意のクラスにインクルードできる

```ruby
# モジュールの定義
>> module MyMethods
>>   def foo
>>     'bar'
>>   end
>> end
=> :foo

# classの定義
>> class MyClass
>>   include MyMethods
>> end
=> MyClass

# MyClassの呼び出し
>> MyClass.new.foo
=> "bar"
```

***

### 5-15-1.定義されている定数に関するメソッド

* `Module.constants`：その時点で定義されている定数を取得する

* `constants`：特定のクラスやモジュールで実行することで、そこで定義されている定数を取得する

* `const_defined?`：指定された定数が定義されているかどうかを調べる

* `const_get`：定義されている定数の値を取り出す

* `const_set`：新たに定数を定義して設定する

* `remove_const`：定義されている定数を取り除く

```ruby
# Module.constantsメソッド
>> Module.constants
=> [:Integer, :Float, :String, :Array,] # 省略済み

# constantsメソッド
>> class MyClass
>>   FOO = 1
>> end
=> 1
>> MyClass.constants
=> [:FOO]

# const_defined?メソッド
>> Object.const_defined?(:ENV)
=> true
>> Object.const_defined?(:ENVIRONMENT)
=> false

# const_getメソッド
>> Object.const_get(:RUBY_VERSION)
=> "2.4.1"

# const_setメソッド
>> Object.const_set(:MY_CONST, 'myconst')
=> "myconst"
>> Object::MY_CONST
=> "myconst"

# 定数を取り除く
>> class MyClass
>>   MYCONST = 1
>>   p remove_const(:MYCONST) # MYCONSTを削除
>>   p MYCONST                # MYCONSTは削除されているので、エラー発生
>> end
1
NameError: uninitialized constant MyClass::MYCONST
```

***

### 5-15-2.メソッドの設定

* `instance_methods`：インスタンスに設定されている`public`、`protected`メソッドの一覧を取得する

* `public_instance_methods`：`public`メソッドの一覧を取得する

* `private_instance_methods`：`private`メソッドの一覧を取得する

* `protected_instance_methods`：`protected`メソッドの一覧を取得する

```ruby
>> Array.instance_methods
=> [:join, :rotate, :rotate!, :sort!, :sort_by!, :collect!, :map!, ] # 省略済み
```

* メソッドの可視性を指定する`public`、`protected`、`private`はメソッド

* 定義済みのメソッドの可視性を後から変更することができる

```ruby
>> class MyClass
>>   private
>>   def foo
>>     puts 'FOO'
>>   end
>>   public :foo
>> end
=> MyClass
>> my_class = MyClass.new
=> #<MyClass:0x007ffe901ff040>
>> my_class.foo
FOO
=> nil
```

* インスタンスの属性として、インスタンス変数と読み取りメソッド、書き込みメソッドを定義するには、

  * `attr_accessor`：読み取りと書き取りメソッド

  * `attr_reader`：読み込みメソッド

  * `attr_writer`：書き込みメソッド

  * `attr`：2番目の引数に`true`を指定すれば読み込み、書き込み両方のメソッド。指定しないか`false`を指定すれば読み込みメソッドのみを定義

```ruby
>> class MyClass
>>   attr_accessor :height
>> end
=> nil
>> my_class = MyClass.new
=> #<MyClass:0x007ffe901bf0a8>
>> my_class.height = 200       # 変数に書き込みをする
=> 200
>> my_class.height             # 変数に値が代入されている
=> 200
```

* `alias_method`：メソッドの別名を定義する。メソッド名を文字列かシンボルで指定できる

* `alias`：予約語であり、直接メソッドを指定できる

```ruby
>> class MyClass
>>   def foo
>>     'foo'
>>   end
>>   alias_method :original_foo, :foo
>>   def foo
>>     'bar' + original_foo
>>   end
>> end
=> :foo
>> m = MyClass.new
=> #<MyClass:0x007ffe90164090>
>> m.foo
=> "barfoo"
```

***

### 5-15-3.評価する

* `eval`：文字列をRubyコードとして評価する。現在のコンテキストで評価する

* `module_eval`、`class_eval`：モジュールやクラスのコンテキストで評価する。メソッドを動的に追加する時などに利用できる

```ruby
>> Array.class_eval do
?>   def foo
>>     'bar'
>>   end
>> end
=> :foo
>> [].foo
=> "bar"
```

* `module_exec`、`class_exec`：モジュールやクラスのコンテキストで評価するときに引数を渡す

```ruby
>> class MyClass
>>   CONST = 1
>> end
=> 1
>> MyClass.class_exec(3) {|i| puts i + self::CONST}
4
=> nil
```

***

### 5-15-4.クラス変数を扱う

* `class_variables`：定義されているクラス変数の一覧を返す

* `class_variables_defined?`：指定されたクラス変数が定義されているかどうかを返す

```ruby
# class_variablesメソッド
>> class MyClass
>>   @@foo = 1
>> end
=> 1
>> MyClass.class_variables
=> [:@@foo]

# class_variables_defined?メソッド
>> class MyClass
>>   @@foo = 1
>> end
=> 1
>> MyClass.class_variable_defined?(:@@foo)
=> true
```

* `class_variables_get`：クラス変数の取得

* `class_variables_set`：クラス変数の設定

* `class_variables_variable`：クラス変数の削除

```ruby
# クラス変数の設定
>> class MyClass
>>   @@var = 'foobar'
>> end
=> "foobar"

# クラス変数の取得、設定、削除のメソッドの設定
>> def MyClass.get
>>   class_variable_get(:@@var)
>> end
=> :get
>> def MyClass.set(var)
>>   class_variable_set(:@@var, var)
>> end
=> :set
>> def MyClass.clear
>>   remove_class_variable(:@@var)
>> end
=> :clear

# 設定したクラス変数の取得→削除→設定→取得
>> MyClass.get
=> "foobar"
>> MyClass.clear
=> "foobar"
>> MyClass.class_variable_defined?(:@@var)
=> false
>> MyClass.set('newvar')
=> "newvar"
>> MyClass.get
=> "newvar"
```

***

### 5-15-5.モジュールの機能を取り込む

* `include`：クラスやモジュール、オブジェクトにモジュールの機能を追加。クラスとそのインスタンスに機能を追加する

* `extend`：`include`メソッドと同じだが、そのオブジェクトのみに機能を追加する

```ruby
# モジュールの定義
>> module MyMethods
>>   def foo
>>     'bar'
>>   end
>> end
=> :foo

# includeメソッドでの取り込み
>> class MyClass
>>   include MyMethods
>> end
=> MyClass
>> MyClass.new.foo
=> "barfoo"

# extendメソッドでの取り込み
>> class NewMyClass; end
=> nil
>> n = NewMyClass.new
=> #<NewMyClass:0x007ffe90225178>
>> n.extend(MyMethods)
=> #<NewMyClass:0x007ffe90225178>
>> n.bar
NoMethodError: undefined method 'bar' for #<NewMyClass:0x007ffe90225178>
>> n.foo
=> "bar"

>> n1 = NewMyClass.new
=> #<NewMyClass:0x007ffe901fed20>
>> n1.foo  # extendしていないので、エラー発生
NoMethodError: undefined method 'foo' for #<NewMyClass:0x007ffe901fed20>
```

* `included`、`extended`：`include`、`extended`メソッドによってそのモジュールの機能がクラスやモジュール、オブジェクトに取り込まれたときに実行されるメソッド

```ruby
# モジュールの定義
>> module MyModule
>>   def self.included(object)
>>     p "#{object} has included #{self}"
>>   end
>> end
=> :included

# includeするクラス
>> class MyClass
>>   include MyModule
>> end
"MyClass has included MyModule"
=> MyClass
```

* `include?`：クラスやモジュールが、指定されたモジュールをインクルードしているかどうかを調べる

* `included_modules`：インクルードしているモジュールの一覧を得る

```ruby
>> module MyModule; end
=> nil
>> class MyClass
>>   include MyModule
>> end
=> MyClass
>> MyClass.include?(MyModule)
=> true
>> MyClass.included_modules
=> [MyModule, Kernel]
```

* `autoload`：未定義の定数が参照されたときに、自動的に特定のファイルをロードするように設定する

* `autoload?`：ファイルがロードされていないときにはそのパス名を、ロードされている場合や指定された定数に`autoload`が指定されていないときには`nil`を返す

```ruby
>> puts open('/Users/MacUser/work/rails/shared_hobby/mymodule.rb').read
module MyModule
  def foo
    puts 'bar'
  end
end
=> nil
>> class MyClass
>>   autoload(:MyModule, "/Users/MacUser/work/rails/shared_hobby/mymodule.rb")
>>   p autoload?(:MyModule) # 1
>>   include MyModule
>>   p autoload?(:MyModule) # 2
>> end
"/Users/MacUser/work/rails/shared_hobby/mymodule.rb"  # 1
nil                                                   # 2
=> nil
```

***

### 5-15-7.祖先クラスを取得する

* `ancestors`：あるクラスの祖先クラスやインクルードしているモジュールの一覧を取得する

```ruby
>> Array.ancestors
=> [Array, Enumerable, Object, Kernel, BasicObject]
```

***

## 5-16.`Enumerable`モジュール

* `Array`、`Hash`クラスにインクルードされている

* 全てのメソッドが`each`メソッドを元に定義されているため、`each`メソッドが定義されているクラスであれば、そのクラスでも利用可能

#### `map`、`collect`メソッド

* 与えられたブロックを評価した結果の配列を返す

```ruby
>> [1, 2, 3, 4, 5].map{|i| i ** 2}
=> [1, 4, 9, 16, 25]
```

#### `each_with_index`メソッド

* 要素とそのインデックスをブロックに渡して繰り返す

```ruby
>> [:a, :b, :c, :d, :e].each_with_index{|v, i| puts "#{v} => #{i}"}
a => 0
b => 1
c => 2
d => 3
e => 4
=> [:a, :b, :c, :d, :e]
```

#### `inject`、`reduce`メソッド

* 自身のたたみこみ演算を行う(初期値と自身の要素を順に組み合わせて結果を返す)

* 引数は、たたみこみを行う際の初期値をとる

```ruby
>> [1, 2, 3, 4, 5].inject(0) {|result, v| result + v ** 2}  # 1から5までの数値の2乗の和を求める(1+4+9+16+25=55)
=> 55
```

#### `each_slice`、`each_cons`メソッド

* `each_slice`：要素を指定された数で区切ってブロックに渡す。要素数が指定された数で割きれない場合は、最後だけ渡される数が少なくなる

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

#### `reverse_each`メソッド

* `each`メソッドとは逆順にブロックに要素を渡して繰り返す

```ruby
>> [1, 2, 3, 4, 5].reverse_each {|i| puts i}
5
4
3
2
1
=> [1, 2, 3, 4, 5]
```

#### `all?`、`any?`、`none?`、`one?`、`member?`、`include?`メソッド

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

#### `find`、`find_index`、`select`などのメソッド

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

#### `sort`、`sort_by`メソッド

* `sort`：要素を`<=>`メソッドで比較して昇順にソートした配列を、新たに生成して返す。ブロックをとる場合は、ブロックの評価結果を元にソートする

* `sort_by`：ブロックの評価結果を`<=>`メソッドで比較して昇順にソートした配列を使って、元の配列をソートした新しい配列を生成して返す

```ruby
>> ["aaa", "b", "cc"].sort{|a, b| a.length <=> b.length}
=> ["b", "cc", "aaa"]
>> ["aaa", "b", "cc"].sort_by{|a| a.length}
=> ["b", "cc", "aaa"]
```

#### `max`、`min`メソッド

* それぞれ要素の最大値と最小値を返す

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

#### `count`メソッド

* 要素数を返す

```ruby
>> [1, 2, 3, 4, 5].count
=> 5
```

#### `cycle`メソッド

* 要素を先頭から順に取り出し、末尾まで到達したら再度先頭に戻り、それを繰り返す

```ruby
>> [:a, :b, :c].cycle{|v| p v}
:a
:b
:c
:a
# 省略
```

#### `group_by`メソッド

* ブロックの評価結果をキーとし、同じキーを持つ要素を配列としたハッシュを返す

```ruby
>> (1..10).group_by{|v| v % 2}
=> {1=>[1, 3, 5, 7, 9], 0=>[2, 4, 6, 8, 10]}
```

#### `zip`メソッド

* 自身と引数に指定した配列から、1つずつ要素を取り出して配列を作り、それを要素とする配列を返す

```ruby
>> [:a, :b, :c].zip([1, 2, 3], ["a", "b", "c"])
=> [[:a, 1, "a"], [:b, 2, "b"], [:c, 3, "c"]]
```

#### `first`、`take`メソッド

* `take`：先頭から指定した数の要素を配列として返す

* `first`：`take`メソッドと同じだが、数を指定しない場合に先頭の要素のみを返す

```ruby
>> [:a, :b, :c].take(2)
=> [:a, :b]
>> [:a, :b, :c].first
=> :a
```

#### `take_while`、`drop`メソッド

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

#### `drop_while`メソッド

* 先頭からブロックを評価し、最初に偽になった要素の手前までを切り捨て、残りの要素を配列として返す

```ruby
>> [:a, :b, :c, :d, :e].drop_while { |e| e != :c}
=> [:c, :d, :e]
```

#### `select`、`reject`メソッド

* `select`：各要素に対してブロックの評価結果が真であった要素を含む配列を返す

* `reject`：ブロックの評価結果が偽であった要素を含む配列を返す

```ruby
>> [1, 2, 3, 4, 5].select { |e| e % 2 == 0 }
=> [2, 4]
>> [1, 2, 3, 4, 5].reject { |e| e % 2 == 0 }
=> [1, 3, 5]
```

#### `lazy`メソッド

* `map`、`select`メソッドなどのメソッドが、遅延評価を行うように再定義される

* 遅延評価になるとそれぞれのメソッドが配列でなく`Enumerator::Lazy`を返すようになるため、メソッドを評価するタイミングを文字通り遅らせることができる

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

***

## 5-17.`Comparable`モジュール

* インクルードしたクラスで比較演算子である`<=>`メソッドを元にオブジェクト同士での比較ができるようになる

* インクルードしたクラスで利用できるインスタンスメソッドは、

  * `<`：負の整数で`true`

  * `<=`：負の整数か0で`true`

  * `==`：0で`true`

  * `>`：正の整数で`true`

  * `>=`：正の整数か0で`true`

  * `between?`：引数`min`と`max`の間にあれば`true`


#### Sampleクラス

* 通常の大小関係と逆の挙動をするクラス

```ruby
>> class Sample
>>   def initialize(value)
>>     @value = value
>>   end
>>
>>   def value
>>     @value
>>   end
>>
>>   def <=>(other)
>>     other.value <=> self.value
>>   end
>> end
=> :<=>
```

***
