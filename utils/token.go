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

type Token struct {
	account string
	token   string
}

type TokenManager struct {
	mutex     sync.Mutex
	tokenDict map[string]Token
}

func (this *TokenManager) TokenManager() {
	this.mutex = sync.Mutex{}
	this.tokenDict = make(map[string]Token)
}

func (this *TokenManager) GetNewToken(account string) string {
	rs := rand.New(rand.NewSource(time.Now().UnixNano()))
	hash := md5.New()
	_, err := io.WriteString(hash, strconv.FormatInt(rs.Int63(), 10))
	if err == nil {
		tkn := Token{}
		tkn.account = account
		tkn.token = fmt.Sprintf("%x", hash.Sum(nil))
		this.mutex.Lock()
		this.tokenDict[account] = tkn
		this.mutex.Unlock()
		return tkn.token
	}
	return ""
}

func (this *TokenManager) CheckToken(account string, tk string) bool {
	if v, b := this.tokenDict[account]; b {
		if v.token == tk {
			this.mutex.Lock()
			delete(this.tokenDict, account)
			this.mutex.Unlock()
			return true
		}
	}
	return false
}
