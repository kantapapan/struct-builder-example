// Package flatfile はフラットファイルのインターフェースアダプターを提供します
package flatfile

// Reader 読み取りインターフェース
type Reader interface {
	// ReadLines はデータ行読み込みを行います。
	ReadLines(path string) (linedata [][]string, err error)
}
