package main

import (
	"log"
	"testing"

	"github.com/covrom/gonec/bincode"
	"github.com/covrom/gonec/core"
	"github.com/covrom/gonec/parser"
)

func TestRun(t *testing.T) {
	env := core.NewEnv()

	script := `
	а = [](0,1000000)
	для н=1 по 1000000 цикл
	  а=а+[н]+[н*10]
	конеццикла
	к=0
	для каждого н из а цикл
	  к=к+н
	конеццикла
	сообщить(к)
	
	#gonec.exe -web -t
	#go tool pprof -svg ./gonec.exe http://localhost:5000/debug/pprof/profile?seconds=10 > cpu.svg
	
	//а=Новый("__функциональнаяструктуратест__",{"ПолеСтрока": "цщушаццке", "ПолеЦелоеЧисло": 3456})
	а=Новый("__функциональнаяструктуратест__")
	а.ПолеСтрока = "авузлхвсзщл"
	сообщить(а.ВСтроку(), а.ПолеСтрока)
	
	# массив с вложенным массивом со структурой и датой
	а=[1, 2, [3, {"а":1}, Дата("2017-08-17T09:23:00+03:00")], 4]
	а[2][1].а=Дата("2017-08-17T09:23:00+03:00")
	# приведение массива или структуры к типу "строка" означает сериализацию в json, со всеми вложенными массивами и структурами
	Сообщить(а, а[2][1].а)
	Сообщить(Строка(а))
	# приведение строки к массиву или структуре означает десериализацию из json
	Сообщить(Массив("[1,2,3,4]"))
	Сообщить(Массив(Строка(а)))
	# приведение массива к одному из типов "число", "булево", "целоечисло"
	# означает формирование массива из элементов с таким типом, 
	# при этом, каждый элемент приводится к целевому типу, сохраняя последовательность элементов в массиве
	а=[1,2,"3"]
	Сообщить(а, а[2])
	Сообщить(Строка(Число(а)))
	Сообщить(Строка(а))
	
	функция фиб(н)
	  если н = 0 тогда
		возврат 0
	  иначеесли н = 1 тогда
		возврат 1
	  конецесли
	  возврат фиб(н-1) + фиб(н-2)
	конецфункции
	
	сообщить(фиб(10))
	
	функция фибт(н0, н1, к)
		если к = 0 тогда
		  возврат н0
		иначеесли к = 1 тогда
		  возврат н1
		конецесли
		возврат фибт(н1, н0+н1, к-1)
	конецфункции
	  
	функция фиб2(н)
		возврат фибт(0, 1, н)
	конецфункции
	
	сообщить(фиб2(10))
	
	функция фиб3(н)
	  если н = 0 тогда
		возврат 0
	  иначеесли н = 1 тогда
		возврат 1
	  конецесли
	  н0, н1 = 0, 1
	  для к = н по 2 цикл
		тмп = н0 + н1
		н0 = н1
		н1 = тмп
	  конеццикла
	  возврат н1
	конецфункции
	
	сообщить(фиб3(10))
	`
	// script := `
	// Модуль ААА
	// а = 1

	// Модуль _

	// Функция ТрехкратныйВозврат()
	// 	абв = 0
	// 	Возврат 10.5, абв, ААА.а
	// КонецФункции

	// п1, п2, п3 = ТрехкратныйВозврат()
	// сообщить(п1,п2,п3)

	// функция фиб(н)
	// 	если н = 0 тогда
	// 		возврат 0
	// 	иначеесли н = 1 тогда
	// 		возврат 1
	// 	конецесли
	// 	возврат фиб(н-1) + фиб(н-2)
	// конецфункции

	// сообщить(фиб(10))

	// функция фибт(н0, н1, к)
	// 	если к = 0 тогда
	// 		возврат н0
	//   	иначеесли к = 1 тогда
	// 		возврат н1
	//   	конецесли
	//   	возврат фибт(н1, н0+н1, к-1)
	// конецфункции

	// функция фиб2(н)
	// 	возврат фибт(0, 1, н)
	// конецфункции

	// сообщить(фиб2(10))

	// функция фиб3(н)
	//   если н = 0 тогда
	// 	возврат 0
	//   иначеесли н = 1 тогда
	// 	возврат 1
	//   конецесли
	//   н0, н1 = 0, 1
	//   для к = н по 2 цикл
	// 	тмп = н0 + н1
	// 	н0 = н1
	// 	н1 = тмп
	//   конеццикла
	//   возврат н1
	// конецфункции

	// сообщить(фиб3(10))
	// `
	parser.EnableErrorVerbose()
	_, stmts, err := bincode.ParseSrc(script)
	if err != nil {
		log.Fatal(err)
	}

	_, err = bincode.Run(stmts, env)
	if err != nil {
		log.Fatal(err)
	}
}
