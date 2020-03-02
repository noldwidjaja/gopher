package main

func main() {
	a := NewServer("8000")
	a.routes()
	a.serve()
}
