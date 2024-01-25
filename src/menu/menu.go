package menu

import "fmt"

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

type User struct {
	Name  string
	Point uint
	Cart  string
}

func Purchase(users ...User) {
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
	for _, value := range Items {
		if num == value.Id {
			if value.Amount > 0 && users[0].Point >= value.Point {
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

func RemainingAmount() {
	fmt.Println("잔여 수량 확인")
	for _, value := range Items {
		fmt.Printf("%v의 잔여 수량은 %v입니다.\n", value.Name, value.Amount)
	}
}

func RemainingMileage(name string, point uint) {
	fmt.Println("잔여 마일리지 확인")
	fmt.Printf("%v님의 잔여 마일리지는 %v점 입니다.", name, point)
}

func DeliveryStatus() {
	fmt.Println("배송 상태 확인")
}

func CheckCart() {
	fmt.Println("장바구니 확인")
}

func ExitProgram() {
	fmt.Println("프로그램 종료")
}
