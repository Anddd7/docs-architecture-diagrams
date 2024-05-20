package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"

	gt "github.com/bas24/googletranslatefree"
)

type TranslateCmd struct {
	Filepath string `required:"" short:"f" type:"path" default:"."`
	Language string `required:"" short:"l" enum:"zh,en" default:"zh"`
}

var cache = sync.Map{}

func (cmd *TranslateCmd) Run(globals *Globals) error {
	if cmd.Language != "zh" && cmd.Language != "en" {
		return errors.New("Only support zh and en now.")
	}

	// scan the files or use specific one
	files := getDiagramFiles(cmd.Filepath)

	// convert to Diagram
	diagrams := []Diagram{}
	for _, path := range files {
		ext := filepath.Ext(path)
		// only support excalidraw files now
		if ext != ".excalidraw" {
			continue
		}
		diagrams = append(diagrams, Excalidraw{path, cmd.Language})
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

	return nil
}

func getDiagramFiles(path string) []string {
	files := []string{}
	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
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
	path     string
	language string
}

func (t Excalidraw) translate() {
	targetPath := filepath.Join(filepath.Dir(t.path), "_i18n", t.language, filepath.Base(t.path))
	if _, err := os.Stat(targetPath); !os.IsNotExist(err) {
		fmt.Printf("File exists, skip: %s\n", targetPath)
		return
	}

	fmt.Printf("Translating: %s\n", t.path)

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

	if result, ok := cache.Load(text); ok {
		return result.(string)
	}

	from, to := "zh", "en"
	if t.language == "zh" {
		from, to = "en", "zh"
	}

	result, _ := gt.Translate(text, from, to)
	cache.Store(text, result)

	return result
}
