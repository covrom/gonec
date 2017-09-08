package core

import (
	"reflect"
	"time"
)

// коллекции и типы вирт. машины

type VMSlice []interface{}

type VMStringMap map[string]interface{}

///////////////////////////////////////
//Дата и время/////////////////////////
///////////////////////////////////////

type VMTime time.Time

var ReflectVMTime = reflect.TypeOf(VMTime{})

func (t VMTime) GobEncode() ([]byte, error) {
	return time.Time(t).GobEncode()
}

func (t *VMTime) GobDecode(data []byte) error {
	tt := time.Time(*t)
	err := (&tt).GobDecode(data)
	if err != nil {
		return err
	}
	*t = VMTime(tt)
	return nil
}

func (t VMTime) MarshalJSON() ([]byte, error) {
	return time.Time(t).MarshalJSON()
}

func (t *VMTime) UnmarshalJSON(data []byte) error {
	tt := time.Time(*t)
	err := (&tt).UnmarshalJSON(data)
	if err != nil {
		return err
	}
	*t = VMTime(tt)
	return nil
}

func (t VMTime) MarshalText() ([]byte, error) {
	return time.Time(t).MarshalText()
}

func (t *VMTime) UnmarshalText(data []byte) error {
	tt := time.Time(*t)
	err := (&tt).UnmarshalText(data)
	if err != nil {
		return err
	}
	*t = VMTime(tt)
	return nil
}

func (t VMTime) String() string {
	return time.Time(t).String()
}

func (t VMTime) Format(layout string) string {
	const bufSize = 64
	var b []byte
	max := len(layout) + 10
	if max < bufSize {
		var buf [bufSize]byte
		b = buf[:0]
	} else {
		b = make([]byte, 0, max)
	}
	b = time.Time(t).AppendFormat(b, layout)
	return string(b)
}

func (t VMTime) Год() int64 {
	return int64(time.Time(t).Year())
}

func (t VMTime) Месяц() int64 {
	return int64(time.Time(t).Month())
}

func (t VMTime) День() int64 {
	return int64(time.Time(t).Day())
}

func (t VMTime) Weekday() int64 {
	//0=воскресенье, 1=понедельник ...
	return int64(time.Time(t).Weekday())
}

func (t VMTime) Неделя() (w int64, y_aligned int64) {
	// по ISO 8601
	// 1-53, выровнены по годам,
	// т.е. конец декабря (три дня) попадает в следующий год,
	// или первые три дня января попадают в предыдущий год
	yy, ww := time.Time(t).ISOWeek()
	return int64(ww), int64(yy)
}

func (t VMTime) ДеньНедели() int64 {
	//1=понедельник, 7=воскресенье ...
	wd := int64(time.Time(t).Weekday())
	if wd == 0 {
		return 7
	}
	return wd
}

func (t VMTime) Квартал() int64 {
	//1-4
	return t.Месяц()/4 + 1
}

func (t VMTime) ДеньГода() int64 {
	//1-366
	return int64(time.Time(t).YearDay())
}

func (t VMTime) Час() int64 {
	return int64(time.Time(t).Hour())
}

func (t VMTime) Минута() int64 {
	return int64(time.Time(t).Minute())
}

func (t VMTime) Секунда() int64 {
	return int64(time.Time(t).Second())
}

func (t VMTime) Миллисекунда() int64 {
	return int64(time.Time(t).Nanosecond()) / 1e6
}

func (t VMTime) Микросекунда() int64 {
	return int64(time.Time(t).Nanosecond()) / 1e3
}

func (t VMTime) Наносекунда() int64 {
	return int64(time.Time(t).Nanosecond())
}

func (t VMTime) UnixNano() int64 {
	return time.Time(t).UnixNano()
}

func (t VMTime) Unix() int64 {
	return time.Time(t).Unix()
}

func (t VMTime) Формат(fmtstr string) string {
	// д (d) - день месяца (цифрами) без лидирующего нуля;
	// дд (dd) - день месяца (цифрами) с лидирующим нулем;
	// ддд (ddd) - краткое название дня недели *);
	// дддд (dddd) - полное название дня недели *);
	// М (M) - номер месяца (цифрами) без лидирующего нуля;
	// ММ (MM) - номер месяца (цифрами) с лидирующим нулем;
	// МММ (MMM) - краткое название месяца *);
	// ММММ (MMMM) - полное название месяца *);
	// к (q) - номер квартала в году;
	// г (y) - номер года без века и лидирующего нуля;
	// гг (yy) - номер года без века с лидирующим нулем;
	// гггг (yyyy) - номер года с веком;
	// ч (h) - час в 12 часовом варианте без лидирующих нулей;
	// чч (hh) - час в 12 часовом варианте с лидирующим нулем;
	// Ч (H) - час в 24 часовом варианте без лидирующих нулей;
	// ЧЧ (HH) - час в 24 часовом варианте с лидирующим нулем;
	// м (m) - минута без лидирующего нуля;
	// мм (mm) - минута с лидирующим нулем;
	// с (s) - секунда без лидирующего нуля;
	// сс (ss) - секунда с лидирующим нулем;
	// млс - миллисекунда с лидирующим нулем

	// var days = [...]string{
	// 	"понедельник",
	// 	"вторник",
	// 	"среда",
	// 	"четверг",
	// 	"пятница",
	// 	"суббота",
	// 	"воскресенье",
	// }

	// var months1 = [...]string{
	// 	"январь",
	// 	"февраль",
	// 	"март",
	// 	"апрель",
	// 	"май",
	// 	"июнь",
	// 	"июль",
	// 	"август",
	// 	"сентябрь",
	// 	"октябрь",
	// 	"ноябрь",
	// 	"декабрь",
	// }

	// var months2 = [...]string{
	// 	"января",
	// 	"февраля",
	// 	"марта",
	// 	"апреля",
	// 	"мая",
	// 	"июня",
	// 	"июля",
	// 	"августа",
	// 	"сентября",
	// 	"октября",
	// 	"ноября",
	// 	"декабря",
	// }

	// hour, min, sec := time.Time(t).Clock()
	// y, m, d := time.Time(t).Date()

	// TODO:

	return ""
}

func (t VMTime) ВычестьДату(t2 VMTime) time.Duration {
	return time.Time(t).Sub(time.Time(t2))
}

func (t VMTime) ДобавитьКДате(dy, dm, dd int) VMTime {
	return VMTime(time.Time(t).AddDate(dy, dm, dd))
}
