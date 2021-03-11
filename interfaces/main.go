package main

import "fmt"

type serVivo interface {
	estaVivo() bool
}

type humano interface {
	respirar()
	pensar()
	comer()
	sexo() string
	estaVivo() bool
}

type animal interface {
	respirar()
	comer()
	esCarnivoro() bool
	estaVivo() bool
}

type vegetal interface {
	clasificacionVegetal() string
}

// Genero Humano
type hombre struct {
	edad       int
	altura     float32
	peso       float32
	respirando bool
	pensando   bool
	comiendo   bool
	esHombre   bool
	vivo       bool
}

type mujer struct {
	hombre
}

func (h *hombre) respirar() { h.respirando = true }
func (h *hombre) comer()    { h.comiendo = true }
func (h *hombre) pensar()   { h.pensando = true }
func (h *hombre) sexo() string {
	if h.esHombre == true {
		return "Hombre"
	}
	return "Mujer"
}
func (h *hombre) estaVivo() bool { return h.vivo }

func humanosRespirando(hu humano) {
	hu.respirar()
	fmt.Printf("Soy un/a %s, y ya estoy respirando \n", hu.sexo())
}

// ------------------------------------
// Genero Animal

type perro struct {
	respirando bool
	comiendo   bool
	carnivoro  bool
	vivo       bool
}

func (p *perro) respirar()         { p.respirando = true }
func (p *perro) comer()            { p.comiendo = true }
func (p *perro) esCarnivoro() bool { return p.carnivoro }
func (p *perro) estaVivo() bool    { return p.vivo }

func animalesRespirar(an animal) {
	an.respirar()
	fmt.Println("Soy un animal y estoy respirando")
}

func animalesCarnivoros(an animal) int {
	if an.esCarnivoro() == true {
		return 1
	}
	return 0
}

// --------------------------------------
// serVivo

func estoyVivo(v serVivo) bool {
	return v.estaVivo()
}

func main() {
	Pedro := new(hombre)
	Pedro.esHombre = true
	Pedro.vivo = true
	humanosRespirando(Pedro)
	fmt.Printf("Estoy vivo = %t\n", estoyVivo(Pedro))

	Maria := new(mujer)
	Maria.vivo = true
	humanosRespirando(Maria)
	fmt.Printf("Estoy viva = %t\n", estoyVivo(Maria))

	totalCarnivoros := 0
	Dogo := new(perro)
	Dogo.carnivoro = true
	Dogo.vivo = true
	animalesRespirar(Dogo)
	totalCarnivoros = +animalesCarnivoros(Dogo)

	fmt.Printf("Total carnivoros %d\n", totalCarnivoros)

	fmt.Printf("Esta vivo = %t", estoyVivo(Dogo))
}
