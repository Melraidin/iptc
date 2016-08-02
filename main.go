/*
	Golang bindings for libiptcdata.
*/
package iptc

/*
#cgo LDFLAGS: -liptcdata

#include <stdlib.h>
#include <libiptcdata/iptc-data.h>

IptcDataSet *get_iptc_dataset(IptcData *iptcData, unsigned int i);
*/
import "C"

import (
	"errors"
	"fmt"
	"unsafe"
)

type Data map[int]map[int]string

var (
	ErrNoIptcData = errors.New("no IPTC data in file %s")
)

/**
 * Opens the given path and attempts to return any IPTC data read.
 */
func Open(file string) (Data, error) {

	cfile := C.CString(file)

	iptcData := C.iptc_data_new_from_jpeg(cfile)

	C.free(unsafe.Pointer(cfile))

	if iptcData == nil {
		return nil, fmt.Errorf(ErrNoIptcData.Error(), file)
	}

	defer func() {
		C.iptc_data_unref(iptcData)
	}()

	return parseIptcData(iptcData)
}

/**
 * Parses an IPTC data blob generating a map of records and tags to
 * string values.
 */
func parseIptcData(iptcData *C.IptcData) (Data, error) {
	parsed := Data{}

	for i := C.uint(0); i < iptcData.count; i++ {
		dataSet := C.get_iptc_dataset(iptcData, i)

		if parsed[int(dataSet.record)] == nil {
			parsed[int(dataSet.record)] = make(map[int]string)
		}

		switch C.iptc_dataset_get_format(dataSet) {
		case C.IPTC_FORMAT_BYTE, C.IPTC_FORMAT_SHORT, C.IPTC_FORMAT_LONG:
			parsed[int(dataSet.record)][int(dataSet.tag)] = fmt.Sprintf("%d", C.iptc_dataset_get_value(dataSet))
		case C.IPTC_FORMAT_BINARY:
			value := make([]C.char, 256)
			C.iptc_dataset_get_as_str(dataSet, &value[0], C.uint(len(value)))

			// Guard against data being shorter than specified.
			dataLen := dataSet.size*3 - 1
			if dataLen < 0 {
				dataLen = 0
			} else {
				if int(dataLen) > len(value)-1 {
					dataLen = C.uint(len(value) - 1)
				}
			}
			parsed[int(dataSet.record)][int(dataSet.tag)] = fmt.Sprintf("%c", value[dataLen])
		default:
			value := make([]C.uchar, 256)
			actualLength := C.iptc_dataset_get_data(dataSet, &value[0], C.uint(len(value)))
			parsed[int(dataSet.record)][int(dataSet.tag)] = fmt.Sprintf("%s", value[:actualLength-1])
		}
	}

	return parsed, nil
}
