import logging

_LOGGER_NAME = "api-logger"


def _setup_custom_logger():
    logger = logging.getLogger(_LOGGER_NAME)
    logger.setLevel(logging.DEBUG)
    return logger


logger = _setup_custom_logger()
