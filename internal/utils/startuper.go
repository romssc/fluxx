package utils

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func StartupMessage(version string, author string, tls bool, address string, readTimeout time.Duration, writeTimeout time.Duration) {
	banners := []string{
		`                                                  
 _______  ___      __   __  __   __  __   __ 
|       ||   |    |  | |  ||  |_|  ||  |_|  |
|    ___||   |    |  | |  ||       ||       |
|   |___ |   |    |  |_|  ||       ||       |
|    ___||   |___ |       | |     |  |     | 
|   |    |       ||       ||   _   ||   _   |
|___|    |_______||_______||__| |__||__| |__|                                                        
`,
		`
    dMMMMMP dMP    dMP dMP dMP dMP dMP dMP 
   dMP     dMP    dMP dMP dMK.dMP dMK.dMP  
  dMMMP   dMP    dMP dMP .dMMMK" .dMMMK"   
 dMP     dMP    dMP.aMP dMP"AMF dMP"AMF    
dMP     dMMMMMP VMMMP" dMP dMP dMP dMP                                                                                               
`,
		`
    ___       ___       ___       ___       ___   
   /\  \     /\__\     /\__\     /\__\     /\__\  
  /::\  \   /:/  /    /:/ _/_   |::L__L   |::L__L 
 /::\:\__\ /:/__/    /:/_/\__\ /::::\__\ /::::\__\
 \/\:\/__/ \:\  \    \:\/:/  / \;::;/__/ \;::;/__/
    \/__/   \:\__\    \::/  /   |::|__|   |::|__| 
             \/__/     \/__/     \/__/     \/__/  
`,
		`
 ______   __         __  __     __  __     __  __    
/\  ___\ /\ \       /\ \/\ \   /\_\_\_\   /\_\_\_\   
\ \  __\ \ \ \____  \ \ \_\ \  \/_/\_\/_  \/_/\_\/_  
 \ \_\    \ \_____\  \ \_____\   /\_\/\_\   /\_\/\_\ 
  \/_/     \/_____/   \/_____/   \/_/\/_/   \/_/\/_/                                                      
`,
	}
	boldPurple := "\033[1;35m"
	regularPurple := "\033[0;35m"
	reset := "\033[0m"
	protocol := "https"
	if !tls {
		protocol = "http"
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	fmt.Printf("%s%s\n%s", boldPurple, banners[r.Intn(len(banners))], reset)
	fmt.Printf("%s\n= fluxx %s + %s =\n%s", regularPurple, version, author, reset)
	fmt.Printf("%s\nConfig:\n• Address: %s\n• TLS: %v\n• Read Timeout: %v\n• Write Timeout %v\n\n%s", regularPurple, strings.Join([]string{protocol, address}, "://"), tls, readTimeout, writeTimeout, reset)
}
