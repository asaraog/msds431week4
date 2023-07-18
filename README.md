# Command-Line Applications and Interface

## Project Summary

Go will be help power our backend web and database servers and distributed service offerings on the cloud. However, data science operations remain a key concern as Python/R remain popular. This project aims to implement basic summary statistics of the [California Housing Prices (Miller 2015)](./housesInput.csv) in Go using the [stats package](https://github.com/montanaflynn/stats). The Go implementation is benchmarked for runtime using 'time' before commands in the command line to compare with [Python's](./runHouses.py) pandas.describe() and [R's](./runHouses.R) summary functions, running each operation 100 times. The results are subsequently written to a .txt file. The filenames can be specified for the Go application on the command line thanks to the [flag package](https://pkg.go.dev/flag).

Python was significantly faster compared to R and Go implementations with 'real' runtimes of 2.27s, 4.10s, 5.02s for Python, R and Go respectively. While Python and R were less verbose and a bit faster than Go, Go's testing package ensured summary statistics as Python for each of the seven variables (value, income, age, rooms, bedrooms, pop, hh) during development. The Data Science team sees test-driven development as an asset in Go and with equivalent statistical results, our concerns about switching to Go are alleviated. We strongly recommend using Go as the primary programming language accross the company.

## Files

*saraogeeweek4.go:* \
Main routine loads input .csv file and computes statistics using [Describe](https://github.com/montanaflynn/stats/blob/master/describe.go) in the stats package. This is nested into a GoDescribe function is used for testing. Statistics for each variable is written to an output .txt file.

*saraogeeweek4_test.go:* \
Unit test for GoDescribe function. This testing routine ensures equivalence with Pyhon/R output coefficients of mean, maximum and minimum is rounded to 2 significant figures with a [roundFloat function](gosamples.dev/round-float/).

*Week4* \
Unix executable file of cross-compiled Go code for Mac/Windows. The -input and -output flag specifies the names of the input .csv and output .txt files.

## Installation

Download or git clone this project onto local machine into folder on local machine.

```
git clone https://github.com/asaraog/msds431week4.git
cd msds431week4
time ./Week4
time python3 runHouses.py 
time  Rscript runHouses.R

./Week4 -input housesInput.csv -output housesOutputGo.txt
```
It is benchmarked using the same command-line 'time' before running each programming language (Go, Python and R). The Go application can be run with user-specified filenames:

```
./Week4 -input housesInput.csv -output housesOutputGo.txt
```

## References

Flynn, Montana. 2023. “Stats - Golang Statistics Package.” 2023. https://github.com/montanaflynn/stats. \
Miller, Thomas. 2015. “Modeling Techniques in Predictive Analytics Chapter 10.” 2015. https://github.com/mtpa/mtpa/tree/master/MTPA_Chapter_10.

