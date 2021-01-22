package versioning

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	cases := []struct {
		ver      string
		semVer   *SemVer
		major    string
		minor    string
		patch    string
		matedata string
	}{
		{
			"1.9.0-caimaoy",
			Create(1, 9, 0, "-caimaoy"),
			"1",
			"9",
			"0",
			"-caimaoy",
		},
		{
			"1.2.0",
			Create(1, 2, 0, ""),
			"1",
			"2",
			"0",
			"",
		},
	}

	for _, c := range cases {
		if c.ver != c.semVer.String() {
			t.Fatalf("expect:%s, got:%s", c.ver, c.semVer.String())
		}
		assert.Equal(t, c.semVer.Major(), c.major)
		assert.Equal(t, c.semVer.Minor(), c.minor)
		assert.Equal(t, c.semVer.Patch(), c.patch)
		assert.Equal(t, c.semVer.Metadata(), c.matedata)
	}
}

func TestParse(t *testing.T) {
	cases := []struct {
		ver    string
		semVer SemVer
	}{
		{"1.0.0-alpha.1",
			SemVer{
				"1",
				"0",
				"0",
				"-alpha.1",
				[3]uint32{1, 0, 0},
			},
		},
		{"1.0.2-alpha",
			SemVer{
				"1",
				"0",
				"2",
				"-alpha",
				[3]uint32{1, 0, 2},
			},
		},
		{"1.0.0+20130313144700",
			SemVer{
				"1",
				"0",
				"0",
				"+20130313144700",
				[3]uint32{1, 0, 0},
			},
		},
		{"1.0.0rc",
			SemVer{
				"1",
				"0",
				"0",
				"rc",
				[3]uint32{1, 0, 0},
			},
		},
	}

	for _, c := range cases {
		semVer, err := Parse(c.ver)
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(*semVer, c.semVer) {
			t.Fatal(c.ver)
		}
		if semVer.String() != c.ver || c.semVer.Compare(semVer, nil) != 0 {
			t.Fatalf("expect:%s, got:%s", c.ver, semVer.String())
		}
	}
}

func TestParseErr(t *testing.T) {
	var a string = "1.2"
	var expectErrString = "invalid semantic version 2: 1.2"

	_, err := Parse(a)
	if err.Error() != expectErrString {
		t.Errorf("Parse(%s), err %v, expect err %s", a, err, expectErrString)
	}
}
