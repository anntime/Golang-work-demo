package main
import "fmt"
type Books struct{
	title string
	author string
	subject string
	book_id int
}
func main(){
	var Book1 Books
	var Book2 Books
	Book1.title = "Go语言"
	Book1.author = "www3"
	Book1.subject = "Go lang"
	Book1.book_id = 2345

	Book2.title = "Python语言"
	Book2.author = "www3"
	Book2.subject = "Python"
	Book2.book_id = 4567
	
	printBook(&Book1)
	printBook(&Book2)

}
func printBook(book *Books){
	fmt.Println("Book title:",book.title);
	fmt.Println("Book title:\n",book.author);
	fmt.Println("Book title: %s\n",book.subject);
	fmt.Println("Book title: %d\n",book.book_id);

}