package builders

import (
	"design_patterns/builder/models"
)

type MyModel models.Human

func NewHuman() *MyModel {
	return &MyModel{}
}

func (human *MyModel) BuildName(name string) *MyModel {
	human.Name = name

	return human
}

func (human *MyModel) BuildEyeColor(eyeColor string) *MyModel {
	human.EyeColor = eyeColor
	return human
}

func (human *MyModel) BuildHeight(height int) *MyModel {
	human.Height = height
	return human
}

func (human *MyModel) BuildWeight(weight int) *MyModel {
	human.Weight = weight
	return human
}
