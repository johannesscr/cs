import pytest
from src.app import create_server
from logging import getLogger

log = getLogger(__name__)


@pytest.fixture(scope='session')
def app():
    _app = create_server(env='test')
    _app.config.TESTING = True
    with _app.app_context():
        log.info(_app.config)
        yield _app
    log.info('teardown APP')


@pytest.fixture(scope='session')
def api(app):
    with app.app_context():
        # Yield the testing client
        yield app.test_client()

    log.info('teardown api')


