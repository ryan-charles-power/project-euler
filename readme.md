# Project Euler

Project Euler is a popular website offering a collection of challenging, mathematical programming problems designed to be solved using computer programs and algorithmic thinking. Aimed at developers and math enthusiasts, these challenges require combining mathematical insights with programming skills to find efficient solutions, often within a "one-minute" CPU time constraint.

This repository contains my solutions to Project Euler problems implemented in Go and Python.  
Each runner includes high-precision timing to track execution performance.

## 🚀 How to Run Project Euler Solutions

### 🐹 Go
The Go runner uses the `time` package for millisecond-precision timing.  

#### How to Run: 
```
go run main.go <solution_number>
```
Example: `go run main.go 1`

### 🐍 Python
The Python runner uses the `time` module for monotonic, high-resolution timing.  

#### How to Run
```
python main.py <solution_number>
```
Example: `python main.py 1`

## 📂 Project Structure
```
.
├── go/
│   ├── utils/                  # Utilities for solutions
│   │   ├── isPrimeNumber.go
│   │   ├── reverseString.go
│   │   └── ...
│   ├── solution1.go            # Individual Go solution logic
│   ├── solution2.go
│   └── ...
├── python/
│   ├── solution1.py            # Individual Python solution logic
│   ├── solution2.py
│   └── ...
├── main.go                     # Go runner & timer
└── main.py                     # Python runner & timer
```

## 📊 Solution Progress
|          | Go |  Python  |
|:--------:|:--:|:--------:|
|Problem 1 | ✅ |    ✅   |
|Problem 2 | ✅ |    ✅   |
|Problem 3 | ✅ |    ✅   |
|Problem 4 | ✅ |    ✅   |
|Problem 5 | ✅ |    ✅   |
|Problem 6 | ✅ |    ✅   |
|Problem 7 | ✅ |    ✅   |
|Problem 11| ✅ |    ✅   |
|Problem 17| ✅ |    ❌   |
|Problem 44| ✅ |    ❌   |

```
✅ Completed  
❌ Not Started
```