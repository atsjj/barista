package barista

import (
  pstruct "github.com/golang/protobuf/ptypes/struct"
)

func String(stringValue string) (*pstruct.Value) {
  return &pstruct.Value {
    &pstruct.Value_StringValue {
      stringValue,
    },
  }
}
