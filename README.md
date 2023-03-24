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
:thumbsup: `Scan(%Username)` => `Scane(0xc00000e0a8)`