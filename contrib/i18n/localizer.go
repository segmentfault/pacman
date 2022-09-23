package i18n

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/segmentfault/pacman/i18n"
	"github.com/segmentfault/pacman/log"

	goI18n "github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v3"
)

// trans global translation
var trans *Translator

// Translator save the language type with localizer mapping
type Translator struct {
	localizes map[i18n.Language]*goI18n.Localizer
	jsonData  map[i18n.Language]any
}

// Dump the translator into json format
func (tr *Translator) Dump(la i18n.Language) ([]byte, error) {
	return json.Marshal(tr.jsonData[la])
}

// Tr to Translate from specified language and string
// TODO: improve multi-threading performance
func (tr *Translator) Tr(la i18n.Language, key string) string {
	l, ok := tr.localizes[la]
	if !ok {
		l = tr.localizes[i18n.DefaultLanguage]
	}

	translation, err := l.Localize(&goI18n.LocalizeConfig{MessageID: key})
	if err != nil {
		log.Warn(err)
		return key
	}

	return translation
}

// NewTranslator new translator from bundle/resource directory
// TODO: singleton and multi-thread safe initialization
func NewTranslator(bundleDir string) (i18n.Translator, error) {
	stat, err := os.Stat(bundleDir)
	if err != nil {
		return nil, err
	}
	if !stat.IsDir() {
		return nil, fmt.Errorf("%s is not a directory", bundleDir)
	}

	entries, err := os.ReadDir(bundleDir)
	if err != nil {
		return nil, err
	}

	trans = &Translator{
		localizes: make(map[i18n.Language]*goI18n.Localizer),
		jsonData:  make(map[i18n.Language]any),
	}

	// set default language is english
	bundle := goI18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("yaml", yaml.Unmarshal)

	// read the bundle resources file from entries
	for _, file := range entries {
		// ignore directory
		if file.IsDir() {
			continue
		}

		// read the resource from single entry file
		buf, err := os.ReadFile(filepath.Join(bundleDir, file.Name()))
		if err != nil {
			return nil, err
		}

		// the default localizes format is yaml
		bundle.MustParseMessageFileBytes(buf, file.Name())

		// TODO: print the log message
		languageName := strings.TrimSuffix(file.Name(), filepath.Ext(file.Name()))

		trans.localizes[i18n.Language(languageName)] = goI18n.NewLocalizer(bundle, languageName)

		// convert localizes language format into json
		j, err := yamlToJson(buf)
		if err != nil {
			return nil, err
		}

		trans.jsonData[i18n.Language(languageName)] = j
	}

	return trans, nil
}
