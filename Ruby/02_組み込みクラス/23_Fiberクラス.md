23 Fiberクラス
=============

## 目次

* [Fiberとは](#0Fiberとは)

* [基本的な振る舞い](#1基本的な振る舞い)



## 0.Fiberとは

* 「ある処理を途中まで実行して、その後任意のタイミングで、前回の続きから処理を行う」という処理の流れを扱う

* プリエンティブ：各スレッドの実行はスケジューラにより自動的に与えられるタイムスライスで実行される(プログラマが制御できない)

* ノンプリエンティブ：ファイバ。軽量なスレッドを提供



## 1.基本的な振る舞い

* `new`：ファイバが生成される。(ブロックを渡す)

* `resume`：生成された時点では実行されないので、実行を開始させる

* ファイバには親子関係があり、`resume`を呼び出した側が親、呼び出されたファイバが子になる

* `yield`：親である呼び出し元へコンテキストが変わる

```ruby
>> f = Fiber.new do
?>       loop do
?>         puts 'hello'
>>         puts 'child -> parent'
>>         Fiber.yield
>>       end
>>     end
=> #<Fiber:0x007fda2b8699a0>
>> 3.times do
?>  puts 'parent -> child'
>>  f.resume
>> end
parent -> child
hello
child -> parent
parent -> child
hello
child -> parent
parent -> child
hello
child -> parent
=> 3
```

| 版     | 年/月/日   |
| ------ | ---------- |
| 初版   | 2018/10/14 |
| 第二版 | 2019/05/11 |
