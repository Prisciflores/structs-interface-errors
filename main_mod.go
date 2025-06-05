// main_mod.go
package main

import (
    "fmt"
    // Importamos el paquete local miutil. El prefijo debe coincidir con el módulo.
    "structs-interface-errors/miutil"
)

func main() {
    // 1) Usar SaludarFormateado del paquete miutil
    saludo := miutil.SaludarFormateado("Priscila")
    fmt.Println(saludo)

    // 2) Usar SumInt para sumar dos números
    resultado := miutil.SumInt(5, 7)
    fmt.Printf("La suma de 5 + 7 es: %d\n", resultado)

    // 3) Podemos combinar ambas funcionalidades
    otro := miutil.SaludarFormateado("Carlos")
    fmt.Println(otro)

    fmt.Println("fin del programa de ejemplo de módulos y paquetes")
}
