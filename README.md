#Youtube Proxy

####Why
My current company is bandwidth constrained across multiple sites because we use an MPLS network. Because of this, we are currently approving youtube.com access on a case-by-case basis and permissioning individuals over the firewall using their IP address. This process is cumbersome to maintain, and as the approved user-base grows, the amount of control/monitoring of the bandwidth we have decreases.

This server is designed to run on a server in the co-location facility and be the only entity permissioned for youtube.com. All employees will be able to access the server, and will be proxied out to youtube. Because all youtube bandwidth will be originating from a single server, we have an increased ability to monitor and control the bandwidth. Further, the server will log every request so monitoring of individuals is simplified as well. 

Also, because I wanted to get my hands dirty in Go.