# godynu
CLI for the free dynamic DNS service Dynu.com

# Usage
```  godynu [command]

Available Commands:
  dns         Work with dynu dynamic dns
  help        Help about any command
  ping        ping API call
  token       get the user API token from dynu

Flags:
      --config string   config file (default is ./config.yaml)
  -h, --help            help for godynu
  -t, --toggle          Help message for toggle
```
  
  
# Add a new dns record and IP

>  godynu dns add -d myname.dynu.net --ip \`curl ifconfig.io\`

```
Flags:
  
  -d, --domain string   Dns record to add
  -h, --help            help for add    
      --ip string       IPV4 to add as resolution for the domain provided
```
