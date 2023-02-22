package toptal

/*
PROBLEM STATEMENT

Given a store with K cashiers. Each cashier processes customers independently
of the other cashiers. Once a cashier is finished with a customer the cashier
can directly continue with the next customer in the queue.

Customers arrive sequentially, in the format [arrivalTime, processingTime],
where arrivalTime denotes the time the customer arrives at the queue to
be assisted by a cashier, processingTime denotes the time the cashier takes
to assist the customer. Let A denote customer arrivals, where
A = [[arrivalTime, processingTime], [0, 3], [3, 4]], such that customer 1
arrives at time 0 and takes 3 units of time to process.

Customers queue and go to the first available cashier. Find the time it takes
for the store to be empty
*/

type cashier struct {
	busy bool
	timeFinished int
}

// findFirstFinished finds and returns the cashier's index who will be
// finished first.
func findFirstFinished(cashierTimes []int) int {
	index := -1
	m := -1
	for i, t := range cashierTimes {
		if t < m || m == -1 {
			index = i
			m = t
		}
	}
	return index
}

//findMax
func findMax(cashiers []cashier) int {
	//fmt.Println(cashiers)
	m := 0
	for _, c := range cashiers {
		if c.timeFinished > m {
			m = c.timeFinished
		}
	}
	return m
}

func findMaxTime(cashierTimes []int) int {
	var m int
	for _, t := range cashierTimes {
		if t > m {
			m = t
		}
	}
	return m
}

func nextArrival(customerArrivals *[][2]int) [2]int {
	var customer [2]int
	if len(*customerArrivals) > 1 {
		customer = (*customerArrivals)[0]
		*customerArrivals = (*customerArrivals)[1:]
	} else if len(*customerArrivals) == 1 {
		customer = (*customerArrivals)[0]
		*customerArrivals = nil
	}
	return customer
}

func nextCustomer(customerArrivals *[][2]int, t int) [2]int {
	//fmt.Println(*customerArrivals)
	if len(*customerArrivals) == 0 {
		return [2]int{0, 0}
	}
	if t >= (*customerArrivals)[0][0] {
		c := (*customerArrivals)[0]
		if len(*customerArrivals) > 0 {
			*customerArrivals = (*customerArrivals)[1:]
		} else {
			*customerArrivals = nil
		}
		return c
	}
	return [2]int{0, 0}
}

func storeEmptyAt(customerArrivals [][2]int, numberOfCashiers int) int {
	cashierTimes := make([]int, numberOfCashiers)

	for len(customerArrivals) > 0 {
		i := findFirstFinished(cashierTimes)
		customer := nextArrival(&customerArrivals)
		// cashier waiting time for customer
		if cashierTimes[i] < customer[0] {
			cashierTimes[i] = customer[0]
		}
		// now add the service time
		cashierTimes[i] = cashierTimes[i] + customer[1]
		//fmt.Println("E", cashierTimes, i, cashierTimes[i], customer)
	}
	return findMaxTime(cashierTimes)
}

func storeEmptyAt2(customerArrivals [][2]int) int {
	t := 0
	cashier1 := cashier{}
	cashier2 := cashier{}
	cashier3 := cashier{}
	for len(customerArrivals) > 0 {
		if t >= cashier1.timeFinished {
			customer := nextCustomer(&customerArrivals, t)
			cashier1.timeFinished = t + customer[1]
			//fmt.Println(cashier1, t, customer)
		}
		if t >= cashier2.timeFinished {
			customer := nextCustomer(&customerArrivals, t)
			cashier2.timeFinished = t + customer[1]
			//fmt.Println(cashier2, t, customer)
		}
		if t >= cashier3.timeFinished {
			customer := nextCustomer(&customerArrivals, t)
			cashier3.timeFinished = t + customer[1]
			//fmt.Println(cashier3, t, customer)
		}
		t++
	}
	cashiers := []cashier{cashier1, cashier2, cashier3}
	return findMax(cashiers)
}