import pandas as pd
N = 100 
with open('housesOutputPy.txt', 'wt') as outfile:
    for i in range(N):
        houses = pd.read_csv("housesInput.csv")
        outfile.write(houses.describe().to_string(header=True, index=True))
        outfile.write("\n")




