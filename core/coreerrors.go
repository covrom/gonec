package core

import (
	"errors"
	"fmt"
)

var (
	VMErrorNeedSinglePacketName = errors.New("Должно быть одно название пакета")
	VMErrorNeedLength           = errors.New("Значение должно иметь длину")
	VMErrorNeedLess             = errors.New("Первое значение должно быть меньше второго")
	VMErrorNeedLengthOrBoundary = errors.New("Должна быть длина диапазона или начало и конец")
	VMErrorNeedFormatAndArgs    = errors.New("Должны быть форматная строка и хотя бы один параметр")
	VMErrorSmallDecodeBuffer    = errors.New("Мало данных для декодирования")

	VMErrorNeedString   = errors.New("Требуется значение типа Строка")
	VMErrorNeedInt      = errors.New("Требуется значение типа ЦелоеЧисло")
	VMErrorNeedDate     = errors.New("Требуется значение типа Дата")
	VMErrorNeedMap    = errors.New("Требуется значение типа Структура")
	VMErrorNeedSlice    = errors.New("Требуется значение типа Массив")
	VMErrorNeedDuration = errors.New("Требуется значение типа Длительность")
	VMErrorNeedSeconds  = errors.New("Должно быть число секунд (допустимо с дробной частью)")
	VMErrorNeedHash     = errors.New("Параметр не может быть хэширован")

	VMErrorIndexOutOfBoundary  = errors.New("Индекс находится за пределами массива")
	VMErrorNotConverted        = errors.New("Приведение к типу невозможно")
	VMErrorUnknownType         = errors.New("Неизвестный тип данных")
	VMErrorIncorrectFieldType  = errors.New("Поле структуры имеет другой тип")
	VMErrorIncorrectStructType = errors.New("Невозможно использовать данный тип структуры")
	VMErrorNotDefined          = errors.New("Не определено")
	VMErrorNotBinaryConverted  = errors.New("Значение не может быть преобразовано в бинарный формат")

	VMErrorNoNeedArgs = errors.New("Параметры не требуются")
	VMErrorNoArgs     = errors.New("Отсутствуют аргументы")

	VMErrorIncorrectOperation = errors.New("Операция между значениями невозможна")
	VMErrorUnknownOperation   = errors.New("Неизвестная операция")

	VMErrorServerNowOnline   = errors.New("Сервер уже запущен")
	VMErrorServerOffline     = errors.New("Сервер уже остановлен")
	VMErrorIncorrectProtocol = errors.New("Неверный протокол")
	VMErrorIncorrectClientId = errors.New("Неверный идентификатор соединения")
	VMErrorIncorrectMessage  = errors.New("Неверный формат сообщения")
	VMErrorEOF  = errors.New("Недостаточно данных в источнике")

	VMErrorServiceNotReady  = errors.New("Сервис не готов") // устанавливается сервисами в случае прекращения работы
	VMErrorServiceAlreadyRegistered  = errors.New("Сервис уже зарегистрирован с таким же ID")
	VMErrorServerAlreadyStarted  = errors.New("Сервер уже запущен")
	VMErrorWrongHTTPMethod  = errors.New("Метод не применим к HTTP-соединению")
	VMErrorNonHTTPMethod  = errors.New("Метод применим только к HTTP-соединению")
)

func VMErrorNeedArgs(n int) error {
	return fmt.Errorf("Неверное количество параметров (требуется %d)", n)
}
