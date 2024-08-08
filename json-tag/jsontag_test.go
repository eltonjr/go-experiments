package jsontag

import (
	"encoding/json"
	"testing"
)

var rawjson = []byte(`{
	"tag": "tag",
	"tag2": "tag2",
	"tag3": "tag3",
	"tag4": "tag4",
	"tag5": "tag5",
	"tag6": "tag6",
	"tag7": "tag7",
	"tag8": "tag8",
	"tag9": "tag9",
	"tag10": "tag10"
}`)

func BenchmarkJsonTag(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var w WithTag
		json.Unmarshal(rawjson, &w)
	}
}

func BenchmarkJsonNoTag(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var w WithoutTag
		json.Unmarshal(rawjson, &w)
	}
}
