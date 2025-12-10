/**
 * @author Marco Cantusci
 * @verion 1.0.0
 * @date 2025-12-09
 * @fileoverview Calculator
 */

// print out options
console.log(
  "Which operation would you like to perform today? (Select by typing the letter in front of the operation.)\n",
  "A. Add\n",
  "B. Subtract\n",
  "C. Multiply\n",
  "D. Divide\n",
  "E. Absolute value\n",
  "F. Round\n",
  "G. Raise to an exponent\n",
  "H. Square root\n",
);

// input
// choosing operation
let choice: string = prompt("What operation do you want to choose: ") || ("\n");
choice = choice.toUpperCase();

// addition
if (choice == "A") {
  const num1String: string = prompt("Enter first number: ") || ("0");
  const num1Number: number = parseFloat(num1String);

  const num2String: string = prompt("Enter second number: ") || ("0");
  const num2Number: number = parseFloat(num2String);

  const result = num1Number + num2Number;
  console.log(`${num1Number} + ${num2Number} = ${result}`);
} else if (choice == "B") { // subtraction
  const num1String: string = prompt("Enter first number: ") || ("0");
  const num1Number: number = parseFloat(num1String);

  const num2String: string = prompt("Enter second number: ") || ("0");
  const num2Number: number = parseFloat(num2String);

  const result = num1Number - num2Number;
  console.log(`${num1Number} - ${num2Number} = ${result}`);
} else if (choice == "C") { // multiplication
  const num1String: string = prompt("Enter first number: ") || ("0");
  const num1Number: number = parseFloat(num1String);

  const num2String: string = prompt("Enter second number: ") || ("0");
  const num2Number: number = parseFloat(num2String);

  const result = num1Number * num2Number;
  console.log(`${num1Number} * ${num2Number} = ${result}`);
} else if (choice == "D") { // division
  const num1String: string = prompt("Enter dividend ") || ("0");
  const num1Number: number = parseFloat(num1String);

  const num2String: string = prompt("Enter divisor: ") || ("0");
  const num2Number: number = parseFloat(num2String);

  const result = num1Number / num2Number;
  console.log(`${num1Number} / ${num2Number} = ${result}`);
} else if (choice == "E") { // Absolute value
  const num1String: string = prompt("Enter number: ") || ("0");
  const num1Number: number = parseFloat(num1String);

  const result = Math.abs(num1Number);
  console.log(`The absolute value of ${num1Number} = ${result}\n`);
} else if (choice == "F") { // Round
  const num1String: string = prompt("Enter number: ") || ("0");
  const num1Number: number = parseFloat(num1String);

  const result = Math.round(num1Number);
  console.log(`The rounded value of ${num1Number} = ${result}\n`);
} else if (choice == "G") { // exponent
  const num1String: string = prompt("Enter base number: ") || ("0");
  const num1Number: number = parseFloat(num1String);

  const num2String: string = prompt("Enter exponent: ") || ("0");
  const num2Number: number = parseFloat(num2String);

  const result = Math.pow(num1Number, num2Number);
  console.log(`${num1Number} ^ ${num2Number} = ${result}`);
} else if (choice == "H") { // square root
  const num1String: string = prompt("Enter number: ") || ("0");
  const num1Number: number = parseFloat(num1String);

  const result = Math.sqrt(num1Number);
  console.log(`Square root of ${num1Number} = ${result}\n`);
} else {
  console.log("Invalid. Please choose between A-H.");
}

console.log("\nDone.");
