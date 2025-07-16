# Bit Manipulation

## Base

Base is a carry counting system with fixed digital symbols and rules. Each
carry counting system will have a base. When the base is $X$, we can it
base-$X$ numeral system, which means the number on each digit will be carried
over when it reaches $X$.

**Commonly used bases**

**Decimal**, in daily life, the most commonly used base system. It has 10
digits: 0, 1, 2, 3, 4, 5, 6, 7, 8 and 9.

**Binary (base 2)**, used in computer science. It has digits: 0 and 1.

**Octal (base 8)**, used in computer science. It has digits: 0, 1, 2, 3, 4,
5, 6 and 7.

**Hexadecimal (base 16)**, used in computer science. It has digits: 0, 1, 2,
3, 4, 5, 6, 7, 8 and 9 in addition there are A, B, C, D, E and F.

## Conversion between bases

### Non-decimal to decimal

To convert a non-decimal number to decimal, we need to add the weighted sum
of each digit. For example, to convert the octal number $720.5_{(8)}$ to
decimal.

$$  
720.5_{(8)} = 7 \times 8^2 + 2 \times 8^1 + 0 \times 8^0 + 5 \times 8^{-1}
= 464.625  
$$

## Decimal to non-decimal

To convert a decimal number to a base-$X$, we need to _convert_ the **integer
part** and the **fractional part**, separately.

To convert the integer part, we will integer divide it by $X$ until the integer
part reaches 0, and record the remainder each time. Traversing in reverse order
will give us the representation in base-$X$ system.

Example $464.625$ to base 8 (octal):
- $464 / 8 = 58$ and $464\ \%\ 8 = 0$.
- $58 / 8 = 7.25$ and $58\ \%\ 8 = 2$.
- $7 / 8 = 0.875$ and $7\ \%\ 8 = 7$.

Therefore, we have as we had seen above 720 for the integer part.

To convert the fractional part, we will multiply the fractional part of the
decimal number by $X$ until it becomes 0 and record the integer part each time.
Traversing the integer part in order will give use the representation in
base-$X$ system.

Example $464.625$ to base 8 (octal):
- $0.625 \times 8 = 5$ with integer 5. The fractional component is 0 so we
  stop.

Example $0.6875$ to base 8 (octal):
- $0.6875 \times 8 = 4,8125$ with integer $4$. Closest to 0.
- $0.8125 \times 8 = 6.5$ with integer $6$.
- $0.5 \times 8 = 4$ with integer $4$.

Therefore, $0.6875_{(10)} = 0.464_{(8)}$.

Note that a finite fraction in one base may become an infinite recurring
fraction in another base. For instance, the decimal $0.2$ wil become
$0.\dot{0}01\dot{1}$

## Common Conversions

The common practice to convert between non-base-10 numbers is to convert to
decimal first and then convert to the target base. Under certain conditions, we
can perform without going through decimal.

Converting binary number to octal or hexadecimal, and converting octal or
hexadecimal number to binary, there is no need to go through the intermediate
decimal step. Each digit in an octal number can be represented as three digits
in a binary number. Each digit in a hexadecimal number can be represented as
four digits in a binary number.

This works because 8 and 16 are both representable in base-2 or binary.
- 8 is `100`.
- 16 is `1000`.

> ![TIP]
> To convert from binary to octal create groups of 3 digits, those are the
> octal digits.
>
> $101110010_{(2)}$ will be `101|110|010` which is `5|6|2` or $562_{(8)}$ =
> $370_{(10)}$.
>
> To convert from binary to hexadecimal create groups of 4 digits, those are
> the hexadecimal digits.
>
> $101110010_{(2)}$ will be `1|0111|0010` which is `1|7|2` or $172_{(16)}$ =
> $370_{(10)}$.

## Math with Bits

Mathematics and specifically addition of bits is done one bit at a time. Any
carry that goes beyond the bits is _discarded_ (just like overflow with
fixed-width integers).
- $1 + 0 = 1$ and $0 + 1 = 1$.
- $0 + 0 = 0$
- $1 + 1 = 10$ which means keep the $0$ and _carry_ the $1$ to the next bit
  position.
- Finally, $1 + 0 + 0 = 1$, $1 + 1 + 0 = 10$ and $1 + 1 + 1 = 11$.
    - With $10$ keep the $0$ and _carry_ the $1$.
    - With $11$ keep the $1$ and _carry_ the $1$.

**Example** add `0010110` and `0000011`

```text
Carry:      11
            ↓↓
Bits :   0010110 = 22 = 16 + 0 + 4 + 0 + 0
       + 0000011 = 3  =  0 + 0 + 0 + 2 + 1
       ---------
         0011001 = 25 = 16 + 8 + 0 + 0 + 1
```

## Binary in Computers

Computers use binary, consisting of two digits: `0` and `1`. A single binary
digit has those two possible values, and a k-digit binary number can take $2^k$
possible values.

## Signed and Unsigned Integers

In signed integers, the highest bit represents the sign, and is called the
_sign bit_. When the bit is 0, it represents a non-negative integer; when the
bit is 1, it represents a negative integer. Digits other than the _sign bit_
are used to represent the size of the integer. In **unsigned integers**, all
the digits are used to represent the size of the number.

For k-bit integers:
- The **range** of a signed integers is $[-2^{k-1}, 2^{k-1}-1]$.
- The **range** of an unsigned integers $[0, 2^{k}-1]$, which is double that of
  signed integers.

### Machine Number and Truth Value

Take the 8-bit (1-byte) binary numbers. The decimal number $+10$ is `00001010`
in binary, and the decimal number $-10$ is `10001010` in binary.

Here the machine number (not what it represents) are `00001010` and `10001010`
respectively. For a machine `00001010` would represent $+10$ and `10001010`
would represent $+138$.

Because the highest bit of a machine number is the sign bit, the machine number
is not necessarily equal to what it represents (or the true value). Therefore,
to distinguish between the two, we name the true value of the machine number
the **truth value**. For example, the truth value of `00001010` is $+10$, and
the truth value of `10001010` is $-10$.

### The Original Code, Inverse Code, and Compliment Code

**Original Code**

The **original code** is the sign bit of the machine number plus the absolute
value of the truth value of the machine number.

**Inverse Code**

The **inverse code** is obtained from the _original code_. THe inverse code of
non-negative number is the same as the original code. The inverse of negative
numbers is to flip every bit of the original code except the sign bit.

- $+10$ original code `00001010` to inverse code `00001010`.
- $-10$ original code `10001010` to inverse code `11110101`.

**Complement Code**

The **complement code** is obtained from the _inverse code_. THe complement code
of non-negative numbers is the same as the original code and the inverse code.
The compliment code of negative numbers is obtained by adding `1` to the inverse
code.

- $+10$ original code `00001010` to inverse code `00001010` to compliment code
  `00001010`.
- $-10$ original code `10001010` to inverse code `11110101` to compliment code
  `11111010`.

Take special note of the negative compliment compared to the inverse.

## Representation in Computers

So far, we have seen three representations of signed binary numbers in
computers: original code, inverse code and complement code. So, what are the
advantages and disadvantages of these representations?

It is easy for the human brain to remember that the highest bit si the sign bit,
but for a computer, finding the sign bit will make computations complicated.
To simplify the calculations, people have developed a method to include the
sign bit in calculations. The original code is the easiest way to follow and
calculate by the human brain, but the original code has two problems:

1. There are representations of $+0$ and $-0$.
2. Performing subtractions with the original code will lead to incorrect
   results.

Introducing the inverse code solves the problem of subtraction errors, but the
issue of dual representation of $0$ remains. The _complement code_ **solves**
both the subtraction error and dual representation error.

Moreover, one more minimum value can be represented. In complement code, there
is no $-0$. Taking an 8-bit binary number, the complement code of $0$ is
`00000000`, while `1000000` represents $-128$.

$-128$ does not have any representation ins hte original code or the inverse
code. The minimum value that can be represented by the original code and the
inverse code with an 8-bit binary number is $-127$.

Therefore, the complement code not only solves the problems of the original
code and the inverse code, but can also represent one additional minimum value.

## Overview of Bit Operations

Computers use binary, which consists fo two numbers `0` and `1`. All the basic
operations are realized through bit operations. There are six types of bit
operations, namely: AND, OR, XOR, negative, left shift, and right shift.

Left and right shift are collectively referred to as shift operations, and shift
operations are further divided into arithmetic shift and logical shift.

Among these bit operations, only the negation is a unary (only one operand or
input) operation, and the rest are binary operations (require two operands or
two inputs).

**AND &**
For each binary bit, when the corresponding bits of both numbers are `1`, the
result is `1`; Otherwise, the result is `0`.
```
0 & 0 = 0
0 & 1 = 0
1 & 1 = 1
```

**OR |**
For each binary bit, when the corresponding bits of both numbers are `0`, the
result is `0`; Otherwise, the result is `1`.
```
0 | 0 = 0
1 | 0 = 1
1 | 1 = 1
```

**XOR ⊕ `^`**
For each binary bit, when the corresponding bits of the two numbers are the
same, the result is `0`; Otherwise, the result is `1`.
```
0 ⊕ 0 = 0
1 ⊕ 0 = 1
0 ⊕ 1 = 1
1 ⊕ 1 = 0
```

**Negation `~`**
Flip each binary it of a number: `0` becomes `1`, and `1` becomes `0`.
```
~1 = 0
~0 = 1
```