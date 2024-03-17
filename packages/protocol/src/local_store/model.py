from peewee import *

from .store_module import db


class Project(Model):
    name = CharField()
    birthday = DateField()

    class Meta:
        database = db
