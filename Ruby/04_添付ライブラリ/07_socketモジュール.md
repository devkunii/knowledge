07 socketモジュール
=================

## 目次

* [Rubyのネットワーク関連ライブラリ](#0Rubyのネットワーク関連ライブラリ)

* [socketモジュールとは](#1socketモジュールとは)

* [TCPを扱うクラス](#2TCPを扱うクラス)

* [UDPを扱うクラス](#3UDPを扱うクラス)

* [プロセス間通信を制御するクラス](#4プロセス間通信を制御するクラス)



## 0.Rubyのネットワーク関連ライブラリ

Rubyのネットワーク関連のライブラリは、以下のものなどを備えている

* socket：プロセス間、ホスト間の通信を行うライブラリ

* net/http：標準的なネットワーク(HTTP)のプロトコルを扱うライブラリ

* uri：ネットワーク内のリソースの場所を示すURIを扱うモジュール



## 1.socketモジュールとは

* プロセス間通信や、ホスト間通信を実現する通信ソケットを扱うライブラリ

* IOクラスを継承したクラス構成になっており、IOクラスと同じインターフェース(puts、getsなど)でデータの送受信を行うことができる

* TCPSocket・TCPServerクラス：TCP通信を行うプログラムを作成するのに適したインターフェースを提供

* UDPSocketクラス：UDP通信を行うプログラムを作成するのに適したインターフェースを提供

* UNIXSocket・UNIXServerクラス：UNIXドメインソケットによるプロセス間通信を行うプログラムを作成するのに適したインターフェースを提供する

* Socketクラス：汎用ソケットへのインターフェースを提供する

![socketライブラリのクラス](./images/6-5/socketライブラリのクラス.jpg)



## 2.TCPを扱うクラス

* `TCPServer`クラス：TCPをサーバー用途で扱う場合に利用する

```ruby
TCPServer.new([host, ]service)
TCPServer.open([host, ]service)
accept
```

例) 10080ポートへの接続用のインスタンスを作成し、`accept`メソッドでクライアントのアクセスを待つ。

> クライアントからの接続があれば、クライアントに対して文字列を送信する

```ruby
require 'socket'

server = TCPServer.new 10080
loop {
  client = server.accept
  client.puts "Hello TCPServer."
  client.close
}
```

* `TCPSocket`クラス：TCPをクライアント用途で扱う

```ruby
TCPSocket.open(host, service [, local_host, local_service])
TCPSocket.new(host, service [, local_host, local_service])
```

例) 先ほどの10080ポートのサーバーに接続し、文字列"Hello TCPServer"を取得する

```ruby
$ ruby 6-5.sample2.rb
Hello TCPServer.
```



## 3.UDPを扱うクラス

* `UDPSocket`クラス：サーバー用、クライアント用の区別はなく使用される。

  * UDPはコネクションレスなデータの送信方式

  * TCPに比べて伝達確認や通信ミスを検出しない分、オーバーヘッドが少なく、素早い通信ができる

  * 途中でデータの送信ミスが起きても大きな影響がない動画配信技術などに使われている



### new/open

* UDPSocketクラスのインスタンスの生成を行う。

* 引数のsocktypeには、アドレスファミリーと呼ばれるネットワークアドレスの種類を指定する

  * デフォルトでは、`Socket::AF_INET`が使用される

|      引数       | アドレスの種類 |
|:---------------|:-------------|
|Socket::AF_INET |IPv4ネットワーク|
|Socket::AF_INET6|IPv6ネットワーク|

```ruby
UDPSocket.new([socktype])
UDPSocket.open([socktype])
```



### bind

* 定義したソケットをホストのポートに関連付ける

```ruby
bind(host, port)
```



### recv

* BasicSocketに定義されているメソッド

* ソケットからデータを受け取り、文字列として返す

* 引数には、受け取るデータの長さを指定する

  * flagsには、受信データ処理のオプションを指定でき、以下のものを指定

    * `Socket::MSG_OOB`(帯域外データを送信する)

    * `Socket::MSG_EOR`(レコードの終了を指示)

  * オプションを指定しない場合には、デフォルト値は0

```ruby
recv(maxlen[, flags])
```

例) UDPSocketを使って10000ポートでUDP通信を待ち受ける

```ruby
require 'socket'

MAX_PACKET = 1024
socket = UDPSocket.new
socket.bind("0.0.0.0", 10000)
print socket.recv(MAX_PACKET)
```



### connect/send

UDPでデータを送信するには、

1. connectメソッドで接続先ホストとポートを指定して、sendメソッドを使用する方法

2. sendメソッドの引数の中に、接続先ホストとポートを指定する方法

* sendメソッドの引数の`dest_sockaddr`は、`Socket.pack_sockaddr_in`を使って生成したソケットアドレス構造体を指定

* sendメソッドの引数のflagsには、送信データのオプションを指定できる

  * recvのflagsと同じオプションを指定でき、指定しない場合は0を指定したものとみなす

```ruby
connect(host, port)
send(mesg, flags, dest_sockaddr = nil)
send(mesg, flags, host, port)
```

例) UDPSocketを使って10000ポートに対してUDPでデータを送信する

```ruby
require 'socket'

socket = UDPSocket.new

socket.send "Hello UDP.\n", 0, "localhost", 10000
socket.close

# サーバー側のコンソール
Hello UDP.
```



## 4.プロセス間通信を制御するクラス

* `UNIXServer`・`UNIXSocket`は、UNIX系OSのみで使用可能

* 通常、プロセスは独立したアドレス空間で動作しているため互いに影響しないようになっている

  * もしプロセス間でデータを共有したい場合や情報のやりとりをしたい場合に、プロセス間通信を使用する

* これらのクラスは、TCPServer・TCPSocketとほとんど同じように扱える



### サーバー側インスタンスの作成(UNIXServer)

* プロセス間のサーバー側インスタンスを作成する

* TCPServerクラスではホストとポート番号を指定したが、UNIXServerクラスでは任意のパス名を用いる

```ruby
UNIXServer.new(path)
UNIXServer.open(path)
accept
```

例) UNIXServerを使って、プロセス間の待ち受けを行う

```ruby
require 'socket'

socket_name = "test_socket"

File.unlink socket_name if File.exist?(socket_name) && File.socket?(socket_name)

server = UNIXServer.new socket_name
loop {
  client = server.accept
  client.puts "Hello UNIXServer."
  client.close
}
```



### クライアント側インスタンスの作成(UNIXSocket)

* プロセス間通信のクライアント側インスタンスを作成する

* 引数には、通信するサーバーが作成したソケット名を指定する

```ruby
require 'socket'

socket = UNIXSocket.new "test_socket"
print socket.gets

# 実行結果
Hello UNIXServer.
```



| 版     | 年/月/日   |
| ------ | ---------- |
| 初版   | 2018/10/21 |
| 第二版 | 2019/05/13 |
