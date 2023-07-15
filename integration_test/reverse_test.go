package test

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"
)

func Test_reverse_normal(t *testing.T) {
	url := "http://127.0.0.1:8088/reverse"
	jsonStr := `{"inputString":"abcccc"}`
	expectedResp := `{"reversedString":"ccccba"}`
	resp, err := http.Post(url, "application/json", bytes.NewBuffer([]byte(jsonStr)))
	if err != nil {
		t.Fatal(err)
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	err = resp.Body.Close()

	if err != nil {
		return
	}

	respString := string(respBytes)

	if respString != expectedResp {
		t.Error("expected resp is " + expectedResp + " but real resp is " + respString + "!")
	} else {
		t.Log("Normal Test Pass!")
	}
}

func Test_zero_normal(t *testing.T) {
	url := "http://127.0.0.1:8088/reverse"
	jsonStr := `{"inputString":""}`
	expectedResp := `{"reversedString":""}`
	resp, err := http.Post(url, "application/json", bytes.NewBuffer([]byte(jsonStr)))
	if err != nil {
		t.Fatal(err)
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	err = resp.Body.Close()

	if err != nil {
		return
	}

	respString := string(respBytes)

	if respString != expectedResp {
		t.Error("expected resp is " + expectedResp + " but real resp is " + respString + "!")
	} else {
		t.Log("Zero Test Pass!")
	}
}
