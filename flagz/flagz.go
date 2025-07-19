package flagz;

import(
	OS     "os"
	IsaTTY "github.com/mattn/go-isatty"
);



type Flagz struct {
	IsTTY  bool
	Colors bool
	Flags  map[string]Flag
}



var parsed *Flagz = nil;

func Parse() *Flagz {
	istty := IsaTTY.IsTerminal(OS.Stdout.Fd()) ||
		IsaTTY.IsCygwinTerminal(OS.Stdout.Fd());
	parsed = &Flagz{
		IsTTY:  istty,
		Colors: istty,
	};
	return parsed;
}



func (flagz *Flagz) Colors() {
	flagz.Colors = true;
}
func (flagz *Flagz) NoColors() {
	flagz.Colors = false;
}


































