Note:
the expectation of a sorting algorithm is to take an unsorted array, and sort them in increasing or decreasing order.
we’re going to compare consecutive items, if they’re out of order we’re going to swap them.
The highest number will bubble its way to the right with each iteration.
N = 5
	     first                            second                           third                            forth
J=0 --------------- J=N-2         J=0 --------------- J=N-3       J=0 --------------- J=N-4        J=0 --------------- J=N-۵
5     4     3     2     1         4     3     2     1     5       3     2     1     4     5        2     1     3     4     5
4     5     3     2     1	      3     4     2     1     5       2     3     1     4     5        1     2     3     4     5
4     3     5     2     1         3     2     4     1     5	      2     1     3     4     5
4     3     2     5     1         3     2     1     4     5
4     3     2     1     5

We’re going to use nested loops, one for iteration.
So we need two loops, one for how many times we should do the process which is N – 1 times.
If we have 5 elements, if we sort 4 of them and they find their places there is only one number left which automatically is 
The smallest number.
and for j We went from j = 0 to j = N – 2 and each time we did a comparison between A[j] and its neighboring element which is A[j+1]
In each element one of them element will gets to its place. We now have a smaller array; we don’t make comparison to the 
End, because 5 is already in the end and if even we do the comparison the result is the same because 5 is largest.
And one for J who iterate through loops and make the comparison, but since after each iteration we will find a biggest 
Number, each time we do this iteration one less time.

O(N^2)