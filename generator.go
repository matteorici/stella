package main

type Generator struct{}

func (g Generator) Create(culture string) City {

	return City{

		Name: pick(Prefixes[culture]) +

			pick(Suffixes[culture]),

	}
}
