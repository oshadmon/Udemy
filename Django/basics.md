# Basic 

## View & Templates 
https://www.tutorialspoint.com/django/django_creating_views.htm
* [templates](myproject/myapp/templates) - directory containing python scripts correlated to html pages 
* [views.py](myproject/myapp/views.py) - python file containing calls to files in [templates](myproject/myapp/templates)

```python
# myapp/templates/hello.html
from django.http import HttpResponse

def hello(request):
   text = """<h1>welcome to my app !</h1>"""
   return HttpResponse(text)


# myapp/views.py 
from django.shortcuts import render

def hello(request):
   return render(request, "myapp/template/hello.html", {})
```

## URL Mapping
https://www.tutorialspoint.com/django/django_url_mapping.htm

* [url.py](myproject/myproject/urls.py) - list of URL links correlated to the application

