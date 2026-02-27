package demo

import "fmt"

type Time struct {
	time string
}

type SubTitle struct {
	title string
}

type Books struct {
	Time
	title     string
	author    string
	subject   string
	book_id   int
	sub_title *SubTitle
}

func demo() {
	var Book1 Books /* Declare Book1 of type Book */
	var Book2 Books /* Declare Book2 of type Book */

	Book3 := Books{
		title:   "C 语言",
		author:  "www.runoob.com",
		subject: "C 语言教程",
		book_id: 6495407,
		Time: Time{
			time: "2023-01-01",
		},
		sub_title: &SubTitle{
			title: "C 语言教程 说明",
		},
	}

	Book := Books{
		Time{
			time: "2026-01-01",
		},
		"js",
		"www.runoob.com",
		"js 语言教程",
		6495407,
		&SubTitle{
			title: "js 语言教程 说明",
		},
	}

	/* book 1 描述 */
	Book1.title = "Go 语言"
	Book1.author = "www.runoob.com"
	Book1.subject = "Go 语言教程"
	Book1.book_id = 6495407

	/* book 2 描述 */
	Book2.title = "Python 教程"
	Book2.author = "www.runoob.com"
	Book2.subject = "Python 语言教程"
	Book2.book_id = 6495700

	/* 打印 Book1 信息 */
	printBook(&Book1)

	/* 打印 Book2 信息 */
	printBook(&Book2)

	printBook(&Book3)
	printBook(&Book)
}
func printBook(book *Books) {
	fmt.Printf("Book title : %s\n", book.title)
	fmt.Printf("Book author : %s\n", book.author)
	fmt.Printf("Book subject : %s\n", book.subject)
	fmt.Printf("Book book_id : %d\n", book.book_id)
	fmt.Printf("Book time : %s\n", book.time)
	// fmt.Printf("Book time : %s\n", book.Time.time)
}
