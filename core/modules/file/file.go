package file

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

//Upload file
func Upload(w http.ResponseWriter, r *http.Request, field string, dir string) string {
	mf, fh, err := r.FormFile(field)
	if err != nil {
		panic(err)
	}
	defer mf.Close()

	fName := fh.Filename

	// create new file
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	t := time.Now()
	tf := t.Format("2006-01-02")

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.Mkdir(dir, os.ModePerm)
	}

	fPath := dir + "/" + tf
	if _, err := os.Stat(fPath); os.IsNotExist(err) {
		os.Mkdir(fPath, os.ModePerm)
	}

	path := filepath.Join(wd, dir, tf, fName)
	nf, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer nf.Close()
	// copy
	mf.Seek(0, 0)
	io.Copy(nf, mf)

	return fPath + "/" + fName
}
