package engine

import (
	"fmt"
	"testing"
)

//能否成功返回页数
func Test_getend(t *testing.T){
	r := &ReturnEnd{"https://www.qidian.com/all"}
	end,err := r.Start()
	if err != nil{
		t.Error("error: ",err)
	}
	//根据网址获取页数
	if end >= 57610{
		fmt.Println("Ok: ",end)
	}
}

//查看性能
func BenchmarkStart(b *testing.B) {
	r := &ReturnEnd{"https://www.qidian.com/all"}
	for i:= 0; i< b.N; i++{
		r.Start()
	}
}