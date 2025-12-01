package main

import (
	"fmt"
)

type bookmarkMap = map[string]string

func main() {
	const menuItems string = "Введите следующее действие: \n 1 - Посмотреть закладки \n 2 - Создать закладку \n 3 - Удалить закладку \n 4 - Выход"
	var param int
	var bookmarks = bookmarkMap{"PurpleSchool": "https://app.purpleschool.ru/"}

	Loop:
	for {
		fmt.Println(menuItems)

        fmt.Scan(&param)

        switch param {
        case 1:
			result := printBookmarks(bookmarks)
			fmt.Println(result)
		case 2:
			bookmarks = addBookmark(bookmarks)
			fmt.Println("Закладка успешно добавлена")
		case 3:
			bookmarks = removeBookmark(bookmarks)
			fmt.Println("Закладка успешно удалена")
		case 4:
			break Loop
        default:
            fmt.Println("Попробуйте еще раз:")
        }
    }
}

func printBookmarks (bookmarks bookmarkMap) string {
	var result string

	if (len(bookmarks) == 0) {
		return "Список закладок пуст"
	}

	for key, value := range bookmarks {
		result += fmt.Sprintf("Ключ: %v, Значение: %v \n", key, value)
	}

	return result
}

func addBookmark (bookmarks bookmarkMap) bookmarkMap {
	var key string
	var value string

	fmt.Print("Введите ключ закладдки:")
    fmt.Scan(&key)
	fmt.Print("Введите значение закладки:")
    fmt.Scan(&value)
	
	bookmarks[key] = value

	return bookmarks
}

func removeBookmark(bookmarks bookmarkMap) bookmarkMap {
	var key string

	fmt.Print("Введите ключ закладдки для удаления:")
    fmt.Scan(&key)

	delete(bookmarks, key)
	
	return bookmarks
}