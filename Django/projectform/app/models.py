from django.db import models

# Create your models here.
class App(models.Model):
    connection = models.CharField(name='Connection Info', max_length=100)
    credentials = models.CharField(name='Credentials', max_length=100)
    commnad = models.TextField(name='Command')

