package args

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var (
	showHelp    bool
	showVersion bool
	ConfigPath  string
	Install     bool
	Uninstall   bool
)

var (
	Version   string
	GitHash   string
	BuildTime string
)

func init() {
	app, _ := filepath.Abs(os.Args[0])
	ext := filepath.Ext(os.Args[0])
	//替换后缀名.exe为.yaml
	cfg := strings.TrimSuffix(app, ext) + ".yaml"

	//log.Println("app.path", app)
	flag.BoolVar(&showVersion, "v", false, "show version")
	flag.BoolVar(&showHelp, "h", false, "show help")
	flag.StringVar(&ConfigPath, "c", cfg, "Configure path")
	flag.BoolVar(&Install, "i", false, "Install service")
	flag.BoolVar(&Uninstall, "u", false, "Uninstall service")
}

func Parse() {
	flag.Parse()
	if showHelp {
		flag.Usage()
		os.Exit(0)
	}
	if showVersion {
		fmt.Printf("Version: %s \n", Version)
		fmt.Printf("Git Hash: %s \n", GitHash)
		fmt.Printf("Build Time: %s \n", BuildTime)
		os.Exit(0)
	}
}
