package menu

import "fmt"

// 구조체 자체도 main.go에서 쓰려면 대문자로 바꿔줘야 함
// 하지만 main.go에서 필드를 사용하려면 필드도 대문자로 고쳐줘야함
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
				//items[index].amount-- 포인터 안만들었을 때
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

//패키지를 사용해서 따로 빼두는 이유는 관련있는 함수나 변수끼리 모아두기 위해서
//함수가 너무 길어지면 나중에 유지보수가 힘들어서 main함수밖에 있는걸 따로 정리해두는 것
//새로 만든 패키지에서 전역변수로 대문자로 선언해야 main함수에서 인식하고 쓸수있음
//main.go에 있는 정보를 새로만든 패키지로 넘기려면 매개변수로 넘길수있고
//새로만든 패키지에서 대문자로 쓰면 main.go에서 사요되는거임
