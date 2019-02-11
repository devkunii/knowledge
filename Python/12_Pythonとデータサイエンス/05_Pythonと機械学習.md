05 Pythonと機械学習
==================

* `機械学習`：人工知能の研究分野の一つ

  * 人工知能の研究は、人間が行なっている認知や判断といった比較的高度な機能を、コンピュータのような機械に行わせることを目的にしている

  * `機械学習`は、そのような研究のうち、コンピュータを使って人間の学習と同じような機能を実現しようという研究において始まった

* 人間の行う学習を、「入力(見る、聞く、触る)を使ってふさわしい出力(反応)を学ぶこと」と定義する

  => コンピュータで同様のことができれば、コンピュータにも学習させることができる

  * コンピュータにとって入力とはデータ(数値)のこと

  * 機械学習は、以下のように表せる

![従来のプログラミングと機械学習](./images/従来のプログラミングと機械学習.png)

* たくさんのデータをコンピュータで扱うときには配列を用いる

* データを処理して、判断に繋げるには確率や統計の手法を使う

* これまでは`R`で行なってきたが、機械学習を適用する分野が多様化するにつれ、Pythonが用いられるようになった

* Pythonと`scikit-learn`を用いて、機械学習を行う



## 機械学習による数値の予測

* 機械学習では、ある値Xに対応して変化する値Yを予測するという問題を扱う

  => 過去の売り上げデータから将来の売り上げを予測することができる

* 実際に観測したXとYデータを多数用意して、法則を学習させることで、Xに対応するYを予測させる

* `最小二乗法`：多項式による近似を行い、元データをうまくなぞる曲線を見つけ出して、値の予測をする

  * `scikit-learn`には、最小二乗法を使った分析器が用意されているので、簡単に使うことができる

* まず、予測を行うNumPyを作る

  * sinに対して標準正規分布に従う乱数を足したデータを作る

  * scikit-learnでは、学習するデータを行として並べる必要がある

  => `[:, np.newaxis]`にしている理由

```python
import numpy as np

# 乱数のシードを設定
np.random.seed(9)
# 0から1まで100個の数値を生成、乱数要素を混ぜる前のx
x_orig = np.linspace(0, 1, 100)

def f(x):
    # xに対応するsinを返す関数
    return np.sin(2*np.pi*x)

# 0から1まで100個のばらけたサンプルデータ(x)を生成
x = np.random.uniform(0, 1, size=100)[:, np.newaxis]
# yに対応するsinに乱数値を足してサンプルデータ(y)を生成
y = f(x)+np.random.normal(scale=0.3, size=100)[:, np.newaxis]
```

* データができたら、後々のために学習用のデータとテスト用のデータに分割する

  * その後、元のsinの線を点線で補いながら、生成したデータをグラフにする

  * この点の間をうまく通る曲線を求める(トレーニング用のデータからデータをうまく説明できる)モデルを求める

```python
%matplotlib inline
import matplotlib.pyplot as plt
from sklearn.model_selection import train_test_split

# 学習用データとテスト用データに分ける
x_train, x_test, y_train, y_test = train_test_split(x, y, test_size=0.8)

# 元のsinとサンプルデータをplot
plt.plot(x_orig, f(x_orig), ls=':')
plt.scatter(x_train, y_train)
plt.xlim((0, 1))
```

![学習データのグラフ](./images/学習データのグラフ.png)

* 次に、データを学習してモデルを作る

  * 最小二乗法の多項式近似という手法を使う

  * データを学習する時に与えるパラメータ(deg=次数)をいくつか試してみる

  => 学習の度合いが変わる

```python
from sklearn.linear_model import LinearRegression
from sklearn.preprocessing import PolynomialFeatures
from sklearn.pipeline import make_pipeline

# 2×2のグラフを描く準備をする
fig, axs = plt.subplots(2, 2, figsize=(8, 5))

# 次数0, 1, 3, 9について学習した結果を表示
for ax, deg in zip(axs.ravel(), [0, 1, 3, 9]):
    # パイプラインを作る
    e = make_pipeline(PolynomialFeatures(deg), LinearRegression())
    # 学習セットで学習をする
    e.fit(x_train, y_train)
    # 元のxを与えて予測
    px = e.predict(x_orig[:, np.newaxis])
    # 予測結果のグラフとテストデータの点を描画
    ax.scatter(x_train, y_train)
    ax.plot(x_orig, px)
    ax.set(xlim=(0, 1), ylim=(-2, 2), ylabel='y', xlabel='x', title='degree={}'.format(deg))

plt.tight_layout()
```

![学習の様子](./images/学習の様子.png)

* 最小二乗法の多項式近似は、与えられたデータから連立方程式を解き、データの真ん中を通る曲線の式を求める

  * 次数を0：傾きのない直線

  * 次数を1：傾きのある直線

  * 次数を上げていくと、曲線が複雑になり、与えたデータに対してより近くを通る曲線の式を求めることができる

  * 曲線の式にXの値を与えると、Yの値が出てくるので、このようにしてXに対応するYの値を予測できる

* テストは以下のように行う

  * 学習データを与えてモデルを作る

  * 作ったモデルでテストデータを予測させ、実データとの誤差を調べてどの程度正確に予測できているかを計測する

  * 次数を変えながら実行して、誤差を二乗して平均した値(平均二乗誤差、RMS)をグラフに表示する

```python
from sklearn.metrics import mean_squared_error

# 実データとの誤差を保存するarray
train_error = np.empty(10)
test_error = np.empty(10)
# 次数0から9について調べる
for deg in range(10):
    # モデルを作る
    e = make_pipeline(PolynomialFeatures(deg), LinearRegression())
    e.fit(x_train, y_train)
    # テストデータを使って、予測値と実際の値の誤差を調べる
    train_error[deg] = mean_squared_error(y_train, e.predict(x_train))
    test_error[deg] = mean_squared_error(y_test, e.predict(x_test))

# グラフを描く
plt.plot(np.arange(10), train_error, ls=':', label='train')
plt.plot(np.arange(10), test_error, ls='-', label='test')
plt.ylim((0, 1))
plt.legend(loc='upper left')
```

![次数と誤差のグラフ](./images/次数と誤差のグラフ.png)

* 比較のため、学習に使ったデータのモデルを使って予測させた場合の誤差を点線で表示している

* 次数が7から先は誤差が広がっている

  => 学習データを過剰に学習して予測できるようになった代わりに、他のデータ(テストデータ)を予測できなくなった

  => `過学習(オーバーフィッティング)`

* グラフを見る限りは、次数を3に取るのが良い



## 機械学習のアルゴリズム



### 回帰

* 売り上げや価格、喫煙率と肺がんの発症率など、連続的な量的データを予測すること

* scikit-learnには、過学習を抑えながら回帰を行うアルゴリズムとして、Lasso(ラッソ回帰)や、Redge(リッジ回帰)がある

* 与えるデータに対して答えを用意するた目の`教師あり学習`と分類される



### 分類

* 学習データにラベルを与えて学習を行うことで、機械学習を使って分類を行うことがきる

* スパム判定やニュースサイトの記事の分類、画像認識などに利用される

* scikit-learnには、

  * SVC(サポートベクター分類器)

  * nearestNeightbors(近傍法)

  * RandomForestClassifier(ランダムフォレスト)

  * naive_bayes(単純ベイズ分類器のライブラリ)

  などがある

* `教師あり学習`と分類される



### クラスタリング

* 近いデータをまとめることを`クラスタリング`という

* 顧客のセグメントを分析したり、クラスタに含まれていないデータを探して以上を検知するなどの応用がある

* scikit-learnには、

  * Kmeans(K平均法)

  * MeanShift(平均変位法)

  などがある

* データだけを与えて規則性を発見するような`教師なし学習`と分類される



## 名前から性別を判定する

* scikit-learnを使った例として、スパム判定にも使われるベイズ理論を使った機械学習を試す

* 男女のラベルをつけた、ひらがなの名前を学習させて、名前から男女を見分ける分析器を作る

  => ひらがなの名前を2文字ずつに分割したデータを使う(教師あり学習)

* arrayの作成

```python
import numpy as np
from sklearn.model_selection import train_test_split

np.random.seed(9)
# 男女のタグ付きひらがなの名前データを読み込む
txtbody = open('names.txt', encoding='utf-8')
# NumPyのarrayに変換
jnames = np.array([x.split() for x in txtbody], dtype='U12')
# 名前と性別に分割
names_train, gender_train, = jnames[:, 1], jnames[:, 0]
```

  * 学習用のデータのことを`ベクトル`と言う

  * 学習用のデータを作ることを`ベクトル化`と言う

* 数値の予測の例では、学習データも予測の対象となるのも数値で、学習しやすい形式を取っていた

  => ベクトル化をする手間がほとんどなかった

* 今回は文字データを使って学習を行うため、事前の下処理が必要

  => 学習データを数値に変換して、目的のアルゴリズムで学習できるように処理する必要がある

  => 名前を2文字に区切った文字列の出現頻度を数値化する

* 元データのベクトル化をする前に、文字列を2文字ずつに分割する関数を作る

```python
def split_in_2words(name):
    return [name[i:i+2] for i in range(len(name)-1)]

split_in_2words("とものり")
```

```python
['とも', 'もの', 'のり']
```

* 次に、学習データを作ってベクトル化の前段階となるデータを作る

  * ひらがなの名前を2文字に分割した文字列の出現回数を数える

  * このような形式のデータおBoW(Bag of Words)と呼ぶ

  * scikit-learnには、文字列の出現回数を数えるためのクラス(CountVectorizer)が用意されている

* 学習用の名前リストを渡して、CountVectorizerオブジェクトを作る

```python
from sklearn.feature_extraction.text import CountVectorizer
bow_t = CountVectorizer(analyzer=split_in_2words).fit(names_train)
```

* これで学習データにある名前を2文字に分割した文字列全てに数値(ID)が振られ、`bow_t`という変数に格納された

* 次に、文字列の出現数を数える

```python
name = 'かんかん'
b1 = bow_t.transform([name])
print(b1[0])
```

```python
(0, 283)	2
(0, 1898)	1
```

* 出力データから、文字列を逆引きしてみる

  * 丸括弧の中の2番目の数字が文字列のID

```python
print(bow_t.get_feature_names()[283])
print(bow_t.get_feature_names()[1898])
```

```python
かん
んか
```

* 実際に学習データを使って文字列の出現回数を数える

```python
names_bow = bow_t.transform(names_train)
```

* 次に、TF-IDFという方法を使ってデータの重み付けと正規化を行うためのオブジェクトを作る

  * 出現数を使って、文字列がどのくらい重要かを示す数値に変換するための下準備を行う

  * scikit-learnに用意されている`TfidfTransformer()`を使い、先ほど作った`names_bow`と言うオブジェクトを渡して`fit()`を呼び出す

```python
from sklearn.feature_extraction.text import TfidfTransformer

tfidf_t = TfidfTransformer().fit(names_bow)
```

* `tfidf_t`でどのような変換がされるかを見るために、重み付けを実行する

```python
tfidf1 = tfidf_t.transform(b1)
print(tfidf1)
```

```python
(0, 1898)	0.530554460022041
(0, 283)	0.8476508508523546
```

* 学習の準備が整ったので、文字列の出現数を重み付けして、ベイズ理論を応用したアルゴリズムで学習をさせる

  * `MultinomialNB`と言う、ナイーブベイズの多項モデルと呼ばれるアルゴリズムを使う

```python
from sklearn.naive_bayes import MultinomialNB
# 文字列の重み付けと正規化を行う
names_tfidf = tfidf_t.transform(names_bow)
# 学習を実行
namegender_detector = MultinomialNB().fit(names_tfidf,  gender_train)
```

* 先ほどの名前を入力する

```python
print(namegender_detector.predict(tfidf1)[0])
```

```python
boy
```

* 次に文字列を与えて性別を予測する関数を使って、色々な名前について試してみる

```python
def predict_gender(name):
  bow = bow_t.transform([name])
  n_tfidf = tfidf_t.transform(bow)
  return namegender_detector.predict(n_tfidf)[0]
```

```python
print(predict_gender('のんな'))
```

```python
girl
```

* 今回作った分析器は、典型的な名前はうまく分類できるが、正答率は80%程度



## 機械学習、データサイエンスとPython

* データサイエンスベン図

![データサイエンスベン図](./images/データサイエンスベン図.png)



| 版 |  年月日   |
|---|----------|
|初版|2019/02/09|
