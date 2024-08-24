let map = fn(arr, mapper) {
  let newArr = []
  let wrapped = fn(wrappedArr, i) {
    if (i == len(arr)) {
      return wrappedArr; 
    }
    let wrappedArr = push(wrappedArr, mapper(arr[i]));
    return wrapped(wrappedArr, i + 1);
  }
  return wrapped(newArr, 0);
}

let forEach = fn(arr, f) {
  let wrapped = fn(i) {
    if (i == len(arr)) {
      return 0; 
    }
    f(i, arr[i])
    wrapped(i + 1)
  }
  wrapped(0);
}

let arrays = [[1, 2, 3], [3, 4], [1, 2], [1], []]
let lens = map(arrays, fn(arr) { len(arr) })

forEach(lens, fn(i, arrLen) {
  println(arrLen)
})
