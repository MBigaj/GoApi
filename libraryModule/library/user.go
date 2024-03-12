package library

type User struct {
	Id         uint16
	Name       string
	Age        uint8
	UserOrders []*Order
}

func (User *User) AddOrder(order *Order) {
	User.UserOrders = append(User.UserOrders, order)
}
