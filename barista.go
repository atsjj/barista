package main

import (
  "github.com/atsjj/pumpkin-spice-latte/barista"
  "github.com/atsjj/pumpkin-spice-latte/sap"
  "github.com/golang/protobuf/proto"
  "log"
  pstruct "github.com/golang/protobuf/ptypes/struct"
)

// func url_for(messageName string) (string) {
//   return "type.googleapis.com/" + messageName
// }

// func to_s(s string) (*any.Any) {
//   stringValue := &wrappers.StringValue { s }
//   typeUrl := url_for(proto.MessageName(stringValue))
//   value, err := proto.Marshal(stringValue)

//   if err != nil {
//     panic(err)
//   }

//   return &any.Any { typeUrl, value }
// }

func to_s(stringValue string) (*pstruct.Value) {
  return &pstruct.Value {
    &pstruct.Value_StringValue {
      stringValue,
    },
  }
}

func main() {
  a := &sap.Payload { &pstruct.Struct { make(map[string]*pstruct.Value) } }
  a.Model.Fields["test"] = to_s("This is a test")

  b := &sap.Payload {}

  c, _ := proto.Marshal(a)

  err := proto.Unmarshal(c, b)

  if err != nil {
    panic(err)
  }

  log.Println("a:")
  log.Println(a.Model.Fields["test"])

  log.Println("b:")
  log.Println(b.Model.Fields["test"])
}
