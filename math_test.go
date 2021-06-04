package main

import "testing"

fun TestSoma(t *testing.T){
	total := Soma(15,15)

	if total != 30 {
		t.Errorf("Resultado da some é invalido: Resultado %d", total, 30)
	}
}