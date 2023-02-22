from src.includes.survivors_count import find_bin, survivors_count


class TestFindBin:
    """
    Test the find bin functionality
    """

    def test_lowest_bin(self):
        """
        GIVEN the value
        WHEN it is less than the first boundary
        THEN return index of 0
        """
        index = find_bin(value=4, bins=[5])
        assert index == 0

        index = find_bin(value=12, bins=[20, 25, 30])
        assert index == 0

    def test_highest_bin_single_bin(self):
        """
        GIVEN the value
        WHEN it is higher than the last boundary
        THEN return index of length of the list of bins + 1
        """
        index = find_bin(value=6, bins=[5])
        assert index == 1

    def test_highest_bin_multiple_bins(self):
        """
        GIVEN the value
        WHEN it is higher than the last boundary
        THEN return index of length of the list of bins + 1
        """
        index = find_bin(value=100, bins=[20, 25, 30])
        assert index == 3

    def test_upper_limit_of_bin_single_bin(self):
        """
        GIVEN the value
        WHEN it is at the upper limit of a boundary
        THEN return index of that boundary

        :example:
            bin_boundaries = [2, 5]
            values = 5

            the boundaries will be greater than the previous boundary if
            there is one and be less than or equal to the upper boundary.
            Thus.
            index 0 is (-inf, 2]
            index 1 is (2, inf)
        """
        index = find_bin(value=20, bins=[20])
        assert index == 0

    def test_upper_limit_of_bin_multiple_bins(self):
        """
        GIVEN the value
        WHEN it is at the upper limit of a boundary
        THEN return index of that boundary

        :example:
            bin_boundaries = [2, 5]
            values = 5

            the boundaries will be greater than the previous boundary if
            there is one and be less than or equal to the upper boundary.
            Thus.
            index 0 is (-inf, 2]
            index 1 is (2, 5]
            index 2 is (5, inf)
        """
        index = find_bin(value=55, bins=[20, 25, 30, 55])
        assert index == 3

    def test_center_of_bin_single_bin(self):
        """
        GIVEN the value
        WHEN it is at the center of a bin
        THEN return index of that bin
        """
        index = find_bin(value=5, bins=[20])
        assert index == 0

        index = find_bin(value=25, bins=[20])
        assert index == 1

    def test_center_of_bin_multiple_bins(self):
        """
        GIVEN the value
        WHEN it is at the center of a bin
        THEN return index of that bin
        """
        index = find_bin(value=33, bins=[25, 40])
        assert index == 1

        index = find_bin(value=27, bins=[20, 25, 30, 55])
        assert index == 2

        index = find_bin(value=41, bins=[20, 25, 30, 55])
        assert index == 3


class TestSurvivorCount:
    """
    Test the survivor count function
    """

    def test_single_entry_single_bin_lower_bin(self):
        """
        GIVEN a dataset with a single entry and a single bin
        WHEN the function is called and the value is in the lower bin
        THEN return a [1, 0]
        """
        data = [
            {'uno': 'uno', 'numberOfKids': 12, 'tres': 100, 'Survived': 1},
        ]
        bin_field = 'numberOfKids'
        bin_boundaries = [20]
        counts = survivors_count(data=data,
                                 bin_field=bin_field,
                                 bin_boundaries=bin_boundaries)
        assert counts == [1, 0]

    def test_single_entry_single_bin_upper_bin(self):
        """
        GIVEN a dataset with a single entry and a single bin
        WHEN the function is called and the value is in the upper bin
        THEN return a [0, 1]
        """
        data = [
            {'uno': 'uno', 'numberOfKids': 21, 'tres': 100, 'Survived': 1},
        ]
        bin_field = 'numberOfKids'
        bin_boundaries = [20]
        counts = survivors_count(data=data,
                                 bin_field=bin_field,
                                 bin_boundaries=bin_boundaries)
        assert counts == [0, 1]

    def test_multiple_entries_single_bin(self):
        """
        GIVEN a dataset with a multiple entries and a single bin
        WHEN the function is called
        THEN return the expected counts
        """
        data = [
            {'uno': 'uno', 'age': 12, 'tres': 100, 'Survived': 1},
            {'uno': 'uno', 'age': 12, 'tres': 100, 'Survived': 0},
            {'uno': 'uno', 'age': 25, 'tres': 100, 'Survived': 1},
            {'uno': 'uno', 'age': 25, 'tres': 100, 'Survived': 0},
            {'uno': 'uno', 'age': 30, 'tres': 5, 'Survived': 1},
            {'uno': 'uno', 'age': 30, 'tres': 5, 'Survived': 0},
        ]
        bin_field = 'age'
        bin_boundaries = [20]
        counts = survivors_count(data=data,
                                 bin_field=bin_field,
                                 bin_boundaries=bin_boundaries)
        assert counts == [1, 2]

    def test_single_entry_multiple_bins(self):
        """
        GIVEN a dataset with a single entry and multiple bins
        WHEN the function is called
        THEN return the expected counts
        """
        data = [
            {'uno': 'uno', 'age': 12, 'tres': 100, 'Survived': 1},
        ]
        bin_field = 'age'
        bin_boundaries = [20, 40, 50]
        counts = survivors_count(data=data,
                                 bin_field=bin_field,
                                 bin_boundaries=bin_boundaries)
        assert counts == [1, 0, 0, 0]

        data = [
            {'uno': 'uno', 'age': 23, 'tres': 45, 'Survived': 1},
        ]
        counts = survivors_count(data=data,
                                 bin_field=bin_field,
                                 bin_boundaries=bin_boundaries)
        assert counts == [0, 1, 0, 0]

        data = [
            {'uno': 'uno', 'age': 41, 'tres': 100, 'Survived': 1},
        ]
        counts = survivors_count(data=data,
                                 bin_field=bin_field,
                                 bin_boundaries=bin_boundaries)
        assert counts == [0, 0, 1, 0]

        data = [
            {'uno': 'uno', 'age': 1020, 'tres': 10, 'Survived': 1},
        ]
        counts = survivors_count(data=data,
                                 bin_field=bin_field,
                                 bin_boundaries=bin_boundaries)
        assert counts == [0, 0, 0, 1]

    def test_single_entry_multiple_bins_did_not_survive(self):
        """
        GIVEN a dataset with a single entry and multiple bins
        WHEN the function is called
        THEN return the expected counts
        """
        data = [
            {'uno': 'uno', 'age': 12, 'tres': 100, 'Survived': 0},
        ]
        bin_field = 'age'
        bin_boundaries = [20, 40, 50]
        counts = survivors_count(data=data,
                                 bin_field=bin_field,
                                 bin_boundaries=bin_boundaries)
        assert counts == [0, 0, 0, 0]

        data = [
            {'uno': 'uno', 'age': 23, 'tres': 45, 'Survived': 0},
        ]
        counts = survivors_count(data=data,
                                 bin_field=bin_field,
                                 bin_boundaries=bin_boundaries)
        assert counts == [0, 0, 0, 0]

        data = [
            {'uno': 'uno', 'age': 41, 'tres': 100, 'Survived': 0},
        ]
        counts = survivors_count(data=data,
                                 bin_field=bin_field,
                                 bin_boundaries=bin_boundaries)
        assert counts == [0, 0, 0, 0]

        data = [
            {'uno': 'uno', 'age': 1020, 'tres': 10, 'Survived': 0},
        ]
        counts = survivors_count(data=data,
                                 bin_field=bin_field,
                                 bin_boundaries=bin_boundaries)
        assert counts == [0, 0, 0, 0]

    def test_multiple_entries_multiple_bins_lower(self):
        """
        GIVEN a dataset with a multiple entries and multiple bins
        WHEN the function is called
        THEN return the expected counts
        """
        data = [
            {'uno': 'uno', 'age': 12, 'tres': 100, 'Survived': 1},
            {'uno': 'uno', 'age': 19, 'tres': 100, 'Survived': 1},
            {'uno': 'uno', 'age': 43, 'tres': 100, 'Survived': 1},
            {'uno': 'uno', 'age': 41, 'tres': 100, 'Survived': 1},
            {'uno': 'uno', 'age': 45, 'tres': 100, 'Survived': 0},
            {'uno': 'uno', 'age': 50, 'tres': 100, 'Survived': 1},
            {'uno': 'uno', 'age': 55, 'tres': 100, 'Survived': 1},
        ]
        bin_field = 'age'
        bin_boundaries = [20, 40, 50]
        counts = survivors_count(data=data,
                                 bin_field=bin_field,
                                 bin_boundaries=bin_boundaries)
        assert counts == [2, 0, 3, 1]

