import csv
import matplotlib.pyplot as plt
import numpy as np

csvFile = "teste1.csv"

with open(csvFile) as file:
    data = list(csv.reader(file))
    data1 = [int(x) for x in data[0]]

csvFile = "teste5.csv"

with open(csvFile) as file:
    data = list(csv.reader(file))
    data2 = [int(x) for x in data[0]]

csvFile = "teste10.csv"

with open(csvFile) as file:
    data = list(csv.reader(file))
    data3 = [int(x) for x in data[0]]

csvFile = "teste20.csv"

with open(csvFile) as file:
    data = list(csv.reader(file))
    data4 = [int(x) for x in data[0]]

csvFile = "teste40.csv"

with open(csvFile) as file:
    data = list(csv.reader(file))
    data5 = [int(x) for x in data[0]]

csvFile = "teste80.csv"

with open(csvFile) as file:
    data = list(csv.reader(file))
    data6 = [int(x) for x in data[0]]

data = [data1, data2, data3, data4, data5, data6]

plt.boxplot(data, showfliers=False, labels=[1, 5, 10, 20, 40, 80])

plt.show()
