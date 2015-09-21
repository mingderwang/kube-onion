
//go:generate ginger $GOFILE
package main

//@ginger
type Kube struct {
	Ginger_Created int32 `json:"ginger_created"`
	Ginger_Id int32 `json:"ginger_id" gorm:"primary_key"`

	Gistid string `json:"gistid"`
	Filename string `json:"filename"`
	Description string `json:"description"`
	Tag string `json:"tag"`
	Like float64 `json:"like"`
}