package sprofile

import (
	"fmt"
	"log"
	"net"
	"os"
	"slices"

	"github.com/sblinch/kdl-go"
)

func ParseServerProfile(path string) (sp ServerProfile, err error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return
	}

	err = kdl.Unmarshal(data, &sp)
	if err != nil {
		return
	}

	return
}

func (sp ServerProfile) Validate(verbose bool) (err error) {

	err = sp.validateMetadata()
	if err != nil {
		return
	}

	err = sp.validateServer()
	if err != nil {
		return
	}

	err = sp.validateOperators()
	if err != nil {
		return
	}

	err = sp.validateListeners()
	if err != nil {
		return
	}

	if verbose {
		sp.printVerboseData()
	}

	return
}

func (sp ServerProfile) validateMetadata() (err error) {
	if sp.Profile.Name == "" {
		return fmt.Errorf("profile metadata missing name")
	}

	if sp.Profile.Description == "" {
		return fmt.Errorf("profile metadata missing description")
	}

	return
}

func (sp ServerProfile) validateServer() (err error) {
	if sp.Server.BindHost == "" {
		return fmt.Errorf("missing server bind-host")
	}

	if sp.Server.BindPort == 0 {
		return fmt.Errorf("missing server bind-port")
	}

	return
}

func (sp ServerProfile) validateOperators() (err error) {
	if sp.Operators.RootPassword == "" {
		return fmt.Errorf("missing root password")
	}

	for username, details := range sp.Operators.Users {
		if details.(map[string]any)["password"] == nil {
			return fmt.Errorf("user %s is missing password field", username)
		}
	}
	return
}

func (sp ServerProfile) validateListeners() (err error) {
	var ports []int64

	for lName, details := range sp.Listeners.HttpListeners {
		if details.(map[string]any)["bind-host"] == nil {
			return fmt.Errorf("listener %s is missing bind-host field", lName)
		}

		// validate valid ip address
		bHost := net.ParseIP(details.(map[string]any)["bind-host"].(string))
		if bHost == nil {
			return fmt.Errorf("listener %s has an invalid bind-host", lName)
		}

		// validate address exists on host
		if details.(map[string]any)["bind-host"].(string) != "127.0.0.1" && details.(map[string]any)["bind-host"].(string) != "0.0.0.0" {
			var sAddrs []string

			addrs, err := net.InterfaceAddrs()
			if err != nil {
				log.Fatal(err)
			}

			for _, addr := range addrs {
				if ipnet, ok := addr.(*net.IPNet); ok {
					if !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
						sAddrs = append(sAddrs, ipnet.IP.String())
					}
				}
			}

			if !slices.Contains(sAddrs, details.(map[string]any)["bind-host"].(string)) {
				return fmt.Errorf("unable to bind to host %s for listener %s", details.(map[string]any)["bind-host"].(string), lName)
			}
		}

		if details.(map[string]any)["bind-port"] == nil {
			return fmt.Errorf("listener %s is missing bind-port field", lName)
		}
		if slices.Contains(ports, details.(map[string]any)["bind-port"].(int64)) {
			return fmt.Errorf("bind port %d used multiple times", details.(map[string]any)["bind-port"].(int64))
		}
		ports = append(ports, details.(map[string]any)["bind-port"].(int64))

		if details.(map[string]any)["bind-port"].(int64) <= 0 || details.(map[string]any)["bind-port"].(int64) > 65535 {
			return fmt.Errorf("invalid bind port  for listener %s", lName)
		}

		if details.(map[string]any)["hosts"] == nil {
			return fmt.Errorf("listener %s is missing hosts field", lName)
		}

		if details.(map[string]any)["port"] == nil {
			return fmt.Errorf("listener %s is missing port field", lName)
		}

		if details.(map[string]any)["user-agent"] == nil {
			return fmt.Errorf("listener %s is missing user-agent field", lName)
		}

		if details.(map[string]any)["headers"] == nil {
			return fmt.Errorf("listener %s is missing headers field", lName)
		}

		if details.(map[string]any)["uris"] == nil {
			return fmt.Errorf("listener %s is missing uris field", lName)
		}

		if details.(map[string]any)["method"] == nil {
			return fmt.Errorf("listener %s is missing method field", lName)
		}

		if details.(map[string]any)["method"].(string) != "post" && details.(map[string]any)["method"].(string) != "get" {
			return fmt.Errorf("invalid method for listener %s", lName)
		}
	}

	return
}

func (sp ServerProfile) printVerboseData() {
	fmt.Println("[=== Profile Metadata ===]")
	fmt.Printf("Profile Name:\t\t%s\n", sp.Profile.Name)
	fmt.Printf("Profile Description:\t%s\n", sp.Profile.Description)
	fmt.Print("\n")

	fmt.Println("[=== Teamserver Config ===]")
	fmt.Printf("Bind Host:\t%s\n", sp.Server.BindHost)
	fmt.Printf("Bind Port:\t%d\n", sp.Server.BindPort)
	fmt.Print("\n")

	fmt.Println("[=== Users ===]")
	fmt.Printf("Root Password:\t%s\n\n", sp.Operators.RootPassword)

	for username, details := range sp.Operators.Users {
		fmt.Printf("User:\t\t%s\n", username)
		fmt.Printf("Password:\t%s\n", details.(map[string]any)["password"].(string))
		fmt.Print("\n")
	}

	fmt.Println("[=== HTTP Listeners ===]")
	for lName, details := range sp.Listeners.HttpListeners {
		fmt.Printf("Listener Name:\t%s\n", lName)
		fmt.Printf("Bind Host:\t%s\n", details.(map[string]any)["bind-host"].(string))
		fmt.Printf("Bind Port:\t%d\n", details.(map[string]any)["bind-port"].(int64))
		fmt.Printf("Hosts:\t\t%s\n", details.(map[string]any)["hosts"].(string))
		fmt.Printf("Port:\t\t%d\n", details.(map[string]any)["port"].(int64))
		fmt.Printf("User Agent:\t%s\n", details.(map[string]any)["user-agent"].(string))
		fmt.Printf("Headers:\t%s\n", details.(map[string]any)["headers"].(string))
		fmt.Printf("Uris:\t\t%s\n", details.(map[string]any)["uris"].(string))
		fmt.Printf("Method:\t\t%s\n", details.(map[string]any)["method"].(string))
		fmt.Print("\n")
	}

	fmt.Print("\n")
}
