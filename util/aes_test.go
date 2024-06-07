package util

import (
	"log"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAes(t *testing.T) {
	origData := []byte("Hello World") // 待加密的数据
	key := []byte("ABCDEFGHIJKLMNOP") // 加密的密钥
	log.Println("原文：", string(origData))

	Convey("TestAes", t, func() {
		Convey("------------------ CBC模式 --------------------", func() {
			encrypted, err := AesEncryptCBC(origData, key)
			So(err, ShouldBeNil)
			// log.Println("密文(hex): ", hex.EncodeToString(encrypted))
			// log.Println("密文(base64): ", base64.StdEncoding.EncodeToString(encrypted))
			decrypted, err := AesDecryptCBC(encrypted, key)
			So(err, ShouldBeNil)
			So(decrypted, ShouldResemble, origData)
		})
		Convey("------------------ ECB模式 --------------------", func() {
			encrypted, err := AesEncryptECB(origData, key)
			So(err, ShouldBeNil)
			decrypted, err := AesDecryptECB(encrypted, key)
			So(err, ShouldBeNil)
			So(decrypted, ShouldResemble, origData)
		})
		Convey("------------------ CFB模式 --------------------", func() {
			encrypted, err := AesEncryptCFB(origData, key)
			So(err, ShouldBeNil)
			decrypted, err := AesDecryptCFB(encrypted, key)
			So(err, ShouldBeNil)
			So(decrypted, ShouldResemble, origData)
		})
	})
}
