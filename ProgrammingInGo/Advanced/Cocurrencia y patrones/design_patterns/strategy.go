package main

import (
    "fmt"
)

type PassordProtector struct {
    userName string
    passwordName string
    hashAlgorithm HashAlgorithm
}

type HashAlgorithm interface {
    Hash(p *PassordProtector)
}

func NewPassordProtector(userName string, passwordName string, hashAlgorithm HashAlgorithm) (*PassordProtector) {
    return &PassordProtector{
        userName: userName,
        passwordName: passwordName,
        hashAlgorithm: hashAlgorithm,
    }
}

func (p *PassordProtector) SetHashAlgorithm(hashAlgorithm HashAlgorithm) {
    p.hashAlgorithm = hashAlgorithm
}

func (p *PassordProtector) Hash(){
    p.hashAlgorithm.Hash(p)
}

type SHA struct {}

func (SHA) Hash(p *PassordProtector) {
    fmt.Printf("Hashing using SHA for %s\n", p.passwordName)
}

type MD5 struct {}

func (MD5) Hash(p *PassordProtector) {
    fmt.Printf("Hashing using MD5 for %s\n", p.passwordName)
}

func main() {
    sha := SHA{}
    md5 := MD5{}
    passordProtector := NewPassordProtector("Eliot", "gmail password", sha)
    passordProtector.Hash()
    passordProtector.SetHashAlgorithm(md5)
    passordProtector.Hash()
}
