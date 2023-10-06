package creates

import (
	"os"
	"strings"
)

func Dirs(p string) error {
	exist, err := ExistPath(p)
	if err != nil {
		return err
	}

	if exist {
		return nil
	}

	if err := os.MkdirAll(p, os.ModePerm); err != nil {

		return err
	}

	return nil

}

func File(fp string) (*os.File, error) {

	split := strings.Split(fp, "/")

	r := strings.Join(split[0:len(split)-1], "/")

	e, err := ExistPath(fp)
	if err != nil {
		return nil, err
	}

	if e {
		os.Remove(fp)
	} else {
		Dirs(r)
	}

	f, err := os.Create(fp)
	if err != nil {
		return nil, err
	}

	return f, nil
}

func ExistPath(path string) (bool, error) {
	_, err := os.Stat(path)

	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}

	return false, err
}
