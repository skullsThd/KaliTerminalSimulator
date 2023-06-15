# Variabili
PYTHON = python3
PIP = pip3
DJANGO = django-admin

# Comandi
install:
    $(PIP) install -r requirements.txt

migrate:
    $(DJANGO) migrate

runserver:
    $(DJANGO) runserver

createsuperuser:
    $(DJANGO) createsuperuser

shell:
    $(DJANGO) shell

test:
    $(DJANGO) test

.PHONY: install migrate runserver createsuperuser shell test
