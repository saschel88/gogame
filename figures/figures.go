// Опиание структур используемых фигур и интерфейса, реализуемого данными структурами
package figures

import (
	"errors"
	"fmt"
	"container/list"
	"github.com/astaxie/beego/logs"
)

const RazmernostPole = 30 //размерность отображаемого поля
// Интерфейс рописывающий дефтельность фигур на поле
type Painter interface {

	//Метод создания фигуры
	//args[0]: Сторона квадрата, Высота прямоугольника, Длина змейки
	//args[1]: Ширина прямоугольника, Координата Х
	//args[2]:  Координата Y
	// pole - указатель на используемое поле
	Create(name string, mapPainter map[string]Painter, pole *[RazmernostPole][RazmernostPole]int, args ...int) error
	// Координаты имеют вид array[X][Y]
	//coordinateX - координата по X
	//coordinateY - координата по Y
	Paint(coordinateX, coordinateY int) error
	// Метод удаления. Возвращаему ошибку взять на контроль (нужна ли она)
	Delete(isFatalDelete bool) error
	// Метод перемещения фигуры.
	// Направления движения 1: Вверх
	//						2: Вниз
	//						3: Вправо
	//						4: Влево
	Move(goTo, howFarToGo int) error
	//Метод получения координат
	TakeKoordinate ()error

}

// Описание фигуры Квадрат с дырой
type Square struct {
	Name        string
	Pole        *[RazmernostPole][RazmernostPole]int
	Coordinates [][2]int
	SideSquare  int //Значение стороны квадрата. Получаем при создании
	MapPainter map[string]Painter

}

//Метод создания фигуры квадрат с дыркой
func (s *Square) Create(name string, mapPainter map[string]Painter, pole *[RazmernostPole][RazmernostPole]int, args ...int) error{
	if _,ok :=mapPainter[name]; ok {
		return errors.New("Фигура с таким именем уже существует")
	}

	cellCountInArray := args[0]*4 - 4    //Подсчет количества ячеек в массиве координатов для квадрата с дыркой
	array := make([][2]int, cellCountInArray) // Инициализация слайса  массивов координат
	coordinateX:=args[1]
	coordinateY:=args[2]
	topLine := args[0]+ coordinateY     //Получаем длину верхней грани
	lateralLine := args[0] + coordinateX //Получаем длину боковой грани
	if coordinateX < 0 || coordinateY < 0 || coordinateX > RazmernostPole || coordinateY > RazmernostPole { //Проверка на принадлежность координат первой ячейки нашему Полю
		return errors.New("Выход начальных координат за пределы Поля") //Проверка на нахождение координат в пределах Поля
	}
	if topLine < 0 || lateralLine < 0 || topLine > RazmernostPole || lateralLine > RazmernostPole { //Проверка на принадлежность координат первой ячейки нашему Полю
		return errors.New("Выход пустого квадрата  за пределы Поля") //Проверка на нахождение координат в пределах Поля
	}
	s.Pole=pole
	tempPole := *s.Pole                       // Получение слепка Поля
	inbexCoordinateArray := 0 //Идекс в массиве координат
	for i := coordinateX; i < lateralLine; i++ {
		for j := coordinateY; j < topLine; j++ {
			if tempPole[i][j] != 0 {
				return errors.New("По данным координатам фигуру построить невозможно, есть препядствия")
			}
			if i == coordinateX || i == lateralLine-1 || j == coordinateY || j == topLine-1 {
				tempPole[i][j] = 1
				array[inbexCoordinateArray][0] = i //Заполнение значений координат Х
				array[inbexCoordinateArray][1] = j //Заполнение значений координат Y
				inbexCoordinateArray++

			}
		}
	}

	s.Name=name
	s.SideSquare=args[0]
	*s.Pole=tempPole
	s.Coordinates=array
	mapPainter[name]=s
	s.MapPainter=mapPainter

	return nil
}


//Метод отображения  фигуры квадрат с дыркой
func (s *Square) Paint(coordinateX, coordinateY int) error {


	if coordinateX < 0 || coordinateY < 0 || coordinateX > RazmernostPole || coordinateY > RazmernostPole { //Проверка на принадлежность координат первой ячейки нашему Полю
		return errors.New("Выход за пределы Поля") //Проверка на нахождение координат в пределах Поля
	}

	tempPole := *s.Pole                       // Получение слепка Поля
	array := s.Coordinates    //Подсчет количества ячеек в массиве координатов для квадрата с дыркой
	topLine := s.SideSquare + coordinateY     //Получаем длину верхней грани
	lateralLine := s.SideSquare + coordinateX //Получаем длину боковой грани
	if topLine < 0 || lateralLine < 0 || topLine > RazmernostPole || lateralLine > RazmernostPole { //Проверка на принадлежность координат первой ячейки нашему Полю
		return errors.New("Выход за пределы Поля") //Проверка на нахождение координат в пределах Поля
	}
	// Проверка на наличие препядствий, мешающих построить квадрат , рисование квадрата с дыркой на Поле и запись координат в контейнер координат
	inbexCoordinateArray := 0 //Идекс в массиве координат
	for i := coordinateX; i < lateralLine; i++ {
		for j := coordinateY; j < topLine; j++ {
			if tempPole[i][j] != 0 {
				return errors.New("По данным координатам фигуру построить невозможно, есть препядствия")
			}
			if i == coordinateX || i == lateralLine-1 || j == coordinateY || j == topLine-1 {
				tempPole[i][j] = 1
				array[inbexCoordinateArray][0] = i //Заполнение значений координат Х
				array[inbexCoordinateArray][1] = j //Заполнение значений координат Y
				inbexCoordinateArray++

			}
		}
	}
	err := s.Delete(false)
	if err != nil {
		return err
	}
	s.Coordinates = array
	*s.Pole = tempPole
	return nil
}

// Удаление имеющегося квадрата с дыркой
func (s *Square) Delete(isFatalDelete bool) error {

	tempPole := *s.Pole               // Получение слепка Поля
	for _, v := range s.Coordinates { //Производим замену по координатам.
		tempPole[v[0]][v[1]] = 0
	}
	*s.Pole = tempPole
	if isFatalDelete {
		delete(s.MapPainter,s.Name)
	}
	return nil
}

//Движение квадрата с дыркой
// Направления движения 1: Вверх
//						2: Вниз
//						3: Вправо
//						4: Влево
func (s *Square) Move(goTo, howFarToGo int) error {

	switch goTo {
	case 1: //Вверх
		err := s.Paint(s.Coordinates[0][0]-howFarToGo, s.Coordinates[0][1])
		if err != nil {
			return err
		}
		return nil
	case 2: //Вниз
		err := s.Paint(s.Coordinates[0][0]+howFarToGo, s.Coordinates[0][1])
		if err != nil {
			return err
		}
		return nil
	case 3: //Вправо
		err := s.Paint(s.Coordinates[0][0], s.Coordinates[0][1]+howFarToGo)
		if err != nil {
			return err
		}
		return nil
	case 4: //Влево
		err := s.Paint(s.Coordinates[0][0], s.Coordinates[0][1]-howFarToGo)
		if err != nil {
			return err
		}
		return nil
	default:
		return errors.New("Указанное направление для движения отсутствует")
	}
}
//Методы вывода координа квадраты с дыркой
func (s *Square) TakeKoordinate () error {
	for k,v := range s.Coordinates{
		fmt.Print(" ",k,"=",v)
	}
	fmt.Println("")
	return nil
}

// Описание фигуры Прямоугольник.

type Rectangle struct {
	Name        string
	Pole        *[RazmernostPole][RazmernostPole]int
	Coordinates [][2]int
	Height      int //Значение ширины прямоугольника. Получаем при создании
	Width       int //Значение высоты прямоугольника. Получаем при создании
	MapPainter map[string]Painter
}
//Метод создания фигуры Прямоугольник
//args[0]: Высота прямоугольника.
//args[1]: Ширина прямоугольника.
//args[2]: Координата Х
//args[3]: Координата Y
// pole - указатель на используемое поле
//Метод создания фигуры прямоугольник
func (r *Rectangle) Create(name string, mapPainter map[string]Painter, pole *[RazmernostPole][RazmernostPole]int, args ...int) error{
	if _,ok :=mapPainter[name]; ok {
		return errors.New("Фигура с таким именем уже существует")
	}

	cellCountInArray := args[0]*args[1]    //Подсчет количества ячеек в массиве координатов для квадрата с дыркой
	array := make([][2]int, cellCountInArray) // Инициализация слайса  массивов координат
	coordinateX:=args[2]
	coordinateY:=args[3]
	topLine := args[1]+ coordinateY     //Получаем длину верхней грани
	lateralLine := args[0] + coordinateX //Получаем длину боковой грани
	if coordinateX < 0 || coordinateY < 0 || coordinateX > RazmernostPole || coordinateY > RazmernostPole { //Проверка на принадлежность координат первой ячейки нашему Полю
		return errors.New("Выход начальных координат за пределы Поля") //Проверка на нахождение координат в пределах Поля
	}
	if topLine < 0 || lateralLine < 0 || topLine > RazmernostPole || lateralLine > RazmernostPole { //Проверка на принадлежность координат первой ячейки нашему Полю
		return errors.New("Выход пустого квадрата  за пределы Поля") //Проверка на нахождение координат в пределах Поля
	}
	r.Pole=pole
	tempPole := *r.Pole                       // Получение слепка Поля
	inbexCoordinateArray := 0 //Идекс в массиве координат

	for i := coordinateX; i < lateralLine; i++ {
		for j := coordinateY; j < topLine; j++ {
			if tempPole[i][j] != 0 {
				return errors.New("По данным координатам фигуру построить невозможно, есть препядствия")
			}
				tempPole[i][j] = 1
				array[inbexCoordinateArray][0] = i //Заполнение значений координат Х
				array[inbexCoordinateArray][1] = j //Заполнение значений координат Y
				inbexCoordinateArray++

		}
	}

	r.Name=name
	r.Height=args[0]
	r.Width=args[1]
	*r.Pole=tempPole
	r.Coordinates=array
	mapPainter[name]=r
	r.MapPainter=mapPainter

	return nil
}

//Метод отображения фигуры прямоугольник
func (r *Rectangle) Paint(coordinateX, coordinateY int) error {


	if coordinateX < 0 || coordinateY < 0 || coordinateX > RazmernostPole || coordinateY > RazmernostPole { //Проверка на принадлежность координат первой ячейки нашему Полю
		return errors.New("Выход за пределы Поля") //Проверка на нахождение координат в пределах Поля
	}

	tempPole := *r.Pole                       // Получение слепка Поля
	array := r.Coordinates    //Подсчет количества ячеек в массиве координатов для прямоугольника
	topLine := r.Height + coordinateY         //Получаем длину верхней грани
	lateralLine := r.Width + coordinateX      //Получаем длину боковой грани
	if topLine < 0 || lateralLine < 0 || topLine > RazmernostPole || lateralLine > RazmernostPole { //Проверка на принадлежность координат первой ячейки нашему Полю
		logs.Info(topLine)
		logs.Warning(lateralLine)
		return errors.New("Выход за пределы Поля") //Проверка на нахождение координат в пределах Поля
	}
	// Проверка на наличие препядствий, мешающих построить прямоугольник , рисование квадрата с дыркой на Поле и запись координат в контейнер координат
	inbexCoordinateArray := 0 //Идекс в массиве координат
	for i := coordinateX; i < lateralLine; i++ {
		for j := coordinateY; j < topLine; j++ {
			if tempPole[i][j] != 0 {
				return errors.New("По данным координатам фигуру построить невозможно, есть препядствия")
			}

				tempPole[i][j] = 1
				array[inbexCoordinateArray][0] = i //Заполнение значений координат Х
				array[inbexCoordinateArray][1] = j //Заполнение значений координат Y
				inbexCoordinateArray++


		}
	}
	err := r.Delete(false)
	if err != nil {
		return err
	}
	r.Coordinates = array
	*r.Pole = tempPole
	return nil
}


// Удаление имеющегося Прямоугольника
func (r *Rectangle) Delete(isFatalDelete bool) error {

	tempPole := *r.Pole               // Получение слепка Поля
	for _, v := range r.Coordinates { //Производим замену по координатам.
		tempPole[v[0]][v[1]] = 0
	}
	*r.Pole = tempPole
	if isFatalDelete {
		delete(r.MapPainter,r.Name)
	}
	return nil
}

//Движение Прямоугольника
// Направления движения 1: Вверх
//						2: Вниз
//						3: Вправо
//						4: Влево
func (r *Rectangle) Move(goTo, howFarToGo int) error {

	switch goTo {
	case 1: //Вверх

		err := r.Paint(r.Coordinates[0][0]-howFarToGo, r.Coordinates[0][1])
		if err != nil {
			return err
		}
		return nil
	case 2: //Вниз

		err := r.Paint(r.Coordinates[0][0]+howFarToGo, r.Coordinates[0][1])
		if err != nil {
			return err
		}
		return nil
	case 3: //Вправо
		err := r.Paint(r.Coordinates[0][0], r.Coordinates[0][1]+howFarToGo)
		if err != nil {
			return err
		}
		return nil
	case 4: //Влево
		err := r.Paint(r.Coordinates[0][0], r.Coordinates[0][1]-howFarToGo)
		if err != nil {
			return err
		}
		return nil
	default:
		return errors.New("Указанное направление для движения отсутствует")
	}
}
//Методы вывода координа прямоугольника
func (r *Rectangle) TakeKoordinate () error {
	for k,v := range r.Coordinates{
		fmt.Print(" ",k,"=",v)
	}
	fmt.Println("")
	return nil
}

// Описание фигуры Змейка.
type Snake struct {
	Name        string
	Pole        *[RazmernostPole][RazmernostPole]int
	Coordinates *list.List //Координаты Змейки
	SnakeLength int
	MapPainter map[string]Painter
}

//Метод создания фигуры Змейка

//args[0]: Длина змейки
//args[1]: Координата Х
//args[2]: Координата Y
// pole - указатель на используемое поле
func (s *Snake) Create(name string, mapPainter map[string]Painter, pole *[RazmernostPole][RazmernostPole]int, args ...int) error{
	if _,ok :=mapPainter[name]; ok {
		return errors.New("Фигура с таким именем уже существует")
	}
	snakeList:=list.List{}
	lenihtSnakeLine := args[0]+args[2]    //Подсчет количества ячеек в массиве координатов для квадрата с дыркой
	coordinateX:=args[1]
	coordinateY:=args[2]
	if coordinateX < 0 || coordinateY < 0 || coordinateX > RazmernostPole || coordinateY > RazmernostPole { //Проверка на принадлежность координат первой ячейки нашему Полю
		return errors.New("Выход начальных координат за пределы Поля") //Проверка на нахождение координат в пределах Поля
	}
	if lenihtSnakeLine < 0 || lenihtSnakeLine > RazmernostPole { //Проверка на принадлежность змейки нашему Полю
		return errors.New("Выход Змейки за пределы Поля")
	}
	s.Pole=pole
	tempPole := *s.Pole
	for i := coordinateY; i < lenihtSnakeLine; i++ {
		if tempPole[coordinateX][i] != 0 {
			return errors.New("По данным координатам фигуру построить невозможно, есть препядствия")
		}
		tempPole[coordinateX][i] = 1
		snakeList.PushBack([2]int{coordinateX,i}) //Добавляем в конец следующий элемент списка
	}



	s.Name=name
	s.SnakeLength=args[0]
	*s.Pole=tempPole
	s.Coordinates=&snakeList
	mapPainter[name]=s
	s.MapPainter=mapPainter

	return nil
}

func (s *Snake) Paint(coordinateX, coordinateY int) error {
	if coordinateX < 0 || coordinateY < 0 || coordinateX > RazmernostPole || coordinateY > RazmernostPole { //Проверка на принадлежность координат первой ячейки нашему Полю
		return errors.New("Выход за пределы Поля") //Проверка на нахождение координат в пределах Поля
	}
	//snakeList:=list.List{}
	snakeList:=*s.Coordinates
	tempPole := *s.Pole                    // Получение слепка Поля

	lenihtSnakeLine := s.SnakeLength + coordinateY //Получаем длину Змейки с учетом расположения на Поле
	// Проверка на наличие препядствий, мешающих построить Змейку , рисование змейки и запись координат в контейнер координа
	if lenihtSnakeLine < 0 || lenihtSnakeLine > RazmernostPole { //Проверка на принадлежность змейки нашему Полю
		return errors.New("Выход Змейки за пределы Поля")
	}
	for i := coordinateY; i < lenihtSnakeLine; i++ {
		if tempPole[coordinateX][i] != 0 {
			return errors.New("По данным координатам фигуру построить невозможно, есть препядствия")
		}
		tempPole[coordinateX][i] = 1
		snakeList.PushBack([2]int{coordinateX,i}) //Добавляем в конец следующий элемент списка
		}

	*s.Pole = tempPole
	s.Coordinates=&snakeList

	return nil
}

// Удаление имеющейся Змейки
func (s *Snake) Delete(isFatalDelete bool) error {
	tempPole := *s.Pole               // Получение слепка Поля
	 elemetOfListCoordinate:=[2]int{}

	for e := s.Coordinates.Front(); e != nil; e = e.Next() { //Проходим по Списку
		elemetOfListCoordinate=e.Value.([2]int)			//Приводим элементы типа interface к типу [2]int
		tempPole[elemetOfListCoordinate[0]][elemetOfListCoordinate[1]]=0 //Обнуляем значение на слепке Поля

	}
	*s.Pole = tempPole //Слепок становится полем
	if isFatalDelete {
		delete(s.MapPainter,s.Name)
	}
	return nil
}

//Движение Змейки
// Направления движения 1: Вверх
//						2: Вниз
//						3: Вправо
//						4: Влево
func (s *Snake) Move(goTo, howFarToGo int) error {
	tempPole:=*s.Pole //Слепок Поля
	snakeList:=s.Coordinates //Слепок списка координат
	firstElementInSnakeList:=snakeList.Front().Value.([2]int) //получаем первый элемент списка и приводим к нужному типу
	coordinateX:=firstElementInSnakeList[0]
	coordinateY:=firstElementInSnakeList[1]

	tempArrayFirst:=[2]int{}
	tempArrayLast:=[2]int{}
	switch goTo {
	case 1: //Вверх
		if coordinateX-howFarToGo<=0 {
			return errors.New("Движение невозможно, достигнута граница поля")
		}
		for i:=coordinateX-1;i>coordinateX-howFarToGo-1 ;i--  {// Проверка на препятствие
			if tempPole[i][coordinateY] !=0{
				return errors.New("На пути следования возникало препятствие!")
			}
			tempPole[i][coordinateY] =1
			tempArrayFirst[0]=i			//Получаем координаты x
			tempArrayFirst[1]=coordinateY //Получаем координаты y
			tempArrayLast=snakeList.Remove(snakeList.Back()).([2]int) 	//Берем последний элемент списк
			tempPole[tempArrayLast[0]][tempArrayLast[1]]=0 //Затираем слепок поля по координатам последнего элемента

			snakeList.PushFront(tempArrayFirst)

		}
		*s.Pole=tempPole
		s.Coordinates=snakeList
		return nil
	case 2: //Вниз
		if coordinateX+howFarToGo>=RazmernostPole {
			return errors.New("Движение невозможно, достигнута граница поля")
		}
		for i:=coordinateX+1;i<coordinateX+howFarToGo+1 ;i++  {// Проверка на препятствие
			if tempPole[i][coordinateY] !=0{
				return errors.New("На пути следования возникало препятствие.")
			}
			tempPole[i][coordinateY] =1
			tempArrayFirst[0]=i			//Получаем координаты x
			tempArrayFirst[1]=coordinateY //Получаем координаты y
			tempArrayLast=snakeList.Remove(snakeList.Back()).([2]int) 	//Берем последний элемент списк
			tempPole[tempArrayLast[0]][tempArrayLast[1]]=0 //Затираем слепок поля по координатам последнего элемента
			snakeList.PushFront(tempArrayFirst)

		}
		*s.Pole=tempPole
		s.Coordinates=snakeList
		return nil

	case 3: //Вправо
		logs.Info(coordinateY+howFarToGo)
		if coordinateY+howFarToGo>=RazmernostPole {
			logs.Info(coordinateY+howFarToGo)
			return errors.New("Движение невозможно, достигнута граница поля")
		}
		for i:=coordinateY+1;i<coordinateY+howFarToGo+1 ;i++  {// Проверка на препятствие
			if tempPole[coordinateX][i] !=0{
				return errors.New("На пути следования возникало препятствие.")
			}
			tempPole[coordinateX][i] =1
			tempArrayFirst[0]=coordinateX			//Получаем координаты x
			tempArrayFirst[1]=i //Получаем координаты y
			tempArrayLast=snakeList.Remove(snakeList.Back()).([2]int) 	//Берем последний элемент списк
			tempPole[tempArrayLast[0]][tempArrayLast[1]]=0 //Затираем слепок поля по координатам последнего элемента
			snakeList.PushFront(tempArrayFirst)

		}
		*s.Pole=tempPole
		s.Coordinates=snakeList
		return nil

	case 4: //Влево

		if coordinateY-howFarToGo<=0 {
			return errors.New("Движение невозможно, достигнута граница поля")
		}

		for i:=coordinateY-1; i > coordinateY-howFarToGo-1 ;i--  {// Проверка на препятствие


		if tempPole[coordinateX][i]!=0 {
				return errors.New("На пути следования возникло препятствие.")
			}

			tempPole[coordinateX][i]=1

			tempArrayFirst[0]=coordinateX //Получаем координаты x
			tempArrayFirst[1]=i //Получаем координаты y

			tempArrayLast=snakeList.Remove(snakeList.Back()).([2]int) 	//Берем последний элемент списк
			tempPole[tempArrayLast[0]][tempArrayLast[1]]=0 //Затираем слепок поля по координатам последнего элемента
			snakeList.PushFront(tempArrayFirst)




		}

		*s.Pole=tempPole

		s.Coordinates=snakeList
		return nil
	default:
		return errors.New("Указанное направление для движения отсутствует")
	}
}

//Методы вывода координа Змейки
func (s *Snake) TakeKoordinate () error {
	for e := s.Coordinates.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	return nil
}
