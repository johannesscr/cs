# Code Challenge

The following is all from the root directory of the server. If cloned from github the root directory would
be `plxs-code-challenge`.

## Initial set-up

The following script is idempotent, `setup.sh` and can be run the the terminal to setup the

```bash
./setup.sh
```

next start the virtual environment for the server with,

```bash
source  ../venv/bin/activate
```

now that the virtual environment is up and running.

## Test Server

To test the server,

```bash
pytest --cov-report term-missing --cov=. tests/
```

all the test should pass.

## Run  Server

To run the server, simply

```bash 
python server.py
```

the server will be running on `http://localhost:5010` or running on localhost listening on port 5010. The port is just
show that the environmental variables are used.

### Endpoints

There is a single enpoint

```text
URL: '/survivorCount' 
Method: POST
Content-Type: application/json
```

The endpoint accepts a body of the form below,where all the fields are required, the API is very strict in that all the
types should be as

```text
{
    "data": [  // type json array of json objects
        {
            "PassengerId": "integer"
            "Survived": "integer",
            "Pclass": "integer",
            "Name": "string",
            "Sex": "string",
            "Age": "integer",
            "SibSp": "integer",
            "Parch": "integer",
            "Ticket": "string"
        },
        ...
    ],
    "binField": "string",
    "binBoundaries": [5, 10]  // type json array of integers
}
```

an example body

```json
{
    "data": [
        {
            "PassengerId": 1,
            "Survived": 0,
            "Pclass": 12,
            "Name": "james",
            "Sex": "male",
            "Age": 23,
            "SibSp": 1,
            "Parch": 0,
            "Ticket": "A/521171"
        },
        {
            "PassengerId": 1,
            "Survived": 1,
            "Pclass": 12,
            "Name": "james",
            "Sex": "male",
            "Age": 23,
            "SibSp": 1,
            "Parch": 0,
            "Ticket": "A/521171"
        }
    ],
    "binField": "Age",
    "binBoundaries": [5, 10]
}
```

#### Example Request

start the server, and run the following `curl` command to make a request to 
the server, the data is that of th example body above

```bash
curl -X POST http://localhost:5010/survivorCount \
    -H "Content-Type: application/json" \
    -d '{"data": [{"PassengerId": 1,"Survived": 0,"Pclass": 12,"Name": "james","Sex": "male","Age": 23,"SibSp": 1,"Parch": 0,"Ticket": "A/521171"},{"PassengerId": 1,"Survived": 1,"Pclass": 12,"Name": "james","Sex": "male","Age": 23,"SibSp": 1,"Parch": 0,"Ticket": "A/521171"}],"binField": "Age","binBoundaries": [5, 10]}'
```

the expected response should be

```bash 
{
  "counts": [
    0,
    0,
    1
  ]
}
```

## Final Note  
The implementation has been built in such a manner to show how I would 
implement more complex systems. This implementation could have been 
implemented a lot more simply with a file structure as
```text
plxs-code-challenge/
|- tests/
|  |- test_server.py
|- requirements.txt
|- server.py
|- setup.sh
```
The `requirements.txt` and `setup.sh` files would remain similar. The 
`server.py` would be able to contain all the needed necessary functions as 
shown below. 

```python
from flask import Flask, jsonify
from flask_cors import CORS

app = Flask(__name__)
cors = CORS()
cors.init_app(app)
 

def find_bin(...):
    ...


def survivor_count(...):
    # use the find_bin function
    ...


@app.route('/survivorCount', methods=['POST'])
def post_survivor_count():
    # basic validation for 400 and 404 responses
    counts = survivor_count(...)
    return jsonify({'counts': counts})

if __name__ == '__main__':
    app.run(host='localhost', port=5010, debug=True)
```

