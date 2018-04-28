
// Опиание структур используемых фигур и интерфейса, реализуемого данными структурами
package figures

import (

	"errors"

)
const RazmernostPole = 50
type Painter interface {
	Create(coordinateX, coordinateY, cellCount int, isCreate bool) error
	//Delete () error
	//Move () error
	// TakeKoordinate ()error
	//testToCreate () error // Возможно стоит добавить в Crete
	//testToMove () error // Возможно стоит добавить в Move
}

// Описание фигуры Квадрат с дырой
type Square struct {
	Name        string
	Pole        *[RazmernostPole][RazmernostPole]int
	Coordinates [][]int
	SideSquare  int //Значение стороны квадрата. Получаем при создании

}

// Описание фигуры Прямоугольник.
type Rectangle struct {
	Name        string
	Pole        *[RazmernostPole][RazmernostPole]int
	Coordinates [][]int
	Height      int //Значение ширины прямоугольника. Получаем при создании
	Width       int //Значение высоты прямоугольника. Получаем при создании
}

// Описание фигуры Змейка.
type Snake struct {
	Name        string
	Pole        *[RazmernostPole][RazmernostPole]int
	Coordinates [][]int //Значение стороны квадрата. Получаем при создании
	SnakeLength int
}

//Метод создания фигуры квадрат с дыркой
func (s *Square) Create(coordinateX, coordinateY, cellCount int, isCreate bool) error {
	tempPole := *s.Pole // Получение слепка Поля
	if isCreate {
		return nil //Дествия при движении
	} else {
		cellCountInArray := s.SideSquare*4 - 4   //Подсчет количества ячеек в массиве координатов для квадрата с дыркой
		array := make([][]int, cellCountInArray) // Инициализация двумерного массива координат
		arrayOfKoordinat := make([]int,2)
		for _, v := range array { //Заполнение дефолтными значениями двумерного массива
			v = append(v, arrayOfKoordinat[0], arrayOfKoordinat[1])

		}

		s.Coordinates = array //Заполнение поля структуры инициализированным массивом координат

		topLine := s.SideSquare + coordinateY     //Получаем длину верхней грани
		lateralLine := s.SideSquare + coordinateX //Получаем длину боковой грани
		//Проверка на наличе препядствий на поле, мешающих построить квадрат
		for i := coordinateX; i < lateralLine; i++ {
			for j := coordinateY; j < topLine; j++ {
				if s.Pole[i][j] != 0 {
					return errors.New("По данным координатам фигуру построить невозможно, есть препядствия")
				}
			}
		}
		for i := coordinateX; i < lateralLine; i++ { //Рисуем 2 вертикальные линии
			s.Pole[i][coordinateY] = 1
			s.Pole[i][topLine-1] = 1

		}
		for j := coordinateY + 1; j < topLine-1; j++ { //Рисуем 2 горизонтальные линии
			s.Pole[coordinateX][j] = 1
			s.Pole[lateralLine-1][j] = 1
		}
		s.Pole = &tempPole
		return nil
	}

}

