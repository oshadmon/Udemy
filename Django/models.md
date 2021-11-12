# Models

* **Model** - your database layout, with additional metadata.
  * **Question** - question and a publication date
  * **Choice** - has two fields: the text of the choice and a vote tally. A choice is correlated to a question 

## Steps
1. Update [polls/models.py](mysite/polls/models.py) to contain code bellow
```python
# polls/models.py
from django.db import models


class Question(models.Model):
    question_text = models.CharField(max_length=200)
    pub_date = models.DateTimeField('date published')


class Choice(models.Model):
    question = models.ForeignKey(Question, on_delete=models.CASCADE)
    choice_text = models.CharField(max_length=200)
    votes = models.IntegerField(default=0)
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


The next step is to play with the [API](api.md) tools Django provides in order to test the models created.