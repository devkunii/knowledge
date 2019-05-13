06 FileUtilsクラス
=================

## 目次

* [FileUtilsクラスとは](#1.FileUtilsクラスとは)

* [ディレクトリの確認・移動を行うコマンド](#2ディレクトリの確認移動を行うコマンド)

* [ファイル・ディレクトリのコピーを行うコマンド](#3ファイルディレクトリのコピーを行うコマンド)

* [ファイル・ディレクトリの移動を行うコマンド](#4ファイルディレクトリの移動を行うコマンド)

* [ファイル・ディレクトリの削除を行うコマンド](#5ファイルディレクトリの削除を行うコマンド)

* [ファイル・ディレクトリの作成を行うコマンド](#6ファイルディレクトリの作成を行うコマンド)

* [ファイル・ディレクトリの権限を変更するコマンド](#7ファイルディレクトリの権限を変更するコマンド)

* [ファイル・ディレクトリへのシンボリックリンクを作成するコマンド](#8ファイルディレクトリへのシンボリックリンクを作成するコマンド)



## 1.FileUtilsクラスとは

* 基本的なファイル操作を集めたライブラリ

* コピー、移動、削除、権限変更などのファイルやディレクトリに対する基本的な操作が行える

* ファイルやディレクトリの操作を、UNIXコマンドライクに呼び出せるようにし、扱いやすくしたもの

* メソッドはモジュール関数として提供されるため、呼び出しは以下のようにできる

```ruby
require 'fileutils'

# 方法1
FileUtils.mv 'foo', 'bar'

# 方法2
include FileUtils
mv 'foo', 'bar'
```

* メソッドにはoptionsを指定でき、ほとんどのメソッドは共通

```ruby
# 詳細をコンソールに出力するメソッド
:verbose => true

# コマンドを実行しない
:noop => true
```



## 2.ディレクトリの確認・移動を行うコマンド

* `cd`：引数dirで指定されたディレクトリへ移動し、成功したら`nil`を返す

  * 指定したディレクトリがない場合は、例外が発生する

* `pwd`：カレントディレクトリを文字列で返す。optionsには`:verbose`を指定できる

```ruby
FileUtils.cd(dif, options = {})
FileUtils.pwd
```



## 3.ファイル・ディレクトリのコピーを行うコマンド

* `cp`：srcで指定したファイルをdestにコピーする

  * 引数がディレクトリの場合は、destディレクトリ以下にsrcをコピーし、

  * destがファイルの場合は、destをsrcで上書きする

* `copy`：`cp`のエイリアスメソッド

  * optionsには`:preserve => true`(元のファイルの属性・タイムスタンプを保持する)、`:noop`、`:verbose`が指定可能

```ruby
FileUtils.cp(src, dest, options = {})
FileUtils.copy(src, dest, options = {})
```

* `cp_r`：`cp`メソッドの再帰オプション`-r`でコピーを行うコマンド

  * 単一または複数のファイルとディレクトリをsrcに指定することができる

  * ディレクトリを指定した場合は、ディレクトリ以下にあるファイルやディレクトリもコピーする。

  * optionsには、`:preserve`、`:noop`、`:verbose`が指定できる

```ruby
FileUtils.cp_r(src, dst, options = {})
```



## 4.ファイル・ディレクトリの移動を行うコマンド

* `mv`：srcに指定したファイルまたはディレクトリをdestに移動する

  * destがディレクトリの場合は、dest以下に移動できる

* `move`：`mv`のエイリアスメソッド

  * optionsには、`:noop`と`:verbose`が指定できる

```ruby
FileUtils.mv(src, dest, options = {})
FileUtils.move(src, dest, options = {})
```



## 5.ファイル・ディレクトリの削除を行うコマンド

* `rm`：listに指定したファイルを削除する

  * 引数listには、単一のファイルパスまたは複数のファイルパスを含む配列を指定することができる

  * ディレクトリを指定すると、例外が発生する

  * optionsには`:force => true`(削除に失敗しても例外を発生させない)、`:noop`、`:verbose`が指定できる

* `remove`：`rm`メソッドのエイリアスメソッド

```ruby
FileUtils.rm(list, options = {})
FileUtils.remove(list, options = {})
```

* `rmdir`：listで指定したディレクトリを削除するメソッド

  * 引数listには単一のディレクトリパスまたは複数のディレクトリパスを含む配列を指定できる

  * ファイルを指定すると例外が発生する。optionsには、`:noop`、`verbose`が指定できる

```ruby
FileUtils.rmdir(list, options = {})
```

* `rm_r`：`rm`メソッドの再帰オプション`-r`で削除を行うコマンド

  * 単一または複数のファイルパスを引数listに指定することができる

  * ディレクトリを指定した場合は、ディレクトリ以下にあるファイルやディレクトリも削除される。

  * optionsには、`:force`、`:noop`、`:verbose`が指定できる

* `rm_rf`：`rm_r`メソッドのoptionsに、`:force => true`を付加して実行したもの

```ruby
FileUtils.rm_r(list, options = {})
FileUtils.rm_rf(list, options = {})
```



## 6.ファイル・ディレクトリの作成を行うコマンド

* `touch`：引数listで指定したファイルの最終変更時刻(mtime)と、アクセス時刻(atime)を変更する

  * 存在しないファイルを指定した場合は、空のファイルを作成する

  * optionsには、`:noop`と`verbose`が指定できる。

```ruby
FileUtils.touch(list, options = {})
```

* `mkdir`：listメソッドに指定した単一または複数の文字列で、ディレクトリを作成する。

  * optionsには、`:noop`と`:verbose`が指定できる。(`mkdir_p`も同様)

```ruby
FileUtils.mkdir(list, options = {})
FileUtils.mkdir_p(list, options = {})
```



## 7.ファイル・ディレクトリの権限を変更するコマンド

* `chown`：引数listに指定したファイルまたはディレクトリの所有権をuser、groupの権限の所有に書き換える

* `chmod`：引数listに指定したファイルまたはディレクトリのパーミッションをmodeに書き換える

  * optionsには、`:noop`と`:verbose`が指定できる。

```ruby
FileUtils.chown(user, group, list, options = {})
FileUtils.chmod(mode, list, options = {})
```



## 8.ファイル・ディレクトリへのシンボリックリンクを作成するコマンド

* `ln_s`：引数srcに指定したファイルまたはディレクトリのシンボリックリンク(別名でアクセス)を引数destに指定した名前で作成する。

  * destが存在するディレクトリを指定した場合は、destディレクトリ以下にsrcという名前のシンボリックリンクを作成する。

  * optionsには、`:force`(失敗した時に例外を発生させない、destが存在する場合は上書きする)、`:noop`、`:verbose`が指定できる

  * `symlink`メソッドは、`ln_s`メソッドのエイリアスメソッド

```ruby
FileUtils.ln_s(src, dest, options = {})
FileUtils.symlink(src, dest, options = {})
```



| 版     | 年/月/日   |
| ------ | ---------- |
| 初版   | 2018/10/20 |
| 第二版 | 2019/05/13 |
