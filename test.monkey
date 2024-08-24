let x = 3;

let newAdder = fn(x) {
  return fn(y) {
    return x + y
  }
}

let addThree = newAdder(3);
addThree(4);

let name = "Igor"
let greet = fn(name) {
  let greeting = "Hello ";
  return greeting + name + "!";
}

greet(name)

if (len(name) > 2) {
  return "OK"
}

