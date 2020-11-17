var ts_button = document.querySelector("button");
// force intention 
var ts_input1 = document.getElementById("num1");
var ts_input2 = document.getElementById("num2");
function add(ts_num1, ts_num2) {
    // Add 2 nunmbers using          
    return ts_num1 + ts_num2;
}
ts_button.addEventListener("click", function () {
    console.log(add(ts_input1.valueAsNumber, ts_input2.valueAsNumber));
});
