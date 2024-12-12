from contextlib import AbstractContextManager, contextmanager
from typing import Callable

from dependency_injector import containers, providers
from loguru import logger
from sqlalchemy import create_engine, Engine, orm
from sqlalchemy.orm import declarative_base, Session

from db.tcp_connection import Base


class Database:

  def __init__(self, db_url: str) -> None:
    self._engine = create_engine(db_url, echo=True)
    self._session_factory = orm.scoped_session(
      orm.sessionmaker(
        autocommit=False,
        autoflush=False,
        bind=self._engine,
      ),
    )

  def create_database(self) -> None:
    Base.metadata.create_all(self._engine)

  @contextmanager
  def session(self) -> Callable[..., AbstractContextManager[Session]]:
    session: Session = self._session_factory()
    try:
      yield session
    except Exception:
      logger.exception("Session rollback because of exception")
      session.rollback()
      raise
    finally:
      session.close()


class DatabaseModule(containers.DeclarativeContainer):
  db = providers.Singleton(Database, "sqlite+pysqlite:///:memory:")
