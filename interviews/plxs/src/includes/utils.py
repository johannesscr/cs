from flask import jsonify
from cerberus import Validator


class Response:
    def __init__(self,
                 http_code: int = None,
                 message: str = None,
                 data: dict = None,
                 errors: dict = None):
        self.status_code = http_code or 200
        self.message = message or 'successful'
        self.data = data or {}
        self.errors = errors or {}

    def create(self):
        # this would normally give a generic response.
        # however the requirement is only return {counts: [..., ...]}
        if self.status_code == 200:
            resp = jsonify(self.data)
        else:
            resp = jsonify(message=self.message,
                           data=self.data,
                           errors=self.errors)
        resp.status_code = self.status_code
        resp.headers.add('Content-Type', 'application/json')
        return resp


def data_validator(schema: dict, data: dict) -> [bool, dict]:
    """
    Function takes a schema for the expected data structure. Then validates
    the data. It purges all data not in the schema and returns a bool
    denoting if the validation has passed

    :param {dict} schema: schema of the expected data structure
    :param {dict} data: the data to be validated
    :return: {dict} the validated and purged data
    """
    v = Validator(schema, purge_unknown=True)
    if v.validate(data):
        return True, v.normalized(data)
    return False, v.errors

