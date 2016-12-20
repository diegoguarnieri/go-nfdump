package nfdump

import (
   "fmt"
   "io"
   "os"
   "time"
   "strconv"
)

func Uint8ToString(u uint8) string {
   return strconv.Itoa(int(u))
}

func EpochtimeToTime(epoch, mili int64) (time.Time, error) {
   i, err := strconv.ParseInt(fmt.Sprintf("%v%03v", epoch, mili), 10, 64)
   if err != nil {
      return time.Time{}, err
   }

   return time.Unix(0, i*int64(time.Millisecond)), nil
}

func HexToInt(h []uint8) int64{
   var str string = "0x"
   for i := 0; i < len(h); i++ {
      str += fmt.Sprintf("%02x", h[i])
   }
   //fmt.Println(str)
   r, _ := strconv.ParseInt(str, 0, 64)
   return r
}

func LoadFile(fileName string) ([]uint8, error) {
   file, err := os.Open(fileName)
   if err != nil {
      return nil, err
   }
   
   var data []uint8
   b := make([]byte, 1)
   i := 0
   for {
      line, err := file.Read(b)

      if err == io.EOF {
         break
      }
      
      data = append(data, b[:line][0])
      i++
   }

   return data, nil
}