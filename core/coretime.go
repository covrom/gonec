package core

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/covrom/gonec/names"
)

const (
	VMNanosecond  VMTimeDuration = 1
	VMMicrosecond                = 1000 * VMNanosecond
	VMMillisecond                = 1000 * VMMicrosecond
	VMSecond                     = 1000 * VMMillisecond
	VMMinute                     = 60 * VMSecond
	VMHour                       = 60 * VMMinute
	VMDay                        = 24 * VMHour
)

// VMTimeDuration - диапазон между отметками времени

type VMTimeDuration time.Duration

var ReflectVMTimeDuration = reflect.TypeOf(VMTimeDuration(0))

func (v VMTimeDuration) vmval() {}

func (v VMTimeDuration) Interface() interface{} {
	return time.Duration(v)
}

func (v VMTimeDuration) Duration() VMTimeDuration {
	return v
}

func (x VMTimeDuration) BinaryType() VMBinaryType {
	return VMDURATION
}

func (v VMTimeDuration) String() string {
	u := uint64(v)
	if u == 0 {
		return "0с"
	}
	neg := v < 0
	if neg {
		u = -u
	}

	var buf [40]byte
	w := len(buf)

	if u < uint64(VMSecond) {
		// Special case: if duration is smaller than a second,
		// use smaller units, like 1.2ms
		var prec int
		w -= 2
		copy(buf[w:], "с")
		switch {
		case u < uint64(VMMicrosecond):
			// print nanoseconds
			prec = 0
			w -= 2 // Need room for two bytes.
			copy(buf[w:], "н")
		case u < uint64(VMMillisecond):
			// print microseconds
			prec = 3
			w -= 4 // Need room for 4 bytes.
			copy(buf[w:], "мк")
		default:
			// print milliseconds
			prec = 6
			w -= 2 // Need room for two bytes.
			copy(buf[w:], "м")
		}
		w, u = fmtFrac(buf[:w], u, prec)
		w = fmtInt(buf[:w], u)
	} else {
		w -= 2
		copy(buf[w:], "с")

		w, u = fmtFrac(buf[:w], u, 9)

		// u is now integer seconds
		w = fmtInt(buf[:w], u%60)
		u /= 60

		// u is now integer minutes
		if u > 0 {
			w -= 2
			copy(buf[w:], "м")
			w = fmtInt(buf[:w], u%60)
			u /= 60

			// u is now integer hours
			if u > 0 {
				w -= 2
				copy(buf[w:], "ч")
				w = fmtInt(buf[:w], u%24)
				u /= 24

				if u > 0 {
					w -= 2
					copy(buf[w:], "д")
					w = fmtInt(buf[:w], u)
				}
			}
		}
	}

	if neg {
		w--
		buf[w] = '-'
	}

	return string(buf[w:])
}

func fmtFrac(buf []byte, v uint64, prec int) (nw int, nv uint64) {
	// Omit trailing zeros up to and including decimal point.
	w := len(buf)
	print := false
	for i := 0; i < prec; i++ {
		digit := v % 10
		print = print || digit != 0
		if print {
			w--
			buf[w] = byte(digit) + '0'
		}
		v /= 10
	}
	if print {
		w--
		buf[w] = '.'
	}
	return w, v
}

func fmtInt(buf []byte, v uint64) int {
	w := len(buf)
	if v == 0 {
		w--
		buf[w] = '0'
	} else {
		for v > 0 {
			w--
			buf[w] = byte(v%10) + '0'
			v /= 10
		}
	}
	return w
}

// EvalBinOp сравнивает два значения или выполняет бинарную операцию
func (x VMTimeDuration) EvalBinOp(op VMOperation, y VMOperationer) (VMValuer, error) {
	switch op {
	case ADD:
		switch yy := y.(type) {
		case VMTimeDuration:
			return VMTimeDuration(int64(x) + int64(yy)), nil
		case VMTime:
			return yy.Add(x), nil
		}
		return VMNil, VMErrorIncorrectOperation
	case SUB:
		switch yy := y.(type) {
		case VMTimeDuration:
			return VMTimeDuration(int64(x) - int64(yy)), nil
		}
		return VMNil, VMErrorIncorrectOperation
	case MUL:
		switch yy := y.(type) {
		case VMInt:
			return VMTimeDuration(int64(x) * int64(yy)), nil
		case VMDecNum:
			return VMTimeDuration(yy.Mul(NewVMDecNumFromInt64(int64(x))).Int()), nil
		case VMTimeDuration:
			return VMTimeDuration(int64(x) * int64(yy)), nil
		}
		return VMNil, VMErrorIncorrectOperation
	case QUO:
		switch yy := y.(type) {
		case VMTimeDuration:
			return VMTimeDuration(int64(x) / int64(yy)), nil
		}
		return VMNil, VMErrorIncorrectOperation
	case REM:
		switch yy := y.(type) {
		case VMTimeDuration:
			return VMTimeDuration(int64(x) % int64(yy)), nil
		}
		return VMNil, VMErrorIncorrectOperation
	case EQL:
		switch yy := y.(type) {
		case VMTimeDuration:
			return VMBool(int64(x) == int64(yy)), nil
		}
		return VMNil, VMErrorIncorrectOperation
	case NEQ:
		switch yy := y.(type) {
		case VMTimeDuration:
			return VMBool(int64(x) != int64(yy)), nil
		}
		return VMNil, VMErrorIncorrectOperation
	case GTR:
		switch yy := y.(type) {
		case VMTimeDuration:
			return VMBool(int64(x) > int64(yy)), nil
		}
		return VMNil, VMErrorIncorrectOperation
	case GEQ:
		switch yy := y.(type) {
		case VMTimeDuration:
			return VMBool(int64(x) >= int64(yy)), nil
		}
		return VMNil, VMErrorIncorrectOperation
	case LSS:
		switch yy := y.(type) {
		case VMTimeDuration:
			return VMBool(int64(x) < int64(yy)), nil
		}
		return VMNil, VMErrorIncorrectOperation
	case LEQ:
		switch yy := y.(type) {
		case VMTimeDuration:
			return VMBool(int64(x) <= int64(yy)), nil
		}
		return VMNil, VMErrorIncorrectOperation
	case OR:
		return VMNil, VMErrorIncorrectOperation
	case LOR:
		return VMNil, VMErrorIncorrectOperation
	case AND:
		return VMNil, VMErrorIncorrectOperation
	case LAND:
		return VMNil, VMErrorIncorrectOperation
	case POW:
		return VMNil, VMErrorIncorrectOperation
	case SHR:
		return VMNil, VMErrorIncorrectOperation
	case SHL:
		return VMNil, VMErrorIncorrectOperation
	}
	return VMNil, VMErrorUnknownOperation
}

func (x VMTimeDuration) ConvertToType(nt reflect.Type) (VMValuer, error) {
	switch nt {
	case ReflectVMString:
		return VMString(x.String()), nil
	case ReflectVMInt:
		return VMInt(int64(x)), nil
	case ReflectVMTimeDuration:
		return x, nil
	case ReflectVMDecNum:
		return NewVMDecNumFromInt64(int64(x)), nil
	}
	return VMNil, VMErrorNotConverted
}

func (x VMTimeDuration) MarshalBinary() ([]byte, error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, int64(x))
	return buf.Bytes(), nil
}

func (x *VMTimeDuration) UnmarshalBinary(data []byte) error {
	var i int64
	if err := binary.Read(bytes.NewReader(data), binary.LittleEndian, &i); err != nil {
		return err
	}
	*x = VMTimeDuration(i)
	return nil
}

func (x VMTimeDuration) GobEncode() ([]byte, error) {
	return x.MarshalBinary()
}

func (x *VMTimeDuration) GobDecode(data []byte) error {
	return x.UnmarshalBinary(data)
}

func (x VMTimeDuration) MarshalText() ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteString(time.Duration(x).String())
	return buf.Bytes(), nil
}

func (x *VMTimeDuration) UnmarshalText(data []byte) error {
	d, err := time.ParseDuration(string(data))
	if err != nil {
		return err
	}
	*x = VMTimeDuration(d)
	return nil
}

func (x VMTimeDuration) MarshalJSON() ([]byte, error) {
	b, err := x.MarshalText()
	if err != nil {
		return nil, err
	}
	return []byte("\"" + string(b) + "\""), nil
}

func (x *VMTimeDuration) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	if len(data) > 2 && data[0] == '"' && data[len(data)-1] == '"' {
		data = data[1 : len(data)-1]
	}
	return x.UnmarshalText(data)
}

//VMTime дата и время

type VMTime time.Time

var ReflectVMTime = reflect.TypeOf(VMTime{})

func Now() VMTime {
	return VMTime(time.Now())
}

func (t VMTime) vmval() {}

func (t VMTime) Interface() interface{} {
	return time.Time(t)
}

func (t VMTime) Time() VMTime {
	return t
}

func (t VMTime) BinaryType() VMBinaryType {
	return VMTIME
}

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

// String аналогичен представлению Го
// "2006-01-02 15:04:05.999999999 -0700 MST m=±ddd.nnnnnnnnn"
func (t VMTime) String() string {
	return time.Time(t).String()
}

// GolangTime возвращает time.Time
func (t VMTime) GolangTime() time.Time {
	return time.Time(t)
}

// Format формат в стиле Го
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

func (t VMTime) MethodMember(name int) (VMFunc, bool) {

	// только эти методы будут доступны из кода на языке Гонец!

	switch names.UniqueNames.GetLowerCase(name) {
	case "год":
		return VMFuncMustParams(0, t.Год), true
	case "месяц":
		return VMFuncMustParams(0, t.Месяц), true
	case "день":
		return VMFuncMustParams(0, t.День), true
	case "неделя":
		return VMFuncMustParams(0, t.Неделя), true
	case "деньнедели":
		return VMFuncMustParams(0, t.ДеньНедели), true
	case "квартал":
		return VMFuncMustParams(0, t.Квартал), true
	case "деньгода":
		return VMFuncMustParams(0, t.ДеньГода), true
	case "час":
		return VMFuncMustParams(0, t.Час), true
	case "минута":
		return VMFuncMustParams(0, t.Минута), true
	case "секунда":
		return VMFuncMustParams(0, t.Секунда), true
	case "миллисекунда":
		return VMFuncMustParams(0, t.Миллисекунда), true
	case "микросекунда":
		return VMFuncMustParams(0, t.Микросекунда), true
	case "наносекунда":
		return VMFuncMustParams(0, t.Наносекунда), true
	case "unixnano":
		return VMFuncMustParams(0, t.ЮниксНано), true
	case "unix":
		return VMFuncMustParams(0, t.Юникс), true
	case "формат":
		return VMFuncMustParams(1, t.Формат), true
	case "вычесть":
		return VMFuncMustParams(1, t.Вычесть), true
	case "добавить":
		return VMFuncMustParams(1, t.Добавить), true
	case "добавитьпериод":
		return VMFuncMustParams(3, t.ДобавитьПериод), true
	case "раньше":
		return VMFuncMustParams(1, t.Раньше), true
	case "позже":
		return VMFuncMustParams(1, t.Позже), true
	case "равно":
		return VMFuncMustParams(1, t.Равно), true
	case "пустая":
		return VMFuncMustParams(0, t.Пустая), true
	case "местное":
		return VMFuncMustParams(0, t.Местное), true
	case "utc":
		return VMFuncMustParams(0, t.ВремяUTC), true
	case "локация":
		return VMFuncMustParams(0, t.Локация), true
	case "влокации":
		return VMFuncMustParams(1, t.ВЛокации), true
	}

	return nil, false
}

func (t VMTime) Year() VMInt {
	return VMInt(time.Time(t).Year())
}

func (t VMTime) Год(args VMSlice, rets *VMSlice, envout *(*Env)) error {
	rets.Append(t.Year())
	return nil
}

func (t VMTime) Month() VMInt {
	return VMInt(time.Time(t).Month())
}

func (t VMTime) Месяц(args VMSlice, rets *VMSlice, envout *(*Env)) error {
	rets.Append(t.Month())
	return nil
}

func (t VMTime) Day() VMInt {
	return VMInt(time.Time(t).Day())
}

func (t VMTime) День(args VMSlice, rets *VMSlice, envout *(*Env)) error {
	rets.Append(t.Day())
	return nil
}

// Weekday 0=воскресенье, 1=понедельник ...
func (t VMTime) Weekday() int64 {
	return int64(time.Time(t).Weekday())
}

func (t VMTime) ISOWeek() (year, week VMInt) {
	yy, ww := time.Time(t).ISOWeek()
	return VMInt(yy), VMInt(ww)
}

// Неделя по ISO 8601
// 1-53, выровнены по годам,
// т.е. конец декабря (три дня) попадает в следующий год,
// или первые три дня января попадают в предыдущий год
func (t VMTime) Неделя(args VMSlice, rets *VMSlice, envout *(*Env)) error {
	yy, ww := t.ISOWeek()
	rets.Append(ww)
	rets.Append(yy)
	return nil
}

func (t VMTime) ДеньНедели(args VMSlice, rets *VMSlice, envout *(*Env)) error {
	//1=понедельник, 7=воскресенье ...
	wd := t.Weekday()
	if wd == 0 {
		rets.Append(VMInt(7))
	} else {
		rets.Append(VMInt(wd))
	}
	return nil
}

func (t VMTime) Quarter() VMInt {
	//1-4
	return VMInt(int64(time.Time(t).Month())/4 + 1)
}

func (t VMTime) Квартал(args VMSlice, rets *VMSlice, envout *(*Env)) error {
	//1-4
	rets.Append(t.Quarter())
	return nil
}

func (t VMTime) YearDay() VMInt {
	//1-366
	return VMInt(time.Time(t).YearDay())
}

func (t VMTime) ДеньГода(args VMSlice, rets *VMSlice, envout *(*Env)) error {
	//1-366
	rets.Append(t.YearDay())
	return nil
}

func (t VMTime) Hour() VMInt {
	return VMInt(time.Time(t).Hour())
}

func (t VMTime) Час(args VMSlice, rets *VMSlice, envout *(*Env)) error {
	rets.Append(t.Hour())
	return nil
}

func (t VMTime) Minute() VMInt {
	return VMInt(time.Time(t).Minute())
}

func (t VMTime) Минута(args VMSlice, rets *VMSlice, envout *(*Env)) error {
	rets.Append(t.Minute())
	return nil
}

func (t VMTime) Second() VMInt {
	return VMInt(time.Time(t).Second())
}

func (t VMTime) Секунда(args VMSlice, rets *VMSlice, envout *(*Env)) error {
	rets.Append(t.Second())
	return nil
}

func (t VMTime) Millisecond() VMInt {
	return VMInt(int64(time.Time(t).Nanosecond()) / 1e6)
}

func (t VMTime) Миллисекунда(args VMSlice, rets *VMSlice, envout *(*Env)) error {
	rets.Append(t.Millisecond())
	return nil
}

func (t VMTime) Microsecond() VMInt {
	return VMInt(int64(time.Time(t).Nanosecond()) / 1e3)
}

func (t VMTime) Микросекунда(args VMSlice, rets *VMSlice, envout *(*Env)) error {
	rets.Append(t.Microsecond())
	return nil
}

func (t VMTime) Nanosecond() VMInt {
	return VMInt(time.Time(t).Nanosecond())
}

func (t VMTime) Наносекунда(args VMSlice, rets *VMSlice, envout *(*Env)) error {
	rets.Append(t.Nanosecond())
	return nil
}

func (t VMTime) UnixNano() VMInt {
	return VMInt(time.Time(t).UnixNano())
}

func (t VMTime) ЮниксНано(args VMSlice, rets *VMSlice, envout *(*Env)) error {
	rets.Append(t.UnixNano())
	return nil
}

func (t VMTime) Unix() VMInt {
	return VMInt(time.Time(t).Unix())
}

func (t VMTime) Юникс(args VMSlice, rets *VMSlice, envout *(*Env)) error {
	rets.Append(t.Unix())
	return nil
}

func (t VMTime) Формат(args VMSlice, rets *VMSlice, envout *(*Env)) error {

	// аргумент - форматная строка
	fmtstr, ok := args[0].(VMString)
	if !ok {
		return VMErrorNeedString
	}

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
		"воскресенье",
		"понедельник",
		"вторник",
		"среда",
		"четверг",
		"пятница",
		"суббота",
	}

	dayssm := [...]string{
		"вс",
		"пн",
		"вт",
		"ср",
		"чт",
		"пт",
		"сб",
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

	src := []rune(string(fmtstr))
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
				res = append(res, []rune(days[t.Weekday()])...)
				i += 4
				continue
			case "ММММ", "MMMM":
				if wasday {
					res = append(res, []rune(months2[t.Month()])...)
				} else {
					res = append(res, []rune(months1[t.Month()])...)
				}
				i += 4
				continue
			case "гггг", "yyyy":
				res = append(res, []rune(strconv.FormatInt(int64(t.Year()), 10))...)
				i += 4
				continue

			}
		}

		if i+3 <= len(src) {
			s = src[i : i+3]
			switch string(s) {
			case "ддд", "ddd":
				res = append(res, []rune(dayssm[t.Weekday()])...)
				i += 3
				continue
			case "МММ", "MMM":
				if wasday {
					res = append(res, []rune(months2[t.Month()])[:3]...)
				} else {
					res = append(res, []rune(months1[t.Month()])[:3]...)
				}
				i += 3
				continue
			case "ссс", "sss":
				sm := strconv.FormatInt(int64(t.Millisecond()), 10)
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
			sm := strconv.FormatInt(int64(t.Quarter()), 10)
			res = append(res, []rune(sm)...)
			i++
			continue
		}
		res = append(res, c)
		i++
	}

	rets.Append(VMString(string(res)))
	return nil
}

func (t VMTime) Sub(t2 VMTime) VMTimeDuration {
	return VMTimeDuration(time.Time(t).Sub(time.Time(t2)))
}

func (t VMTime) Вычесть(args VMSlice, rets *VMSlice, envout *(*Env)) error {
	t2, ok := args[0].(VMTime)
	if !ok {
		return VMErrorNeedDate
	}
	rets.Append(t.Sub(t2))
	return nil
}

func (t VMTime) Add(d VMTimeDuration) VMTime {
	return VMTime(time.Time(t).Add(time.Duration(d)))
}

func (t VMTime) Добавить(args VMSlice, rets *VMSlice, envout *(*Env)) error {
	d, ok := args[0].(VMTimeDuration)
	if !ok {
		return VMErrorNeedDuration
	}
	rets.Append(t.Add(d))
	return nil
}

func (t VMTime) ДобавитьПериод(args VMSlice, rets *VMSlice, envout *(*Env)) error { //(dy, dm, dd int) VMTime {
	dy, ok := args[0].(VMInt)
	if !ok {
		return VMErrorNeedInt
	}
	dm, ok := args[1].(VMInt)
	if !ok {
		return VMErrorNeedInt
	}
	dd, ok := args[2].(VMInt)
	if !ok {
		return VMErrorNeedInt
	}
	rets.Append(VMTime(time.Time(t).AddDate(int(dy), int(dm), int(dd))))
	return nil
}

func (t VMTime) Before(d VMTime) bool {
	return time.Time(t).Before(time.Time(d))
}

func (t VMTime) Раньше(args VMSlice, rets *VMSlice, envout *(*Env)) error { //(d VMTime) bool {
	t2, ok := args[0].(VMTime)
	if !ok {
		return VMErrorNeedDate
	}
	rets.Append(VMBool(t.Before(t2)))
	return nil
}

func (t VMTime) After(d VMTime) bool {
	return time.Time(t).After(time.Time(d))
}

func (t VMTime) Позже(args VMSlice, rets *VMSlice, envout *(*Env)) error { //(d VMTime) bool {
	t2, ok := args[0].(VMTime)
	if !ok {
		return VMErrorNeedDate
	}
	rets.Append(VMBool(t.After(t2)))
	return nil
}

func (t VMTime) Equal(d VMTime) bool {
	// для разных локаций тоже работает, в отличие от =
	return time.Time(t).Equal(time.Time(d))
}

func (t VMTime) Равно(args VMSlice, rets *VMSlice, envout *(*Env)) error { //(d VMTime) bool {
	// для разных локаций тоже работает, в отличие от =
	t2, ok := args[0].(VMTime)
	if !ok {
		return VMErrorNeedDate
	}
	rets.Append(VMBool(t.Equal(t2)))
	return nil
}

func (t VMTime) IsZero() bool {
	return time.Time(t).IsZero()
}

func (t VMTime) Пустая(args VMSlice, rets *VMSlice, envout *(*Env)) error { //() bool {
	rets.Append(VMBool(t.IsZero()))
	return nil
}

func (t VMTime) Local() VMTime {
	return VMTime(time.Time(t).Local())
}

func (t VMTime) Местное(args VMSlice, rets *VMSlice, envout *(*Env)) error { //() VMTime {
	rets.Append(t.Local())
	return nil
}

func (t VMTime) UTC() VMTime {
	return VMTime(time.Time(t).UTC())
}

func (t VMTime) ВремяUTC(args VMSlice, rets *VMSlice, envout *(*Env)) error { //() VMTime {
	rets.Append(t.UTC())
	return nil
}

func (t VMTime) LocationString() string {
	return time.Time(t).Location().String()
}

func (t VMTime) Локация(args VMSlice, rets *VMSlice, envout *(*Env)) error { //() string {
	rets.Append(VMString(t.LocationString()))
	return nil
}

func (t VMTime) InLocation(name string) VMTime {
	loc, err := time.LoadLocation(name)
	if err != nil {
		panic(err)
	}
	return VMTime(time.Time(t).In(loc))
}

func (t VMTime) ВЛокации(args VMSlice, rets *VMSlice, envout *(*Env)) error { //(name string) VMTime {
	name, ok := args[0].(VMString)
	if !ok {
		return VMErrorNeedString
	}
	loc, err := time.LoadLocation(string(name))
	if err != nil {
		return err
	}
	rets.Append(VMTime(time.Time(t).In(loc)))
	return nil
}

// EvalBinOp сравнивает два значения или выполняет бинарную операцию
func (x VMTime) EvalBinOp(op VMOperation, y VMOperationer) (VMValuer, error) {
	switch op {
	case ADD:
		switch yy := y.(type) {
		case VMDurationer:
			return x.Add(yy.Duration()), nil
		}
		return VMNil, VMErrorIncorrectOperation
	case SUB:
		switch yy := y.(type) {
		case VMDurationer:
			return x.Add(VMTimeDuration(-int64(yy.Duration()))), nil
		case VMTime:
			return x.Sub(yy), nil
		}
		return VMNil, VMErrorIncorrectOperation
	case MUL:
		return VMNil, VMErrorIncorrectOperation
	case QUO:
		return VMNil, VMErrorIncorrectOperation
	case REM:
		return VMNil, VMErrorIncorrectOperation
	case EQL:
		switch yy := y.(type) {
		case VMDateTimer:
			return VMBool(x.Equal(yy.Time())), nil
		}
		return VMNil, VMErrorIncorrectOperation
	case NEQ:
		switch yy := y.(type) {
		case VMDateTimer:
			return VMBool(!x.Equal(yy.Time())), nil
		}
		return VMNil, VMErrorIncorrectOperation
	case GTR:
		switch yy := y.(type) {
		case VMDateTimer:
			return VMBool(x.After(yy.Time())), nil
		}
		return VMNil, VMErrorIncorrectOperation
	case GEQ:
		switch yy := y.(type) {
		case VMDateTimer:
			return VMBool(x.Equal(yy.Time()) || x.After(yy.Time())), nil
		}
		return VMNil, VMErrorIncorrectOperation
	case LSS:
		switch yy := y.(type) {
		case VMDateTimer:
			return VMBool(x.Before(yy.Time())), nil
		}
		return VMNil, VMErrorIncorrectOperation
	case LEQ:
		switch yy := y.(type) {
		case VMDateTimer:
			return VMBool(x.Equal(yy.Time()) || x.Before(yy.Time())), nil
		}
		return VMNil, VMErrorIncorrectOperation
	case OR:
		return VMNil, VMErrorIncorrectOperation
	case LOR:
		return VMNil, VMErrorIncorrectOperation
	case AND:
		return VMNil, VMErrorIncorrectOperation
	case LAND:
		return VMNil, VMErrorIncorrectOperation
	case POW:
		return VMNil, VMErrorIncorrectOperation
	case SHR:
		return VMNil, VMErrorIncorrectOperation
	case SHL:
		return VMNil, VMErrorIncorrectOperation
	}
	return VMNil, VMErrorUnknownOperation
}

func (x VMTime) ConvertToType(nt reflect.Type) (VMValuer, error) {
	switch nt {
	case ReflectVMString:
		// сериализуем в json
		b, err := json.Marshal(x)
		if err != nil {
			return VMNil, err
		}
		return VMString(string(b)), nil
		// case ReflectVMInt:
	case ReflectVMTime:
		return x, nil
		// case ReflectVMBool:
		// case ReflectVMDecNum:
		// case ReflectVMSlice:
		// case ReflectVMStringMap:
	}

	return VMNil, VMErrorNotConverted
}
