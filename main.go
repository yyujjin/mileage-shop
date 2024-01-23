package main

import "fmt"

type user struct {
	name  string
	point int
	cart  string
}

// 1. 포인터 함수가 되려면 내보내는 데이터타입 앞에 * 붙이고
func newUser(name string) user {
	d := user{
		name, 1000000, "",
	}
	return d
}

func main() {
	user1 := newUser("박유진")
	user2 := newUser("박수현")
	user3 := newUser("박정숙")
	//2. 여기에도 구조체 이름 앞에 * 붙이기 => why? 포인터 구조체들이 모여있는 배열이라서 (데이터 타입이 달라서 일반 구조체와 포이터 구조체를 같은 배열안에 담을 수 없음 )
	users := []user{}
	users = append(
		users,
		user1,
		user2,
		user3,
	)
	for i := 0; i < len(users); i++ {
		fmt.Printf("%v : %v\n", users[i].name, users[i].point)
	}

	type item struct {
		name   string
		point  int
		amount int
	}

	items := []item{
		{"텀블러", 10000, 30},
		{"롱패딩", 500000, 20},
		{"투미 백팩", 400000, 20},
		{"나이키 운동화", 150000, 50},
		{"빼빼로", 1200, 500},
	}

	fmt.Println(items)

	var menu int

	for {
		fmt.Scan(&menu)
		switch menu {
		case 1:
			fmt.Println("구매")
		case 2:
			fmt.Println("잔여 수량 확인")
		case 3:
			fmt.Println("잔여 마일리지 확인")
			fmt.Printf("%v님의 잔여 마일리지는 %v점 입니다.", users[0].name, users[0].point)
		case 4:
			fmt.Println("배송 상태 확인")
		case 5:
			fmt.Println("장바구니 확인")
		case 6:
			fmt.Println("프로그램 종료")
			return
		}
	}

	// 사용자가 3을 입력하기 전까진 계속 메뉴를 입력받을 수 있음
	// 실행 -> 1 -> 구매 화면~ -> 엔터 -> 메뉴를 다시 입력 가능 -> 2 -> 장바구니~~ -> 엔터 -> 메뉴 입력 가능 -> 3 -> 프로그램이 종료~
}
