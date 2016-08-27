package gobind

import (
	"testing"
)

type Result struct {
	Foo  string `json:"foo" tag:"0"`
	Foo2 string `json:"foo2" tag:"1"`
	Foo3 string `json:"foo3" tag:"2"`
	Foo4 string `json:"foo4" tag:"3"`
	Foo5 string `json:"foo5" tag:"4"`
	Foo6 string `json:"foo6" tag:"5"`
}

type Data struct {
	Foo  string `json:"foo"`
	Foo2 string `json:"foo2"`
	Foo3 string `json:"foo3"`
	Foo4 string `json:"foo4"`
	Foo5 string `json:"foo5"`
	Foo6 string `json:"foo6"`
}

func TestBind(t *testing.T) {
	var r Result
	data := Data{"foo", "foo2", "foo3", "foo4", "foo5", "foo6"}

	err := Bind(&r, data, data, data, data, data, data)
	if err != nil {
		t.Fatal("bind fail, %s", err)
		return
	}

	if r.Foo != "foo" {
		t.Fatal("bind fail, r.Foo is not equal with foo, got:%s", r.Foo)
	}
	if r.Foo2 != "foo2" {
		t.Fatal("bind fail, r.Foo2 is not equal with foo2, got:%s", r.Foo2)
	}
	if r.Foo3 != "foo3" {
		t.Fatal("bind fail, r.Foo2 is not equal with foo2, got:%s", r.Foo2)
	}
	if r.Foo4 != "foo4" {
		t.Fatal("bind fail, r.Foo2 is not equal with foo2, got:%s", r.Foo2)
	}
	if r.Foo5 != "foo5" {
		t.Fatal("bind fail, r.Foo2 is not equal with foo2, got:%s", r.Foo2)
	}

}

func BenchmarkBind(b *testing.B) {
	type Result struct {
		Foo  string `json:"foo" tag:"0"`
		Foo2 string `json:"foo2" tag:"1"`
	}

	type Data struct {
		Foo  string `json:"foo"`
		Foo2 string `json:"foo2"`
	}

	data := Data{"foo", "foo2"}

	for i := 0; i < b.N; i++ {
		var r Result
		err := Bind(&r, data, data)
		if err != nil {
			b.Fatal("bind error")
			return
		}
	}
}

func BenchmarkDirectAssign(b *testing.B) {
	data := Data{"foo", "foo2", "foo3", "foo4", "foo5", "foo6"}
	for i := 0; i < b.N; i++ {
		var r Result
		r.Foo = data.Foo
		r.Foo2 = data.Foo2
		r.Foo3 = data.Foo3
		r.Foo4 = data.Foo4
		r.Foo5 = data.Foo5
		r.Foo6 = data.Foo6
	}
}
