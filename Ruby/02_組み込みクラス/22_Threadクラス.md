22 Threadクラス
==============

## 目次

* [Threadクラスとは](#0Threadクラスとは)

* [スレッドの生成](#1スレッドの生成)

* [スレッドの状態](#2スレッドの状態)

* [スレッドの実行・一時停止・終了](#3スレッドの実行一時停止終了)

* [スレッド終了時のensure節](#4スレッド終了時のensure節)

* [スレッド中の例外](#5スレッド中の例外)

* [スレッドのデッドロック](#6スレッドのデッドロック)

* [スレッドのリスト](#7スレッドのリスト)

* [スレッドの切り替え](#8スレッドの切り替え)

* [スレッドの終了を待つ](#9スレッドの終了を待つ)

* [スレッドの優先度](#10スレッドの優先度)

* [スレッド固有のデータ](#11スレッド固有のデータ)



## 0.Threadクラスとは

* スレッドを表すクラス

* 使用することで、並行プログラミングが可能になる

* プログラムの開始と同時に生成されるスレッドを、メインスレッドと呼ぶ

* メインスレッドが終了するときには、他のスレッドも含めてプログラム自体が終了する

* スレッドの実行はスケジューリングされており、管理方法はプラットフォームに依存する

* 現在実行中のスレッドを、カレントスレッドという



## 1.スレッドの生成

* `Thread.new`：ブロックを渡すことで、オブジェクトであるスレッドが生成される

  * 引数をブロックに渡すことができる

  * 他にも、`Thread.fork`、`Thread.start`なども使用可能

```ruby
>> t = Thread.new { sleep 1 }
=> #<Thread:0x007ff3c3833da8@(irb):6 run>
>> t = Thread.new(3) {|t| sleep t }
=> #<Thread:0x007ff3c4169cd8@(irb):7 run>
```



## 2.スレッドの状態

* 生成したスレッドは、実行状態を持っている

* `status`：実行状態を調べる

|   状態  |                                    意味                                    |
|:-------|:---------------------------------------------------------------------------|
|run     |実行中または実行可能状態。生成直後やrun・wakeupメソッドで起こったスレッドはこの状態になる|
|sleep   |一時停止状態。stop・joinメソッドにより一時停止されたスレッドの状態                    |
|aborting|終了処理中。killメソッドなどで終了されるスレッドは、一時的にこの状態になる              |
|false   |killメソッドで終了したり、正常終了したスレッドの場合に返る                           |
|nil     |例外などで以上終了したスレッドの場合にはnilが返る                                   |

* `alive?`：スレッドが生きているかどうかの確認

* `stop?`：スレッドが終了もしくは一時停止していることを確認

```ruby
>> t = Thread.new { sleep 100 }
=> #<Thread:0x007fb5ec035e68@(irb):1 run>
>> t.status
=> "sleep"
>> t.alive?
=> true
```



## 3.スレッドの実行・一時停止・終了

* 生成したスレッドは実行中となる

* `stop`・`join`：他のスレッドを待っている場合に、一時停止状態となる

```ruby
>> t = Thread.new do
?>       Thread.stop
>>       puts "OK\n"
>>     end
=> #<Thread:0x007fb5ec14a5d8@(irb):4 sleep_forever>
```

* `run`：スレッドを再開する。

* `wakeup`：スレッドを開始するが、実行可能状態にするだけで、実行されるかどうかは他のスレッドの状況に依存

```ruby
>> t = Thread.new do
?>       Thread.stop
>>       puts "OK\n"
>>     end
=> #<Thread:0x007fb5ec1397d8@(irb):8 run>
>> t.run
OK
=> #<Thread:0x007fb5ec1397d8@(irb):8 dead>
```

* `kill`：スレッドを終了する

  * 指定したスレッドの`exit`メソッドを呼び出す

* `exit`：カレントスレッドの`exit`メソッドを呼び出す

```ruby
>> t = Thread.new do
?>       Thread.exit
>>     end
=> #<Thread:0x007fb5ec8b0cf0@(irb):13 dead>
>> t
=> #<Thread:0x007fb5ec8b0cf0@(irb):13 dead>
```



## 4.スレッド終了時のensure節

* スレッド生成時にensure節がある場合、スレッドが終了するときに実行される

* 正常に終了する場合だけでなく、他のスレッドから`kill`などで終了させられた場合も同様

```ruby
>> t = Thread.new do
?>       begin
?>         sleep 10000
>>       ensure
?>         puts 'Killed'
>>       end
>>     end
=> #<Thread:0x007fb5ec8aa800@(irb):17 sleep>
>> t.kill
Killed=> #<Thread:0x007fb5ec8aa800@(irb):17 sleep>
>>
t
=> #<Thread:0x007fb5ec8aa800@(irb):17 dead>
```



## 5.スレッド中の例外

* スレッドの中で例外が発生した時、通常はそのスレッドのみが警告なしに終了される(rescue節は別)

* ただし、そのスレッドを`join`メソッドで待っている他のスレッドがある場合、待っているスレッドに対して同じ例外が再度発生する

```ruby
>> t = Thread.new { Thread.pass; raise "Raise exception" }
=> #<Thread:0x007fb5ec8a2560@(irb):26 run>
>> e = Thread.new do
?>       begin
?>         t.join
>>       rescue => ex
>>         puts ex.message
>>       end
>>     end
Raise exception=> #<Thread:0x007fb5eb06bd98@(irb):27 sleep>
```



### スレッド中の例外が発生した場合の、プログラム自体を終了させる方法

1. `Thread.abort_on_exception`メソッドに`true`を設定

2. 特定のスレッドの`abort_on_exception`メソッドに`true`を設定する

3. グローバル変数`$DEBUG`を`true`にし、プログラムを`-d`オプション付きで実行する

> `-d`：デバッグモードでスクリプトを実行します。`$DEBUG` を `true` にします。


* `raise`：スレッドで強制的に例外を発生させる

  * 実行したスレッドの中で、指定した例外を発生させる

```ruby
>> Thread.new { sleep 1; Thread.main.raise "Error" }; begin; sleep; rescue => ex; puts ex.message; end
Error
=> nil
```



## 6.スレッドのデッドロック

### デッドロック

Rubyでは、以下の状態になっている場合、デッドロックとみなして`fatal`が発生してプログラムが終了する

* スレッドが複数ある

* 全てのスレッドが停止状態である

* IO待ちのスレッドは存在しない

### 実際のコード

```ruby
>> t = Thread.new { Thread.stop }
=> #<Thread:0x007fda2a8eb910@(irb):3 sleep_forever>
>> Thread.stop
fatal: No live threads left. Deadlock?
```

* ただし、メインスレッドのみが`Thread.stop`で停止している状態では、`fatal`は発生しない



## 7.スレッドのリスト

* `list`：実行中のスレッドのリストを取得する

  * スレッドを生成していない場合は、メインスレッドのみ表示される

```ruby
>> Thread.list
=> [#<Thread:0x007fda2b07ef88 run>, #<Thread:0x007fda2a8eb910@(irb):3 sleep_forever>]
>> t = Thread.new { Thread.stop }
=> #<Thread:0x007fda2a8d0728@(irb):7 run>
>> Thread.list
=> [#<Thread:0x007fda2b07ef88 run>, #<Thread:0x007fda2a8eb910@(irb):3 sleep_forever>, #<Thread:0x007fda2a8d0728@(irb):7 sleep_forever>]
```

* `main`：メインスレッドが得られる

* `current`：カレントスレッドが得られる

```ruby
>> Thread.main
=> #<Thread:0x007fda2b07ef88 run>
>> t = Thread.new { sleep 100 }
=> #<Thread:0x007fda2a8ba748@(irb):10 run>
>> Thread.current
=> #<Thread:0x007fda2b07ef88 run>
```



## 8.スレッドの切り替え

* `pass`：実行中のスレッドの状態を変えずに、他のスレッドに実行権を譲る

```ruby
>> t = Thread.new { Thread.pass; raise "Raise exception" }
=> #<Thread:0x007fda2a8a2dc8@(irb):12 run>
>> t.status
=> nil
```



## 9.スレッドの終了を待つ

* 通常スレッドは生成と同時に実行されるが、ラウンドロビンによりいつ終了するか分からない

  > ラウンドロビン：１つのドメインに複数のIPアドレスを設定して、コンピュータの負荷を軽減した状態(1つの資源を順番に利用する)

* `join`：そのスレッドの実行終了までカレントスレッドを停止することができる

```ruby
>> t = Thread.new { Thread.pass; 10.times {|i| puts i ** 2; sleep 1 } }; t.join
0
1
4
9
16
25
36
49
64
81
=> #<Thread:0x007fda2b147cf8@(irb):19 dead>
```

* `value`：スレッドの終了を待つことができる。スレッドのブロックの評価結果を返す



## 10.スレッドの優先度

* `priority`：スレッドの優先度を得る

  * 値が大きいほど、優先度が高くなる

* `priority=`：優先度を設定することができる

  * メインスレッドのデフォルト値は0で、新しく生成されたスレッドは親スレッドの値を引き継ぐ

```ruby
>> t = Thread.new(10000) {|t| sleep t }
=> #<Thread:0x007fda2a878730@(irb):20 run>
>> t.priority
=> 0
>> t.priority = 1000
=> 1000
```



## 11.スレッド固有のデータ

* スレッドでは、データに固有の名前を付けて保持することができる

* これを用いて、スレッド間で値をやりとりできる

```ruby
>> t = Thread.current
=> #<Thread:0x007fda2b07ef88 run>
>> t[:foo] = "Bar"
=> "Bar"
>> t[:foo]
=> "Bar"
```

* `key?`：データが保持されているかどうかを調べる

```ruby
>> t = Thread.current
=> #<Thread:0x007fda2b07ef88 run>
>> t[:foo] = "Bar"
=> "Bar"
>> t.key?(:foo)
=> true
```

* `keys`：保持している名前を取得する

```ruby
>> t = Thread.current
=> #<Thread:0x007fda2b07ef88 run>
>> t[:foo] = "Bar"
=> "Bar"
>> t.keys
=> [:foo]
```

| 版     | 年/月/日   |
| ------ | ---------- |
| 初版   | 2018/10/14 |
| 第二版 | 2019/05/11 |
