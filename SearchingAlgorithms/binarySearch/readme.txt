Note:
it gets used when you have a sorted array (it's essential that the array be sorted)
consider an array by the size of N elements.
at first you have a starting point which is 0, an end point which is N - 1 and a middle point which is (start +‌ end ) / 2.
and you have a target element, instead of going through all of the element one by one,
you take the middle point element and cut your array in two halves.
first check if the middle element is your target or not,
if it is, well then return it's index, if it's not you have to choose which side you're going to look, left or right?
you check if the middle one is bigger or smaller than your target element.
since your array is sorted, if your target is smaller than your middle one then you have to look in the left side,
if your target is bigger than your middle one then you have to look in the right side.
so you ommited half of the array +‌ the middle one.
if your target was smaller than middle one, the starting point remain the same, but you change the endpoint to:
end = middle - 1
this makes your end point one element before the middle one.
if your target was bigger than middle one, the end point remain the same, but you change the start point to:
start = middle + 1
this makes your star point one element after the middle one.
keep doing this while your start point is <= than your end point.

O(Log N)