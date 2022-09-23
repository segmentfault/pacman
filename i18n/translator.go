package i18n

import "strings"

const (
	LanguageChinese Language = "zh_CN"
	LanguageEnglish Language = "en_US"
	LangEng                  = LanguageEnglish
	LangChn                  = LanguageChinese

	DefaultLanguage = LanguageEnglish
	DefaultLang     = DefaultLanguage
	DefLang         = DefaultLanguage
)

type Language string

// Abbr get language type abbreviation
func (t Language) Abbr() string {
	s := string(t)
	if idx := strings.Index(s, "_"); idx > 0 {
		return s[0:idx]
	}
	return s
}

type Translator interface {
	Tr(Language, string) string
	Dump(Language) ([]byte, error)
}
