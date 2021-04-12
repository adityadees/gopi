package main

import "gopi/config/routes"

func main() {
	eg := routes.Routing.GetRoutes(routes.Routing{})

	_ = eg.Start(":1323") 
}