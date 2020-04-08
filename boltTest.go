package main

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
)

func main() {
	var (
		db  *bolt.DB
		err error
	)
	//1.打开数据库
	if db, err = bolt.Open("testDB", 0600, nil); err != nil {
		log.Panic("打开数据库失败")
	}
	defer db.Close()
	//2.将要操作数据库
	db.Update(func(tx *bolt.Tx) error {
		//2.找到抽屉bucket(如果没有 就创建)
		bucket := tx.Bucket([]byte("b1"))
		if bucket == nil {
			//没有抽屉 需要创建
			if bucket, err = tx.CreateBucket([]byte("b1")); err != nil {
				log.Panic("创建失败")
			}
		}
		//3.写数据
		bucket.Put([]byte("1111"), []byte("222222"))
		bucket.Put([]byte("22222"), []byte("22223333322"))
		return nil
	})

	//4.读数据
	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("b1"))
		if bucket == nil {
			log.Panic("bucket无法找到")
		}
		//直接读取数据
		v1 := bucket.Get([]byte("1111"))
		v2 := bucket.Get([]byte("22222"))

		fmt.Printf("%s\n", v1)
		fmt.Printf("%s\n", v2)

		return nil
	})
}
