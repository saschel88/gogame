package main

import (
	"gogame/newtypes"
	"fmt"
)

func main() {


	pole := [figures.RazmernostPole][figures.RazmernostPole]int{} //Наше поле
	//имитация значений полученных от пользователя
	//sideTop:=4
	//pole[5][8]=1
	x := 5
	y := 3
	sideSquare := 4
	sideWidht := 3
	sideHeight := 4
	lenightSnake:=5
	goTo:=1
	howFarCount:=3
	name := "Square1"
	name2 := "Rectangle1"
	name3 := "Snake1"
	mapStructName := make(map[string]figures.Painter) //карта

	//Типы фигуо: 1-Квадрат с дыркой
	//			  2-Прямоугольник
	//			  3-Змейка
	figureType := 3

	if _, ok := mapStructName[name2]; ok {
		// Тут описывается движение фогуры
	} else {
		if x < 0 || y < 0 || x > figures.RazmernostPole || y > figures.RazmernostPole { //Проверка на принадлежность координат первой ячейки нашему Полю
			fmt.Println("Начальные координаты находятся вне Поля") //Доработать выход для к меню
		} else {
			if figureType == 1 {

				newKvadrat := new(figures.Square) //Инициализация указателя на Структуру квадрата с дыркой
				newKvadrat.Pole = &pole           //передача адреса Нашего поля в объект структуры квадрат с дыркой
				newKvadrat.Name = name
				newKvadrat.SideSquare = sideSquare

				err := newKvadrat.Create(x, y)
				if err != nil {
					fmt.Println(err)
				}
				mapStructName[newKvadrat.Name] = newKvadrat

				for _, v := range pole { //Вывод Нашего поля на экран
					fmt.Println(v)
				}

				//Move()
				err=newKvadrat.Move(goTo,howFarCount)
				if err!=nil {
					fmt.Println(err)
				}
				fmt.Println("************************************************")
				for _, v := range pole { //Вывод Нашего поля на экран
					fmt.Println(v)
				}

				err=mapStructName[name].Delete()
				if err != nil {
					fmt.Println(err)
				} else {
					delete(mapStructName, newKvadrat.Name)
				}
				fmt.Println("**********************************************")
			} else if figureType == 2 {
				fmt.Println(2)
				newRectangle := new(figures.Rectangle) //Инициализация указателя на Структуру квадрата с дыркой
				newRectangle.Pole = &pole              //передача адреса Нашего поля в объект структуры квадрат с дыркой
				newRectangle.Name = name2
				newRectangle.Width = sideWidht
				newRectangle.Height = sideHeight
				err := newRectangle.Create(x, y)
				if err != nil {
					fmt.Println(err)
				}
				mapStructName[newRectangle.Name] = newRectangle

				for _, v := range pole { //Вывод Нашего поля на экран
					fmt.Println(v)
				}

				//Move()
				err=mapStructName[newRectangle.Name].Move(goTo,howFarCount)
				if err != nil {
					fmt.Println(err)
				}
				fmt.Println("************************************************")
				for _, v := range pole { //Вывод Нашего поля на экран
					fmt.Println(v)
				}
					fmt.Println("************************************************")
				err = newRectangle.Delete()
				if err != nil {
					fmt.Println(err)
				} else {
					delete(mapStructName, newRectangle.Name)
				}
				fmt.Println("**********************************************")
				// Для прямоугольника
			} else if figureType == 3 {
				fmt.Println(3)
				newSnake := new(figures.Snake) //Инициализация указателя на Структуру квадрата с дыркой
				newSnake.Pole = &pole              //передача адреса Нашего поля в объект структуры квадрат с дыркой
				newSnake.Name = name3
				newSnake.SnakeLength = lenightSnake


				err := newSnake.Create(x, y)
				if err != nil {
					fmt.Println(err)
				}
				mapStructName[newSnake.Name] = newSnake

				for _, v := range pole { //Вывод Нашего поля на экран
					fmt.Println(v)
				}

				fmt.Println("*************************************")
				err=newSnake.Move(goTo,howFarCount)
				if err != nil {
					fmt.Println(err)
				}
				for _, v := range pole { //Вывод Нашего поля на экран
					fmt.Println(v)
				}
				fmt.Println("********************************************")
				err = mapStructName[name3].Delete()
				if err != nil {
					fmt.Println(err)
				} else {
					delete(mapStructName, newSnake.Name)
				}
				fmt.Println("**********************************************")

			}

		}

		for _, v := range pole { //Вывод Нашего поля на экран
			fmt.Println(v)
		}

	}




}
