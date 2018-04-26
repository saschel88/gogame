package main

import (
	"fmt"
	"errors"
)

type kvadrat struct {
	storona int		//Значение стороны квадрата. Получаем при создании
	firstCell [2]int
	Pole *[20][20]int
}

func (k *kvadrat) Create (n,x,y int) error {  //метод создания квалрат с дыркой
	err:=k.testToCreate(n,x,y)
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
	return nil
	}
}

func (k *kvadrat) testToCreate(n,x,y int) error { //Метод проверки возможности создания фигуры
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

func main()  {
	pole:=[20][20]int{}
	pole[1][4]=1
	newKvadrat:=new(kvadrat)
	newKvadrat.Pole=&pole

	err:=newKvadrat.Create(5,1,1)
	if err!=nil {
		fmt.Println(err)
	}
	//fmt.Println(newKvadrat)
	for _,v :=range pole{
		fmt.Println(v)
	}

}

