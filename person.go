// person.go
package main

import "fmt"

// Persona con nombre y edad
type Persona struct {
    Nombre string
    Edad   int
}

// Saludar imprime una frase con nombre y edad (receptor por valor)
func (p Persona) Saludar() {
    fmt.Printf("Hola, soy %s y tengo %d años.\n", p.Nombre, p.Edad)
}

// Cumplir Años incrementa la edad de la Persona (receptor por puntero)
func (p *Persona) CumplirAnios() {
    p.Edad++ // modifica el struct original
}

func main() {
    // Instanciar por valor
    p1 := Persona{Nombre: "Priscila", Edad: 32}
    p1.Saludar()           // Hola, soy Priscila y tengo 32 años.
    p1.CumplirAnios()      // p1.Edad pasa a 33
    fmt.Printf("Ahora tengo %d años.\n", p1.Edad)

    // Instanciar por puntero
    p2 := &Persona{Nombre: "Fernando", Edad: 33}
    p2.Saludar()           // Hola, soy Fernando y tengo 33 años.
    p2.CumplirAnios()      // p2.Edad pasa a 34
    p2.Saludar()           // Hola, soy Fernando y tengo 34 años.
}
