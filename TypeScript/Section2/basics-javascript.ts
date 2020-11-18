// declare variables 
const number1 = 5;
const number2 = 2.8;
const printResult = true; '
let results = null

function add(num1:number, num2:number, printBool:boolean){
    /** 
    * Basic function that adds 2 values 
    * In JavaScript if at least 1 of the values is 'string', then the function concatenates num1 and num2
    * If param has ":", then an error is returned when value set isn't of the correct type
    */
    if (typeof num1 !== "number" || typeof num2 !== "number"){
        /**
         * For JavaScript you need to check the data type. 
         * If invalid type returns error 
         */
        throw new Error('Incorrect input')
    }
    const result = num1 + num2; 
    if(printBool){
        console.log('Result is: ' + result);
    }
    else{
        return result;
    }
}

// set result of function to variabless
results = add(number1, number2, printResult);

// print result
console.log(results);

