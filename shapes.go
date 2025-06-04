// shapes.go
package main

import (
    "fmt"
    "math"
)

// Forma es una interfaz que declara dos métodos: Area y Perimetro.
// Cualquier tipo que tenga estos dos métodos satisface la interfaz Forma.
type Forma interface {
    Area() float64
    Perimetro() float64
}

// Rectangulo representa un rectángulo con base y altura.
type Rectangulo struct {
    Base   float64
    Altura float64
}

// Los métodos de Rectangulo (con receptor por valor) calculan área y perímetro.

func (r Rectangulo) Area() float64 {
    // Área = base * altura
    return r.Base * r.Altura
}

func (r Rectangulo) Perimetro() float64 {
    // Perímetro = 2*(base + altura)
    return 2 * (r.Base + r.Altura)
}

// Circulo representa un círculo con un radio.
type Circulo struct {
    Radio float64
}

// Los métodos de Circulo (con receptor por valor) calculan área y perímetro.

func (c Circulo) Area() float64 {
    // Área = π * radio^2
    return math.Pi * c.Radio * c.Radio
}

func (c Circulo) Perimetro() float64 {
    // Perímetro = 2 * π * radio
    return 2 * math.Pi * c.Radio
}

// ResumenForma recibe cualquier implementación de Forma
// y muestra su tipo, área y perímetro.
func ResumenForma(f Forma) {
    // Aquí usamos un type assertion para identificar qué tipo concreto
    // implementa la interfaz. Esto es opcional, sólo para demostración.
    switch forma := f.(type) {
    case Rectangulo:
        fmt.Println("→ Detalles de Rectángulo:")
        fmt.Printf("   Base: %.2f, Altura: %.2f\n", forma.Base, forma.Altura)
    case Circulo:
        fmt.Println("→ Detalles de Círculo:")
        fmt.Printf("   Radio: %.2f\n", forma.Radio)
    default:
        fmt.Println("→ Detalle de forma desconocida")
    }

    // Imprimimos el área y el perímetro usando la interfaz
    fmt.Printf("Área:      %.2f\n", f.Area())
    fmt.Printf("Perímetro: %.2f\n", f.Perimetro())
    fmt.Println()
}

func main() {
    // Creamos instancias de Rectángulo y Círculo
    r := Rectangulo{Base: 4, Altura: 3}
    c := Circulo{Radio: 2.5}

    // Llamamos a ResumenForma con cada una. Ambas satisfacen la interfaz Forma.
    fmt.Println("=== RESUMEN DE FORMAS ===")
    ResumenForma(r)
    ResumenForma(c)

    // También podemos guardar distintas Formas en un slice y procesarlas en bucle:
    formas := []Forma{r, c}

    fmt.Println("=== BUCLE SOBRE SLICE DE FORMAS ===")
    for i, f := range formas {
        fmt.Printf("Forma #%d:\n", i+1)
        ResumenForma(f)
    }
}
