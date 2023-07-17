N = 100 
sink("housesOutputR.txt")
for (i in 1:N) {
    houses = read.csv(file = "housesInput.csv", header = TRUE)
    print(summary(houses)) 
}
sink()
