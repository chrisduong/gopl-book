# CHAPTER 8: Goroutines and Channels

## Keywords to remember:

- Communicating Sequential processes (**CSP**)
- *goroutine*: each concurrently executing activity. it is *similar to a OS thread*. But they are actually different.
- **Minor keywords**: *main goroutine*

## 8.1 Goroutines

## 8.2 Example: Concurrent Clock Server

### **First example:** *sequential clock server*

The listener's *Accept* method blocks until an incoming connection request is made, then returns a *net.Conn* object representing the connection

- ~> mean for sequential work, the Listener wait for the next connection being made, to do the specific work. It block until the connection is dropped
- The *for* loop will make the Server keep checking for new connection. Instead of exiting after the first connection is done

The *handleConn* function handles one complete client connection. The *for* loop will make it keep writing Times to client

The second client must wait until the first client is finished because the server is *sequential*

### Second example: concurrent Clock server

Added Goroutine to handle multiple connections concurrently.

```go
go handleConn(conn)
```

## 8.3 Example: Concurrent Echo Server

The clock server used one goroutine per connection. Echo server would use *multiple goroutines per connection*.

**Notice** that the third shout from the client is not dealt with until the second shout has petered out, which is not very realistic.

Add a goroutine in the func *handleConn* in the server side.

```go
go echo(c, Input.Text(), 1*time.Second)
```

The arguments to the function started by *go* (*MEAN echo(...)*) are evaluated when the go statement itself is executed; thus `input.Text()` is evaluated in the **main goroutine** (*overlapping happened!*)

- ~> inside the *subroutine handleConn* it was sequential if not using Goroutine.

We had to <u>consider carefully that it is safe to call methods of net.Conn concurrently</u>

## 8.4 Channels

If goroutines are the activities of a concurrent Go program, *channels* are the connections between them.

A channel is a communication mechanism that lets one goroutine send values to another goroutine.

```go
ch := make(chan int)
```

A channel is a *reference* to the data structure created by *make*. When we copy a channel or pass one as an argument to a function, we are copying a reference, so caller and callee refer to the *same data structure*.

Kind of channels: Unbuffered channel (*synchronous channels*) and Buffered channel with *capacity*

### 8.4.1 Unbuffered Channels

A send operation on an unbuffered channel blocks the sending goroutine until another goroutine executes a corresponding receive on the same channel.

The benefit is stronger synchronization guarantees.

### 8.4.2 Pipelines

### 8.4.3 Unidirectional Channel Types

### 8.4.4 Bufferred Channels

The assembly line metaphor is a useful one for channels and goroutines.

**Beware of Deadlock** if not using the sufficient buffer capacity