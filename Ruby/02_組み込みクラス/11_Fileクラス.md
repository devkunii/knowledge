11 Fileクラス
============

## 目次

* [Fileクラスとは](#0Fileクラスとは)

* [ファイルを開いて読み込む](#1ファイルを開いて読み込む)

* [ファイルのモード](#2ファイルのモード)

* [ファイルのエンコーディング](#3ファイルのエンコーディング)

* [ファイルに書き込む](#4ファイルに書き込む)

* [ファイルの属性を取得する](#5ファイルの属性を取得する)

* [ファイルをテストする](#6ファイルをテストする)

* [ファイルの属性を設定する](#7ファイルの属性を設定する)

* [ファイルのパスを絶対パスに展開する](#8ファイルのパスを絶対パスに展開する)

* [ファイルを削除する・リネームする](#9ファイルを削除するリネームする)

* [ファイルをロックする](#10ファイルをロックする)



## 0.Fileクラスとは

ファイルの読み取りや書き込み、新規作成や削除など、ファイルを扱うクラス



## 1.ファイルを開いて読み込む

* `File.open`、`File.new`メソッド：ファイルを開く。

* 引数としてファイル名だけを与えると、読み取りモード`"r"`で開く。

* ファイルが存在しない場合は、エラーが発生する。

* ファイルを開くとファイルオブジェクトが返り、

  * `read`メソッド：ファイルの内容を取得

  * `close`メソッド：ファイルを閉じる

* ファイルの入出力時には、エンコーディングが有効になる

* `File.open`メソッドにブロックを与えると、ブロック終了時に自動的にファイルを閉じることができる


```ruby
>> file = File.open("README.md")
=> #<File:README.md>
>> file.read
=> "# 省略"
>> file.close
=> nil

# 入出力時のエンコーディング
>> Encoding.default_external
=> #<Encoding:UTF-8>
>> file = File.open("README.md")
=> #<File:README.md>
>> file.read
=> "# 省略"
>> file.read.encoding
=> #<Encoding:UTF-8>

# ファイルをブロックで開く
>> File.open("README.md"){|file| file.read}
=> "# 省略"
```



## 2.ファイルのモード

* `File.open`メソッドの2番目の引数は、ファイルを開くモードを指定できる

  * `"r"`：読み込みモード

  * `"w"`：書き込みモード。既存ファイルの場合は、ファイルの内容を空にする

  * `"a"`：追記モード。常にファイルの末尾に追加される

  * `"r+"`：読み書きモード。ファイルの読み書き位置が先頭になる

  * `"w+"`：読み書きモード。`"r+"`と同じだが、既存ファイルの場合はファイルの内容が空になる

  * `"a+"`：読み書きモード。ファイルの読み込み位置は先頭に、書き込み位置は常に末尾になる


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

###  `w+`

* `w+`は新規作成・読み込み + 書き込みモードで開く

* 既にファイルが存在する場合は、空になる

* `IO#seek`はファイルポインタを指定の位置に移動する

* `IO:SEEK_SET`がファイルの先頭からの位置を指定する識別子

> `recode 1`を書き込み後にファイルの先頭にファイルポインタを移動し、`recode 2`で上書きしています。



### `a+`

* `a+`はファイルを読み込みモード + 追記書き込みモードで開く

* ファイルの読み込みは、ファイルの先頭から行いますが、書き込みは、ファイルの末尾に行います。



## 3.ファイルのエンコーディング

* モードの後ろに以下の項目を指定可能

  * ファイルのエンコーディング(外部エンコーディング)

  * 読み込んだ時のエンコーディング(内部エンコーディング)

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



## 4.ファイルに書き込む

* `write`メソッド：ファイルに文字を書き込む

  * ファイルオブジェクトのメソッド

* `IO`クラスに、他のメソッドの記述あり

```ruby
>> File.open('new-file', "w") {|file| file.write "This is new file."}
=> 17
```



## 5.ファイルの属性を取得する

**ファイルの属性を取得するメソッド**

* `File.basename`：指定されたパスからファイル名を取得する

* `File.dirname`：指定されたパスからディレクトリ名を取得する

* `File.extname`：指定されたパスから拡張子を取得する

```ruby
>> p File.basename("/Users/MacUser/work/knowledge/ruby/組み込みクラス/mistake.md")
"mistake.md"
=> "mistake.md"
>> p File.dirname("/Users/MacUser/work/knowledge/ruby/組み込みクラス/mistake.md")
"/Users/MacUser/work/knowledge/ruby/組み込みクラス"
=> "/Users/MacUser/work/knowledge/ruby/組み込みクラス"
>> p File.extname("/Users/MacUser/work/knowledge/ruby/組み込みクラス/mistake.md")
".md"
=> ".md"
```

* `File.split`：指定されたパスからディレクトリ名とファイル名の配列を取得する

* `File.join`：引数で与えられた文字列を、`File::SEPARATOR`で連結

  * 通常`/`が設定されている。

  * ただし、前の文字列が`/`で終わる場合は入れない

```ruby
# splitメソッド
>> p File.split("/usr/local/bin/ruby")
["/usr/local/bin", "ruby"]
=> ["/usr/local/bin", "ruby"]
>> p File.split("ruby")
[".", "ruby"]
=> [".", "ruby"]

# joinメソッド
>> puts File.join("/", "user", "bin")
/user/bin
=> nil
```

* `File.stat`、`File.lstat`：ファイルの属性を示す`File::Stat`クラスのオブジェクトを返す

* `File.atime`、`File.ctime`、`File.mtime`：それぞれのファイルの最終アクセス時刻、状態が変更された時刻、最終更新時刻を取得する



### ファイルオブジェクトのメソッドによる取得

* `path`：ファイルを開くときに使用したパスを返す

* `lstat`：ファイルの属性を示す`File::Stat`クラスのオブジェクトを返す

  * `actime`、`ctime`、`mtime`：それぞれのファイルの最終アクセス時刻、状態が変更された時刻、最終更新時刻を取得する

  * `size`：ファイルサイズ

```ruby
>> filename = "foo"
=> "foo"
>> File.open(filename, "w").close
=> nil
>> st = File.stat(filename)
=> #<File::Stat dev=0x1000004, ino=8607491803, mode=0100644, nlink=1, uid=501, gid=20, rdev=0x0, size=0, blksize=4194304, blocks=0, atime=2018-09-29 19:06:42 +0900, mtime=2018-09-29 19:06:42 +0900, ctime=2018-09-29 19:06:42 +0900, birthtime=2018-09-29 19:06:42 +0900>
>> p st.mtime
2018-09-29 19:06:42 +0900
=> 2018-09-29 19:06:42 +0900
>> p st.size
0
=> 0
```



## 6.ファイルをテストする

### Filetestモジュールとは

* ファイルの存在確認や、ディレクトリかどうかの判定など、ファイルをテストするメソッド



### Filetestモジュールのメソッド

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



## 7.ファイルの属性を設定する

* `File.chmod`メソッド：ファイルの属性を変更する

* `File.chown`メソッド：ファイルの所有者を変更する

* `File.utime`メソッド：ファイルの最終アクセス時刻や更新時刻を設定する

```ruby
# File.chmod、File.chownメソッド
>> File.chmod(0644, 'README.md')  # 0644：所有者は読み書きの両方をできるが、それ以外は読み込みのみ
=> 1
>> File.chown(501, 20, 'README.md')
=> 1

# File.utimeメソッド
>> File.utime(Time.now, Time.now, 'README.md')
=> 1
```



## 8.ファイルのパスを絶対パスに展開する

* `File.expand_path`メソッド：指定されたパスを絶対パスに展開する

```ruby
>> File.expand_path('README.md')
=> "/Users/MacUser/work/rails/shared_hobby/README.md"
```



## 9.ファイルを削除する・リネームする

* `delete`、`unlink`メソッド：指定されたファイルを削除する

  * 削除に失敗した場合は、エラーが発生する

* `truncate`メソッド：ファイルを指定したバイト数に切り詰める。

* `rename`メソッド：1つ目の引数で指定したファイル名を、2つ目の引数で指定したファイル名に変更する。

  * リネーム先のファイルが存在する場合は、ファイルを上書きする

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



## 10.ファイルをロックする

* `flock`メソッド：ファイルをロックする

  * 引数にはロック方法を指定する。

```ruby
>> File.open('README.md', "w") {|file| file.flock(File::LOCK_EX)}
=> 0
```



| 版     | 年/月/日   |
| ------ | ---------- |
| 第二版 | 2019/05/11 | 
