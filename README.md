# Computor v1

## Overview

This project is part of the **42 School curriculum** and focuses on implementing a simple **polynomial equation solver** in Go.  
The program handles equations of **degree 2 or lower**, such as:

It parses the equation, reduces it to its **simplest form**, determines its **degree**, and computes its **solutions** (real or complex).

---

## Features

- Parses polynomial equations written in the form:
- Displays:
- **Reduced form** of the equation (all terms on one side).
- **Polynomial degree**.
- **Discriminant** information (for degree 2).
- **Solutions** (real or complex).
- Handles:
- Zero, positive, and negative coefficients.
- Infinite or no solutions (e.g., `42 * X^0 = 42 * X^0` → all real numbers).
- Implements **custom square root function** (Newton’s method), no `math.Sqrt` used.

---

## Examples

```bash
$ ./computor "5 * X^0 + 4 * X^1 - 9.3 * X^2 = 1 * X^0"
Reduced form: 4 * X^0 + 4 * X^1 - 9.3 * X^2 = 0
Polynomial degree: 2
Discriminant is strictly positive, the two solutions are:
0.905239
-0.475131

$ ./computor "5 * X^0 + 4 * X^1 = 4 * X^0"
Reduced form: 1 * X^0 + 4 * X^1 = 0
Polynomial degree: 1
The solution is:
-0.25

$ ./computor "8 * X^0 - 6 * X^1 + 0 * X^2 - 5.6 * X^3 = 3 * X^0"
Reduced form: 5 * X^0 - 6 * X^1 + 0 * X^2 - 5.6 * X^3 = 0
Polynomial degree: 3
The polynomial degree is strictly greater than 2, I can't solve.

```

## Quick Start with Docker

Pull the official Go image
docker pull golang:latest

Run your Go code directly from your project directory
docker run --rm -v $(pwd):/app -w /app golang go run main.go

Interactive Development Shell

For longer development sessions:

**docker run -it --rm -v $(pwd):/app -w /app golang bash**
