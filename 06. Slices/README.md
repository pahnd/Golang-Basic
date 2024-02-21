Slice:
Has a dynamic length (it can shrink or grow);
The length of a slice is not part of its type and it belongs to runtime;
An uninirialized slice is equal to nil (its zero values is nil).
Both a slice and an array can contain only the same type of elements
We can create a keyed slice like a keyed array;


Slice's Backing (Underlying) Array
When createing a slice, behind the scenes Go creates a hidden array called Backing Array.
The backing array stores the elements, not the slice.
Go implements a slice as a data structure called slice header.


Slice Header contains 3 fields:
1. the address of the backing array (pointer).
2. the length of the slice. len() returns it.
3. the capacity of the slice. cap() return it.


Slice Header is the runtime representation of a slice.
A nil slice doesn't have backing array.