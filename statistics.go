package main

import "fmt"

func PrintStats(

	count int,

) {

	fmt.Println()

	fmt.Println(

		"Statistics\n",

	)

	fmt.Printf(

		"Generated: %d\n",

		count,

	)

	fmt.Printf(

		"Culture: %s\n",

		SelectedCulture,

	)
}
