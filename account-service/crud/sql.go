package crud

import (
	"fmt"
	"io"
	"log"
	"os"
	"reflect"
)

type QueriesStruct struct {
	SELECT_USER_BY_EMAIL string `filename:"SELECT_USER_BY_EMAIL"`
	INSERT_USER          string `filename:"INSERT_USER"`
	SELECT_USER_BY_CREDS string `filename:"SELECT_USER_BY_CREDS"`
	DELETE_USER_BY_ID    string `filename:"DELETE_USER_BY_ID"`
	INSERT_TELEGRAM_SESSION    string `filename:"INSERT_TELEGRAM_SESSION"`
	DELETE_TELEGRAM_SESSION_BY_ID string `filename:"DELETE_TELEGRAM_SESSION_BY_ID"`
}

var SQL = QueriesStruct{}

func Init() {
	if err := readQueries(); err != nil {
		panic(err)
	}
	log.Println(SQL)
}

func readQueries() error {
	var queriesPath = fmt.Sprintf(
		"crud%vSQL%v", string(os.PathSeparator), string(os.PathSeparator),
	)

	SQL = QueriesStruct{}
	st := reflect.TypeOf(SQL)

	for i := 0; i < st.NumField(); i++ {
		field := st.Field(i)

		if filename, ok := field.Tag.Lookup("filename"); ok {
			val := reflect.ValueOf(&SQL).Elem().Field(i)

			file, err := os.Open(fmt.Sprintf("%s%s.sql", queriesPath, filename))
			if err != nil {
				return err
			}
			q, err := io.ReadAll(file)
			if err != nil {
				return err
			}

			if string(q) == "" {
				return fmt.Errorf(fmt.Sprintf("queries field %v is empty", field.Name))
			}
			val.SetString(string(q))
		} else {
			return fmt.Errorf(fmt.Sprintf("queries field %v has missing or wrong tag", field.Name))
		}
	}

	return nil
}
