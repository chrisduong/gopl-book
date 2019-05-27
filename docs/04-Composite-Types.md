- [Chapter 4: Composite Types](#chapter-4-composite-types)
  - [4.1 Arrays](#41-arrays)
  - [4.2 Slices](#42-slices)
  - [4.4 Structs](#44-structs)
    - [4.4.1 Struct Literals](#441-struct-literals)

# Chapter 4: Composite Types

## 4.1 Arrays

An array is a fixed-length sequence of zero or more elements of a particular type. Because of their fixed length, arrays are rarely used directly in Go.

## 4.2 Slices

Slices (*underlying array*) represent variable-length sequences whose elements all have the same type.

[Slices from the ground up](https://dave.cheney.net/2018/07/12/slices-from-the-ground-up)

Go’s slice type differs from its array counterpart in two important ways:

- Slices do not have a fixed length. A slice’s length is not declared as part of its type, rather it is held within the slice itself and is recoverable with the built-in function len.2
- Assigning one slice variable to another does not make a copy of the slices contents. This is because a slice does not directly hold its contents. Instead a slice holds a pointer to its underlying array3 which holds the contents of the slice.
- The **slice header** value.

## 4.4 Structs

“A struct is an aggregate data type that groups together zero or more named values of arbitrary types as a single entity. Each value is called a field”

**NOTES:**

- The individual fields are accessed using *dot notation*
- Field order is significatnt to type identity.
- The name of a struct field is exported if it begins with a capital letter.
- An aggregate value cannot contain itself.
  - But S may declare a field of the pointer type *S, which lets us create recursive data structures like linked lists and trees. SEE: `ch4/treesort`

### 4.4.1 Struct Literals

First form: `type Point struct{ X, Y int }`

Second form (preferred): `p := Point{1, 2}`

## 4.5 JSON

JavaScript Object Notation (JSON) is standard notation of sending and receiving structured information.

Converting a Go data structure like *movies* to JSON is called *marshalling*

A *field tag* is a string of metadata associated at compile time with the field of a struct. 

