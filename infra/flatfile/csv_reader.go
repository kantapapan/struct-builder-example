package flatfile

import (
	"encoding/csv"
	"os"

	"github.com/kantapapan/struct-builder-example/configs"
	"github.com/kantapapan/struct-builder-example/ifaces/flatfile"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

// CsvReader はファイル読取実装です。
type CsvReader struct {
	cfg configs.Config
}

// NewCsvReader はファイルリーダーを生成します。
func NewCsvReader(cfg configs.Config) flatfile.Reader {
	return &CsvReader{cfg: cfg}
}

// ReadLines は指定行数読み込んだデータスライスを返します。
func (f *CsvReader) ReadLines(path string) (lines [][]string, err error) {

	fr, err := os.Open(path)
	if err != nil {
		return
	}
	defer fr.Close()
	r := csv.NewReader(transform.NewReader(fr, japanese.ShiftJIS.NewDecoder()))
	r.Comma = ','
	//r.Comment = '#'
	records, err := r.ReadAll()

	if err != nil {
		return
	}

	lines = [][]string{}
	for _, r := range records {
		lines = append(lines, r)
	}

	return lines, nil
}
