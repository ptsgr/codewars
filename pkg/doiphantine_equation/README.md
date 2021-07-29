# [Diophantine Equation](https://www.codewars.com/kata/554f76dca89983cc400000bb)

___5 kyu___

In mathematics, a [Diophantine equation](https://en.wikipedia.org/wiki/Diophantine_equation) is a polynomial equation, usually with two or more unknowns, such that only the integer solutions are sought or studied.

In this kata we want to find all integers `x, y` `(x >= 0, y >= 0)` solutions of a diophantine equation of the form:

x^2^ - 4 * y^2^ = n

(where the unknowns are `x` and `y`, and `n` is a given positive number) in decreasing order of the positive `x~i~`.

If there is no solution return `[]` or `"[]"` or `""`. (See "RUN SAMPLE TESTS" for examples of returns).

## Examples

```bash
solEquaStr(90005) --> "[[45003, 22501], [9003, 4499], [981, 467], [309, 37]]"
solEquaStr(90002) --> "[]"
```

## Hint

$x^2-4*y^2=(x-2*y)*(x+2*y)$
