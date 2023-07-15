package test

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"runtime"
	"testing"
)

func BenchmarkReverse(b *testing.B) {
	url := "http://127.0.0.1:8088/reverse"
	jsonStr := `{"inputString":"123456"}`
	for i := 0; i < b.N; i++ {
		resp, err := http.Post(url, "application/json", bytes.NewBuffer([]byte(jsonStr)))
		if err != nil {
			b.Fatal(err)
		}
		_, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			b.Fatal(err)
		}
		err = resp.Body.Close()
		if err != nil {
			return
		}
	}
}

func BenchmarkReverseParallel(b *testing.B) {
	url := "http://127.0.0.1:8088/reverse"
	jsonStr := `{"inputString":"123456"}`
	runtime.GOMAXPROCS(8)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			resp, err := http.Post(url, "application/json", bytes.NewBuffer([]byte(jsonStr)))
			if err != nil {
				b.Fatal(err)
			}
			_, err = ioutil.ReadAll(resp.Body)
			if err != nil {
				b.Fatal(err)
			}
			err = resp.Body.Close()
			if err != nil {
				return
			}
		}
	})
}
