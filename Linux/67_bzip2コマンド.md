04 bzip2コマンド
===============

* `bzip2`コマンド：`gzip`コマンドと同様に、ファイルの圧縮・展開を行うためのコマンド

  * bzip2形式は、gzip形式よりも圧縮率が高く、データ量をより小さくできる

  * ただし、圧縮や展開にはgzipより長い時間がかかる

  ```bash
  bzip2 <圧縮元ファイル>
  ```

  * 例)bzip2コマンドで圧縮する

  ```bash
  $ bzip2 ps.txt
  $ ls -l
  -rw-r--r--   1 MacUser  staff  10208  4 14 13:28 ps.txt.bz2
  ```

* `d`オプション：圧縮したファイルを元に戻す

  * `bunzip2`コマンドも同様の動作を行う

  ```bash
  $ bzip2 -d ps.txt.bz2
  ```

* `c`オプション：標準出力に出力する

  ```bash
  $ bzip2 -c ps.txt > ps_test.txt.bz2
  ```



## tarとbzip2を組み合わせる

* `j`オプション：`tar`コマンドでbzip2形式でファイル圧縮する

  ```bash
  $ tar cjf dir1.tar.bz2 dir1
  ```



| 版 |  年/月/日 |
|----|----------|
|初版|2019/04/14|
