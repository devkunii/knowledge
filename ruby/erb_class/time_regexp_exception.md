## 5-11.`Time`クラス

* 時刻を表すクラス

* 時刻は、世界標準時の1970年1月1日午前0時(起算時)からの経過秒数で保持される

* タイムゾーンとして、 **UTC** か **地方時刻(ローカルタイム)** を指定することができる

### 5-11-1.`Time`オブジェクトの生成

* `Time.new`、`Time.now`：現在時刻の`Time`オブジェクトを生成して返す。タイムゾーンはローカルタイム(日本なら、`+0900`)

* `Time.at`：引数で指定した起算時からの秒数に対応する`Time`オブジェクトを生成して返す。さらに精度が必要な場合には、2番目の引数にマイクロ秒を指定

* `Time.mktime`、`Time.local`：与えられた引数に対応する`Time`オブジェクトを生成して返す

  * 引数が7個までの場合、先頭から`年`、`月`、`日`、`時`、`分`、`秒`、`マイクロ秒`を指定できる

  * 年のみ省略できない。他の引数を省略した場合は、`1`、`1`、`0`、`0`、`0`、`0`が指定されたとみなす

  * 月に関しては、`"Jan"`や`"Feb"`などの英語の月名の省略形も指定できる

  * 引数が10個の場合は、先頭から`秒`、`分`、`時`、`日`、`月`、`年`、`曜日に対応する数値`、`年日`、`夏時間かどうかの真理値`、`タイムゾーン`を指定できる

  * ただし、`曜日に対応する数値`、`年日`、`夏時間かどうかの真理値`、`タイムゾーン`に関しては無視される

* `Time.gm`、`Time.utc`：引数の数やその順序は`Time.mktime`と同じだが、生成される`Time`オブジェクトのタイムゾーンが **UTC** になる

```ruby
# Time.nowメソッド
>> Time.now
=> 2018-09-15 12:08:25 +0900

# Time.atメソッド
>> Time.at(1234567890)
=> 2009-02-14 08:31:30 +0900
>> Time.at(1234567890, 1234567890)  # マイクロ秒指定
=> 2009-02-14 08:52:04 +0900

# Time.mktimeメソッド(引数7個)
>> Time.mktime(2017)          # 年のみ指定
=> 2017-01-01 00:00:00 +0900
>> Time.mktime(2017, 7, 7)    # 年、月、日のみ指定
=> 2017-07-07 00:00:00 +0900

# Time.mktimeメソッド(引数10個)
>> Time.mktime(0, 0, 0, 7, 7, 2017, 4, 188, false, "JST")
=> 2017-07-07 00:00:00 +0900

# Time.gmメソッド
>> Time.gm(2017)
=> 2017-01-01 00:00:00 UTC
>> Time.gm(2017, 7, 7)
=> 2017-07-07 00:00:00 UTC
```

***

### 5-11-2.`Time`オブジェクトの属性を取得する

* `year`：年を取得

* `mon`、`month`：月を取得

* `day`、`mday`：日を取得

* `hour`：時を取得

* `min`：分を取得

* `sec`：秒を取得

* `usec`、`tv_usec`：マイクロ秒を取得

```ruby
>> t = Time.mktime(2017, 1, 2, 3, 4, 5, 6)
=> 2017-01-02 03:04:05 +0900
>> t.year     # 年を取得
=> 2017
>> t.mday     # 日を取得
=> 2
>> t.sec      # 秒を取得
=> 5
```

* `wday`：曜日に対応する数値を返す。日曜(0)〜土曜(6)が対応

* `yday`：1月1日からの日数を返す

```ruby
>> t = Time.mktime(2017, 1, 2, 3, 4, 5, 6)
=> 2017-01-02 03:04:05 +0900
>> t.wday      # 月曜日
=> 1
>> t.yday      # 1月1日から2日目
=> 2
```

* `isdst`、`dst?`：夏時間かどうかを返す

* `gmt?`、`utc?`：タイムゾーンが **UTC** かどうかを返す

```ruby
>> t = Time.mktime(2017, 1, 2, 3, 4, 5, 6)
=> 2017-01-02 03:04:05 +0900
>> t.dst?
=> false
>> t.gmt?
=> false
```

* `gmtoff`、`gmt_offset`：UTC時刻との差を秒単位の数値として返す

```ruby
>> t = Time.mktime(2017, 1, 2, 3, 4, 5, 6)
=> 2017-01-02 03:04:05 +0900
>> t.gmtoff
=> 32400
>> t.gmtoff / 3600  # 時に変更
=> 9
```

***

### 5-11-3.タイムゾーンを変更する

* `localtime`：タイムゾーンをローカルタイムに変更

* `gmtime`、`utc`：タイムゾーンをUTCに変更

```ruby
>> t = Time.mktime(2017, 1, 2, 3, 4, 5, 6)
=> 2017-01-02 03:04:05 +0900
>> t.localtime
=> 2017-01-02 03:04:05 +0900
>> t.gmtime
=> 2017-01-01 18:04:05 UTC
```

* `getlocal`：タイムゾーンをローカルタイムに変更した新しい`Time`オブジェクトを返す

* `getgm`、`getutc`：タイムゾーンをUTCに変更した新しい`Time`オブジェクトを返す

```ruby
>> t = Time.mktime(2017, 1, 2, 3, 4, 5, 6)
=> 2017-01-02 03:04:05 +0900
>> t.object_id
=> 70329225673020
>> t1 = t.getlocal
=> 2017-01-02 03:04:05 +0900
>> t1.object_id
=> 70329225651440
```

***

### 5-11-4.`Time`オブジェクトの演算

* `+`：指定した秒数後の`Time`オブジェクトを返す

* `-`：指定した秒数前の`Time`オブジェクトを返す

```ruby
>> t = Time.mktime(2017, 1, 2, 3, 4, 5, 6)
=> 2017-01-02 03:04:05 +0900
>> t + 60 * 60 * 3
=> 2017-01-02 06:04:05 +0900
>> t - 60 * 60 * 3
=> 2017-01-02 00:04:05 +0900

# Timeオブジェクト同士の差
>> t1 = Time.mktime(2011, 1, 2, 3, 4, 5, 6)
=> 2011-01-02 03:04:05 +0900
>> t2 = Time.mktime(2011, 2, 3, 4, 5, 6, 7)
=> 2011-02-03 04:05:06 +0900
>> t2 - t1
=> 2768461.000001
```

***

### 5-11-5.`Time`オブジェクトの変換

* `to_i`、`tv_sec`：起算時からの秒数を整数で返す

* `to_f`：起算時からの秒数を浮動小数点数で返す

```ruby
>> t = Time.mktime(2017, 1, 2, 3, 4, 5, 6)
=> 2017-01-02 03:04:05 +0900
>> t.to_i
=> 1483293845
>> t.to_f
=> 1483293845.000006
```

* `to_a`：先頭から、秒、分、時、日、月、年、曜日に対応する数値、年日、夏時間かどうかの真理値、タイムゾーン順の配列を返す

```ruby
>> t = Time.mktime(2017, 1, 2, 3, 4, 5, 6)
=> 2017-01-02 03:04:05 +0900
>> t.to_a
=> [5, 4, 3, 2, 1, 2017, 1, 2, false, "JST"]
```

* `to_s`：UNIXコマンドのdateコマンドのような形式の文字列を返す

```ruby
>> t = Time.mktime(2017, 1, 2, 3, 4, 5, 6)
=> 2017-01-02 03:04:05 +0900
>> t.to_s
=> "2017-01-02 03:04:05 +0900"
```

* `strftime`：指定したフォーマットに従って文字列に変換した結果を返す

```ruby
>> t = Time.mktime(2017, 1, 2, 3, 4, 5, 6)
=> 2017-01-02 03:04:05 +0900
>> t.strftime("%Y年%m月%d日 %H時%M分%S秒")
=> "2017年01月02日 03時04分05秒"
```

この表は、後に画像として貼る
2018/9/15

***

## 5-12.`Regexp`クラス

* 正規表現オブジェクトを扱うクラス

* 正規表現を使って文字列やデータのマッチングを行うときに、使用

### 5-12-1.正規表現オブジェクトを生成

* 正規表現は、正規表現リテラルを使って表現する。リテラルの末尾には、オプションが指定できる。

* オプションは、

  * `i`：大文字小文字の違いを無視

  * `m`：正規表現の`.`で改行にマッチさせる

  * `x`：空白や`#`から始まるコメントを無視する

  などがある。また、オプションの複数指定もできる

```ruby
>> a = /abcdefg/i  # 大文字小文字の違いを無視する
=> /abcdefg/i
>> a.class
=> Regexp
```

* `Regexp.new`、`Regexp.compile`：正規表現オブジェクトを生成する。2つ目の引数に、マッチングのオプションを指定できる。

  * `Regexp::IGNORECASE`：大文字小文字の違いを無視する

  * `Regexp::MULTILINE`：正規表現の`.`が改行にマッチするようになる

  * `Regexp::EXTENDED`：バックスラッシュでエスケープされていない空白と、`#`から改行までを無視する

  * 論理和を使って複数指定も可能

  * マッチングするときの文字コードを、3番目の引数で指定することも可能

```ruby
>> a = Regexp.new("abcdefg", Regexp::MULTILINE | Regexp::IGNORECASE)
=> /abcdefg/mi
```

***

### 5-12-2.正規表現オブジェクトでマッチングする

* `match`：正規表現オブジェクトで文字列とマッチングさせる。マッチした場合には`MatchData`オブジェクトを、しなかった場合には`nil`を返す

* `=~`：正規表現オブジェクトで文字列とマッチングさせる。マッチすればマッチした位置のインデックスが、しなかった場合は`nil`を返す

* `===`：正規表現オブジェクトで文字列とマッチングさせる。マッチすれば`true`、しなかった場合は`false`が返る

* `~`：特殊変数`$_`とマッチングする

```ruby
# matchメソッド
>> a = Regexp.new("abc")
=> /abc/
>> a.match("abcdef")
=> #<MatchData "abc">

# =~メソッド
>> a = Regexp.new("abc")
=> /abc/
>> a =~ "abcdef"          # 0番目の文字がマッチした
=> 0
>> "abcdefg" =~ a         # 0番目の文字がマッチした
=> 0

# ===メソッド
>> a = Regexp.new("abc")
=> /abc/
>> a === "abcdef"
=> true

# ~メソッド
>> $_ = "abcdefg"
=> "abcdefg"
>> a = Regexp.new("abc")
=> /abc/
>> ~ a
=> 0
```

***

### 5-12-3.正規表現の特殊文字をエスケープする

* `Regexp.escape`、`Regexp.quote`：ピリオド`.`、カッコ`[]`などでマッチングする際に、これらの文字を自動的にエスケープする

```ruby
>> Regexp.escape("array.push(hash[key])")
=> "array\\.push\\(hash\\[key\\]\\)"
```

***

### 5-12-4.マッチした結果を取得する

* `Regexp.last_match`：正規表現でマッチした結果を取得。`MatchData`オブジェクト(現在のスコープ(トップレベルやクラス・モジュール・メソッド定義)の中で最後に行った正規表現のマッチ結果)を返す。特殊変数`$_`でも取得できる

```ruby
>> /abcdefg/ =~ "abcdefghijklmnopqrstuvwxyz"
=> 0
>> Regexp.last_match
=> #<MatchData "abcdefg">
>> $~
=> #<MatchData "abcdefg">
```

* `Regexp.last_match`メソッドに整数値を与えると、該当のマッチ文字列が得られる。

  * `0`であれば正規表現にマッチした文字列

  * それ以降の整数では、カッコにマッチした部分文字列が得られる。これらの文字列はそれぞれ特殊変数`$&`、`$1`、`$2`などでも取得可能

```ruby
>> /(abc)d(efg)/ =~ "abcdefghijklmnopqrstuvwxyz"
=> 0
>> Regexp.last_match(0)
=> "abcdefg"
>> $&
=> "abcdefg"
>> Regexp.last_match(1)
=> "abc"
>> $1
=> "abc"
>> $2
=> "efg"
```

***

### 5-12-5.正規表現の論理和を求める

* `Regexp.union`：複数の正規表現を結合し、そのどれかにマッチするような新しい正規表現を求める

```ruby
>> a = Regexp.new("abc")
=> /abc/
>> b = Regexp.new("ABC")
=> /ABC/
>> c = Regexp.union(a, b)
=> /(?-mix:abc)|(?-mix:ABC)/
>> c =~ "abc"
=> 0
>> Regexp.last_match
=> #<MatchData "abc">
```

***

### 5-12-6.正規表現オブジェクトのオプションや属性を取得する

* `options`：正規表現オブジェクトを生成する時に設定したオプションである`Regexp::IGNORECASE`、`Regexp::MULTILINE`、`Regexp::EXTENDED`の論理和を返す

* `casefold?`：オプション`Regexp::IGNORECASE`が設定してあるかどうかを返す

* `encoding`：正規表現オブジェクトがコンパイルされている文字コードを`Encoding`オブジェクトとして返す

* `source`：正規表現の元となった文字列表現を返す

  * `to_s`：他の正規表現に埋め込んでも元の意味が保たれるような形式

  * `inspect`：`to_s`メソッドよりも自然な形式な文字列になるが、元の意味は保たれない

```ruby
# optionsメソッド
>> a = Regexp.new("abcdefg", Regexp::MULTILINE | Regexp::IGNORECASE)
=> /abcdefg/mi
>> a.options
=> 5
>> a.options & Regexp::IGNORECASE
=> 1
>> a.options & Regexp::EXTENDED
=> 0

# casefold?メソッド
>> a = Regexp.new("abcdefg")
=> /abcdefg/
>> a.casefold?
=> false
>> a = Regexp.new("abcdefg", Regexp::MULTILINE | Regexp::IGNORECASE)
=> /abcdefg/mi
>> a.casefold?
=> true

# encodingメソッド
>> a = Regexp.new("ルビー")
=> /ルビー/
>> a.encoding
=> #<Encoding:UTF-8>
>> a = Regexp.new("ルビー".encode("EUC-JP"))
=> /\x{A5EB}\x{A5D3}\x{A1BC}/
>> a.encoding
=> #<Encoding:EUC-JP>

# sourceメソッド
>> a = Regexp.new("abcdefg", Regexp::MULTILINE | Regexp::IGNORECASE)
=> /abcdefg/mi
>> a.source
=> "abcdefg"
>> a.to_s
=> "(?mi-x:abcdefg)"
>> a.inspect
=> "/abcdefg/mi"
```

***

## 5-13.`Exception`クラス

* 全ての例外クラスのスーパークラス。エラーが発生した場合や、`raise`メソッドで例外を発生した時に、このクラスのオブジェクトが生成される

#### 例外クラスの自作

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

***

#### エラーメッセージを指定する

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

#### エラーメッセージを取得する

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

***

#### バックトレースを取得

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

***

#### バックトレースを取得(オリジナルの情報追加)

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

***
