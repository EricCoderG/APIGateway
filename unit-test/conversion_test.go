package test

import (
	conversion "api-gateway/conversion-service/kitex_gen/conversion"
	convHandler "api-gateway/unit-test/utils"
	"context"
	"testing"
)

func Test_convertCase_func_normal(t *testing.T) {
	testString := "abcd"
	expectedResp := "ABCD"
	csi := new(convHandler.ConversionServiceImpl)
	req := &conversion.ConversionRequest{
		InputString: testString,
	}
	resp, err := csi.ConvertCase(context.Background(), req)
	if err != nil {
		t.Error(err)
	}
	if resp.ConvertedString != expectedResp {
		t.Errorf("expected resp is %s but resp in fact is %s", expectedResp, resp.ConvertedString)
	}
}

func Test_convertCase_func_Mixed(t *testing.T) {
	testString := "AbCd"
	expectedResp := "ABCD"
	csi := new(convHandler.ConversionServiceImpl)
	req := &conversion.ConversionRequest{
		InputString: testString,
	}
	resp, err := csi.ConvertCase(context.Background(), req)
	if err != nil {
		t.Error(err)
	}
	if resp.ConvertedString != expectedResp {
		t.Errorf("expected resp is %s but resp in fact is %s", expectedResp, resp.ConvertedString)
	}
}

func Test_convertCase_func_zero(t *testing.T) {
	testString := ""
	expectedResp := ""
	csi := new(convHandler.ConversionServiceImpl)
	req := &conversion.ConversionRequest{
		InputString: testString,
	}
	resp, err := csi.ConvertCase(context.Background(), req)
	if err != nil {
		t.Error(err)
	}
	if resp.ConvertedString != expectedResp {
		t.Errorf("expected resp is %s but resp in fact is %s", expectedResp, resp.ConvertedString)
	}
}
