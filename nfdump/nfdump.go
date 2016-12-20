package nfdump

import (
   "time"
)

const flowHead uint8 = 0x34

type Flow struct {
   StartDate time.Time
   EndDate time.Time
   Protocol int
   SrcIp string
   SrcPort int
   DstIp string
   DstPort int
   Packets int64
   Bytes int64
   Flows int
}

func Decode(data []uint8) (map[int]Flow, error) {
   flow := make(map[int]Flow)

   i := 354
   for i < len(data) {
      if data[i] == flowHead {
         miliIni := HexToInt([]uint8{data[i + 7], data[i + 6]})
         miliFim := HexToInt([]uint8{data[i + 9], data[i + 8]})
         dateIniEpoch := HexToInt([]uint8{data[i + 13], data[i + 12], data[i + 11], data[i + 10]})
         dateFimEpoch := HexToInt([]uint8{data[i + 17], data[i + 16], data[i + 15], data[i + 14]})
         startDate, err := EpochtimeToTime(dateIniEpoch, miliIni)
         if err != nil {
            return nil, err
         }
         endDate, err := EpochtimeToTime(dateFimEpoch, miliFim)
         if err != nil {
            return nil, err
         }

         protocol := int(HexToInt([]uint8{data[i + 21], data[i + 20]}))
         srcPort  := int(HexToInt([]uint8{data[i + 23], data[i + 22]}))
         dstPort  := int(HexToInt([]uint8{data[i + 25], data[i + 24]}))
         srcIp    := Uint8ToString(data[i + 29]) + "." + Uint8ToString(data[i + 28]) + "." + Uint8ToString(data[i + 27]) + "." + Uint8ToString(data[i + 26])
         dstIp    := Uint8ToString(data[i + 33]) + "." + Uint8ToString(data[i + 32]) + "." + Uint8ToString(data[i + 31]) + "." + Uint8ToString(data[i + 30])
         packets  := HexToInt([]uint8{data[i + 37], data[i + 36], data[i + 35], data[i + 34]})
         bytes    := HexToInt([]uint8{data[i + 41], data[i + 40], data[i + 39], data[i + 38]})
         flows    := int(HexToInt([]uint8{data[i + 51], data[i + 50]}))

         f := Flow {
            StartDate: startDate,
            EndDate: endDate,
            Protocol: protocol,
            SrcIp: srcIp,
            SrcPort: srcPort,
            DstIp: dstIp,
            DstPort: dstPort,
            Packets: packets,
            Bytes: bytes,
            Flows: flows,
         }
         flow[len(flow)] = f
      }

      i = i + 52
   }

   return flow, nil
}

func Encode() {

}