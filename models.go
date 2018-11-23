package main

//创建了一个book实体，包括书名，json，作者，页数
type Book struct {
	// The main identifier for the Book. This will be unique.
	ISDN   string `json:"isdn"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Pages  int    `json:"pages"`
}


//存储书籍的事例

var bookstore = make(map[string]*Book)
