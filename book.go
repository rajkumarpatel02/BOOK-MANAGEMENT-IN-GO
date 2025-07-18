package modules

type Book struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Auther      string `json:"author"`
	Publication string `json:"publication"`
}

var books []Book
var nextID int64 = 1

func GetAllBooks() []Book {
	return books
}

func GetBookById(id int64) (*Book, interface{}) {
	for _, b := range books {
		if b.ID == id {
			return &b, nil
		}
	}
	return nil, nil
}

func (b *Book) CreateBook() *Book {
	b.ID = nextID
	nextID++
	books = append(books, *b)
	return b
}

func DeleteBook(id int64) Book {
	var deleted Book
	var updated []Book
	for _, b := range books {
		if b.ID == id {
			deleted = b
			continue
		}
		updated = append(updated, b)
	}
	books = updated
	return deleted
}
