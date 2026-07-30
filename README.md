# Dice Roller CLI (Go)

A simple, interactive Command Line Interface (CLI) application written in Go that simulates rolling a customizable die a user-specified number of times.

## Features

- **Customizable Die Sides**: Choose any number of sides (e.g., 6-sided, 20-sided die).
- **Multiple Rolls**: Specify how many times the die should be rolled in a single run.
- **Input Validation**: Keeps asking until the user inputs a valid integer.
- **Modern Go Standard**: Built using Go's modern `math/rand/v2` library and range-over-int loops (`range repetir`).

## Requirements

- **Go 1.22** or higher (required for `math/rand/v2` and `range` over integer features).

## Quick Start

Run the application directly using `go run`:

```bash
go run main.go
