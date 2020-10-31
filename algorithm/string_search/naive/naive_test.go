package naive

import "testing"

func TestStringSearch(t *testing.T) {
	cases := []struct {
		SearchString string
		Pattern      string
		Index        int
	}{
		{"abcdefg", "cde", 2},
		{"cdcdecg", "cde", 2},
	}

	for _, c := range cases {
		if index := StringSearch(c.SearchString, c.Pattern); index != c.Index {
			t.Fatalf("%s中搜索%s失败  %d!=%d", c.SearchString, c.Pattern, index, c.Index)
		}
	}
}

func BenchmarkStringSearch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringSearch("abcdefg",  "cde")
	}
}
