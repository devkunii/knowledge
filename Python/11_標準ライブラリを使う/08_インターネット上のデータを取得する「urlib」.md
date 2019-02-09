08 インターネット上のデータを取得する「urllib」
========================================

* `urllib`モジュール：インターネット上のデータを取得するために利用する

  * WebやFTPを使ってデータを取得したり、データをPOSTしたりしてCGIやWebサービスを操作できる

* WebやFTP上のデータをファイルに保存したいだけならば、`request.()`関数を利用する

  => URLを指定して関数を呼び出すと、ファイルの作成から保存までを実行してくれる

* URLで指定されたサイトからデータを取得してPythonで処理するには、`request.urlopen()`関数を使う

  * 戻り値としてファイル風のオブジェクトを返す

  * 返ってきたオブジェクトに対して、`read()`メソッドなどを使って読み込んで、文字列にコピーして利用する

  * プロキシー(Proxy)を使ったアクセスをすることもできる

  * ネットワークを利用した処理では、完了するまで処理を中断する

  * ネットワークが繋がっていないなどの理由で接続先から応答が返ってこない場合は、処理が中断してしまうので注意する

* CGIやWebサービスなどにデータをPOSTしたい時は、`parse.urlencode()`と`request.urlopen()`を組み合わせて使う

  * POSTするデータをディクショナリなどにして`parse.urlencode()`に渡してデータを作成し、

  * `urlopen.urlopen()`を呼び出すときに引数として添えてデータをPOSTする



## WebやFTPからファイルを取得する

* `urlretrieve()`：URLを指定してファイルを取得する

  * WebやFTPから、URLを指定してファイルを取得できる

  * 保存用のファイル名はオプションになっているが、省略すると **一時ファイル用のディレクトリ** に保存しようとする

    => 実用上はファイル名を指定して利用する

  * 指定したURLのデータのみを取得する(HTMLや書かれた画像などは取得しない)

  * この関数は、戻り値として2つの値をタプルとして返す

    * 1つ目の戻り値は、保存したファイルのパス

    * 2つ目の戻り値は、レスポンスのヘッダ情報を取り出すのに利用する

  * `urlopen()`が返すオブジェクトに対して、`info()`を呼び出して返ってくるオブジェクトと同じ

  * POST用のデータを引数として与えると、URLに対してPOSTリクエストを送ることができる

  ```python
  urllib.request.urlretrieve(url[, ファイル[, POST用のデータ]])
  ```



## `request.urlretrieve()`を使ったサンプルコード

* `request.urlretrieve()`を使って、Web上のデータをファイルに保存する

* `request`モジュールの他、`urllib`モジュールに含まれる`parse`モジュールをインポートしている

* URLのパスをスラッシュ`/`で区切り、簡易にファイル名を取り出すために利用している

```python
from urllib import request
from urllib import parse
url = 'http://dname.com/somefile.zip'

# URLを分割して、ファイル名を取得
filename = parse.urlparse(url)[2].split('/')[-1]

# ファイル名を確認
filename
>>> 'somefile.zip'

# ファイル名を取得、カレントディレクトリに保存
request.urlretrieve(url, filename)
>>> ('somefile.zip', <http.client.HTTPMessage object at ox1012c8190)
```



## WebやFTPからデータを読み込む

* `urlopen()`：URLから取得したデータをオブジェクトで返す

  * WebやFTPからデータを取得してPythonで処理をしたい場合に、`request.urlopen()`を使うと便利

  * この関数は、インターネットから取得したデータを、読み込み可能な **ファイル風のオブジェクト** に格納して返す

  * 返ってきたオブジェクトは、`read()`や`readlines()`などの読み込みを行う処理に限って、`open()`関数が返すオブジェクトと同様に扱うことができる

    => 読み込まれたデータはバイト型となる

    => 文字列として処理をするためには、エンコードを指定して文字列型に変更する必要がある

  * `urlopen()`が返すファイル風のオブジェクトには、`シーク位置`があることに注意する

    => 一度`read()`などを使って読み込みを行なった上で、再度読み込みを行なっても、`シーク位置`が末尾にセットしているので、データは返ってこない

  * `urlopen()`に第2引数(POST用のデータ)を与えると、URLに対してPOSTリクエストを送ることができる

  * 第3引数に秒数を指定すると、一定時間応答がない場合に処理を中止できる

  ```python
  request.urlopen(url[, POST用のデータ[, タイムアウト]])
  ```



## `urlopen()`が返すオブジェクトで利用できるメソッド

* `urlopen()`が返すオブジェクトに対しては、以下のメソッド呼び出しを行うことができる

* `F`は`urlopen()`が返すファイル風オブジェクト、`[]`内の引数はオプション(省略可能)

* `read()`：データを連続的に読み込む

  * データを読み込み、文字列として返す

  * 引数としてサイズを指定しないと、データを最後まで読み込む

  * ファイルオブジェクトの`read()`と同じ

  ```python
  F.read([整数のサイズ])
  ```

* `readline()`：データから1行読み込む

  * データから1行を読み込んで、文字列を返す

  * ファイルオブジェクトの`readline()`と同じ

  ```python
  F.readline([整数のサイズ])
  ```

* `readlines()`：行単位で連続的に読み込む

  * データから複数行を読み込む

  * 文字列を要素として含んだリストを返す

  * ファイルオブジェクトの`readlines()`と同じ

  ```python
  F.readlines([整数のサイズ])
  ```

* `geturl()`：URLを返す

  * 取得したデータのURLを返す

  * `urlopen()`は、HTTPヘッダを使ったりリダイレクトをサポートしている

  * 引数として与えたURLがリダイレクトされた場合、リダイレクトした先のURLを知りたいときに利用すると便利

  ```python
  F.geturl()
  ```

* `info()`：メタ情報を返す

  * データを取得したときに送られてきたメタ情報を持つオブジェクトを返すメソッド

  * レスポンス時に送られたヘッダ情報を取得する場合に利用する

  * このメソッドが返すオブジェクトは、ディクショナリのようにアクセスできる

  * `keys()`を使ってプロパティ名の一覧を取得したり、`R.info()['content-length']`のようにしてヘッダの値を取得する

  * `info()`メソッドの返すオブジェクトの実体は、`mimetools.Message`クラスなど、URLスキーマに対応したクラスのインスタンス



## BASIC認証

* Web上のファイルなどにアクセスをするとき、BASIC認証を利用したい場合がある

* 簡易に済ませたいならば、`request.urlopen()`や`request.urlretrieve`に引数として渡すURLにユーザ名とパスワードを埋め込むと良い

```python
http://ユーザ名:パスワード@example.com/foo/bar.html
```

* ただし、このようにすると認証に必要なソースコードが埋め込まれてしまうので注意する

* `URLopener`や`FancyURLopener`というクラスを継承したクラスを利用する方法もある



## データをPOSTする

* `urlencode()`：ディクショナリからクエリ文字列を作る

  * ディクショナリまたはシーケンスからURLエンコードされたクエリ文字列を作る

    => クエリ文字列は、「キー=値」のペアを「`&`」で繋げたもの(`?`は含まれない)

  * スペースや日本語のようなマルチバイト文字列は、「`%`」で始まる文字列に変換される

  * `request.urlopen()`、`request.urlretrieve()`を使い、POSTメソッドでCGIやWebサービスなどにデータを送信したいときに利用すると便利

  * 引数には、ディクショナリまたはキーワードと値2つの要素を持ったシーケンスを渡す

  * シーケンスを渡す場合には、引数`doseq`にTrueを指定する

  * 引数`safe`にはURLエンコードから除外する文字列を指定する(デフォルトでは何も設定されていない)

  * 引数`encoding`にはディクショナリやシーケンスに含まれる要素の文字コードを指定する(デフォルトでは`utf-8`)

  * 引数`errors`には、`encoding`で指定した文字コードがサポートされていない文字が出現した時の対処法を指定する(デフォルトは`strict`)

  * この関数はURLエンコードに関して`quote_plus()`関数と同じ処理を行う(空白文字列は半角のプラス`+`に変換される)

  ```python
  parse.urlencode(ディクショナリ、またはシーケンス[, doseq[, safe[, encoding[, errors]]]])
  ```

* 例)`urlopen()`でデータをPOSTする

```python
from urllib import request
from urllib import parse

# クエリのディクショナリを作る
postdic = {'name': 'someone', 'email': 'foo@bar.com'}

# ディクショナリを変換
postdata = parse.urlencode(postdic)

# 内容を確認
postdata
>>> 'name=someone&email=foo%40bar.com'

# データを指定してPOST
file = request.urlopen('http://service.com/process.cgi', postdata)
```



## その他の関数

* `quote()`：文字列をURLエンコードする

  * 引数として渡した文字列をURLエンコードして返す

  * アルファベットと数字、および「`_`」、「`-`」は変換を行わない

  * 引数`encoding`、`errors`については、`urlencode()`関数と同じだが、`quote()`メソッドの場合、`safe`のデフォルト値は`/`になっている

  * 例)Google検索用URLの先頭部分に検索文字列をURLエンコードして追加する

  ```python
  >>> from urllib import parse
  >>> url = "https://www.google.com/webhp?ie=UTF-8#q="
  >>> url += parse.quote('python サンプルコード')
  >>> url
  'https://www.google.com/webhp?ie=UTF-8#q=python%20%E3%82%B5%E3%83%B3%E3%83%97%E3%83%AB%E3%82%B3%E3%83%BC%E3%83%89'
  ```

* `quote_plus()`：文字列をURLエンコードする

  * `quote()`の処理に加えて、空白文字列を半角のプラス記号`+`に置き換える

  * HTMLフォームに入力されたデータのように空白を含む可能性がある場合はこちらを使用すると良い

  * `quote()`関数と異なり、引数`safe`にはデフォルト値が設定されていない

  ```python
  parse.quote_plus(文字列[, safe[, encoding[, errors]]])
  ```

* `uniquote()`：文字列をURLデコードする

  * URLエンコードされた文字列を、通常の文字列に変換する

  * `quote()`関数の逆の処理を行う

  * `encoding`には変換先の文字コードを指定し、省略すると「`utf-8`」になる

  * また、`errors`は変換時にエラーが生じた場合の処理方法を指定する

    => デフォルトは`replace`

  ```python
  parse.uniquote(文字列[, encoding[, errors]])
  ```

* `uniquote_plus()`：文字列をURLデコードする

  * `uniquote()`関数と同じ処理を行うが、半角のプラス記号`+`は空白に変換される

  ```python
  parse.uniquote_plus(文字列[, encoding[, errors]])
  ```



| 版 |  年月日   |
|---|----------|
|初版|2019/02/09|
