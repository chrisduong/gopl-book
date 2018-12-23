# Branch `mydev` to add exercises in the book

## Resources been consulted

[torbiak/gopl](https://github.com/torbiak/gopl)

[gopl-solutions](https://github.com/vinceyuan/gopl-solutions)

## Chapter 6

### Bit Vector (aka. Bit Array)

**SEE:** 

- [https://wiki.python.org/moin/BitArrays](https://wiki.python.org/moin/BitArrays).
  
  - Whatever they are called, these useful objects are often the most compact way to store data. If you can depict your data as boolean values, and can correlate each value with a unique integer, a bit array is a natural choice.

- [Bit Shifting](https://python-reference.readthedocs.io/en/latest/docs/operators/bitwise_left_shift.html)

- [Bit Manipulation](https://wiki.python.org/moin/BitManipulation)

- [Bit Clear]:
  
// Use bit clear AND NOT &^ to get the bits that are in source AND NOT dest (order matters)
// 26     = 00011010
// 6      = 00000110
// turn (negated)     11111001
// => 26 &^ 6 = 00011000 