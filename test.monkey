let map = fn(arr, mapper) {
  let newArr = []
  let i = 0

  for (i < len(arr)) {
    newArr = push(newArr, mapper(i, arr[i]));
    i = i + 1;
  }

  return newArr
}

let filter = fn(arr, condition) {
  let newArr = []
  let i = 0

  for (i < len(arr)) {
    if (condition(arr[i])) {
      newArr = push(newArr, arr[i]);
    } 
    i = i + 1;
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
