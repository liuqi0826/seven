package utils

import (
	"crypto/md5"
	"fmt"
	"io"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

type TokenManager struct {
	sync.Mutex

	tokens map[string]string
}

func (this *TokenManager) TokenManager() {
	this.tokens = make(map[string]string)
}
func (this *TokenManager) Create(id string) string {
	rs := rand.New(rand.NewSource(time.Now().UnixNano()))
	hash := md5.New()
	_, err := io.WriteString(hash, strconv.FormatInt(rs.Int63(), 10))
	if err == nil {
		token := fmt.Sprintf("%x", hash.Sum(nil))
		this.tokens[id] = token
		return token
	}
	return ""
}
func (this *TokenManager) Update(id string) string {
	return this.Create(id)
}
func (this *TokenManager) Delete(id string) string {
	if tk, ok := this.tokens[id]; ok {
		this.Lock()
		delete(this.tokens, id)
		this.Unlock()
		return tk
	}
	return ""
}
func (this *TokenManager) Check(id string, token string) bool {
	if tk, ok := this.tokens[id]; ok {
		if tk == token {
			return true
		}
	}
	return false
}
