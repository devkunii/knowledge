15 Exceptionクラス
=================

## 目次

* [Exceptionクラスとは](#0Exceptionクラスとは)

* [例外クラスの自作](#1例外クラスの自作)

* [エラーメッセージを指定する](#2エラーメッセージを指定する)

* [エラーメッセージを取得する](#3エラーメッセージを取得する)

* [バックトレースを取得](#4バックトレースを取得)

* [オリジナルの情報追加したバックトレースを取得](#5オリジナルの情報追加したバックトレースを取得)



## 0.Exceptionクラスとは

* 全ての例外クラスのスーパークラス

* エラーが発生した場合や、`raise`メソッドで例外を発生した時に、このクラスのオブジェクトが生成される



## 1.例外クラスの自作

```ruby
>> class MyError < RuntimeError; end
=> nil
>> begin
?>   raise MyError
>> rescue => ex
>>   p ex
>> end
#<MyError: MyError>
=> #<MyError: MyError>
```



## 2.エラーメッセージを指定する

```ruby
>> class MyError < RuntimeError; end
=> nil
>> begin
>>   raise MyError.exception('エラーが発生しました。')
>> rescue => ex
>>   p ex
>> end
#<MyError: エラーが発生しました。>
=> #<MyError: エラーが発生しました。>
```

***

## 3.エラーメッセージを取得する

* `message`、`to_s`、`to_str`：例外オブジェクトに設定されているエラーメッセージを取得

* オブジェクトにエラーメッセージを設定するには、`new`、`exception`メソッドの呼び出し時に引数で指定

```ruby
>> class MyError < RuntimeError; end
=> nil
>> begin
?>   raise MyError.exception('エラーが発生しました。')
>> rescue => ex
>>   p ex.message
>> end
"エラーが発生しました。"
=> "エラーが発生しました。"
```



## 4.バックトレースを取得

* `backtrace`：発生した例外のバックトレース情報を取得。これらは、配列で返る。

```ruby
>> class MyError < RuntimeError; end
=> nil
>> begin
?> raise MyError.exception('エラーが発生しました。')
>> rescue => ex
>> p ex.backtrace
>> end
["(irb):41:in `irb_binding'", "/Users/MacUser/.rbenv/versions/2.4.1/lib/ruby/2.4.0/irb/workspace.rb:87:in `eval'",] # 省略
=> ["(irb):41:in `irb_binding'", "/Users/MacUser/.rbenv/versions/2.4.1/lib/ruby/2.4.0/irb/workspace.rb:87:in `eval'",] # 省略
>>
```



## 5.オリジナルの情報追加したバックトレースを取得

* `set_backtrace`：バックトレース情報にオリジナルの情報を設定する。ただし、それまでのバックトレース情報は上書きされることに注意。

```ruby
>> class MyError < RuntimeError; end
=> nil
>> begin
?>   raise MyError.exception('エラーが発生しました。')
>> rescue => ex
>>   ex.set_backtrace("This is new backtrace.")
>>   p ex.backtrace
>> end
["This is new backtrace."]
=> ["This is new backtrace."]
```



| 版     | 年/月/日   |
| ------ | ---------- |
| 第二版 | 2019/05/11 | 
