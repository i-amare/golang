# Golang Deep DIve🧪🐹

This repo is **me exploring Go** from *basic syntax* to *more complex topics & features*—mostly by building small, focused programs and mini-projects.

It loosely follows (and riffs on) this Udemy course:

- [Go - The Complete Guide (Udemy)](https://www.udemy.com/course/go-the-complete-guide/)

If you’re browsing: expect **many small modules**, lots of `main.go` files, and the occasional “why is this pointer doing that?” moment.

## Repo structure 🗂️

Folders are organized like course “chapters”:

- `01.getting-started/` – hello world
- `02.go-essentials/` – console apps, input/output, control flow, basic math
- `03.go-packages/` – splitting code into packages + basic module usage
- `04.pointers/` – pointers & dereferencing
- `05.structs/` – structs, methods, constructors, embedding
- `06.interfaces/` – interfaces + small persistence examples
- `07.data-structures/` – arrays, slices, maps
- `08.functions/` – recursion, variadic funcs, higher-order funcs/closures, generics
- `09.price-calculator/` – file IO + JSON output from computed results
- `10.concurrency/` – goroutines, channels, and a small concurrent “job” processor
- `rest-api/` – a Gin REST API backed by SQLite

Most of these folders are **their own Go module** (they contain a `go.mod`), so you run them from inside the folder you care about.

## What I’ve covered (so far) ✅

Based on the code in this repo, here’s what I’ve implemented while working through the course:

- **Basics**
  - `package main`, `func main()`, imports, `fmt.Print*`
  - Variables and types (e.g. `int`, `float64`, `string`)
  - Reading user input with `fmt.Scan*`
- **Control flow**
  - `for` loops, `switch`, `continue`, early `return`
- **Functions**
  - Multiple return values
  - Recursion (`08.functions/factorial/`)
  - Variadic functions + spread operator (`numbers...`) (`08.functions/sumup/`)
  - Higher-order functions + closures (`08.functions/number-transforms/`)
  - **Generics** with type constraints (`08.functions/sumup/`)
- **Data structures**
  - Arrays vs slices, slicing, `append` and `...` (`07.data-structures/arrays/`)
  - Maps + iteration (`07.data-structures/maps/`)
- **Packages & modules**
  - Creating your own packages (e.g. `utils/`, `user/`, `todo/`, `note/`)
  - Importing local packages within a module (`03.go-packages/bank/`, `05.structs/`)
- **Pointers**
  - Taking addresses (`&x`), dereferencing (`*p`), pointer parameters (`04.pointers/`)
- **Structs**
  - Struct constructors (e.g. `NewUser`)
  - Methods with value vs pointer receivers (mutating state)
  - Embedding structs (`Admin` embedding `User`) (`05.structs/`)
- **Interfaces**
  - Defining an interface and accepting it as a parameter
  - Swapping implementations (file vs command input/output manager style) (`06.interfaces/`, `10.concurrency/concurrency-project/`)
- **File IO + JSON**
  - Reading/writing files (`os.ReadFile`, `os.WriteFile`, `os.Open` + `bufio.Scanner`)
  - JSON marshal/encode to files (`encoding/json`) (`06.interfaces/`, `09.price-calculator/`, `10.concurrency/concurrency-project/`)
- **REST API (Gin + SQLite)**
  - Routing + handlers with Gin (`GET/POST/PUT/DELETE`) (`rest-api/routes/`)
  - JSON binding/validation tags (`binding:"required"`)
  - SQLite setup via `database/sql` + `github.com/mattn/go-sqlite3` (`rest-api/db/`)
  - Simple vendor CRUD + a `.http` file for manual API calls (`rest-api/api-tests/vendor.http`)
- **Concurrency**
  - Starting goroutines (`go fn()`)
  - Signaling completion with channels (`chan bool`) (`10.concurrency/concurrency-example/`)
  - Fan-out work across multiple goroutines (per tax rate) (`10.concurrency/concurrency-project/`)

## Why this exists 😄

Because the best way for me to learn is to build lots of tiny things, break them, fix them, and slowly accumulate Go muscle memory.

