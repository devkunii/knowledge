04 cutコマンド
=============

* `cut`コマンド：入力行の一部分を切り出して出力するコマンド

  ```bash
  cut -d <区切り文字> -f <フィールド番号> [<ファイル名>]
  ```

  * `<区切り文字>`：指定した文字で入力行を分割し、その中の`<フィールド番号>`で指定したフィールドだけを出力する

  * 例)`-d , -f 3`：区切り文字を「`,`」と見なして3番目のフィールドだけが表示される

    => CSVファイルの特定のカラムを取り出すときに使うことができる

    ```bash
    cut -d , -f 3 file.csv
    ```

* `-d`オプションを付けずに区切り文字を指定しない場合は、デフォルト値として`タブ`が区切り文字としてみなされる

* 例)`/etc/password`：ユーザの情報を「:」区切りで記述したファイル

  ```bash
  $ cat /etc/passwd
  nobody:*:-2:-2:Unprivileged User:/var/empty:/usr/bin/false
  root:*:0:0:System Administrator:/var/root:/bin/sh
  daemon:*:1:1:System Services:/var/root:/usr/bin/false
  # 以下省略
  ```

  * このファイルのそれぞれのフィールドは、次のような意味を持っている

  1. ユーザ名

  1. パスワード(本当のパスワードではなく、ダミー値)

  1. ユーザID

  1. グループID

  1. コメント(本名など)

  1. ホームディレクトリ

  1. ログインシェル

* この中で、ログインシェルだけを出力したい場合には、次のように「:」区切りで7番目のフィールドを指定する

  ```bash
  $ cut -d : -f 7 /etc/passwd
  /usr/bin/false
  /bin/sh
  /usr/bin/false
  /usr/sbin/uucico
  ```

* 出力するフィールド番号は、カンマを付けて複数指定できる

  * 例)フィールド1番のユーザ名、6番のホームディレクトリ、7番のログインシェルを表示する

  ```bash
  $ cut -d : -f 1,6,7 /etc/passwd
  nobody:/var/empty:/usr/bin/false
  root:/var/root:/bin/sh
  daemon:/var/root:/usr/bin/false
  _uucp:/var/spool/uucp:/usr/sbin/uucico
  ```

* このように`cut`コマンドは、元のデータから特定の部分だけを取り出したい場合に便利



| 版 |  年/月/日 |
|----|----------|
|初版|2019/03/11|
