# FMT Cheat Sheet

## General
- %v: default format for value
- %T: type of value
- %d: decimal integer
- %s: string
- %%: percent sign = to add a percent sign to the output

## Booleans
- %t: true or false = see the boolean value

## Integers
- %b: binary integer (base 2)
- %o: octal integer (base 8)
- %d: decimal integer (base 10)
- %x: hexadecimal integer (base 16, lowercase)
<!-- - %X: hexadecimal integer (base 16, uppercase) -->

## Floating Point
- %e: scientific notation (e.g. 1.234456e+78)
- %f: decimal point and fraction, no exponent (e.g. 123.456)
- %g: compact representation of %e or %f, for large exponents (e.g. 123.456 or 1.23456e+78)

## Strings
- %s: string (default format)
- %q: double-quoted string with Go syntax (e.g. "Hello World")

## Width and Precision
- %f: default width, default precision (e.g. %10.2f = 10 characters wide, 2 decimal places)
- %9f: 9 characters wide, default precision (e.g. %9f = 9 characters wide, 6 decimal places)
- %10s: 10 characters wide (right-aligned)
- %.2f: default width, 2 decimal places
- %10.2f: default width, 2 decimal places (e.g. %10.2f = 10 characters wide, 2 decimal places)
- %9.f: 9 characters wide, no decimal places (e.g. %9.f = 9 characters wide, 0 decimal places)

## Padding and Alignment
- %09d: 9 characters wide, padded with zeros (e.g. %09d = 9 characters wide, padded with zeros)
- %-4d: 4 characters wide, left-aligned (e.g. %-4d = 4 characters wide, left-aligned)

## Methods
- Sprintf(): format and return a string
- Printf(): format and print to standard output
