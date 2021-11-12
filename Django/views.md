# Views 

https://docs.djangoproject.com/en/3.2/intro/tutorial03/

## Argument Options

1. Update [polls/views.py](mysite/polls/views.py) to support different types of questions
```python
from django.http import HttpResponse

def detail(request, question_id):
    return HttpResponse("You're looking at question %s." % question_id)

def results(request, question_id):
    response = "You're looking at the results of question %s."
    return HttpResponse(response % question_id)

def vote(request, question_id):
    return HttpResponse("You're voting on question %s." % question_id)
```

2. Updated [polls/urls.py](mysite/polls/urls.py) to support the newly added modules 
```python
from django.urls import path

from . import views

urlpatterns = [
    # ex: /polls/5/
    path('<int:question_id>/', views.detail, name='detail'),
    # ex: /polls/5/results/
    path('<int:question_id>/results/', views.results, name='results'),
    # ex: /polls/5/vote/
    path('<int:question_id>/vote/', views.vote, name='vote'),
]
```

3. Validate the behavior is as expected by visiting each of the sites
   * Question details - http://127.0.0.1:8000/polls/5/
   * Question results - http://127.0.0.1:8000/polls/5/results/
   * [Question vote - http://127.0.0.1:8000/polls/5/vote/

## 