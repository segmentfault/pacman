package i18n

import (
	"os"
	"testing"

	"github.com/segmentfault/pacman/i18n"
	"github.com/stretchr/testify/assert"
)

func TestNewTranslator(t *testing.T) {
	translator, err := NewTranslator("./testdata/")
	assert.NoError(t, err)
	assert.Equal(t, translator.Tr(i18n.LanguageChinese, "base.success"), "成功")
	assert.Equal(t, translator.Tr(i18n.LanguageEnglish, "base.success"), "success")
	assert.Equal(t, translator.Tr(i18n.LanguageChinese, "base.test"), "{{ bad_value }} original value")
	assert.Equal(t, translator.Tr(i18n.LanguageChinese, "base.not_found"), "base.not_found")
}

func TestAddTranslator(t *testing.T) {
	enUS, err := os.ReadFile("./testdata/en_US.yaml")
	assert.NoError(t, err)

	zhCN, err := os.ReadFile("./testdata/zh_CN.yaml")
	assert.NoError(t, err)

	err = AddTranslator(enUS, "en_US.yaml")
	assert.NoError(t, err)

	err = AddTranslator(zhCN, "zh_CN.yaml")
	assert.NoError(t, err)

	assert.Equal(t, GlobalTrans.Tr(i18n.LanguageChinese, "base.success"), "成功")
	assert.Equal(t, GlobalTrans.Tr(i18n.LanguageEnglish, "base.success"), "success")
}
