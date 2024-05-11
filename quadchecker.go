package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

func main() {
	// Définir les noms des exécutables
	executables := []string{"./quadA", "./quadB", "./quadC", "./quadD", "./quadE"}

	// Lire la sortie de l'exécutable précédent à partir de l'entrée standard
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println("Erreur lors de la lecture de l'entrée standard :", err)
		return
	}
	targetOutput := string(input)

	// Lire les valeurs de x et y à partir de l'entrée standard
	var x, y string
	fmt.Fscanf(os.Stdin, "%s %s\n", &x, &y)

	// Exécuter les autres exécutables et comparer la sortie
	for _, execName := range executables {
		cmd := exec.Command(execName, x, y)
		var out bytes.Buffer
		cmd.Stdout = &out
		err := cmd.Run()
		if err != nil {
			fmt.Printf("Erreur lors de l'exécution de %s : %v\n", execName, err)
			continue
		}
		output := out.String()

		if strings.EqualFold(targetOutput, output) {
			fmt.Printf("[%s] [%s] [%s] || ", execName, x, y)
		}
	}
	fmt.Println()
}
