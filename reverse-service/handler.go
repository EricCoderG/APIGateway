package main

import (
	reverse "api-gateway/reverse-service/kitex_gen/reverse"
	"context"
)

// ReverseServiceImpl implements the last service interface defined in the IDL.
type ReverseServiceImpl struct{}

// ReverseString implements the ReverseServiceImpl interface.
func (s *ReverseServiceImpl) ReverseString(ctx context.Context, req *reverse.ReverseRequest) (resp *reverse.ReverseResponse, err error) {
	// TODO: Your code here...
	return
}
