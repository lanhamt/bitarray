#bitarray
Implementation of bitarray data structure for go.

##Usage
Import by either adding to gopath for libraries or importing 
directly from github (specific instructions for import can be 
found here: https://golang.org/doc/code.html).


Initialize bitarray with given length (all bits initialized to 0). 
Supports getting bit at index in array and setting bit at index 
to either 0 or 1. 
```
ba := BitArray(10) // 0000000000
ba.SetBit(0, 1)    // 1000000000
ba.SetBit(3, 1)    // 1001000000
ba.GetBit(9)       // returns false
ba.GetBit(3)       // returns true
ba.GetLength()     // returns 10
```

##Testing
To test with included unit test, descend into directory and 
execute `go test`.

Note: test framework does not test out of bounds errors since 
golang's testing does not accomodate `panic`. 
