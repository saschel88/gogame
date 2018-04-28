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
	Delete () error
	//Move () error
	// TakeKoordinate ()error

}

// Описание фигуры Квадрат с дырой
type Square struct {
	Name        string
	Pole        *[RazmernostPole][RazmernostPole]int
	Coordinates [][2]int
	SideSquare  int //Значение стороны квадрата. Получаем при создании

}

// Описание фигуры Прямоугольник.
type Rectangle struct {
	Name        string
	Pole        *[RazmernostPole][RazmernostPole]int
	Coordinates [][2]int
	Height      int //Значение ширины прямоугольника. Получаем при создании
	Width       int //Значение высоты прямоугольника. Получаем при создании
}

// Описание фигуры Змейка.
type Snake struct {
	Name        string
	Pole        *[RazmernostPole][RazmernostPole]int
	Coordinates [][2]int //Значение стороны квадрата. Получаем при создании
	SnakeLength int
}

//Метод создания фигуры квадрат с дыркой
func (s *Square) Create(coordinateX, coordinateY int) error {
	tempPole := *s.Pole // Получение слепка Поля
			cellCountInArray := s.SideSquare*4 - 4    //Подсчет количества ячеек в массиве координатов для квадрата с дыркой
		array := make([][2]int, cellCountInArray) // Инициализация слайса  массивов координат
		s.Coordinates = array //Заполнение поля структуры инициализированным массивом координат

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
	tempPole := *s.Pole // Получение слепка Поля
	for _,v :=range s.Coordinates{ //Производим замену по координатам.
		tempPole[v[0]][v[1]]=0
	}
	*s.Pole = tempPole
	return nil
}

//Метод создания фигуры прямоугольник
func (r *Rectangle) Create(coordinateX, coordinateY int) error {
	tempPole := *r.Pole // Получение слепка Поля
	cellCountInArray := r.Width*r.Height    //Подсчет количества ячеек в массиве координатов для квадрата с дыркой
	array := make([][2]int, cellCountInArray) // Инициализация слайса  массивов координат
	for _, v := range array {                 //Заполнение дефолтными значениями двумерного массива
		v[0] = 0
		v[1] = 0

	}
	r.Coordinates = array //Заполнение поля структуры инициализированным массивом координат
	topLine := r.Height + coordinateX     //Получаем длину верхней грани
	lateralLine := r.Width + coordinateY //Получаем длину боковой грани
	// Проверка на наличие препядствий, мешающих построить квадрат , рисование квадрата с дыркой на Поле и запись координат в контейнер координат
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
