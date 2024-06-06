package main

import (
	"log"
	"net/http"

	"golang.org/x/net/html"
)

// getProverbs используется для получения мапы поговорок с сайта go-proverbs.github.io
func getProverbs() map[int]string {
	var proverbs = make(map[int]string)

	response, err := http.Get("https://go-proverbs.github.io/")
	if err != nil {
		log.Fatalf("ошибка получения ответа: %v\n", err)
	}
	defer response.Body.Close()

	if response.Status != "200 OK" {
		log.Fatalf("ответ не получен, код ответа: %v\n", response.Status)
		return proverbs
	}

	body, err := html.Parse(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	// extractProverbs рекурсивно обходит узлы HTML и извлекает текст вложенных узлов <a> внутри узлов <h3>
	var extractProverbs func(*html.Node)
	extractProverbs = func(n *html.Node) {
		// Проверяем, является ли текущий узел <a> и находится ли он внутри <h3>
		if n.Data == "a" && n.Parent.Data == "h3" {
			// Если условие выполняется, то обходим всех потомков узла <a> и извлекаем их текст
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				// Добавляем текст поговорки в proverbs
				proverbs[len(proverbs)+1] = c.Data
			}
		}
		// Рекурсивный вызов функции для обхода всех дочерних узлов текущего узла
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			extractProverbs(c)
		}
	}
	// Обходим HTML-структуру из ответа сайта
	extractProverbs(body)
	return proverbs
}
