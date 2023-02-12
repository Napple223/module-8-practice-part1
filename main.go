package main

import (
	"fmt"
	"module-8-practice/electronic"
)

func main() {
	i := electronic.NewApplePhone("", "Iphone 14", "", "")
	a := electronic.NewAndroidPhone("Samsung", "A70", "", "Android")
	r := electronic.NewRadioPhone("Alcatel", "XP-55", "station", 12)
	printCharacteristics(&i)
	pc1(&i)
	printCharacteristics(&a)
	pc1(&a)
	printCharacteristics(&r)
	pc2(&r)
}

func printCharacteristics(p electronic.Phone) {
	fmt.Printf("Бренд:%v\nМодель:%v\nТип:%v\n", p.Brand(), p.Model(), p.Type())
}

func pc1(s electronic.Smartphone) {
	fmt.Printf("Операционная система:%s\n", s.OS())
}

func pc2(sp electronic.StationPhone) {
	fmt.Printf("Количество кнопок:%v\n", sp.ButtonsCount())
}
