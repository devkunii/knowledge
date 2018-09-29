# 5-14.`Proc`クラス

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

## 5-14-1.ブロック付きメソッドの引数として利用する

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

## 5-14-2.手続きオブジェクトの中での処理の中断

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

## 5-14-3.`Proc.new`以外の手続きオブジェクト生成

* `lambda`、`proc`：Kernelモジュールのメソッド。手続きオブジェクトを生成する。

* いくつかの場面で振る舞いが異なる

### 手続きオブジェクトにおける引数の数

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

### 手続きオブジェクトの中でのジャンプ構文

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

# 5-15.`Module`クラス

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

## 5-15-1.定義されている定数に関するメソッド

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

## 5-15-2.メソッドの設定

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

## 5-15-3.評価する

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

## 5-15-4.クラス変数を扱う

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

## 5-15-5.モジュールの機能を取り込む

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

## 5-15-7.祖先クラスを取得する

* `ancestors`：あるクラスの祖先クラスやインクルードしているモジュールの一覧を取得する

```ruby
>> Array.ancestors
=> [Array, Enumerable, Object, Kernel, BasicObject]
```

***
