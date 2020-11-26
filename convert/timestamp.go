package convert

import "unicode/utf8"

func Nano2Milli(nanoTimeStampStr string) string {
	length := utf8.RuneCountInString(nanoTimeStampStr)
	if length < 6 {
		return nanoTimeStampStr
	}
	return nanoTimeStampStr[:length-6]
}

func Milli2Nano(milliTimestamp string) string {
	return milliTimestamp + "000000"
}
