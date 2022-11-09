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

// GlobalTrans global translation
var (
	GlobalTrans = &Translator{
		localizes: make(map[i18n.Language]*goI18n.Localizer),
		jsonData:  make(map[i18n.Language]any),
	}
	Bundle = goI18n.NewBundle(language.English)
)

func init() {
	Bundle.RegisterUnmarshalFunc("yaml", yaml.Unmarshal)
}

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

// AddTranslator add translator to global translations
// translation is content
// language is language file name like en_US.yaml
func AddTranslator(translation []byte, language string) (err error) {
	// the default localizes format is yaml
	_, err = Bundle.ParseMessageFileBytes(translation, language)
	if err != nil {
		return err
	}

	languageName := strings.TrimSuffix(language, filepath.Ext(language))

	GlobalTrans.localizes[i18n.Language(languageName)] = goI18n.NewLocalizer(Bundle, languageName)

	// convert localizes language format into json
	j, err := yamlToJson(translation)
	if err != nil {
		return err
	}
	GlobalTrans.jsonData[i18n.Language(languageName)] = j
	return
}

// NewTranslator new translator from Bundle/resource directory
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

	// read the Bundle resources file from entries
	for _, file := range entries {
		// ignore directory
		if file.IsDir() {
			continue
		}
		// ignore non-YAML file
		if filepath.Ext(file.Name()) != ".yaml" {
			continue
		}

		// read the resource from single entry file
		buf, err := os.ReadFile(filepath.Join(bundleDir, file.Name()))
		if err != nil {
			return nil, err
		}

		// the default localizes format is yaml
		if _, err := Bundle.ParseMessageFileBytes(buf, file.Name()); err != nil {
			return nil, fmt.Errorf("parse language message file [%s] failed: %s", file.Name(), err)
		}

		languageName := strings.TrimSuffix(file.Name(), filepath.Ext(file.Name()))

		GlobalTrans.localizes[i18n.Language(languageName)] = goI18n.NewLocalizer(Bundle, languageName)

		// convert localizes language format into json
		j, err := yamlToJson(buf)
		if err != nil {
			return nil, err
		}

		GlobalTrans.jsonData[i18n.Language(languageName)] = j
	}
	return GlobalTrans, nil
}
