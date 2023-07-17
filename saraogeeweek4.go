package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/montanaflynn/stats"
)

func main() {
	//Runs 100 times
	for j := 0; j < 1; j++ {
		//Reads in each line of input as an array of strings
		file, _ := os.Open("housesInput.csv")
		reader := csv.NewReader(file)
		records, _ := reader.ReadAll()
		defer file.Close() //It is idiomatic to close files after opening

		//Initializing of variable arrays and percentile values
		var values []float64
		var incomes []float64
		var ages []float64
		var rooms []float64
		var bedrooms []float64
		var pops []float64
		var hhs []float64
		var percentiles = []float64{25, 50, 75} //same values as for R, Python

		for i := range records {
			//Converts records to float64 from strings
			value, _ := strconv.ParseFloat(records[i][0], 64)
			income, _ := strconv.ParseFloat(records[i][1], 64)
			age, _ := strconv.ParseFloat(records[i][2], 64)
			room, _ := strconv.ParseFloat(records[i][3], 64)
			bedroom, _ := strconv.ParseFloat(records[i][4], 64)
			pop, _ := strconv.ParseFloat(records[i][5], 64)
			hh, _ := strconv.ParseFloat(records[i][6], 64)

			//Appends value from record to initialized arrays
			values = append(values, value)
			incomes = append(incomes, income)
			ages = append(ages, age)
			rooms = append(rooms, room)
			bedrooms = append(bedrooms, bedroom)
			pops = append(pops, pop)
			hhs = append(hhs, hh)
		}

		//Not sure if better way to deal with zero that is initialized, but this code removes
		//the zero that is in these values when they are specified as a float64 array in line 50
		values = values[1:]
		incomes = incomes[1:]
		ages = ages[1:]
		rooms = rooms[1:]
		bedrooms = bedrooms[1:]
		pops = pops[1:]
		hhs = hhs[1:]

		//Prints descriptions in a txt file with stats.Describe function
		f, _ := os.Create("housesOutputGo.txt")

		values_descrip := GoDescribe(values, false, &percentiles)
		f.WriteString("values")
		f.WriteString("\n")
		f.WriteString(Stringed(values_descrip))
		f.WriteString("\n")

		incomes_descrip := GoDescribe(incomes, false, &percentiles)
		f.WriteString("income")
		f.WriteString("\n")
		f.WriteString(Stringed(incomes_descrip))
		f.WriteString("\n")

		age_descrip := GoDescribe(ages, false, &percentiles)
		f.WriteString("age")
		f.WriteString("\n")
		f.WriteString(Stringed(age_descrip))
		f.WriteString("\n")

		rooms_descrip := GoDescribe(rooms, false, &percentiles)
		f.WriteString("rooms")
		f.WriteString("\n")
		f.WriteString(Stringed(rooms_descrip))
		f.WriteString("\n")

		bedrooms_descrip := GoDescribe(bedrooms, false, &percentiles)
		f.WriteString("bedrooms")
		f.WriteString("\n")
		f.WriteString(Stringed(bedrooms_descrip))
		f.WriteString("\n")

		pop_descrip := GoDescribe(pops, false, &percentiles)
		f.WriteString("pop")
		f.WriteString("\n")
		f.WriteString(Stringed(pop_descrip))
		f.WriteString("\n")

		hh_descrip := GoDescribe(hhs, false, &percentiles)
		f.WriteString("hh")
		f.WriteString("\n")
		f.WriteString(Stringed(hh_descrip))
		f.WriteString("\n")
	}
}

// Other imported code from stats package
type Description struct { //type imported from stats package for modified Stringed function
	Count                  int
	Mean                   float64
	Std                    float64
	Max                    float64
	Min                    float64
	DescriptionPercentiles []descriptionPercentile
	AllowedNaN             bool
}
type descriptionPercentile struct { ////type imported from stats package for modified Stringed function
	Percentile float64
	Value      float64
}

func Stringed(d *stats.Description) string { //modified to remove NaN fields from stats package describe.go
	var str string
	str += fmt.Sprintf("count\t%d\n", d.Count)
	str += fmt.Sprintf("mean\t%.*f\n", 2, d.Mean)
	str += fmt.Sprintf("std\t%.*f\n", 2, d.Std)
	str += fmt.Sprintf("max\t%.*f\n", 2, d.Max)
	str += fmt.Sprintf("min\t%.*f\n", 2, d.Min)
	for _, percentile := range d.DescriptionPercentiles {
		str += fmt.Sprintf("%.2f%%\t%.*f\n", percentile.Percentile, 2, percentile.Value)
	}
	return str
}

// Function created for testing framework to ensure consistent results with Python/R
func GoDescribe(a stats.Float64Data, b bool, c *[]float64) (d *stats.Description) {
	f, _ := stats.Describe(a, b, c)
	return f
}
