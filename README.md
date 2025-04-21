# 概要
GO(フレームワーク : Gin) + Docker + Docker-compose + Nginx + Mysqlの環境を構築した。その動作手順について、記載する

# 前提条件
- [Homebrew](https://brew.sh/ja/)がインストールされている
- [Docker Desktop](https://formulae.brew.sh/cask/docker)がインストールされている
- [Go言語（Golang）](https://formulae.brew.sh/formula/go)がインストールされている

# 必要ファイルの作成
- `phpmyadmin/conf`に`.env`を作成して、下記の内容を記述する
```sh
PMA_ARBITRARY=1
PMA_HOSTS=db # ⇦ docker-composeのサービス名を記載する
PMA_USER=root # ⇦ 任意の値を付与しても良い
PMA_PASSWORD=password # ⇦ 任意の値を付与しても良い
```

- `db/conf`に`.env`を作成して、下記の内容を記述する
```sh
MYSQL_DATABASE=example_db # ⇦ 任意の値を付与しても良い
MYSQL_ROOT_PASSWORD=password # ⇦ 任意の値を付与しても良い
TZ=Asia/Tokyo
```

# 動作手順
- Dockerのイメージを作成した後、起動を行う
```sh
# イメージ作成
docker compose build --no-cache

# 起動
docker compose up -d
```

# 補足コマンド
- コンテナー内に入るコマンド
```sh
# go環境
docker exec -it web1 /bin/bash

# nginx環境
docker exec -it nginx /bin/bash

# db環境
docker exec -it db /bin/bash
```


- `go.mod`の更新コマンド
```sh
# コンテナー内に入り行う
go mod tidy

# コンテナー内に入らずに行う
docker compose exec web1 go mod tidy
```

- 削除
```sh
docker-compose down --rmi all --volumes  --remove-orphans && docker volume prune -f
```

- 全体削除
```sh
docker system prune -f
```