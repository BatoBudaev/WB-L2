package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

// Функция для загрузки страницы и сохранения её содержимого в файл
func downloadPage(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("не удалось получить страницу: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("не удалось прочитать тело ответа: %w", err)
	}

	err = os.WriteFile("page.html", body, 0644)
	if err != nil {
		return fmt.Errorf("не удалось записать файл: %w", err)
	}

	return nil
}

// Рекурсивная функция для парсинга ссылок из HTML-документа
func parseLinks(n *html.Node) []string {
	var links []string
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
				break
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = append(links, parseLinks(c)...)
	}

	return links
}

// Функция для сканирования сайта и вывода всех найденных ссылок
func crawlSite(url string) error {
	htmlContent, err := fetchHTML(url)
	if err != nil {
		return fmt.Errorf("не удалось получить HTML: %w", err)
	}

	// Парсинг HTML-контента из io.Reader
	doc, err := html.Parse(bytes.NewReader(htmlContent))
	if err != nil {
		return fmt.Errorf("не удалось парсить HTML: %w", err)
	}

	// Получение всех ссылок из HTML-документа
	links := parseLinks(doc)

	// Вывод всех найденных ссылок
	for _, link := range links {
		fmt.Println(link)
	}

	return nil
}

// Функция для получения HTML-контента по URL
func fetchHTML(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("не удалось получить URL: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("не удалось прочитать тело ответа: %w", err)
	}

	return body, nil
}

func main() {
	url := "http://mail.ru"
	err := crawlSite(url)
	if err != nil {
		fmt.Printf("Ошибка сканирования сайта: %v\n", err)
	} else {
		fmt.Println("Сканирование завершено успешно")
	}

	err = downloadPage(url)
	if err != nil {
		fmt.Printf("Ошибка скачивания сайта: %v\n", err)
	} else {
		fmt.Println("Скачивание завершено успешно")
	}
}
