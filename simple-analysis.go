// package main

// import "fmt"

//	func main() {
//		// Print to the console
//		fmt.Println("Welcome to KDnuggets")
//	}
package main

import (
	"fmt"
	"os"
	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
)

// xklckms

func main() {
	// Loading the CSV file
	f, err := os.Open("adult.csv")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	df := dataframe.ReadCSV(f)

	// Filter the data: individuals with education level "HS-grad"
	hsGrad := df.Filter(dataframe.F{Colname: "education", Comparator: series.Eq, Comparando: "HS-grad"})
	fmt.Println("\nFiltered DataFrame (HS-grad):")
	fmt.Println(hsGrad)

	// calculating the average age of individuals in the dataset
	avgAge := df.Col("age").Mean()
	fmt.Printf("\nAverage age: %.2f\n", avgAge)

	// Describing the data
	fmt.Println("\nGenerate descriptive statistics:")
	description := df.Describe()
	fmt.Println(description)

}

// Before running the code, we have to install all the packages used in the above code. For that, we will run:

// $ go mod tidy
