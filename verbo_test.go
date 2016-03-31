package verbo

import (
	"testing"
)

func equal(t *testing.T, src, dest string) {
	if src != dest {
		t.Errorf("Wrong result: %s should be: %s", src, dest)
	}
}

func equalInt(t *testing.T, src, dest int) {
	if src != dest {
		t.Errorf("Wrong result: %d should be: %d", src, dest)
	}
}

func equalArray(t *testing.T, src, dest []string) {
	if len(src) != len(dest) {
		t.Errorf("Wrong result: %q should be: %q", src, dest)
	}
}

func TestIsBlank(t *testing.T) {
	if !IsBlank(" ") {
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

func TestChop(t *testing.T) {
	equalInt(t, len(Chop("whitespace", 2)), 5)
	equalInt(t, len(Chop("whitespace", 3)), 4)
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

	equal(t, CleanDiacritics("ä"), "a")
	equal(t, CleanDiacritics("Ä"), "A")
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

func TestLeftPad(t *testing.T) {
	equal(t, LeftPad("1", 8, ""), "       1")
	equal(t, LeftPad("1", 8, "0"), "00000001")
}

func TestLevenstein(t *testing.T) {
	equalInt(t, Levenshtein("Godfather", "Godfather"), 0)
  equalInt(t, Levenshtein("Godfather", "Godfathe"), 1)
  equalInt(t, Levenshtein("Godfather", "odfather"), 1)
  equalInt(t, Levenshtein("Godfather", "godfather"), 1)
  equalInt(t, Levenshtein("Godfather", "Gdfthr"), 3)
  equalInt(t, Levenshtein("seven", "eight"), 5)
}

func TestLines(t *testing.T) {
	equalInt(t, len(Lines("Hello\nWorld")), 2)
  equalInt(t, len(Lines("Hello\rWorld")), 2)
  equalInt(t, len(Lines("Hello World")), 1)
  equalInt(t, len(Lines("\r\n\n\r")), 4)
  equalInt(t, len(Lines("Hello\r\r\nWorld")), 3)
  equalInt(t, len(Lines("Hello\r\rWorld")), 3)
}

func TestPad(t *testing.T) {
	equal(t, Pad("1", 8, "", "left"), "       1")
	equal(t, Pad("1", 8, "0", "left"), "00000001")
	equal(t, Pad("1", 8, "0", "left"), "00000001")
	equal(t, Pad("1", 8, "0", "right"), "10000000")
	equal(t, Pad("1", 8, "0", "both"), "00001000")
	equal(t, Pad("foo", 8, "0", "both"), "000foo00")
	equal(t, Pad("foo", 7, "0", "both"), "00foo00")
}

func TestPred(t *testing.T) {
	equal(t, Pred("b"), "a")
	equal(t, Pred("B"), "A")
	equal(t, Pred(","), "+")
}

/*
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
*/

func TestRepeat(t *testing.T) {
	equal(t, Repeat("foo", 0, ""), "")
	equal(t, Repeat("foo", 3, ""), "foofoofoo")
}

func TestReverse(t *testing.T) {
	equal(t, Reverse("foo"), "oof" )
  equal(t, Reverse("foobar"), "raboof" )
  equal(t, Reverse("foo bar"), "rab oof" )
  equal(t, Reverse("saippuakauppias"), "saippuakauppias" )
}

func TestRightPad(t *testing.T) {
	equal(t, RightPad("1", 8, ""), "1       ")
	equal(t, RightPad("1", 8, "0"), "10000000")
	equal(t, RightPad("foo", 8, "0"), "foo00000")
	equal(t, RightPad("foo", 7, "0"), "foo0000")
}

func TestSlugify(t *testing.T) {
	equal(t, Slugify("Jack & Jill like numbers 1,2,3 and 4 and silly characters ?%.$!/"), "jack-jill-like-numbers-1-2-3-and-4-and-silly-characters")
	equal(t, Slugify("Un éléphant à l\"orée du bois"), "un-elephant-a-l-oree-du-bois")
	equal(t, Slugify("I know latin characters: á í ó ú ç ã õ ñ ü ă ș ț"), "i-know-latin-characters-a-i-o-u-c-a-o-n-u-a-s-t")
	equal(t, Slugify("I am a word too, even though I am but a single letter: i!"), "i-am-a-word-too-even-though-i-am-but-a-single-letter-i")
	equal(t, Slugify("Some asian 天地人 characters"), "some-asian-characters")
	equal(t, Slugify("SOME Capital Letters"), "some-capital-letters")
}

func TestSucc(t *testing.T) {
	equal(t, Succ("a"), "b")
  equal(t, Succ("A"), "B")
  equal(t, Succ("+"), ",")
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

func TestUnquote(t *testing.T) {
	equal(t, Unquote("\"foo\"", ""), "foo")
  equal(t, Unquote("\"\"foo\"\"", ""), "\"foo\"")
  equal(t, Unquote("\"1\"", ""), "1")
  equal(t, Unquote("\"foo\"", "\""), "foo")
}

func TestWords(t *testing.T) {
	equalArray(t, Words("I love you!", ""), []string{"I", "love", "you!"})
	equalArray(t, Words(" I    love   you!  ", ""), []string{"I", "love", "you!"})
	equalArray(t, Words("I_love_you!", "_"), []string{"I", "love", "you!"})
	equalArray(t, Words("I-love-you!", "-"), []string{"I", "love", "you!"})
}
