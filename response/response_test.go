package response

import (
	"encoding/json"
	"testing"
)

func TestNormalizeArrayIfNeeded(t *testing.T) {
	var nilArr []int
	cases := []struct {
		a interface{}
		b string
	}{
		{[]int{1, 2}, `[1,2]`},
		{[]int{}, `[]`},
		{nilArr, `[]`},
		{nil, `null`},
		{"hello", `"hello"`},
		{map[string]interface{}{"hello": "hi!"}, `{"hello":"hi!"}`},
		{map[string][]string{"hello": {"world", "!"}}, `{"hello":["world","!"]}`},
		{map[string][]string{"hello": {}}, `{"hello":[]}`},
	}

	for _, oneCase := range cases {
		nVal := normalizeArrayIfNeeded(oneCase.a)
		res, _ := json.Marshal(nVal)
		if string(res) != oneCase.b {
			t.Errorf("a: %#v; expected: %#v; got: %#v", oneCase.a, oneCase.b, string(res))
		}
	}
}
