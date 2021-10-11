package main

func main() {
	app := App{}
	app.connectDatabase()
	app.initialiseRoutes()
	app.run()
}
