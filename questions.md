
# Questions

#### 1. How do we run the code in our project?

`go run <filename>`

#### 2. What does package main mean?

 - package == project == workspace
 - If you are working on one discreate app, right now that will constitute a
   package. 
 - Package can have many related files in it. Example Package Main, can have main.go
   support.go, helper.go etc. Each of these will have a `package main` as the first
   line. 

-  The reason we call the package main and not the folder name. 
   -  In go we have two types of packages 
      -  Executable and Reusable 
         -  Executable = Generates a file that we can run -> `main` is key word!
         -  Reusable = Code used as 'helpers'. Good place to put 
            -  reusable logic. (code dependencies, library of re-usable code)
- package main --> go build --> main.exe (windows) 
- package test --> go build --> no executable file gets generated
  - The word `main` allows us to make the package executable, by generating an 
  - executable file. 
- Additionally it also has a  function named `main`, which is ran automatically
- when the program runs. Entry point for execution when the program runs.
 
#### 3. What does 'import fmt' mean?

 - `format` - Standard libary similar to print in python
 - `import` command can be used to import other re-usable packages


#### 4. What's that 'func' thing?

  `func main (){}`
  - `func` - Tells go we're about to declare a function 
  - `main` - Sets the name of the function 
  - `()` - List of arguments to pass the function
  - `{}` -  Function body calling the function runs this code within the brackets
  
#### 5. How is the main.go file organized?

  1. Package declaration - `package main`
  2. Import other packages that we need - `import fmt`
  3. Declare functions, tell Go to do things

    func main() {
      fmt.Println("Hi there!")
    }

#### Types of DS in go to present an array 

-  Array - Fixed length list of things
-  Slice - An array that can grow or shrink == similar to a list in python
-  Array and slice must be defined with a data type. 
  - Every element in a slice must be of a same type.

#### Cards OOP Approach 

- Deck Class --> Which define an instance of Deck of cards.
                  - Each instance of the Deck class will have some properties.
                  - And defines some function to work on the deck of the cards.
  
In Go:

type deck []string - Tell Go we want to create an array of strings and add a bunch 
of functions specifically made to work with it.
 - Extend the slice and add some properties 

Functions with 'deck' as a 'receiver' - A function with a receiver is like a method 
- a function that belongs to an instance.

```
func (d deck) print(){
  for i, card:= range d {
    fmt.Println(i, card)
  }
}
```
Here d is reference to the actual copy of the variable we used of the string type.
which in our case was cards, so we could replace d with cards as well!!!! but by 
convention we have to keep it to 1-2 words. 

- The actual copy of the deck we'are working with is available 
  in the function as a variable called 'd'
- Every variable of type 'deck' can call this function on itself.

##### Slices

- Slices are zero-indexed
  
- fruits = "apple", "banana", "grape", "orange"
- fruits[0] -> "apple"
- fruits[3] -> "orange"


#### Byte Slice 

- string of character ~ to byte slice 
- String gets converted to ascii characters. 
- Tyoe conversation using go --> []byte("Hi there!) --> Type we want(Vaue we have)
- For converting deck to []byte
  - deck -> []string -> string -> []byte
    - Here string is concatiation of the slice of string
      - we can convert back to slice of byte []byte  


#### Error 

- byteSlice, err := os.ReadFile(filename)
- value of type 'error' if nothing went wrong it will have a value of 'nil'
  - 'nil' is similar to null in python. 
  
#### Testing

- newDeck
  - What do I care about for this newDeck function 
    - easy assertions 
      - Check the len 
      - First card check 
      - Last card check 

---

1. Create a deck
2. Save a file
3. File created
4. Attempt to load file 
5. Crash

- Clean up is required by you.
  
#### Delete a file 