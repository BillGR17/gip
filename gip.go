package main

import (
  "net"
  "log"
  "os/exec"
  "strings"
)

func main() {
  addr,e:=net.LookupIP("domaingoes.here")
  if e!= nil {
    log.Fatal("Coulnt get ip from host")
  } else {
    out,e:=exec.Command("/bin/bash","-c","iptables -L -n --line-numbers|grep HOME|tr -s '[:space:]'|tr -d '\n'").Output()
    if e!= nil {
      log.Fatal("Iptables Search Failed")
    }
    o:=strings.Split(string(out)," ")
    line,ip := o[0],o[4]
    if ip!= addr[0].String() {
      exec.Command("/bin/bash","-c","iptables -R INPUT "+line+" -j ACCEPT -s "+addr[0].String()+" -m comment --comment 'HOME'").Run();
    }
  }
}
