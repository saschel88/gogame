package main

import (

	"fmt"
	"homework/gogame/newtypes"
)


func main() {

	pole := [figures.RazmernostPole][figures.RazmernostPole]int{} //Наше поле
	//имитация значений полученных от пользователя
	//sideTop:=4
	//pole[5][8]=1
	x := 14
	y := 0
	sideSquare := 4
	//cellGoTo := 0
	name := "Square1"
	mapStructName := make(map[string]figures.Painter) //карта

	//Типы фигуо: 1-Квадрат с дыркой
	//			  2-Прямоугольник
	//			  3-Змейка
	figureType:=2


	if _, ok := mapStructName[name]; ok {
			// Тут описывается движение фогуры
	} else {
		if x < 0 || y < 0 || x > figures.RazmernostPole || y > figures.RazmernostPole { //Проверка на принадлежность координат первой ячейки нашему Полю
			fmt.Println("Начальные координаты находятся вне Поля") //Доработать выход для к меню
		} else {
			if figureType==1 {


			newKvadrat := new(figures.Square) //Инициализация указателя на Структуру квадрата с дыркой
			newKvadrat.Pole = &pole   //передача адреса Нашего поля в объект структуры квадрат с дыркой
			newKvadrat.Name = name
			newKvadrat.SideSquare = sideSquare

			err := newKvadrat.Create(x, y)
			if err != nil {
				fmt.Println(err)
			}
			mapStructName[newKvadrat.Name]=newKvadrat

			for _, v := range pole { //Вывод Нашего поля на экран
				fmt.Println(v)
			}
			newKvadrat.Delete()
			fmt.Println("**********************************************")
			} else if figureType==2{
				fmt.Println(2)
				// Для прямоугольника
			}else if figureType==3{
				//Для Змейки
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
