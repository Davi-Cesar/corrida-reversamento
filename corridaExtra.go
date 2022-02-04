package main

// Problema 2: Uma Corrida de Revezamento Extra
import (
	"fmt"
	"sync"
	"time"
)

var waitGroup sync.WaitGroup // Pacote waitGroup serva para suspender um Go rotina até que as outras finalize
// Pegando o tempo de cada time
var start = time.Now()
var timeAzul = time.Since(start)
var timeVermelho = time.Since(start)

func main() {

	waitGroup.Add(6) // Numero de Go rotinas que a main tem que esperar
	execute()        // Função que vai exercutaras go rotinas
	waitGroup.Wait() // Suspende a main, para esperar as outras Go rotinas

	if timeAzul > timeVermelho {
		fmt.Println("---------------------------------")
		fmt.Println("Equipe vermelha é a vencedora!!")
		fmt.Println("---------------------------------")
	} else if timeAzul < timeVermelho {
		fmt.Println("---------------------------------")
		fmt.Println("Equipe Azul é a vencedora!!")
		fmt.Println("---------------------------------")
	} else {
		fmt.Println("---------------------------------")
		fmt.Println("Equipes Empataram!!")
		fmt.Println("---------------------------------")
	}
}

func execute() {

	go func() {
		corrida("azul", 1, 0, 20)
		waitGroup.Done() // Go rotina finaliza
		go func() {
			corrida("azul", 2, 20, 40)
			waitGroup.Done()
			go func() {
				corrida("azul", 3, 40, 60)
				waitGroup.Done()
			}()
		}()
		timeEquipeAzul := time.Since(start)
		timeAzul = timeEquipeAzul
	}()

	start = time.Now()
	go func() {
		corrida("vermelha", 1, 0, 20)
		waitGroup.Done()
		go func() {
			corrida("vermelha", 2, 20, 40)
			waitGroup.Done()
			go func() {
				corrida("vermelha", 3, 40, 60)
				waitGroup.Done()
			}()
		}()
		timeEquipeVermelha := time.Since(start)
		timeVermelho = timeEquipeVermelha
	}()

}
func corrida(equipe string, corredor int, inicio int, fim int) { // a funação responsavel pela corrida das equipes

	for i := inicio; i <= fim; i++ {
		if i != inicio {
			time.Sleep(1e8)
			fmt.Println("Equipe", equipe, "Corredor", corredor, "percorrendo", i, "km")
		}
	}
	if corredor == 3 { // se for o ultimo
		fmt.Println("fim da corrida equipe", equipe)
	} else {
		fmt.Println("corredor", corredor, "da equipe", equipe, "passou o bastaão para corredor", corredor+1)
	}
}
