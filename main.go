package main

import "fmt"

func main() {

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
			fmt.Println("구매 화면입니다.")
		case 2:
			fmt.Println("장바구니 화면입니다.")
		case 3:
			fmt.Println("프로그램이 종료됩니다.")
			return

		}
	}

	// 사용자가 3을 입력하기 전까진 계속 메뉴를 입력받을 수 있음
	// 실행 -> 1 -> 구매 화면~ -> 엔터 -> 메뉴를 다시 입력 가능 -> 2 -> 장바구니~~ -> 엔터 -> 메뉴 입력 가능 -> 3 -> 프로그램이 종료~
}
