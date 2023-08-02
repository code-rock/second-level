package main

import (
	"flag"
	"fmt"
	"main/internal/server"
	"main/internal/store"
	"os"

	"github.com/spf13/viper"
)

func main() {
	// Получение пути к конфигу из атребутов командной строки
	var configFile *string = flag.String("c", "config.yaml", "Setting the configuration file")
	flag.Parse()

	// Возврат описания файла или информации об ошибке
	_, err := os.Stat(*configFile)
	if err == nil {
		// Использования пути конфига заданного пользователем
		fmt.Println("Using User Specified Configuration file!")
		viper.SetConfigFile(*configFile)
	} else {
		// Использования дефолтного имени аргумента командной строки
		viper.SetConfigName(*configFile)
		// Для поиска соответствующего конфига в текущей дериктории
		viper.AddConfigPath(".")
	}

	// Чтение и синтаксичский анализ YAML-файла
	err = viper.ReadInConfig()
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	// Возвращает файл, используемый для заполнения реестра конфигурации.
	fmt.Printf("Using config: %s\n", viper.ConfigFileUsed())

	var PORT interface{}
	if viper.IsSet("server.port") {
		// Получение номера порта для подключения к серверу
		PORT = viper.Get("server.port")
		fmt.Println("server.val_port", PORT)
	} else {
		fmt.Println("server.port not set!")
	}

	// Создания хранилища данных
	es := store.New()
	// Создание http-сервера с доступом в хранилище данных
	serverEvents := server.New(es)
	// Запуск http-сервера на порту
	serverEvents.Api(PORT)
}
