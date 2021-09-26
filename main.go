package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type RespJson struct {
	Value string
}

func httpRequest(url string) []byte{
	//Создать http клиента
	client := http.Client{}
	//Создать get запрос к серверу
	resp, err := client.Get(url)
	//Обработчик ошибки запроса к серверу
	if err != nil {
		fmt.Println(err)
		return nil
	}
	//Закрыть тело ответа в конце
	defer resp.Body.Close()
	//Конвертация тела http ответа в строчку
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	bodyJson := buf.String()

	return []byte(bodyJson)
}

func ParseRespJson(bodyJson []byte) []byte{
	//Создать переменную типа структуры для парсинга json ответа
	var res RespJson
	//Распасить json ответ в переменную
	json.Unmarshal(bodyJson,&res)
	return []byte(res.Value)
}


func main() {
	//============================================CTEGORY===============================================================
	//Получить список категорий
	bCat:=httpRequest("https://api.chucknorris.io/jokes/categories")
	//Убираем символы "[" и "]" из строки
	str:=strings.ReplaceAll(string(bCat), "[","")
	str=strings.ReplaceAll(str, "]","")
	//Создаю массив категорий
	split := strings.Split(str, ",")

	//Перебор категорий
	for i := 0; i <= len(split)-1; i++ {
		//Выполняем http запрос указанное кол-во раз
		category := strings.ReplaceAll(split[i],"\"","")
		bRespCat := httpRequest("https://api.chucknorris.io/jokes/random?category=" + category)
		//Распарсить json
		bRespParseCat := ParseRespJson(bRespCat)
		//Создать текстовый файл
		file, err := os.Create(category + ".txt")
		//Обработчик ошибки запроса к серверу
		if err != nil {
			fmt.Println(err)
			return
		}
		//Закрыть текстовый файл в конце
		defer file.Close()
		//Записать текст в файл
		file.WriteString(string(bRespParseCat))
	}
	//============================================CTEGORY===============================================================

	//============================================RANDOM================================================================
	//Получаем значение переменной среды (если это значение пользователь установил)
	getEnvNum:=os.Getenv("SET_NUMBER_JOKES")
	//Преобразовать строку в числовой тип для использования в счетчике цикла
	n, err := strconv.Atoi(getEnvNum)
	//Обработчик ошибки в случае преобразования типов данных (если ошибка то используется дефолтное значение = 5)
	if err != nil {
		n=5
	}
	//Выполняем http запрос указанное кол-во раз
	for i := 1; i <= n; i++ {
		bRandom:=httpRequest("https://api.chucknorris.io/jokes/random")
		//Распарсить json
		bRespParseRandom:=ParseRespJson(bRandom)
		//Печать случайной шутки в консоль
		fmt.Printf("%s\n",string(bRespParseRandom))
	}
	//============================================RANDOM================================================================
}