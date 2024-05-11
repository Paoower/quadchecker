package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 3 {
		return
	}
	x, _ := strconv.Atoi(os.Args[1])
	y, _ := strconv.Atoi(os.Args[2])
	QuadC(x, y)
	CreateTempFile(x, y)
}

func QuadC(x, y int) {
	// print first line
	if x > 0 && y > 0 {
		PrintLine(x, 'A', 'A', 'B')
		// print intermediate lines
		for i := 0; i < y-2; i++ {
			PrintIntermediateLine(x, 'B')
		}
		// print last line
		if y > 1 {
			PrintLine(x, 'C', 'C', 'B')
		}
	}
}

func PrintLine(x int, leftCorner, rightCorner, inter rune) {
	z01.PrintRune(leftCorner)
	if x > 1 {
		for i := 0; i < x-2; i++ {
			z01.PrintRune(inter)
		}
		z01.PrintRune(rightCorner)
	}
	z01.PrintRune('\n')
}
func PrintIntermediateLine(x int, edge rune) {
	for i := 0; i <= x-1; i++ {
		if i == 0 || i == x-1 {
			z01.PrintRune(edge)
		} else {
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune('\n')
}

func CreateTempFile(x, y int) {
	// Obtenir le répertoire de travail courant
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Erreur lors de l'obtention du répertoire de travail courant :", err)
		os.Exit(1)
	}

	// Créer un fichier non temporaire avec un nom fixe dans le répertoire de travail courant pour stocker les valeurs de x et y
	fileName := dir + "/quad-x-y.txt"
	err = ioutil.WriteFile(fileName, []byte(fmt.Sprintf("%s %s\n", x, y)), 0644)
	if err != nil {
		fmt.Println("Erreur lors de l'écriture dans le fichier :", err)
		os.Exit(1)
	}
}
