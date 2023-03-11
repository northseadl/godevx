package setx

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	set := NewHashSet[string]("aa")
	set.Add()
	fmt.Println(set.Slice())
}
