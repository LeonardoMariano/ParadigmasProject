package main

import "fmt"

func (a *Aluno) calculoMediaFinal() float32 {
	var a1 mediaA1
	var a2 mediaA1
	var mf mediaA1

	a1.total = (a.p1 * 0.7) + (a.ma1 * 0.2) + (a.mb1 * 0.1)
	a2.total = (a.p2 * 0.7) + (a.ma2 * 0.2) + (a.mb2 * 0.1)
	mf.total = ((a1.total + (2 * a2.total)) / 3)
	fmt.Println("media total", mf.total)

	return mf.total
}

func (a *Aluno) calculoFrequencia() float32 {

	var frequenciaAluno float32

	frequenciaAluno = (((a.qtdeAulas - a.qtdeFaltas) / a.qtdeAulas) * 100.0)
	fmt.Println("freq aluno", frequenciaAluno)

	return frequenciaAluno

}

func (a *Aluno) situacaoAluno() {
	if (a.mediaFinal >= 5.0) && (a.frequencia >= 75) {
		fmt.Println("APROVADO")
	} else if ((a.mediaFinal < 5.0) && (a.mediaFinal >= 3.0)) && a.frequencia >= 75 {
		fmt.Println("RECUPERAÇÃO")
	} else if (a.mediaFinal < 3.0) && (a.frequencia >= 75) {
		fmt.Println("REPROVADO POR NOTA")
	} else {
		fmt.Println("REPROVADO POR FALTA")
	}

}
