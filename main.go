package main

func main() {

	repository :=

		Repository{}

	cities :=

		repository.Generate()

	PrintCities(

		cities,

	)

	PrintStats(

		len(cities),

	)
}
