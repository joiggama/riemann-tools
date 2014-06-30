package utils

import(
  "io/ioutil"
)

func ReadFile(filename string) string {
	file, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	return string(file)
}
