package main

import (
	"fmt"
	"module-8-practice/electronic"
)

func main() {
	i := electronic.NewApplePhone("", "Iphone 14", "", "")
	//a := electronic.NewAndroidPhone("Samsung", "A70", "", "Android")
	//r := electronic.NewRadioPhone("Alcatel", "XP-55", "station", 12)
	printCharacteristics(&i)
	//printCharacteristics(&a)
}

func printCharacteristics(p electronic.Phone) {
	fmt.Printf("Бренд:%v\nМодель:%v\nТип:%v\n", p.Brand(), p.Model(), p.Type())
}

func pr(s electronic.Smartphone) {

}
