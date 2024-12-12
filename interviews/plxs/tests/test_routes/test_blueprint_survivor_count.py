from logging import getLogger

log = getLogger(__name__)


class TestBlueprintSurvivorCount:
    """
    Test the survivor count blueprint
    """

    def test_invalid_data_format(self, api):
        """
        GIVEN the endpoint is called
        WHEN the data payload is of the incorrect format
        THEN return a 400 response
        """
        with api:
            payload = {}
            res = api.post('/survivorCount',
                           json=payload,
                           headers={'Content-Type': 'application/json'})
            assert res.status_code == 400
            assert res.json['message'] == 'BadRequest'

    def test_invalid_body_type(self, api):
        """
        GIVEN the endpoint is called
        WHEN when the request body is not of type json
        THEN return a 400 response
        """
        with api:
            payload = {
                'data': [
                    {
                        'PassengerId': 1,
                        'Survived': 1,
                        'Pclass': 12,
                        'Name': 'james',
                        'Sex': 'male',
                        'Age': 23,
                        'SibSp': 1,
                        'Parch': 0,
                        'Ticket': 'A/521171'
                    }
                ],
                'binField': 'Name',
                'binBoundaries': [5, 10]
            }
            res = api.post('/survivorCount',
                           data=payload)
            log.info(res.json)
            assert res.status_code == 400
            assert res.json['errors']['content-type'] == [
                'incorrect body type',
                'request body type json'
            ]

    def test_invalid_field_type(self, api):
        """
        GIVEN the endpoint is called
        WHEN when the bin field is not of type integer
        THEN return a 404 response
        """
        with api:
            payload = {
                'data': [
                    {
                        'PassengerId': 1,
                        'Survived': 1,
                        'Pclass': 12,
                        'Name': 'james',
                        'Sex': 'male',
                        'Age': 23,
                        'SibSp': 1,
                        'Parch': 0,
                        'Ticket': 'A/521171'
                    }
                ],
                'binField': 'Name',
                'binBoundaries': [5, 10]
            }
            res = api.post('/survivorCount',
                           json=payload)
            log.info(res.json)
            assert res.status_code == 404
            assert res.json['message'] == 'NotFound'

    def test_valid_data_format(self, api):
        """
        GIVEN the endpoint is called
        WHEN the data payload is of the correct format
        THEN return 200 with the survivor count
        """
        with api:
            payload = {
                'data': [
                    {
                        'PassengerId': 1,
                        'Survived': 1,
                        'Pclass': 12,
                        'Name': 'james',
                        'Sex': 'male',
                        'Age': 23,
                        'SibSp': 1,
                        'Parch': 0,
                        'Ticket': 'A/521171'
                    }
                ],
                'binField': 'Age',
                'binBoundaries': [5, 10]
            }
            res = api.post('/survivorCount',
                           json=payload)
            assert res.status_code == 200
            assert isinstance(res.json['counts'], list)
            assert res.json['counts'] == [0, 0, 1]

    def test_only_count_survived(self, api):
        """
        GIVEN the endpoint is called
        WHEN the data payload is of the correct format, only if Survived is
            1 should it be counted
        THEN return 200 with the survivor count
        """
        with api:
            payload = {
                'data': [
                    {
                        'PassengerId': 1,
                        'Survived': 0,
                        'Pclass': 12,
                        'Name': 'james',
                        'Sex': 'male',
                        'Age': 23,
                        'SibSp': 1,
                        'Parch': 0,
                        'Ticket': 'A/521171'
                    }
                ],
                'binField': 'Age',
                'binBoundaries': [5, 10]
            }
            res = api.post('/survivorCount',
                           json=payload)
            assert res.status_code == 200
            assert isinstance(res.json['counts'], list)
            assert res.json['counts'] == [0, 0, 0]

    def test_bin_allocation(self, api):
        """
        GIVEN the correct payload is sent
        WHEN when the bins are specified as 5 and 10, with 3 observations such that
            (-inf, 5] = 1
            (5, 10] = 2
            (10, inf) = 0
        THEN that bin count should be returned
        """
        with api:
            payload = {
                'data': [
                    {
                        'PassengerId': 1,
                        'Survived': 1,
                        'Pclass': 12,
                        'Name': 'james',
                        'Sex': 'male',
                        'Age': 5,
                        'SibSp': 1,
                        'Parch': 0,
                        'Ticket': 'A/521171'
                    },
                    {
                        'PassengerId': 1,
                        'Survived': 1,
                        'Pclass': 12,
                        'Name': 'james',
                        'Sex': 'male',
                        'Age': 7,
                        'SibSp': 1,
                        'Parch': 0,
                        'Ticket': 'A/521171'
                    },
                    {
                        'PassengerId': 1,
                        'Survived': 0,
                        'Pclass': 12,
                        'Name': 'james',
                        'Sex': 'male',
                        'Age': 6,
                        'SibSp': 1,
                        'Parch': 0,
                        'Ticket': 'A/521171'
                    }
                ],
                'binField': 'Age',
                'binBoundaries': [5, 10]
            }
            res = api.post('/survivorCount',
                           json=payload)
            assert res.status_code == 200
            assert isinstance(res.json['counts'], list)
            assert res.json['counts'] == [1, 1, 0]
