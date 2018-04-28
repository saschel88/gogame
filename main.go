package main

import (

	"fmt"
	"homework/gogame/newtypes"
)


func main() {

	pole := [figures.RazmernostPole][figures.RazmernostPole]int{} //Наше поле
	//имитация значений полученных от пользователя
	//sideTop:=4
	pole[5][8]=1
	x := 4
	y := 7
	sideSquare := 4
	cellGoTo := 0
	name := "Square1"
	mapStructName := make(map[string]figures.Painter) //карта

	if _, ok := mapStructName[name]; ok {

	} else {
		if x < 0 || y < 0 || x > figures.RazmernostPole || y > figures.RazmernostPole { //Проверка на принадлежность координат первой ячейки нашему Полю
			fmt.Println("Начальные координаты находятся вне Поля") //Доработать выход для к меню
		} else {
			newKvadrat := new(figures.Square) //Инициализация указателя на Структуру квадрата с дыркой
			newKvadrat.Pole = &pole   //передача адреса Нашего поля в объект структуры квадрат с дыркой
			newKvadrat.Name = name
			newKvadrat.SideSquare = sideSquare

			err := newKvadrat.Create(x, y, cellGoTo, ok)
			if err != nil {
				fmt.Println(err)
			}
		}

		for _, v := range pole { //Вывод Нашего поля на экран
			fmt.Println(v)
		}
	}
	/*n:=10
	array:=make([][]int,n)
	arrayOfKoordinat:=[2]int{}
	array[0]=append(array[0],arrayOfKoordinat[0],arrayOfKoordinat[1])

	fmt.Println(array[0][1])*/

}
