package product

import (
	"fmt"
	"mileage-shop/src/user"
)

type Item struct {
	Id     int
	Name   string
	Point  uint
	Amount int
}

var Items = []*Item{
	{1, "텀블러", 10000, 30},
	{2, "롱패딩", 500000, 20},
	{3, "투미 백팩", 400000, 20},
	{4, "나이키 운동화", 150000, 50},
	{5, "빼빼로", 1200, 500},
}

var Cart []Item

func Purchase(users ...user.User) {
	fmt.Println("구매")
	for _, value := range Items {
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
				return
			}
		}
	}
	fmt.Printf("구매할 상품 : %v\n", num)
	// 바로 구매? 지금 그대로. 장바구니에 담기? 장바구니에 추가하고 추가되었습니다. 무조건 하나씩만 구매 가능
	// 장바구니에 있는 물품 목록을 출력.
	for _, value := range Items {
		if num == value.Id {
			if value.Amount > 0 && users[0].Point >= value.Point {
				for {
					var isPutCart int
					fmt.Println("원하시는 번호를 입력하세요. 1.바로 구매 2.장바구니에 담기")
					fmt.Scan(&isPutCart)
					if isPutCart == 1 {
						break
					} else {
						fmt.Println("장바구니에 추가되었습니다.")
						Cart = append(Cart, Item{value.Id, value.Name, value.Point, 1})
						return
					}
				}
				users[0].Point -= value.Point
				value.Amount--
				fmt.Println("구매가 완료되었습니다.")
			} else if value.Amount < 0 {
				fmt.Println("잔여 수량이 부족하여 구매가 불가능합니다.")
			} else {
				fmt.Println("보유 포인트가 부족하여 구매가 불가능합니다.")
			}
		}
	}
}
