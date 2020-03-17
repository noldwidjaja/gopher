package main

func main() {

	db := newDatabase()
	db.run()

	a := NewServer("8000")
	a.routes()
	a.serve()
}
