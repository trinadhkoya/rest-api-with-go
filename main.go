package main

import (
	router "rest-api/http"
)

var httpRouter = router.NewMuxRouter()

func main() {
	const port string = "8000"
	httpRouter.SERVE(port)
}
