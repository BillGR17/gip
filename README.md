# gip
Grabs IP from Domain and updates the iptables

In order for this to work you need to have a comment with ``HOME`` (unless you change the script to something else) in your iptables
Then when the script runs it will get the ip from the domain name and checks it if its different than the last time  
and if it is it will update the iptable

You need the Go languange installed

for ubuntu 
``sudo apt install golang``

for arch 
``sudo pacman -S go``

For a single run ``go run gip.go``
If you want exec file 
``go build gip.go``
``./gip``
