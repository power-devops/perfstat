// Copyright 2020 Power-Devops.com. All rights reserved.
// Use of this source code is governed by the license
// that can be found in the LICENSE file.
/*
Package perfstat is Go interface to IBM AIX libperfstat.
To use it you need AIX with installed bos.perf.libperfstat. You can check, if is installed using the following command:

	$ lslpp -L bos.perf.perfstat

The package is written using Go 1.14.7 and AIX 7.2 TL5. It should work with earlier TLs of AIX 7.2, but I
can't guarantee that perfstat structures in the TLs have all the same fields as the structures in AIX 7.2 TL5.

For documentation of perfstat on AIX and using it in programs refer to the official IBM documentation:
https://www.ibm.com/support/knowledgecenter/ssw_aix_72/performancetools/idprftools_perfstat.html
*/
package perfstat
