FROM golang:latest

ENV LANG=C.UTF-8 \
    TZ=Asia/Tokyo \
    WEB1_ROOT=/web1 \
    APP_ENV=development

WORKDIR ${WEB1_ROOT}

# 開発支援ツールをインストール（AirとDelve）
RUN go install github.com/air-verse/air@latest && \
    go install github.com/go-delve/delve/cmd/dlv@latest

# 依存関係のみ先にコピーしてダウンロード（キャッシュ用）
COPY go.mod go.sum ./
RUN go mod download 

# アプリケーションソースをコピー
COPY . .

RUN go mod tidy

# ホットリロードをデフォルトコマンドに設定（.air.tomlがある場合読み込まれる）
CMD ["air", "-c", ".air.toml"]