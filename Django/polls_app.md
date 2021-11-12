# Django Polls App 

https://docs.djangoproject.com/en/3.2/intro/tutorial01/

1. Create project - `django-admin startproject mysite1`
 ```buildoutcfg
mysite/
    manage.py
    mysite/
        __init__.py
        settings.py
        urls.py
        asgi.py
        wsgi.py
```

2. Test connection: `python3.9 manage.py runserver`


3. Create a new app for polls: `cd mysite1 ; python3.9 manage.py startapp polls`
```buildoutcfg
polls/
    __init__.py
    admin.py
    apps.py
    migrations/
        __init__.py
    models.py
    tests.py
    views.py
```


4. Update content in [polls/views.py](mysite1/polls/views.py) to contain a request
```python
from django.http import HttpResponse


def index(request):
    return HttpResponse("Hello, world. You're at the polls index.")
```

5. Update content in [polls/urls.py](mysite1/polls/urls.py) to contain call to [polls/views.py](mysite/polls/views.py) 
```python
from django.urls import path

from . import views

urlpatterns = [
    path('', views.index, name='index'),
]
```

6. Update [mysite/urls.py](mysite1/mysite1/urls.py) to contain calls include `polls.urls`
```python
from django.urls import include, path

urlpatterns = [
    path('polls/', include('polls.urls')),
]
```

**Note**: I added the [admin configs](admin_interface.md) between steps 2 & 3 