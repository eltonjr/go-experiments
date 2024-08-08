package jsontag

type WithTagSmall struct {
	Tag  string `json:"tag"`
	Tag2 string `json:"tag2"`
	Tag3 string `json:"tag3"`
	Tag4 string `json:"tag4"`
	Tag5 string `json:"tag5"`
}

type WithoutTagSmall struct {
	Tag  string
	Tag2 string
	Tag3 string
	Tag4 string
	Tag5 string
}

var rawjsonSmall = []byte(`{
	"tag": "tag",
	"tag2": "tag2",
	"tag3": "tag3",
	"tag4": "tag4",
	"tag5": "tag5"
}`)
