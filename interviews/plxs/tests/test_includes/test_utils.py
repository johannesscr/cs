from src.includes.utils import Response, data_validator


class TestResponse:
    """
    Test the Response Class
    """
    def test_default(self):
        """
        GIVEN a response is created
        WHEN no parameters are passed
        THEN return the default response of successful
        """
        res = Response().create()
        assert res.status_code == 200
        assert res.headers.get('Content-Type') == 'application/json'
        assert res.json == {}

    def test_valid_data(self):
        """
        GIVEN the response is created
        WHEN parameters are passed
        THEN return a response with those parameters
        """
        res = Response(http_code=403,
                       message='Unauthorised',
                       data={'survivors': 0}).create()
        assert res.status_code == 403
        assert res.headers.get('Content-Type') == 'application/json'
        assert res.json['message'] == 'Unauthorised'
        assert res.json['data'] == {'survivors': 0}
        assert res.json['errors'] == {}


class TestDataValidator:
    """Test the data validator function"""

    def test_purging_data(self):
        """
        GIVEN a schema and data
        WHEN there is data fields that are not in the schema
        THEN validated the data and purge all fields not in the schema
        """
        data = {
            'first_name': 'james',
            'last_name': 'bond'
        }
        schema = {
            'first_name': {
                'type': 'string'
            }
        }
        validated, data = data_validator(schema=schema, data=data)
        assert validated
        assert data == {'first_name': 'james'}

    def test_basic_validation(self):
        """
        GIVEN a schema and data
        WHEN the data has errors
        THEN return false and a dict of errors
        """
        data = {
            'last_name': 'bond'
        }
        schema = {
            'first_name': {
                'type': 'string',
                'required': True
            }
        }
        validated, data = data_validator(schema=schema, data=data)
        assert validated is False
        assert data == {'first_name': ['required field']}

