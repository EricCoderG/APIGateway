package test

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"
)

func Test_conversion_normal(t *testing.T) {
	url := "http://127.0.0.1:8088/convert"
	jsonStr := `{"inputString":"abcccc"}`
	expectedResp := `{"convertedString":"ABCCCC"}`
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

func Test_conversion_ZERO(t *testing.T) {
	url := "http://127.0.0.1:8088/convert"
	jsonStr := `{"inputString":""}`
	expectedResp := `{"convertedString":""}`
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
		t.Log("ZERO Test Pass!")
	}
}

func Test_conversion_Mixed(t *testing.T) {
	url := "http://127.0.0.1:8088/convert"
	jsonStr := `{"inputString":"aBcCcC"}`
	expectedResp := `{"convertedString":"ABCCCC"}`
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
		t.Log("Mixed Test Pass!")
	}
}
