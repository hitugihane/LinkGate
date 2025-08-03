# LinkGate

## 📦 使用技術

- [Go 1.24](https://go.dev/)
- [Dev Containers](https://containers.dev/)
- [Visual Studio Code](https://code.visualstudio.com/)
- 拡張機能（自動インストール）:
  - Go
  - GitLens
  - TODO Tree
  - Error Lens
  - Code Spell Checker
  - 日本語言語パック

---

## 🛠️ 前提条件

以下をインストールしておいてください。

- [Visual Studio Code](https://code.visualstudio.com/)
- [Dev Containers 拡張機能](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers)
- [Docker](https://www.docker.com/)

---

## 🚀 開発環境の立ち上げ手順

### 1. リポジトリをクローン

```bash
git clone https://github.com/fugafuga-2025/LinkGate.git
```
```bash
cd LinkGate
```

### 2. 開発環境の立ち上げ
1. Docker Desktopを立ち上げる


2. 右下に表示される`コンテナーで再度開く`を押すか、
VS Code 上でコマンドパレットを開き、`Dev Containers: Reopen in Container` を選択

初回起動時は Docker イメージのビルドや拡張機能のインストールに数分かかります。