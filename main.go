package main

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/gogo/protobuf/proto"
	"github.com/ulikunitz/xz"
)

func main() {
	_ = jwt.StandardClaims{}
	_ = proto.CompactTextString
	_, _ = xz.NewReader(nil)

	fmt.Println("vuln-go demo")
}
