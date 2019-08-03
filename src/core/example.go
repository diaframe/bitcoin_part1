package core

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"github.com/boltdb/bolt"
	"log"
)

func (block *Block)Serialize()string{
	var result bytes.Buffer
	encoder :=gob.NewEncoder(&result)
	err :=encoder.Encode(block)
	if err!=nil{
		fmt.Println(err)
	}
	return result.String()
}

func main(){
	//1.打开数据库
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Update(func(tx *bolt.Tx) error {
		//2.通过Bucket()方法打开BlockBucket表
		b := tx.Bucket([]byte("BlockBucket"))
		//3.通过Put()方法往表里面存储数据
		if b != nil {
			err := b.Put([]byte("l"), []byte("Send $200 TO Fengyingcong"))
			err =  b.Put([]byte("ll"), []byte("Send $100 TO Bruce"))
			if err != nil {
				log.Panic("数据存储失败..")
			}
		}

		return nil
	})
	//更新失败
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("存储成功")
}