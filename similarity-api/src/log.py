import logging

_LOGGER_NAME = "api-logger"


def _setup_custom_logger():
    format = logging.Formatter(
        '%(asctime)s - %(name)s - %(levelname)s - %(message)s')
    handler = logging.StreamHandler()
    handler.setFormatter(format)
    logger = logging.getLogger(_LOGGER_NAME)
    logger.setLevel(logging.DEBUG)
    logger.addHandler(handler)
    return logger


logger = _setup_custom_logger()
