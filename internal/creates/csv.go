package creates

import (
	"encoding/csv"
	"reflect"

	"github.com/bipo38/isaac-tui/internal/manipulation"
)

type Writter interface {
	Write() error
}

type Csv[T any] struct {
	Name     string
	Path     string
	Category []T
}

func (el *Csv[T]) Write() error {

	h := manipulation.GetHeaders(el.Category[0])

	fn := manipulation.ParserFileName(el.Name, "csv")

	_, fp := manipulation.RouteParser(el.Path, fn)

	f, err := File(fp)
	if err != nil {
		return err
	}

	w := csv.NewWriter(f)
	if err := w.Write(h); err != nil {
		return err
	}

	for _, v := range el.Category {

		s := format(v, h)

		if err := w.Write(s); err != nil {
			return err
		}
	}

	w.Flush()

	f.Close()

	return nil
}

func format[T any](item T, headers []string) []string {

	var el []string

	e := reflect.ValueOf(item)

	for _, h := range headers {

		el = append(el, e.FieldByName(h).String())
	}

	return el

}
