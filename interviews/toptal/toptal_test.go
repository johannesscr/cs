package toptal

import "testing"

func TestStoreEmptyAt(t *testing.T) {
	tt := []struct {
		name string
		customerArrivals [][2]int
		numberOfCashiers int
		output int
	}{
		{
			name: "test one",
			customerArrivals: [][2]int{
				{0, 7}, {1, 3}, {1, 4},
			},
			numberOfCashiers: 3,
			output: 7,
		},
		{
			name: "test two",
			customerArrivals: [][2]int{},
			numberOfCashiers: 3,
			output: 0,
		},
		{
			name: "test three",
			customerArrivals: [][2]int{
				{0, 2}, {1, 2}, {2, 3},
			},
			numberOfCashiers: 3,
			output: 5,
		},
		{
			name: "test four",
			customerArrivals: [][2]int{
				{0, 2}, {0, 2}, {2, 3}, {2, 5},
			},
			numberOfCashiers: 3,
			output: 7,
		},
		{
			name: "test five",
			customerArrivals: [][2]int{
				{0, 2}, {5, 2},
			},
			numberOfCashiers: 3,
			output: 7,
		},
		{
			name: "test six",
			customerArrivals: [][2]int{
				{0, 2}, {0, 6}, {0, 3}, {2, 4}, {5, 3},
			},
			numberOfCashiers: 4,
			output: 8,
		},
		{
			name: "test seven",
			customerArrivals: [][2]int{
				{0, 3}, {0, 5}, {0, 2}, {2, 3}, {2, 1}, {4, 2}, {5, 3}, {5, 7}, {5, 2}, {5, 5}, {30, 10},
			},
			numberOfCashiers: 3,
			output: 40,
		},
		{
			name: "test eight",
			customerArrivals: [][2]int{
				{0, 3}, {0, 5}, {0, 2}, {20, 3}, {20, 1}, {40, 2}, {50, 3}, {50, 7}, {50, 2}, {50, 5}, {300, 10},
			},
			numberOfCashiers: 3,
			output: 310,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			o := storeEmptyAt(tc.customerArrivals, tc.numberOfCashiers)
			if o != tc.output {
				t.Errorf("Expected %d got %d", tc.output, o)
			}
		})
	}
}

func TestStoreEmptyAt2(t *testing.T) {
	tt := []struct {
		name string
		customerArrivals [][2]int
		numberOfCashiers int
		output int
	}{
		{
			name: "test one",
			customerArrivals: [][2]int{
				{0, 7}, {1, 3}, {1, 4},
			},
			numberOfCashiers: 3,
			output: 7,
		},
		{
			name: "test two",
			customerArrivals: [][2]int{},
			numberOfCashiers: 3,
			output: 0,
		},
		{
			name: "test three",
			customerArrivals: [][2]int{
				{0, 2}, {1, 2}, {2, 3},
			},
			numberOfCashiers: 3,
			output: 5,
		},
		{
			name: "test four",
			customerArrivals: [][2]int{
				{0, 2}, {0, 2}, {2, 3}, {2, 5},
			},
			numberOfCashiers: 3,
			output: 7,
		},
		{
			name: "test five",
			customerArrivals: [][2]int{
				{0, 2}, {5, 2},
			},
			numberOfCashiers: 3,
			output: 7,
		},
		{
			name: "test seven",
			customerArrivals: [][2]int{
				{0, 3}, {0, 5}, {0, 2}, {2, 3}, {2, 1}, {4, 2}, {5, 3}, {5, 7}, {5, 2}, {5, 5}, {30, 10},
			},
			numberOfCashiers: 3,
			output: 40,
		},
		{
			name: "test eight",
			customerArrivals: [][2]int{
				{0, 3}, {0, 5}, {0, 2}, {20, 3}, {20, 1}, {40, 2}, {50, 3}, {50, 7}, {50, 2}, {50, 5}, {300, 10},
			},
			numberOfCashiers: 3,
			output: 310,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			o := storeEmptyAt2(tc.customerArrivals)
			if o != tc.output {
				t.Errorf("Expected %d got %d", tc.output, o)
			}
		})
	}
}

func BenchmarkStoreEmptyAt(b *testing.B) {
	customerArrivals := [][2]int{
		{0, 3}, {0, 5}, {0, 2}, {2, 3}, {2, 1}, {4, 2}, {5, 3}, {5, 7}, {5, 2}, {5, 5}, {30, 10},
	}
	for i := 0; i < b.N; i++ {
		storeEmptyAt(customerArrivals, 3)
	}
}

func BenchmarkStoreEmptyAt_2(b *testing.B) {
	customerArrivals := [][2]int{
		{0, 3}, {0, 5}, {0, 2}, {20, 3}, {20, 1}, {40, 2}, {50, 3}, {50, 7}, {50, 2}, {50, 5}, {300, 10},
	}
	for i := 0; i < b.N; i++ {
		storeEmptyAt(customerArrivals, 3)
	}
}

func BenchmarkStoreEmptyAt2(b *testing.B) {
	customerArrivals := [][2]int{
		{0, 3}, {0, 5}, {0, 2}, {2, 3}, {2, 1}, {4, 2}, {5, 3}, {5, 7}, {5, 2}, {5, 5}, {30, 10},
	}
	for i := 0; i < b.N; i++ {
		storeEmptyAt2(customerArrivals)
	}
}

func BenchmarkStoreEmptyAt2_2(b *testing.B) {
	customerArrivals := [][2]int{
		{0, 3}, {0, 5}, {0, 2}, {20, 3}, {20, 1}, {40, 2}, {50, 3}, {50, 7}, {50, 2}, {50, 5}, {300, 10},
	}
	for i := 0; i < b.N; i++ {
		storeEmptyAt2(customerArrivals)
	}
}