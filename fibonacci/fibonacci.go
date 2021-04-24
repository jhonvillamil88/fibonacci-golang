package main

import (
	"fmt"
	"log"
	"strconv"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	var count int
	fmt.Println("Iniciando aplicacion para generar serie de fibonacci")
	fmt.Print("Ingresar cantidad de numeros a calcular: ")
	var input string
	fmt.Scanln(&input)
	count, err := strconv.Atoi(input)
	if err == nil {
		makeManual(count)
	} else {
		fmt.Println(err)
		fmt.Println("Ocurrio un error al capturar la cantidad de numeros")
	}

}

func print(output int) {
	fmt.Print(output)
}

func makeManual(count int) {
	var before int = 0
	var after int = 1
	var aux int = 0
	var values []int
	fmt.Println("Serie fibonacci ")
	for i := 0; i < count; i++ {
		fmt.Print(" ")
		print(before)
		fmt.Print(" ")
		aux = before
		before = after
		after = aux + after
		values = append(values, before)
	}
	linePlot(values, count)
	fmt.Println("")
}

func makePoint(values []int, count int) plotter.XYs {
	pts := make(plotter.XYs, count)
	for i := range pts {
		pts[i].Y = float64(values[i])
		pts[i].X = float64(i)
	}
	return pts
}

func linePlot(values []int, count int) {

	points := makePoint(values, count)
	p := plot.New()
	p.Title.Text = "Fibonacci grafica "
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Serie fibonacci"
	p.Add(plotter.NewGrid())

	filled, err := plotter.NewLine(points)
	if err != nil {
		log.Panic(err)
	}
	// filled.FillColor = color.RGBA{R: 196, G: 255, B: 196, A: 255}

	p.Add(filled)

	err = p.Save(200, 200, "Fibonacci-"+fmt.Sprint(count)+".png")
	if err != nil {
		log.Panic(err)
	}
}

func histPlot(values plotter.Values) {
	// var p plot
	p := plot.New()

	p.Title.Text = "histogram plot"

	hist, err := plotter.NewHist(values, 20)
	if err != nil {
		// fmt.Println("here!")
		panic(err)
	}
	// fmt.Println("here1!")
	p.Add(hist)
	// fmt.Println("here2!")

	if err := p.Save(3*vg.Inch, 3*vg.Inch, "hist.png"); err != nil {
		// fmt.Println("here3!")
		panic(err)
	}
	// fmt.Println("here4!")
}
