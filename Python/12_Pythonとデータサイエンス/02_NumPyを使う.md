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

  $`A = \left[
    \begin{array}{rrr}
      0 & 1 & 2 & 3 \\
    \end{array}
  \right]`$

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

  $`A = \left[
    \begin{array}{rrr}
      0 & 0 & 0 \\
      0 & 0 & 0 \\
      0 & 0 & 0
    \end{array}
  \right]`$

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

  $`A = \left[
    \begin{array}{rrr}
      0. & 0. & 0. \\
      0. & 0. & 0. \\
      0. & 0. & 0.
    \end{array}
  \right]`$

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

  $`A = \left[
    \begin{array}{rrr}
      0 & 1 & 2 \\
      3 & 4 & 5 \\
      6 & 7 & 8
    \end{array}
  \right]`$

  $`A\times T = \left[
    \begin{array}{rrr}
      0 & 3 & 6 \\
      1 & 4 & 7 \\
      2 & 5 & 8
    \end{array}
  \right]`$



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


$`A = \left[
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
\right]`$

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

$`A = \left[
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
\right]\\`$

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

  $`A = \left[
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
  \right]`$

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
  $`A = \left[
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
  \right]`$

* NumPyに搭載されている関数を使うことで、arrayの合計を計算することができる

  * 1次元であっても多次元であっても、同様に合計を計算できる

  ```python
  a = np.arange(9).reshape(3, 3)
  np.sum(a)
  ```

  ```python
  36
  ```

  $`A = \left[
    \begin{array}{rrr}
      0 & 1 & 2 \\
      3 & 4 & 5 \\
      6 & 7 & 8 \\
    \end{array}
  \right]
  \\
  0+1+2+3+4+5+6+7+8=36`$

  * `np.sum()`に、`axis`(軸)という引数を渡すと、軸を指定して合計を計算する方向を指定できる

  * 例)軸ごとに計算する

  ```python
  a = np.arange(9).reshape(3, 3)
  np.sum(a, axis=0)
  ```

  ```python
  array([ 9, 12, 15])
  ```

  $`A = \left[
    \begin{array}{rrr}
      0 & 1 & 2 \\
      3 & 4 & 5 \\
      6 & 7 & 8 \\
    \end{array}
  \right]
  \\
  0+3+6=9
  \\
  1+4+7=12
  \\
  2+5+8=15
  \\
  B= \left[
    \begin{array}{rrr}
      9 & 12 & 15 \\
    \end{array}
  \right]
  \\`$


  * `np.sum()`の他に、

    * `np.mean()`：平均を計算する

    * `np.median()`：中央値を計算する

    * `np.max()`：最大値を計算する

    * `np.min()`：最小値を計算する

    => どの関数も`axis`という引数をとり、軸を指定することができる

  * `np.dot()`関数では、arrayを行列に見立てて積を計算することができる

    => 行列の積を計算するための演算子`@`を使うと、`np.dot()`と同じような操作が可能



## 要素へのアクセス

* arrayの要素へは、リスト型と同じくインデックスを使ってアクセスできる

* 多次元の配列は、`[1][2]`のようにインデックスを連ねた方法の他、カンマで複数のインデックスを渡すことでもアクセスできる

* 例)要素へアクセスする

```python
a = np.arange(9).reshape(3, 3)
a[1, 2]
```

```python
5
```
$`A = \left[
  \begin{array}{rrr}
    0 & 1 & 2 \\
    3 & 4 & 5 \\
    6 & 7 & 8 \\
  \end{array}
\right]`$

* インデックスを使ってarrayの要素を指定して代入を行うと、リストと同じく要素の入れ替えが行える

  * 代入によってarrayの型(dtype)は変わらない

  * 例)整数要素を持つarrayの要素に、1.5のような浮動小数点数を代入しても、小数点以下を切り捨てた状態で代入が行われる。また、要素の削除は行えない

* arrayではリスト型と同じくスライスも使える

  * スライスを組み合わせると、多次元配列の一部を多次元配列として取り出せる

  * 例)2×2の部分を取り出す

  ```python
  a[1:, 1:3]
  ```

  ```python
  array([[4, 5],
       [7, 8]])
  ```

  $`A = \left[
    \begin{array}{rrr}
      4 & 5 \\
      7 & 8 \\
    \end{array}
  \right]`$

* また、インデックスとしてリストを渡すと、リスト上のインデックスに見立てて複数の値を取り出せる

  * 例)1から9までのarrayを作り、インデックスとしてリストを渡し、偶数だけ取り出す

  ```python
  d = np.arange(1, 10)
  d[[1, 3, 5, 7]]
  ```

  ```python
  array([2, 4, 6, 8])
  ```

  $`A = \left[
    \begin{array}{rrr}
      2 & 4 & 6 & 8\\
    \end{array}
  \right]`$



## arrayの連結

* NumPyのarray同士を連結するには、関数が用意されている

* `np.hstack()`：横方向に連結する


  ```python
  a = np.arange(4).reshape(2, 2)
  b = np.arange(5, 9).reshape(2, 2)
  np.hstack((a, b))
  ```

  ```python
  array([[0, 1, 5, 6],
       [2, 3, 7, 8]])
  ```

  $`A = \left[
    \begin{array}{rrr}
      0 & 1\\
      2 & 3\\
    \end{array}
  \right]
  \\
  B = \left[
    \begin{array}{rrr}
      5 & 6\\
      7 & 8\\
    \end{array}
  \right]
  \\
  hstack = \left[
   \begin{array}{rrr}
     0 & 1 & 5 & 6\\
     2 & 3 & 7 & 8\\
   \end{array}
  \right]`$

* `np.vstack()`：縦方向に連結する

```python
a = np.arange(4).reshape(2, 2)
b = np.arange(5, 9).reshape(2, 2)
np.vstack((a, b))
```

```python
array([[0, 1],
       [2, 3],
       [5, 6],
       [7, 8]])
```

$`vstack = \left[
 \begin{array}{rrr}
   0 & 1\\
   2 & 3\\
   5 & 6\\
   7 & 8\\
 \end{array}
\right]`$

* 1次元の配列であれば、`np.column_stack()`という関数を使って縦に積むことができる



## arrayのコピー

* array同士の代入

  * 例)あるarray(a)を他の変数(b)に代入する。この時、bのarrayに対して操作を行うとaのarrayはどうなるか

```python
a = np.zeros(4)
b = a
b += 1
a
```

```python
array([1., 1., 1., 1.])
```

  * bに対して行なった操作が、aにも実行されている

    => Pythonの代入はコピーではなく参照なので、`b = a`とすることで`a`も`b`も同じarrayオブジェクトを指すようになる

    => スライスについても同じ

* `copy`：あるarrayを別のオブジェクトとして取り出す

```python
a = np.zeros(4)
b = a.copy()
b += 1
print(a)
```

```python
[0. 0. 0. 0.]
```



| 版 |  年月日   |
|---|----------|
|初版|2019/02/09|
