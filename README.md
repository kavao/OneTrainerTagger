# FineTuning用元データタグ付け補助 for OneTrainer

## 概要
このツールは、AI画像生成モデルのファインチューニング用データセット作成を支援するGUIアプリケーションです。画像ファイルと同名のテキストファイルを一括で作成し、画像の説明文（プロンプト）を自動生成します。

主な機能：
- 画像ファイル名からプロンプトを自動生成
- アンダースコア区切りをカンマ区切りに変換
- UUIDやシリアル番号の自動除去
- 共通タグの一括追加

## 使い方
1. アプリケーションを起動
2. 「ディレクトリを選択」ボタンで画像ファイルが含まれるフォルダーを選択
3. 「追加するタグ」欄に共通で付けたいタグを入力（デフォルトで`, best quality, masterpiece`が設定済み）
4. 「処理実行」ボタンをクリックして処理を開始

例：
- 入力ファイル名：`bamboo_forest_07fcd401-1bff-42fd-bab8-8d07bac1fece_0.png`
- 生成されるテキストファイル名：`bamboo_forest_07fcd401-1bff-42fd-bab8-8d07bac1fece_0.txt`
- テキストファイルの内容：`bamboo, forest, best quality, masterpiece`

### 必要な環境
- Go 1.21以上
- Git

## 対応ファイル形式
- JPG/JPEG
- PNG
- GIF

## コンパイル方法

### リポジトリのクローン

```bash
git clone [リポジトリURL]
cd [リポジトリ名]
```

### 依存関係のインストール

```bash
go mod download
go mod tidy
```

### デバッグ
```bash
go run main.go
```

### ビルド

```bash
go build
```

### 実行
```bash
./imageprocessor # Linuxの場合
imageprocessor.exe # Windowsの場合
```

## ライセンス
MIT License

Copyright (c) 2024 satoshi.okawa

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

