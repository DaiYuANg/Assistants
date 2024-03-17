from injector import Module, Binder
from peewee import *

db = SqliteDatabase('./people.db')

class StoreModule(Module):
    def configure(self, binder: Binder):
        binder.bind(SqliteDatabase, to=db)
