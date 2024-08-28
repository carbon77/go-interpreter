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
    if (filterFn(arr[i])) {
      newArr = push(newArr, arr[i]);
    }
  }
  return newArr;
}

let arrays = [ [1, 2, 3], [3, 4], [], [3, 4, 1, 4] ]
let lens = map(arrays, fn(arr) { return len(arr); })
print(1)
print(lens) 
print(1)

let filtered = filter(arrays, fn(arr) { 
  print("Arr: ", arr[0])
  return true;
})
print(filtered)
print(123)
