// declare variables 
var number1 = '5';
var number2 = 2.8;
function add(num1, num2) {
    /** s
    basic function that adds 2 values
    in JavaScript if at least 1 of the values is 'string', then the function concatenates num1 and num2
    If param has ":", then an error is returned when value set isn't of the correct type
    */
    return num1 + num2;
}
// set result of function to variabless
var result = add(number1, number2);
// print result
console.log(result);
