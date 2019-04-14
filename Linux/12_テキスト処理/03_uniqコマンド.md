03 uniqコマンド
==============

* `uniq`コマンド：連続した同じ内容の行を省くコマンド

  * IDリストなどのファイルから重複を取り除くためによく利用される

  * 例)重複行を含む`file2`というファイルがあるとする

  ```txt
  Hokkaido
  Hokkaido
  Aomori
  Iwate
  Iwate
  Iwate
  Miyagi
  ```

  * このファイルに対して`uniq`コマンドを使うと、次のように重複していた「Hokkaido」と「Iwate」が1行だけになる

  ```bash
  $ uniq file2
  Hokkaido
  Aomori
  Iwate
  Miyagi
  ```

* なお、`uniq`コマンドは **同じ内容の行が連続している場合にのみ重複を取り除く** ことに注意する

  * 例)「Hokkaido」が離れた位置にあるため、`uniq`コマンドを実行後も重複行が残っている

  ```txt
  Hokkaido
  Hokkaido
  Aomori
  Miyagi
  Hokkaido
  ```

  * 重複行が連続していない場合は、取り除かれない

  ```bash
  $ uniq file3
  Hokkaido
  Aomori
  Miyagi
  Hokkaido
  ```

  * このような場合には、一旦`sort`コマンドを使って対象ファイルの内容を並び替えておくと、ファイル全体から重複行を取り除くことができる

  ```bash
  $ sort file3 | uniq
  Aomori
  Hokkaido
  Miyagi
  ```

* なお、`sort`コマンドでは、重複行を一度しか表示しない`-u`オプションを利用することで、`uniq`コマンドを使わずに重複行を取り除くことができる

  ```bash
  $ sort -u file3
  Aomori
  Hokkaido
  Miyagi
  ```



### 重複行を数える

* `-c`オプション：重複している行数を数えて表示するためのオプション

  * 対象が何らかのデータを記録しているファイルならば、各データがそれぞれ何件あるかを表示することができる

  * 例)ファイル中に「Hokkaido」という行が2行、「Iwate」とうい行が3行あることを`-c`オプションで表示している

  ```bash
  $ sort file2 | uniq -c
   1 Aomori
   2 Hokkaido
   3 Iwate
   1 Miyagi
  ```

  * この`uniq -c`の出力は、さらにパイプで`sort`コマンドに渡す手法がよく使われる

    * こうすると、`uniq -c`で表示される行頭の出現回数でソートされるので、結果として重複行の多い順(`-rn`)、あるいは少ない順(`-n`)に表示することができる

    ```bash
    $ sort file2 | uniq -c | sort -rn
    3 Iwate
    2 Hokkaido
    1 Miyagi
    1 Aomori
    ```

    ```bash
    $ sort file2 | uniq -c | sort -n
    1 Aomori
    1 Miyagi
    2 Hokkaido
    3 Iwate
    ```

* このように`sort`コマンドと`uniq`コマンドを組み合わせることで、データファイルから簡単に出現回数のランキングを作成することができる



| 版 |  年/月/日 |
|----|----------|
|初版|2019/03/11|
