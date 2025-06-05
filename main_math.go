// main_math.go
package main

import (
    "fmt"
    // Importamos el paquete interno calculadora.
    // El prefijo debe coincidir con el módulo definido en go.mod.
    "structs-interface-errors/calculadora"
)

func main() {
    // 1) Uso de Sumar y Restar
    sum := calculadora.Sumar(8, 5)
    fmt.Printf("8 + 5 = %d\n", sum)

    diff := calculadora.Restar(15, 7)
    fmt.Printf("15 - 7 = %d\n", diff)

    // 2) Uso de Factorial con un caso válido
    fact, err := calculadora.Factorial(6)
    if err != nil {
        fmt.Println("Error al calcular factorial:", err)
    } else {
        fmt.Printf("Factorial(6) = %d\n", fact)
    }

    // 3) Intento de Factorial con número negativo
    _, errNeg := calculadora.Factorial(-2)
    if errNeg != nil {
        fmt.Println("Error esperado:", errNeg)
    }

    fmt.Println("Fin del programa principal de matemáticas.")
}
