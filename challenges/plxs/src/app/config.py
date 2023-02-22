from dotenv import dotenv_values
from collections import namedtuple


def get_config(env):
    """
    Function take the environment name from which to load the config
    :param env: name of the dot env file
    :return: object of environmental variables
    """
    env_config = dotenv_values(env)
    config = namedtuple("Config", env_config.keys())(*env_config.values())
    return config
