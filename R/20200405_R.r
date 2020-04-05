par(family= "HiraKakuProN-W3")

# data <- read.table(         # 店舗別売上.txtをデータフレームとしてdataに代入
#   "/Users/kunii.sotaro/Desktop/r_sample/chap04/sec01/dispersion/販売台数.txt",
#   header=T,                 # 1行目は列名であることを指定
#   fileEncoding="CP932"      # 文字コードをShift_JISに指定
# )
# hist(data$車種A)            # 車種Aのヒストグラムを作成
# hist(data$車種B)            # 車種Bのヒストグラムを作成

# me_A = mean(data$車種A)     # 車種Aの平均を求める
# num_A <- data$車種A         # 車種Aの販売台数をベクトルに代入
# dev_A <- num_A - me_A       # 車種Aの偏差を求める
# cat("偏差：",dev_A, "\n")   # 出力

# me_B = mean(data$車種B)     # 車種Bの平均を求める
# num_B <- data$車種B         # 車種Bの販売台数をベクトルに代入
# dev_B <- num_B - me_B       # 車種Bの偏差を求める
# cat("偏差：",dev_B, "\n")   # 出力

# 車種Aの分散を求める
# dspr_A <- sum(dev_A^2) / length(data$車種A)
# cat("分散：", dspr_A, "\n")
# 車種Bの分散を求める
# dspr_B <- sum(dev_B^2) / length(data$車種A)
# cat("分散：", dspr_B, "\n")

# getDisper <- function(x) {       # 分散を返す関数
#   dev <- x - mean(x)             # 偏差を求める
#   return(sum(dev^2) / length(x)) # 分散を返す
# }

# data <- read.table(           # 店舗別売上.txtをdataに代入
#   "販売台数.txt",
#   header=T,                   # 1行目は列名であることを指定
#   fileEncoding="CP932"        # 文字コードをShift_JISに指定
# )

# num_A  <- data$車種A          # 車種Aのデータをベクトルに代入
# dspr_A <- getDisper(num_A)    # 車種Aの分散を求める
# sd_A   <- sqrt(dspr_A)        # 車種Aの標準偏差を求める
# cat("標準偏差：", sd_A, "\n") # 出力
#
# num_B  <- data$車種B          # 車種Bのデータをベクトルに代入
# dspr_B <- getDisper(num_B)    # 車種Bの分散を求める
# sd_B   <- sqrt(dspr_B)        # 車種Bの標準偏差を求める
# cat("標準偏差：", sd_B, "\n") # 出力

# getSd <- function(x) {                  # 標準偏差を返す関数
#   return(                               # 戻り値を返す
#     sqrt(                               # 分散の平方根
#       sum((x - mean(x))^2) / length(x)) # 分散の計算式
#     )
# }

# data <- read.table(           # 店舗別売上.txtをdataに代入
#   "販売台数.txt",
#   header=T,                   # 1行目は列名であることを指定
#   fileEncoding="CP932"        # 文字コードをShift_JISに指定
# )

# num_A <- data$車種A           # 車種Aのデータをベクトルに代入
# sd_A <- getSd(num_A)          # 車種Aの標準偏差を求める
# cat("標準偏差：", sd_A, "\n") # 出力
#
# num_B <- data$車種B           # 車種Bのデータをベクトルに代入
# sd_B <- getSd(num_B)          # 車種Bの標準偏差を求める
# cat("標準偏差：", sd_B, "\n") # 出力

# getSd <- function(x) {                  # 標準偏差を返す関数
#   return(                               # 戻り値を返す
#     sqrt(                               # 分散の平方根
#       sum((x - mean(x))^2) / length(x)) # 分散の計算式
#   )
# }

# data <- read.table(           # 店舗別売上.txtをdataに代入
#   "販売台数.txt",
#   header=T,                   # 1行目は列名であることを指定
#   fileEncoding="CP932"        # 文字コードをShift_JISに指定
# )

# num_A  <- data$車種A          # 車種Aのデータをベクトルに代入
# mean_A <- mean(num_A)         # 車種Aの平均を求める
# sd_A   <- getSd(num_A)        # 車種Aの標準偏差を求める
# std_A  <-
#   (num_A - mean_A) / sd_A     # 標準化(データ-平均)÷標準偏差)
# print(std_A)                  # 出力

# num_B  <- data$車種B          # 車種Bのデータをベクトルに代入
# mean_B <- mean(num_B)         # 車種Bの平均を求める
# sd_B   <- getSd(num_B)        # 車種Bの標準偏差を求める
# std_B  <-
#   (num_B - mean_B) / sd_A     # 標準化(データ-平均)÷標準偏差)
# print(std_B)                  # 出力

# getSd <- function(x) {                  # 標準偏差を返す関数
#   return(                               # 戻り値を返す
#     sqrt(                               # 分散の平方根
#       sum((x - mean(x))^2) / length(x)) # 分散の計算式
#   )
# }

# data <- read.table(       # 来店者数.txtをdataに代入
#   "/Users/kunii.sotaro/Desktop/r_sample/chap04/sec03/store traffic/来店者数.txt",
#   header=T,               # 1行目は列名であることを指定
#   fileEncoding="CP932"    # 文字コードをShift_JISに指定
# )
#
# num  <- data$来店者数     # 来店者数のデータをベクトルに代入
# sd   <- getSd(num)        # 来店者数の標準偏差を求める
# mean <- mean(num)         # 来店者数の平均
# 標準化係数をデータフレームにバインド
new_frm <- cbind(
              # もとのデータフレーム
              data,
              # 標準化係数のデータフレーム
              data.frame(
                # 標準化の計算
                (num - mean) / sd
              )
           )

           # 0～2を0.01刻みにしたシーケンスを生成
           n <- seq(0, 2, by = 0.01)
           # 標準正規分布の確率密度関数で確率を求める
           dn <- dnorm(n, mean=0, sd=1)

           # 区間面積用の数値ベクトルを作る
           s_area <- as.numeric(NULL)
           # 累積面積用の数値ベクトルを作る
           t_area <- as.numeric(NULL)

           # 区間面積と累計面積を計算
           i <- 0                             # カウンター変数
           for (value in n){                  # nの要素のぶんだけ繰り返す
             if(value == 0){                  # nが0であれば実行
               s_area <- 0                    # 区間面積を0とする
               t_area <- 0                    # 累計面積を0とする
               i <- i + 1                     # カウンターの値を1増やす
             } else {                         # nの要素が0以外の場合の処理
               # 区間面積のベクトルに区間面積を追加
               s_area <- c(
                 s_area,                      # 代入先のベクトル
                 (dn[i] + dn[i+1]) * 0.01 / 2 # 区間面積を台形の面積で近似
                 )
               # 累計面積のベクトルに累計面積を追加
               t_area <- c(
                 t_area,                      # 代入先のベクトル
                 t_area[i] + s_area[i+1]      # 累計の面積を計算
                 )
               i <- i + 1                     # カウンターの値を1増やす
             }
           }
           # 標準正規分布の数表を作成
           dframe <- data.frame(
             x_value      = n,                # 0～2までの連続値
             section_area = s_area,           # 区間面積
             total_area   = t_area            # 累計面積
             )
