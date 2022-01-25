from django.http import HttpResponse
from django.shortcuts import render
from .forms import AppForm

# Create your views here.

def newapp(request):
    if request.method == 'POST':
        form = AppForm(request.POST)
        if form.is_valid():
            form.save()

    elif request.method == 'GET':
        form = AppForm()
        context = {
            'form': form
        }

    return render(request, 'app/new.html', context)