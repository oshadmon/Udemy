# Admin Interface

The Admin interface depends on the django.countrib module. To have it working you need to make sure some modules are 
imported in the INSTALLED_APPS and MIDDLEWARE_CLASSES tuples of the `myproject/settings.py` file.

https://www.tutorialspoint.com/django/django_admin_interface.htm

## Packages
* **[Installed Apps](myproject/myproject/settings.py#L33)** - Django tools required to be installed  
```python3
INSTALLED_APPS = (
   'django.contrib.admin',
   'django.contrib.auth',
   'django.contrib.contenttypes',
   'django.contrib.sessions',
   'django.contrib.messages',
   'django.contrib.staticfiles',
   'myapp'
)
```

* **[Middleware](myproject/myproject/settings.py#L33)** - Django services required to be configured
```python
MIDDLEWARE_CLASSES = (
   'django.contrib.sessions.middleware.SessionMiddleware',
   'django.middleware.common.CommonMiddleware',
   'django.middleware.csrf.CsrfViewMiddleware',
   'django.contrib.auth.middleware.AuthenticationMiddleware',
   'django.contrib.messages.middleware.MessageMiddleware',
   'django.middleware.clickjacking.XFrameOptionsMiddleware',
)
```

## Configure Admin interface 

0. Make sure you are in `myproject` directory


1. Initiate database - syncdb will create necessary tables or collections depending on your db type, necessary for the admin interface to run.
   * **Command**: `python3.9 manage.py migrate` 
   
**Note**: To create a superuser execute: `python3.9 manage.py createsuperuser`

2. Run server - `cd $HOME/myproject ; python3.9 manage.py runserver`


3. Go to admin page & access using the created superuser - http://127.0.0.1:8000/admin/  


That interface will let you administrate Django groups and users, and all registered models in your app.

The interface gives you the ability to do at least the "CRUD" (Create, Read, Update, Delete) operations on your models.

