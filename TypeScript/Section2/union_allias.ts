// An allias for a union of data-types
type number_string = number | string; 

function combine(input1:number | string, input2:number_string){
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


let results:any;
results = combine(5, 3); 
if(results === add(5, 3)){
    console.log(results)
}
results = combine(5, 'a')
console.log(results)
