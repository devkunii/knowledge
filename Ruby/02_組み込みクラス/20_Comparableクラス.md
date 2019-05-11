20 Comparableモジュール
======================

## 目次

* [Comparableモジュールとは](#0Comparableモジュールとは)

* [between?メソッド](#1between?メソッド)



## 0.Comparableモジュールとは

* インクルードしたクラスで比較演算子である`<=>`メソッドを元にオブジェクト同士での比較ができるようになる

* インクルードしたクラスで利用できるインスタンスメソッドは、

  * `<`：負の整数で`true`

  * `<=`：負の整数か0で`true`

  * `==`：0で`true`

  * `>`：正の整数で`true`

  * `>=`：正の整数か0で`true`

  * `between?`：引数`min`と`max`の間にあれば`true`



## 1.between?メソッド

* `between?`：レシーバ`obj`の値が引数`min`と`max`の間に含まれればtrue、そうでなければfalseを返す

  * `obj`が`min`または`max`と等しいときはtrueを返す

  * 実際には、演算子`<`と`>`を使って`obj < min`または`obj > min`ならfalse、それ以外ならtrueを返す

  * このメソッドはComparableをインクルードしている数値クラス、String、Timeなどで扱える

```ruby
>> str = "hello"
=> "hello"
>> p str.between?("helicopter", "help")
true
=> true
>> p str.between?("help", "here")
false
=> false
```

```ruby
>> num = 123
=> 123
>> p num.between?(100, 150)
true
=> true
>> p num.between?(123, 150)
true
=> true
```



### Sampleクラス(例)

* 通常の大小関係と逆の挙動をするクラス

```ruby
>> class Sample
>>   def initialize(value)
>>     @value = value
>>   end
>>
>>   def value
>>     @value
>>   end
>>
>>   def <=>(other)
>>     other.value <=> self.value
>>   end
>> end
=> :<=>
```



| 版     | 年/月/日   |
| ------ | ---------- |
| 第二版 | 2019/05/11 |
