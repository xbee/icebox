package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"

	// _ "github.com/mattn/go-sqlite3"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
	"github.com/btcsuite/btcutil"

	"conseweb.com/go-bip44"
	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"
)

const (
	dbpath   = "db/foo.db"
	secretFn = "ss/secret.dat"
)

var (
	masterKey *bip32.Key
	db        *gorm.DB
)

type Addr struct {
	// gorm.Model
	ID         uint32 `gorm:"AUTO_INCREMENT, PRIMARY_KEY"` // set id to auto incrementable
	Type       uint32 `gorm:"not null"`                    // set type to unique and not null
	IsExternal uint32 `gorm:"default:0"`
	Index      uint32 `gorm:"not null"`
	Name       string
}

// 该表存储在设备侧，静态配置数据，还可以存储币种的加密配置信息
type Product struct {
	// gorm.Model
	ID     uint32 `gorm:"AUTO_INCREMENT, PRIMARY_KEY"` // set id to auto incrementable
	Type   uint32 `gorm:"unique;not null"`             // 对应bip32的coin_type
	Symbol string `gorm:"unique;not null"`             // 对应币种的代号, 如比特币是: btc
	Name   string `gorm:"unique;not null"`             // 对应币种的全称, 如比特币是: bitcoin
	Path   string `gorm:"not null"`                    // 对应币种的account derivation path
}

func InitDB(filepath string) *gorm.DB {
	db, err := gorm.Open("sqlite3", filepath)
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Product{})
	db.AutoMigrate(&Addr{})

	// Create
	db.Create(&Product{Type: 0, Symbol: "btc", Name: "bitcoin", Path: "m/44'/0'/0'/"})
	db.Create(&Product{Type: 1, Symbol: "test", Name: "testnet", Path: "m/44'/1'/0'/"})
	db.Create(&Product{Type: 2, Symbol: "ltc", Name: "litecoin", Path: "m/44'/2'/0'/"})
	db.Create(&Product{Type: 3, Symbol: "doge", Name: "dogecoin", Path: "m/44'/3'/0'/"})
	db.Create(&Product{Type: 5, Symbol: "dsh", Name: "dash", Path: "m/44'/5'/0'/"})
	db.Create(&Product{Type: 9, Symbol: "xcp", Name: "counterparty", Path: "m/44'/9'/0'/"})
	db.Create(&Product{Type: 60, Symbol: "eth", Name: "ethereum", Path: "m/44'/60'/0'/"})
	db.Create(&Product{Type: 61, Symbol: "etc", Name: "ethereum classic", Path: "m/44'/61'/0'/"})

	return db
}

type Transaction struct {
	TxId               string `json:"txid"`
	SourceAddress      string `json:"source_address"`
	DestinationAddress string `json:"destination_address"`
	Amount             int64  `json:"amount"`
	UnsignedTx         string `json:"unsignedtx"`
	SignedTx           string `json:"signedtx"`
}

func CreateTransaction(secret string, destination string, amount int64, txHash string) (Transaction, error) {
	var transaction Transaction
	wif, err := btcutil.DecodeWIF(secret)
	if err != nil {
		return Transaction{}, err
	}
	addresspubkey, _ := btcutil.NewAddressPubKey(wif.PrivKey.PubKey().SerializeUncompressed(), &chaincfg.MainNetParams)
	sourceTx := wire.NewMsgTx(wire.TxVersion)
	sourceUtxoHash, _ := chainhash.NewHashFromStr(txHash)
	sourceUtxo := wire.NewOutPoint(sourceUtxoHash, 0)
	sourceTxIn := wire.NewTxIn(sourceUtxo, nil, nil)
	destinationAddress, err := btcutil.DecodeAddress(destination, &chaincfg.MainNetParams)
	sourceAddress, err := btcutil.DecodeAddress(addresspubkey.EncodeAddress(), &chaincfg.MainNetParams)
	if err != nil {
		return Transaction{}, err
	}
	destinationPkScript, _ := txscript.PayToAddrScript(destinationAddress)
	sourcePkScript, _ := txscript.PayToAddrScript(sourceAddress)
	sourceTxOut := wire.NewTxOut(amount, sourcePkScript)
	sourceTx.AddTxIn(sourceTxIn)
	sourceTx.AddTxOut(sourceTxOut)
	sourceTxHash := sourceTx.TxHash()
	redeemTx := wire.NewMsgTx(wire.TxVersion)
	prevOut := wire.NewOutPoint(&sourceTxHash, 0)
	redeemTxIn := wire.NewTxIn(prevOut, nil, nil)
	redeemTx.AddTxIn(redeemTxIn)
	redeemTxOut := wire.NewTxOut(amount, destinationPkScript)
	redeemTx.AddTxOut(redeemTxOut)
	sigScript, err := txscript.SignatureScript(redeemTx, 0, sourceTx.TxOut[0].PkScript, txscript.SigHashAll, wif.PrivKey, false)
	if err != nil {
		return Transaction{}, err
	}
	redeemTx.TxIn[0].SignatureScript = sigScript
	flags := txscript.StandardVerifyFlags
	vm, err := txscript.NewEngine(sourceTx.TxOut[0].PkScript, redeemTx, 0, flags, nil, nil, amount)
	if err != nil {
		return Transaction{}, err
	}
	if err := vm.Execute(); err != nil {
		return Transaction{}, err
	}
	var unsignedTx bytes.Buffer
	var signedTx bytes.Buffer
	sourceTx.Serialize(&unsignedTx)
	redeemTx.Serialize(&signedTx)
	transaction.TxId = sourceTxHash.String()
	transaction.UnsignedTx = hex.EncodeToString(unsignedTx.Bytes())
	transaction.Amount = amount
	transaction.SignedTx = hex.EncodeToString(signedTx.Bytes())
	transaction.SourceAddress = sourceAddress.EncodeAddress()
	transaction.DestinationAddress = destinationAddress.EncodeAddress()
	return transaction, nil
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func InitDevice(path string) {
	// if secretFn does not exists then init device
	// detect if file exists
	var _, err = os.Stat(path)

	// create file if not exists
	if !os.IsNotExist(err) {
		// read file
		var key *bip32.Key
		key, err = loadMasterPrivKey(path)
		publicKey := key.PublicKey()

		// Display mnemonic and keys
		fmt.Println("Loading key from secure storage ...")
		fmt.Println("Master private key: ", key)
		fmt.Println("Master public key: ", publicKey)
		return
	}
	// Generate a mnemonic for memorization or user-friendly seeds
	entropy, _ := bip39.NewEntropy(256)
	mnemonic, _ := bip39.NewMnemonic(entropy)

	// Generate a Bip32 HD wallet for the mnemonic and a user supplied password
	seed := bip39.NewSeed(mnemonic, "Secret")

	masterKey, _ := bip32.NewMasterKey(seed)
	publicKey := masterKey.PublicKey()

	// Display mnemonic and keys
	fmt.Println("Mnemonic: ", mnemonic)
	fmt.Println("Master private key: ", masterKey)
	fmt.Println("Master public key: ", publicKey)

	secret := masterKey.String()
	err = ioutil.WriteFile(secretFn, []byte(secret), 0644)
	check(err)
}

func deleteDbFile(path string) {
	// delete file
	var err = os.Remove(path)
	if isError(err) {
		return
	}

	fmt.Println("==> done deleting file")
}

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}

func loadMasterPrivKey(fn string) (key *bip32.Key, err error) {
	var data []byte
	data, err = ioutil.ReadFile(fn)
	check(err)

	key, err = bip32.B58Deserialize(string(data))
	if err != nil {
		return nil, err
	}
	return key, nil
}

// 创建一个新的
func CreateAddr(coin, chain, index uint32, name string) (*bip32.Key, error) {
	db.Create(&Addr{Type: coin, Index: index, Name: name})

	mkey, err := loadMasterPrivKey(secretFn)
	if mkey != nil {
		var key *bip32.Key
		key, err = bip44.NewKeyFromMasterKey(mkey, coin, 0, chain, index)
		if err != nil {
			return nil, err
		}
		return key, nil
	}
	return nil, err
}

func main() {
	InitDevice(secretFn)

	// _, err := CreateTransaction("5HusYj2b2x4nroApgfvaSfKYZhRbKFH41bVyPooymbC6KfgSXdD", "1KKKK6N21XKo48zWKuQKXdvSsCf95ibHFa", 91234, "81b4c832d70cb56ff957589752eb4125a4cab78a25a8fc52d6a09e5bd4404d48")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// data, _ := json.Marshal(transaction)
	// fmt.Println(string(data))

	db = InitDB(dbpath)
	defer db.Close()
	defer deleteDbFile(dbpath)

	// Read
	var product Product
	db.First(&product, "type = ?", 0) // find product with code l121

	// var key *bip32.Key
	var idx uint32
	idx = 32
	key, err := CreateAddr(product.Type, 0, idx, "default")
	if err != nil {
		fmt.Errorf("%s", err)
	}
	pk := key.PublicKey()
	CreateAddr(product.Type, 0, idx+1, "bussiness")

	var addr Addr
	db.First(&addr, "type = ?", 0)
	// db.Create(&Addr{Type: product.Type, Index: idx, Name: "default"})

	fmt.Printf("%+v\n", addr)
	fmt.Printf("%+v\n", product)
	fmt.Printf("%s\n", pk.String())
}
