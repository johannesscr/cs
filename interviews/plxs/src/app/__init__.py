import logging
from flask import Flask
from src.includes.extends import cors
from src.app.config import get_config

logging.basicConfig(level=logging.DEBUG)


def create_server(env='.env'):
    """
    Creates an application instance of the server.
    """

    app = Flask(__name__)
    # setup server config
    app.config.from_object(get_config(env))

    # initialise extensions
    cors.init_app(app)

    # register error handlers
    from src.routes.error_handler import (error_handler_400,
                                          error_handler_404)
    app.register_error_handler(400, error_handler_400)
    app.register_error_handler(404, error_handler_404)

    # register blueprints
    from src.routes.survivor_count import survivor_count_
    app.register_blueprint(survivor_count_)

    return app
