const js_button = document.querySelector("button");
const js_input1 = document.getElementById("num1");
const js_input2 = document.getElementById("num2");

function add_concate(js_num1, js_num2) {
  // Concat 2 values into a single string
  return js_num1 + js_num2;
}

function add_numeric(js_num1, js_num2) {
  // Add to numeric values
  return +js_num1 + +js_num2;
}

js_button.addEventListener("click", function() {
  console.log(add_concate(js_input1.value, js_input2.value));
  console.log(add_numeric(js_input1.value, js_input2.value));
});
