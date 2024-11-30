package views

import (
	"log"
	"os"
	"path/filepath"
	"regexp"

	"github.com/gin-contrib/multitemplate"
)

const viewsDir = "internal/views"
const layoutsDir = "internal/views/layouts"
const viewExt = ".html"

// それぞれのViewでは、
// {{ template "xxx_start" . }}...{{ template "xxx_end" . }}
// でコンテンツを囲むことでレイアウトを適用できる。
//
// nameマッピング: "home/index" -> "internal/views/home/index.html"
func CreateViews() multitemplate.Renderer {
	renderer := multitemplate.NewRenderer()

	layouts := readLayouts()
	views := readViews()

	for _, view := range views {
		keyRegexp := regexp.MustCompile(filepath.Join(viewsDir, "(.*)"+viewExt))
		key := keyRegexp.FindStringSubmatch(view)[1]

		files := append([]string{view}, layouts...)

		renderer.AddFromFiles(key, files...)
	}

	return renderer
}

func readLayouts() []string {
	return readFilePaths(layoutsDir, "")
}

func readViews() []string {
	return readFilePaths(viewsDir, layoutsDir)
}

func readFilePaths(dir string, skipDir string) []string {
	FilePaths := []string{}

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		if info.IsDir() && path == skipDir {
			return filepath.SkipDir
		}
		if !info.IsDir() && filepath.Ext(path) == viewExt {
			FilePaths = append(FilePaths, path)
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	return FilePaths
}
