## Introduction


### How to define a var in Golang?


In Go, defining variables and constants is straightforward, but the syntax you choose depends on the context—specifically, whether you are writing code inside or outside a function. 

#### **1. Defining Variables**

Go provides a few different ways to declare variables. 

**The `var` Keyword (Standard Declaration)**
You can use the `var` keyword followed by the variable name and its type. This method can be used both inside and outside of functions (at the package level).

```go
// Explicitly typing the variable
var age int = 25
var name string = "Alice"

// Declaring without initializing (Go assigns a "zero value" like 0 or "")
var score int 
```

**Type Inference**
If you assign a value right away, Go is smart enough to infer the type. You can leave the type out completely:

```go
var age = 25          // Go knows this is an int
var name = "Alice"    // Go knows this is a string
```

**Short Variable Declaration (`:=`)**
This is the most common way to declare variables **inside a function**. It drops the `var` keyword entirely and infers the type. 
*Note: You cannot use `:=` outside of a function.*

```go
func main() {
    age := 25
    name := "Alice"
}
```

**Multiple Variable Declaration**
You can declare multiple variables on a single line or in a block for cleaner code:

```go
// Single line
var x, y, z int = 1, 2, 3

// Block declaration (great for grouping package-level variables)
var (
    isActive bool   = true
    maxUsers int    = 100
    server   string = "localhost"
)
```

#### **2. Defining Constants**

Constants are declared using the `const` keyword. Unlike variables, their values cannot be changed once they are set. They must be evaluated at compile-time, meaning you can't assign the result of a runtime function to a constant.

**Standard Constant Declaration**
Like variables, you can either explicitly type them or let Go infer the type.

```go
const pi = 3.14159
const greeting = "Hello, World!"

// Typed constant
const maxConnections int = 50
```

**Multiple Constant Declaration**
Just like with variables, you can group constants together in a block:

```go
const (
    StatusOk       = 200
    StatusNotFound = 404
    ServerName     = "MyGoServer"
)
```

**Pro-Tip: Using `iota` for Enumerated Constants**
Go has a special identifier called `iota` that is used to create incrementally numbered constants. It resets to `0` whenever the word `const` appears.

```go
const (
    Sunday = iota // 0
    Monday        // 1
    Tuesday       // 2
    Wednesday     // 3
)
```




### **Typing: Static Typing vs Dynamic typing**


#### **1. Dynamic typing**

Type checking occurs at runtime. Values have types, but variables not. The same variable can hold a numer on line 1- and a tring on line 20. Pyhton and JS are the most pupular program languages below a example in JS:

Here we have no error
```js
var age = 10
console.log(age)
age = "matheus"
console.log(age)
```

#### **2. Static typing**

Type checking occurs at compile time. Variables are bound to a specific type. If a variable start as Integer . it stays an Integer

Here we have an error
```go
import (
	"fmt"
)

func main() {
	age := 10
	fmt.Println(age)
	age = "matheus"
	fmt.Println(age)
}
//./prog.go:12:8: cannot use "matheus" (untyped string constant) as int value in assignment
//Go build failed.
```

Now to fix you have to create a new var to storage the new value
```go
import (
	"fmt"
)

func main() {
	age := 10
	fmt.Println(age)
	name := "matheus"
	fmt.Println(name)
}

```

### **Typing: Weak Typing vs Strong typing**

In our days we have different kind of progamn language, one kind isWeak Typing which does the conversion under the hoode, and the Strong typing that means that the language have to be explicit convert to not raise any error.

#### **1 Weak Typing**

So, what are the upsides of using weak typing, and why does it exist? We can focus solely on JS, when dealing with webpages, we often receive a lot of unexpected data from the server. To avoid breaking the web for silly reasons, JS accepts 'weird' operations as a trade-off to keep running. The downside is that errors can be suppressed. See the examples below.

Examples

```js
"3" + 1     // Result: "31" (It converted the 1 to a string and concatenated)
"3" - 1     // Result: 2    (It converted the "3" to a number and subtracted)
"3" + 1 - 1 // Result: 310  ("3"+1 becomes "31". "31" - 1 becomes 310... wait, what? "31" - 1 is 30! I messed up the joke, it's 30!)
```

#### **2 String Typing**

Strong typing requires that you only perform operations between the same data types, which makes total sense for avoiding bugs. If you try to add a string to a number, it won't be accepted. Even Python, which is very flexible, doesn't allow adding an integer to a string. I’ve included two examples using Golang below


This will fail because we cannot sum up string with number in Golang
```go
func main() {
	numberString := "3"
	number := 1
	fmt.Println(numberString + number)

    //invalid operation: numberString + number (mismatched types string and int)
    //Go build failed.
}
```

How to fix it?

We have to convert before try sum up those values
```go
import (
	"fmt"
	"strconv"
)

func main() {
	stringNumber := "3"
	number := 1

	convValue, err := strconv.Atoi(stringNumber)
	if err != nil {
		panic(err)
	}
	fmt.Println(convValue + number)
}
```





