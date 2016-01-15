package router551_test

import (
	"github.com/go51/router551"
	"testing"
)

func TestMethod(t *testing.T) {
	get := router551.GET
	post := router551.POST
	put := router551.PUT
	delete := router551.DELETE
	command := router551.COMMAND
	unknown := router551.UNKNOWN

	if get.String() != "GET" {
		t.Errorf("HTTP メソッドが正常に定義されていません。GET => %s", get.String())
	}
	if post.String() != "POST" {
		t.Errorf("HTTP メソッドが正常に定義されていません。POST => %s", post.String())
	}
	if put.String() != "PUT" {
		t.Errorf("HTTP メソッドが正常に定義されていません。PUT => %s", put.String())
	}
	if delete.String() != "DELETE" {
		t.Errorf("HTTP メソッドが正常に定義されていません。DELETE => %s", delete.String())
	}
	if command.String() != "COMMAND" {
		t.Errorf("HTTP メソッドが正常に定義されていません。COMMAND => %s", command.String())
	}
	if unknown.String() != "UNKNOWN" {
		t.Errorf("HTTP メソッドが正常に定義されていません。UNKNOWN => %s", unknown.String())
	}
}

func TestLoad(t *testing.T) {
	r1 := router551.Load()
	r2 := router551.Load()
	r3 := &router551.Router{}

	if r1 == nil {
		t.Error("インスタンスの生成に失敗しました。")
	}

	if r2 == nil {
		t.Error("インスタンスの生成に失敗しました。")
	}

	if r1 != r2 {
		t.Error("インスタンスの生成に失敗しました。")
	}

	if r1 == r3 {
		t.Error("インスタンスの生成に失敗しました。")
	}
}

func BenchmarkLoad(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = router551.Load()
	}
}
