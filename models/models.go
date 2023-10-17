package models

type Student struct {
	Name  string  `json:"name"`
	Age   uint64  `json:"age"`
	Grade float32 `json:"grade"`
}
