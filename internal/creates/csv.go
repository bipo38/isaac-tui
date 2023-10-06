package creates

import (
	"encoding/csv"
	"isaac-scrapper/internal/manipulation"
	"os"
)

func Csv(c interface{}, fp, fn string) (*csv.Writer, *os.File, error) {
	headers := manipulation.GetHeaders(c)

	fn = manipulation.ParserFileName(fn, "csv")

	_, fp = manipulation.RouteParser(fp, fn)

	f, err := File(fp)
	if err != nil {
		return nil, nil, err
	}

	w := csv.NewWriter(f)
	if err := w.Write(headers); err != nil {
		return nil, nil, err
	}

	return w, f, nil

}
