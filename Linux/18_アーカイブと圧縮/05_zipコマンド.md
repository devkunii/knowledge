05 zipコマンド
============

* `zip`コマンド：アーカイブと圧縮を同時に行うコマンド

  * 複数のファイルをまとめて1つのファイルに圧縮できる

  * zipで圧縮したファイルには慣習として`.zip`という拡張子を付け、これを「zipファイル」と呼ぶ



## zipファイルの作成

* `zip`コマンドでファイルやディレクトリを圧縮するには、以下のように記述する

  ```bash
  $ zip -r <圧縮ファイル名> <圧縮対象パス>
  ```

  * 例)`dir1`というディレクトリを、`dir1.zip`というファイルに圧縮する

  ```bash
  $ zip -r dir1.zip dir1
  adding: dir1/ (stored 0%)
  adding: dir1/file-5.txt (stored 0%)
  adding: dir1/file-4.txt (stored 0%)
  adding: dir1/file-3.txt (stored 0%)
  adding: dir1/file-2.txt (stored 0%)
  adding: dir1/file-1.txt (stored 0%)
  ```

  * `r`オプションは、指定したディレクトリの下に含まれるファイルもまとめて圧縮するためのオプション

    => zipファイルを作成する時は、常に`-r`オプションを利用する

* `zipinfo`コマンド：作成した圧縮ファイルの内容を確認する

  ```bash
  $ zipinfo dir1.zip
  Archive:  dir1.zip
  Zip file size: 950 bytes, number of entries: 6
  drwxr-xr-x  3.0 unx        0 bx stor 19-Apr-14 11:51 dir1/
  -rw-r--r--  3.0 unx        0 bx stor 19-Apr-14 11:51 dir1/file-5.txt
  -rw-r--r--  3.0 unx        0 bx stor 19-Apr-14 11:51 dir1/file-4.txt
  -rw-r--r--  3.0 unx        0 bx stor 19-Apr-14 11:51 dir1/file-3.txt
  -rw-r--r--  3.0 unx        0 bx stor 19-Apr-14 11:51 dir1/file-2.txt
  -rw-r--r--  3.0 unx        0 bx stor 19-Apr-14 11:51 dir1/file-1.txt
  6 files, 0 bytes uncompressed, 0 bytes compressed:  0.0%
  ```

* `unzip`コマンド：圧縮したzipファイルを展開する

  ```bash
  $ unzip dir1.zip
  Archive:  dir1.zip
  replace dir1/file-5.txt? [y]es, [n]o, [A]ll, [N]one, [r]ename: y
   extracting: dir1/file-5.txt         
  replace dir1/file-4.txt? [y]es, [n]o, [A]ll, [N]one, [r]ename: y
   extracting: dir1/file-4.txt         
  replace dir1/file-3.txt? [y]es, [n]o, [A]ll, [N]one, [r]ename: y
   extracting: dir1/file-3.txt         
  replace dir1/file-2.txt? [y]es, [n]o, [A]ll, [N]one, [r]ename: y
   extracting: dir1/file-2.txt         
  replace dir1/file-1.txt? [y]es, [n]o, [A]ll, [N]one, [r]ename: y
   extracting: dir1/file-1.txt
  ```

* `q`オプション：ファイル名を表示しないオプション

  * `unzip`コマンドも同様

  ```bash
  $ zip -rq dir1.zip dir1
  ```



### パスワード付きzipファイル

* `e`オプション：パスワード付きのzipファイルを作成する

  ```bash
  $ zip -er dir1.zip dir1
  Enter password:
  Verify password:
  updating: dir1/ (stored 0%)
  updating: dir1/file-5.txt (stored 0%)
  updating: dir1/file-4.txt (stored 0%)
  updating: dir1/file-3.txt (stored 0%)
  updating: dir1/file-2.txt (stored 0%)
  updating: dir1/file-1.txt (stored 0%)
  ```

  * パスワード付きzipファイルは、`unzip`コマンドで展開できる

  ```bash
  $ unzip dir1.zip
  Archive:  dir1.zip
  [dir1.zip] dir1/file-5.txt password:
  replace dir1/file-5.txt? [y]es, [n]o, [A]ll, [N]one, [r]ename: y
   extracting: dir1/file-5.txt         
  replace dir1/file-4.txt? [y]es, [n]o, [A]ll, [N]one, [r]ename: y      
   extracting: dir1/file-4.txt         
  replace dir1/file-3.txt? [y]es, [n]o, [A]ll, [N]one, [r]ename: y
   extracting: dir1/file-3.txt         
  replace dir1/file-2.txt? [y]es, [n]o, [A]ll, [N]one, [r]ename: y
   extracting: dir1/file-2.txt         
  replace dir1/file-1.txt? [y]es, [n]o, [A]ll, [N]one, [r]ename: y
   extracting: dir1/file-1.txt  
  ```



| 版 |  年/月/日 |
|----|----------|
|初版|2019/04/14|
