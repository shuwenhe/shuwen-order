package router

import (
	"net/http"

	"github.com/shuwenhe/shuwen-shop/controller"
	"github.com/spf13/viper"
)

func Run() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("views/static"))))
	http.Handle("/pages/", http.StripPrefix("/pages/", http.FileServer(http.Dir("views/pages"))))

	Order()

	port := viper.GetString("server.port")
	http.ListenAndServe(port, nil) // Listen and serve
}

func Order()  {
	http.HandleFunc("/checkout", controller.Checkout)
	http.HandleFunc("/getOrders", controller.GetOrders)       // Get all orders
	http.HandleFunc("/getOrderInfo", controller.GetOrderInfo) // Get all orders
	http.HandleFunc("/getOrderByUserID", controller.GetOrderByUserID)
	http.HandleFunc("/sendOrder", controller.SendOrder)
	http.HandleFunc("/takeOrder", controller.TakeOrder)
}