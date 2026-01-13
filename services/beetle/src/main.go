package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"sigs.k8s.io/yaml"
)

func main() {
	// Paths where the ConfigMap files will be mounted
	// paths := map[string]string{
	// 	"global":  "/etc/config/global-config.yaml",
	// 	"service": "/etc/config/service-config.yaml",
	// }
	//
	// out := make(map[string]interface{})
	//
	// for key, path := range paths {
	// 	data, err := ioutil.ReadFile(path)
	// 	if err != nil {
	// 		log.Printf("WARN: could not read %s (%s): %v", key, path, err)
	// 		continue
	// 	}
	//
	// 	var parsed interface{}
	// 	if err := yaml.Unmarshal(data, &parsed); err != nil {
	// 		log.Fatalf("ERROR: failed to parse %s: %v", path, err)
	// 	}
	//
	// 	out[key] = parsed
	// }
	//
	// pretty, err := json.MarshalIndent(out, "", "  ")
	// if err != nil {
	// 	log.Fatalf("ERROR: failed to marshall JSON: %v", err)
	// }
	//
	// fmt.Println(string(pretty))
	fmt.Println("hello creature...")
	//
	// // Keep container alive so you can check logs
	// for {
	// 	time.Sleep(10 * time.Minute)
	// }
}

