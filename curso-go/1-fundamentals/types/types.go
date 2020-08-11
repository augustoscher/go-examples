package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	// 1. Integers
	fmt.Println(1, 3, 1000)
	fmt.Println("Integer literals: ", reflect.TypeOf(32000))

	// sem sinal (só positivos)
	// uint8 = byte (8 bits)
	// uint16 = short (2 bytes)
	// uint32 = int (4 bytes)
	// uint64 = long (8 bytes)

	var b byte = 255
	fmt.Println("byte ", reflect.TypeOf(b))

	// com sinal int8 int16 int32 int64
	i1 := math.MaxInt64
	fmt.Println("O valor máximo de int64 é:", i1)
	fmt.Println("O tipo é:", reflect.TypeOf(i1))

	// representa o mapeamento de um valor na tabela unicode (int32)
	var i2 rune = 'a'
	fmt.Println("O tipo rune é:", reflect.TypeOf(i2)) //int32
	fmt.Println("O rune é:", i2)                      //97

	//2. Números Reais (float32, float64)
	var x float32 = 49.99
	fmt.Println("O valor do float32 é:", x)
	// toda vez que delcaramos um float, ele instancia uma float64:
	fmt.Println("Pegando o literal de 49.99 é:", reflect.TypeOf(49.99)) //float64

	//3. Boolean
	bo := true
	fmt.Println("Boolean literal:", reflect.TypeOf(bo)) //bool
	fmt.Println("Valor:", bo)

	// 4. String (somente aspas duplas)
	s1 := "Olá, tudo bem com você?"
	fmt.Println("String literal:", reflect.TypeOf(s1)) //string
	fmt.Println("Valor de String:", s1)
	fmt.Println("Tamanho de String:", len(s1))

	// String multipas linhas
	s2 := `Olá,
	tudo bem com vocẽ?`
	fmt.Println("String multipas linhas:", s2)
	fmt.Println("Tamanho de String multi linhas:", len(s2))

	// 5. char - construção com aspas simples
	char := 'a'
	fmt.Println("Não existe char, e sim int32:", reflect.TypeOf(char)) //int32
	fmt.Println("value de char:", char)                                //97
}
