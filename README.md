# QuadChecker

## Procédure d'installation.
Pour build `quadA.go`, `quadB.go`, `quadC.go`, `quadD.go`, `quadE.go` & `quadchecker.go` executer le MakeFile :

```bash
sh MakeFile
```

## Executer la démo.

```bash
sh Demo
```

## Main

### isStarting

`isStarting` est une variable booléenne globale permettant de gérer les séparateurs "||" dans le résultat, de sorte à éviter d'afficher "||" sur la première sortie.

Elle est initialisée sur `false` et devient `true` à partir de la deuxième sortie.

```go
var isStarting bool = false
```
### OutputOfExec

Met dans la variable `originalOutput` la sortie standard du terminal résultante de la fonction quadA-E en type string.

*Exemple:*

```bash
./quadA 5 6
o---o
|   |
|   |
|   |
|   |
o---o
```
`main`
```go
func main() {
	originalOutput := OutputOfExec()
```
`OutputOfExec`
```go
func OutputOfExec() string {
	input, _ := ioutil.ReadAll(os.Stdin)
	return string(input)
}
```
### ReadQuadTxt

Chacune des fonctions quadA-E génère un fichier texte temporaire contenant les valeurs de `x` et `y` demandées par la ligne de commande via `quadchecker`

`ReadQuadTxt` ouvre donc le fichier en question, s'il n'est pas capable de l'ouvrir alors ça n'est pas une fonction quad.

*"Not a quad function"*

`main`
```go
	txtData, err := ReadQuadTxt()
	if err != nil {
		fmt.Println("Not a quad function")
		return
	}
	fmt.Println(originalOutput)
```
`ReadQuadTxt`
```go
func ReadQuadTxt() ([]byte, error) {
	txtData, err := ioutil.ReadFile("quad-x-y.txt")
	return txtData, err
}
```
### XYExtract

Le fichier texte temporaire apparaissant sous cette forme :

```txt
%!s(int=5) %!s(int=6)
```
La fonction `XYExtract` permet à la fois d'extraire `x` et `y` de cette chaîne de caractère mais aussi de les convertir en int. Avec la fonction `Atoi`.

`main`
```go
	xInt, yInt := XYExtract(txtData)
```
`XYExtract`
```go
func XYExtract(txtData []byte) (int, int) {
	xStr := strings.Trim(string(txtData[:len(txtData)/2]), "%!s(int=) ")
	xInt, _ := strconv.Atoi(xStr)
	yStr := strings.Trim(string(txtData[len(txtData)/2:]), "%!s(int=) \n")
	yInt, _ := strconv.Atoi(yStr)
	return xInt, yInt
}
```
### executables(variable)

Cette chaîne de string contient tous les quadA-E executables.

`main`
```go
	executables := []string{"quadA", "quadB", "quadC", "quadD", "quadE"}
```

### ExecAllQuads

Cette fonction prend en argument `xInt` `yInt` `executables` `originalOutput` via une boucle `for` execute avec les mêmes données `x` et `y` toutes les fonctions quadA-E

`main`
```go
	ExecAllQuads(xInt, yInt, executables, originalOutput)
```

On met dans la variable `cmd` la commande de terminal que l'on veut executer :
```go
"./"+execName, strconv.Itoa(xInt), strconv.Itoa(yInt)
```
Équivalent à taper dans le terminal *./quadA-E x y*.
avec `cmd.Run()` la ligne de commande est exécutée pour chaque fonction quadA-E jusqu'à la fin de la boucle et stockée à chaque itération dans la variable `output`.

À chaque itération la fonction [strings.EqualFold](https://pkg.go.dev/strings#EqualFold) compare la sortie standard de la fonction testée au départ avec la fonction testée durant cette itération.

Si elles sont identiques alors elle les imprime.

`ExecAllQuads`

```go
func ExecAllQuads(xInt, yInt int, executables []string, originalOutput string) {
	for _, execName := range executables {
		cmd := exec.Command("./"+execName, strconv.Itoa(xInt), strconv.Itoa(yInt))
		var out bytes.Buffer
		cmd.Stdout = &out
		cmd.Run()
		output := out.String()
		if strings.EqualFold(originalOutput, output) {
			if !isStarting {
				fmt.Printf("[%s] [%d] [%d]", execName, xInt, yInt)
				isStarting = true
			} else {
				fmt.Printf(" || [%s] [%d] [%d]", execName, xInt, yInt)
			}
		}
	}
	fmt.Println()
}
```
### os.Remove

Supprime le fichier temporaire.

`main`
```go
	os.Remove("quad-x-y.txt")                                            // Supprimer le fichier quad-x-y.txt
}
```