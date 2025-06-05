// calculadora/calc_test.go
package calculadora

import "testing"

// TestSumar verifica que Sumar devuelva la suma correcta.
func TestSumar(t *testing.T) {
    if Sumar(2, 3) != 5 {
        t.Errorf("Sumar(2,3) = %d; se esperaba 5", Sumar(2, 3))
    }
}

// TestRestar verifica que Restar devuelva a - b
func TestRestar(t *testing.T) {
    if Restar(10, 4) != 6 {
        t.Errorf("Restar(10,4) = %d; se esperaba 6", Restar(10, 4))
    }
}

// TestFactorialCasos verifica varios escenarios de Factorial
func TestFactorialCasos(t *testing.T) {
    // Caso n positivo
    fact, err := Factorial(5)
    if err != nil || fact != 120 {
        t.Errorf("Factorial(5) = %d, %v; se esperaba (120, nil)", fact, err)
    }
    // Caso n = 0
    fact0, err0 := Factorial(0)
    if err0 != nil || fact0 != 1 {
        t.Errorf("Factorial(0) = %d, %v; se esperaba (1, nil)", fact0, err0)
    }
    // Caso n negativo
    _, errNeg := Factorial(-3)
    if errNeg == nil {
        t.Errorf("Factorial(-3) no devolvió error; se esperaba error")
    }
}
