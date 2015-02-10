package gobg

import (
	"reflect"
	"testing"
)

func TestGobGeneralEncoding(t *testing.T) {
	type Test struct {
		Name string
		Age  int32
	}
	test := Test{"Joe", 32}

	encOutput, err := GobGeneralEncoder(test)
	if err != nil {
		t.Fatalf("err isn't nil", err)
	}

	var decOutput Test
	err = GobGeneralDecoder(encOutput, Test{}, &decOutput)
	if err != nil {
		t.Fatalf("err isn't nil", err)
	}
	if !reflect.DeepEqual(decOutput, test) {
		t.Error("decOutput is ", decOutput, " Expected ", test)
	}
}
