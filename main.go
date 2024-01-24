package main

import (
	"fmt"
	"mileage-shop/src/menu"
)

func newUser(name string) user {
	d := user{
		name, 1000000, "",
	}
	return d
}

type user struct {
	name  string
	point uint
	cart  string
}

func main() {

	var user1 = newUser("박유진")
	var user2 = newUser("박수현")
	var user3 = newUser("박정숙")

	var users = []user{}
	users = append(
		users,
		user1,
		user2,
		user3,
	)

	for i := 0; i < len(users); i++ {
		fmt.Printf("%v : %v\n", users[i].name, users[i].point)
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
			fmt.Println("구매")
			for _, value := range menu.Items {
				fmt.Printf("물품%v : %v, 가격: %v원, 잔여 수량: %v\n", value.Id, value.Name, value.Point, value.Amount)
			}
			var num int

			fmt.Print("구매 할 물품 번호를 입력하세요: ")
			fmt.Scan(&num)
			if num < 1 || num > 5 {
				for {
					fmt.Println("잘못된 입력입니다. 1~5까지의 숫자를 입력해주세요.(메뉴로 돌아가려면 0번을 누르세요)")
					fmt.Scan(&num)
					if num >= 1 && num <= 5 {
						break
					}
					if num == 0 {
						break
					}
				}
				if num == 0 {
					continue
				}
			}
			fmt.Printf("구매할 상품 : %v\n", num)
			for _, value := range menu.Items {
				if num == value.Id {
					if value.Amount > 0 && users[0].point >= value.Point {
						users[0].point -= value.Point
						value.Amount--
						//items[index].amount-- 포인터 안만들었을 때
						fmt.Println("구매가 완료되었습니다.")
					} else if value.Amount < 0 {
						fmt.Println("잔여 수량이 부족하여 구매가 불가능합니다.")
					} else {
						fmt.Println("보유 포인트가 부족하여 구매가 불가능합니다.")
					}
				}
			}
		case 2:
			menu.RemainingAmount()
		case 3:
			menu.RemainingMileage(users[0].name, users[0].point)
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
