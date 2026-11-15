package actioninfo

import (
	"fmt"
	"log"
)

type DataParser interface {
	Parse(datastring string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	for _, data := range dataset {
		if err := dp.Parse(data); err != nil {
			log.Println("ошибка парсинга данных:", err)
			continue
		}

		info, err := dp.ActionInfo()
		if err != nil {
			log.Println("ошибка формирования информации об активности:", err)
			continue
		}

		fmt.Println(info)
	}
}
