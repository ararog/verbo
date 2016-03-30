package verbo

import (
	"testing"
	"strings"
)

func equal(t *testing.T, src, dest string) {
	if src != dest {
		t.Errorf("Wrong result: %s", src)
	}
}

func TestIsBlank(t *testing.T) {
  if ! IsBlank(" ") {
    t.Errorf("Wrong result: %s", "should not go here")
  }
}

func TestCamelize(t *testing.T) {

	equal(t, Camelize("the_camelize_string_method", false), "theCamelizeStringMethod")
	equal(t, Camelize("webkit-transform", false), "webkitTransform")
	equal(t, Camelize("-the-camelize-string-method", false), "TheCamelizeStringMethod")
	equal(t, Camelize("_the_camelize_string_method", false), "TheCamelizeStringMethod")
	equal(t, Camelize("The-camelize-string-method", false), "TheCamelizeStringMethod")
	equal(t, Camelize("the camelize string method", false), "theCamelizeStringMethod")
	equal(t, Camelize(" the camelize  string method", false), "theCamelizeStringMethod")
	equal(t, Camelize("the camelize   string method", false), "theCamelizeStringMethod")
	equal(t, Camelize(" with   spaces", false), "withSpaces")
	equal(t, Camelize("_som eWeird---name-", false), "SomEWeirdName")
}

func TestCapitalize(t *testing.T) {

	equal(t, Capitalize("fabio", false), "Fabio")
	equal(t, Capitalize("fabio", false), "Fabio")
	equal(t, Capitalize("FOO", false), "FOO")
	equal(t, Capitalize("FOO", false), "FOO")
	equal(t, Capitalize("foO", false), "FoO")
	equal(t, Capitalize("FOO", true), "Foo")
	equal(t, Capitalize("foO", true), "Foo")
	equal(t, Capitalize("f", false), "F")
	equal(t, Capitalize("f", true), "F")
	equal(t, Capitalize("f", false), "F")
}

func TestClassify(t *testing.T) {

	equal(t, Classify("some_class_name"), "SomeClassName")
	equal(t, Classify("my wonderfull class_name"), "MyWonderfullClassName")
	equal(t, Classify("my wonderfull.class.name"), "MyWonderfullClassName")
	equal(t, Classify("myLittleCamel"), "MyLittleCamel")
	equal(t, Classify("myLittleCamel.class.name"), "MyLittleCamelClassName")
}

func TestClean(t *testing.T) {
	equal(t, Clean(" foo    bar   "), "foo bar")
}

func TestClearDiacritics(t *testing.T) {

	from := "ąàáäâãåæăćčĉęèéëêĝĥìíïîĵłľńňòóöőôõðøśșşšŝťțţŭùúüűûñÿýçżźž"
	to := "aaaaaaaaaccceeeeeghiiiijllnnoooooooossssstttuuuuuunyyczzz"

	equal(t, CleanDiacritics(from), to)
	equal(t, CleanDiacritics(strings.ToUpper(from)), strings.ToUpper(to))

	equal(t, CleanDiacritics("ä"), "a")
	equal(t, CleanDiacritics("Ä Ø"), "A O")
	equal(t, CleanDiacritics("1 foo ääkkönen"), "1 foo aakkonen")
	equal(t, CleanDiacritics("Äöö ÖÖ"), "Aoo OO")
	equal(t, CleanDiacritics(" ä "), " a ")
}

func TestDasherize(t *testing.T) {
	equal(t, Dasherize("the_dasherize_string_method"), "the-dasherize-string-method")
  equal(t, Dasherize("TheDasherizeStringMethod"), "-the-dasherize-string-method")
  equal(t, Dasherize("thisIsATest"), "this-is-a-test")
  equal(t, Dasherize("this Is A Test"), "this-is-a-test")
  equal(t, Dasherize("thisIsATest123"), "this-is-a-test123")
  equal(t, Dasherize("123thisIsATest"), "123this-is-a-test")
  equal(t, Dasherize("the dasherize string method"), "the-dasherize-string-method")
  equal(t, Dasherize("the  dasherize string method  "), "the-dasherize-string-method")
  equal(t, Dasherize("téléphone"), "téléphone")
  equal(t, Dasherize("foo$bar"), "foo$bar")
  equal(t, Dasherize("input with a-dash"), "input-with-a-dash")
}

func TestDecapitalize(t *testing.T) {
	equal(t, Decapitalize("Fabio"), "fabio")
	equal(t, Decapitalize("FOO"), "fOO")
}

func TestHumanize(t *testing.T) {
	equal(t, Humanize("the_humanize_string_method"), "The humanize string method")
	equal(t, Humanize("ThehumanizeStringMethod"), "Thehumanize string method")
	equal(t, Humanize("-ThehumanizeStringMethod"), "Thehumanize string method")
	equal(t, Humanize("the humanize string method"), "The humanize string method")
	equal(t, Humanize("the humanize_id string method_id"), "The humanize id string method")
	equal(t, Humanize("the  humanize string method  "), "The humanize string method")
	equal(t, Humanize("   capitalize dash-CamelCase_underscore trim  "), "Capitalize dash camel case underscore trim")
}

func TestPred(t *testing.T) {
	equal(t, Pred("b"), "a")
	equal(t, Pred("B"), "A")
	equal(t, Pred(","), "+")
}

func TestPrune(t *testing.T) {
	equal(t, Prune("Hello, cruel world", 6, " read more"), "Hello read more");
  equal(t, Prune("Hello, world", 5, "read a lot more"), "Hello, world");
  equal(t, Prune("Hello, world", 5, ""), "Hello...");
  equal(t, Prune("Hello, world", 8, ""), "Hello...");
  equal(t, Prune("Hello, cruel world", 15, ""), "Hello, cruel...");
  equal(t, Prune("Hello world", 22, ""), "Hello world");
  equal(t, Prune("Привет, жестокий мир", 6, " read more"), "Привет read more");
  equal(t, Prune("Привет, мир", 6, "read a lot more"), "Привет, мир");
  equal(t, Prune("Привет, мир", 6, ""), "Привет...");
  equal(t, Prune("Привет, мир", 8, ""), "Привет...");
  equal(t, Prune("Привет, жестокий мир", 16, ""), "Привет, жестокий...");
  equal(t, Prune("Привет, мир", 22, ""), "Привет, мир");
  equal(t, Prune("alksjd!!!!!!....", 100, ""), "alksjd!!!!!!....");
}

func TestRepeat(t *testing.T) {
	equal(t, Repeat("foo", 0, ""), "")
  equal(t, Repeat("foo", 3, ""), "foofoofoo")
}

func TestSlugify(t *testing.T) {
	equal(t, Slugify("Jack & Jill like numbers 1,2,3 and 4 and silly characters ?%.$!/"), "jack-jill-like-numbers-1-2-3-and-4-and-silly-characters")
  equal(t, Slugify("Un éléphant à l\"orée du bois"), "un-elephant-a-l-oree-du-bois")
  equal(t, Slugify("I know latin characters: á í ó ú ç ã õ ñ ü ă ș ț"), "i-know-latin-characters-a-i-o-u-c-a-o-n-u-a-s-t")
  equal(t, Slugify("I am a word too, even though I am but a single letter: i!"), "i-am-a-word-too-even-though-i-am-but-a-single-letter-i")
  equal(t, Slugify("Some asian 天地人 characters"), "some-asian-characters")
  equal(t, Slugify("SOME Capital Letters"), "some-capital-letters")
}

func TestSwapCase(t *testing.T) {
	equal(t, SwapCase("AaBbCcDdEe"), "aAbBcCdDeE")
  equal(t, SwapCase("Hello World"), "hELLO wORLD")
}

func TestUnderscored(t *testing.T) {
	equal(t, Underscored("the-underscored-string-method"), "the_underscored_string_method")
	equal(t, Underscored("theUnderscoredStringMethod"), "the_underscored_string_method")
	equal(t, Underscored("TheUnderscoredStringMethod"), "the_underscored_string_method")
	equal(t, Underscored(" the underscored  string method"), "the_underscored_string_method")
}
