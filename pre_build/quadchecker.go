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

func main() {
	// Lire la sortie de l'exécutable précédent à partir de l'entrée standard
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println("Erreur lors de la lecture de l'entrée standard :", err)
		os.Exit(1)
	}
	targetOutput := string(input)

	// Lire les valeurs de x et y à partir du fichier quad-x-y.txt
	data, err := ioutil.ReadFile("quad-x-y.txt")
	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier :", err)
		os.Exit(1)
	}

	// Extraire les valeurs de x et y à partir des chaînes de caractères formatées avec %!s(int=...)
	xStr := strings.Trim(string(data[:len(data)/2]), "%!s(int=) ")
	xInt, err := strconv.Atoi(xStr)
	if err != nil {
		fmt.Println("Erreur lors de l'extraction de la valeur de x à partir de la chaîne de caractères formatée :", err)
		os.Exit(1)
	}
	yStr := strings.Trim(string(data[len(data)/2:]), "%!s(int=) \n")
	yInt, err := strconv.Atoi(yStr)
	if err != nil {
		fmt.Println("Erreur lors de l'extraction de la valeur de y à partir de la chaîne de caractères formatée :", err)
		os.Exit(1)
	}

	// Définir les noms des exécutables
	executables := []string{"./quadA", "./quadB", "./quadC", "./quadD", "./quadE"}

	// Exécuter les autres exécutables et comparer la sortie
	for _, execName := range executables {
		// Exécuter l'exécutable avec les valeurs de x et y extraites
		cmd := exec.Command(execName, strconv.Itoa(xInt), strconv.Itoa(yInt))
		var out bytes.Buffer
		cmd.Stdout = &out
		err := cmd.Run()
		if err != nil {
			fmt.Printf("Erreur lors de l'exécution de %s : %v\n", execName, err)
			continue
		}
		output := out.String()

		// Comparer la sortie avec la sortie cible en utilisant strings.EqualFold
		if strings.EqualFold(targetOutput, output) {
			fmt.Printf("[%s] [%d] [%d] || ", execName, xInt, yInt)
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
