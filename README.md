# Monkey language interpreter

This is Monkey language interpreter written in Go. This interpreter is based on book "Writing An Interpreter In Go" by Thorsten Ball

## Features
* Literals: integer, boolean, string, arrays, hash tables
* let statements
* Conditions (if, else)
* Functions
* Loops
* Operators: +, -, /, *, <, >, ==
* Index operator: getting and setting values
* Built-in functions: print, exit, len, push

## Examples
### Map and filter functions
```js
let map = fn(arr, mapper) {
  let newArr = [];
  for (let i = 0; i < len(arr); i = i + 1) {
    newArr = push(newArr, mapper(arr[i]));
  }
  return newArr;
}

let filter = fn(arr, filterFn) {
  let newArr = [];
  for (let i = 0; i < len(arr); i = i + 1) {
    if (!filterFn(arr[i])) {
      continue;
    } else {
      newArr = push(newArr, arr[i]);
    }
  }
  return newArr;
}

let arrays = [ [1, 2, 3], [3, 4], [], [3, 4, 1, 4] ]
let lens = map(arrays, fn(arr) { return len(arr); })
print(lens)  // [ 3, 2, 0, 4 ]

let filtered = filter(arrays, fn(arr) { return arr[len(arr) - 1] == 4 })
print(filtered)  // [ [3, 4], [3, 4, 1, 4] ]  
```

### Closures
```
let newAdder = fn(x) {
  return fn(num) {
    return num + x;
  ];
};

let addTwo = newAdder(2);
let addThree = newAdder(3);

print(addTwo(2))  // 4
print(addThree(10)) // 13
```
