// composition.go
package main

import "fmt"

// ------------------------------------------------------------
// 1. Structs básicos
// ------------------------------------------------------------

// Direccion representa la dirección de una persona.
type Direccion struct {
    Calle  string
    Ciudad string
    Codigo string
}

// PersonaSimple tiene campos básicos de una persona.
type PersonaSimple struct {
    Nombre string
    Edad   int
}

// Saludar es un método asociado a PersonaSimple.
func (p PersonaSimple) Saludar() {
    fmt.Printf("Hola, soy %s y tengo %d años.\n", p.Nombre, p.Edad)
}

// ------------------------------------------------------------
// 2. Composición: embebiendo structs
// ------------------------------------------------------------

// PersonaCompleta “embebe” PersonaSimple y Direccion.
// Gracias a esto, PersonaCompleta hereda los campos y métodos de ambos,
// excepto que desambiguaremos “Info” definiendo uno propio más abajo.
type PersonaCompleta struct {
    PersonaSimple // embed de PersonaSimple
    Direccion     // embed de Direccion
    Email         string // campo adicional propio
}

// ------------------------------------------------------------
// 3. Métodos “propios” de PersonaCompleta
// ------------------------------------------------------------

// Contactar imprime un mensaje que usa campos de PersonaSimple y Direccion.
func (pc PersonaCompleta) Contactar() {
    fmt.Printf(
        "Enviando correo a %s: hola %s, vives en %s, %s. ¡Saludos!\n",
        pc.Email, pc.Nombre, pc.Calle, pc.Ciudad,
    )
}

// Info unifica las implementaciones de PersonaSimple.Info y Direccion.Info,
// evitando la ambigüedad al pertenecer a la interfaz Informacion.
func (pc PersonaCompleta) Info() string {
    infoPersona := pc.PersonaSimple.Info()   // "Persona: Nombre, Edad"
    infoDireccion := pc.Direccion.Info()     // "Dirección: Calle, Ciudad, CP Código"
    return fmt.Sprintf("%s, Email: %s, %s", infoPersona, pc.Email, infoDireccion)
}

// ------------------------------------------------------------
// 4. Composición + interfaces
// ------------------------------------------------------------

// Informacion es una interfaz que exige un método Info() string.
type Informacion interface {
    Info() string
}

// Implementamos “Info” en PersonaSimple.
func (p PersonaSimple) Info() string {
    return fmt.Sprintf("Persona: %s, Edad: %d", p.Nombre, p.Edad)
}

// Implementamos “Info” en Direccion.
func (d Direccion) Info() string {
    return fmt.Sprintf("Dirección: %s, %s, CP %s", d.Calle, d.Ciudad, d.Codigo)
}

// printInfo recibe cualquier tipo que cumpla Informacion.
func printInfo(i Informacion) {
    fmt.Println(" Info:", i.Info())
}

// ------------------------------------------------------------
// 5. main: demostración
// ------------------------------------------------------------

func main() {
    // 5.1. Instanciar solo PersonaSimple y Direccion
    pSim := PersonaSimple{Nombre: "Ana", Edad: 25}
    dir := Direccion{Calle: "Av. Siempre Viva 742", Ciudad: "Springfield", Codigo: "12345"}

    // Llamamos a Saludar (método de PersonaSimple)
    pSim.Saludar()
    // Llamamos a Info() en cada struct porque cumplen Informacion
    printInfo(pSim)
    printInfo(dir)

    fmt.Println()

    // 5.2. Instanciar PersonaCompleta
    pc := PersonaCompleta{
        PersonaSimple: PersonaSimple{Nombre: "Carlos", Edad: 30},
        Direccion:     Direccion{Calle: "Calle Falsa 123", Ciudad: "Shelbyville", Codigo: "67890"},
        Email:         "carlos@ejemplo.com",
    }

    // 5.2.1. Podemos usar directamente el método Saludar que viene de PersonaSimple
    pc.Saludar() // Go auto‐desreferencia y busca Saludar en PersonaSimple

    // 5.2.2. También usamos el método Contactar propio de PersonaCompleta
    pc.Contactar()

    // 5.2.3. Accedemos a campos embebidos sin calificar:
    fmt.Printf("%s vive en %s, CP %s\n", pc.Nombre, pc.Ciudad, pc.Codigo)

    fmt.Println()

    // 5.2.4. PersonaCompleta ya tiene su propia Info() definida,
    // así que al tratar a PersonaCompleta como Informacion Go usará esta implementación.
    printInfo(pc)

    // 5.3. Slice de Informacion con varios tipos
    fmt.Println("\n--- Slice de Informacion ---")
    var lista []Informacion
    lista = append(lista, pSim, dir, pc) // Ya no hay ambigüedad porque PersonaCompleta.Info está definido

    for i, elemento := range lista {
        fmt.Printf("Elemento %d: %s\n", i+1, elemento.Info())
    }
}
