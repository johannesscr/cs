from flask import Blueprint, request, abort
from src.includes.utils import Response, data_validator
from src.includes.survivors_count import survivors_count

survivor_count_ = Blueprint('survivor_count', __name__)


@survivor_count_.route('/survivorCount', methods=['POST'])
def post_survivor_count():
    """
    Endpoint is to return the survivor count for a given field grouped by
    bins. That is creating a histogram based on a field with specified
    number of bins.
    """
    if not request.is_json:
        abort(400, {
            'content-type': ['incorrect body type', 'request body type json']
        })
    schema = {
        'data': {
            'type': 'list',
            'schema': {
                'type': 'dict',
                'schema': {
                    'PassengerId': {
                        'type': 'integer',
                        'required': True
                    },
                    'Survived': {
                        'type': 'integer',
                        'required': True
                    },
                    'Pclass': {
                        'type': 'integer',
                        'required': True
                    },
                    'Name': {
                        'type': 'string',
                        'required': True
                    },
                    'Sex': {
                        'type': 'string',
                        'required': True
                    },
                    'Age': {
                        'type': 'integer',
                        'required': True
                    },
                    'SibSp': {
                        'type': 'integer',
                        'required': True
                    },
                    'Parch': {
                        'type': 'integer',
                        'required': True
                    },
                    'Ticket': {
                        'type': 'string',
                        'required': True
                    }
                }
            },
            'minlength': 1,
            'required': True
        },
        'binBoundaries': {
            'type': 'list',
            'schema': {'type': 'integer'},
            'required': True
        },
        'binField': {
            'type': 'string',
            'required': True
        }
    }
    validated, data = data_validator(schema=schema, data=request.json)
    if validated is False:
        abort(400, data)

    if not isinstance(data['data'][0][data['binField']], int):
        abort(404, {
            data['binField']: [
                'each of the data {} must be of type '
                'integer'.format(data['binField'])
            ]
        })

    counts = survivors_count(data=data['data'],
                             bin_field=data['binField'],
                             bin_boundaries=data['binBoundaries'])
    return Response(message='survivors counted successfully',
                    data={'counts': counts}).create()
