# 这是一个示例 Python 脚本。
import sys

from dependency_injector import containers, providers
from dotenv import load_dotenv
from loguru import logger

from db.container import DatabaseModule
from ui.module import UIModule

# from db.container import DatabaseContainer
# from ui.module import Container

load_dotenv()
logger.add(
  sys.stdout, colorize=True, format="<green>{time}</green> <level>{message}</level>"
)

class ApplicationContainer(containers.DeclarativeContainer):
  ui_module = providers.Container(UIModule)
  db_module = providers.Container(DatabaseModule)


if __name__ == "__main__":
  logger.debug("That's it, beautiful and simple logging!")
  container = ApplicationContainer()
  uiContainer = UIModule()
  # databaseContainer = DatabaseContainer()

  app = uiContainer.app()

  window = uiContainer.window()

  window.center()
  window.show()

  sys.exit(app.exec())
