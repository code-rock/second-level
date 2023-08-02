package wget

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

func writeInFolder(name string, data []byte) {
	folderPath := strings.Split(name, "/")
	if len(folderPath) > 1 {
		dirPath := strings.Join(folderPath[0:len(folderPath)-1], "/")
		err := os.MkdirAll(fmt.Sprintf("./wget/folder/%s", dirPath), 0777)
		if err != nil {
			log.Fatal(err)
		}
	}
	f, err := os.Create(fmt.Sprintf("./wget/folder/%s", name))
	if err != nil {
		log.Fatal(err)
	}

	k, err := f.Write(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(k)
}

func addToQuery(file []byte) []string {
	var links []string
	re := regexp.MustCompile(`(src|href)=['"]./([^'"]*)['"]`)
	match := re.FindAllStringSubmatch(string(file), -1)

	if len(match) > 1 {
		for _, val := range match {
			links = append(links, val[2])
		}
		return links
	} else {
		return []string{}
	}
}

func getFile(url string) ([]byte, error) {
	req, err := http.Get(url)
	if err != nil {
		return []byte{}, err
	}
	defer req.Body.Close()
	b, _ := ioutil.ReadAll(req.Body)
	return b, nil
}

func Download(url string, name string) {
	var query []string
	b, _ := getFile(fmt.Sprintf("%s/%s", url, name))
	writeInFolder(name, b)
	newLinks := addToQuery(b)
	if len(newLinks) == 0 {
		return
	} else {
		query = append(query, newLinks...)
		for i := 0; i < len(query); i++ {
			Download(url, query[i])
		}
	}
}
