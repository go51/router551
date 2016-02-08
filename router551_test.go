package router551_test

import (
	"github.com/go51/container551"
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

func ActionFunc(c *container551.Container) interface{} {
	ret := []string{"Action", "Function", "Test", "Interface"}

	return ret
}

func TestAdd(t *testing.T) {
	r := router551.Load()
	r.Add(router551.GET, "index", "/", ActionFunc)
	r.Add(router551.POST, "top", "/top", ActionFunc)
	r.Add(router551.PUT, "account", "/account/:account_id:", ActionFunc)
	r.Add(router551.DELETE, "account_action", "/account/:account_id:/:action:", ActionFunc)
	r.Add(router551.COMMAND, "command", "command:test", ActionFunc)
	r.Add(router551.GET|router551.POST|router551.PUT|router551.DELETE, "all", "/all/:action:/detail/:no:", ActionFunc)

	route := r.FindRouteByName("GET", "index")
	if route == nil {
		t.Errorf("ルーティングの設定が失敗しました。")
	}
	if route.Name() != "index" {
		t.Errorf("ルーティングの設定が失敗しました。")
	}
	if len(route.Keys()) != 0 {
		t.Errorf("ルーティングの設定が失敗しました。")
	}

	route = r.FindRouteByName("POST", "top")
	if route == nil {
		t.Errorf("ルーティングの設定が失敗しました。")
	}
	if route.Name() != "top" {
		t.Errorf("ルーティングの設定が失敗しました。")
	}
	if len(route.Keys()) != 0 {
		t.Errorf("ルーティングの設定が失敗しました。")
	}

	route = r.FindRouteByName("PUT", "account")
	if route == nil {
		t.Errorf("ルーティングの設定が失敗しました。")
	}
	if route.Name() != "account" {
		t.Errorf("ルーティングの設定が失敗しました。")
	}
	if len(route.Keys()) != 1 {
		t.Errorf("ルーティングの設定が失敗しました。")
	}
	if route.Keys()[0] != "account_id" {
		t.Errorf("ルーティングの設定が失敗しました。")
	}

	route = r.FindRouteByName("DELETE", "account_action")
	if route == nil {
		t.Errorf("ルーティングの設定が失敗しました。")
	}
	if route.Name() != "account_action" {
		t.Errorf("ルーティングの設定が失敗しました。")
	}
	if len(route.Keys()) != 2 {
		t.Errorf("ルーティングの設定が失敗しました。")
	}
	if route.Keys()[0] != "account_id" {
		t.Errorf("ルーティングの設定が失敗しました。")
	}
	if route.Keys()[1] != "action" {
		t.Errorf("ルーティングの設定が失敗しました。")
	}

	route = r.FindRouteByName("COMMAND", "command")
	if route == nil {
		t.Errorf("ルーティングの設定が失敗しました。")
	}
	if route.Name() != "command" {
		t.Errorf("ルーティングの設定が失敗しました。")
	}
	if len(route.Keys()) != 0 {
		t.Errorf("ルーティングの設定が失敗しました。")
	}

	getRoute := r.FindRouteByName("GET", "all")
	postRoute := r.FindRouteByName("POST", "all")
	putRoute := r.FindRouteByName("PUT", "all")
	deleteRoute := r.FindRouteByName("DELETE", "all")

	if getRoute != postRoute {
		t.Errorf("ルーティングの設定が失敗しました。")
	}
	if postRoute != putRoute {
		t.Errorf("ルーティングの設定が失敗しました。")
	}
	if putRoute != deleteRoute {
		t.Errorf("ルーティングの設定が失敗しました。")
	}
	if deleteRoute != getRoute {
		t.Errorf("ルーティングの設定が失敗しました。")
	}

}

func TestFindActionByPathMatch(t *testing.T) {
	r := router551.Load()
	r.Add(router551.GET, "index", "/", ActionFunc)
	r.Add(router551.POST, "top", "/top", ActionFunc)
	r.Add(router551.PUT, "account", "/account/:account_id:", ActionFunc)
	r.Add(router551.DELETE, "account_action", "/account/:account_id:/:action:", ActionFunc)
	r.Add(router551.COMMAND, "command", "command:test", ActionFunc)
	r.Add(router551.GET|router551.POST|router551.PUT|router551.DELETE, "all", "/all/:action:/detail/:no:", ActionFunc)

	route := r.FindRouteByPathMatch("GET", "/")
	if route.Name() != "index" {
		t.Errorf("ルートの取得に失敗しました。")
	}
	if len(route.Keys()) != 0 {
		t.Errorf("ルートの取得に失敗しました。")
	}

	route = r.FindRouteByPathMatch("POST", "/top")
	if route.Name() != "top" {
		t.Errorf("ルートの取得に失敗しました。")
	}
	if len(route.Keys()) != 0 {
		t.Errorf("ルートの取得に失敗しました。")
	}

	route = r.FindRouteByPathMatch("PUT", "/account/13")
	if route.Name() != "account" {
		t.Errorf("ルートの取得に失敗しました。")
	}
	if len(route.Keys()) != 1 {
		t.Errorf("ルートの取得に失敗しました。")
	}

	route = r.FindRouteByPathMatch("DELETE", "/account/13/get")
	if route.Name() != "account_action" {
		t.Errorf("ルートの取得に失敗しました。")
	}
	if len(route.Keys()) != 2 {
		t.Errorf("ルートの取得に失敗しました。")
	}

	route = r.FindRouteByPathMatch("COMMAND", "command:test")
	if route.Name() != "command" {
		t.Errorf("ルートの取得に失敗しました。")
	}
	if len(route.Keys()) != 0 {
		t.Errorf("ルートの取得に失敗しました。")
	}

	route = r.FindRouteByPathMatch("GET", "/all/get/detail/1")
	if route.Name() != "all" {
		t.Errorf("ルートの取得に失敗しました。")
	}
	if len(route.Keys()) != 2 {
		t.Errorf("ルートの取得に失敗しました。")
	}

	url := r.Url("index")
	if url != "/" {
		t.Errorf("URL 生成に失敗しました。/ => %s", url)
	}
	url = r.Url("top")
	if url != "/top" {
		t.Errorf("URL 生成に失敗しました。/top => %s", url)
	}
	url = r.Url("account", "1")
	if url != "/account/1" {
		t.Errorf("URL 生成に失敗しました。/account/1 => %s", url)
	}
	url = r.Url("account_action", "1", "lock")
	if url != "/account/1/lock" {
		t.Errorf("URL 生成に失敗しました。/account/1/lock => %s", url)
	}
}
