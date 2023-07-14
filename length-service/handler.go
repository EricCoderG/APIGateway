package main

import (
	length "api-gateway/length-service/kitex_gen/length"
	"context"
)

// LengthServiceImpl implements the last service interface defined in the IDL.
type LengthServiceImpl struct{}

// CalculateLength implements the LengthServiceImpl interface.
func (s *LengthServiceImpl) CalculateLength(ctx context.Context, req *length.LengthRequest) (resp *length.LengthResponse, err error) {
	resp = &length.LengthResponse{
		Length: int32(len(req.InputString)),
	}
	return
}
