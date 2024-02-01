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

type CartItem struct {
	id     int
	name   string
	amount int
}

var Cart []CartItem

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
			fmt.Println("잘못된 입력입니다. 1~5까지의 숫자를 입력해주세요.(메뉴로 돌아가려면 0번을 누르세요.)")
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

	// 장바구니에 있는 물품 목록을 출력.
	for _, value := range Items {
		if num == value.Id {
			if value.Amount > 0 && users[0].Point >= value.Point {
				var isPutCart int
				fmt.Println("원하시는 번호를 입력하세요. 1.바로 구매 2.장바구니에 담기 (구매를 취소하려면 0번을 누르세요.)")
				fmt.Scan(&isPutCart)
				switch isPutCart {
				case 1:
					PurchaseItem(&users[0], value)
				case 2:
					PutToCart(value)
				case 0:
					return
				}
			} else if value.Amount < 0 {
				fmt.Println("잔여 수량이 부족하여 구매가 불가능합니다.")
			} else {
				fmt.Println("보유 포인트가 부족하여 구매가 불가능합니다.")
			}
		}
	}
}
func RemainingAmount() {
	fmt.Println("잔여 수량 확인")
	for _, value := range Items {
		fmt.Printf("%v의 잔여 수량은 %v입니다.\n", value.Name, value.Amount)
	}
}

func DeliveryStatus() {
	fmt.Println("배송 상태 확인")
}

func CheckCart(users []user.User) {
	fmt.Println("장바구니 확인")
	for index := range Cart {
		fmt.Printf("제품 ID : %v, 제품 이름 : %v, 장바구니 수량 : %v \n", Cart[index].id, Cart[index].name, Cart[index].amount)
	}
	var num int
	for {
		fmt.Println("구매를 원하시면 1번 , 메뉴로 돌아가려면 0번을 누르세요.")
		fmt.Scan(&num)
		if num == 1 || num == 0 {
			break
		}
	}
	switch num {
	case 1:
		var selectedItem int
		for {
			fmt.Println("구매를 원하시는 제품의 ID를 입력하세요 (메뉴로 돌아가려면 0번을 누르세요.) ")
			fmt.Scan(&selectedItem)
			if selectedItem == 0 {
				return
			}
			foundItem := findItem(selectedItem)
			if foundItem == -1 {
				fmt.Println("입력하신 ID가 장바구니에 존재하지 않습니다.")
			} else {
				break
			}
		}
		for _, value := range Items {
			if selectedItem == value.Id {
				if value.Amount > 0 && users[0].Point >= value.Point {
					PurchaseItem(&users[0], value)
					foundIndex := findItem(selectedItem)
					Cart = append(Cart[:foundIndex], Cart[foundIndex+1:]...)
				} else if value.Amount < 0 {
					fmt.Println("잔여 수량이 부족하여 구매가 불가능합니다.")
					return
				} else {
					fmt.Println("보유 포인트가 부족하여 구매가 불가능합니다.")
					return
				}
			}
		}
	case 0:
		return
	}

}

func PurchaseItem(user *user.User, value *Item) {
	user.Point -= value.Point
	value.Amount--
	fmt.Println("구매가 완료되었습니다.")
}

func PutToCart(item *Item) {
	foundIndex := findItem(item.Id)
	if foundIndex == -1 {
		Cart = append(Cart, CartItem{item.Id, item.Name, 1})
	} else {
		Cart[foundIndex].amount += 1
	}

	fmt.Println("장바구니에 추가되었습니다.")
}

func findItem(itemId int) int {
	foundIndex := -1
	for index := range Cart {
		if Cart[index].id == itemId {
			foundIndex = index
			break
		}
	}
	return foundIndex
}
