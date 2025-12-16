package model

type Employee struct {
	Name        string `json:"name"`
	Designation string `json:"designation"`
	Age         int32  `json:"age"`
}
