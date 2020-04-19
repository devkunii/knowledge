# 顧客id
id <- c(1:3)

# 名前
name <- list("秀和太郎",
             "築地花子",
             "宗田解析")

# 住所
add <- list("中央区築地100-1",
            "中央区築地本町200",
            "中央区日本橋99")

# リストを作成
add_book <- list("顧客リスト", id, name, add)

# リスト要素の取り出し
list1 <- add_book[[1]]
list2 <- add_book[[2]]
list3 <- add_book[[3]]
list4 <- add_book[[4]]

# リスト要素の特定の要素を取り出す
cat1_id <- add_book[[2]][[1]]
cat1_name <- add_book[[3]][[1]]
cat1_add <- add_book[[4]][[1]]

# リストの第1、第2要素をリストとして取り出す
var1 <- add_book[c(1, 2)]

# リストとして取得した要素をベクトルにする
var2 <- unlist(add_book[3])

# リストの要素を変更する
add_book[[3]][[1]] <- "築地太郎"

# リストの要素、サブリストの要素を削除する
add_book[[3]][[1]] <- NULL
add_book[[3]] <- NULL

add_book2 <- list(
                  id = c(1:3),
                  name = list("秀和太郎", "築地花子", "宗田解析"),
                  add = list("中央区築地100-1",
                             "中央区築地本町200",
                             "中央区日本橋99")
                  )
