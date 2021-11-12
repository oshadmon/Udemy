# Models & Database 

* **Model** - your database layout, with additional metadata.
  * **Question** - question and a publication date
  * **Choice** - has two fields: the text of the choice and a vote tally. A choice is correlated to a question 

https://docs.djangoproject.com/en/3.2/intro/tutorial02/

## Steps
1. Update [polls/models.py](mysite/polls/models.py) to contain code bellow
```python
# polls/models.py
import datetime
from django.db import models
from django.utils import timezone

class Question(models.Model):
    question_text = models.CharField(max_length=200)
    pub_date = models.DateTimeField('date published')
    def __str__(self):
        return self.question_text

    def was_published_recently(self):
        return self.pub_date >= timezone.now() - datetime.timedelta(days=1)

class Choice(models.Model):
    question = models.ForeignKey(Question, on_delete=models.CASCADE)
    choice_text = models.CharField(max_length=200)
    votes = models.IntegerField(default=0)
    def __str__(self):
        return self.choice_text
```

2. Update `INSTALLED_APPS` in [mysite/settings.py](mysite/mysite/settings.py) to contain 'polls.apps.PollsConfig'

3. Update polls code in manage.py - `python3.9 manage.py makemigrations polls`
```buildoutcfg
# expected output
Migrations for 'polls':
  polls/migrations/0001_initial.py
    - Create model Question
    - Create model Choice
```

4. Update database as a result of the migration in step 3 - `python3.9 manage.py sqlmigrate polls 0001` 

5. Repeat step 3 (`python3.9 manage.py makemigrations polls`) to attach the created tables

## API
0. Assert that the web application is running - `cd mysite ; python3.9 manage.py runserver`


1. Access the API interface - `python3.9 manage.py shell`


2. Create a new question (based on [polls/models.py](mysite/polls/models.py)) and save in the database
```python
from polls.models import Choice, Question
from django.utils import timezone

# create a new question
q = Question(question_text="What's new?", pub_date=timezone.now())

# save question
q.save()

# view all questions 
Question.objects.all()
```

