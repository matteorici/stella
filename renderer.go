package main

import "fmt"

func PrintCities(

	cities []City,

) {

	fmt.Println(

		"Selected Culture\n",

	)

	fmt.Println(

		SelectedCulture,

	)

	fmt.Println()

	fmt.Println(

		"Generated Cities\n",

	)

	for _, city := range cities {

		fmt.Println(

			city.Name,

		)

	}
}
