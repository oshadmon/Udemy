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

## Interactive Views

* A view is responsible for doing one of two things: 
  1. returning an _HttpResponse_ object containing the content for the requested page
  2. raising an exception such as _Http404_

**Example**: displays the latest 5 poll questions in the system, separated by commas, according to publication date
1. Update [polls/views.py](mysite/polls/views.py) to have an updated index 
```python
from django.shortcuts import render
from django.template import loader

def index(request):
    latest_question_list = Question.objects.order_by('-pub_date')[:5]
    template = loader.get_template('polls/index.html')
    context = {
        'latest_question_list': latest_question_list,
    }
    return render(request, 'polls/index.html', {})
```

2. Create [polls/index.html](mysite/polls/templates/polls/index.html) which contains to read [polls/views.py](mysite/polls/views.py)
```html
{% if latest_question_list %}
    <ul>
    {% for question in latest_question_list %}
        <li><a href="/polls/{{ question.id }}/">{{ question.question_text }}</a></li>
    {% endfor %}
    </ul>
{% else %}
    <p>No polls are available.</p>
{% endif %}
```

3. 