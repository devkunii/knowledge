MacOS homebrew
==============

## 目次

* [旧パッケージを使用する場合](#旧パッケージを使用する場合)



## 旧パッケージを使用する場合

> 例：Python3.7ではなく、Python3.6を使用したい場合

1. Homebrewのリポジトリのあるディレクトリまで移動

  ```bash
  $ cd /usr/local/Homebrew/Library/Taps/homebrew/homebrew-core/Formula/
  ```

2. 対象のパッケージのコミットを取得する

  * `git log [パッケージ名].rb`を使用して、バージョンを見ることができる

  * GitHubのコミット履歴から、ファイルを見て決定することもある

    > 例：`option`については、GitHubのファイルに指定あり(https://github.com/Homebrew/homebrew-core/tree/master/Formula)

3. 使用するコミットまで戻す

  ```bash
  $ git checkout [ハッシュ] [パッケージ名].rb
  ```

4. 現バージョンのアンインストール

  ```bash
  $ brew uninstall [パッケージ名]
  ```

5. 旧バージョンをインストールする

  * 先頭の`HOMEBREW_NO_AUTO_UPDATE=1`を指定すると、アップデートせずに現在のコミットでインストールできる

  ```bash
  $ HOMEBREW_NO_AUTO_UPDATE=1 brew install [パッケージ名]
  ```

6. バージョンの確認

  ```bash
  $ [パッケージ名] --version
  ```

7. リポジトリの修復

  * Homebrewのリポジトリが古い状態になっているので元に戻す

  ```bash
  $ git reset HEAD
  $ git checkout .
  ```

> 参考記事：Homebrewのインストール方法が変わってしまった
> https://creepfablic.site/2021/03/13/homebrew-python-rebase/

> 間違ってPython関連で`rm /usr/local/bin/`下の内容を消した時の対処法
> https://qiita.com/hiro_rookie/items/f237c786cc5254de48ad


| 版     | 年/月/日   |
| ------ | ---------- |
| 初版   | 2019/07/13 |
| 第二版 | 2021/03/29 |
