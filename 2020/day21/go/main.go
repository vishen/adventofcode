package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {

	data, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		log.Fatalf("unable to open file: %v", err)
	}

	run(data)
}

type allergen struct {
	name                string
	count               int
	possibleIngredients map[string]int
}

func newAllergen() *allergen {
	return &allergen{possibleIngredients: make(map[string]int)}
}

func run(data []byte) {

	allergens := map[string]*allergen{}
	ingredientsAppear := map[string]int{}

	for _, line := range bytes.Split(data, []byte{'\n'}) {
		if len(line) == 0 {
			continue
		}
		lineSplit := bytes.Split(line, []byte{'('})

		ingredients := []string{}
		for _, ingredient := range bytes.Split(lineSplit[0], []byte{' '}) {
			if len(ingredient) == 0 {
				continue
			}
			ingredientsAppear[string(ingredient)]++
			ingredients = append(ingredients, string(ingredient))
		}

		allergies := []string{}
		for _, allergy := range bytes.Split(lineSplit[1], []byte{' '})[1:] {
			allergies = append(allergies, string(allergy[:len(allergy)-1]))
		}

		for _, allergy := range allergies {
			a, ok := allergens[allergy]
			if !ok {
				a = newAllergen()
				allergens[allergy] = a
			}
			a.count++
			for _, ingredient := range ingredients {
				a.possibleIngredients[ingredient]++
			}
		}
	}

	likelyUsed := map[string]int{}
	likelyUnused := map[string]int{}
	for _, a := range allergens {
		for i, count := range a.possibleIngredients {
			if a.count == count {
				likelyUsed[i]++
			} else {
				likelyUnused[i]++
			}
		}
	}
	//fmt.Println(likelyUsed)
	//fmt.Println(likelyUnused)

	total := 0
	for possiblyUnused, _ := range likelyUnused {
		if _, ok := likelyUsed[possiblyUnused]; ok {
			continue
		}
		count := ingredientsAppear[possiblyUnused]
		//fmt.Println(possiblyUnused, count)
		total += count
	}
	fmt.Println("Part 1: ", total)
}
