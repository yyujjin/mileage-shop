package main

import (
	"fmt"
	"mileage-shop/src/menu"
)

func newUser(name string) menu.User {
	d := menu.User{
		name, 1000000, "",
	}
	return d
}

func main() {

	var user1 = newUser("박유진")
	var user2 = newUser("박수현")
	var user3 = newUser("박정숙")

	var users = []menu.User{}
	users = append(
		users,
		user1,
		user2,
		user3,
	)

	for i := 0; i < len(users); i++ {
		fmt.Printf("%v : %v\n", users[i].Name, users[i].Point)
	}

	var inputMenu int

	for {
		fmt.Println(
			"메뉴의 번호를 입력하세요.\n" +
				"1.구매\n" +
				"2.잔여 수량 확인\n" +
				"3.잔여 마일리지 확인\n" +
				"4.배송 상태 확인\n" +
				"5.장바구니 확인\n" +
				"6.프로그램 종료")

		fmt.Scan(&inputMenu)

		switch inputMenu {
		case 1:
			menu.Purchase(users...)
		case 2:
			menu.RemainingAmount()
		case 3:
			menu.RemainingMileage(users[0].Name, users[0].Point)
		case 4:
			menu.DeliveryStatus()
		case 5:
			menu.CheckCart()
		case 6:
			menu.ExitProgram()
			return
		default:
			fmt.Println("잘못된 입력입니다. 1~6까지의 숫자를 입력해 주세요.")
		}
	}
}
