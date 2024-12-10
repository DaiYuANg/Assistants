from dependency_injector import containers, providers
from pony.orm import Database


class DatabaseContainer(containers.DeclarativeContainer):
    database = Database()
    db_url = "database.sqlite"
    database.bind(provider="sqlite", filename=db_url, create_db=True)
    database.generate_mapping(create_tables=True)
    db = providers.Singleton(db=database)
