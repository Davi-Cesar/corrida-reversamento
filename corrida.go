package main

// Problema 2: Uma Corrida de Revezamento
import (
	"fmt"
	"sync"
	"time"
)

var waitGroup sync.WaitGroup // Pacote waitGroup serva para suspender um Go rotina até que as outras finalize
func main() {

	waitGroup.Add(3) // Numero de Go rotinas que a main tem que esperar
	execute()        // Função que vai exercutaras go rotinas
	waitGroup.Wait() // Suspende a main, para esperar as outras Go rotinas

}

func execute() {
	go func() {
		corrida(1, 0, 20)
		waitGroup.Done() // Go rotina finaliza
		go func() {
			corrida(2, 20, 40)
			waitGroup.Done()
			go func() {
				corrida(3, 40, 60)
				waitGroup.Done()
			}()
		}()

	}()
}

func corrida(corredor int, inicio int, fim int) { // a funação responsavel pela corrida das equipes

	for i := inicio; i <= fim; i++ {
		if i != inicio {
			time.Sleep(1e8)
			fmt.Println("Corredor", corredor, "percorrendo", i, "km")
		}
	}
	if corredor == 3 { // se for o ultimo
		fmt.Println("-----------------")
		fmt.Println("Fim da corrida!")
		fmt.Println("-----------------")
	} else {
		fmt.Println("corredor", corredor, "passou o bastaão para corredor", corredor+1)
	}
}
