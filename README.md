# Byte Array to String Converter

A simple Go program that converts byte arrays into their string representation. It accepts input either as command-line arguments or through stdin.

## Installation

```bash
go install github.com/nicksherron/byteprint@latest
```

## Usage

The program can be used in two ways:

### 1. Command Line Arguments

Pass the byte array as command-line arguments:

```bash
byteprint 102,111,111,109,97,110,99,104,117
```

Or with spaces:

```bash
byteprint "102, 111, 111, 109, 97, 110, 99, 104, 117"
```

Both will output: `foomanchu`

### 2. Standard Input

You can also pipe or type numbers directly into stdin:

```bash
echo "102 111 111" | byteprint
```

Or type numbers manually (press Ctrl+D or Cmd+D when done):
```bash
byteprint
102
111
111
```
