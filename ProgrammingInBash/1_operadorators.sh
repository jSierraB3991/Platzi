#! /bin/bash

num=10
num2=5

echo "Operadores Aritmeticos"
echo "Numeros A: num y B: num2"
echo "La suma de A y B = $((num + num2))"
echo "La resta de A y B = $((num - num2))"
echo "La multiplicacion de A y B = $((num * num2))"
echo "La division de A y B = $((num / num2))"
echo "El resto de la division de A y B = $((num % num2))"

echo "Operadores Relacionales"
echo "Numeros A: num y B: num2"
echo "A > B = $((num > num2))"
echo "A < B = $((num < num2))"
echo "A >= B = $((num >= num2))"
echo "A <= B = $((num <= num2))"
echo "A == B = $((num == num2))"
echo "A != B = $((num != num2))"

echo "Operadores de Asignacion"
echo "Numeros A: num y B: num2"
echo "Sumar A: $num += B: $num2 = $((num += num2))"
echo "Restar A: $num -= B: $num2 = $((num - num2))"
echo "Multiplicar A: $num *= B: $num2 = $((num *= num2))"
echo "Dividir A: $num /= B: $num2 = $((num /= num2))"
echo "Residuo A: $num %= B: $num2 = $((num %= num2))"
