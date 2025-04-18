package logger

import (
	"io"
	"net"
	"os"
	"sync"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	globalLogger zerolog.Logger
	once         sync.Once
)

// Инициализация логгера
func InitLogger(serviceName string) zerolog.Logger {
	once.Do(func() {
		var writer io.Writer
		conn, err := net.Dial("tcp", "logstash:5000")

		// Если не удалось подключиться к Logstash, пытаемся переподключиться
		if err != nil {
			log.Warn().Err(err).Msg("Не удалось подключиться к Logstash, пробуем переподключиться...")
			go attemptReconnect()                          // Пытаемся переподключиться в отдельной горутине
			writer = zerolog.ConsoleWriter{Out: os.Stdout} // В случае ошибки выводим в консоль
		} else {
			writer = zerolog.MultiLevelWriter(
				zerolog.ConsoleWriter{Out: os.Stdout},
				conn,
			)
		}

		// Инициализация глобального логгера
		globalLogger = zerolog.New(writer).
			With().
			Timestamp().
			Str("service", serviceName).
			Logger()

		log.Logger = globalLogger
	})

	return globalLogger
}

// Получить глобальный логгер
func GetLogger() zerolog.Logger {
	return globalLogger
}

// Функция для переподключения к Logstash с повторной попыткой
func attemptReconnect() {
	for {
		conn, err := net.Dial("tcp", "logstash:5000")
		if err == nil {
			log.Info().Msg("Подключено к Logstash")
			// Если подключение успешно, обновляем глобальный логгер
			writer := zerolog.MultiLevelWriter(
				zerolog.ConsoleWriter{Out: os.Stdout},
				conn,
			)

			globalLogger = zerolog.New(writer).
				With().
				Timestamp().
				Logger()

			log.Logger = globalLogger
			break // Завершаем цикл, если подключение успешно
		}
		// Пробуем переподключиться через 5 секунд
		log.Warn().Err(err).Msg("Ошибка подключения к Logstash, попытка переподключения через 5 секунд")
		time.Sleep(5 * time.Second)
	}
}
