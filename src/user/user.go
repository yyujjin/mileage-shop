package user

import (
	"fmt"
)

type User struct {
	Name  string
	Point uint
	Cart  string
}

func RemainingMileage(name string, point uint) {
	fmt.Println("잔여 마일리지 확인")
	fmt.Printf("%v님의 잔여 마일리지는 %v점 입니다.", name, point)
}
