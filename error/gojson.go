package error

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func getJson(errJSONPath string) {
	data, error := readAll(errJSONPath)
	if error != nil {

		//todo
		//err := readFileError
		//log.err(err)

		return
	}

	var es = make(map[string]One)
	err := json.Unmarshal(data, &es)
	if err != nil {
		//todo
		//log.error
		return
	}
	for k, v := range es {
		errs[k]=v
	}

	return
}

func readAll(filePth string) ([]byte, error) {
	f, err := os.Open(filePth)
	defer f.Close()
	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(f)
}