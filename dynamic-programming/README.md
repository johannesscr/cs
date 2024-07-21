# Dynamic Programming

Dynamic programming is an optimisation technique is an algorithm that uses
something called a *cache*. A *cache* is an in-memory store of recent values or 
observations made in the algorithm. It can be thought of as a piece of paper
to keep note of recent values used in the algorithm.

**Dynamic Programming** is a method of solving problems by breaking it down
into a collection of sub-problems. Solving each of those sub-problems just once,
in case you need to solve it again, then you can just retrieve the solution from
memory.

## Memoization

**Caching** is a way to store values so that you can use them later on. 
**Memoization** is synonymous to caching, to speed up programs/algorithms and 
to keep a collection of values that are easily accessible. Memoization is
storing the return value for its given input values. The combined set of input
values are the for the key or hash to retrieve the output values.

Determining whether dynamic programming and memoization is applicable:
1. Can the problem be divided into sub problems?
2. Recursive solution?
3. Are there repetitive sub problems?
4. Memoize sub problems?
5. \*Demand a raise from your boss!

Some Leetcode questions
1. [house robber](https://leetcode.com/problems/house-robber/)
2. [best time to buy and sell stocks](https://leetcode.com/problems/best-time-to-buy-and-sell-stock/)
3. [climbing stairs](https://leetcode.com/problems/climbing-stairs/)
