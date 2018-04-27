package main

import (
	"errors"
	"fmt"
)
type Pointer interface {
	 Create (n,x,y,typeFiguri int, name string) error
	 //Delete () error
	 //Move () error
	// TakeKoordinate ()error
	//testToCreate () error // Возможно стоит добавить в Crete
	//testToMove () error // Возможно стоит добавить в Move
}
type kvadrat struct {
	Name string
	storona int		//Значение стороны квадрата. Получаем при создании
	firstCell [2]int
	Pole *[20][20]int
}



func (k *kvadrat) Create (n,x,y, typeFiguri int, name string) error {  //метод создания квалрат с дыркой
	err:=k.testToCreate(n,x,y,2)
	if err!=nil {
			return err
		} else {
		vertical:=x+n		//положение квадрата по вертикали
		gorizontal:=y+n		//положение квадрата по горизонтали
	for i:=x;i<vertical ;i++  {
		k.Pole[i][y]=1
		k.Pole[i][gorizontal-1]=1
	}
	for j:=y+1;j<gorizontal-1 ;j++  {
		k.Pole[x][j]=1
		k.Pole[vertical-1][j]=1
	}
	k.storona=n
	k.firstCell[0]=x
	k.firstCell[1]=y
	k.Name=name			//Временная заглушка, данная информация
	return nil
	}
}

func (k *kvadrat) testToCreate(n,x,y, typeFiguri int) error { //Метод проверки возможности создания фигуры
	vertical:=x+n		//положение квадрата по вертикали
	gorizontal:=y+n		//положение квадрата по горизонтали
	for i:=x;i<vertical ;i++  {
		for j:=y;j<gorizontal ;j++  {
			if k.Pole[i][j]!=0 {
				return errors.New("По данным координатам фигуру построить невозможно, занято")
			}
		}
	}
return nil
}
/*func (k *kvadrat) testToMove (moveTo, moveToNCell int) error{
	if k.firstCell[0]-moveToNCell<=0 {
		return errors.New("Перемещение вверх невозможно, достигается верхняя граница")
	} else {
		for i:=k.firstCell[0]-1;i>k.firstCell[0]-moveToNCell ;i--  {

		}
	}
}*/

/*func (k *kvadrat) Move (name string, moveTo, moveToNCell int) error {
	switch moveTo {
	case 1:	//move to Up
	case 2:
	case 3:
	case 4:
	default:

	}
}*/
func main()  {
	/*pole:=[20][20]int{} //Наше поле
	pole[1][4]=1
	newKvadrat:=new(kvadrat)	//Инициализация указателя на Структуру квадрата с дыркой
	newKvadrat.Pole=&pole  //передача алреса Нашего поля в объект структуры квадрат с дыркой

	err:=newKvadrat.Create(5,1,1, 2,"name")
	if err!=nil {
		fmt.Println(err)
	}

	for _,v :=range pole{	//ВЫвод НАшего поля на экран
		fmt.Println(v)
	}*/

	n:=10
	array:=make([][]int,n)
	arrayOfKoordinat:=[2]int{}
	array[0]=append(array[0],arrayOfKoordinat[0],arrayOfKoordinat[1])

	fmt.Println(array[0][1])
	/*for _,v:=range array {
		v=append(v,arrayOfKoordinat)
		fmt.Println(array)
	}*/

}

