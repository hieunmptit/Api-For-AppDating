package interfaces

import "github.com/jinzhu/gorm"

type Profile struct {
	gorm.Model
	Name     string
	Gender   string
	Age      int
	Birthday string
	Location string
	PLo      int
	PGender  string
	PAge     int
	Status   []int
}
