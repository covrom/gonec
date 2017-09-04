package bincode

import (
	"encoding/json"
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/covrom/gonec/ast"
	envir "github.com/covrom/gonec/env"
)

func InvokeNumber(lit string) (interface{}, error) {
	if strings.Contains(lit, ".") || strings.Contains(lit, "e") || strings.Contains(lit, "E") {
		v, err := strconv.ParseFloat(lit, 64)
		if err != nil {
			return v, err
		}
		return v, nil
	}
	var i int64
	var err error
	if strings.HasPrefix(lit, "0x") {
		i, err = strconv.ParseInt(lit[2:], 16, 64)
	} else {
		i, err = strconv.ParseInt(lit, 10, 64)
	}
	if err != nil {
		return i, err
	}
	return i, nil
}

func ToString(v interface{}) string {
	if s, ok := v.(string); ok {
		return s
	}
	if v == nil {
		return "Неопределено"
	}
	if b, ok := v.(bool); ok {
		if b {
			return "Истина"
		} else {
			return "Ложь"
		}
	}
	return fmt.Sprint(v)
}

func ToBool(v interface{}) bool {

	switch v.(type) {
	case float32, float64:
		return ToFloat64(v) != 0.0
	case int, int32, int64:
		return ToInt64(v) != 0
	case bool:
		return v.(bool)
	case string:
		vlow := strings.ToLower(v.(string))
		if vlow == "true" || vlow == "истина" {
			return true
		}
		if ToInt64(v) != 0 {
			return true
		}
	}
	return false
}

func ToFloat64(v interface{}) float64 {
	switch x := v.(type) {
	case float32:
		return float64(x)
	case float64:
		return x
	case int:
		return float64(x)
	case int8:
		return float64(x)
	case int16:
		return float64(x)
	case int32:
		return float64(x)
	case int64:
		return float64(x)
	case uint:
		return float64(x)
	case uint8:
		return float64(x)
	case uint16:
		return float64(x)
	case uint32:
		return float64(x)
	case uint64:
		return float64(x)
	}
	return 0.0
}

func ToInt64(v interface{}) int64 {
	switch x := v.(type) {
	case float32:
		return int64(x)
	case float64:
		return int64(x)
	case int:
		return int64(x)
	case int8:
		return int64(x)
	case int16:
		return int64(x)
	case int32:
		return int64(x)
	case int64:
		return x
	case uint:
		return int64(x)
	case uint8:
		return int64(x)
	case uint16:
		return int64(x)
	case uint32:
		return int64(x)
	case uint64:
		return int64(x)
	case string:
		var i int64
		var err error
		if strings.HasPrefix(x, "0x") {
			i, err = strconv.ParseInt(x, 16, 64)
		} else {
			i, err = strconv.ParseInt(x, 10, 64)
		}
		if err == nil {
			return i
		}
	}
	return 0
}

func IsNum(v interface{}) bool {
	switch v.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr, float32, float64:
		return true
	}
	return false
}

func Equal(lhsV, rhsV interface{}) bool {
	if lhsV == rhsV {
		return true
	}

	if IsNum(lhsV) && IsNum(rhsV) {
		if reflect.TypeOf(rhsV).ConvertibleTo(reflect.TypeOf(lhsV)) {
			rhsV = reflect.ValueOf(rhsV).Convert(reflect.TypeOf(lhsV)).Interface()
		}
	}
	return reflect.DeepEqual(lhsV, rhsV)
}

func GetMember(v reflect.Value, name int, stmt ast.Pos) (reflect.Value, error) {

	m, _ := ast.MethodByNameCI(v, name)
	// ошибку не обрабатываем, т.к. ищем поле
	if !m.IsValid() {
		if v.Kind() == reflect.Ptr {
			v = v.Elem()
		}
		if v.Kind() == reflect.Struct {
			var err error
			m, err = ast.FieldByNameCI(v, name)
			if err != nil || !m.IsValid() {
				return envir.NilValue, NewStringError(stmt, "Метод или поле не найдено: "+ast.UniqueNames.Get(name))
			}
		} else if v.Kind() == reflect.Map {
			m = v.MapIndex(reflect.ValueOf(ast.UniqueNames.Get(name)))
			if !m.IsValid() {
				return envir.NilValue, NewStringError(stmt, "Значение по ключу не найдено")
			}
		} else {
			return envir.NilValue, NewStringError(stmt, "У значения нет полей")
		}
	}
	return m, nil
}

func LeftRightBounds(rb, re reflect.Value, vlen int) (ii, ij int, err error) {
	// границы как в python:
	// положительный - имеет максимум до длины (len)
	// отрицательный - считается с конца с минимумом -длина
	// если выходит за макс. границу - возвращаем пустой слайс
	// если выходит за мин. границу - считаем =0

	// правая граница как в python - исключается

	// левая граница включая
	if !rb.IsValid() {
		ii = 0
	} else {
		if rb.Kind() != reflect.Int && rb.Kind() != reflect.Int64 {
			return 0, 0, fmt.Errorf("Индекс должен быть целым числом")
		}
		ii = int(rb.Int())
	}

	switch {
	case ii > 0:
		if ii >= vlen {
			ii = vlen - 1
		}
	case ii < 0:
		ii += vlen
		if ii < 0 {
			ii = 0
		}
	}
	// правая граница не включая
	if !re.IsValid() {
		ij = vlen
	} else {
		if re.Kind() != reflect.Int && re.Kind() != reflect.Int64 {
			return 0, 0, fmt.Errorf("Индекс должен быть целым числом")
		}
		ij = int(re.Int())
	}

	switch {
	case ij > 0:
		if ij > vlen {
			ij = vlen
		}
	case ij < 0:
		ij += vlen
		if ij < 0 {
			ij = 0
		}
	}
	return
}

func SliceAt(v, rb, re, defval reflect.Value) (reflect.Value, error) {

	vlen := v.Len()

	ii, ij, err := LeftRightBounds(rb, re, vlen)
	if err != nil {
		return defval, err
	}

	if ij < ii {
		return defval, fmt.Errorf("Окончание диапазона не может быть раньше его начала")
	}

	return v.Slice(ii, ij), nil
}

func StringToRuneSliceAt(v, rb, re reflect.Value) (fullrune []rune, ii, ij int, err error) {

	r := []rune(v.String())
	vlen := len(r)

	ii, ij, err = LeftRightBounds(rb, re, vlen)
	if err != nil {
		return
	}

	if ij < ii {
		err = fmt.Errorf("Окончание диапазона не может быть раньше его начала")
		return
	}

	return r, ii, ij, nil
}

func StringAt(v, rb, re, defval reflect.Value) (reflect.Value, error) {

	r, ii, ij, err := StringToRuneSliceAt(v, rb, re)
	if err != nil {
		return defval, err
	}

	return reflect.ValueOf(string(r[ii:ij])), nil
}

func EvalBinOp(op int, lhsV, rhsV reflect.Value) (interface{}, error) {
	// log.Println(OperMapR[op])
	if !lhsV.IsValid() || !rhsV.IsValid() {
		if !rhsV.IsValid() && !rhsV.IsValid() {
			// в обоих значениях nil
			return true, nil
		} else {
			// одно из значений nil, а второе нет
			return false, nil
		}
	}

	switch op {

	// TODO: математика множеств и графов

	case ADD:
		if lhsV.Kind() == reflect.String || rhsV.Kind() == reflect.String {
			return ToString(lhsV.Interface()) + ToString(rhsV.Interface()), nil
		}
		if (lhsV.Kind() == reflect.Array || lhsV.Kind() == reflect.Slice) && (rhsV.Kind() != reflect.Array && rhsV.Kind() != reflect.Slice) {
			return reflect.Append(lhsV, rhsV).Interface(), nil
		}
		if (lhsV.Kind() == reflect.Array || lhsV.Kind() == reflect.Slice) && (rhsV.Kind() == reflect.Array || rhsV.Kind() == reflect.Slice) {
			return reflect.AppendSlice(lhsV, rhsV).Interface(), nil
		}
		if lhsV.Kind() == reflect.Float64 || rhsV.Kind() == reflect.Float64 {
			return ToFloat64(lhsV.Interface()) + ToFloat64(rhsV.Interface()), nil
		}
		return ToInt64(lhsV.Interface()) + ToInt64(rhsV.Interface()), nil
	case SUB:
		if lhsV.Kind() == reflect.Float64 || rhsV.Kind() == reflect.Float64 {
			return ToFloat64(lhsV.Interface()) - ToFloat64(rhsV.Interface()), nil
		}
		return ToInt64(lhsV.Interface()) - ToInt64(rhsV.Interface()), nil
	case MUL:
		if lhsV.Kind() == reflect.String && (rhsV.Kind() == reflect.Int || rhsV.Kind() == reflect.Int32 || rhsV.Kind() == reflect.Int64) {
			return strings.Repeat(ToString(lhsV.Interface()), int(ToInt64(rhsV.Interface()))), nil
		}
		if lhsV.Kind() == reflect.Float64 || rhsV.Kind() == reflect.Float64 {
			return ToFloat64(lhsV.Interface()) * ToFloat64(rhsV.Interface()), nil
		}
		return ToInt64(lhsV.Interface()) * ToInt64(rhsV.Interface()), nil
	case QUO:
		return ToFloat64(lhsV.Interface()) / ToFloat64(rhsV.Interface()), nil
	case REM:
		return ToInt64(lhsV.Interface()) % ToInt64(rhsV.Interface()), nil
	case EQL:
		return Equal(lhsV.Interface(), rhsV.Interface()), nil
	case NEQ:
		return Equal(lhsV.Interface(), rhsV.Interface()) == false, nil
	case GTR:
		return ToFloat64(lhsV.Interface()) > ToFloat64(rhsV.Interface()), nil
	case GEQ:
		return ToFloat64(lhsV.Interface()) >= ToFloat64(rhsV.Interface()), nil
	case LSS:
		return ToFloat64(lhsV.Interface()) < ToFloat64(rhsV.Interface()), nil
	case LEQ:
		return ToFloat64(lhsV.Interface()) <= ToFloat64(rhsV.Interface()), nil
	case OR:
		return ToInt64(lhsV.Interface()) | ToInt64(rhsV.Interface()), nil
	case LOR:
		if x := ToBool(lhsV.Interface()); x {
			return x, nil
		} else {
			return ToBool(rhsV.Interface()), nil
		}
	case AND:
		return ToInt64(lhsV.Interface()) & ToInt64(rhsV.Interface()), nil
	case LAND:
		if x := ToBool(lhsV.Interface()); x {
			return ToBool(rhsV.Interface()), nil
		} else {
			return x, nil
		}
	case POW:
		if lhsV.Kind() == reflect.Float64 {
			return math.Pow(ToFloat64(lhsV.Interface()), ToFloat64(rhsV.Interface())), nil
		}
		return int64(math.Pow(ToFloat64(lhsV.Interface()), ToFloat64(rhsV.Interface()))), nil
	case SHR:
		return ToInt64(lhsV.Interface()) >> uint64(ToInt64(rhsV.Interface())), nil
	case SHL:
		return ToInt64(lhsV.Interface()) << uint64(ToInt64(rhsV.Interface())), nil
	default:
		return nil, fmt.Errorf("Неизвестный оператор")
	}
}

// TODO: переделать с учетом здесь новых функций ToInt, ToBool и т.п., убрать рефлексию

func TypeCastConvert(rv reflect.Value, nt reflect.Type, skipCollections bool, defval reflect.Value) (reflect.Value, error) {
	rvkind := rv.Kind()
	if skipCollections && (rvkind == reflect.Array || rvkind == reflect.Slice ||
		rvkind == reflect.Map || rvkind == reflect.Struct || rvkind == reflect.Chan) {
		return rv, nil
	}
	if rvkind == reflect.Interface || rvkind == reflect.Ptr {
		rv = rv.Elem()
		rvkind = rv.Kind()
	}
	// учитываем случай двойной вложенности указателя или интерфейса в указателе
	if rvkind == reflect.Interface || rvkind == reflect.Ptr {
		rv = rv.Elem()
		rvkind = rv.Kind()
	}
	if rvkind == nt.Kind() {
		return rv, nil
	}

	switch rvkind {
	case reflect.Array, reflect.Slice:
		switch nt.Kind() {
		case reflect.String:
			// сериализуем в json
			b, err := json.Marshal(rv.Interface())
			if err != nil {
				return rv, err
			}
			return reflect.ValueOf(string(b)), nil
		default:
			// преобразуем в такой же слайс, но с типизированными значениями, и копируем их с новым типом
			rs := reflect.MakeSlice(reflect.SliceOf(nt), rv.Len(), rv.Cap())
			for i := 0; i < rv.Len(); i++ {
				iv := rv.Index(i)
				// конверсия вложенных массивов и структур не производится
				rsi, err := TypeCastConvert(iv, nt, true, defval)
				if err != nil {
					return rv, err
				}
				sv := rs.Index(i)
				if sv.CanSet() {
					sv.Set(rsi)
				}
				//rs = reflect.Append(rs, rsi)
			}
			return rs, nil
		}
	case reflect.Chan:
		// возвращаем новый канал с типизированными значениями и прежним размером буфера
		return reflect.MakeChan(reflect.ChanOf(reflect.BothDir, nt), rv.Cap()), nil
	case reflect.Map:
		switch nt.Kind() {
		case reflect.String:
			// сериализуем в json
			b, err := json.Marshal(rv.Interface())
			if err != nil {
				return rv, err
			}
			return reflect.ValueOf(string(b)), nil
		case reflect.Struct:
			// для приведения в структурные типы - можно использовать мапу для заполнения полей
			rs := reflect.New(nt) // указатель на новую структуру
			//заполняем экспортируемые неанонимные поля, если их находим в мапе
			for i := 0; i < nt.NumField(); i++ {
				f := nt.Field(i)
				if f.PkgPath == "" && !f.Anonymous {
					setv := reflect.Indirect(rv.MapIndex(reflect.ValueOf(f.Name)))
					if setv.Kind() == reflect.Interface {
						setv = setv.Elem()
					}
					fv := rs.Elem().FieldByName(f.Name)
					if setv.IsValid() && fv.IsValid() && fv.CanSet() {
						if fv.Kind() != setv.Kind() && setv.Type().ConvertibleTo(fv.Type()) {
							setv = setv.Convert(fv.Type())
						}
						fv.Set(setv)
					}
				}
			}
			return rs, nil
		}
	case reflect.String:
		switch nt.Kind() {
		case reflect.Float64:
			if rv.Type().ConvertibleTo(nt) {
				return rv.Convert(nt), nil
			}
			f, err := strconv.ParseFloat(ToString(rv), 64)
			if err == nil {
				return reflect.ValueOf(f), nil
			}
		case reflect.Array, reflect.Slice:
			//парсим json из строки и пытаемся получить массив
			var rm []interface{}
			if err := json.Unmarshal([]byte(ToString(rv)), &rm); err != nil {
				return rv, err
			}
			return reflect.ValueOf(rm), nil
		case reflect.Map:
			//парсим json из строки и пытаемся получить мапу
			var rm map[string]interface{}
			if err := json.Unmarshal([]byte(ToString(rv)), rm); err != nil {
				return rv, err
			}
			return reflect.ValueOf(rm), nil
		case reflect.Struct:
			//парсим json из строки и пытаемся получить указатель на структуру
			rm := reflect.New(nt).Interface()
			if err := json.Unmarshal([]byte(ToString(rv)), rm); err != nil {
				return rv, err
			}
			return reflect.ValueOf(rm), nil
		case reflect.Int64:
			if rv.Type().ConvertibleTo(nt) {
				return rv.Convert(nt), nil
			}
			i, err := strconv.ParseInt(ToString(rv), 10, 64)
			if err == nil {
				return reflect.ValueOf(i), nil
			}
			f, err := strconv.ParseFloat(ToString(rv), 64)
			if err == nil {
				return reflect.ValueOf(int64(f)), nil
			}
		case reflect.Bool:
			s := strings.ToLower(ToString(rv))
			if s == "истина" {
				return reflect.ValueOf(true), nil
			}
			if rv.Type().ConvertibleTo(reflect.TypeOf(1.0)) && rv.Convert(reflect.TypeOf(1.0)).Float() > 0.0 {
				return reflect.ValueOf(true), nil
			}
			b, err := strconv.ParseBool(s)
			if err == nil {
				return reflect.ValueOf(b), nil
			}
			return reflect.ValueOf(false), nil
		default:
			if rv.Type().ConvertibleTo(nt) {
				return rv.Convert(nt), nil
			}
		}
	case reflect.Bool:
		switch nt.Kind() {
		case reflect.String:
			if ToBool(rv) {
				return reflect.ValueOf("Истина"), nil
			} else {
				return reflect.ValueOf("Ложь"), nil
			}
		case reflect.Int64:
			if ToBool(rv) {
				return reflect.ValueOf(int64(1)), nil
			} else {
				return reflect.ValueOf(int64(0)), nil
			}
		case reflect.Float64:
			if ToBool(rv) {
				return reflect.ValueOf(float64(1.0)), nil
			} else {
				return reflect.ValueOf(float64(0.0)), nil
			}
		}
	case reflect.Float32, reflect.Float64,
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		// числа конвертируются стандартно
		switch nt.Kind() {
		case reflect.Bool:
			return reflect.ValueOf(ToBool(rv)), nil
		default:
			if rv.Type().ConvertibleTo(nt) {
				return rv.Convert(nt), nil
			}
		}
	case reflect.Struct:
		if t, ok := rv.Interface().(time.Time); ok {
			// это дата/время - конвертируем в секунды (целые или с плавающей запятой) или в формат RFC3339
			switch nt.Kind() {
			case reflect.String:
				return reflect.ValueOf(t.Format(time.RFC3339)), nil
			case reflect.Int64:
				return reflect.ValueOf(t.Unix()), nil
			case reflect.Float64:
				return reflect.ValueOf(float64(t.UnixNano()) / 1e9), nil
			}
		} else {
			switch nt.Kind() {
			case reflect.Map:
				// структура может быть приведена в мапу
				rs := make(map[string]interface{})
				rtyp := rv.Type()
				for i := 0; i < rtyp.NumField(); i++ {
					f := rtyp.Field(i)
					fv := rv.Field(i)
					if f.PkgPath == "" && !f.Anonymous {
						rs[f.Name] = fv.Interface()
					}
				}
				return reflect.ValueOf(rs), nil
			case reflect.String:
				// сериализуем структуру в json
				b, err := json.Marshal(rv.Interface())
				if err != nil {
					return rv, err
				}
				return reflect.ValueOf(string(b)), nil

			}
		}
	}
	return defval, fmt.Errorf("Приведение типа недопустимо")
}
