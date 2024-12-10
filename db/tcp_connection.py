from dependency_injector.wiring import inject, Container
from pony.orm import Required, Set
from pony.orm.core import Entity


class Person(Entity):
    name = Required(str)
    age = Required(int)
    cars = Set("Car")
