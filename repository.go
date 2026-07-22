package main

type Repository struct{}

func (r Repository) Generate() []City {

	var cities []City

	gen := Generator{}

	for i := 0; i < CityCount; i++ {

		cities = append(

			cities,

			gen.Create(

				SelectedCulture,

			),

		)

	}

	return cities
}
