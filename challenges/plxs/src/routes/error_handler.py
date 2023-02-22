from src.includes.utils import Response


def error_handler_400(errors):
    """
    Function to return a generic 400 response

    :param {dict} errors: each key is an error point with a
        list of error descriptions
    :return: response
    """
    return Response(http_code=400,
                    message='BadRequest',
                    errors=errors.description).create()


def error_handler_404(errors):
    """
    Function to return a generic 404 response

    :param {dict} errors: each key is an error point with a
        list of error descriptions
    :return: response
    """
    return Response(http_code=404,
                    message='NotFound',
                    errors=errors.description).create()
