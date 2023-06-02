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
   func <func_name>() {

        // code 
        
   } 
   ```

   Function consisting of some code
