# Django 

Based on https://www.tutorialspoint.com/django/index.htm

## Requirement
https://www.tutorialspoint.com/django/django_environment.htm

*  Install django - `pip3 install django` 


## Create a Project
https://www.tutorialspoint.com/django/django_creating_project.htm 

1. Create project - create a "myproject" folder with the following structure
   * **Command**: `django-admin startproject myproject`
```buildoutcfg
myproject/
   manage.py        <-- project local django-admin for interacting with your project via command line
   myproject/       <-- python package of your project
      __init__.py   
      settings.py   <--  name indicates, your project settings. 
      urls.py       <-- All links of your project and the function to call.
      wsgi.py       <-- Support for WSGI deployment 
```

2. manage.py help information: `python3.9 manage.py help`


3. Run server to make sure it works: `python3.9 manage.py runserver` 
   * To validate, in your browser goto: http://127.0.0.1:8000/


## Creating an App  

