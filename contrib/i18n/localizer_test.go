package i18n

import (
	"testing"

	"github.com/segmentfault/pacman/i18n"
	"github.com/stretchr/testify/assert"
)

func TestNewTranslator(t *testing.T) {
	translator, err := NewTranslator("./testdata/")
	assert.NoError(t, err)
	assert.Equal(t, translator.Tr(i18n.LanguageChinese, "base.success"), "成功")
	assert.Equal(t, translator.Tr(i18n.LanguageEnglish, "base.success"), "success")
}
