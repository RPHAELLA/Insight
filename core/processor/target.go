package processor

import (
	"fmt"
	"net"
	"strings"
	"../util"
	"../config"
)

/*
Function: To process target by mapping them into IP and URL. 
Perform DNS resolve

[Future enhancement]
Option to perform subdomain search and add to list
*/
func SetTargets(addresses string) map[string]map[string][]string {
	command := fmt.Sprintf("nmap -sn %s --randomize-hosts", addresses)
	return process_targets(util.Execute(command))
}

func process_targets(result []string) map[string]map[string][]string {
    addresses := make(map[string]map[string][]string)
	for _, line := range result {
		if strings.Contains(line, "Nmap scan report for") {
			s := strings.Split(line, " ")
			ip_address := s[len(s)-1]
			
			domain := []string{""}

			switch len(s) {
            case 5:
				domain, _ = net.LookupAddr(ip_address)
			case 6: 
			    domain[0] = s[len(s)-2]
			}
			
			sub_key := make(map[string][]string)
			sub_key["domain"] = domain
			addresses[ip_address] = sub_key

			if len(domain) != 0 {
				util.CreateFolder([]string{config.C_BASE, strings.Join([]string{util.Strip(ip_address, "()")," - ", domain[0]}, "")})
			} else {
				util.CreateFolder([]string{config.C_BASE, ip_address})
			}
		}
	}
	
	return addresses
}

func check_ipv4(host string) bool {
	address := net.ParseIP(host)
	if address.To4() == nil {
		return false
	}
	return true
}