package main

import (
	"rest-api/controller"
	router "rest-api/http"
)

var httpRouter router.Router = router.NewMuxRouter()
var utilController controller.UtilsController = controller.NewUtilsController()

func main() {
	const port string = ":5000"

	httpRouter.GET("/health", utilController.GetHealth)
	httpRouter.GET("/", utilController.Welcome)
	httpRouter.GET("/status", utilController.GetHealth)

	httpRouter.SERVE(port)
}
