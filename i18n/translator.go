package i18n

import "strings"

const (
	LanguageChinese            Language = "zh_CN"
	LanguageChineseTraditional Language = "zh_TW"
	LanguageEnglish            Language = "en_US"
	LanguageGerman             Language = "de_DE"
	LanguageSpanish            Language = "es_ES"
	LanguageFrench             Language = "fr_FR"
	LanguageItalian            Language = "it_IT"
	LanguageJapanese           Language = "ja_JP"
	LanguageKorean             Language = "ko_KR"
	LanguagePortuguese         Language = "pt_PT"
	LanguageRussian            Language = "ru_RU"
	LanguageVietnamese         Language = "vi_VN"

	DefaultLanguage = LanguageEnglish
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
