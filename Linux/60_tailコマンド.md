06 tailコマンド
==============

* `tail`コマンド：ファイルの末尾を表示するコマンド

  * オプションを指定しない場合は、ファイルの末尾10行が表示される

  ```bash
  $ tail /etc/passwd
  _datadetectors:*:257:257:DataDetectors:/var/db/datadetectors:/usr/bin/false
  _captiveagent:*:258:258:captiveagent:/var/empty:/usr/bin/false
  _ctkd:*:259:259:ctkd Account:/var/empty:/usr/bin/false
  _applepay:*:260:260:applepay Account:/var/db/applepay:/usr/bin/false
  _hidd:*:261:261:HID Service User:/var/db/hidd:/usr/bin/false
  _cmiodalassistants:*:262:262:CoreMedia IO Assistants User:/var/db/cmiodalassistants:/usr/bin/false
  _analyticsd:*:263:263:Analytics Daemon:/var/db/analyticsd:/usr/bin/false
  _fpsd:*:265:265:FPS Daemon:/var/db/fpsd:/usr/bin/false
  _timed:*:266:266:Time Sync Daemon:/var/db/timed:/usr/bin/false
  _reportmemoryexception:*:269:269:ReportMemoryException:/var/db/reportmemoryexception:/usr/bin/false
  ```

  * `-n`オプション：`tail`コマンドで表示する行数を指定する

    * 例)`-n 1`と指定して、ファイルの末尾1行のみを表示

    ```bash
    $ tail -n 1 /etc/passwd
    _reportmemoryexception:*:269:269:ReportMemoryException:/var/db/reportmemoryexception:/usr/bin/false
    ```

* `head`コマンド：ファイルの先頭部分を表示する

  * デフォルトでは、先頭10行を表示する

  * `-n`オプションで表示する行数を指定できる

  ```bash
  $ head -n 1 /etc/passwd
  ##
  ```



### ファイルへの追記を監視する

* `-f`オプション：追記されるたびにその内容を表示してファイル監視をすることができる

  * ログの出力やデータ収集などで常時追記されていくファイルを表示することができる

  ```bash
  tail -f <ファイル名>
  ```

  * これにより、ファイルの内容が書き換えられると、その都度 **リアルタイムに表示** される

  * 例)`output.log`というファイルへの追記を監視する

  ```bash
  $ touch output.log
  $ tail -f output.log
                        # ここでカーソルが止まる
  ```

  * コマンドは終了せずにカーソルが止まったように見える

    * この状態で別のターミナルから、`output.log`へ「Hello」という文字列を追記してみる

    ```bash
    echo Hello >> output.log
    ```

  * 先ほど`tail -f`を実行していたターミナル上に、追記内容が表示される

  ```bash
  $ tail -f output.log
  Hello

  ```
* これは、Linuxの運用などでログファイルの監視によく使われる

* このコマンドを終了するには、`Ctrl`+`c`を押す



| 版 |  年/月/日 |
|----|----------|
|初版|2019/03/11|
