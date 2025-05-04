package testutil

import "flag"

var Update = flag.Bool("UPDATE", false, "update golden files by writing actual output to expected files")
