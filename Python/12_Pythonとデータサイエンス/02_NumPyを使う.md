02 NumPyを使う
=============

* NumPyに内臓されている`array`(`ndarray`)を使うと、多次元の配列を使った処理を手軽に、かつ高速に実行できる

  * 科学技術計算を初めてとして、確率や統計、それらを元にした機械学習や人工知能、データサイエンスでは、大量のデータを処理するのに多くの配列を使う

  => NumPyのarrayを使うことが、そのような分野でPythonを活用するための入り口となる

* `array`はPythonのリスト型と似たデータ型

  * インデックスやスライスを使った要素のアクセスなど、部分的にリスト型と同じ操作ができる

  * 乱数などを使った生成、配列に対する演算、行列演算や線形代数演算など、リスト型が持たない便利な機能を持っている



## NumPyのarrayを生成する

* `array`の作成

  * シーケンス(リスト型)からarrayを作成する

  ```python
  import numpy as np
  a = np.array([0, 1, 2, 3])
  a
  ```

  ```python
  array([0, 1, 2, 3])
  ```

  $$
  A = \left[
    \begin{array}{rrr}
      0 & 1 & 2 & 3 \\
    \end{array}
  \right]
  $$

  * シーケンスのシーケンスを渡すと、2次元のarrayを生成することができる

  ```python
  import numpy as np
  b = np.array([[0, 0, 0], [0, 0, 0], [0, 0, 0]])
  b
  ```

  ```python
  array([[0, 0, 0],
       [0, 0, 0],
       [0, 0, 0]])
  ```

  $$
  A = \left[
    \begin{array}{rrr}
      0 & 0 & 0 \\
      0 & 0 & 0 \\
      0 & 0 & 0
    \end{array}
  \right]
  $$

* `array`オブジェクトには、いくつかのアトリビュートがある

  * アトリビュートには、arrayの情報が入っている

  ```python
  print(b.ndim)      # 次元数
  print(b.shape)     # 各次元の要素数
  print(b.size)      # サイズ
  print(b.dtype)     # 型
  ```

  ```python
  2
  (3, 3)
  9
  int64
  ```

  * データ型は初期化に使うデータの種類で自動的に決まる

    * arrayを生成するときに`dtype`という引数を渡すと、型を変更できる

* 表.NumPyのarray生成関数

| 関数名                                    | 解説                                                                                                                                                             |
| :---------------------------------------- | :--------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `np.matrix()`                             | "1 2; 3 5"のような文字列からarrayを生成する                                                                                                                      |
| `np.arange([開始値,]終了値[, ステップ])`  | 増減する数値を使ってarrayを生成する。組み込み関数の`range()とほぼ同じ引数が使える。ステップ数には、整数だけでなく、小数を含む数を使える`                         |
| `np.ones(要素数)`                         | 要素数分の1で埋め尽くされたarrayを生成する。引数にタプルやリストを渡すと、多次元のarrayを生成する                                                                |
| `np.zeros(要素数)`                        | `np.ones()`と似た関数で、0で埋め尽くされたarrayを生成する                                                                                                        |
| `np.linspace(開始値, 終了値, 要素数)`     | 開始値から終了値まで、区間を均等に要素数分並べたarrayを作る                                                                                                      |
| `np.random.rand(要素数0[, 要素数1, ...])` | 0から1までの乱数を使い、arrayを生成する。要素数を複数個渡すと、多次元のarrayを生成できる。`np.random.randn()`を使うと、標準正規分布に従う乱数からarrayを生成する |

* 一度作ったarrayは、`reshape()`メソッドを呼び出すことで形を返すことができる

  * 例)2次元配列に変換する

    => arrayのshepeアトリビュートを、`b2.shape = 3, 3`のように書き換えることでも、`reshape()`メソッドと同じ操作を実行できる

  ```python
  b2 = np.zeros(9).reshape(3, 3)
  b2
  ```

  ```python
  array([[0., 0., 0.],
       [0., 0., 0.],
       [0., 0., 0.]])
  ```

  $$
  A = \left[
    \begin{array}{rrr}
      0. & 0. & 0. \\
      0. & 0. & 0. \\
      0. & 0. & 0.
    \end{array}
  \right]
  $$

  * arrayオブジェクトの`T`というアトリビュートには、X軸とY軸を入れ替えた(列と行を入れ替え90度回転した)配列が入っている

  ```python
  a = np.arange(9).reshape(3, 3)
  a
  a.T
  ```

  ```python
  # 作成した配列
  array([[0, 1, 2],
       [3, 4, 5],
       [6, 7, 8]])

  # 90度回転した配列
  array([[0, 3, 6],
       [1, 4, 7],
       [2, 5, 8]])
```

  $$
  A = \left[
    \begin{array}{rrr}
      0 & 1 & 2 \\
      3 & 4 & 5 \\
      6 & 7 & 8
    \end{array}
  \right]
  $$

  $$
  A\times$T$ = \left[
    \begin{array}{rrr}
      0 & 3 & 6 \\
      1 & 4 & 7 \\
      2 & 5 & 8
    \end{array}
  \right]
  $$



## arrayを使った演算

* NumPyのarrayと演算子を組み合わせると、配列の各要素に対する演算が実行できる

* 例)各要素に1を足す

```python
a = np.arange(1, 10)
a+1
```

```python
array([ 2,  3,  4,  5,  6,  7,  8,  9, 10])
```

$$
A = \left[
  \begin{array}{rrr}
    1 & 2 & 3 & 4 & 5 & 6 & 7 & 8 & 9\\
  \end{array}
\right]\\
+
\left[
\begin{array}{rrr}
  1 & 1 & 1 & 1 & 1 & 1 & 1 & 1 & 1\\
\end{array}
\right]\\
= \left[
\begin{array}{rrr}
  2 & 3 & 4 & 5 & 6 & 7 & 8 & 9 & 10\\
\end{array}
\right]
$$

* 複数のarrayを組み合わせると、各要素を使った演算ができる

  * ただし、同じ型(shape)のarrayを組み合わせる必要がある

  * リストとリストを足し算すると連結になるが、NumPyのarrayでは加減乗除となる

```python
a = np.arange(1, 10)
b = np.arange(1, 10)
a+b
```

```python
array([ 2,  4,  6,  8, 10, 12, 14, 16, 18])
```

$$
A = \left[
  \begin{array}{rrr}
    1 & 2 & 3 & 4 & 5 & 6 & 7 & 8 & 9\\
  \end{array}
\right]\\
B = \left[
  \begin{array}{rrr}
    1 & 2 & 3 & 4 & 5 & 6 & 7 & 8 & 9\\
  \end{array}
\right]\\
\\
A+B=\left[
  \begin{array}{rrr}
    2 & 4 & 6 & 8 & 10 & 12 & 14 & 16 & 18\\
  \end{array}
\right]\\
$$

* array同士の演算は形が異なると実行できない

  * しかし、列、または行のどちらかが同じ要素数を持つarray同士であれば演算が可能

  * 例)3×3のarrayに、3×1のarrayをかける

  ```python
  a = np.ones(9).reshape(3, 3)      # 1だけで構成された3×3のarrayを作る
  b = np.arange(1, 4)               # 1,2,3のarrayを作る
  a*b                               # 掛け算した結果を表示
  ```

  ```python
  array([[1., 2., 3.],
       [1., 2., 3.],
       [1., 2., 3.]])
  ```

  $$
  A = \left[
    \begin{array}{rrr}
      1 & 1 & 1 \\
      1 & 1 & 1 \\
      1 & 1 & 1
    \end{array}
  \right]
  \\
  B = \left[
  \begin{array}{rrr}
      1 \\
      2 \\
      3 \\
    \end{array}
  \right]
  \\
  A\times B =\left[
    \begin{array}{rrr}
      1 & 2 & 3 \\
      1 & 2 & 3 \\
      1 & 2 & 3
    \end{array}
  \right]
  $$

  => これを、`ブロードキャスティング`と呼ぶ

  * 例)3×1と1×3のarrayを組み合わせて、9×9のarrayを作る

  ```python
  np.zeros((3, 1))*np.zeros((1, 3))
  ```

  ```python
  array([[0., 0., 0.],
       [0., 0., 0.],
       [0., 0., 0.]])
  ```
  $$
  A = \left[
    \begin{array}{rrr}
      0 & 0 & 0 \\
    \end{array}
  \right]
  \\
  B = \left[
  \begin{array}{rrr}
      0 \\
      0 \\
      0 \\
    \end{array}
  \right]
  \\
  A\times B =\left[
    \begin{array}{rrr}
      0 & 0 & 0 \\
      0 & 0 & 0 \\
      0 & 0 & 0
    \end{array}
  \right]
  $$

* NumPyに搭載されている関数を使うことで、arrayの合計を計算することができる

  * 1次元であっても多次元であっても、同様に合計を計算できる

  ```python
  
  ```
