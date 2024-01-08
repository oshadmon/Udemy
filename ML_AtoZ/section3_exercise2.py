# Importing the necessary libraries
import pandas
import numpy
from sklearn.impute import SimpleImputer


# Load the dataset
dataset = pandas.read_csv('csv_files/pima-indians-diabetes.csv')

# Identify missing data (assumes that missing data is represented as NaN)
missing_data = dataset.isnull().sum()

# Print the number of missing entries in each column
print(f"Missing Data:\n{missing_data}")


# Configure an instance of the SimpleImputer class
imputer = SimpleImputer(missing_values=numpy.nan, strategy='mean')

# Fit the imputer on the DataFrame
imputer.fit(dataset)

# Apply the transform to the DataFrame
dataset_imputer = imputer.transform(dataset)

#Print your updated matrix of features
print(f"Updated matrix:\n{dataset_imputer}")