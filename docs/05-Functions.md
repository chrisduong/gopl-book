- [Functions](#functions)
  - [5.1 Function Declarations](#51-function-declarations)
    - [HTML Node Spec](#html-node-spec)
  - [5.2 Recursion](#52-recursion)
    - [5.3 Multiple Return Values](#53-multiple-return-values)
    - [5.4 Errors](#54-errors)
    - [Error-Handling Strategies](#error-handling-strategies)
    - [5.4.2 End of File (EOF)](#542-end-of-file-eof)
  - [5.5 Function Values](#55-function-values)
    - [NEW: String format way - variable padded strings](#new-string-format-way---variable-padded-strings)
    - [Using a function value](#using-a-function-value)
  - [5.6 Anonymous Functions](#56-anonymous-functions)
    - [Example of Anonymous Function](#example-of-anonymous-function)
  - [5.7 Variadic Functions](#57-variadic-functions)
  - [Exercises](#exercises)
    - [ex5.10](#ex510)
    - [Exercise 5.15](#exercise-515)

# Functions

A function lets us wrap up a sequence of statements as a unit.

A function hides its implementation details from its users

## 5.1 Function Declarations

```go
func name(parameter-list) (result-list) {
    body
}
```

The type of a function is sometimes called its *signature*.

Arguments are passed by *value*

Keywords:

- `named result` and `bare return`.

Example code:

- gopl.io/ch5/findlinks1
- [DOM nodes diagram](https://www.w3schools.com/js/js_htmldom_navigation.asp)
- gopl.io/ch5/outline

### HTML Node Spec

[HTML Spec](https://html.spec.whatwg.org/multipage/syntax.html)

## 5.2 Recursion

Functions may be recursive, that is, they may call themselves, either directly or indirectly.

Recursion is a powerful technique for many problems, and of course it’s *essential for processing recursive data structures*

**NOTE:**

- It is IMPORTANT to know the Stack call (Caller and Callee) in recursion way. So you can know if it can cover all the cases you want, or traverse in the right way.

### 5.3 Multiple Return Values

- **bare return**

### 5.4 Errors

Some functions always succeed at their task.

For many other functions, even in a well-written program, success is not caused because it depends on factors beyond the programmer's control. For e.g, function that does I/O.

Reported using `exceptions`, not ordinary values.

- TODO: need to learn about about this page 377. For e.g. "exception mechanism of sorts",..

### Error-Handling Strategies

**First strategy:**

- SHOULD construct a _descriptive error message_.
- “error messages are frequently chained together, message strings should not be capitalized and newlines should be avoided”

**Second strategy:**

- “For errors that represent transient or unpredictable problems, it may make sense to retry the failed operation, possibly with a delay between tries”

**Third strategy:**

- “if progress is impossible, the caller can print the error and stop the program gracefully, but this course of action should generally be reserved for the main package of a program.”

**Fourth strategy:**

- “in some cases, it’s sufficient just to log the error and then continue.

**Fifth strategy:**

- In rare cases, we can safely ignore an error entirely.

“Get into the habit of considering errors after every function call, and when you deliberately ignore one, document your intention clearly.”

### 5.4.2 End of File (EOF)

## 5.5 Function Values

Functions are *first-class* values in Go: like other values, function values have *types*, and they may be *assigned to variables* or **passed** to or returned from functions. _A function value may be called like any other function_ (`f(3)`)

### NEW: String format way - variable padded strings

`fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)`

- `%*s` will be padded with a number (`depth*2`) of space (`""`).

### Using a function value

```go
func square(n int) int {
  return n * n
}

f := square
fmt.Println(f(3)) // "9"
```

- **nil value** of the function, ONLY happens when function is declared but not defined

- Function values let us parameterize our functions over not just data, but behavior too
  - Std Library - `strings.Map`. SEE: Strings package source for more.

- We can separate the logic for tree traversal from the logic for the action to be applied to each node, letting us _resuse the traversal with different actions_.

Code `gopl.io/ch5/outline2`

## 5.6 Anonymous Functions

A _function literal_ is written like a function declaration, but without a name following the func keyword.

More importantly, functions defined in this way have *access to the entire lexical environment*, so the inner function can refer to variables from the enclosing function, as this example shows

```go
func squares() func() int {
    var x int
    return func() int {
        x++
        return x * x
    }
}

func main() {
    f := squares()
    fmt.Println(f()) // "1"
    fmt.Println(f()) // "4"
    fmt.Println(f()) // "9"
    fmt.Println(f()) // "16"
}
```

The squares example demonstrates that function values are not just code but can have ***state*** (state is store inside variable `f`)

- These *hidden variables references*(`x`) are why we classify functions as reference types and why function values are not comparable (*state can be changed*)
- Function values like these are implemented using a technique called *closure*.

- XXX: *Here again we see an example where the lifetime of a variable is not determined by its scope*: the variable x exists after squares has returned within main, even though x is hidden inside f.

### [Example of Anonymous Function](ch5/example-of-anonymouse-function)

## 5.7 Variadic Functions

## Exercises

### ex5.10

Is the example of nondeterministic, as when you traverse (range) a Map, it will be randomly.

### Exercise 5.15

“Write variadic functions max and min, analogous to sum. What should these functions do when called with no arguments? Write variants that require at least one argument.”

