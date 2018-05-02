package main

import (

	"fmt"
	"homework/gogame/figures"
	"homework/gogame/textconstants"

)

func showPole(pole *[figures.RazmernostPole][figures.RazmernostPole]int)  {
	for _, v := range pole { //Вывод Нашего поля на экран
		fmt.Println(v)
	}
}

func main() {
	mapStructName := make(map[string]figures.Painter) //карта
	pole := [figures.RazmernostPole][figures.RazmernostPole]int{}
/*	pole := [figures.RazmernostPole][figures.RazmernostPole]int{} //Наше поле
	//имитация значений полученных от пользователя
	//sideTop:=4
	//pole[5][8]=1
	x := 5
	y := 3
	sideSquare := 4
	sideWidht := 3
	sideHeight := 4
	lenightSnake:=5
	goTo:=2
	howFarCount:=4
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
				if err= mapStructName[name].TakeKoordinate();err!=nil {
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
				if err= mapStructName[name2].TakeKoordinate();err!=nil {
					fmt.Println(err)
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


				err := newSnake.Create(10, 12)
				if err != nil {
					fmt.Println(err)
				}
				mapStructName[newSnake.Name] = newSnake

				newKvadrat := new(figures.Square) //Инициализация указателя на Структуру квадрата с дыркой
				newKvadrat.Pole = &pole           //передача адреса Нашего поля в объект структуры квадрат с дыркой
				newKvadrat.Name = name
				newKvadrat.SideSquare = sideSquare

				err = newKvadrat.Create(1, 1)
				if err != nil {
					fmt.Println(err)
				}
				mapStructName[newKvadrat.Name] = newKvadrat

				for _, v := range pole { //Вывод Нашего поля на экран
					fmt.Println(v)
				}

				fmt.Println("*************************************")
				//err=mapStructName[name].Move(2,5)
				//if err != nil {
				//	fmt.Println(err)
				//}

				fmt.Println("*************************************")
				err=mapStructName[name3].Move(1,6)
				if err != nil {
					fmt.Println(err)
				}
				fmt.Println("*************************************")
				for _, v := range pole { //Вывод Нашего поля на экран
					fmt.Println(v)
				}
				err=mapStructName[name].Move(3,4)
				if err != nil {
					fmt.Println(err)
				}
				err=mapStructName[name3].Move(4,4)

				if err != nil {
					fmt.Println(err)
				}
				fmt.Println("*************************************")
				/*for _, v := range pole { //Вывод Нашего поля на экран
					fmt.Println(v)
				}*/
				/*if err= mapStructName[name3].TakeKoordinate();err!=nil {
					fmt.Println(err)
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
*/

	//Переход в Главное меню
	GoToMainMenu := false
	// Выход из программы
	GoToExit := false
	//Ответ пользователя
	var userResponse int
	//Имя квадрата
	squareName:=""
	//Сторона квадрата
	squareSide:=0
	//Имя прямоугольника
	rectangleName:=""

	//Высота прямоугольника
	rectangleHigh:=0
	//Ширина прямоугольника
	rectangleWidht:=0
	//Имя Змейки
	snakeName:=""
	//Длина Змейки
	snakeLengh:=0
	//Координаты
	coordinateX:=0
	coordinateY:=0

	// Цикл основной программы
	for !GoToExit {
		GoToExit=false

		fmt.Println(textconstants.MainMenu)
		fmt.Print(textconstants.EnterTheNumber)
		fmt.Scanf("%d", &userResponse)

		switch userResponse {
		// Вывод поля
		case 1:
		//Создать фигуру
		case 2:

			for !GoToMainMenu {
				GoToMainMenu = false
				fmt.Println(textconstants.CreateFigureHeader)
				fmt.Print(textconstants.EnterTheNumber)
				fmt.Scanf("%d", &userResponse)
				coordinateY=0
				coordinateX=0
				switch userResponse {
				case 1: //Пустой квадрат
					squareName=""
					fmt.Print(textconstants.SquareFigureHeader)
					fmt.Print(textconstants.NameFigure)
					fmt.Scanf("%s",&squareName)

					// Проверка на наличие карты или Create()
					fmt.Print(textconstants.SquareSideLenightInput)
					fmt.Print(textconstants.EnterTheNumber)
					fmt.Scanf("%d", &squareSide)
					fmt.Print(textconstants.CoordinateFigureX)
					fmt.Scanf("%d", &coordinateX)
					fmt.Println(textconstants.CoordinateFigureY)
					fmt.Scanf("%d", &coordinateY)
					newSquare:=new(figures.Square)
					if err:=newSquare.Create(squareName,mapStructName,&pole,squareSide,coordinateX,coordinateY);err!=nil{
						fmt.Println(err)
					}//Create()

					showPole(&pole)//Вывод поля

					GoToMainMenu=true
				case 2: //Прямоугольник
					rectangleName=""
					fmt.Println(textconstants.RectangleFigureHeader)
					fmt.Println(textconstants.NameFigure)
					fmt.Scanf("%s",&rectangleName)
					// Проверка на наличие карты или Create()
					// Ввод высоты прямоугольника
					fmt.Println(textconstants.RectangleHighLenightInput)
					fmt.Println(textconstants.EnterTheNumber)
					fmt.Scanf("%d", &rectangleHigh)
					// Ввод Ширины прямоугольника
					fmt.Println(textconstants.RectangleWidthLenightInput)
					fmt.Println(textconstants.EnterTheNumber)
					fmt.Scanf("%d", &rectangleWidht)
					// Ввод координат
					fmt.Println(textconstants.CoordinateFigureX)
					fmt.Scanf("%d", &coordinateX)
					fmt.Println(textconstants.CoordinateFigureY)
					fmt.Scanf("%d", &coordinateY)
					//Create()
					//Вывод поля
				case 3: //Змейка
					snakeName=""
					fmt.Println(textconstants.SnakeFigureHeader)
					fmt.Println(textconstants.NameFigure)
					fmt.Scanf("%s",&snakeName)
					// Проверка на наличие карты или Create()
					// Ввод длины Змейки
					fmt.Println(textconstants.SnakeLenightInput)
					fmt.Println(textconstants.EnterTheNumber)
					fmt.Scanf("%d", &snakeLengh)
					// Ввод координат
					fmt.Println(textconstants.CoordinateFigureX)
					fmt.Scanf("%d", &coordinateX)
					fmt.Println(textconstants.CoordinateFigureY)
					fmt.Scanf("%d", &coordinateY)
					//Create()
					//Вывод поля
					GoToMainMenu=true

				case 0: //Выход в Главное меню
					GoToMainMenu = true

				default: //Переход к меню Выбор типа фигур
					fmt.Println(textconstants.DefaultValue)
					continue
				}
			}

		case 3:
			fmt.Println(textconstants.DeleteFigureHeader)
			fmt.Println(textconstants.NameFigure)
			fmt.Print(textconstants.EnterTheNumber)
			fmt.Scanf("%d", &userResponse)
		case 4:
			fmt.Println(textconstants.MoveFigureHeader)
			fmt.Println(textconstants.NameFigure)
			fmt.Print(textconstants.EnterTheNumber)
			fmt.Scanf("%d", &userResponse)
		case 5:
			fmt.Println(textconstants.TakeFigureCoordinate)
			fmt.Println(textconstants.NameFigure)
			fmt.Print(textconstants.EnterTheNumber)
			fmt.Scanf("%d", &userResponse)
		case 0:
			fmt.Println(textconstants.GoodBay)
			GoToExit = true
		default:
			fmt.Println(textconstants.DefaultValue)

		}

	}

}


