// Опиание структур используемых фигур и интерфейса, реализуемого данными структурами
package figures

import (
	"errors"
	"fmt"
)

const RazmernostPole = 20 //размерность отображаемого поля
// Интерфейс рописывающий дефтельность фигур на поле
type Painter interface {
	// Координаты имеют вид array[X][Y]
	//coordinateX - координата по X
	//coordinateY - координата по Y
	Create(coordinateX, coordinateY int) error
	// Метод удаления. Возвращаему ошибку взять на контроль (нужна ли она)
	Delete() error
	Move(goTo, howFarToGo int) error
	// TakeKoordinate ()error

}

// Описание фигуры Квадрат с дырой
type Square struct {
	Name        string
	Pole        *[RazmernostPole][RazmernostPole]int
	Coordinates [][2]int
	SideSquare  int //Значение стороны квадрата. Получаем при создании

}

//Метод создания фигуры квадрат с дыркой
func (s *Square) Create(coordinateX, coordinateY int) error {
	if coordinateX < 0 || coordinateY < 0 || coordinateX > RazmernostPole || coordinateY > RazmernostPole { //Проверка на принадлежность координат первой ячейки нашему Полю
		return errors.New("Выход за пределы Поля") //Проверка на нахождение координат в пределах Поля
	}
	tempPole := *s.Pole                       // Получение слепка Поля
	cellCountInArray := s.SideSquare*4 - 4    //Подсчет количества ячеек в массиве координатов для квадрата с дыркой
	array := make([][2]int, cellCountInArray) // Инициализация слайса  массивов координат
	s.Coordinates = array                     //Заполнение поля структуры инициализированным массивом координат

	topLine := s.SideSquare + coordinateY     //Получаем длину верхней грани
	lateralLine := s.SideSquare + coordinateX //Получаем длину боковой грани
	// Проверка на наличие препядствий, мешающих построить квадрат , рисование квадрата с дыркой на Поле и запись координат в контейнер координат
	inbexCoordinateArray := 0 //Идекс в массиве координат
	for i := coordinateX; i < lateralLine; i++ {
		for j := coordinateY; j < topLine; j++ {
			if tempPole[i][j] != 0 {
				return errors.New("По данным координатам фигуру построить невозможно, есть препядствия")
			}
			if i == coordinateX || i == lateralLine-1 || j == coordinateY || j == topLine-1 {
				tempPole[i][j] = 1
				s.Coordinates[inbexCoordinateArray][0] = i //Заполнение значений координат Х
				s.Coordinates[inbexCoordinateArray][1] = j //Заполнение значений координат Y
				inbexCoordinateArray++

			}
		}
	}
	fmt.Println(s.Coordinates)
	*s.Pole = tempPole
	return nil
}

// Удаление имеющегося квадрата с дыркой
func (s *Square) Delete() error {
	tempPole := *s.Pole               // Получение слепка Поля
	for _, v := range s.Coordinates { //Производим замену по координатам.
		tempPole[v[0]][v[1]] = 0
	}
	*s.Pole = tempPole
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
		err := s.Delete()
		if err != nil {
			return err
		}
		err = s.Create(s.Coordinates[0][0]-howFarToGo, s.Coordinates[0][1])
		if err != nil {
			return err
		}
		return nil
	case 2: //Вниз
		err := s.Delete()
		if err != nil {
			return err
		}
		err = s.Create(s.Coordinates[0][0]+howFarToGo, s.Coordinates[0][1])
		if err != nil {
			return err
		}
		return nil
	case 3: //Вправо
		err := s.Delete()
		if err != nil {
			return err
		}
		err = s.Create(s.Coordinates[0][0], s.Coordinates[0][1]+howFarToGo)
		if err != nil {
			return err
		}
		return nil
	case 4: //Влево
		err := s.Delete()
		if err != nil {
			return err
		}
		err = s.Create(s.Coordinates[0][0], s.Coordinates[0][1]-howFarToGo)
		if err != nil {
			return err
		}
		return nil
	default:
		return errors.New("Указанное направление для движения отсутствует")
	}
}

// Описание фигуры Прямоугольник.
type Rectangle struct {
	Name        string
	Pole        *[RazmernostPole][RazmernostPole]int
	Coordinates [][2]int
	Height      int //Значение ширины прямоугольника. Получаем при создании
	Width       int //Значение высоты прямоугольника. Получаем при создании
}

//Метод создания фигуры прямоугольник
func (r *Rectangle) Create(coordinateX, coordinateY int) error {
	if coordinateX < 0 || coordinateY < 0 || coordinateX > RazmernostPole || coordinateY > RazmernostPole { //Проверка на принадлежность координат первой ячейки нашему Полю
		return errors.New("Выход за пределы Поля") //Проверка на нахождение координат в пределах Поля
	}
	tempPole := *r.Pole                       // Получение слепка Поля
	cellCountInArray := r.Width * r.Height    //Подсчет количества ячеек в массиве координатов для Прямоугольника
	array := make([][2]int, cellCountInArray) // Инициализация слайса  массивов координат
	r.Coordinates = array                     //Заполнение поля структуры инициализированным массивом координат
	topLine := r.Height + coordinateX         //Получаем длину верхней грани
	lateralLine := r.Width + coordinateY      //Получаем длину боковой грани
	// Проверка на наличие препядствий, мешающих построить Прямоугольник , рисование Прямоугольника на Поле и запись координат в контейнер координат
	inbexCoordinateArray := 0 //Идекс в массиве координат
	for i := coordinateX; i < lateralLine; i++ {
		for j := coordinateY; j < topLine; j++ {
			if tempPole[i][j] != 0 {
				return errors.New("По данным координатам фигуру построить невозможно, есть препядствия")
			}
			tempPole[i][j] = 1
			r.Coordinates[inbexCoordinateArray][0] = i //Заполнение значений координат Х
			r.Coordinates[inbexCoordinateArray][1] = j //Заполнение значений координат Y
			inbexCoordinateArray++

		}
	}
	fmt.Println(r.Coordinates)
	*r.Pole = tempPole
	return nil
}

// Удаление имеющегося Прямоугольника
func (r *Rectangle) Delete() error {
	tempPole := *r.Pole               // Получение слепка Поля
	for _, v := range r.Coordinates { //Производим замену по координатам.
		tempPole[v[0]][v[1]] = 0
	}
	*r.Pole = tempPole
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
		err := r.Delete()
		if err != nil {
			return err
		}
		err = r.Create(r.Coordinates[0][0]-howFarToGo, r.Coordinates[0][1])
		if err != nil {
			return err
		}
		return nil
	case 2: //Вниз
		err := r.Delete()
		if err != nil {
			return err
		}
		err = r.Create(r.Coordinates[0][0]+howFarToGo, r.Coordinates[0][1])
		if err != nil {
			return err
		}
		return nil
	case 3: //Вправо
		err := r.Delete()
		if err != nil {
			return err
		}
		err = r.Create(r.Coordinates[0][0], r.Coordinates[0][1]+howFarToGo)
		if err != nil {
			return err
		}
		return nil
	case 4: //Влево
		err := r.Delete()
		if err != nil {
			return err
		}
		err = r.Create(r.Coordinates[0][0], r.Coordinates[0][1]-howFarToGo)
		if err != nil {
			return err
		}
		return nil
	default:
		return errors.New("Указанное направление для движения отсутствует")
	}
}

// Описание фигуры Змейка.
type Snake struct {
	Name        string
	Pole        *[RazmernostPole][RazmernostPole]int
	Coordinates [][2]int //Координаты Змейки
	SnakeLength int
}

//Метод создания фигуры Змейка
func (s *Snake) Create(coordinateX, coordinateY int) error {
	if coordinateX < 0 || coordinateY < 0 || coordinateX > RazmernostPole || coordinateY > RazmernostPole { //Проверка на принадлежность координат первой ячейки нашему Полю
		return errors.New("Выход за пределы Поля") //Проверка на нахождение координат в пределах Поля
	}
	tempPole := *s.Pole                    // Получение слепка Поля
	array := make([][2]int, s.SnakeLength) // Инициализация слайса  массивов координат
	s.Coordinates = array                  //Заполнение поля структуры инициализированным массивом координат

	lenihtSnakeLine := s.SnakeLength + coordinateY //Получаем длину Змейки с учетом расположения на Поле
	// Проверка на наличие препядствий, мешающих построить Змейку , рисование змейки и запись координат в контейнер координа
	inbexCoordinateArray := 0 //Идекс в массиве координат
	for i := coordinateY; i < lenihtSnakeLine; i++ {
		if tempPole[coordinateX][i] != 0 {
			return errors.New("По данным координатам фигуру построить невозможно, есть препядствия")
		}
		tempPole[coordinateX][i] = 1
		s.Coordinates[inbexCoordinateArray][0] = coordinateX //Заполнение значений координат Х
		s.Coordinates[inbexCoordinateArray][1] = i           //Заполнение значений координат Y
		inbexCoordinateArray++
	}
	fmt.Println(s.Coordinates)
	*s.Pole = tempPole
	return nil
}

// Удаление имеющегося квадрата с дыркой
func (s *Snake) Delete() error {
	tempPole := *s.Pole               // Получение слепка Поля
	for _, v := range s.Coordinates { //Производим замену по координатам.
		tempPole[v[0]][v[1]] = 0
	}
	*s.Pole = tempPole
	return nil
}

//Движение квадрата с дыркой
// Направления движения 1: Вверх
//						2: Вниз
//						3: Вправо
//						4: Влево
func (s *Snake) Move(goTo, howFarToGo int) error {
	//tempPole:=*s.Pole //Слепок Поля
	switch goTo {
	case 1: //Вверх
		if s.Coordinates[0][0]-howFarToGo < 0 { //Проверка на нахождение координат в пределах Поля
			return errors.New("Выход за пределы Поля")
		}
	//	tempCoordinate:=s.Coordinates //Слепок массива координат
 	/*	for i:=; ;  {
			//Здесь будет цикл, в которм будет проверятся каждая ячейка на пути на наличие препятствия, в случае его отсутсвия змейка будет смещаться пузырьком.
			//Если будет препятствие, тоы выход.
			}
*/

	default:
		return errors.New("Указанное направление для движения отсутствует")
	}
}
