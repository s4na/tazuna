# tazuna

`tazuna` と打つと `hello world` と表示されるだけの小さな CLI ツール。

## インストール

### Homebrew (予定)

```sh
brew tap s4na/tazuna
brew install tazuna
```

### go install

```sh
go install github.com/s4na/tazuna@latest
```

### ソースからビルド

```sh
git clone https://github.com/s4na/tazuna.git
cd tazuna
go build -o tmp/tazuna .
```

## 使い方

```sh
$ tazuna
hello world
```
