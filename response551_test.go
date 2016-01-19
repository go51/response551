package response551_test

import (
	"fmt"
	"github.com/go51/response551"
	"net/http"
	"testing"
)

func TestRedirect(t *testing.T) {
	uri := "https://golang.org/"

	ret300 := response551.Redirect(uri, 300)
	ret301 := response551.Redirect(uri, 301)
	ret302 := response551.Redirect(uri, 302)
	ret303 := response551.Redirect(uri, 303)

	if ret300.Code() != 302 {
		t.Error("リダイレクトレスポンス構造体の生成に失敗しました。")
	}
	if ret300.Text() != http.StatusText(302) {
		t.Error("リダイレクトレスポンス構造体の生成に失敗しました。")
	}
	if ret300.Uri() != uri {
		t.Error("リダイレクトレスポンス構造体の生成に失敗しました。")
	}

	if ret301.Code() != 301 {
		t.Error("リダイレクトレスポンス構造体の生成に失敗しました。")
	}
	if ret301.Text() != http.StatusText(301) {
		t.Error("リダイレクトレスポンス構造体の生成に失敗しました。")
	}
	if ret301.Uri() != uri {
		t.Error("リダイレクトレスポンス構造体の生成に失敗しました。")
	}

	if ret302.Code() != 302 {
		t.Error("リダイレクトレスポンス構造体の生成に失敗しました。")
	}
	if ret302.Text() != http.StatusText(302) {
		t.Error("リダイレクトレスポンス構造体の生成に失敗しました。")
	}
	if ret302.Uri() != uri {
		t.Error("リダイレクトレスポンス構造体の生成に失敗しました。")
	}

	if ret303.Code() != 302 {
		t.Error("リダイレクトレスポンス構造体の生成に失敗しました。")
	}
	if ret303.Text() != http.StatusText(302) {
		t.Error("リダイレクトレスポンス構造体の生成に失敗しました。")
	}
	if ret303.Uri() != uri {
		t.Error("リダイレクトレスポンス構造体の生成に失敗しました。")
	}
}

func TestError(t *testing.T) {
	ret403 := response551.Error(403, "アクセスするには会員登録をおこないログインする必要があります。")
	ret404 := response551.Error(404, "指定されたカテゴリーは存在しません。")

	if ret403.Code() != 403 {
		t.Error("エラーレスポンス構造体の生成に失敗しました。")
	}
	if ret403.Text() != http.StatusText(403) {
		t.Error("エラーレスポンス構造体の生成に失敗しました。")
	}
	if ret403.Message() != "アクセスするには会員登録をおこないログインする必要があります。" {
		t.Error("エラーレスポンス構造体の生成に失敗しました。")
	}
	if ret403.String() != "アクセスするには会員登録をおこないログインする必要があります。" {
		t.Error("エラーレスポンス構造体の生成に失敗しました。")
	}

	if ret404.Code() != 404 {
		t.Error("エラーレスポンス構造体の生成に失敗しました。")
	}
	if ret404.Text() != http.StatusText(404) {
		t.Error("エラーレスポンス構造体の生成に失敗しました。")
	}
	if ret404.Message() != "指定されたカテゴリーは存在しません。" {
		t.Error("エラーレスポンス構造体の生成に失敗しました。")
	}
	if ret404.String() != "指定されたカテゴリーは存在しません。" {
		t.Error("エラーレスポンス構造体の生成に失敗しました。")
	}

	if fmt.Sprint(ret404) != "指定されたカテゴリーは存在しません。" {
		t.Error("エラーレスポンス構造体の生成に失敗しました。")
	}
}
