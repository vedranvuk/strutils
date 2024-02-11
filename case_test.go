package strutils

import "testing"

var ops int = 1e6

type sample struct {
	str, out string
}

func TestCamelcase(t *testing.T) {
	samples := []sample{
		{"sample text", "sampleText"},
		{"sample-text", "sampleText"},
		{"sample_text", "sampleText"},
		{"sample___text", "sampleText"},
		{"sampleText", "sampleText"},
		{"inviteYourCustomersAddInvites", "inviteYourCustomersAddInvites"},
		{"sample 2 Text", "sample2Text"},
		{"   sample   2    Text   ", "sample2Text"},
		{"   $#$sample   2    Text   ", "sample2Text"},
		{"SAMPLE 2 TEXT", "sample2Text"},
		{"___$$Base64Encode", "base64Encode"},
		{"FOO:BAR$BAZ", "fooBarBaz"},
		{"FOO#BAR#BAZ", "fooBarBaz"},
		{"something.com", "somethingCom"},
		{"$something%", "something"},
		{"something.com", "somethingCom"},
		{"•¶§ƒ˚foo˙∆˚¬", "foo"},
	}

	for _, sample := range samples {
		if out := CamelCase(sample.str); out != sample.out {
			t.Errorf("got %q from %q, expected %q", out, sample.str, sample.out)
		}
	}
}

func BenchmarkCamelcase(t *testing.B) {
	for i := 0; i < t.N; i++ {
		CamelCase("some sample text here_noething:too$amazing")
	}
}

func TestSnakecase(t *testing.T) {
	samples := []sample{
		{"@49L0S145_¬fwHƒ0TSLNVp", "49l0s145_fw_h_0tslnvp"},
		{"lk0B@bFmjrLQ_Z6YL", "lk0_b_b_fmjr_lq_z6yl"},
		{"samPLE text", "sam_ple_text"},
		{"sample text", "sample_text"},
		{"sample-text", "sample_text"},
		{"sample_text", "sample_text"},
		{"sample___text", "sample_text"},
		{"sampleText", "sample_text"},
		{"inviteYourCustomersAddInvites", "invite_your_customers_add_invites"},
		{"sample 2 Text", "sample_2_text"},
		{"   sample   2    Text   ", "sample_2_text"},
		{"   $#$sample   2    Text   ", "sample_2_text"},
		{"SAMPLE 2 TEXT", "sample_2_text"},
		{"___$$Base64Encode", "base64_encode"},
		{"FOO:BAR$BAZ", "foo_bar_baz"},
		{"FOO#BAR#BAZ", "foo_bar_baz"},
		{"something.com", "something_com"},
		{"$something%", "something"},
		{"something.com", "something_com"},
		{"•¶§ƒ˚foo˙∆˚¬", "foo"},
		{"CStringRef", "cstring_ref"},
		{"5test", "5test"},
		{"test5", "test5"},
		{"THE5r", "the5r"},
		{"5TEst", "5test"},
		{"_5TEst", "5test"},
		{"@%#&5TEst", "5test"},
		{"edf_6N", "edf_6n"},
		{"f_pX9", "f_p_x9"},
		{"p_z9Rg", "p_z9_rg"},
		{"2FA Enabled", "2fa_enabled"},
		{"Enabled 2FA", "enabled_2fa"},
	}

	for _, sample := range samples {
		if out := separatorCase(sample.str, underscoreByte); out != sample.out {
			t.Errorf("got %q from %q, expected %q", out, sample.str, sample.out)
		}
	}
}

func BenchmarkUnchangedLong(b *testing.B) {
	var s = "invite_your_customers_add_invites"
	b.SetBytes(int64(len(s)))
	for n := 0; n < b.N; n++ {
		separatorCase(s, underscoreByte)
	}
}

func BenchmarkUnchangedSimple(b *testing.B) {
	var s = "sample_text"
	b.SetBytes(int64(len(s)))
	for n := 0; n < b.N; n++ {
		separatorCase(s, underscoreByte)
	}
}

func BenchmarkModifiedUnicode(b *testing.B) {
	var s = "ß_ƒ_foo"
	b.SetBytes(int64(len(s)))
	for n := 0; n < b.N; n++ {
		separatorCase(s, underscoreByte)
	}
}
func BenchmarkModifiedLong(b *testing.B) {
	var s = "inviteYourCustomersAddInvites"
	b.SetBytes(int64(len(s)))
	for n := 0; n < b.N; n++ {
		separatorCase(s, underscoreByte)
	}
}

func BenchmarkModifiedLongSpecialChars(b *testing.B) {
	var s = "FOO:BAR$BAZ__Sample    Text___"
	b.SetBytes(int64(len(s)))
	for n := 0; n < b.N; n++ {
		separatorCase(s, underscoreByte)
	}
}

func BenchmarkModifiedSimple(b *testing.B) {
	var s = "sample text"
	b.SetBytes(int64(len(s)))
	for n := 0; n < b.N; n++ {
		separatorCase("sample text", underscoreByte)
	}
}

func BenchmarkModifiedUnicode2(b *testing.B) {
	var s = "ẞ•¶§ƒ˚foo˙∆˚¬"
	b.SetBytes(int64(len(s)))
	for n := 0; n < b.N; n++ {
		separatorCase(s, underscoreByte)
	}
}

func BenchmarkLeadingUnderscoresDigitUpper(b *testing.B) {
	var s = "_5TEst"
	b.SetBytes(int64(len(s)))
	for n := 0; n < b.N; n++ {
		separatorCase(s, underscoreByte)
	}
}

func BenchmarkDigitUpper(b *testing.B) {
	var s = "5TEst"
	b.SetBytes(int64(len(s)))
	for n := 0; n < b.N; n++ {
		separatorCase(s, underscoreByte)
	}
}

func BenchmarkDigitUpper2(b *testing.B) {
	var s = "lk0B@bFmjrLQ_Z6YL"
	b.SetBytes(int64(len(s)))
	for n := 0; n < b.N; n++ {
		separatorCase(s, underscoreByte)
	}
}
