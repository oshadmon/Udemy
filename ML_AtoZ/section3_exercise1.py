# Importing the necessary libraries
import pandas 

# Loading the Iris dataset
dataset = pandas.read_csv("csv_files/iris.csv")


# Creating the matrix of features (X) and the dependent variable vector (y)
X = dataset.iloc[:, :-1].values
y = dataset.iloc[:, -1].values



# Printing the matrix of features and the dependent variable vector
print(X)
print(y)