let books = [
  {
    "author": "Thorsten Ball",
    "title": "Writing An Interpreter In Go"
  },
  {
    "author": "Joshua Bloch",
    "title": "Effective Java"
  }
]

let printBook = fn(book) {
  print("Book[author=" + book["author"] + ", title=" + book["title"] + "]")
}

for (let i = 0; i < len(books); i = i + 1) {
  printBook(books[i])
}
