package main

import (
	"flag"
	"fmt"
	"wget/wget"
)

func main() {
	url := flag.String("url", "https://code-rock.github.io/invasiya-view/", "Set sites url.")
	flag.Parse()
	fmt.Println(*url)
	wget.Download(*url, "index.html")
}
