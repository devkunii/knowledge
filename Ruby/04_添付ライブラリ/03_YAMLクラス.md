03 YAMLクラス
============

## 目次

* [YAMLとは](#0YAMLとは)

* [YAMLの読み込み](#1YAMLの読み込み)

* [YAMLの読み込み](#2YAMLの読み込み)



## 0.YAMLとは

* xmlよりもさらに簡単な記法でデータの階層構造を表現できるフォーマット

* ハッシュ・配列の組み合わせをスペースによるインデントで表現したもの

```YAML
# 配列["Red", "Green", "Blue"]のYAML表現
- Red
- Green
- Blue

# ハッシュ{"country" => "Japan", "state" => "Tokyo"}のYAML表現
country: Japan
state: Tokyo
```



## 1.YAMLの読み込み

* YAMLの入出力に関するメソッドは、全てクラスメソッドなので、インスタンスを生成しないで呼び出すことが可能



### loadメソッド

* `load`：YAMLデータが記録された文字列またはIOインスタンスからHashクラスのインスタンスを生成

  * 文字列またはIOに複数のYAMLが記録されていても、最初のYAMLのみが生成(残りは無視)

```ruby
require 'yaml'

yaml_data = <<-DATA
- Red
- Green
- Blue
---
- Yellow
- Pink
- White
DATA
p YAML.load(yaml_data)
=> ["Red", "Green", "Blue"]
```



### load_fileメソッド

* `load_file`：指定したパスのYAMLファイルからオブジェクトを生成する

  * `YAML.load`メソッドと同様に、ファイルに複数のYAMLが記録されていても、最初のYAMLのみが生成(残りは無視)

```ruby
# 6-3.sample.yml
- Red
- Green
- Blue
---
- Yellow
- Pink
- White

# 6-3.test.rb
require 'yaml'

p YAML.load_file "6-3.sample.yml"
=> ["Red", "Green", "Blue"]
```



### load_streamメソッド

* `load_stream`：ioに記録された複数のYAMLデータを順番にロードし、結果を`YAML::Stream`のインスタンスで返す

```ruby
require 'yaml'

p yaml_stream = YAML.load_stream(File.open "6-3.sample.yml")
=> [["Red", "Green", "Blue"], ["Yellow", "Pink", "White"]]
```



### load_documentsメソッド

* `load_documents`：引数のioに記録された複数のYAMLデータを順番にロードし、それぞれをブロック内で処理することができる

```ruby
require 'yaml'

YAML.load_documents(File.open "6-3.sample.yml") do |yaml|
  p yaml.first
end

# 実行結果
"Red"
"Yellow"
```



## 2.YAMLの読み込み

### dumpメソッド

* `dump(obj, io = nil)`：単一のインスタンスを文字列またはIOインスタンスに出力することができる。

  * 引数ioを省略すると、yaml形式に変化した文字列を返し、ioを指定すると指定した出力先にyaml形式のデータを書き込む

```ruby
require 'yaml'

colors = ["Red", "Green", "Blue"]
p YAML.dump colors
=> "---\n- Red\n- Green\n- Blue\n"

File.open("6-3.sample.yml", "w+") do |f|
  YAML.dump(colors, f)
end

# 実行結果
---
- Red
- Green
- Blue
```

### dump_streamメソッド

* `dump_stream(*objs)`：複数のオブジェクトを文字列に出力することができる

```ruby
require 'yaml'

colors1 = ["Red", "Green", "Blue"]
colors2 = ["Yellow", "Pink", "White"]

p YAML.dump_stream colors1, colors2

# 実行結果
"---\n- Red\n- Green\n- Blue\n---\n- Yellow\n- Pink\n- White\n"
```



| 版     | 年/月/日   |
| ------ | ---------- |
| 初版   | 2018/10/19 |
| 第二版 | 2019/05/13 |
