# Testing

https://docs.djangoproject.com/en/3.2/intro/tutorial05/

## Exposing an Issue 
**Issue**: When adding a future question, the expected output should be `False`, but it currently returns `True`

1. Manually confirm there's a bug in [polls/views.py](mysite3/polls/views.py)
```bash
python3.9 manage.py shell

>>> import datetime
>>> from django.utils import timezone
>>> from polls.models import Question
>>> # create a Question instance with pub_date 30 days in the future
>>> future_question = Question(pub_date=timezone.now() + datetime.timedelta(days=30))
>>> # was it published recently?
>>> future_question.was_published_recently()
True
```

2. Update [polls/tests.py](mysite3/polls/tests.py) to contain the test shown above
```python
import datetime

from django.test import TestCase
from django.utils import timezone

from .models import Question


class QuestionModelTests(TestCase):

    def test_was_published_recently_with_future_question(self):
        """
        was_published_recently() returns False for questions whose pub_date
        is in the future.
        """
        time = timezone.now() + datetime.timedelta(days=30)
        future_question = Question(pub_date=time)
        self.assertIs(future_question.was_published_recently(), False)
```

3. To test: `python3.9 manage.py test polls` 

4. Because the test failed because the code should see if question is recent or not. in this case it was not (+30 days)
   * To resolve - update [polls/models.py](mysite3/polls/models.py) to contain `was_published_recently` in `Question` class
   ```python
    import datetime
    from django.db import models
    from django.utils import timezone

    class Question(models.Model):
        question_text = models.CharField(max_length=200)
        pub_date = models.DateTimeField('date published')
        def __str__(self):
            return self.question_text
        def was_published_recently(self):
            now = timezone.now()
            return now - datetime.timedelta(days=1) <= self.pub_date <= now
    ```
   * rerun "Step 3" - `python3.9 manage.py test polls`

**Note**: [polls/tests.py](mysite3/polls/tests.py) contains 2 more examples based on the 
[comprehensive tests documentation](https://docs.djangoproject.com/en/3.2/intro/tutorial05/#more-comprehensive-tests)

## Test Client  

