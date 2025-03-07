package main

import (
    "fmt"
    "net/http"
)

// Readme (структура проекта)
// .env
// docker
// логирование в файл
// тестирование с помощью функций
// swager

// проверка на не существующий роутинг
// проверка на sql инъекции ???
// проверка на код ошибки

// time_table - id name_tablet time period start end 

func main() {
    fmt.Println("Hello, World!")
    http.HandleFunc('/', index)

    if err:= http.ListenAndServe(":8080", nil); err != nil {
        //
    }
}