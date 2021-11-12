# API 

0. Assert that the web application is running - `cd mysite ; python3.9 manage.py runserver`

2. Access the API interface - `python3.9 manage.py shell`

3. Create a new question (based on [polls/models.py](mysite/polls/models.py)) and save in the database
```python
from polls.models import Choice, Question
from polls.models import Choice, Question

# create a new question
q = Question(question_text="What's new?", pub_date=timezone.now())

# save question
q.save()

# view all questions 
Question.objects.all()
```