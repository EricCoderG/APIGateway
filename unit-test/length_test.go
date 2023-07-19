package test

import (
	length "api-gateway/length-service/kitex_gen/length"
	lenHandler "api-gateway/unit-test/utils"
	"context"
	"testing"
)

func Test_calculateLength_func_normal(t *testing.T) {
	testString := "abcd"
	expectedResp := 4
	lsi := new(lenHandler.LengthServiceImpl)
	req := &length.LengthRequest{
		InputString: testString,
	}
	resp, err := lsi.CalculateLength(context.Background(), req)
	if err != nil {
		t.Error(err)
	}
	if resp.Length != int32(expectedResp) {
		t.Errorf("expected resp is %d but resp in fact is %d", expectedResp, resp.Length)
	}

}

func Test_calculateLength_func_zero(t *testing.T) {
	testString := ""
	expectedResp := 0
	lsi := new(lenHandler.LengthServiceImpl)
	req := &length.LengthRequest{
		InputString: testString,
	}
	resp, err := lsi.CalculateLength(context.Background(), req)
	if err != nil {
		t.Error(err)
	}
	if resp.Length != int32(expectedResp) {
		t.Errorf("expected resp is %d but resp in fact is %d", expectedResp, resp.Length)
	}

}
