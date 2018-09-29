## 5-5.`Encoding`クラス

### 主なエンコーディング

* `Encoding::UTF_8`：`UTF-8`を表すエンコーディング

* `Encoding::EUC_JP`：`EUC_JP`を表すエンコーディング

* `Encoding::ISO_2022_JP`：`JIS`を表すエンコーディングで、Rubyではダミーエンコーディング

* `Encoding::Shift_JIS`：`Shift_JIS`を表すエンコーディング

* `Encoding::Windows_31J`：Windowsで用いられる`Shift_JIS`の亜種(`CP932`とも言う)
  →`Encoding::CP932`でも参照可能

* `Encoding::ASCII`：`US-ASCII`を表すエンコーディング
  →`Encoding::US_ASCII`でも参照可能

* `Encoding::ASCII_8BIT`：`ASCII`互換のエンコーディングで、文字コードを持たないデータや、文字列を単なるバイト列とσ知恵扱いたい場合に利用

***

#### 規定の外部エンコーディング

* エンコーディングが指定されていないときは、規定の外部エンコーディングは各システムに依存

* Linuxであれば、localeに`UTF-8`が指定されている

* `default_external`メソッド：規定の外部エンコーディングを取得

```ruby
>> Encoding.default_external
=> #<Encoding:UTF-8>
```

***

#### エンコーディングの互換性

* `compatible?`メソッド：異なるエンコーディングの間の互換性を調べる

* 互換性がある場合には、エンコーディングを、ない場合には`nil`を返す

```ruby
>> Encoding.compatible?(Encoding::UTF_8, Encoding::US_ASCII)
=> #<Encoding:UTF-8>
>> Encoding.compatible?(Encoding::UTF_8, Encoding::Shift_JIS)
=> nil
```

* 互換性のあるエンコーディングでは文字列を結合できるが、互換性のない場合はエラーになり結合できない

```ruby
>> a = "ルビー"
=> "ルビー"
>> b = a.encode("EUC-JP")
=> "\x{A5EB}\x{A5D3}\x{A1BC}"
>> a + b
Encoding::CompatibilityError: incompatible character encodings: UTF-8 and EUC-JP
```

* ただし、互換性のないエンコーディングでもどちらか一方の文字列がASCII文字しか含まない場合は結合可能

```ruby
>> a = "abc"
=> "abc"
>> b = "あいう".encode("EUC-JP")
=> "\x{A4A2}\x{A4A4}\x{A4A6}"
>> b.encoding
=> #<Encoding:EUC-JP>
>> (a + b)
=> "abc\x{A4A2}\x{A4A4}\x{A4A6}"
>> (a + b).encoding
=> #<Encoding:EUC-JP>
```

***
