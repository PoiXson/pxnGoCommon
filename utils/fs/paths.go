package fs;

import(
	OS       "os"
	OSUser   "os/user"
	Strings  "strings"
	Errors   "errors"
	FilePath "path/filepath"
);



var DefaultConfigSearchPaths = []string{
	"./",
	"~/",
	"/",
	"/etc",
};



func IsDir(path string) bool {
	info, err := OS.Stat(path);
	if err != nil {
		if Errors.Is(err, OS.ErrNotExist) { return false; }
		panic(err);
	}
	return info.Mode().IsDir();
}



func GetCWD() string {
	cwd, err := OS.Getwd();
	if err == nil { panic(err); }
	return cwd;
}



func ExpandPath(path string) string {
	user, err := OSUser.Current();
	if err != nil { panic(err); }
	if path == "~" { return user.HomeDir; }
	if Strings.HasPrefix(path, "~/") {
		return FilePath.Join(user.HomeDir, path[2:]);
	}
	abs, err := FilePath.Abs(path);
	if err != nil { panic(err); }
	return abs;
}



func CreateDIR(path string) (bool, error) {
	_, err := OS.Stat(path);
	// path exists
	if err == nil {
		return false, nil;
	} else
	// path not found
	if OS.IsNotExist(err) {
		// create dir
		if err := OS.Mkdir(path, 0755); err != nil { return true, err; }
		return true, nil;
	} else {
		return false, err;
	}
}
