let counter = fn(x) {
  if (x > 100) {
    return true
  }
  return counter(x + 1)
}

counter(0)
