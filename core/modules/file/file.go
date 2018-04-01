package file

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

//Upload file
func Upload(w http.ResponseWriter, r *http.Request, field string, dir string) string {
	mf, fh, err := r.FormFile(field)
	if err != nil {
		return ""
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

	// add time prefix if file exist
	if _, err := os.Stat(fPath + "/" + fName); err == nil {
		time := time.Now().Unix()
		pTime := strconv.FormatInt(time, 10)
		fName = pTime + "-" + fName
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

// Delete file
func Delete(path string) {
	var err = os.Remove(path)
	if err != nil {
		return
	}
}
