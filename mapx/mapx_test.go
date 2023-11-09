package mapx

import (
	"testing"
)

var (
	intCases = []struct {
		caseMap     map[int]int
		wantKeys    []int
		wantValues  []int
		wantHasKeys bool
	}{
		{
			caseMap: map[int]int{
				1: 4,
				2: 3,
				3: 2,
				4: 1,
			},
			wantKeys:    []int{1, 2, 3, 4},
			wantValues:  []int{4, 3, 2, 1},
			wantHasKeys: true,
		},
	}
)

func TestHasKeys(t *testing.T) {
	for _, intCase := range intCases {
		if got := HasKeys(intCase.caseMap, intCase.wantKeys...) == intCase.wantHasKeys; !got {
			t.Errorf("HasKeys want %v, but got %v", intCase.wantHasKeys, got)
		}
	}
}
