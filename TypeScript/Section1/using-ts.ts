const ts_button = document.querySelector("button");

// force intention 
const ts_input1 = document.getElementById("num1")! as HTMLInputElement;
const ts_input2 = document.getElementById("num2")! as HTMLInputElement;

function add(ts_num1:number, ts_num2:number) {
    // Add 2 nunmbers using          
    return ts_num1 + ts_num2;
}

ts_button.addEventListener("click", function() {
    console.log(add(ts_input1.valueAsNumber, ts_input2.valueAsNumber));
});