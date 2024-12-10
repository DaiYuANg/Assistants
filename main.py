# 这是一个示例 Python 脚本。
import sys

from dotenv import load_dotenv
from loguru import logger

from db.container import DatabaseContainer
from ui.container import Container

load_dotenv()
logger.add(
    sys.stdout, colorize=True, format="<green>{time}</green> <level>{message}</level>"
)
if __name__ == "__main__":
    logger.debug("That's it, beautiful and simple logging!")
    uiContainer = Container()
    databaseContainer = DatabaseContainer()

    app = uiContainer.app()

    window = uiContainer.window()

    window.center()
    window.show()

    sys.exit(app.exec())
