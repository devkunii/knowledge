02 Rubyインタプリタの実行
======================

## rubyインタプリタの実行

```ruby
ruby [オプション][--][プログラムファイル][引数]
```



## 主なコマンドラインオプション

* `-h`：ヘルプの表示

* `-v`：rubyのバージョンの表示

* `-c`：指定されたファイルが文法的に正しいかを確認(実行しない)

* `-e`：指定された文字列をrubyプログラムとして実行

* `-r`：スクリプト実行時に、指定されたファイルを実行する

* `-d`：デバックモードでの実行

* `-w`：冗長モードでの出力

* `-l`：`$LOAD_PATH`に、指定された文字列を追加

例)冗長モードの出力

* `test.rb`

```ruby
a = 3
b = 2
puts a
```

> ターミナル上

```ruby
$ ruby -w test.rb  #冗長モードでの出力
test.rb:2: warning: assigned but unused variable - b
3

$ ruby -W test.rb  #全ての警告を出力
test.rb:2: warning: assigned but unused variable - b
3

$ ruby -W2 test.rb  #全ての警告を出力
test.rb:2: warning: assigned but unused variable - b
3

$ ruby -W1 test.rb  #重要な警告のみ
3

$ ruby -W0 test.rb  #警告を出力しない
3
```



### コマンドラインオプションの`-r`

* **`require`と`load`**

* `-r`はスクリプトの実行時に、指定されたライブラリを読み込む

* 指定されたファイルを絶対パスor相対パスで指定



### コマンドラインオプションの`-l`

* 指定されたファイルを相対パスで指定した時に、組み込み変数

* `Kernel#require`・`Kernel#load`を用いる時の探索パスが配列で格納されている

* `$LOAD_PATH`に格納されたパス順に探す

```ruby
>> $LOAD_PATH
=> ["/usr/local/Cellar/rbenv/1.1.1/rbenv.d/exec/gem-rehash",
   "/Users/MacUser/.rbenv/versions/2.4.1/lib/ruby/gems/2.4.0/gems/did_you_mean-1.1.0/lib",
   "/Users/MacUser/.rbenv/versions/2.4.1/lib/ruby/site_ruby/2.4.0",
   "/Users/MacUser/.rbenv/versions/2.4.1/lib/ruby/site_ruby/2.4.0/x86_64-darwin17",
   "/Users/MacUser/.rbenv/versions/2.4.1/lib/ruby/site_ruby", "/Users/MacUser/.rbenv/versions/2.4.1/lib/ruby/vendor_ruby/2.4.0",
   "/Users/MacUser/.rbenv/versions/2.4.1/lib/ruby/vendor_ruby/2.4.0/x86_64-darwin17",
   "/Users/MacUser/.rbenv/versions/2.4.1/lib/ruby/vendor_ruby", "/Users/MacUser/.rbenv/versions/2.4.1/lib/ruby/2.4.0",
   "/Users/MacUser/.rbenv/versions/2.4.1/lib/ruby/2.4.0/x86_64-darwin17"]
```

#### 順番

1. Rubyのインストールフォルダ

1. Rubyを実行したフォルダ(カレントディレクトリ)

> この際に、読み込み順序を拡張する際は`-l`オプションを使用

> このオプションで指定されたフォルダを最優先で検索

例)rubyインタプリタが参照する主な環境変数

* `RUBYOPT`：デフォルトで指定するコマンドラインオプションを指定

* `RUBYLIB`：デフォルトでライブラリを検索するパスを指定

* `PATH`：Ruby上から実行したコマンドを検索するパスを指定



| 版     | 年/月/日   |
| ------ | ---------- |
| 初版   | 2018/08/11 |
| 第二版 | 2019/05/09 |
