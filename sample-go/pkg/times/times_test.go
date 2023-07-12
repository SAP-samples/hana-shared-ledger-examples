package times_test

import (
	"github.com/stretchr/testify/assert"
	"hana-shared-ledger-sample/pkg/times"
	"testing"
	"time"
)

func TestTimeIsInIntervall(t *testing.T) {
	check := assert.New(t)
	now := time.Now().UTC()

	r := times.TimeIsInIntervall(now, 1*time.Minute)
	check.True(r, "time should be in interval")

	r = times.TimeIsInIntervall(now.Add(-70*time.Second), 1*time.Minute)
	check.False(r, "time should not be in interval")

	r = times.TimeIsInIntervall(now.Add(70*time.Second), 1*time.Minute)
	check.False(r, "time should not be in interval")
}

func TestToString(t *testing.T) {
	check := assert.New(t)
	testTime := time.Date(2020, 8, 11, 12, 13, 14, 123456789, time.UTC)
	timeString := times.ToString(testTime)

	check.Equal("2020-08-11 12:13:14.1234567", timeString)
}

func TestParseString(t *testing.T) {
	check := assert.New(t)
	_, err := times.ParseString("11.03.1983 12:13:14")
	check.Error(err)

	testString := "2020-08-11 12:13:14.1234567"
	res, err := times.ParseString(testString)
	check.NoError(err)
	check.Equal(testString, times.ToString(res))

	testString = "2020-08-11 12:13:14.123456789"
	res, err = times.ParseString(testString)
	check.Error(err)
}
