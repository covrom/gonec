package bincode

import (
	"encoding/json"
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/covrom/gonec/builtins"
	posit "github.com/covrom/gonec/pos"
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
	if v == nil {
		return "Неопределено"
	}
	if s, ok := v.(string); ok {
		return s
	}
	if b, ok := v.(bool); ok {
		if b {
			return "true" // для совместимости с другими платформами
		} else {
			return "false" // для совместимости с другими платформами
		}
	}
	return fmt.Sprint(v)
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

func GetMember(v reflect.Value, name int, stmt posit.Pos) (reflect.Value, error) {

	// m, _ := ast.MethodByNameCI(v, name)
	// // ошибку не обрабатываем, т.к. ищем поле
	// if !m.IsValid() {
	// 	if v.Kind() == reflect.Ptr {
	// 		v = v.Elem()
	// 	}
	// 	if v.Kind() == reflect.Struct {
	// 		var err error
	// 		m, err = ast.FieldByNameCI(v, name)
	// 		if err != nil || !m.IsValid() {
	// 			return envir.NilValue, NewStringError(stmt, "Метод или поле не найдено: "+env.UniqueNames.Get(name))
	// 		}
	// 	} else if v.Kind() == reflect.Map {
	// 		m = v.MapIndex(reflect.ValueOf(env.UniqueNames.Get(name)))
	// 		if !m.IsValid() {
	// 			return envir.NilValue, NewStringError(stmt, "Значение по ключу не найдено")
	// 		}
	// 	} else {
	// 		return envir.NilValue, NewStringError(stmt, "У значения нет полей")
	// 	}
	// }
	// return m, nil
	panic("TODO")
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

func SliceAt(v, rb, re reflect.Value) (interface{}, error) {

	vlen := v.Len()

	ii, ij, err := LeftRightBounds(rb, re, vlen)
	if err != nil {
		return nil, err
	}

	if ij < ii {
		return nil, fmt.Errorf("Окончание диапазона не может быть раньше его начала")
	}

	return v.Slice(ii, ij).Interface(), nil
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

func StringAt(v, rb, re reflect.Value) (interface{}, error) {

	r, ii, ij, err := StringToRuneSliceAt(v, rb, re)
	if err != nil {
		return nil, err
	}

	return string(r[ii:ij]), nil
}

func EvalBinOp(op core.VMOperation, lhsV, rhsV reflect.Value) (interface{}, error) {
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
	lk := lhsV.Kind()
	rk := rhsV.Kind()
	lvi := lhsV.Interface()
	rvi := rhsV.Interface()

	switch op {

	// TODO: математика множеств и графов

	case core.ADD:
		if lk == reflect.String || rk == reflect.String {
			return ToString(lvi) + ToString(rvi), nil
		}
		if (lk == reflect.Array || lk == reflect.Slice) && (rk != reflect.Array && rk != reflect.Slice) {
			return reflect.Append(lhsV, rhsV).Interface(), nil
		}
		if (lk == reflect.Array || lk == reflect.Slice) && (rk == reflect.Array || rk == reflect.Slice) {
			return reflect.AppendSlice(lhsV, rhsV).Interface(), nil
		}
		if lk == reflect.Float64 || rk == reflect.Float64 {
			return ToFloat64(lvi) + ToFloat64(rvi), nil
		}
		if lk == reflect.Struct && rk == reflect.Int64 {
			// проверяем на дату + число
			if lhsV.Type().AssignableTo(core.ReflectVMTime) {
				if rhsV.Type().AssignableTo(reflect.TypeOf(time.Duration(0))) {
					// это Duration
					return time.Time(lvi.(core.VMTime)).Add(time.Duration(rhsV.Int())), nil
				} else {
					// это было число в секундах
					return time.Time(lvi.(core.VMTime)).Add(time.Duration(1e9 * rhsV.Int())), nil
				}
			}
		}
		return ToInt64(lvi) + ToInt64(rvi), nil
	case core.SUB:
		if lk == reflect.Float64 || rk == reflect.Float64 {
			return ToFloat64(lvi) - ToFloat64(rvi), nil
		}
		if lk == reflect.Struct && rk == reflect.Int64 {
			// проверяем на дату + число
			if lhsV.Type().AssignableTo(core.ReflectVMTime) {
				if rhsV.Type().AssignableTo(reflect.TypeOf(time.Duration(0))) {
					// это Duration
					return time.Time(lvi.(core.VMTime)).Add(time.Duration(-rhsV.Int())), nil
				} else {
					// это было число в секундах
					return time.Time(lvi.(core.VMTime)).Add(time.Duration(-1e9 * rhsV.Int())), nil
				}
			}
		}
		if lk == reflect.Struct && rk == reflect.Struct {
			// проверяем на дата - дата
			if lhsV.Type().AssignableTo(core.ReflectVMTime) && rhsV.Type().AssignableTo(core.ReflectVMTime) {
				return time.Time(lvi.(core.VMTime)).Sub(time.Time(rvi.(core.VMTime))), nil
			}
		}
		return ToInt64(lvi) - ToInt64(rvi), nil
	case core.MUL:
		if lk == reflect.String && (rk == reflect.Int || rk == reflect.Int32 || rk == reflect.Int64) {
			return strings.Repeat(ToString(lvi), int(ToInt64(rvi))), nil
		}
		if lk == reflect.Float64 || rk == reflect.Float64 {
			return ToFloat64(lvi) * ToFloat64(rvi), nil
		}
		return ToInt64(lvi) * ToInt64(rvi), nil
	case core.QUO:
		return ToFloat64(lvi) / ToFloat64(rvi), nil
	case core.REM:
		return ToInt64(lvi) % ToInt64(rvi), nil
	case core.EQL:
		return Equal(lvi, rvi), nil
	case core.NEQ:
		return Equal(lvi, rvi) == false, nil
	case core.GTR:
		return ToFloat64(lvi) > ToFloat64(rvi), nil
	case core.GEQ:
		return ToFloat64(lvi) >= ToFloat64(rvi), nil
	case core.LSS:
		return ToFloat64(lvi) < ToFloat64(rvi), nil
	case core.LEQ:
		return ToFloat64(lvi) <= ToFloat64(rvi), nil
	case core.OR:
		return ToInt64(lvi) | ToInt64(rvi), nil
	case core.LOR:
		if x := ToBool(lvi); x {
			return x, nil
		} else {
			return ToBool(rvi), nil
		}
	case core.AND:
		return ToInt64(lvi) & ToInt64(rvi), nil
	case core.LAND:
		if x := ToBool(lvi); x {
			return ToBool(rvi), nil
		} else {
			return x, nil
		}
	case core.POW:
		if lk == reflect.Float64 {
			return math.Pow(ToFloat64(lvi), ToFloat64(rvi)), nil
		}
		return int64(math.Pow(ToFloat64(lvi), ToFloat64(rvi))), nil
	case core.SHR:
		return ToInt64(lvi) >> uint64(ToInt64(rvi)), nil
	case core.SHL:
		return ToInt64(lvi) << uint64(ToInt64(rvi)), nil
	default:
		return nil, fmt.Errorf("Неизвестный оператор")
	}
}

func TypeCastConvert(v interface{}, nt reflect.Type, skipCollections bool) (interface{}, error) {
	rv := reflect.ValueOf(v)
	rvkind := rv.Kind()

	if skipCollections && (rvkind == reflect.Array || rvkind == reflect.Slice ||
		rvkind == reflect.Map || rvkind == reflect.Struct || rvkind == reflect.Chan) {
		return v, nil
	}
	if rvkind == reflect.Interface || rvkind == reflect.Ptr {
		rv = rv.Elem()
		rvkind = rv.Kind()
		v = rv.Interface()
	}
	// учитываем случай двойной вложенности указателя или интерфейса в указателе
	if rvkind == reflect.Interface || rvkind == reflect.Ptr {
		rv = rv.Elem()
		rvkind = rv.Kind()
		v = rv.Interface()
	}
	if rvkind == nt.Kind() {
		return v, nil
	}

	switch rvkind {
	case reflect.Array, reflect.Slice:
		switch nt.Kind() {
		case reflect.String:
			// сериализуем в json
			b, err := json.Marshal(v)
			if err != nil {
				return nil, err
			}
			return string(b), nil
		default:
			// преобразуем в такой же слайс, но с типизированными значениями, и копируем их с новым типом
			rs := reflect.MakeSlice(reflect.SliceOf(nt), rv.Len(), rv.Cap())
			for i := 0; i < rv.Len(); i++ {
				iv := rv.Index(i).Interface()
				// конверсия вложенных массивов и структур не производится
				rsi, err := TypeCastConvert(iv, nt, true)
				if err != nil {
					return nil, err
				}
				sv := rs.Index(i)
				if sv.CanSet() {
					sv.Set(reflect.ValueOf(rsi))
				}
				//rs = reflect.Append(rs, rsi)
			}
			return rs.Interface(), nil
		}
	case reflect.Chan:
		// возвращаем новый канал с типизированными значениями и прежним размером буфера
		return reflect.MakeChan(reflect.ChanOf(reflect.BothDir, nt), rv.Cap()).Interface(), nil
	case reflect.Map:
		switch nt.Kind() {
		case reflect.String:
			// сериализуем в json
			b, err := json.Marshal(v)
			if err != nil {
				return nil, err
			}
			return string(b), nil
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
						if fv.Kind() != setv.Kind() {
							if setv.Type().ConvertibleTo(fv.Type()) {
								setv = setv.Convert(fv.Type())
							} else {
								return nil, fmt.Errorf("Поле структуры имеет другой тип")
							}
						}
						fv.Set(setv)
					}
				}
			}
			return rs.Interface(), nil
		}
	case reflect.String:
		if nt.AssignableTo(core.ReflectVMTime) {
			tt, err := time.Parse(time.RFC3339, rv.String())
			if err == nil {
				return core.VMTime(tt), nil
			} else {
				panic(err)
			}
		}
		switch nt.Kind() {
		case reflect.Float64:
			if rv.Type().ConvertibleTo(nt) {
				return rv.Convert(nt).Interface(), nil
			}
			f, err := strconv.ParseFloat(ToString(v), 64)
			if err == nil {
				return f, nil
			}
		case reflect.Array, reflect.Slice:
			//парсим json из строки и пытаемся получить массив
			var rm core.VMSlice
			if err := json.Unmarshal([]byte(ToString(v)), &rm); err != nil {
				return nil, err
			}
			return rm, nil
		case reflect.Map:
			//парсим json из строки и пытаемся получить мапу
			var rm core.VMStringMap
			if err := json.Unmarshal([]byte(ToString(v)), rm); err != nil {
				return nil, err
			}
			return rm, nil
		case reflect.Struct:
			//парсим json из строки и пытаемся получить указатель на структуру
			rm := reflect.New(nt).Interface()
			if err := json.Unmarshal([]byte(ToString(v)), rm); err != nil {
				return nil, err
			}
			return rm, nil
		case reflect.Int64:
			if rv.Type().ConvertibleTo(nt) {
				return rv.Convert(nt).Interface(), nil
			}
			i, err := strconv.ParseInt(ToString(v), 10, 64)
			if err == nil {
				return i, nil
			}
			f, err := strconv.ParseFloat(ToString(v), 64)
			if err == nil {
				return int64(f), nil
			}
		case reflect.Bool:
			s := strings.ToLower(ToString(v))
			if s == "истина" || s == "true" {
				return true, nil
			}
			if s == "ложь" || s == "false" {
				return false, nil
			}
			if rv.Type().ConvertibleTo(reflect.TypeOf(1.0)) && rv.Convert(reflect.TypeOf(1.0)).Float() > 0.0 {
				return true, nil
			}
			b, err := strconv.ParseBool(s)
			if err == nil {
				return b, nil
			}
			return false, nil
		default:
			if rv.Type().ConvertibleTo(nt) {
				return rv.Convert(nt).Interface(), nil
			}
		}
	case reflect.Bool:
		switch nt.Kind() {
		case reflect.String:
			if ToBool(v) {
				return "true", nil // для совместимости с другими платформами
			} else {
				return "false", nil // для совместимости с другими платформами
			}
		case reflect.Int64:
			if ToBool(v) {
				return int64(1), nil
			} else {
				return int64(0), nil
			}
		case reflect.Float64:
			if ToBool(v) {
				return float64(1.0), nil
			} else {
				return float64(0.0), nil
			}
		}
	case reflect.Float32, reflect.Float64,
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		// приведение к дате
		if nt.AssignableTo(core.ReflectVMTime) {
			switch rvkind {
			case reflect.Float32, reflect.Float64:
				rti := int64(rv.Float())
				rtins := int64((rv.Float() - float64(rti)) * 1e9)
				return core.VMTime(time.Unix(rti, rtins)), nil
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
				reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				rti := rv.Int()
				return core.VMTime(time.Unix(rti, 0)), nil
			}
		}
		// числа конвертируются стандартно
		switch nt.Kind() {
		case reflect.Bool:
			return ToBool(v), nil
		default:
			if rv.Type().ConvertibleTo(nt) {
				return rv.Convert(nt).Interface(), nil
			}
		}
	case reflect.Struct:
		t, ok := v.(time.Time)
		if !ok {
			var t2 core.VMTime
			t2, ok = v.(core.VMTime)
			if ok {
				t = time.Time(t2)
			}
		}
		if ok {
			if nt.AssignableTo(core.ReflectVMTime) {
				return core.VMTime(t), nil
			}
			// это дата/время - конвертируем в секунды (целые или с плавающей запятой) или в формат RFC3339
			switch nt.Kind() {
			case reflect.String:
				return t.Format(time.RFC3339), nil
			case reflect.Int64:
				return t.Unix(), nil // до секунд
			case reflect.Float64:
				return float64(t.UnixNano()) / 1e9, nil // после запятой - наносекунды
			}
		} else {
			switch nt.Kind() {
			case reflect.Map:
				// структура может быть приведена в мапу
				rs := make(core.VMStringMap)
				rtyp := rv.Type()
				for i := 0; i < rtyp.NumField(); i++ {
					f := rtyp.Field(i)
					fv := rv.Field(i)
					if f.PkgPath == "" && !f.Anonymous {
						rs[f.Name] = core.ReflectToVMValue(fv)
					}
				}
				return rs, nil
			case reflect.String:
				// сериализуем структуру в json
				b, err := json.Marshal(v)
				if err != nil {
					return nil, err
				}
				return string(b), nil

			}
		}
	}
	return nil, fmt.Errorf("Приведение типа недопустимо")
}
