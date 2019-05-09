00 基本的なRubyの扱い
===================

## irbの起動

```ruby
$ irb
>>
```



## 変数と定数

### 変数

何度でも値を変更できる

```ruby
>> a = 1
>> print a + 1
2=> nil
>> a = 3
=> 3
```

> 変数の型(stringなど)を指定していなくても、値が型の情報を管理する
>
> →動的型言語

色々なコード

```ruby
>> a = 1
=> 1
>> a = "ruby"
=> "ruby"
>> print a
ruby=> nil
```

空白を用意した場合

```ruby
>> a = 1
=> 1
>> print ( a + 1 )
2=> nil
>>
```



### 定数

変更の際に警告が出てくる。基本的に値を変更しない。

```ruby
>> A = 1
=> 1
>> print A + 1
2=> nil
>> A = 2
(irb):5: warning: already initialized constant A
(irb):3: warning: previous definition of A was here
=> 2
```



## 関数とクラス

### 関数

```ruby
>> def add(a, b)
>>   a + b
>> end
=> :add
>> print add(1, 2)
3=> nil
```

省略した形

```ruby
>> def add(a, b); a + b; end
=> :add
>> print add(1, 2)
3=> nil
```



### クラス

```ruby
>> class Hoge                #クラスHogeの定義
>>   def test                #testメソッドの定義
>>     1
>>   end
>> end
=> :test
>> hoge = Hoge.new           #Hogeクラスのインスタンスを生成
=> #<Hoge:0x007ff3890a5ad0>
>> hoge.test                 #Hogeクラスのtestメソッドを実行
=> 1
```



| 版     | 年/月/日   |
| ------ | ---------- |
| 初版   | 2018/08/11 |
| 第二版 | 2019/05/09 |
