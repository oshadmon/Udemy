# Core Types 

## JavaScript 
* `number`: Both floats and integers are number 
* `string`: anything that's quoted ("", '', ``)
    * `` is considerd tempelate litterals 
* `boolean`: true or false 
    * 0 and none is `falsy` boolean value 

## TypeScript
* Declerations of parameters in function support type validaton
```
function add(num1:number, num2:number){
    /** s
    basic function that adds 2 values 
    in JavaScript if at least 1 of the values is 'string', then the function concatenates num1 and num2
    If param has ":", then an error is returned when value set isn't of the correct type
    */
    return num1 + num2;  
}
```