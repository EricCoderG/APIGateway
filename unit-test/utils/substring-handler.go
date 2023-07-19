package utils

import (
	substring "api-gateway/substring-service/kitex_gen/substring"
	"context"
	"strings"
)

// SubstringServiceImpl implements the last service interface defined in the IDL.
type SubstringServiceImpl struct{}

// FindSubstring implements the SubstringServiceImpl interface.
func (s *SubstringServiceImpl) FindSubstring(ctx context.Context, req *substring.SubstringRequest) (resp *substring.SubstringResponse, err error) {
	mainStr := req.MainString
	subStr := req.SubString

	resp = &substring.SubstringResponse{
		Positions: findPositions(mainStr, subStr),
	}
	return
}

// findPositions 返回子串在主串中出现的所有位置。
func findPositions(mainStr, subStr string) []int32 {
	var positions []int32
	for i := 0; i < len(mainStr); i++ {
		if strings.HasPrefix(mainStr[i:], subStr) {
			positions = append(positions, int32(i))
		}
	}
	return positions
}
