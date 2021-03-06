package bundle

import (
	"testing"

	"github.com/mash/go-i18n/i18n/language"
	"github.com/mash/go-i18n/i18n/translation"
)

func TestMustLoadTranslationFile(t *testing.T) {
	t.Skipf("not implemented")
}

func TestLoadTranslationFile(t *testing.T) {
	t.Skipf("not implemented")
}

func TestParseTranslationFileBytes(t *testing.T) {
	t.Skipf("not implemented")
}

func TestAddTranslation(t *testing.T) {
	t.Skipf("not implemented")
}

func TestMustTfunc(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expected MustTfunc to panic")
		}
	}()
	New().MustTfunc("invalid")
}

func TestTfunc(t *testing.T) {
	b := New()
	translationID := "translation_id"
	englishTranslation := "en-US(translation_id)"
	englishLanguage := language.MustParse("en-US")[0]
	frenchLanguage := language.MustParse("fr-FR")[0]
	spanishLanguage := language.MustParse("es")[0]
	b.AddTranslation(englishLanguage, testNewTranslation(t, map[string]interface{}{
		"id":          translationID,
		"translation": englishTranslation,
	}))
	frenchTranslation := "fr-FR(translation_id)"
	b.AddTranslation(frenchLanguage, testNewTranslation(t, map[string]interface{}{
		"id":          translationID,
		"translation": frenchTranslation,
	}))
	spanishTranslation := "es(translation_id)"
	b.AddTranslation(spanishLanguage, testNewTranslation(t, map[string]interface{}{
		"id":          translationID,
		"translation": spanishTranslation,
	}))

	tests := []struct {
		languageIDs      []string
		valid            bool
		result           string
		expectedLanguage *language.Language
	}{
		{
			[]string{"invalid"},
			false,
			translationID,
			nil,
		},
		{
			[]string{"invalid", "invalid2"},
			false,
			translationID,
			nil,
		},
		{
			[]string{"invalid", "en-US"},
			true,
			englishTranslation,
			englishLanguage,
		},
		{
			[]string{"en-US", "invalid"},
			true,
			englishTranslation,
			englishLanguage,
		},
		{
			[]string{"en-US", "fr-FR"},
			true,
			englishTranslation,
			englishLanguage,
		},
		{
			[]string{"invalid", "es"},
			true,
			spanishTranslation,
			spanishLanguage,
		},
		{
			[]string{"zh-CN,fr-XX,es"},
			true,
			spanishTranslation,
			spanishLanguage,
		},
	}

	for i, test := range tests {
		tf, err := b.Tfunc(test.languageIDs[0], test.languageIDs[1:]...)
		if err != nil && test.valid {
			t.Errorf("Tfunc(%v) = error{%q}; expected no error", test.languageIDs, err)
		}
		if err == nil && !test.valid {
			t.Errorf("Tfunc(%v) = nil error; expected error", test.languageIDs)
		}
		if result := tf(translationID); result != test.result {
			t.Errorf("translation %d was %s; expected %s", i, result, test.result)
		}
		lang := b.SupportedLanguage(test.languageIDs[0], test.languageIDs[1:]...)
		if (lang == nil && test.expectedLanguage != nil) ||
			(lang != nil && test.expectedLanguage == nil) ||
			(lang != nil && test.expectedLanguage != nil && lang.String() != test.expectedLanguage.String()) {
			t.Errorf("lang %d was %s; expected %s", i, lang, test.expectedLanguage)
		}
	}
}

func testNewTranslation(t *testing.T, data map[string]interface{}) translation.Translation {
	translation, err := translation.NewTranslation(data)
	if err != nil {
		t.Fatal(err)
	}
	return translation
}

/*

func bundleFixture(t *testing.T) *Bundle {
	l, err := NewLocaleFromString("ar-EG")
	if err != nil {
		t.Errorf(err.Error())
	}
	return &Bundle{
		Locale: l,
		localizedStrings: map[string]*LocalizedString{
			"a": &LocalizedString{
				ID: "a",
			},
			"b": &LocalizedString{
				ID:          "b",
				Translation: "translation(b)",
			},
			"c": &LocalizedString{
				ID: "c",
				Translations: map[PluralCategory]*PluralTranslation{
					Zero:  NewPluralTranslation("zero(c)"),
					One:   NewPluralTranslation("one(c)"),
					Two:   NewPluralTranslation("two(c)"),
					Few:   NewPluralTranslation("few(c)"),
					Many:  NewPluralTranslation("many(c)"),
					Other: NewPluralTranslation("other(c)"),
				},
			},
			"d": &LocalizedString{
				ID: "d",
				Translations: map[PluralCategory]*PluralTranslation{
					Zero: NewPluralTranslation("zero(d)"),
					One:  NewPluralTranslation("one(d)"),
				},
			},
		},
	}
}
*/
