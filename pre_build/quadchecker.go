package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

var isStarting bool = false

func main() {
	// Lire la sortie de l'executable precedent à partir de l'entree standard
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println("Erreur lors de la lecture de l'entree standard :", err)
		os.Exit(1)
	}
	targetOutput := string(input)
	if !(targetOutput >= "A" && targetOutput <= "E") {
		fmt.Println("Not a quad function")
		return
	}

	// Lire les valeurs de x et y à partir du fichier quad-x-y.txt
	data, err := ioutil.ReadFile("quad-x-y.txt")
	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier :", err)
		os.Exit(1)
	}

	// Extraire les valeurs de x et y à partir des chaînes de caractères formatees avec %!s(int=...)
	xStr := strings.Trim(string(data[:len(data)/2]), "%!s(int=) ")
	xInt, err := strconv.Atoi(xStr)
	if err != nil {
		fmt.Println("Erreur lors de l'extraction de la valeur de x à partir de la chaîne de caractères formatee :", err)
		os.Exit(1)
	}
	yStr := strings.Trim(string(data[len(data)/2:]), "%!s(int=) \n")
	yInt, err := strconv.Atoi(yStr)
	if err != nil {
		fmt.Println("Erreur lors de l'extraction de la valeur de y à partir de la chaîne de caractères formatee :", err)
		os.Exit(1)
	}

	// Definir les noms des executables
	executables := []string{"quadA", "quadB", "quadC", "quadD", "quadE"}

	// Executer les autres executables et comparer la sortie
	for _, execName := range executables {
		// Executer l'executable avec les valeurs de x et y extraites
		cmd := exec.Command("./"+execName, strconv.Itoa(xInt), strconv.Itoa(yInt))
		var out bytes.Buffer
		cmd.Stdout = &out
		err := cmd.Run()
		if err != nil {
			fmt.Printf("Erreur lors de l'execution de %s : %v\n", execName, err)
			continue
		}
		output := out.String()

		// Comparer la sortie avec la sortie cible en utilisant strings.EqualFold
		if strings.EqualFold(targetOutput, output) {
			if !isStarting {
				fmt.Printf("[%s] [%d] [%d]", execName, xInt, yInt)
				isStarting = true
			} else {
				fmt.Printf(" || [%s] [%d] [%d]", execName, xInt, yInt)
			}
		}
	}
	fmt.Println()

	// Supprimer le fichier quad-x-y.txt
	err = os.Remove("quad-x-y.txt")
	if err != nil {
		fmt.Println("Erreur lors de la suppression du fichier :", err)
		os.Exit(1)
	}
}
