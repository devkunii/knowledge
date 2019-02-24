03 scikit-learn
===============

* `scikit-learn`：機械学習に置いて最も重要なライブラリ

  * オープンソースプロジェクト(自由に利用し再配布し、誰でもソースコードを見て、裏側で何が起こっているかを確かめることができる)

  * 常に開発と改良が続けられており、非常に活発なユーザコミュニティを持つ

  * 様々な最先端の機械学習アルゴリズムが用意されている

  * 個々のアルゴリズムに対して[包括的なドキュメント](http://scikit-learn.org/stable/documentation)も用意されている

  * 産業界でも大学でも広く使われており、様々なチュートリアルやコード例がweb上に存在する

  * 他の科学技術向けPythonツール群を組み合わせて使うこともできる

* [scikit-learnのユーザガイド](http://scikit-learn.org/stable/user_guide.html)



## 1. scikit-learnのインストール

* scikit-learnは、`NumPy`と`SciPy`という2つのPythonパッケージに依存している

* グラフ描写とインタラクティブな開発を行うには、`matplotlib`、`IPython`、`Jupiter Notebook`をインストールする必要がある

* 必要なパッケージがはじめから含まれている下記のパッケージ済みディストリビューションのいずれかを使用する

1. [Anaconda](https://store.contiuum.io/cshop/anaconda/)

  * 大規模データ処理、予測解析、科学技術計算向けのPythonディストリビューション

  * `NumPy`、`SciPy`、`matplotlib`、`IPython`、`Jupiter Notebook`、`scikit-learn`が含まれている

  * Mac OS、Windows、Linux用が用意されている

  * 現在商用のIntel MKLライブラリが無料で含まれている

    => scikit-learnに含まれる多くのアルゴリズムが大幅に高速化される

1. [Enthought Canopy](https://www.enthought.com/products/canopy/)

  * `NumPy`、`SciPy`、`matplotlib`、`pandas`、`IPython`が含まれている

  * 無償のバージョンには`scikit-learn`が含まれていない(学位が授与できる組織に属していれば無料)

  * Python2.7用で、Mac OS、Windows、Linuxで利用できる

1. [Python(x, y)](http://python-xy.github.io/)

  * Windows向けの無償の科学技術用Pythonディストリビューション

  * `NumPy`、`SciPy`、`matplotlib`、`pandas`、`IPython`、`scikit-learn`が含まれている



| 版 |  年月日   |
|---|----------|
|初版|2019/02/16|
