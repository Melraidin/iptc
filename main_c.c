#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#include <libiptcdata/iptc-data.h>

IptcDataSet *get_iptc_dataset(IptcData *iptcData, unsigned int);

IptcDataSet *get_iptc_dataset(IptcData *iptcData, unsigned int i) {
	return iptcData->datasets[i];
}
