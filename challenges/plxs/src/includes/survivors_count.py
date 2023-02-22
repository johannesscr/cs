from math import inf


def find_bin(value: int, bins: list) -> int:
    """
    Function takes a value of type integer, then compares it to the list as
    a set of boundaries, then returns the index of the bin to be incremented.

    :param {int} value: value to be used for comparison
    :param {list} bins: list of integer bin values
    :return: {int} index
    """
    bins.sort()
    # multiple bins
    if len(bins) > 1:
        for i, upper_bound in enumerate(bins):
            if i == 0:
                lower_bound = -inf
            elif i == (len(bins) - 1):
                if value <= upper_bound:
                    return i
                return i + 1
            else:
                lower_bound = bins[i - 1]

            if lower_bound < value <= upper_bound:
                return i

    # single bin condition
    if value <= bins[0]:
        return 0
    return 1


def survivors_count(data: list, bin_field: str, bin_boundaries: list):
    """
    Function goes through the list of data items. Then based on the
    bin_field compare the value to that of the provided bins and increments
    the bin the item falls into.

    NB: edge case not considered is if there is a mixed set of variable
        types for the single field. Assume if the first is of type integer,
        then it will be consistently so throughout the dataset.

    :param {list} data: list of data items.
    :param {str} bin_field: the key of the item field.
    :param {list} bin_boundaries: list of integer bins.
    :return:
    """
    bins = [0 for _ in range(len(bin_boundaries) + 1)]
    for item in data:
        index = find_bin(value=item[bin_field], bins=bin_boundaries)
        bins[index] += item['Survived']
    return bins
