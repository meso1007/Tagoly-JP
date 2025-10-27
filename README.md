# 🏷️ tagoly-jp

`tagoly-jp` は、ファイルや画像などに簡単にタグを付けて管理できる **軽量コマンドラインツール** です。  
Go 言語で開発されており、macOS・Linux・Windows で動作します。

---

## 🚀 主な機能

- ファイルへのタグ追加・削除
- タグによるファイル検索
- ファイルごとのタグ一覧表示
- 複数プラットフォーム対応（macOS / Linux / Windows）
- 日本語 CLI 表示に対応

---

## 📦 インストール方法

### ① リリースページからダウンロード
[最新のリリースはこちら](https://github.com/YOUR_USERNAME/tagoly-jp/releases)

環境に合わせて以下のファイルをダウンロードしてください：

| OS | アーキテクチャ | ファイル名 |
|----|----------------|------------|
| macOS (Apple Silicon) | arm64 | tagoly-jp-darwin-arm64 |
| macOS (Intel) | amd64 | tagoly-jp-darwin-amd64 |
| Linux | amd64 | tagoly-jp-linux-amd64 |
| Windows | amd64 | tagoly-jp.exe |

### ② 実行権限を付与（macOS / Linux）
```bash
chmod +x tagoly-jp-darwin-arm64
```

### ③ 任意の場所に移動（例：/usr/local/bin）
``` bash
sudo mv tagoly-jp-darwin-arm64 /usr/local/bin/tagoly-jp
```