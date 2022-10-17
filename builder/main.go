package main

import (
	"design_patterns/builder/builders"
	"log"
)

func main() {
	human := builders.NewHuman().BuildEyeColor("green").BuildHeight(175).BuildWeight(78).BuildName("akif")

	log.Printf("%+v", human)
}
