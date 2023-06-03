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



