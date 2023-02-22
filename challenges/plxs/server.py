import os  # pragma no cover
from src.app import create_server  # pragma no cover

env = os.getenv('ENV')  # pragma no cover
app = create_server()  # pragma no cover

if __name__ == '__main__':  # pragma no cover
    app.run(host=app.config['HOST'],
            port=app.config['PORT'],
            debug=app.config['DEBUG'])
