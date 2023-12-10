package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"

	gt "github.com/bas24/googletranslatefree"
)

var _filepath string
var _language string
var cache = make(map[string]string)

func main() {
	// get the args
	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Println("Usage: translate <filepath> <language>")
		return
	}
	_filepath = args[0]
	_language = args[1]

	// scan the files or use specific one
	filepaths := getFilepaths()

	// convert to Diagram
	diagrams := []Diagram{}
	for _, path := range filepaths {
		ext := filepath.Ext(path)
		// only support excalidraw files now
		if ext != ".excalidraw" {
			continue
		}
		diagrams = append(diagrams, Excalidraw{path})
	}

	// parallel translate
	wg := new(sync.WaitGroup)
	for _, dia := range diagrams {
		wg.Add(1)
		go func(d Diagram) {
			defer func() {
				wg.Done()
			}()
			d.translate()
		}(dia)
	}
	wg.Wait()
}

func getFilepaths() []string {
	files := []string{}
	filepath.Walk(_filepath, func(path string, info os.FileInfo, err error) error {
		if path == "." || info.Name()[0] == '.' || info.Name()[0] == '_' ||
			strings.Contains(path, "_i18n") {
			// fmt.Println("Skip dir/file:", path)
			return nil
		}

		files = append(files, path)

		return nil
	})
	return files
}

type Diagram interface {
	translate()
}

type Excalidraw struct {
	path string
}

func (t Excalidraw) translate() {
	targetPath := filepath.Join(filepath.Dir(t.path), "_i18n", _language, filepath.Base(t.path))
	if _, err := os.Stat(targetPath); !os.IsNotExist(err) {
		// fmt.Printf("File exists, skip: %s\n", targetPath)
		return
	}

	file, _ := os.Open(t.path)
	defer file.Close()

	body := make(map[string]interface{})
	_ = json.NewDecoder(file).Decode(&body)

	// translate the elements[*].text, and elements[*].originalText
	body["elements"] = t.translateElements(body["elements"].([]interface{}))

	// create folder if not exists
	if _, err := os.Stat(filepath.Dir(targetPath)); os.IsNotExist(err) {
		_ = os.MkdirAll(filepath.Dir(targetPath), 0755)
	}

	file, _ = os.Create(targetPath)
	defer file.Close()

	_ = json.NewEncoder(file).Encode(body)
}

func (t Excalidraw) translateElements(elements []interface{}) []interface{} {
	for _, e := range elements {
		element := e.(map[string]interface{})
		if element["type"] == "text" {
			element["text"] = t.translateText(element["text"].(string))
			element["originalText"] = t.translateText(element["originalText"].(string))
		}
	}
	return elements
}

func (t Excalidraw) translateText(text string) string {
	// ignore text if empty or spaces
	if len(text) == 0 || len(strings.TrimSpace(text)) == 0 {
		return text
	}

	// ignore if only contains +,-,*,/
	if len(text) == 1 && strings.ContainsAny(text, "+-*/") {
		return text
	}

	if result, ok := cache[text]; ok {
		return result
	}

	from, to := "zh", "en"
	if _language == "zh" {
		from, to = "en", "zh"
	}

	result, _ := gt.Translate(text, from, to)
	cache[text] = result

	return result
}

func confirmOverwrite(path string) bool {
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		fmt.Println("File exists:", path)
		fmt.Print("Overwrite? (y/n): ")
		var response string
		fmt.Scanln(&response)
		if response != "y" {
			fmt.Println("Skip file:", path)
			return false
		}
	}
	return true
}
