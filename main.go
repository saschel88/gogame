package main

import (

	"fmt"
	"gogame/figures"
	"gogame/textconstants"

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
	//Имя Фигуры
	figureName:=""

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
	// Направление перемещения
	goTo :=0
	// Количество ячеек на которое необходимо перевезти фигуру
	howFarToGo:=0

	// Цикл основной программы
	for !GoToExit {
		GoToExit=false
		GoToMainMenu = false
		fmt.Println(textconstants.MainMenu)
		fmt.Print(textconstants.EnterTheNumber)
		fmt.Scanf("%d", &userResponse)

		switch userResponse {
		// Вывод поля
		case 1:
			showPole(&pole)
		//Создать фигуру
		case 2:

			for !GoToMainMenu {

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
					if squareSide<2{
						fmt.Println("Квадрат с стороной", squareSide, " построть невозвможно")
						break
					}
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

					// Ввод высоты прямоугольника
					fmt.Print(textconstants.RectangleHighLenightInput)
					fmt.Println(textconstants.EnterTheNumber)
					fmt.Scanf("%d", &rectangleHigh)
					// Ввод Ширины прямоугольника
					fmt.Print(textconstants.RectangleWidthLenightInput)
					fmt.Println(textconstants.EnterTheNumber)
					fmt.Scanf("%d", &rectangleWidht)
					if (rectangleHigh<2 && rectangleWidht<2)||rectangleHigh<0||rectangleWidht<0{
						fmt.Println("Данный прямоугольник построть невозвможно")
						break
					}
					// Ввод координат
					fmt.Println(textconstants.CoordinateFigureX)
					fmt.Scanf("%d", &coordinateX)
					fmt.Println(textconstants.CoordinateFigureY)
					fmt.Scanf("%d", &coordinateY)
					newRectangle:=new(figures.Rectangle)
					if err:=newRectangle.Create(rectangleName,mapStructName,&pole,rectangleHigh,rectangleWidht,coordinateX,coordinateY);err!=nil{
						fmt.Println(err)
					}
					//Вывод поля
					showPole(&pole)//Вывод поля

					GoToMainMenu=true
				case 3: //Змейка
					snakeName=""
					fmt.Print(textconstants.SnakeFigureHeader)
					fmt.Println(textconstants.NameFigure)
					fmt.Scanf("%s",&snakeName)
					// Ввод длины Змейки
					fmt.Print(textconstants.SnakeLenightInput)
					fmt.Println(textconstants.EnterTheNumber)
					fmt.Scanf("%d", &snakeLengh)
					if snakeLengh<0{
						fmt.Println("Данную  Змейку невозвможно")
						break
					}
					// Ввод координат
					fmt.Println(textconstants.CoordinateFigureX)
					fmt.Scanf("%d", &coordinateX)
					fmt.Println(textconstants.CoordinateFigureY)
					fmt.Scanf("%d", &coordinateY)
					newSnake:=new(figures.Snake)
					if err:=newSnake.Create(snakeName,mapStructName,&pole,snakeLengh,coordinateX,coordinateY);err!=nil{
						fmt.Println(err)
					}
					//Вывод поля
					showPole(&pole)//Вывод поля

					GoToMainMenu=true

				case 0: //Выход в Главное меню
					GoToMainMenu = true

				default: //Переход к меню Выбор типа фигур
					fmt.Println(textconstants.DefaultValue)
					continue
				}
			}

		case 3:// Удаление фигуры
			figureName=""
			fmt.Print(textconstants.DeleteFigureHeader)
			fmt.Println(textconstants.NameFigure)
			fmt.Scanf("%s", &figureName)
			if _,ok:=mapStructName[figureName]; !ok { //Проверка на существовынии указанной фтгуры
				fmt.Println("Фигура с именем", figureName, " не существует")
				continue
			}
			if err:=mapStructName[figureName].Delete(true);err!=nil{
				fmt.Println(err)
			}
		case 4:// Перемещение фигуры
			figureName=""
			goTo=0
			howFarToGo=0
			fmt.Print(textconstants.MoveFigureHeader)
			fmt.Println(textconstants.NameFigure)
			fmt.Scanf("%s", &figureName)
			if _,ok:=mapStructName[figureName]; !ok { //Проверка на существовынии указанной фтгуры
				fmt.Println("Фигура с именем", figureName, " не существует")
				continue
			}

			//вывод вариантов направления движения
			fmt.Print(textconstants.MoveFigureDirection)
			fmt.Print(textconstants.EnterTheNumber)
			fmt.Scanf("%d", &goTo)
			fmt.Print(textconstants.MoveFigureValue)
			fmt.Print(textconstants.EnterTheNumber)
			fmt.Scanf("%d", &howFarToGo)
			if err:=mapStructName[figureName].Move(goTo,howFarToGo); err!=nil{
				fmt.Println(err)

			}
			showPole(&pole)

		case 5:
			figureName=""
			fmt.Print(textconstants.TakeFigureCoordinate)
			fmt.Println(textconstants.NameFigure)
			fmt.Scanf("%s", &figureName)
			if _,ok:=mapStructName[figureName]; !ok { //Проверка на существовынии указанной фтгуры
				fmt.Println("Фигура с именем", figureName, " не существует")
				continue
			}
			mapStructName[figureName].TakeKoordinate()
		case 0:
			fmt.Println(textconstants.GoodBay)
			GoToExit = true
		default:
			fmt.Println(textconstants.DefaultValue)

		}

	}

}


