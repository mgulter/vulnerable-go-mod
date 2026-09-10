package main

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/gogo/protobuf/proto"
	"github.com/ulikunitz/xz"
	"gopkg.in/yaml.v2"
)

func main() {
	_ = jwt.StandardClaims{}
	_ = proto.CompactTextString
	_, _ = xz.NewReader(nil)
	var config map[string]bool
	_ = yaml.Unmarshal([]byte("enabled: true"), &config)

	fmt.Println("vuln-go demo")
}
