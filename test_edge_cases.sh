#!/bin/bash
echo "Testing edge cases..."

echo "1. Zero coefficients:"
./computor "0 * X^2 + 0 * X^1 + 0 * X^0 = 0"
./computor "0 * X = 5"

echo "2. Negative coefficients:"
./computor "-X^2 = 0"
./computor "X^2 = -4"

echo "3. Decimal coefficients:"  
./computor "1.5 * X^2 + 2.7 * X = 0"

echo "4. Complex solutions:"
./computor "X^2 + 1 = 0"
