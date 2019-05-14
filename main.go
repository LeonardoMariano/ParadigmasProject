package main

import "fmt"

func main() {

	var continua = true
	var aluno Aluno

	for continua == true {

		fmt.Println("Informe a nota da Prova 1:")
		fmt.Scan(&aluno.p1)
		if aluno.p1 < 0 || aluno.p1 > 10 {
			break
		}
		fmt.Println("Informe a nota da Ma1:")
		fmt.Scan(&aluno.ma1)
		if aluno.ma1 < 0 || aluno.ma1 > 10 {
			break
		}
		fmt.Println("Informe a nota da Mb1:")
		fmt.Scan(&aluno.mb1)
		if aluno.mb1 < 0 || aluno.mb1 > 10 {
			break
		}
		fmt.Println("Informe a nota da Prova 2:")
		fmt.Scan(&aluno.p2)
		if aluno.p2 < 0 || aluno.p2 > 10 {
			break
		}
		fmt.Println("Informe a nota da Ma2:")
		fmt.Scan(&aluno.ma2)
		if aluno.ma2 < 0 || aluno.ma2 > 10 {
			break
		}
		fmt.Println("Informe a nota da Mb2:")
		fmt.Scan(&aluno.mb2)
		if aluno.mb2 < 0 || aluno.mb2 > 10 {
			break
		}
		fmt.Println("Informe a quantidade de aulas ministradas no semestre:")
		fmt.Scan(&aluno.qtdeAulas)
		if aluno.qtdeAulas < 0 {
			break
		}
		fmt.Println("Informe a quantidade de faltas do aluno no semestre:")
		fmt.Scan(&aluno.qtdeFaltas)
		if aluno.qtdeFaltas < 0 {
			break
		}

		aluno.mediaFinal = aluno.calculoMediaFinal()
		fmt.Println("A média final do aluno é: ", aluno.mediaFinal)
		aluno.frequencia = aluno.calculoFrequencia()

		aluno.situacaoAluno()

		fmt.Println("Você deseja fazer uma nova consulta? (true/false)")
		fmt.Scan(&continua)

	}
}
