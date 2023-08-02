package wget

import (
	"os"
	"testing"
)

func TestDownloadSite(t *testing.T) {
	dir := "./folder"
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		t.Errorf(dir, "does not exist")
	}

	url := "https://code-rock.github.io/invasiya-view/"
	htmlFile := "./folder/index.html"
	Download(url, "index.html")
	if _, err := os.Stat(htmlFile); err != nil {
		t.Errorf("File does not exist")
	}
}
