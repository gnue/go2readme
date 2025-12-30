package main

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/gnue/fieldfmt"
	"github.com/gnue/go2readme/version"
)

type Version struct {
	Version   string `field:"Version,omitempty"`
	Sum       string `field:"Checksum,omitempty"`
	GoVersion string `field:"Go Vision,omitempty"`
	Revision  string `field:"Git commit,omitempty"`
}

var (
	versionInfo *Version
)

func init() {
	info := version.New()
	versionInfo = &Version{Version: info.Main.Version, Sum: info.Main.Sum, Revision: info.ReversionString(7, "-dirty"), GoVersion: info.GoVersion}
}

func versionPrint() {
	w := tabwriter.NewWriter(os.Stdout, 0, 4, 1, ' ', 0)

	fmt.Fprintln(w, "Version:")
	fmt.Println(versionInfo)
	fieldfmt.Fprintf(w, " %v:\t%v\n", versionInfo)
	w.Flush()
}
