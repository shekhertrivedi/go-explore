package cdc

import (
	"fmt"
)

func FuncFromOtherPackage() {
	fmt.Println("Printing from other package")
	funcFromOtherPackage1()
}

func funcFromOtherPackage1() {
	fmt.Println("Printing from other package another function")
}
