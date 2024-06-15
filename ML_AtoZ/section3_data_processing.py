import numpy
import pandas
from sklearn.impute import SimpleImputer
from sklearn.compose import ColumnTransformer
from sklearn.preprocessing import OneHotEncoder
from sklearn.preprocessing import LabelEncoder

# import data set
dataset = pandas.read_csv('csv_files/Data.csv')
x = dataset.iloc[:, :-1].values
y = dataset.iloc[:, -1].values

# if missing value, replace with the average value
imputer = SimpleImputer(missing_values=numpy.nan, strategy='mean')
imputer.fit(x[:, 1:3]) # only numerical columns
x[:, 1:3] = imputer.transform(x[:, 1:3])

# encoding independent variables
ct = ColumnTransformer(transformers=[('encoder', OneHotEncoder(), [0])], remainder='passthrough')
x = numpy.array(ct.fit_transform(x))
print(x)

# encoding dependent - convert yes/no to 1/0
le = LabelEncoder()
y = le.fit_transform(y)
print(y)


