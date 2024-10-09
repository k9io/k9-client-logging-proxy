
Join the Key9 Slack channel
---------------------------

[![Slack](./images/slack.png)](https://key9identity.slack.com/)


What is they Key9 Client Logging Proxy
--------------------------------------

Key9 Client Logging is a service that "proxies" data from Key9 Tail (k9-tail). Key9 Tail is a service that collects SSH logs from syslog and sends them to Key9. These logs are used to determine which public keys are used to access a machine within your organization. 

Use cases:
----------

Some networks might have restrictions that prohibit machines from accessing the Key client logging API directly.   In those cases, the Key9 client loggin proxy can be used as a centralized API access point for all machines within a restricted network. 

This proxy should only be used within networks with egress network restrictions. 

What software uses they Key9 client logging proxy?
--------------------------------------------------

The proxy can be used by k9-tail.

Building and installing the Key9 client logging proxy
-----------------------------------------------------

Make sure you have Golang installed! 

<pre>
$ go mod init k9-client-logging-proxy
$ go mod tidy
$ go build
$ sudo mkdir -p /opt/k9/etc /opt/k9/bin
$ sudo cp etc/k9-client-logging-proxy.yaml /opt/k9/etc
$ sudo cp k9-client-logging-proxy /opt/k9/bin
$ sudo /opt/k9/bin/k9-client-logging-proxy 	 # Run from the command line... Control-C exits
$ sudo cp k9-client-logging-proxy.service /etc/systemd/system
$ sudo systemctl enable k9-client-logging-proxy
$ sudo systemctl start k9-client-logging-proxy
</pre>

Prebuild Key9 client logging proxy
----------------------------------

If you are unable to access a Golang compiler, you can download pre-built/pre-compiled binaries. These binaries are available for various architectures (i386, amd64, arm64, etc) and multiple operating systems (Linux, Solaris, NetBSD, etc).

You can find those binaries at: https://github.com/k9io/k9-binaries/tree/main/k9-client-logging-proxy

You will need a copy of the 'k9-proxy' configuation file.  That is located at: 

https://github.com/k9io/k9-client-logging-proxy/blob/main/etc/k9-client-logging-proxy.yaml

