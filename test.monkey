let map = fn(arr, mapper) {
  let newArr = []

  for (let i = 0; i < len(arr); i = i + 1) {
    newArr = push(newArr, mapper(i, arr[i]));
  }

  return newArr
}

let filter = fn(arr, condition) {
  let newArr = []

  for (let i = 0; i < len(arr); i = i + 1) {
    if (condition(arr[i])) {
      newArr = push(newArr, arr[i]);
    } 
  }

  return newArr;
}

let arrays = [
  [1, 2, 3], 
  [3, 4], 
  [1, 2], 
  [1], 
  [],
  [3, 4, 5]
]
let lens = map(arrays, fn(idx, arr) { len(arr) })

print("Arrays: ", arrays)
print("Lengths of arrays: ", lens)
print("Filtered: ", filter(lens, fn(l) { l > 2 }))
