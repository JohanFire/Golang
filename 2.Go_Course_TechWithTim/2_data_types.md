
# Numbers
## Integers
### Unsigned Integers (no negative numbers)
- uint8 / byte: 0 to 255
- uint16: 0 to 65535
- uint32: 0 to 4294967295
- uint64: 0 to 18446744073709551615

### Signed Integers (negative and positive numbers)
- int8: -128 to 127
- int16: -32768 to 32767
- int32: -2147483648 to 2147483647
- int64: -9223372036854775808 to 9223372036854775807

### 3 Machine Dependent Types
- uint: 32 or 64 bits
    32 bit unsigned integer (0 to 4294967295) on 32-bit machines, 64 bit unsigned integer (0 to 18446744073709551615) on 64-bit machines
- int: same as uint, but signed
- uintptr: an unsigned integer to store the uniterpreted bits of a pointer value.
    unsigned integer the size of a pointer (32 or 64 bits)

## Floating Point Numbers
- float32: 32 bit floating point number (1.2e-38 to 3.4e+38)
- float64: 64 bit floating point number (2.2e-308 to 1.8e+308)

## Complex  (imaginary numbers)
- complex64: 32 bit real and imaginary parts (float32 + float32i)
- complex128: 64 bit real and imaginary parts (float64 + float64i)
- complex: same as complex128, but machine dependent (32 or 64 bits)

# Strings
- string: a sequence of bytes (UTF-8 encoded Unicode characters)
    - string literals: "Hello World" or `Hello World` ( ` backticks allow multi-line strings)
    - string concatenation: + operator
    - string length: len(string)
    - string indexing: string[index] (0 based index)
    - string slicing: string[start:end] (start inclusive, end exclusive)
    - string comparison: ==, !=, <, >, <=, >=
    - string formatting: fmt.Sprintf("Hello %s", name) or fmt.Printf("Hello %s", name)

# Booleans
- bool: true or false
    - boolean operators: && (and), || (or), ! (not)