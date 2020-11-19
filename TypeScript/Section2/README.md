# Core Types 
* [basics.ts](basics.ts) - covers `numbers`, `strings` and `booleans`. 

## JavaScript 
* `number`: Both floats and integers are number
    * By default all numbers are floats (ie 5 = 5.0 and 5 != 5.5)
* `string`: anything that's quoted ("", '', ``)
    * `` is considerd tempelate litterals 
* `boolean`: true or false 
    * 0 and none is `falsy` boolean value 
* `const` is a static value, whereas `let` is a changeable variable
* Sample Function using JavaScript  
```
function add(num1, num2){
    /** 
    * Basic function that adds 2 values 
    * In JavaScript if at least 1 of the values is 'string', then the function concatenates num1 and num2
    * For JavaScript you need to check the data type. 
    * If invalid type returns error 
    */
    if (typeof num1 !== "number" || typeof num2 !== "number"){
        throw new Error('Incorrect input')
    }
    return num1 + num2;  
}

```
* `object`: key-value pair ({key: value})
* `typeof ${value}` informs type of value (like `type(value)` in Python)
* `array`: A list of values ( [] )
    * can be specific type or `any` 

## TypeScript
* Declerations of parameters in function support type validaton
```
function add(num1:number, num2:number){
    /** 
    basic function that adds 2 values 
    in JavaScript if at least 1 of the values is 'string', then the function concatenates num1 and num2
    If param has ":", then an error is returned when value set isn't of the correct type
    */
    return num1 + num2;  
}
```
* TypeScript is staticly type 
* `tuple`: a fixed length/type `array` 
* `enum`: Enuumarted global constatn identifiers (`enum{new, old}`)
* `any`: Store any type of value 
**Key Difference:** JavaScript uses "dynamic types" (resolved at runtime), while TypeScript uses "static types" (set during development) 
* `union`: allowing for a combination of value types. 
```
function combine(input1:number | string, input2:number | string){
    /**
     * Add 2 values together. 
     *  If values both values are of type number do arithmetics
     *  If at least on of the values is of type string, concatinate strings
     */ 
    let result:any = null; 
    if(typeof input1 === 'number' && typeof input2 === 'number'){
        result = input1 + input2; 
    }
    else{
        result = input1.toLocaleString() + input2.toLocaleString();
    }
    return result 
}
``` 
* `litteral`: specifying not only the type but also the actual value (ex. `const value = 2.5`)
* `type`: allows to declare our on type 
```
type number_string = number | string; 
function combine(input1:number_string, input2:number_string){
    /**
     * Add 2 values together. 
     *  If values both values are of type number do arithmetics
     *  If at least on of the values is of type string, concatinate strings
     */ 
    let result:any = null; 
    if(typeof input1 === 'number' && typeof input2 === 'number'){
        result = input1 + input2; 
    }
    else{
        result = input1.toLocaleString() + input2.toLocaleString();
    }
    return result 
}
