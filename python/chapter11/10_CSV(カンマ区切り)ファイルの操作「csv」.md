10 CSV(カンマ区切り)ファイルの操作「csv」
===================================

* `csv`は、複数の要素をカンマ`,`で区切って並べたCSVファイルを扱うためのモジュール

  * 表計算ソフトやデータベースからファイルを書き出す際によく利用されるファイル形式

  * CSVファイルの改行や区切り文字などの形式に標準は存在しない

    => アプリケーションによって異なる形式のファイルを利用している

* `csv`モジュールでは、`dialect`という仕組みを使って、異なる形式のCSVファイルを扱う

  * Excelの書き出すファイル形式に対応した`dialect`が用意されていて、利用できる

  * 独自の`dialect`も作ることができる

* `reader()`：CSVファイルのデータを連続的に読み込む

  * CSVファイルをファイルオブジェクトとして指定して呼び出す

    => `readerオブジェクト`と呼ばれるオブジェクトを返す

    * `reader`オブジェクトでは、for文に添えるなどして利用する

    * `reader`オブジェクトはイテレータの一種で、CSVファイルを1行ずつ読み込んで処理を進める

    * `reader`オブジェクトがCSVファイルの行を読み込むと、要素に分割してリストに格納して返す

  * `dialect`という引数はオプション

    * 引数を指定しないと、ExcelのCSVファイルを読み込むための設定を利用する

    * `csv`には他に、「`excel-tab`」というExcelのタブ区切りファイルを読み込む設定が内蔵されている

  ```python
  csv.reader(ファイルオブジェクト[, dialect])
  ```

  * 例)`reader`関数でCSVファイルを開く

  ```python
  import csv
  csvfile = open("text.csv", encoding="utf-8")
  for row in csv.reader(csvfile):
    print(row)
  ```

* `writer()`：`writer`オブジェクトを返す

  * 要素を設定に従って区切ったCSVファイルを書き出すために利用する`writerオブジェクト`を返す関数

  * 書き込みができるモードで開いたファイルオブジェクトを引数に与えて呼びだす

  * オプションの引数`dislect`には、CSVファイルの形式を指定する


  ```python
  csv.writer(ファイルオブジェクト[, dialect])
  ```

  * 実際の書き込み処理は、`writerオブジェクト`の以下のメソッドを使用する

    => `W`：`writerオブジェクト`を表す

* `writerow()`：1行書き込む

  * シーケンスに引数を与えると、要素を`writerオブジェクト`を作るときに指定したファイルオブジェクトに対して書き込む

  ```python
  W.writerow(シーケンス)
  ```

* `writerows()`：複数行書き込む

  * シーケンスを引数に渡すと、要素を区切って複数の行を書き込む

  ```python
  W.writerows(シーケンス)
  ```

* 例)`writer()`関数で1行ずつ書き出す

```python
import csv
csvfile2 = open("test2.csv", "w", encoding="utf-8")

writer = csv.writer(csvfile2)
for row in seq:
  writer.writerow(row)
```



| 版 |  年月日   |
|---|----------|
|初版|2019/02/09|
