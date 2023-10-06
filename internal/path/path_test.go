package path_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

	const (
		path1 string = "/some/path/to/" // prefix and suffix
		path2 string = "/some/path/to" // prefix
		path3 string = "some/path/to/" // suffix
		path4 string = "some/path/to" // none
		path5 string = "./some/path/to/" // point prefix and suffix
		path6 string = "./some/path/to" // point prefix
	


		//windows paths

		path7 string = "C:\\some\\path\\to\\" // absolute prefix and suffix
		path8 string = "C:\\some\\path\\to" // absolute prefix
		path9 string = "some\\path\\to\\" // suffix
		path10 string = "some\\path\\to" // none
		path11 string = ".\\some\\path\\to\\" // point prefix and suffix
		path12 string = ".\\some\\path\\to" // point prefix

	)





//WORKING HERE
func TestNormalizePath(t *testing.T ){






	assert.Equal(t, true, true)
}