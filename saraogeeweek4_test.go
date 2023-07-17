package main

import (
	"encoding/csv"
	"math"
	"os"
	"strconv"
	"testing"
)

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

func TestGoDescribe(t *testing.T) {
	// Reads in data into an array of floats as in main program
	file, _ := os.Open("housesInput.csv")
	reader := csv.NewReader(file)
	records, _ := reader.ReadAll() //in string format
	defer file.Close()             //It is idiomatic to close files after opening

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
	//Removes zero initialization value
	values = values[1:]
	incomes = incomes[1:]
	ages = ages[1:]
	rooms = rooms[1:]
	bedrooms = bedrooms[1:]
	pops = pops[1:]
	hhs = hhs[1:]

	// Calculates descriptions using stats package
	values_descrip := GoDescribe(values, true, &percentiles)
	incomes_descrip := GoDescribe(incomes, true, &percentiles)
	age_descrip := GoDescribe(ages, true, &percentiles)
	rooms_descrip := GoDescribe(rooms, true, &percentiles)
	bedrooms_descrip := GoDescribe(bedrooms, true, &percentiles)
	pop_descrip := GoDescribe(pops, true, &percentiles)
	hh_descrip := GoDescribe(hhs, true, &percentiles)

	//Checks values with Python values rounded to 2 decimal places from housesOutputPy.txt
	if values_descrip.Min != 14999.00 || values_descrip.Max != 500001.00 || roundFloat(values_descrip.Mean, 2) != 206855.82 {
		t.Error("Value calculation error")
	}

	if roundFloat(incomes_descrip.Min, 2) != 0.50 || roundFloat(incomes_descrip.Max, 2) != 15.00 || roundFloat(incomes_descrip.Mean, 2) != 3.87 {
		t.Error("Income calculation error")
	}

	if age_descrip.Min != 1.0 || age_descrip.Max != 52.0 || roundFloat(age_descrip.Mean, 2) != 28.64 {
		t.Error("Age calculation error")
	}

	if rooms_descrip.Min != 2.0 || rooms_descrip.Max != 39320.00 || roundFloat(rooms_descrip.Mean, 2) != 2635.76 {
		t.Error("Rooms calculation error")
	}

	if bedrooms_descrip.Min != 1.0 || bedrooms_descrip.Max != 6445.00 || roundFloat(bedrooms_descrip.Mean, 2) != 537.90 {
		t.Error("Bedrooms calculation error")
	}

	if pop_descrip.Min != 3.0 || pop_descrip.Max != 35682.00 || roundFloat(pop_descrip.Mean, 2) != 1425.48 {
		t.Error("pop calculation error")
	}

	if hh_descrip.Min != 1.0 || hh_descrip.Max != 6082.00 || roundFloat(hh_descrip.Mean, 2) != 499.54 {
		t.Error("hh calculation error")
	}

	if hh_descrip.Count != 20640 || values_descrip.Count != 20640 || pop_descrip.Count != 20640 || rooms_descrip.Count != 20640 || bedrooms_descrip.Count != 20640 || age_descrip.Count != 20640 || incomes_descrip.Count != 20640 {
		t.Error("Counting error")
	}
}
