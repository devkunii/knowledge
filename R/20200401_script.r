vct1 <- c(1, 2, 3, 4, 5, 6)
vct2 <- c(10, 20, 30, 40, 50, 60)
vct3 <- c(100, 200, 300, 400, 500, 600)

mtx1 <- matrix(vct1)

mtx2 <- matrix(vct1, nrow=2)

mtx3 <- matrix(vct1, ncol=2)

array1 <- array(
                # vct1, vct2, vct3を要素にしたベクトル
                c(vct1, vct2, vct3),
                # 行数の2, 列数の3, 行列の数3を要素にしたベクトル
                (dim = c(2, 3, 3))
)

mtx4 <- matrix(vct1, 2, 2)

mtx5 <- matrix(vct1,
               nrow = 2,
               byrow = TRUE)

mtx6 <- rbind(vct1, vct2, vct3)

mtx7 <- cbind(vct1, vct2, vct3)

branch <- c(
          "初台店",
          "幡谷店",
          "吉祥寺店",
          "笹塚店",
          "明大前店"
          )

sales <- c(
           2024,
           2164,
           6465,
           2186,
           2348
          )

df <- data.frame(branch=branch, salses=sales)

df$branch # 列データを要素として取り出す
df[[1]]   # 列データを要素として取り出す
df[,1]    # 列データを要素として取り出す
df[1]     # 列データをリストとして取り出す

df[1,]    # 行データをリストとして取り出す

data <- read.table(             # 店舗別売上.txtをdataに代入
  "/Users/kunii.sotaro/Desktop/r_sample/chap02/sec02/load_file/店舗別売上.txt",
  header=TRUE,                  # 1行目は列名であることを指定
  fileEncoding="CP932"          # 文字コードをShift_JISに指定
)

num <- -10

if (num < 0){
  num <- num * -1
}

num <- 10

if (num < 0) {
  num <- num * -1
} else if (num > 0){
  num <- num * -1
}

num <- "-10"

if (is.numeric(num) & num < 0) {
  num <- num * -1
} else if (is.numeric(num) & num > 0) {
  num <- num * -1
} else {
  num <- as.numeric(num)
}

for (word in c("おはよう！", "こんにちは", "わんばんこ")) {
  # print(word)
}

data <- read.table(    # 定着度.txtをdataに代入
  "/Users/kunii.sotaro/Desktop/r_sample/chap02/sec03/for/定着度.txt",
  header=T,            # 1行目は列名であることを指定
  fileEncoding="CP932" # 文字コードをShift_JISに指定
)

j <- length(data[1,])  # 列の数を調べる

for(i in c(1:j)) {
  assign(
    sprintf("x%d", i), # xに連番を付けた名前を作る
    data[,i]           # データフレムの1列目から代入する
  )
}

show1 <- function() {
  print("Hello!")
}

show2 <- function(word1, word2) {
  print(word1)
  print(word2)
}

taxin <- function(val) {
  tax_in <- val * 1.08
  return(tax_in)
}

tax_in <- taxin(100)
