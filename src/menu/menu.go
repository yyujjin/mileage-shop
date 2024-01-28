package menu

import (
	"fmt"
	"mileage-shop/src/product"
)

func RemainingAmount() {
	fmt.Println("잔여 수량 확인")
	for _, value := range product.Items {
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
	fmt.Println(product.Cart)
}

func ExitProgram() {
	fmt.Println("프로그램 종료")
}
