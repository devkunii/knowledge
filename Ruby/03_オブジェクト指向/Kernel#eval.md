Kernel#eval
===========

## Kernel#eval

* `eval`：ブロックの代わりにRubyのコードの文字列受け取る(コード文字列)

  渡されたコード文字列を実行して、その結果を戻す

  ```ruby
  >> array = [10, 20]
  => [10, 20]
  >> element = 30
  => 30
  >> eval("array << element")
  => [10, 20, 30]
  ```

* REST Clientの例

  ```ruby
  def get(path, *args, &b)
    r[path].get(*args, &b)
  end
  ```

  これをそれぞれのResourceで何度も定義するのは無理がある

  以下のコードで対応可能

  ```ruby
  POSSIBLE_VERBS = ['get', `put`, `post`, `delete`]

  POSSIBLE_VERBS.each do |m|
    eval <<-end_eval
      def #{m}(path, *args, &b)
        r[path].#{m}(*args, &b)
      end
    end_eval
  end
  ```

* Bindingオブジェクトの例

  * `Binding`：スコープをオブジェクトにまとめたもの

  * `Binding`を作ってローカルスコープを取得することで、そのスコープを持ち回すことができる

  * `eval`と一緒に組み合わせて使えば、後からそのスコープでコードを実行することができる

  * Kernel#Bindingメソッドで生成される

  ```ruby
  >> class MyClass
  >>   def my_method
  >>     @x = 1
  >>     binding
  >>   end
  >> end
  => :my_method

  >> b = MyClass.new.my_method
  => #<Binding:0x007f8b981fd1c0>
  ```

  Bindingオブジェクトには、スコープは含まれているが、コードは含まれていない

  ->ブロックよりも純粋な **クロージャ**

  取得したスコープでコードを評価するには、`eval`の引数にBindingを渡す

  ```ruby
  >> eval "@x", b
  => 1
  ```

  `TOPLEVEL_BINDING`：トップレベルのスコープのBinding(Rubyが事前に定義している定数)

  ->トップレベルのスコープにプログラムのどこからでもアクセスできる

  ```ruby
  >> class AnotherClass
  >>   def my_method
  >>     eval "self", TOPLEVEL_BINDING
  >>   end
  >> end
  => :my_method

  >> AnotherClass.new.my_method
  => main
  ```

* irbの例

  * `irb`：標準入力やファイルをパースして、それぞれの行を`eval`に渡すシンプルなプログラム

    ->「コードプロセッサ」という

  ```ruby
  # irbのevalを呼び出す部分
  eval(statements, @binding, file, line)
  ```

  * `statements`：Rubyのコード行

  * `Binding`：irbがコードを異なるコンテキストで実行する時に使用する

    ->特定のオブジェクトでirbセッションをネストして開くような時

    ->既存のirbのセッションの中で、irbとオブジェクトの名前を入力すれば、以降のコマンドはそのコンテキストで評価される

    ->`instance_eval`と似ている

  * `file`・`line`：例外が発生した時にスタックトレースを調整するために使用する

  ```ruby
  >> x = 1 / 0
  ZeroDivisionError: divided by 0
  	from (irb):8:in '/'
  ```

* 「コード文字列」と「ブロック」

  * `eval`：コード文字列を常に必要とする

  * `class_eval`・`instance_eval`：コード文字列orブロックのいずれかを受け取る

  * コード文字列は、ブロックと同じように、ローカル変数にアクセスすることができる

  ```ruby
  >> array = ['a', 'b', 'c']
  => ["a", "b", "c"]
  >> x = 'd'
  => "d"
  >> array.instance_eval "self[1] = x"
  => "d"

  >> array
  => ["a", "d", "c"]
  ```

  * 可能であれば、コード文字列を避ける



## evalの問題点

* コード文字列は、「シンタクッスハイライト」や「自動補完」といったエディタの機能が使えない

  ->実行時に予期せずに失敗するような脆弱性のあるプログラムになる可能性



### コードインジェクション

* Arrayのメソッドを確認する例

```ruby
>> def explore_array(method)
>>   code = "['a', 'b', 'c'].#{method}"
>>   puts "Evaluating: #{method}"
>>   eval code
>> end
=> :explore_array
>>
>> loop { p explore_array(gets.chomp)}

find_index("b")
Evaluating: find_index("b")
1

map! {|e| e.next }
Evaluating: map! {|e| e.next }
["b", "c", "d"]

object_id; Dir.glob("*")
Evaluating: object_id; Dir.glob("*")
["github", "ruby", "chemical"]       # プライベートな情報(現在実行しているディレクトリの一覧)
```

* 悪意のあるユーザーが、作成されたコンピュータの脅威になる任意のコードを実行できてしまう

  ->脆弱性を突く行為：「コードインジェクション攻撃」



### コードインジェクションの対策

```ruby
POSSIBLE_VERBS = ['get', `put`, `post`, `delete`]

POSSIBLE_VERBS.each do |m|
  eval <<-end_eval
    def #{m}(path, *args, &b)
      r[path].#{m}(*args, &b)
    end
  end_eval
end
```

* 対策1：「動的メソッド」&「動的ディスパッチ」で置き換える

```ruby
POSSIBLE_VERBS.each do |m|
  define_method m do |path, *args, &b|
    r[path].send(m, *args, &b)
  end
end
```

```ruby
def explore_array(method, *arguments)
  ['a', 'b', 'c'].send(method, *arguments)
end
```



### オブジェクトの汚染とセーフレベル

* Rubyは、潜在的に安全ではないオブジェクト(特に外部から来たオブジェクト)に自動的に **汚染** の印をつける

  * ウェブフォーム

  * ファイル

  * コマンドライン

  * システム変数(プログラムが読み込んだ文字列)

* 汚染された文字列を操作して新しい文字列を作ると、その新しい文字列も汚染される

  ->`tainted?`：オブジェクトが汚染されているかどうかを確認

```ruby
# ユーザー入力を読み込む
user_input = "User input: #{gets()}"
puts user_input.tainted?

x = 1
true
```

* セーフレベル：オブジェクトの汚染を補完する

  ->`$SAFE`：0〜3の4つのレベルが存在する

  -> 2では、ファイル操作はほとんど認められていない

  -> 0より大きいセーフレベルでは、Rubyは汚染した文字列を評価できない

```ruby
$SAFE = 1
user_input = "User input: #{gets()}"
eval user_input

x = 1
SecurityError: Insecure operation - eval
```

* 安全性を自分で調整するには、コード文字列を評価する前に明示的に汚染を除去(`Object#untaint`を呼び出す)、

  セーフレベルに頼って、ディスクアクセスのような危険な操作を抑止する

* ERBの例

```rhtml
<p><strong>Wake up!</strong>Its a nice sunny <%= Time.new.strftime("%A") %>.</p>
```

```ruby
require 'erb'
erb = ERB.new(File.read('template.rhtml'))
erb.run

=> <p><strong>Wake up!</strong>Its a nice sunny Friday.</p>
```

  * 実行部分

  ```ruby
  class ERB
    def result(b = new_toplevel)
      if @safe_level
        proc {
          $SAFE = @safe_level
          eval(@src, b, (@filename || '(erb)'), 0)
        }.call
      else
        eval(@src, b, (@filename || '(erb)'), 0)
      end
    end
  end
  ```

  * `new_toplevel`：`TOPLEVEL_BINDING`のコピーを戻すメソッド

  * `@src`：ERBのタグの中身

  * `@safe_level`：ユーザーが必要とするセーフレベル

    * 未設定：タグの中身がそのまま評価される

    * 設定済：ERBはサンドボックス(`eval`用に制御した環境)を作る

      その中で、グローバルのセーフレベルをユーザーの指定と一致させ、Procをクリーンルームにして、別のスコープでコードを実行



## フックメソッド

* オブジェクトモデルは、

  * クラスを継承される

  * モジュールがクラスにMix-inされる

  * メソッドが定義される

  * 定義が解除される

  * 削除される

  など、コードを実行すると様々なイベントが起きる

* これらのイベントをキャッチすることができる

```ruby
>> class String
>>   def self.inherited(subclass)
>>     puts "#{self}は#{subclass}に継承された"
>>   end
>> end
=> :inherited

>> class MyClass < String; end
StringはMyClassに継承された
=> nil
```

* `inherited`：Classクラスのインスタンスメソッド。クラスが継承された時にRubyが呼び出してくれる

* デフォルトでは何もしないので、自分のコードでオーバーライドして使う

* 「フックメソッド」：特定のイベントにフックを掛ける



### フックの例

* `Module#included`・`Module#prepend`をオーバーライドすれば、モジュールのライフサイクルにプラグインできる

```ruby
>> module M1
>>   def self.included(othermod)
>>     puts "M1 は #{othermod} にインクルードされた"
>>   end
>> end
=> :included

>> module M2
>>   def self.prepended(othermod)
>>     puts "M2 は #{othermod} にプリペンドされた"
>>   end
>> end
=> :prepended

>> class C
>>   include M1
>>   prepend M2
>> end
M1 は C にインクルードされた
M2 は C にプリペンドされた
=> C
```

* `Module#extended`をオーバーライドすれば、モジュールがオブジェクトを拡張した時にコードを実行できる

* `Module#method_added・method_removed・method_undefined`をオーバーライドすれば、メソッドに関連したイベントを実行できる

```ruby
>> module M
>>   def self.method_added(method)
>>     puts "新しいメソッド：M##{method}"
>>   end
>>   def my_method; end
>> end
新しいメソッド：M#my_method
=> :my_method
```

* これらのフックは、オブジェクトのクラスに住むインスタンスメソッドにしか使えない

  ->オブジェクトの特異クラスに住む特異メソッドでは動作しない

* 特異メソッドのイベントをキャッチするには、

  * `Kernel#singleton_method_added`

  * `Kernel#singleton_method_removed`

  * `Kernel#singleton_method_undefinded`



### Module#included

* もっとも広く使われているフック

* VCRの例

  * HTTP呼び出しを記録&再生するgem

```ruby
module VCR
  class Request
    include Normalizers::Body
    # 省略
  end
end
```

* Bodyモジュールは、HTTPメッセージボディーを扱う`body_from`などのメソッドを定義している

* モジュールをインクルードすると、これらのメソッドがRequestクラスのメソッドになる

  ->Requestが、Normalizers::Bodyをインクルードすることにより、クラスメソッドを手に入れる

* クラスがモジュールをインクルードすると、「インスタンスメソッド」が手に入るが、Normalizers::Bodyはクラスメソッドはどうすればクラスメソッドを定義しているのか

```ruby
module VCR
  module Normalizers
    module Body
      def self.included(klass)
        klass.extend ClassMethods
      end

      module ClassMethods
        def body_from(hash_or_string)
          # 省略
        end
      end
    end
  end
end
```

* BodyにはClassMethodsという名前の内部クラスが存在しており、そこに`body_from`などの通常のインスタンスメソッドが定義さている

* Bodyには、`included`というフックメソッドがある。RequestがBodyをインクルードすると

  1. Rubyが、Bodyの`included`フックを呼び出す

  1. フックがRequestに戻り、ClassMethodsモジュールをエクステンドする

  1. `extend`メソッドが、ClassMethodsのメソッドをRequestの特異クラスにインクルードする

* `body_from`などのインスタンスメソッドがRequestの特異クラスにMix-inされる

  ->実質、Requestのクラスメソッドになる



| 版 |   年月日  |
|----|----------|
|初版|2018/11/24|
