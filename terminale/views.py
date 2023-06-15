from django.shortcuts import render

# Create your views here.
def terminale(request):
	return render(request, 'terminale/terminal.html ')