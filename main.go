package main

import (
	"fmt"
	"mileage-shop/src/product"
	"mileage-shop/src/user"
)

func newUser(name string) user.User {
	d := user.User{
		name, 1000000, "",
	}
	return d
}

func ExitProgram() {
	fmt.Println("프로그램 종료")
}

func main() {

	var user1 = newUser("박유진")
	var user2 = newUser("박수현")
	var user3 = newUser("박정숙")

	var users = []user.User{}
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

	// 패키지를 분리하는 이유? 유지보수가 편함!
	// 몇개의 패키지로 분리해야 하지?
	// 구매 (+ 배송상태 확인), 장바구니, 사용자(마일리지) // 만들고 없애고 반복할수 있음 -> 리팩토링한다.
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
			product.Purchase(users...)
		case 2:
			product.RemainingAmount()
		case 3:
			user.RemainingMileage(users[0].Name, users[0].Point)
		case 4:
			product.DeliveryStatus()
		case 5:
			product.CheckCart()
		case 6:
			ExitProgram()
			return
		default:
			fmt.Println("잘못된 입력입니다. 1~6까지의 숫자를 입력해 주세요.")
		}
	}
}
