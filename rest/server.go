package rest
import (
	"ecomerce/config"
	"ecomerce/rest/middleware"
	"fmt"
	"net/http"
	"strconv"
)
func Start(cnf config.Config){
	manager := middleware.NewManager()
	manager.Use(
		middleware.PreFlight,
		middleware.Cors,
		middleware.Logger,
	)
	mux := http.NewServeMux()
	wrapMux:=manager.WrapMux(mux)
	InitRoute(mux, manager) // for all router
	port :=":"+ strconv.Itoa(cnf.HttpPort)
	fmt.Println("Server is running at http://localhost:",cnf.HttpPort)

	err := http.ListenAndServe(port,wrapMux )
	if err != nil {
		fmt.Println("Error running server:", err)
	}
}