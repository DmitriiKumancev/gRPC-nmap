package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Ullaakut/nmap/v3"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	scanner, err := nmap.NewScanner(
		ctx,
		nmap.WithTargets("scanme.nmap.org"),
		nmap.WithPorts("80"),
		nmap.WithServiceInfo(),
		nmap.WithScripts("vulners"),
	)
	if err != nil {
		log.Fatalf("unable to create nmap scanner: %v", err)
	}
	result, _, err := scanner.Run()
	if err != nil {
		log.Fatalf("unable to run nmap scan: %v", err)
	}

	for _, host := range result.Hosts {
		if len(host.Ports) == 0 || len(host.Addresses) == 0 {
			continue
		}

		for _, port := range host.Ports {

			fmt.Printf("Port.ID:%v \n", port.ID)
			fmt.Printf("port.Service.Name: %s \n", port.Service.Name)
			fmt.Printf("port.Service.Name: %s \n", port.Service.Version)

			for _, script := range port.Scripts {

				if script.ID == "vulners" {


					for _, table := range script.Tables[0].Tables {

						for _, el := range table.Elements {

							fmt.Println(el.Key, "-", el.Value)
						}

					}

				}
			}

		}

	}

}
