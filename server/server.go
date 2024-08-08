package server

func Init() {

	r := SetRoutes()
	r.Run(":8000")
}
