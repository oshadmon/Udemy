from django.shortcuts import render

# Create your views here.
from django.http import HttpResponse
import requests

def index(request):
    try:
        r = requests.get('http://76.214.179.41:2051', headers={'command': 'get status', 'User-Agent': 'AnyLog/1.23' })
    except Exception as e:
        output = "Failed to execute GET status against '76.214.179.41:2051' (Error: %s)" % e
    else:
        if int(r.status_code) != 200:
            output = "Failed to execute GET status against '76.214.179.41:2051' (Network Error: %s)" % r.status_code
        else:
            try:
                output =  r.text
            except Exception as e:
                output = "Failed to extract GET status results from '76.214.179.41:2051' (Error: %s)" % e

    return HttpResponse(output)
