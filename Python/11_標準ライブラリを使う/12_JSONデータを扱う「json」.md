12 JSONデータを扱う「json」
========================

* JSON形式のデータをPythonで扱うときには、`json`モジュールを利用する

* `json`モジュールは、

  * Pythonの組み込み型のデータと、JSONデータの相互変換をする機能を提供

  * 数値、文字列、リスト、ディクショナリといったデータをJSON形式に変換

  * JSON形式のデータをPythonのデータに変換

  することができる

* Pythonの`json`モジュールが提供するインターフェースは、pickleモジュールの設計によく似ている

  => pickleの使い方を知っていると、jsonも抵抗なく使える



## JSONをPythonのデータ型に変換する

* `loads()`：JSON文字列をPythonオブジェクトに変換する

  * JSON文字列をPythonのデータ型に変換し、戻り値として返す

  ```python
  json.loads(JSON文字列)
  ```

| JSONの型     | Pythonの型     |
| :----------- | :------------- |
| object       | ディクショナリ |
| array        | リスト         |
| string       | 文字列         |
| number(int)  | 整数型         |
| number(real) | float型        |
| true         | True           |
| false        | False          |
| null         | None           |

* `load()`：JSON文字列を含むファイルをPythonオブジェクトに変換する

  * ファイルまたはファイル風オブジェクトにあるJSON文字列を、Pythonのデータ型に変換し、戻り値として返す

  ```python
  json.load(ファイル)
  ```



## Pythonのデータ型をJSONに変換する

* `dumps()`：PythonオブジェクトをJSON文字列に変換する

  * 引数に渡したオブジェクトをJSON文字列に変換する

  * JSONに変換できない型のオブジェクトを与えると例外(TypeError)を発生する

  * `skipkeys`にTrueを渡すと、日本語のような非ASCII文字をエスケープして出力する(デフォルトの動作)

  * `ensure_ascii`にFalseを渡すと、非ASCII文字はそのまま出力される

  * JSON文字列のエンコードを指定するには、`ensure_ascii`にFalseを渡してJSON文字列を生成し、`encode()`を使って目的のエンコードに変換したバイト型文字列を得る

  ```python
  json.dumps(Pythonオブジェクト[, オプションの引数...])
  ```

* `dump()`：JSON文字列に変換してファイルに書き出す

  * 引数に渡したオブジェクトをJSON文字列に変換し、指定されたファイルに書き出す

  ```python
  json.dump(Pythonオブジェクト, ファイル)
  ```



## jsonの使用例

* Webから読み込んだデータをPythonのデータ型に変換する

* 例)GitHubから読み込み

```python
from urllib.request import urlopen
from json import loads

url = 'https://api.github.com/users/gvanrossum/repos'
body = request.urlopen(url).read()
body = body.decode('utf-8')
repos = loads(body)
for r in repos:
  print(r['name'])
```



| 版 |  年月日   |
|---|----------|
|初版|2019/02/09|
