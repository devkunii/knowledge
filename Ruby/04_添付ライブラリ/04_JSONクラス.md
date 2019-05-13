04 JSONクラス
=============

## 目次

* [JSONとは](#1JSONとは)

* [JSONの読み込み](#2JSONの読み込み)

* [JSONへの書き込み](#3JSONへの書き込み)



## 1.JSONとは

* YAMLと同様に、簡単な記法でデータの階層構造を表現できるフォーマット

* ブラウザとサーバー間でデータを交換する形式として使用される

```ruby
# 配列["Red", "Green", "Blue"]のjson表現
["Red", "Green", "Blue"]

# ハッシュ{"country" => "Japan", "state" => "Tokyo"}のjson表現
{"country": "Japan", "state": "Tokyo"}
```



## 2.JSONの読み込み

### loadメソッド

* `load(source, proc = nil)`：文字列からJSONを読み込む。

  * `source`はJSON形式の文字列を指定する(to_str、to_io、readメソッドを持つオブジェクトなら、文字列以外も指定可能)

  * `proc`には要素を読み込むごとに呼び出されるProcオブジェクトを渡すことができる

```ruby
require 'json'

json = <<-DATA
["Red", "Green", "Blue"]
DATA

p JSON.load(json)
=> ["Red", "Green", "Blue"]
```

* `load`にprocを渡すと、Procには解釈されたオブジェクトから順番に渡される

```ruby
require 'json'

json = <<-DATA
["Red", "Green", "Blue"]
DATA

JSON.load(json, lambda{|x| p x})

# 実行結果
"Red"
"Green"
"Blue"
["Red", "Green", "Blue"]
```


### parseメソッド

* `parse(source, options = {})`：文字列からJSONを読み込む。`load`とは異なり、`source`には文字列しか渡すことができない

  * `options`には解釈する時のオプションを指定することができる

  * YAMLとは違い、ファイル名から直接読み込むメソッドはないが、`load`メソッドの引数にioオブジェクトを指定してデータを読み込むことができる

```ruby
# 6-3.sample.json
["Red", "Green", "Blue"]

# 6-3.test.rb
require 'json'

p JSON.load(File.open("6-3.sample.json"))
=> ["Red", "Green", "Blue"]
```



## 3.JSONへの書き込み

### dumpメソッド

* `dump(obj, io = nil, limit = nil)`：`obj`に書き出したい対象のオブジェクトを指定する

  * 文字列で結果を得たい場合は、第2・3引数を省略して呼び出す

```ruby
require 'json'

array = ["Red", "Green", "Blue"]
p JSON.dump(array)
=> "[\"Red\",\"Green\",\"Blue\"]"
```

* ファイル等に書き出したい時は、ioにIOオブジェクトを指定する

* `limit`はダンプするオブジェクトの参照の深さを指定することができ、`limit`以上深くリンクしたオブジェクトを

* ダンプしようとすると、`ArgumentError`が発生する

```ruby
require 'json'

array = ["Red", "Green", "Blue"]
File.open("6-3.sample.json", "w+") do |f|
  JSON.dump(array, f)
end

# 実行結果
["Red","Green","Blue"]
```



| 版     | 年/月/日   |
| ------ | ---------- |
| 初版   | 2018/10/19 |
| 第二版 | 2019/05/13 |
