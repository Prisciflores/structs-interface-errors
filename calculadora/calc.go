// calculadora/calc.go
package calculadora

import "fmt"

// Sumar recibe dos enteros y devuelve su suma.
func Sumar(a, b int) int {
    return a + b
}

// Restar recibe dos enteros y devuelve a - b.
func Restar(a, b int) int {
    return a - b
}

// Factorial calcula el factorial de n (n >= 0).
// Si n < 0, retorna un error.
func Factorial(n int) (int, error) {
    if n < 0 {
        return 0, fmt.Errorf("no se puede calcular factorial de un número negativo: %d", n)
    }
    resultado := 1
    for i := 2; i <= n; i++ {
        resultado *= i
    }
    return resultado, nil
}
