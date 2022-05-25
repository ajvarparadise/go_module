package meta

import "fmt"

type MetaList map[string]string

func Data() MetaList {
	return MetaList{
		"name":     "Admir Husovic",
		"mail":     "mail.admir.husovic@gmail.com",
		"location": "STOCKHOLM, SWEDEN",
		"role":     "Fullstack-developer",
	}
}

func PrintMeta() {
	data := Data()
	for label, value := range data {
		fmt.Printf("%s: %s\n", label, value)
	}
}
