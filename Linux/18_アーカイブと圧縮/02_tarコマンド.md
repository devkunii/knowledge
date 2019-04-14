02 tarコマンド
=============

* `tar`コマンド：アーカイブファイルを扱うためのコマンド

  * 複数のファイルやディレクトリをアーカイブファイルにまとめる

  * 逆にアーカイブファイルから元のファイルを取り出すために利用する



## 練習用ファイルの準備

* 練習用ファイルの作成

  ```bash
  $ mkdir dir1
  $ touch dir1/file-{1..5}.txt
  $ ls dir1
  file-1.txt	file-2.txt	file-3.txt	file-4.txt	file-5.txt
  ```

  * `ブレース展開`：連番のリストを生成することができる

  ```bash
  {<開始数値>..<終了数値>}
  ```

  * つまり、先ほどの`touch`コマンドは次のように展開されてから実行される

  ```bash
  touch dir1/file-1.txt	dir1/file-2.txt	dir1/file-3.txt	dir1/file-4.txt	dir1/file-5.txt
  ```

  * この他にも、文字列もブレース展開で扱うことができる

    * 例)aからeまでのリストをブレース展開で生成

  ```bash
  $ echo {a..e}.txt
  a.txt b.txt c.txt d.txt e.txt
  ```

  * また、ブレース内をカンマで区切ることで文字列のリストを生成できる

    * 例)拡張子`txt`、`log`、`dat`のファイルを生成する

  ```bash
  $ echo sample.{txt,log,dat}
  sample.txt sample.log sample.dat
  ```



## アーカイブファイルの作成

* ファイルをアーカイブするには、以下のコマンドを実行する

  ```bash
  tar cf <アーカイブファイル> <アーカイブファイル元パス>
  ```

  * `c`：新しくアーカイブファイルを作成するように指定

  * `f`：ファイルを意味する

    * `f <アーカイブファイル>`として、新しく作成するアーカイブファイル名を指定する

    * `tar`コマンドでファイルを引数に取る場合には、必ずこの`f`を付ける

  * 例)`dir1`というディレクトリと、その下にある5つのファイルをアーカイブファイルにまとめる

  => アーカイブファイルは`dir1.tar`

  ```bash
  $ tar cf dir1.tar dir1
  ```

> `tar`コマンドのオプソhんは、ハイフンを付けても付けなくても構わない



## アーカイブファイルの内容確認

* `-t`オプション：作成したアーカイブファイルの内容を確認する

  * `list`の`t`から名称が付けられている

  * `c`オプションでファイルをアーカイブしたら、目的のファイルが正しくアーカイブされているかどうか、`t`オプションで確認する習慣を付けておく

  ```bash
  $ tar tf dir1.tar
  dir1/
  dir1/file-5.txt
  dir1/file-4.txt
  dir1/file-3.txt
  dir1/file-2.txt
  dir1/file-1.txt
  ```



## アーカイブの展開

* `x`オプション：アーカイブファイルを展開して元のファイルやディレクトリを取り出す

  * `extract`の`x`から名称が付けられている

  ```bash
  tar xf <アーカイブファイル>
  ```

  * 例)`dir1`ディレクトリを一旦削除してから、アーカイブファイルを展開する

  ```bash
  $ rm -rf dir1
  $ tar xf dir1.tar
  $ ls dir1
  file-1.txt	file-2.txt	file-3.txt	file-4.txt	file-5.txt
  ```

  * なお、`tar`で展開されるファイル名と同じファイル名が存在している場合、そのファイルが上書きされてしまうことに注意する



## ファイルリストを表示するvオプション

* `v`オプション：`c`オプションによるアーカイブ作成や、`x`オプションによる展開時に、対象となったファイルリストを表示する

  ```bash
  $ tar cvf dir1.tar dir1
  a dir1
  a dir1/file-5.txt
  a dir1/file-4.txt
  a dir1/file-3.txt
  a dir1/file-2.txt
  a dir1/file-1.txt
  ```

* `t`オプション：アーカイブファイル確認で`v`オプションを併用すると、ファイル名だけでなくファイル属性が合わせて表示される

  ```bash
  $ tar tvf dir1.tar
  drwxr-xr-x  0 MacUser staff       0  4 14 11:51 dir1/
  -rw-r--r--  0 MacUser staff       0  4 14 11:51 dir1/file-5.txt
  -rw-r--r--  0 MacUser staff       0  4 14 11:51 dir1/file-4.txt
  -rw-r--r--  0 MacUser staff       0  4 14 11:51 dir1/file-3.txt
  -rw-r--r--  0 MacUser staff       0  4 14 11:51 dir1/file-2.txt
  -rw-r--r--  0 MacUser staff       0  4 14 11:51 dir1/file-1.txt
  ```



## ファイル属性の保持

* `tar`コマンドは単にファイルをコピーするだけでなく、ファイルのパーミッションやオーナー、タイムスタンプなどのファイルの属性もそのままアーカイブする

* しかし、所有者がrootでオーナーのみにしか読み込みパーミッションが付いているファイルでは、一般ユーザではアーカイブできない

* ディレクトリツリーの完全なバックアップを取得するならば、root権限で`tar`コマンドを実行することが必要になる



| 版 |  年/月/日 |
|----|----------|
|初版|2019/04/14|
