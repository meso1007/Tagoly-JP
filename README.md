# Tagoly-JP

Tagoly は、Git コミットをスマートに支援する CLI ツールです。  
スコープ検出・カスタムタグ対応・インタラクティブな選択など、手動でのコミットメッセージ作成を効率化します。

## 🚀 主な機能

- 🔍 **自動スコープ検出**  
  変更されたファイルパスから自動的にスコープを判定

- 🏷️ **カスタムタグ対応**  
  `.tagolyrc` に自分専用のタグを定義可能（例：ci, perf など）

- 💬 **対話的コミット生成**  
  commit type / scope / message を順にインタラクティブに選択

- 🧠 **スマートスコープ選択**  
  複数スコープが含まれる場合は最も多いものを自動選択  
  必要に応じて手動で選択可能

## 📦 インストール方法（手動）

### macOS (Apple Silicon)
```bash
mv tagoly-jp-darwin-arm64 /usr/local/bin/tagoly && chmod +x /usr/local/bin/tagoly
```

## macOS (Intel)
```bash
mv tagoly-jp-darwin-amd64 /usr/local/bin/tagoly && chmod +x /usr/local/bin/tagoly
```

### Linux
```bash
mv tagoly-jp-linux-amd64 /usr/local/bin/tagoly && chmod +x /usr/local/bin/tagoly
```

## Windows(PowerShell)
```bash
Move-Item .\tagoly.exe "C:\Program Files\tagoly\tagoly.exe"
```

## 設定ファイル .tagolyrc
```json
{
  "customTags": ["ci", "perf"]
}
```

## 使用方法
```bash
git add .
tagoly
```

## リポジトリ
https://github.com/meso1007/tagoly-jp


---
