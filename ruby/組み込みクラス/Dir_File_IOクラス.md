# 5-10.`Dir`クラス、`File`クラス、`IO`クラス

* `Dir`クラス：ディレクトリの移動や作成、ディレクトリ内のファイル一覧の取得など、ディレクトリを扱うクラス

* `File`クラス：ファイルの読み取りや書き込み、新規作成や削除など、ファイルを扱うクラス

* `IO`クラス：`File`クラスのスーパークラスで、ファイルやプロセスなどとの入出力を扱うクラス

***

## 5-10-1.`Dir`クラス

### ディレクトリを開く、閉じる

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

### 開いているディレクトリのパスの取得

* `path`メソッド：開いているディレクトリのパスを取得

* `Dir.open`メソッドはブロックを取ることができ、この場合はブロックを出るときに自動的に閉じられる

```ruby
>> Dir.open("/usr/local/bin"){|dir| puts dir.path}
/usr/local/bin
=> nil          # 自動的に閉じられる
```

***

### カレントディレクトリの取得

* `Dir.pwd`、`Dir.getwd`メソッド：カレントディレクトリを取得する

```ruby
>> Dir.pwd
=> "/Users/MacUser/work/rails/shared_hobby"
```

***

### カレントディレクトリの移動

* `Dir.chdir`メソッド：カレントディレクトリを指定されたディレクトリに変更する。

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

### ディレクトリの作成

* `mkdir`メソッド：指定したパスのディレクトリを作成する。2つ目の引数にパーミッション(mode)を指定可能

* 通常、パーミッションは3桁の8進数で指定。実際のパーミッションは、指定された値と`unmask`をかけた値(`mode & ~unmask`)となる

パーミッションがよくわからないので、省略

2018/9/10

```ruby
>> Dir.mkdir("/tmp/foo")
=> 0
>> Dir.mkdir("/tmp/bar", 0755)
=> 0
```

***

### ディレクトリの削除

* `rmdir`メソッド：ディレクトリを削除する

```ruby
>> Dir.mkdir("/tmp/foo")
=> 0
>> Dir.rmdir("/tmp/foo")
=> 0
```

***

## 5-10-2.`File`クラス

### ファイルを開いて読み込む

* `File.open`、`File.new`メソッド：ファイルを開く。

* 引数としてファイル名だけを与えると、読み取りモード`"r"`で開く。

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

***

### ファイルのモード

* `File.open`メソッドの2番目の引数は、ファイルを開くモードを指定できる

  * `"r"`：読み込みモード

  * `"w"`：書き込みモード。既存ファイルの場合は、ファイルの内容を空にする

  * `"a"`：追記モード。常にファイルの末尾に追加される

  * `"r+"`：読み書きモード。ファイルの読み書き位置が先頭になる

  * `"w+"`：読み書きモード。`"r+"`と同じだが、既存ファイルの場合はファイルの内容が空になる

  * `"a+"`：読み書きモード。ファイルの読み込み位置は先頭に、書き込み位置は常に末尾になる

※Silverで間違えている

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

### ファイルのエンコーディング

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

### ファイルに書き込む

* `write`メソッド：ファイルに文字を書き込む。ファイルオブジェクトのメソッド

* `IO`クラスに、他のメソッドの記述あり

```ruby
>> File.open('new-file', "w") {|file| file.write "This is new file."}
=> 17
```

***

### ファイルの属性を取得する

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

* `File.join`：引数で与えられた文字列を、`File::SEPARATOR`で連結。通常`/`が設定されている。ただし、前の文字列が`/`で終わる場合は入れない

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

**ファイルオブジェクトのメソッドによる取得**

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

***

### ファイルをテストする(`Filetest`モジュール)

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

### ファイルの属性を設定する

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

***

### ファイルのパスを絶対パスに展開する

* `File.expand_path`メソッド：指定されたパスを絶対パスに展開する

```ruby
>> File.expand_path('README.md')
=> "/Users/MacUser/work/rails/shared_hobby/README.md"
```

***

### ファイルを削除する、リネームする

* `delete`、`unlink`メソッド：指定されたファイルを削除する。削除に失敗した場合は、エラーが発生する

* `truncate`メソッド：ファイルを指定したバイト数に切り詰める。

* `rename`メソッド：1つ目の引数で指定したファイル名を、2つ目の引数で指定したファイル名に変更する。
リネーム先のファイルが存在する場合は、ファイルを上書きする

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

### ファイルをロックする

* `flock`メソッド：ファイルをロックする。引数にはロック方法を指定する。

```ruby
>> File.open('README.md', "w") {|file| file.flock(File::LOCK_EX)}
=> 0
```

***

## 5-10-3.`IO`クラス

* `File`クラスのスーパークラスであり、基本的な入出力機能を備えたクラス

* 多くのメソッドは`File`クラスでも利用できる

* 標準出力(`STDOUT`)、標準入力(`STDIN`)、標準エラー出力(`STDERR`)は`IO`クラスのオブジェクト

***

### `IO`を開く

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

### `IO`からの入力


* `IO.read`、`read`：`IO`から内容を読み込む。長さが指定されていれば、その長さだけ読み込む。
  長さを指定した場合のみ、バイナリ読み込みとなり、エンコーディングが **ASCII-8BIT** となる

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
  →Silverで間違えている！！

* `read`：長さが指定されていない場合は`""`、指定されている場合は`nil`が返る

* `readchar`、`readline`：`EOFError`エラーが発生する
  →Silverで間違えている！！

* `readlines`：空配列`[]`が返る

* `getbyte`：`nil`が返る

* `readbyte`：`EOFError`エラーが発生する

***

#### `IO`への出力

* `write`：`IO`に対して引数の文字列を出力する。引数が文字列以外の場合は、`to_s`メソッドで文字列化して出力
  →出力が成功すると、出力した文字列のバイト数を返す

* `puts`：`IO`に対して複数のオブジェクトを出力する。引数が文字列や配列でない場合、`to_ary`メソッドにより配列化し、
  次に各要素を`to_s`メソッドにより文字列化して出力する

* `print`：`IO`に対して複数のオブジェクトを出力する。`puts`メソッドと異なり、複数のオブジェクトが指定されると、
  各オブジェクトの間に`$,`の値を出力する。`$\`に値が設定されていれば最後に出力する。
  引数が文字列でない場合には、`to_s`メソッドで文字列化して出力する

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

* `flush`：`IO`の内部バッファをフラッシュ(強制的に出力)して出力する

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
>> io.read                 # ioを全て読み込んだため、最終行に達する
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

※Silverで間違えている

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
