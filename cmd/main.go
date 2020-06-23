package main

import (
	"fmt"
	"reflect"

	"github.com/kantapapan/struct-builder-example/configs"
	fgw "github.com/kantapapan/struct-builder-example/ifaces/flatfile"
	fileInfra "github.com/kantapapan/struct-builder-example/infra/flatfile"
)

// 適当なCSVファイルを読み込んでヘッダ行を構造体の属性値にセット
// 行データを構造体に詰め込んでいく
func main() {

	c, _ := configs.Load()
	path := "/data/storage/efs-root/maintenance/master/example/sample.csv"

	var reader fgw.Reader
	reader = fileInfra.NewCsvReader(c)
	lines, err := reader.ReadLines(path)
	if err != nil {
		fmt.Printf("%#v\n", err)
	}

	b := NewStructBuilder()

	// 構造体フィールドを生成する
	for _, v := range lines[0] {
		b.AddField(v, reflect.TypeOf(""))
	}

	header := lines[0]
	fmt.Printf("ヘッダ行: %#v\n", header)
	fmt.Printf("ヘッダ行要素数: %d\n ", len(header))

	// NOTE: example : 型情報を分析 設定する値はなんでもよい
	//b.AddField("Name", reflect.TypeOf("")) // <- string
	//b.AddField("Age", reflect.TypeOf(123)) // <- int

	data := b.Build()
	p := data.NewInstance()

	//fmt.Printf("%#v\n", p)

	// 行データを構造体にセットする
	var res []interface{}
	for index, ln := range lines {
		if index == 0 {
			continue
		}
		p.SetString(b.field[0].Name, ln[0])
		p.SetString(b.field[1].Name, ln[1])
		p.SetString(b.field[2].Name, ln[2])
		res = append(res, p.Value())
	}

	// NOTE: example
	//fmt.Println(i.Value())   //  値渡し
	//fmt.Println(i.Pointer()) //  ポインタ渡し

	fmt.Println("---")
	fmt.Printf("%#v\n", res) // 構造体出力
}

// ======

// StructBuilder ...
type StructBuilder struct {
	field []reflect.StructField
}

// NewStructBuilder ...
func NewStructBuilder() *StructBuilder {
	return &StructBuilder{}
}

// AddField ...
func (b *StructBuilder) AddField(fname string, ftype reflect.Type) {
	b.field = append(
		b.field,
		reflect.StructField{
			Name: fname,
			Type: ftype,
		})
}

// Build ...
func (b *StructBuilder) Build() Struct {
	strct := reflect.StructOf(b.field)
	index := make(map[string]int)
	for i := 0; i < strct.NumField(); i++ {
		index[strct.Field(i).Name] = i
	}
	return Struct{strct, index}
}

// Struct ...
type Struct struct {
	strct reflect.Type
	index map[string]int
}

// NewInstance ...
func (s *Struct) NewInstance() *Instance {
	instance := reflect.New(s.strct).Elem()
	return &Instance{instance, s.index}
}

// Instance ...
type Instance struct {
	internal reflect.Value
	index    map[string]int
}

// Field ...
func (i *Instance) Field(name string) reflect.Value {
	return i.internal.Field(i.index[name])
}

// SetString ...
func (i *Instance) SetString(name, value string) {
	i.Field(name).SetString(value)
}

// SetBool ...
func (i *Instance) SetBool(name string, value bool) {
	i.Field(name).SetBool(value)
}

// SetInt ...
func (i *Instance) SetInt(name string, value int) {
	i.Field(name).SetInt(int64(value))
}

// SetFloat ...
func (i *Instance) SetFloat(name string, value float64) {
	i.Field(name).SetFloat(value)
}

// Value ...
func (i *Instance) Value() interface{} {
	return i.internal.Interface()
}

// Pointer ...
func (i *Instance) Pointer() interface{} {
	return i.internal.Addr().Interface()
}
