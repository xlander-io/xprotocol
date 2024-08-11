package triedb

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/xlander-io/cache"
	"github.com/xlander-io/hash"
	"github.com/xlander-io/kv"
	"github.com/xlander-io/kv_leveldb"
	"github.com/xlander-io/triedb"
)

/*


1. token

path:
/token-0.0.1/{token_addr}

val:
{
    creator_addr,
    block_height,
    status,                 // normal , stopped , expiring , expired
    expire_block_height,    // blockheight to be expired
    total_minted,
    total_burned,
    mint_max_limit,         //    upper bound is uint64
}

TX (1,0,{mint_max_limit}) , {0x1234,signature}  //

*/

type Token_0_0_1 struct {
	Token_name          string `json:"token_name"`
	Creator_addr_hex    string `json:"creator_addr_hex"`
	Block_height        int    `json:"block_height"`
	Status              string `json:"status"`
	Expire_block_height int    `json:"expire_block_height"`
	Total_minted        int    `json:"total_minted"`
}

func TestToken(t *testing.T) {

	kvdb, err := kv_leveldb.NewDB("./test2.db")
	if err != nil {
		panic(err)
	}

	c, err := cache.New(nil)
	if err != nil {
		panic(err)
	}
	tdb, err := triedb.NewTrieDB(kvdb, c, nil)

	if err != nil {
		panic(err)
	}

	token_name := "token-0.0.1"
	//

	token_addr := hash.NewHashFromString("0xd72cd83afba0dcfeff95d463adcb2b8def6aa623")
	token_0_0_1 := &Token_0_0_1{
		Token_name:          "coin1",
		Creator_addr_hex:    hash.NewHashFromString("0xDB65702A9b26f8a643a31a4c84b9392589e03D7c").Hex(),
		Block_height:        12,
		Status:              "normal",
		Expire_block_height: 100,
		Total_minted:        1000,
	}
	//

	token_addr2 := hash.NewHashFromString("0x4838b106fce9647bdf1e7877bf73ce8b0bad5f97")
	token_0_0_1_2 := &Token_0_0_1{
		Token_name:          "coin2",
		Creator_addr_hex:    hash.NewHashFromString("0xDB65702A9b26f8a643a31a4c84b9392589e03D7c").Hex(),
		Block_height:        12,
		Status:              "normal",
		Expire_block_height: 50,
		Total_minted:        2000,
	}

	//
	account := "0xdac17f958d2ee523a2206206994597c13d831ec7"

	token_0_0_1_json, _ := json.Marshal(token_0_0_1)

	fmt.Println(tdb.Put(triedb.Path([]byte(token_name), []byte(token_addr.Bytes())), token_0_0_1_json, false))

	fmt.Println(tdb.Put(triedb.Path([]byte("account"), []byte(account), []byte(token_name), []byte(token_addr.Hex())), []byte("10bytes"), false))
	//

	token_0_0_2_json, _ := json.Marshal(token_0_0_1_2)

	fmt.Println(tdb.Put(triedb.Path([]byte(token_name), []byte(token_addr2.Bytes())), token_0_0_2_json, false))

	fmt.Println(tdb.Put(triedb.Path([]byte("account"), []byte(account), []byte(token_name), []byte(token_addr2.Hex())), []byte("10bytes"), false))

	///////////
	result, _, _ := tdb.Get(triedb.Path([]byte(token_name), []byte(token_addr.Bytes())))

	result_token_0_0_1 := &Token_0_0_1{}

	json.Unmarshal(result, result_token_0_0_1)

	fmt.Println(result_token_0_0_1)

	//make a change

	result_token_0_0_1.Expire_block_height = 10000
	result_token_0_0_1.Total_minted += 1000
	token_0_0_1_json2, _ := json.Marshal(result_token_0_0_1)

	fmt.Println(tdb.Put(triedb.Path([]byte(token_name), []byte(token_addr.Bytes())), token_0_0_1_json2, false))

	result2, _, _ := tdb.Get(triedb.Path([]byte(token_name), []byte(token_addr.Bytes())))

	result2_token_0_0_1 := &Token_0_0_1{}

	json.Unmarshal(result2, result2_token_0_0_1)

	fmt.Println(result2_token_0_0_1)

	//

	root_hash, updated, deleted, _ := tdb.Commit()

	fmt.Println("update2 len:", len(updated))

	b := kv.NewBatch()

	for key, val := range updated {
		fmt.Println("to update:", hash.NewHashFromBytes([]byte(key)).Hex())
		b.Put([]byte(key), val)
	}

	for key, val := range deleted {
		fmt.Println("to del:", val.Hex())
		b.Delete([]byte(key))
	}

	err = kvdb.WriteBatch(b, true)
	if err != nil {
		fmt.Println("del batch err", err)
	}

	kvdb.Close()

	kvdb2, err := kv_leveldb.NewDB("./test2.db")
	if err != nil {
		panic(err)
	}

	c2, err := cache.New(nil)
	if err != nil {
		panic(err)
	}
	tdb2, err := triedb.NewTrieDB(kvdb2, c2, &triedb.TrieDBConfig{
		Root_hash: root_hash,
		Read_only: true,
	})

	if err != nil {
		panic(err)
	}

	///iter
	iter, iter_err := tdb2.NewIterator(triedb.Path([]byte("account"), []byte(account), []byte(token_name)))

	fmt.Println("iter:", iter, " err:", iter_err)

	for {
		fmt.Println("/////////////////////////")
		fmt.Println(iter.FullPathFlatStr())
		fmt.Println("/////////////////////////")
		has_next, _ := iter.Next()
		if !has_next {
			break
		}
	}

}
