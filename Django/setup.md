# Setting Up Project & Application in Django

**Project**: a collection of configuration and apps for a particular website. Project contain multiple apps. An app can be in multiple projects.

**App(lication)**: A web application that does something  Examples: 
   - Weblog system 
   - a database of public records 
   - a small poll app
   
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


3. Run server: 
   * To validate, in your browser goto: http://127.0.0.1:8000/
```bash
python3.9 manage.py runserver

# To user an connection other than 127.0.0.1:8000 specify IP:PORT connnection info
python3.9 manage.py runserver ${IP}:${PORT}
``` 
 


## Creating an App  

https://www.tutorialspoint.com/django/django_apps_life_cycle.htm

1. Create application - create a "myapp" project with the following structure: 
      * **Command**: `cd myproject ; python3.9 manage.py startapp myapp`
```buildoutcfg
myapp/
   __init__.py 
   admin.py     <-- make the app modifiable in the admin interface
   models.py    <-- application models are stored
   tests.py     <-- unit tests
   views.py     <-- application views
```
