let x = 3;

let newAdder = fn(x) {
  return fn(y) {
    return x + y
  }
}

let addThree = newAdder(3);
addThree(4);
