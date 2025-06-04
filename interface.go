package main

import "fmt"

// Pagable define el comportamiento de un método Pagar
type Pagable interface {
    Pagar(monto float64) error
}

// Tarjeta simula un pago con tarjeta de crédito
type Tarjeta struct {
    Titular string
    Numero  string
}

func (t *Tarjeta) Pagar(monto float64) error {
    // Aquí iría la lógica real; simulamos un pago exitoso
    fmt.Printf("Pagando %.2f con tarjeta de %s (****%s)\n",
        monto, t.Titular, t.Numero[len(t.Numero)-4:])
    return nil
}

// Transferencia simula un pago por transferencia bancaria
type Transferencia struct {
    Banco   string
    Cuenta  string
}

func (t *Transferencia) Pagar(monto float64) error {
    // Simulamos validación de datos
    if t.Cuenta == "" {
        return fmt.Errorf("cuenta bancaria inválida")
    }
    fmt.Printf("Transferencia de %.2f a cuenta %s del banco %s\n",
        monto, t.Cuenta, t.Banco)
    return nil
}

func procesarPago(p Pagable, monto float64) {
    // Llamamos al método Pagar de la interfaz
    if err := p.Pagar(monto); err != nil {
        fmt.Println("Error al pagar:", err)
    } else {
        fmt.Println("Pago realizado con éxito")
    }
}

func main() {
    t := &Tarjeta{Titular: "Priscila Flores", Numero: "1234567890123456"}
    procesarPago(t, 150.75)

    fmt.Println() // salto de línea

    tr := &Transferencia{Banco: "Banco XYZ", Cuenta: ""}
    procesarPago(tr, 200.00)

    // Ahora con cuenta válida
    tr.Cuenta = "0987654321"
    procesarPago(tr, 200.00)
}
