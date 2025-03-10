package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

var isStarting bool = false // Variable globale permettant de gérer les séparateurs "||" dans le résultat.

func main() {
	originalOutput := OutputOfExec() // Met dans une variable la sortie standard du résultat de la fonction en type string.
	txtData, err := ReadQuadTxt()    // Ouverture du fichier quad-x-y.txt généré par l'executable quad[A-E] pour transmettre les coordonnées x et y à quadchecker.
	if err != nil {
		fmt.Println("Not a quad function") // Si le fichier n'est pas trouvé, c'est que ça n'est pas un quad[A-E]
		return
	}
	fmt.Println(originalOutput)                                          // Print de la sortie du quad. ex: "o---o"
	x, y := XYExtract(txtData)                                           // Extraire les valeurs de x et y à partir des chaînes de caractères formatees avec %!s(int=...) + conversion string to int(Atoi).
	executables := []string{"quadA", "quadB", "quadC", "quadD", "quadE"} // Definir les noms des executables
	ExecAllQuads(x, y, executables, originalOutput)                      // Executer les autres executables et comparer la sortie
	os.Remove("quad-x-y.txt")                                            // Supprimer le fichier quad-x-y.txt
}

func OutputOfExec() string {
	input, _ := ioutil.ReadAll(os.Stdin)
	return string(input)
}

func ReadQuadTxt() ([]byte, error) {
	txtData, err := ioutil.ReadFile("quad-x-y.txt")
	return txtData, err
}

func XYExtract(txtData []byte) (string, string) {
	x := strings.Trim(string(txtData[:len(txtData)/2]), "%!s(int=) ")   // extrait x de la chaîne "%!s(int=x) "
	y := strings.Trim(string(txtData[len(txtData)/2:]), "%!s(int=) \n") // extrait y de la chaîne "%!s(int=y) "
	return x, y
}

func ExecAllQuads(x, y string, executables []string, originalOutput string) {
	for _, execName := range executables {
		// Executer l'executable avec les valeurs de x et y extraites
		cmd := exec.Command("./"+execName, x, y)
		var out bytes.Buffer
		cmd.Stdout = &out
		cmd.Run()
		output := out.String()
		// Comparer la sortie avec la sortie cible en utilisant strings.EqualFold
		if strings.EqualFold(originalOutput, output) {
			if !isStarting {
				fmt.Printf("[%s] [%s] [%s]", execName, x, y)
				isStarting = true
			} else {
				fmt.Printf(" || [%s] [%s] [%s]", execName, x, y)
			}
		}
	}
	fmt.Println() // Saut de ligne.
}
