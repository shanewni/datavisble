package engine

import (
	"os"
	"testing"
)

//单元测试，需要更改测试函数文件路径 "../crack/" + a.ttfFontStyle + FontTail
func TestAuthor_CreateFontFile(t *testing.T) {
	a := &Author{ttfFontStyle:"lryMhtHe"}
	a.CreateFontFile()
	_, err := os.Open("../crack/lryMhtHe.ttf")
	if err != nil{
		t.Error("TestAuthor_CreateFontFile error:",err)
	}
}

//性能测试，需要更改测试函数文件路径 "../crack/" + a.ttfFontStyle + FontTail
func BenchmarkAuthor_CreateFontFile(b *testing.B) {
	for i:=0; i< b.N; i ++{
		a := &Author{ttfFontStyle:"lryMhtHe"}
		a.CreateFontFile()
	}
}

//单元测试，需要更改测试函数文件路径 "../crack/"
func TestAuthor_BuildFontFile(t *testing.T) {
	a := &Author{ttfFontStyle:"lryMhtHe"}
	a.BuildFontFile()
	_, err := os.Open("../crack/lryMhtHe.ttx")
	if err != nil{
		t.Error("TestAuthor_BuildFontFile error:",err)
	}
}

//性能测试，需要更改测试函数文件路径 "../crack/" + a.ttfFontStyle + FontTail
func BenchmarkAuthor_BuildFontFile(b *testing.B) {
	for i:=0; i< b.N; i ++{
		a := &Author{ttfFontStyle:"lryMhtHe"}
		a.BuildFontFile()
	}
}

//单元测试，需要更改测试函数文件路径 "../crack/"
func TestAuthor_CraCkXML(t *testing.T) {
	a := &Author{ttfFontStyle:"lryMhtHe"}
	mp := a.CraCkXML()
	if mp == nil{
		t.Error("TestAuthor_CraCkXML error:")
	}
}

//性能测试，需要更改测试函数文件路径 "../crack/" + a.ttfFontStyle + FontTail
func BenchmarkAuthor_CraCkXML(b *testing.B) {
	for i:=0; i< b.N; i ++{
		a := &Author{ttfFontStyle:"lryMhtHe"}
		a.CraCkXML()
	}
}