package arguments

import (
	"flag"
)

type Arguments struct {
	Debug			bool
	IsServer		bool
	Host			string
	Port			string
	CertPath		string
	KeyPath			string
}

func GetArguments() *Arguments {
	args := &Arguments{}

	flag.BoolVar(&args.Debug, "debug", true, "Run in debug mode?")
	flag.BoolVar(&args.IsServer, "is_server", true, "Is the current client server or not")
	flag.StringVar(&args.Host, "host", "localhost", "Host location")
	flag.StringVar(&args.Port, "port", "3000", "Port of the host location")
	flag.StringVar(&args.CertPath, "cert_path", "/", "Path to the certification .pem file")
	flag.StringVar(&args.KeyPath, "key_path", "/", "Path to the .key file")

	flag.Parse()

	return args
}