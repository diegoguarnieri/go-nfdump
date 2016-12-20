package main

import (
   "fmt"
   "log"
   "os"
   "go-nfdump/nfdump"
)

func main() {

   args := os.Args[1:]

   if len(args) < 1 {
      printUsage()
   } else {

      for _, fileName := range args {
         data, err := nfdump.LoadFile(fileName)
         if err != nil {
            log.Fatalf("Error [%v] opening file: %v\n", err, fileName)
         }

         flows, nil := nfdump.Decode(data)
         if err != nil {
            log.Fatalf("Error [%v] decoding file: %v\n", err, fileName)
         }

         for i := 0; i < len(flows); i++ {
            v := flows[i]
            fmt.Printf("[%v - %v] prot: %v %v:%v <--> %v:%v pcks: %v bytes: %v flows: %v\n", v.StartDate, v.EndDate, v.Protocol, v.SrcIp, v.SrcPort, v.DstIp, v.DstPort, v.Packets, v.Bytes, v.Flows)
         }
      }
   }
}

func printUsage() {
   fmt.Println("Usage: go-nfdump [FILE]...")
}