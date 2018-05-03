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
				snakeLengh=0
				squareSide=0
				rectangleHigh=0
				rectangleWidht=0
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


