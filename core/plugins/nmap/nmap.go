package main

import (
    "fmt"
	"time"
	"sync"
    "regexp"
    "strings"
    "../../util"
    "../../config"
    // "github.com/gookit/color"
)

/*
	Engage run 
	
	if got dns -> subdomain
	Check subdomain takeover
*/

/*
This method is invoked when the plugin is loaded at the Engage stage
Function: To sweep / check online status of the target host(s)
*/
func Engage(options map[string]string) *map[string]string {
    command := fmt.Sprintf("nmap -sn %s --randomize-hosts", options["target"])
	process_engage(util.Execute(command))
	process_nmap(options)
    return &options
}

/*
This is a supporting method for Engage()
Function: To create a directory for each online host
*/
func process_engage(result []string) {
	for _, line := range result {
		if strings.Contains(line, "Nmap scan report for") {
            s := strings.Split(line, " ")
			switch len(s) {
            case 5:
                util.CreateFolder([]string{config.C_BASE, s[len(s)-1]})
            case 6: 
                util.CreateFolder([]string{config.C_BASE, strings.Join([]string{util.Strip(s[len(s)-1], "()"), " - ",s[len(s)-2]}, "")})
            }
		}
    }
}

/*
This method is invoked when the plugin is loaded at the Engage stage
Function: To perform network service scan on each online host
          Current host(s) are scanned sequentially from A-Z and 0-9
*/
func process_nmap(options map[string]string) *map[string]string {
	directory_list := util.GetDirectoryList()

	var wg sync.WaitGroup
	wg.Add(len(directory_list))
	done := make(chan bool)
	
	var mode string

	switch options["mode"] {
	case "quiet":
		mode = fmt.Sprintf("-sS -Pn -n -vvv -T2 -p-")
	case "quick": 
		mode = fmt.Sprintf("-sS -T5 -vvv")
	case "comprehensive": 
		mode = fmt.Sprintf("-sS -Pn -n -vvv -A -p-")
	default:
		mode = fmt.Sprintf("-sS -Pn -n")
	}

    for _, directory := range directory_list {
		go func(directory string, mode string) {
			defer wg.Done()
			util.CreateFolder([]string{config.C_BASE, directory, "nmap"})

			target, _ := util.FormatFolderName(directory)
			directory_path := util.FormatPath([]string{config.C_BASE, directory})
			output := util.FormatPath([]string{config.C_BASE, directory, "nmap", fmt.Sprintf("nmap_%s_%s", target, time.Now().Format("02012006150405"))})

			command := fmt.Sprintf("nmap %s -oA \"%s\" %s", mode, output, target)
			process_map(util.Execute(command), directory_path, done)
		}(directory, mode)
    }

    for range directory_list {
        <-done
    }
    close(done)
    
    return &options
}

func process_map(result []string, directory string, done chan bool) {
    services := make(map[string][]string)
	for _, line := range result {
		if strings.Contains(line, "tcp") && strings.Contains(line, "open") && !strings.Contains(line, "Discovered") {
			s := strings.Fields(line)
			for _, s_type := range process_map_standardise_service(s[2]) {
				services[s_type] = append(services[s_type], strings.Split(s[0],"/")[0])
			}
		}
	}
    util.Save(directory, "services.txt", services)
    done <- true
}

func process_map_standardise_service(service string) []string {
	flag := false
	s_types := []string{}

	if match, _ := regexp.MatchString("http", service); match { s_types = append(s_types, "http"); flag = true }
	if match, _ := regexp.MatchString("https|ftps|ssl", service); match { s_types = append(s_types, "ssl"); flag = true }
	if match, _ := regexp.MatchString("ftp", service); match { s_types = append(s_types, "ftp"); flag = true }
	if match, _ := regexp.MatchString("mysql", service); match { s_types = append(s_types, "mysql"); flag = true }
	if match, _ := regexp.MatchString("ms-sql", service); match { s_types = append(s_types, "mssql"); flag = true }
	if match, _ := regexp.MatchString("dns", service); match { s_types = append(s_types, "dns"); flag = true }
	if match, _ := regexp.MatchString("telnet", service); match { s_types = append(s_types, "telnet"); flag = true }
	if match, _ := regexp.MatchString("smb", service); match { s_types = append(s_types, "smb"); flag = true }
	if match, _ := regexp.MatchString("msrdp|ms-wbt-server", service); match { s_types = append(s_types, "msrdp"); flag = true }
	if match, _ := regexp.MatchString("smtp", service); match { s_types = append(s_types, "smtp"); flag = true }
	if match, _ := regexp.MatchString("snmp", service); match { s_types = append(s_types, "snmp"); flag = true }
	if match, _ := regexp.MatchString("ssh", service); match { s_types = append(s_types, "ssh"); flag = true }
	if match, _ := regexp.MatchString("msrpc|rpcbind", service); match { s_types = append(s_types, "msrpc"); flag = true }
	if match, _ := regexp.MatchString("netbios-ssn", service); match { s_types = append(s_types, "netbios"); flag = true }
	if match, _ := regexp.MatchString("ipp", service); match { s_types = append(s_types, "cups"); flag = true }
	if match, _ := regexp.MatchString("java-rmi", service); match { s_types = append(s_types, "java-rmi"); flag = true }
	if match, _ := regexp.MatchString("vnc", service); match { s_types = append(s_types, "vnc"); flag = true }
	if match, _ := regexp.MatchString("oracle", service); match { s_types = append(s_types, "oracle"); flag = true }
	if match, _ := regexp.MatchString("kerberos", service); match { s_types = append(s_types, "kerberos"); flag = true }
	if match, _ := regexp.MatchString("ldap", service); match { s_types = append(s_types, "ldap"); flag = true }

	if flag == false { s_types = append(s_types, service); }
	return s_types
}