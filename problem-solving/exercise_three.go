package problem_solving

/*
	#1: Sort10 schools around your house by distance
*/
/*
Insertion sort, fast on smaller datasets, O(1) space complexity.
Could also be nearly sorted.
*/

/*
	#2: eBay sorts listings by current Bid amount
*/
/*
Radix/Counting, a bid is a relatively small number, and those algorithms
are very good at sorting integers for a "small" fixed length.
*/

/*
	#3: Sort scores on ESPN
*/
/*
Quicksort: because scores could have decimals and random.
*/

/*
	#4: Massive database (can't fit into your memory) needs to sort through
	past year's user data.
*/
/*
Merge sort: won't need to sort in memory, so not too big a worry, but we want
the performance of O(n log n)
*/

/*
	#5: Almost sorted Udemy review data needs to update and add 2 new reviews.
*/
/*
Insertion sort: although large dataset, assumption that most of previous data
is already sorted.
*/

/*
	#6: Temperature records for the past 50 years in Canada
*/
/*
Radix/Counting sort: if no decimal places.
or
Quick sort: for in memory to save on space complexity.
*/

/*
	#7: Large username database need to be sorted. Data is very random.
*/
/*
Mergesort: if we have enough memory.
or
Quicksort: if not to worried about worst case.
*/

/*
	#8: You want to teach sorting algorithms for the first time.
*/
/*
Bubble sort and selection sort.
*/
