# Go

#### *How to run a go file ?*

``` go build <filename> ``` : Used to compile and build the executable file which could be run separately

``` go run <filename> ``` : Used to compile and run

#### *Fundamentals*

1. ``` package <name> ``` : Package is used to group the set of files under a single umbrella
    
    * There are 2 types of packages : Executable and Reusable
    * Executable package has ``` <name> ``` as '*main*'

2. ``` import "<package_name>" ``` : gives our code access to the code inside *<package_name>*

3. ``` 
   func <func_name>() <return_type> {

        // code 
        
   } 
   ```

   Function consisting of some code

   Note: Functions returning two values can have ``` (<return_type1>, <return_type2>) ``` in the method signature. Similar thing can be followed for function returning multiple values

4. ``` var message string = "Hi Kishan!" ``` \
   ``` message := "Hi Kishan!" ```
   
   are equivalent variable declaration , intializations

   Note : ``` := ``` should only be used when we are first initializing a variable

5. *Slices* : Slices are arrays of elements containing same data type but can be of varying length 

    ``` cars := []string { "audi" , "bmw"}   // Slice declaration and initilization ``` \
    ``` cars = append(cars, "toyota")        // Adding elements into slice``` 

    * *Iterating the slices*
    ``` 
    for index,car := range cars {
        fmt.Println(index,car)
    }
    ```  

    Note : 1. If we don't want ``` index ``` to be used inside loop we can replace it with *undescore* \
           $~~~~~~~~~$ 2. If we don't want the element of the slice directly we can skip using ```car``` in the loop signature

6. #### *Go is not an OOP Language* . 
    We can make use of Custom Type declarations and Receiver functions to achieve the complex functionalities

7. *Custom Type Declaration*

    ``` type cars []string ``` : Declaring a type cars which is a slice of string

8. *Receiver Functions*

    ```
    func (c cars) print() {
        for _, car := c {
            fmt.Println(car)
        }
    }

    // Receiever function for type 'cars'
    ``` 

    * Receiver functions can be used by the var's of a particular type to be called
    * The reciever argument is by conventionally a single / double letter character(s) signifying that particular type

9. Go is a pass-by-value language

10. ``` &variable ``` : Gives the memory address of the variable \
    ``` *pointer ``` : Gives the value at the memory address of the pointer

    ``` *type ``` : Indicates that we deal with the pointer of that type \
    ``` *varibale ``` : Indicates the value

    Address to value : ``` *address ``` \
    Value to address : ``` &value ```

11. Reference Types : slices , maps , channels , functions \
    Value Types : int , bool , string , struct

     *How Slice works without pointers as refernce type ?* \
         Slice maintians a reference to the underlying array, so although when we try to manipulate the slice according to GO ( which is a pass by value language) it creates another copy but here the pointer to underlying array is same . This is a catch which makes it a reference type


