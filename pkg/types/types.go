package types

// Money представляет собой денежную сумму в минимальных единицах (cents, coins, dirams...)
type Money int64

//Category представляет собой категорию, в которой был совершён платёж (auto, pharmacy, restaurants...)
type Category string

//Payment представляет информацию о платеже
type Payment struct {
	ID       int
	Amount   Money
	Category Category
}