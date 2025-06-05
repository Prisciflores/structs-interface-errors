# structs-interface-errors

Ejercicios de la **Semana 2** que profundizan en Go con:

- Punteros, structs y métodos
- Interfaces y composición
- Organización en módulos y paquetes
- Manejo de errores y pruebas unitarias

---

## 📂 Contenido del directorio

```plaintext
structs-interface-errors/
├── go.mod
├── person.go           # Ejemplo de struct con métodos (receptores por valor y puntero)
├── interface.go        # Ejemplo de interfaces (Pagable) y tipos que la implementan
├── composition.go      # Demostración de composición (embedded structs + conflictos de Info)
├── main_math.go        # Programa principal que usa el paquete interno 'calculadora'
├── calculadora/        # Paquete interno con funciones matemáticas y sus tests
│   ├── calc.go         # Funciones: Sumar, Restar, Factorial (con manejo de error)
│   └── calc_test.go    # Tests unitarios para Sumar, Restar y Factorial
└── README.md

---

## 📂 Descripción de cada archivo

go.mod
Define el módulo raíz como structs-interface-errors.
```
module structs-interface-errors

go 1.22

```


