package main

import (
	"fmt"
)

type user struct {
	name  string
	point uint
	cart  string
}

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
		id     int
		name   string
		point  uint
		amount int
	}
	//Go 언어에서 슬라이스는 참조 타입입니다. 여기에만 * 붙여주면 된대. 쓰려는 곳에 & 이거 안붙여줘도 된다는데 뮤슨말?
	//그럼 밑에 value에서는 새 값이 나타난거잖아 for range 안에서?
	//*이걸 적어줌으로써 value도 items의 값을 참조하게 되고 그 참조한 값으로 값을 변경해줘서 바뀌는건가?
	items := []*item{
		{1, "텀블러", 10000, 30},
		{2, "롱패딩", 500000, 20},
		{3, "투미 백팩", 400000, 20},
		{4, "나이키 운동화", 150000, 50},
		{5, "빼빼로", 1200, 500},
	}

	var menu int

	for {
		fmt.Println(
			"메뉴의 번호를 입력하세요.\n" +
				"1.구매\n" +
				"2.잔여 수량 확인\n" +
				"3.잔여 마일리지 확인\n" +
				"4.배송 상태 확인\n" +
				"5.장바구니 확인\n" +
				"6.프로그램 종료")
		fmt.Scan(&menu)
		switch menu {
		case 1:
			fmt.Println("구매")
			for _, value := range items {
				fmt.Printf("물품%v : %v, 가격: %v원, 잔여 수량: %v\n", value.id, value.name, value.point, value.amount)
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
			for _, value := range items {
				if num == value.id {
					if value.amount > 0 && users[0].point >= value.point {
						users[0].point -= value.point
						value.amount--
						//items[index].amount-- 포인터 안만들었을 때
						fmt.Println("구매가 완료되었습니다.")
					} else if value.amount < 0 {
						fmt.Println("잔여 수량이 부족하여 구매가 불가능합니다.")
					} else {
						fmt.Println("보유 포인트가 부족하여 구매가 불가능합니다.")
					}
				}
			}
		case 2:
			fmt.Println("잔여 수량 확인")
			for _, value := range items {
				fmt.Printf("%v의 잔여 수량은 %v입니다.", value.name, value.amount)
			}
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
		default:
			fmt.Println("잘못된 입력입니다. 1~6까지의 숫자를 입력해 주세요.")
		}
	}
}
