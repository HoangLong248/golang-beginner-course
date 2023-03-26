## Format Ouput - printf
    `fmt.Printf("Some text with a variable %s, myVariable)`
- It takes a template string that contains the text that need to be formatted
- Plus some annotation verbs ( or placeholder ) that tells the fmt functions how to format the variable passed in

## Data Types
- In any programming languages you have mutiple data types
- Difference: which data types do they support exactly?
    + Strings
    + Integers
    + Booleans
    + Maps 
    + Arrays
    + ...
- Each data type can do different things and behave differently.
- Go is a statically typed language => You need to tell Go Compiler, the data type when declaring the variable
<br/>:thumbsup: **More Robus, reduces the likelihood of errors.**
<br/>:thumbsup: **Helps developers to catch type mismatches sooner (at compiler time)**
<br/>:thumbsup: **Discover mistakes at compile time, NOT at runtime**

### "Syntatic Sugar" in Programming
    `conferenceName := "Go Conference" Syntatic Sugar`
- A term to describe a feature in a language that let's you do smth more easily
- Makes the language "sweeter" for human use
- But doesn't add any new functionality that it didn't already have

## Geting The User Input
### "fmt" package
- Different functions for:
    + Formatted Input and Ouput (I/O)
    + Print Messages
    + Collect User Input
    + Write into a File
### What is a Pointer?
- Where does the values get stored?
- A pointer is a variable that points to the memory address of another variable
- A special variable
:thumbsup: `Scan(%Username)` => `Scan(0xc00000e0a8)`

## Array & Slices in Go
- Data structures to store collection of elements in a single variable.
- We want to store the entered user data in some kind of a list
    - who is attending the conference?
    - who booked the tickets?
### Arrays in Go
- Fixed size ( size = how many elements the array can hold )
    `var variable_name[size]variable_type`
    `var bookings = [50]string{}`
- Only the same data type can be stored
- Adding and accessing elements by their index (position)
### Slices in Go
- Slice in an **abstraction of an Array**
- Slice are more flexible and powerful:
    - **variable-length**
    - get an sub-array of its own
- Slices are also **index-based** and have a size, but is **resize when needed**
### append
- Adds the element(s) **at the end of the slices**
- **Grows the slice if a greater capacity is needed** and returns the updated slice value.

## Loops
- In general, languages provide various control structures to control the applications flow
- A loop statement allows us to execute code multiple times, in a loop.
- You only have the **"for loop"** in Go.
### ranges
- Range interates over elements for different data structures (so not only arrays and slices)
- For arrays and slices, range provides the index and value for each element.
### strings.Fields()
- Splits the string with white space as separator.
- And returns a slice with the split elements.
### Blank Identifier
- To ignore a variable you don't want to use.

## Conditions (If/else) and Boolean Data Type
```
if condition {
    // code to be executed if condition is true
} else if condition {
    // code to be executed if condition is true
} 
else {
    // code to be executed if condition is false
}
```
### Conditions
- Expressions that evaluate to either **true** or **false**

| Operator | Description |
|---|---|
| == | Compares two things |
| != | Not equal |
| < | Less than |
| <= | Less than or equal to |
| > | Greater than |
| >= | Greater than or equal to |
### "break" statement
- **Terminates the for loop**
- And continues with the code right after the for loop (in our case we don't have any code)

### "continue" statement
- Causes loop to skip the remainder of its body.
- And immediately retesting its condition (in our infinite loop our condition is always true)

### End of a Loop
- A loop continues as long as a condition is true

## User Input Validation
- We need to expect that and **make sure our application can handle bad user input**
- We always need to validate user input.

### Logical Operator -- &&
- Called logical **AND** operator
- **Both conditions must be true**, for the whole expression to be true

### Logical Operator -- ||
- Called logical **OR** operator
- **Any of the conditions must be true**, for the whole expression to be true

## Logical Operator -- !
- Called logical **NOT** operator
- Reverses the boolean value.

## Switch Statement
- Allows a variable to be tested for equality against a list of values
- Default handles the case, if no match is found.

## Functions
- Encapsulated code into own container (= function). Which logically belongs together!
- Like variable name, you should give a function a **descriptive name.**
- **Call the function by its name**, whenever you want to execute this block of code.
- Every program has a least on function, which is the **main()** function
- Function is only executed, when "called"!.
- You can call a function as many times you want.
- So a function is also used to reduce code duplication.

### Parameters
- Information can be passed into functions as parameter.
- Parameters are also called arguments.

<br/>:thumbsup: => Cleaner
<br/>:thumbsup: => More Descriptive

### Return Values
- A function can return data as result.
- So a function can take an input and return an output.

### :boom: Returning multiple Values
- A Go function can return multiple values

### Package Level Variables
- Defined at the top **outside all functions**

### :star: Best Practice:
- Define variable as "local" as possible

### More Use Cases for Functions
- Group logic that belongs together.
- Reuse logic and so reducing duplication of code.

## Packages in Go
- Go programs are organized into packages
- A package is **a collection of Go files**

### Scope: Package level
- Again: Variables and Functions defined outside any function, can be accessed in all other files **within the same package**

### Multiple Packages in your app
- Organize your app
- logically group your code.

### Exporting a variable
- Make it available for all packages in the app
- **Capitalize first letter**

## 3 Levels of Scope

- Local
- Package
- Global

### Local
- Declaration **within function**
- Declaration **within block** (e.g for, if-else)

### Package
- Declaration **outside all function**

### Global
- Declaration **outside all functions & uppercase first letter**

## Maps
- Maps unique keys to values.
- You can retrieve the value by using its key value.
- All keys have the same data type
- All values have the same data type
- We can not mix data types

## Struct Data Type
- Stands for "Structure"
- Can hold mixed data types

### "type" statement - Custom types
- The **type keyword creates a new type**, with the name you specify
- In fact, you could also create a type based on every other data type like int, string, etc
```
type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}
```
### Defining a structure
- Mixed data type
- Defining a structure (which fields) of the User Type
- It's like a **lightweight class**, which e.g. doesn't support inheritance

## Goroutines - Concurrency in Go

### "time" - functionality for time
- The **sleep** function stops or blocks the current "thread" (goroutine) execution for the defined duration.

- **Handle this blocking code with** Goroutines
- **By default: Sequential** code execution
- Processing 1 task or 1 code after another.
- **Concurrency** to make our program more efficient

### Single Thread

![Screenshot](image/single_thread.png)

- **Step by step** execution of the code

### "go" keyword

- "go..." starts a new goroutines
- A goroutine is a **lightweight thread** managed by the Go runtime

<br/> :thumbsup: Stays Responsive
<br/> :thumbsup: Faster

### Synchronzing the Goroutines

- By default, the main goroutine does **NOT** wait for other goroutines

### Waitgroup

- Waits for the launched gorountine to finish
- Package "sync" provides basic synchronization functionality
- **Add**: Sets the number of goroutines to wait for (**increase** the counter by the provided number)
- **Wait**: Blocks until the WaitGroup counter is 0
- **Done**: Decrements the WaitGroup counter by 1. So this is called by the goroutine to indicate that it's fininshed.

### Goroutine

- Go is using, what's called a **"Green thread"**
- **Abstraction** of an actual thread.
- Managed by the go runtime, we are only interacting with these high level goroutines.
- Cheaper && lighweight.
- You can **run** hundreds of thousands or **millions goroutines without affecting the performance**

### Channels
- **Built-in** functionality for goroutines to talk with one another.