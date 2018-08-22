package main

import "fmt"

/*
Trying to use Something() in another file gives this error:
GOROOT=C:\Go #gosetup
GOPATH=C:\Users\condo\go #gosetup
C:\Go\bin\go.exe build -i -o C:\Users\condo\AppData\Local\Temp\___go_build_importswin_go.exe C:/Users/condo/go/src/goven/biznatch/importswin.go #gosetup
# command-line-arguments
biznatch\importswin.go:19:8: undefined: Something

Compilation finished with exit code 2


https://stackoverflow.com/questions/41365400/how-to-run-the-whole-project-in-gogland
https://stackoverflow.com/questions/28153203/golang-undefined-function-declared-in-another-file

These do not owrk:
package: goven\biznatch
package: C:\Users\condo\go\src\goven\biznatch

directory: C:/Users/condo/go/src/goven/
directory: C:\Users\condo\go\src\goven
directory: C:\Users\condo\go\src\goven\biznatch
runnerw.exe: CreateProcess failed with error 216: This version of %1 is not compatible with the version of Windows you're running. Check your computer's system information and then contact the software publisher.

*/
func main() {
	fmt.Println("Importing something.")

	fmt.Println("if it's in the same package, just go.")

	s1 := Something{"wicked this way"}

	fmt.Println(s1.name)
}
