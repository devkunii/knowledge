08 uriモジュール
===============

## 目次

* [uriモジュールとは](#1uriモジュールとは)

* [parse/URI(uri_str)](#2parseURIuri_str)

* [split](#3split)

* [encode_www_form_component](#4encode_www_form_component)

* [decode_www_form_component](#5decode_www_form_component)



## 1.uriモジュールとは

* URIを扱うライブラリ(現行のRubyのuriモジュールはURLのみ対応)

* URLはネットワーク上のリソースの場所を示し、スキームとスキーム毎に定められた表現形式を持つ

![URIの例](./images/6-5/URIの例.jpg)

* Rubyでは以下のスキームに対応しており、それ以外のスキームには汎用のパーサが使われる

* これらのスキームのパーサは、URIモジュール以下に別々のクラスとして定義されている

  > どのパーサを使用するかは、URIモジュールが与えられたスキームから決定するので、個別に呼び出す必要はない

![Rubyが対応しているスキーム](./images/6-5/Rubyが対応しているスキーム.jpg)



## 2.parse/URI(uri_str)

* `URI.parse`：与えられたuri_strから該当するURIサブクラスのインスタンスを生成して返す

例)

```ruby
require 'uri'

>> uri = URI.parse "https://docs.ruby-lang.org/ja/man/html/index.html"

>> p uri
=> #<URI::HTTPS https://docs.ruby-lang.org/ja/man/html/index.html>
>> p uri.scheme
=> "https"
>> p uri.host
=> "docs.ruby-lang.org"
>> p uri.port
=> 443
>> p uri.path
=> "/ja/man/html/index.html"
```

* スキーム：httpsなので、内部では`URI::HTTPS`クラスのパーサが個々の要素に分割する

* HTTPSのパーサは、RFCの企画にしたがって厳密にパースするため、ホスト名として使用できない文字が含まれていると、正しくパースしない

* URI(uri_str)メソッドは、uriロード時にKernel内に定義され、内部でURI.parse(uri_str)を呼び出す

例) `URI.parse`で正しくパースできない例

> パースできてしまう

```ruby
require 'uri'

uri = URI.parse "https://docs.ruby_lang.org/ja/man/html/index.html"

>> p uri
=> #<URI::HTTPS https://docs.ruby_lang.org/ja/man/html/index.html>
>> p uri.scheme
=> "https"
>> p uri.host
=> "docs.ruby_lang.org"
>> p uri.port
=> 443
>> p uri.path
=> "/ja/man/html/index.html"
```



## 3.split

* URLを以下の要素に分割した配列を返す

* 該当しない場合は、nilが入る

```ruby
URI.split(url)

# 配列
[scheme, userinfo, host, port, registory, path, opaque, query, fragment]
```

例)

```ruby
require 'uri'

uri = URI.split "https://docs.ruby_lang.org/ja/man/html/index.html"

p uri

# 実行結果
["https", nil, "docs.ruby_lang.org", nil, nil, "/ja/man/html/index.html", nil, nil, nil]
```



## 4.encode_www_form_component

* URIで使用できるASCII以外の文字列データを扱う場合、16進数で表記したバイトコードを「%xx」という形で表記することが定められている

  > これを「URIエンコード」という

* `URI.encode_www_form_component`：引数strで指定した文字列をURLエンコードした文字列で返す

  > マルチバイト文字を使用したURIをエンコードする場合は、デフォルトではそのプログラムが書かれた文字コードによってエンコードされる

```ruby
URI.encode_www_form_component(str)
```

* エンコードを指定する場合は、文字列を変換する必要がある

* どの文字コードを用いるかは、通信先のサーバーなどの実装によって異なる

```ruby
require 'uri'

>> p "http://www.example.com/" + URI.encode_www_form_component("Ruby技術者試験対策教科書")
=> "http://www.example.com/Ruby%E6%8A%80%E8%A1%93%E8%80%85%E8%A9%A6%E9%A8%93%E5%AF%BE%E7%AD%96%E6%95%99%E7%A7%91%E6%9B%B8"
>> p "http://www.example.com/" + URI.encode_www_form_component("Ruby技術者試験対策教科書".encode("EUC-JP"))
=> "http://www.example.com/Ruby%B5%BB%BD%D1%BC%D4%BB%EE%B8%B3%C2%D0%BA%F6%B6%B5%B2%CA%BD%F1"
```



## 5.decode_www_form_component

* URLエンコードされた文字列を元の文字列に戻す

```ruby
URI.decode_www_form_component(str)
```

例)

```ruby
require 'uri'

uri = "http://www.example.com/Ruby%E6%8A%80%E8%A1%93%E8%80%85%E8%A9%A6%E9%A8%93%E5%AF%BE%E7%AD%96%E6%95%99%E7%A7%91%E6%9B%B8"

parsed_uri = URI.parse(uri)

p URI.decode_www_form_component(parsed_uri.path[1..-1])

# 実行結果
$ ruby 6-5.sample.rb
"Ruby技術者試験対策教科書"
```



| 版     | 年/月/日   |
| ------ | ---------- |
| 初版   | 2018/10/21 |
| 第二版 | 2019/05/13 |
