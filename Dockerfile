FROM mcr.microsoft.com/devcontainers/go:1-1.24-bookworm

# 必要なパッケージのアップデートと追加パッケージインストール
RUN apt-get update && apt-get install -y \
    git \
    curl \
    && rm -rf /var/lib/apt/lists/*

# Go CLIツールをインストール
RUN go install github.com/swaggo/swag/cmd/swag@latest \
    && go install github.com/evilmartians/lefthook@latest \
    && go install golang.org/x/tools/cmd/goimports@latest \
    && go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest \
    && go install github.com/air-verse/air@latest \
    && go install github.com/go-task/task/v3/cmd/task@latest

# 環境変数にgoバイナリを追加（既に設定されているが明示的に追加）
ENV PATH=$PATH:/go/bin

# 作業ディレクトリを設定
WORKDIR /workspace