package main

import (
	"github.com/go-martini/martini"
	"code.google.com/p/goprotobuf/proto"
	"net/http"
)

func main() {
	m := martini.Classic()
	m.Get("/books/:id", book)
	m.Run()
}

func book(w http.ResponseWriter, r *http.Request) {
	book := &Book{
		BookId: proto.String("1"),
		Title: proto.String("Harry Potter And Something!"),
		PageCount: proto.Int32(322),
		AuthorName: proto.String("Foo"),
	}
	data, _ := proto.Marshal(book)

	_, _ = w.Write(data)
}
