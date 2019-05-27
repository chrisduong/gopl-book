# CHAPTER 2: PROGRAM STRUCTURE

In Go, as in any other programming language, one builds large programs from a small set of basic constructs.

## 2.1 Names

A name befines with a letter or an underscore.

25 *keywords*

## 2.2 Declarations

A *declaration* names a program entity and specifies some or all of its properties

## 2.3 Variables

### 2.3.2 Pointers

A *pointer* value is the *address* of a variable.

**NOTE**: 

- not every value has an address, but every variable does. (HINT: value is taken care by Compiler)
- THINK *pointer* as the shortcut when you don't need to define an variable, OR when you only need a value

#### When to use pointer receiver or value receiver

[https://flaviocopes.com/golang-methods-receivers/](https://flaviocopes.com/golang-methods-receivers/)

##### **When to use pointer receiver**

- Modify the receiver
- Optimization: if the struct is very large

##### **When Value Receivers are better**

- Value receivers are concurrency safe.

Variable are sometimes described as *addressable* values.

The zero value for a pointer of any type is *nil*.

Pointers are comparable.

#### Resources

- [There is no pass-by-reference in Go](https://dave.cheney.net/2017/04/29/there-is-no-pass-by-reference-in-go)

## 2.5 Type Declarations

A type declaration defines a new ***named type*** that has the same underlying type as an existing type. The named type provides a way to separate different and perhaps incompatible uses of the underlying type <u>so that they can’t be mixed unintentionally</u>.

 “type name **underlying-type**”

A ***conversion*** from one type to another is allowed if both have the same underlying type, or if both are unnamed pointer types that point to variables of the same underlying type. These conversions change the type but not the representation of the value. (**NOTE:** loosely couple between type and its value ).

Conversions are also allowed between numeric types, and between string and some slice types, as we will see in the next chapter. <u>These conversions may change the representation of the value.</u>

## 2.6 Packages and Files

**Exercise 2.1 - page 159**

We would use episode to calculate the accurate of our conversions. See TEST file for more detail.

#### 2.6.1 Imports

**Exercise 2.2 - page 164**

- **Technique:** *type assertion*

