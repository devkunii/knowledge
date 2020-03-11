05 CSVクラス
===========

## 目次

* [CSVとは](#1CSVとは)

* [読み書き可能なメソッド](#2読み書き可能なメソッド)

* [読み込みのみ可能なメソッド](#3読み込みのみ可能なメソッド)



## 1.CSVとは

* Comma Separated Values

* データを感まで区切ったテキストデータ

* 古くから多くの表計算ソフトや、データベースソフトなどで使用されているデータフォーマット

```csv
abc,def,efg,hijk
123,456,789,10
ABCDEFG,,HIJKLM,,
```



## 2.読み書き可能なメソッド

### openメソッド

* `open`：CSVをモードで指定した形式でオープンし、ブロックで処理する

> modeで指定できる形式

|mode|        意味       |
|----|:------------------|
|'r' |読み込み            |
|'w' |書き込み            |
|'rb'|バイナリ読み込みモード|
|'rw'|バイナリ書き込みモード|

* 読み込み時はfilenameで指定したファイルをオープンし、各行を配列としてブロック引数に渡す

  * ブロックを渡さなかった場合は、CSVオブジェクトが返る

```ruby
CSV.open(filename, mode = "rb", options = {}){|csv| ...}
CSV.open(filename, mode = "rb", options = {})
CSV.open(filename, options = {}){|csv| ...}
CSV.open(filename, options = {})
```

* デフォルトの区切り文字は`CRLF/LF`、`CR`を区切りにするには、optionsに`{ row_sep: ?\r}`を渡す必要がある

```ruby
require 'csv'

CSV.open("6-3.sample.csv") do |csv|
  csv.each do |row|
    p row
  end
end

# 実行結果
["abc", "def", "efg", "hijk"]
["123", "456", "789", "10"]
["ABCDEFG", nil, "HIJKLM", nil, nil]
```

* 書き込み時はfilenameで指定したファイルをオープンし、CSVオブジェクトをブロック引数に渡す

* ブロックを渡さなかった場合は、CSVオブジェクトが返る

```ruby
require 'csv'

CSV.open("6-3.sample.csv", "w") do |row|
  row << ["abc", "def", "ghijk"]
  row << ["lmn", "opq", "rstuv"]
end

# 実行結果
abc,def,ghijk
lmn,opq,rstuv
```



## 3.読み込みのみ可能なメソッド

### foreachメソッド

* `foreach`：読み込みを行う。ブロック引数にEnumeratorオブジェクトを受け取り、それ経由で各行を受け取ることができる

  * ブロックを渡さなかった場合は、Enumeratorオブジェクトが返る

```ruby
CSV.foreach(path, options = {}){|row| ...}
```



### read・readlinesメソッド

* `read`・`readlines`：読み込みを行う。内部の処理は異なるが、どちらも読み込み結果を配列の配列に格納して返す

  * optionsに`|headers: true|`を渡すと、`CSV::Table`が返る

```ruby
require 'csv'

p CSV.read("6-3.sample.csv")
=> [["abc", "def", "efg", "hijk"], ["123", "456", "789", "10"], ["ABCDEFG", nil, "HIJKLM", nil, nil]]

p CSV.readlines("6-3.sample.csv")
=> [["abc", "def", "efg", "hijk"], ["123", "456", "789", "10"], ["ABCDEFG", nil, "HIJKLM", nil, nil]]

csv_table = CSV.read("6-3.sample.csv", headers: true)
p csv_table
=> #<CSV::Table mode:col_or_row row_count:3>
p csv_table.to_a
=> [["abc", "def", "efg", "hijk"], ["123", "456", "789", "10"], ["ABCDEFG", nil, "HIJKLM", nil, nil]]
```



| 版     | 年/月/日   |
| ------ | ---------- |
| 初版   | 2018/10/19 |
| 第二版 | 2019/05/13 |
