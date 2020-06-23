# struct-builder-example

## Description
CSVファイルを読み込んで、ヘッダ行を解析して任意の構造体に変換するプログラム

## Usage

```
go run cmd/main.go 
```

### Output

```golang
ヘッダ行: []string{
    "FullName",
    "NickName",
    "Description"
}
    ヘッダ行要素数: 3
     ---
    []interface {}{
        struct { FullName string; NickName string; Description string } { FullName: "安倍晋三", NickName: "漆黒", Description: "とんでもない" }, 
        struct { FullName string; NickName string; Description string }{FullName: "竹中平蔵", NickName: "売国奴", Description: "とんでもない"
    }
```

