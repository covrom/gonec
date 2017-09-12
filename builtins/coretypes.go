package core

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/shopspring/decimal"
)

// иерархия базовых типов вирт. машины
type (
	// VMValuer корневой тип всех значений, доступных вирт. машине
	VMValuer interface {
		vmval()
	}

	// VMInterfacer корневой тип всех значений,
	// которые могут преобразовываться в значения для функций на языке Го в родные типы Го
	VMInterfacer interface {
		VMValuer
		Interface() interface{} // в типах Го, может возвращать в т.ч. nil
	}

	// VMFromGoParser может парсить из значений на языке Го
	VMFromGoParser interface {
		VMValuer
		ParseGoType(interface{}) // используется для указателей, т.к. парсит в их значения
	}

	// VMParser может парсить из строки
	VMParser interface {
		VMValuer
		Parse(string) error // используется для указателей, т.к. парсит в их значения
	}

	VMChaner interface {
		VMValuer
		Send(VMValuer)
		Recv() VMValuer
		TrySend(VMValuer) bool
		TryRecv() (VMValuer, bool)
	}

	// конкретные типы виртуальной машины

	// VMStringer строка
	VMStringer interface {
		VMInterfacer
		String() string
	}

	// VMNumberer число, внутреннее хранение в int64 или decimal формате
	VMNumberer interface {
		VMInterfacer
		Int() int64
		Float() float64
		Decimal() VMDecimal
	}

	// VMBooler сообщает значение булево
	VMBooler interface {
		VMInterfacer
		Bool() bool
	}

	// VMSlicer может быть представлен в виде слайса Гонец
	VMSlicer interface {
		VMInterfacer
		Slice() VMSlice
	}

	// VMMaper может быть представлен в виде структуры Гонец
	VMStringMaper interface {
		VMInterfacer
		StringMap() VMStringMap
	}

	// VMFuncer это функция Гонец
	VMFuncer interface {
		VMInterfacer
		Func() VMFunc
	}

	// VMDateTimer это дата/время
	VMDateTimer interface {
		VMInterfacer
		Time() VMTime
	}

	// VMChanMaker может создать новый канал
	VMChanMaker interface {
		VMInterfacer
		MakeChan(int) VMChaner
	}
)

// коллекции и типы вирт. машины

// VMInt для ускорения работы храним целочисленное представление отдельно от decimal
type VMInt int64

func (x VMInt) vmval() {}

func (x VMInt) Interface() interface{} {
	return int64(x)
}

func (x *VMInt) ParseGoType(v interface{}) {
	switch vv := v.(type) {
	case int:
		*x = VMInt(vv)
	case int8:
		*x = VMInt(vv)
	case int16:
		*x = VMInt(vv)
	case int32:
		*x = VMInt(vv)
	case int64:
		*x = VMInt(vv)
	case uint:
		*x = VMInt(vv)
	case uint8:
		*x = VMInt(vv)
	case uint16:
		*x = VMInt(vv)
	case uint32:
		*x = VMInt(vv)
	case uint64:
		*x = VMInt(vv)
	case uintptr:
		*x = VMInt(vv)
	case float32:
		*x = VMInt(int64(vv))
	case float64:
		*x = VMInt(int64(vv))
	default:
		rv := reflect.Indirect(reflect.ValueOf(v))
		if rv.Kind() == reflect.Interface {
			rv = rv.Elem()
		}
		*x = VMInt(rv.Int()) // выдаст панику, если это не число
	}
}

func (x VMInt) String() string {
	return strconv.FormatInt(int64(x), 10)
}

func (x VMInt) Int() int64 {
	return int64(x)
}

func (x VMInt) Float() float64 {
	return float64(x)
}

func (x VMInt) Decimal() VMDecimal {
	return VMDecimal(decimal.New(int64(x), 0))
}

func (x VMInt) Bool() bool {
	return x > 0
}

func (x VMInt) MakeChan(size int) VMChaner {
	return make(VMChan, size)
}

func (x VMInt) Time() VMTime {
	return VMTime(time.Unix(int64(x), 0))
}

func (x *VMInt) Parse(s string) error {
	i64, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return err
	}
	*x = VMInt(i64)
	return nil
}

// VMDecimal с плавающей токой, для финансовых расчетов высокой точности (decimal)
type VMDecimal decimal.Decimal

func (x VMDecimal) vmval() {}

func (x VMDecimal) Interface() interface{} {
	return decimal.Decimal(x)
}

func (x *VMDecimal) ParseGoType(v interface{}) {
	switch vv := v.(type) {
	case int:
		*x = VMDecimal(decimal.New(int64(vv), 0))
	case int8:
		*x = VMDecimal(decimal.New(int64(vv), 0))
	case int16:
		*x = VMDecimal(decimal.New(int64(vv), 0))
	case int32:
		*x = VMDecimal(decimal.New(int64(vv), 0))
	case int64:
		*x = VMDecimal(decimal.New(vv, 0))
	case uint:
		*x = VMDecimal(decimal.New(int64(vv), 0))
	case uint8:
		*x = VMDecimal(decimal.New(int64(vv), 0))
	case uint16:
		*x = VMDecimal(decimal.New(int64(vv), 0))
	case uint32:
		*x = VMDecimal(decimal.New(int64(vv), 0))
	case uint64:
		*x = VMDecimal(decimal.New(int64(vv), 0))
	case uintptr:
		*x = VMDecimal(decimal.New(int64(vv), 0))
	case float32:
		*x = VMDecimal(decimal.NewFromFloat(float64(vv)))
	case float64:
		*x = VMDecimal(decimal.NewFromFloat(vv))
	default:
		rv := reflect.Indirect(reflect.ValueOf(v))
		if rv.Kind() == reflect.Interface {
			rv = rv.Elem()
		}
		if rv.Kind() == reflect.Float32 || rv.Kind() == reflect.Float64 {
			*x = VMDecimal(decimal.NewFromFloat(rv.Float()))
		} else {
			*x = VMDecimal(decimal.New(rv.Int(), 0)) // выдаст панику, если это не число
		}
	}
}

func (x VMDecimal) String() string {
	return decimal.Decimal(x).String()
}

func (x VMDecimal) Int() int64 {
	return decimal.Decimal(x).IntPart()
}

func (x VMDecimal) Float() float64 {
	f64, ok := decimal.Decimal(x).Float64()
	if !ok {
		panic("Невозможно получить значение с плавающей запятой 64 бит")
	}
	return f64
}

func (x VMDecimal) Decimal() VMDecimal {
	return x
}

func (x VMDecimal) Bool() bool {
	return decimal.Decimal(x).GreaterThan(decimal.Zero)
}

func (x VMDecimal) MakeChan(size int) VMChaner {

	return make(VMChan, size)
}

func (x VMDecimal) Time() VMTime {
	intpart := decimal.Decimal(x).IntPart()
	nanopart := decimal.Decimal(x).Sub(decimal.New(intpart, 0)).Mul(decimal.New(1e9, 0)).IntPart()
	return VMTime(time.Unix(intpart, nanopart))
}

func (x *VMDecimal) Parse(s string) error {
	d, err := decimal.NewFromString(s)
	if err != nil {
		return err
	}
	*x = VMDecimal(d)
	return nil
}

// VMString строки
type VMString string

func (x VMString) vmval() {}

func (x VMString) Interface() interface{} {
	return string(x)
}

func (x VMString) String() string {
	return string(x)
}

func (x VMString) Int() int64 {
	i64, err := strconv.ParseInt(string(x), 10, 64)
	if err != nil {
		panic(err)
	}
	return i64
}

func (x VMString) Float() float64 {
	f64, err := strconv.ParseFloat(string(x), 64)
	if err != nil {
		panic(err)
	}
	return f64
}

func (x VMString) Decimal() VMDecimal {
	d, err := decimal.NewFromString(string(x))
	if err != nil {
		panic(err)
	}
	return VMDecimal(d)
}

func (x VMString) MakeChan(size int) VMChaner {
	return make(VMChan, size)
}

func (x VMString) Time() VMTime {
	t, err := time.ParseInLocation("2006-01-02T15:04:05", string(x), time.Local)
	if err == nil {
		return VMTime(t)
	}
	t, err = time.Parse(time.RFC3339, string(x))
	if err == nil {
		return VMTime(t)
	}
	t, err = time.ParseInLocation("20060102150405", string(x), time.Local)
	if err == nil {
		return VMTime(t)
	}
	t, err = time.ParseInLocation("20060102", string(x), time.Local)
	if err == nil {
		return VMTime(t)
	}
	t, err = time.ParseInLocation("02.01.2006", string(x), time.Local)
	if err == nil {
		return VMTime(t)
	}
	t, err = time.ParseInLocation("02.01.2006 15:04:05", string(x), time.Local)
	if err == nil {
		return VMTime(t)
	}
	t, err = time.Parse(time.RFC1123, string(x))
	if err == nil {
		return VMTime(t)
	}
	panic("Неверный формат даты и времени")
}

func (x *VMString) Parse(s string) error {
	*x = VMString(s)
	return nil
}

func (x VMString) Slice() VMSlice {
	var rm []interface{}
	if err := json.Unmarshal([]byte(x), &rm); err != nil {
		panic(err)
	}
	return VMSlice(rm)
}

func (x VMString) StringMap() VMStringMap {
	var rm map[string]interface{}
	if err := json.Unmarshal([]byte(x), rm); err != nil {
		panic(err)
	}
	return VMStringMap(rm)
}

// VMChan - канал для передачи любого типа вирт. машины
type VMChan chan VMValuer

func (x VMChan) vmval() {}

func (x VMChan) Send(v VMValuer) {
	x <- v
}

func (x VMChan) Recv() VMValuer {
	return <-x
}

func (x VMChan) TrySend(v VMValuer) (ok bool) {
	select {
	case x <- v:
		ok = true
	default:
		ok = false
	}
	return
}

func (x VMChan) TryRecv() (v VMValuer, ok bool) {
	select {
	case v = <-x:
		ok = true
	default:
		ok = false
	}
	return
}

// старые типы

type VMSlice []interface{}

type VMStringMap map[string]interface{}

type VMChannel chan interface{}

// Функции такого типа создаются на языке Гонец,
// их можно использовать в стандартной библиотеке, проверив на этот тип
type VMFunc func(args ...interface{}) (interface{}, error)

func (f VMFunc) String() string {
	return fmt.Sprintf("[Функция: %p]", f)
}

///////////////////////////////////////
//Дата и время/////////////////////////
///////////////////////////////////////

type VMTime time.Time

var ReflectVMTime = reflect.TypeOf(VMTime{})

func (t VMTime) MarshalBinary() ([]byte, error) {
	return time.Time(t).MarshalBinary()
}

func (t *VMTime) UnmarshalBinary(data []byte) error {
	tt := time.Time(*t)
	err := (&tt).UnmarshalBinary(data)
	if err != nil {
		return err
	}
	*t = VMTime(tt)
	return nil
}

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
	// формат в стиле Го
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

func (t VMTime) GolangTime() time.Time {
	return time.Time(t)
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

	// К (Q) - номер квартала в году;
	// г (y) - номер года без века и лидирующего нуля;
	// гг (yy) - номер года без века с лидирующим нулем;
	// гггг (yyyy) - номер года с веком;

	// ч (h) - час в 24 часовом варианте без лидирующих нулей;
	// чч (hh) - час в 24 часовом варианте с лидирующим нулем;

	// м (m) - минута без лидирующего нуля;
	// мм (mm) - минута с лидирующим нулем;

	// с (s) - секунда без лидирующего нуля;
	// сс (ss) - секунда с лидирующим нулем;
	// ссс (sss) - миллисекунда с лидирующим нулем

	days := [...]string{
		"", //0-го не бывает
		"понедельник",
		"вторник",
		"среда",
		"четверг",
		"пятница",
		"суббота",
		"воскресенье",
	}

	months1 := [...]string{
		"", //0-го не бывает
		"январь",
		"февраль",
		"март",
		"апрель",
		"май",
		"июнь",
		"июль",
		"август",
		"сентябрь",
		"октябрь",
		"ноябрь",
		"декабрь",
	}

	months2 := [...]string{
		"", //0-го не бывает
		"января",
		"февраля",
		"марта",
		"апреля",
		"мая",
		"июня",
		"июля",
		"августа",
		"сентября",
		"октября",
		"ноября",
		"декабря",
	}

	dayssm := [...]string{
		"пн",
		"вт",
		"ср",
		"чт",
		"пт",
		"сб",
		"вс",
	}

	src := []rune(fmtstr)
	res := make([]rune, 0, len(src)*2)
	wasday := false
	hour, min, sec := time.Time(t).Clock()
	y, m, d := time.Time(t).Date()

	i := 0
	for i < len(src) {
		var s []rune

		if i+4 <= len(src) {
			s = src[i : i+4]
			switch string(s) {
			case "дддд", "dddd":
				res = append(res, []rune(days[t.ДеньНедели()])...)
				i += 4
				continue
			case "ММММ", "MMMM":
				if wasday {
					res = append(res, []rune(months2[t.Месяц()])...)
				} else {
					res = append(res, []rune(months1[t.Месяц()])...)
				}
				i += 4
				continue
			case "гггг", "yyyy":
				res = append(res, []rune(strconv.FormatInt(t.Год(), 10))...)
				i += 4
				continue

			}
		}

		if i+3 <= len(src) {
			s = src[i : i+3]
			switch string(s) {
			case "ддд", "ddd":
				res = append(res, []rune(dayssm[t.ДеньНедели()])...)
				i += 3
				continue
			case "МММ", "MMM":
				if wasday {
					res = append(res, []rune(months2[t.Месяц()])[:3]...)
				} else {
					res = append(res, []rune(months1[t.Месяц()])[:3]...)
				}
				i += 3
				continue
			case "ссс", "sss":
				sm := strconv.FormatInt(t.Миллисекунда(), 10)
				if len(sm) < 3 {
					sm = strings.Repeat("0", 3-len(sm)) + sm
				}
				res = append(res, []rune(sm)...)
				i += 3
				continue
			}
		}

		if i+2 <= len(src) {
			s = src[i : i+2]
			switch string(s) {
			case "дд", "dd":
				sm := strconv.Itoa(d)
				if len(sm) < 2 {
					sm = "0" + sm
				}
				res = append(res, []rune(sm)...)
				i += 2
				wasday = true
				continue
			case "ММ", "MM":
				sm := strconv.Itoa(int(m))
				if len(sm) < 2 {
					sm = "0" + sm
				}
				res = append(res, []rune(sm)...)
				i += 2
				continue
			case "гг", "yy":
				sm := strconv.Itoa(int(y % 100))
				if len(sm) < 2 {
					sm = "0" + sm
				}
				res = append(res, []rune(sm)...)
				i += 2
				continue
			case "чч", "hh":
				sm := strconv.Itoa(int(hour))
				if len(sm) < 2 {
					sm = "0" + sm
				}
				res = append(res, []rune(sm)...)
				i += 2
				continue
			case "мм", "mm":
				sm := strconv.Itoa(int(min))
				if len(sm) < 2 {
					sm = "0" + sm
				}
				res = append(res, []rune(sm)...)
				i += 2
				continue
			case "сс", "ss":
				sm := strconv.Itoa(int(sec))
				if len(sm) < 2 {
					sm = "0" + sm
				}
				res = append(res, []rune(sm)...)
				i += 2
				continue
			}
		}

		c := src[i]
		switch c {
		case 'д', 'd':
			sm := strconv.Itoa(d)
			res = append(res, []rune(sm)...)
			i++
			wasday = true
			continue
		case 'М', 'M':
			sm := strconv.Itoa(int(m))
			res = append(res, []rune(sm)...)
			i++
			continue
		case 'г', 'y':
			sm := strconv.Itoa(int(y % 100))
			res = append(res, []rune(sm)...)
			i++
			continue
		case 'ч', 'h':
			sm := strconv.Itoa(int(hour))
			res = append(res, []rune(sm)...)
			i++
			continue
		case 'м', 'm':
			sm := strconv.Itoa(int(min))
			res = append(res, []rune(sm)...)
			i++
			continue
		case 'с', 's':
			sm := strconv.Itoa(int(sec))
			res = append(res, []rune(sm)...)
			i++
			continue
		case 'К', 'Q':
			sm := strconv.FormatInt(t.Квартал(), 10)
			res = append(res, []rune(sm)...)
			i++
			continue
		}
		res = append(res, c)
		i++
	}

	return string(res)
}

func (t VMTime) Вычесть(t2 VMTime) time.Duration {
	return time.Time(t).Sub(time.Time(t2))
}

func (t VMTime) Добавить(d time.Duration) VMTime {
	return VMTime(time.Time(t).Add(d))
}

func (t VMTime) ДобавитьПериод(dy, dm, dd int) VMTime {
	return VMTime(time.Time(t).AddDate(dy, dm, dd))
}

func (t VMTime) Раньше(d VMTime) bool {
	return time.Time(t).Before(time.Time(d))
}

func (t VMTime) Позже(d VMTime) bool {
	return time.Time(t).After(time.Time(d))
}

func (t VMTime) Равно(d VMTime) bool {
	// для разных локаций тоже работает, в отличие от =
	return time.Time(t).Equal(time.Time(d))
}

func (t VMTime) Пустая() bool {
	return time.Time(t).IsZero()
}

func (t VMTime) Местное() VMTime {
	return VMTime(time.Time(t).Local())
}

func (t VMTime) UTC() VMTime {
	return VMTime(time.Time(t).UTC())
}

func (t VMTime) Локация() string {
	return time.Time(t).Location().String()
}

func (t VMTime) ВЛокации(name string) VMTime {
	loc, err := time.LoadLocation(name)
	if err != nil {
		panic(err)
	}
	return VMTime(time.Time(t).In(loc))
}
