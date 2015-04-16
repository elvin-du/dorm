package orm

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"reflect"
	"strings"
)

var db *sql.DB = nil

func init() {
	log.SetFlags(log.Lshortfile)
	var err error = nil
	db, err = sql.Open("mysql", "root:@(localhost:3306)/test")
	if nil != err {
		log.Fatalln(err)
	}
}

//func ParseStruct(i interface{}) {
//	tp := reflect.TypeOf(i)
//	if tp.Kind() != reflect.Ptr && tp.Kind() != reflect.Struct {
//		log.Panicln("non-ptr or not-struct")
//	}

//	v := reflect.ValueOf(i)

//	log.Println(v.Elem().FieldByName("Name").String())
//	log.Println(tp.Elem().FieldByName("Name"))

//	feild, ok := tp.Elem().FieldByName("Name")
//	if ok {
//		log.Println(feild.Tag.Get("mysql"))
//		log.Println(feild.Tag.Get("json"))
//		log.Println(feild.Name)
//	}
//}

func SelectOne(i interface{}, query string, a ...interface{}) {
	rows, err := db.Query(query, a...)
	if nil != err {
		log.Println(err)
		return
	}

	cols, err := rows.Columns()
	if nil != err {
		log.Println(err)
		return
	}

	is := make([]interface{}, len(cols))
    tp := reflect.TypeOf(i)
	val := reflect.ValueOf(i)
	for i, v := range cols {
		v = strings.ToUpper(v[0:1]) + v[1:]
		is[i] = val.Elem().FieldByName(v).Addr().Interface()
	}

	if rows.Next() {
		err = rows.Scan(is...)
		if nil != err {
			log.Println(err)
			return
		}
	}
}

func InsertSql(i interface{}) (string, error) {
	tp := reflect.TypeOf(i)
	if tp.Kind() != reflect.Ptr && tp.Kind() != reflect.Struct {
		log.Println("non-ptr or not-struct")
		return "", errors.New("non-ptr or not-struct")
	}

	length := tp.Elem().NumField()
	fields := []reflect.StructField{}
	for i := 0; i < length; i++ {
		field := tp.Elem().Field(i)
		fields = append(fields, field)
	}

	val := reflect.ValueOf(i)
	names := []string{}
	values := []string{}
	for _, field := range fields {
		n := field.Tag.Get("mysql")
		names = append(names, n)

		switch val.Elem().FieldByName(field.Name).Type().Kind() {
		case reflect.String:
			v := val.Elem().FieldByName(field.Name).String()
			values = append(values, "'"+v+"'")
		case reflect.Uint, reflect.Uint8:
			v := val.Elem().FieldByName(field.Name).Uint()
			values = append(values, fmt.Sprintf("%d", v))
		case reflect.Int:
			v := val.Elem().FieldByName(field.Name).Int()
			values = append(values, fmt.Sprintf("%d", v))
		}

	}

	insertSQL := "INSERT INTO " + strings.ToLower(tp.Elem().Name()) + "("

	vals := strings.Join(values, ",")
	ns := strings.Join(names, ",")

	insertSQL += ns + ") " + "VALUES(" + vals + ")"
	log.Println(insertSQL)
	return insertSQL, nil
}
