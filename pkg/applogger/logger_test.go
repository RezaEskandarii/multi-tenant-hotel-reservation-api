package applogger

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

const (
	infoLogTxt    = "this_text_is_test_for_info_level_log"
	errorLogTxt   = "this_text_is_test_for_error_level_log"
	debugLogTxt   = "this_text_is_test_for_debug_level_log"
	warningLogTxt = "this_text_is_test_for_warning_level_log"
)

func TestCanWriteLog(t *testing.T) {

	logger := AppLogger{}

	t.Run("test_can_write_info_log", func(t *testing.T) {
		logger.LogInfo(infoLogTxt)
		_, err := os.Stat(getLogFileName())
		assert.Nil(t, err)
		assert.NotEqual(t, err, os.ErrNotExist)

		content, err := ioutil.ReadFile(getLogFileName())
		assert.Nil(t, err)
		assert.True(t, strings.Contains(string(content), infoLogTxt))
		fmt.Println(string(content))
	})
}
