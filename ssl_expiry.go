package main

import (
	"crypto/tls"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

// Basic type is a HostPort string like "localhost:8443"
type HostPort string

// Probes is a list of HostPorts
type Probes []HostPort

// And a ProbeConfig struct is a mapping of a Probes list to the  "probes:" heading in a YAML config
type ProbeConfig struct {
	Probes `yaml:"probes"`
}

func probe(hp HostPort) (timestamp int64, err error) {
	// Given a HostPort, attempt to connect via SSL.
	// On failure, return an error
	// On success, return the certificate's expiration date as a Unix timestamp
	config := tls.Config{
		InsecureSkipVerify: true,
	}
	conn, err := tls.Dial("tcp", string(hp), &config)
	if err != nil {
		log.Printf("ERROR dial: %s\n", err)
		return timestamp, err
	}
	defer conn.Close()
	log.Printf("client: connected to %s\n ", conn.RemoteAddr())
	s := conn.ConnectionState()
	for _, cert := range s.PeerCertificates {
		log.Printf("CN: %s\n", cert.Subject)
		timestamp = cert.NotAfter.Unix() // thus we naturally always get the last timestamp
	}
	return timestamp, err
}

func main() {
	conffile := "./probes.yaml"
	data, err := ioutil.ReadFile(conffile)
	if err != nil {
		log.Printf("ERROR cannot read file %s: %v\n", conffile, err)
	}
	config := ProbeConfig{}
	unmarerr := yaml.Unmarshal(data, &config)
	if unmarerr != nil {
		log.Printf("ERROR shit: %v\n", unmarerr)
	}
	log.Printf("--- config:\n%v\n\n", config)
	for i, hp := range config.Probes {
		log.Printf("Probe %d: %s\n", i, hp)
		ts, err := probe(hp)
		if err != nil {
			log.Printf("ERROR Failed to probe %s\n", hp)
		}
		log.Printf("Probe %d timestamp: %d\n", i, ts)
	}

}
