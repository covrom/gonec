package core

import (
	"bytes"
	"reflect"
	"sync"
	"time"

	"github.com/boltdb/bolt"
	"github.com/covrom/gonec/names"
)

// VMBoltDB - группа ожидания исполнения горутин
type VMBoltDB struct {
	sync.Mutex
	db *bolt.DB
	tx *bolt.Tx
}

var ReflectVMBoltDB = reflect.TypeOf(VMBoltDB{})

func (x *VMBoltDB) vmval() {}

func (x *VMBoltDB) Interface() interface{} {
	return x
}

func (x *VMBoltDB) String() string {
	return "Файловая база данных BoltDB"
}

func (x *VMBoltDB) Open(filename string) (err error) {
	x.Lock()
	defer x.Unlock()
	x.db, err = bolt.Open(filename, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}
	return nil
}

func (x *VMBoltDB) Close() {
	x.Lock()
	defer x.Unlock()
	if x.db != nil {
		x.db.Close()
		x.db = nil
	}
}

func (x *VMBoltDB) Begin(writable bool) (err error) {
	x.Lock()
	defer x.Unlock()
	if x.tx != nil {
		return VMErrorTransactionIsOpened
	}
	x.tx, err = x.db.Begin(writable)
	if err != nil {
		x.tx = nil
	}
	return
}

func (x *VMBoltDB) Commit() error {
	x.Lock()
	defer x.Unlock()
	if x.tx == nil {
		return VMErrorTransactionNotOpened
	}
	err := x.tx.Commit()
	x.tx = nil
	return err
}

func (x *VMBoltDB) Rollback() error {
	x.Lock()
	defer x.Unlock()
	if x.tx == nil {
		return VMErrorTransactionNotOpened
	}
	x.tx.Rollback()
	x.tx = nil
	return nil
}

func (x *VMBoltDB) CreateTableIfNotExists(name string) (*VMBoltTable, error) {
	if x.tx == nil {
		return nil, VMErrorTransactionNotOpened
	}
	b, err := x.tx.CreateBucketIfNotExists([]byte(name))
	t := &VMBoltTable{name: name, b: b}
	return t, err
}

func (x *VMBoltDB) DeleteTable(name string) error {
	if x.tx == nil {
		return VMErrorTransactionNotOpened
	}
	return x.tx.DeleteBucket([]byte(name))
}

func (x *VMBoltDB) BackupDBToFile(name string) error {
	if x.tx == nil {
		return VMErrorTransactionNotOpened
	}
	return x.tx.CopyFile(name, 0644)
}

func (x *VMBoltDB) MethodMember(name int) (VMFunc, bool) {

	// только эти методы будут доступны из кода на языке Гонец!
	switch names.UniqueNames.GetLowerCase(name) {
	case "открыть":
		return VMFuncMustParams(1, x.Открыть), true
	case "закрыть":
		return VMFuncMustParams(0, x.Закрыть), true
	case "начатьтранзакцию":
		return VMFuncMustParams(1, x.НачатьТранзакцию), true
	case "зафиксироватьтранзакцию":
		return VMFuncMustParams(0, x.ЗафиксироватьТранзакцию), true
	case "отменитьтранзакцию":
		return VMFuncMustParams(0, x.ОтменитьТранзакцию), true
	case "таблица":
		return VMFuncMustParams(1, x.Таблица), true
	case "удалитьтаблицу":
		return VMFuncMustParams(1, x.УдалитьТаблицу), true
	case "полныйбэкап":
		return VMFuncMustParams(1, x.ПолныйБэкап), true
	}
	return nil, false
}

func (x *VMBoltDB) Открыть(args VMSlice, rets *VMSlice, envout *(*Env)) error {
	v, ok := args[0].(VMString)
	if !ok {
		return VMErrorNeedString
	}
	return x.Open(string(v))
}

func (x *VMBoltDB) Закрыть(args VMSlice, rets *VMSlice, envout *(*Env)) error {
	x.Close()
	return nil
}

func (x *VMBoltDB) НачатьТранзакцию(args VMSlice, rets *VMSlice, envout *(*Env)) error {
	v, ok := args[0].(VMBool)
	if !ok {
		return VMErrorNeedBool
	}
	return x.Begin(bool(v))
}

func (x *VMBoltDB) ЗафиксироватьТранзакцию(args VMSlice, rets *VMSlice, envout *(*Env)) error {
	return x.Commit()
}

func (x *VMBoltDB) ОтменитьТранзакцию(args VMSlice, rets *VMSlice, envout *(*Env)) error {
	return x.Rollback()
}

func (x *VMBoltDB) Таблица(args VMSlice, rets *VMSlice, envout *(*Env)) error {
	v, ok := args[0].(VMString)
	if !ok {
		return VMErrorNeedString
	}
	t, err := x.CreateTableIfNotExists(string(v))
	if err != nil {
		return err
	}
	rets.Append(t)
	return nil
}

func (x *VMBoltDB) УдалитьТаблицу(args VMSlice, rets *VMSlice, envout *(*Env)) error {
	v, ok := args[0].(VMString)
	if !ok {
		return VMErrorNeedString
	}
	return x.DeleteTable(string(v))
}

func (x *VMBoltDB) ПолныйБэкап(args VMSlice, rets *VMSlice, envout *(*Env)) error {
	v, ok := args[0].(VMString)
	if !ok {
		return VMErrorNeedString
	}
	return x.BackupDBToFile(string(v))
}

// VMBoltTable реализует функционал Bucket для BoltDB
type VMBoltTable struct {
	name string
	b    *bolt.Bucket
}

func (x *VMBoltTable) vmval() {}

func (x *VMBoltTable) Interface() interface{} {
	return x
}

func (x *VMBoltTable) String() string {
	return "Таблица '" + x.name + "' файловой базы данных BoltDB"
}

func (x *VMBoltTable) Set(k string, v VMBinaryTyper) error {
	i := []byte{byte(v.BinaryType())}
	ii, err := v.MarshalBinary()
	if err != nil {
		return err
	}
	return x.b.Put([]byte(k), append(i, ii...))
}

func parseBoltValue(sl []byte) (VMValuer, error) {
	if len(sl) < 1 {
		return nil, VMErrorWrongDBValue
	}
	tt := sl[0]
	var bb []byte
	if len(sl) > 1 {
		bb = sl[1:]
	}
	vv, err := VMBinaryType(tt).ParseBinary(bb)
	if err != nil {
		return VMNil, err
	}
	return vv, nil
}

func (x *VMBoltTable) Get(k string) (VMValuer, bool, error) {
	sl := x.b.Get([]byte(k))
	if sl == nil {
		return VMNil, false, nil
	}
	vv, err := parseBoltValue(sl)
	return vv, true, err
}

func (x *VMBoltTable) Delete(k string) error {
	return x.b.Delete([]byte(k))
}

func (x *VMBoltTable) NextId() (VMInt, error) {
	id, err := x.b.NextSequence()
	return VMInt(id), err
}

func (x *VMBoltTable) GetPrefix(pref string) (VMStringMap, error) {
	c := x.b.Cursor()
	vsm := make(VMStringMap)
	for k, v := c.Seek([]byte(pref)); k != nil && bytes.HasPrefix(k, []byte(pref)); k, v = c.Next() {
		vx, err := parseBoltValue(v)
		if err != nil {
			return vsm, err
		}
		vsm[string(k)] = vx
	}
	return vsm, nil
}

func (x *VMBoltTable) GetRange(kmin, kmax string) (VMStringMap, error) {
	c := x.b.Cursor()
	vsm := make(VMStringMap)
	for k, v := c.Seek([]byte(kmin)); k != nil && bytes.Compare(k, []byte(kmax)) <= 0; k, v = c.Next() {
		vx, err := parseBoltValue(v)
		if err != nil {
			return vsm, err
		}
		vsm[string(k)] = vx
	}
	return vsm, nil
}

func (x *VMBoltTable) GetAll() (VMStringMap, error) {
	c := x.b.Cursor()
	vsm := make(VMStringMap)
	for k, v := c.First(); k != nil; k, v = c.Next() {
		vx, err := parseBoltValue(v)
		if err != nil {
			return vsm, err
		}
		vsm[string(k)] = vx
	}
	return vsm, nil
}

func (x *VMBoltTable) SetByMap(m VMStringMap) error {

	mm := make(map[string]VMBinaryTyper)
	for ks, vs := range m {
		v, ok := vs.(VMBinaryTyper)
		if !ok {
			return VMErrorNeedBinaryTyper
		}
		mm[ks] = v
	}

	for ks, vs := range mm {

		i := []byte{byte(vs.BinaryType())}
		ii, err := vs.MarshalBinary()
		if err != nil {
			return err
		}
		err = x.b.Put([]byte(ks), append(i, ii...))
		if err != nil {
			return err
		}
	}

	return nil
}

func (x *VMBoltTable) MethodMember(name int) (VMFunc, bool) {

	// только эти методы будут доступны из кода на языке Гонец!
	switch names.UniqueNames.GetLowerCase(name) {
	case "получить":
		return VMFuncMustParams(1, x.Получить), true
	case "установить":
		return VMFuncMustParams(2, x.Установить), true
	case "удалить":
		return VMFuncMustParams(1, x.Удалить), true
	case "следующийидентификатор":
		return VMFuncMustParams(0, x.СледующийИдентификатор), true
	case "получитьдиапазон":
		return VMFuncMustParams(2, x.ПолучитьДиапазон), true
	case "получитьпрефикс":
		return VMFuncMustParams(1, x.ПолучитьПрефикс), true
	case "получитьвсе":
		return VMFuncMustParams(0, x.ПолучитьВсе), true
	case "установитьструктуру":
		return VMFuncMustParams(1, x.УстановитьСтруктуру), true
	}
	return nil, false
}

func (x *VMBoltTable) Получить(args VMSlice, rets *VMSlice, envout *(*Env)) error {
	v, ok := args[0].(VMString)
	if !ok {
		return VMErrorNeedString
	}
	rv, ok, err := x.Get(string(v))
	if err != nil {
		return err
	}
	rets.Append(rv)
	rets.Append(VMBool(ok))
	return nil
}

func (x *VMBoltTable) Установить(args VMSlice, rets *VMSlice, envout *(*Env)) error {
	vk, ok := args[0].(VMString)
	if !ok {
		return VMErrorNeedString
	}
	vv, ok := args[1].(VMBinaryTyper)
	if !ok {
		return VMErrorNeedBinaryTyper
	}
	return x.Set(string(vk), vv)
}

func (x *VMBoltTable) Удалить(args VMSlice, rets *VMSlice, envout *(*Env)) error {
	v, ok := args[0].(VMString)
	if !ok {
		return VMErrorNeedString
	}
	return x.Delete(string(v))
}

func (x *VMBoltTable) СледующийИдентификатор(args VMSlice, rets *VMSlice, envout *(*Env)) error {
	v, err := x.NextId()
	rets.Append(v)
	return err
}

func (x *VMBoltTable) ПолучитьДиапазон(args VMSlice, rets *VMSlice, envout *(*Env)) error {
	vmin, ok := args[0].(VMString)
	if !ok {
		return VMErrorNeedString
	}
	vmax, ok := args[1].(VMString)
	if !ok {
		return VMErrorNeedString
	}
	vsm, err := x.GetRange(string(vmin), string(vmax))
	if err != nil {
		return err
	}
	rets.Append(vsm)
	return nil
}

func (x *VMBoltTable) ПолучитьПрефикс(args VMSlice, rets *VMSlice, envout *(*Env)) error {
	pref, ok := args[0].(VMString)
	if !ok {
		return VMErrorNeedString
	}
	vsm, err := x.GetPrefix(string(pref))
	if err != nil {
		return err
	}
	rets.Append(vsm)
	return nil
}

func (x *VMBoltTable) ПолучитьВсе(args VMSlice, rets *VMSlice, envout *(*Env)) error {
	vsm, err := x.GetAll()
	if err != nil {
		return err
	}
	rets.Append(vsm)
	return nil
}

func (x *VMBoltTable) УстановитьСтруктуру(args VMSlice, rets *VMSlice, envout *(*Env)) error {
	v, ok := args[0].(VMStringMap)
	if !ok {
		return VMErrorNeedMap
	}
	return x.SetByMap(v)
}
