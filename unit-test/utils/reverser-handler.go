package utils

import (
	"api-gateway/reverse-service/kitex_gen/reverse"
	"context"
)

// ReverseServiceImpl implements the last service interface defined in the IDL.
type ReverseServiceImpl struct{}

// ReverseString implements the ReverseServiceImpl interface.
func (s *ReverseServiceImpl) ReverseString(ctx context.Context, req *reverse.ReverseRequest) (resp *reverse.ReverseResponse, err error) {
	resp = &reverse.ReverseResponse{
		ReversedString: Reverse(req.InputString),
	}
	return
}

func Reverse(s string) string {
	a := []rune(s)
	// todo: 使用微服务
	lenA := len(a)
	for i, j := 0, lenA-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	return string(a)
}
