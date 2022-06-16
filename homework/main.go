package main

import "fmt"

var (
	apple_priсe float32 = 5.99 //Цена яблока
	pear_price  uint8   = 7    //Цена груши
	cash        uint8   = 23   //Наличествующие деньги
)

func main() {
	summ := apple_priсe*9 + float32(pear_price)*8
	fmt.Println("Скільки грошей треба витратити, щоб купити 9 яблук та 8 груш?", summ)
	number_of_pears := cash / pear_price
	fmt.Println("Скільки груш ми можемо купити?", number_of_pears)
	number_of_apples := uint8(float32(cash) / apple_priсe)
	fmt.Println("Скільки яблук ми можемо купити?", number_of_apples)
	can_buy := float32(cash) >= (apple_priсe*2 + float32(pear_price)*2)
	fmt.Println("Чи ми можемо купити 2 груші та 2 яблука?", can_buy)
}
