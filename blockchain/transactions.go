package blockchain

import (
	"errors"
	"ralo/utils"
	"ralo/wallet"
	"time"
)

const (
	minerReward int = 50
)

type mempool struct {
	Txs []*Tx
}

var Mempool *mempool = &mempool{}

var ErrorNoMoney = errors.New("Not enough money")
var ErrorNotValid = errors.New("Tx Invalid")

type Tx struct {
	Id        string   `json:"id"`
	Timestamp int      `json:"timestamp"`
	TxIns     []*TxIn  `json:"txIns"`
	TxOuts    []*TxOut `json:"txOuts"`
}

type TxIn struct {
	TxId      string `json:"txid"`
	Index     int    `json:"index"`
	Signature string `json:"signature"`
}

type UTxOut struct {
	TxID   string
	Index  int
	Amount int
}

type TxOut struct {
	Address string `json:"address"`
	Amount  int    `json:"amount"`
}

func (t *Tx) getId() {
	t.Id = utils.Hash(t)
}

func (t *Tx) sign() {
	for _, txIn := range t.TxIns {
		txIn.Signature = wallet.Sign(t.Id, wallet.Wallet())
	}
}

func validate(tx *Tx) bool {
	vaild := true
	for _, txIn := range tx.TxIns {
		prevTx := FindTx(Blockchain(), txIn.TxId)
		if prevTx == nil {
			vaild = false
			break
		}
		address := prevTx.TxOuts[txIn.Index].Address
		vaild = wallet.Verify(txIn.Signature, tx.Id, address)
		if vaild == false {
			break
		}
	}

	return vaild
}

func isOnMempool(uTxOut *UTxOut) bool {
	exist := false
Outer:
	for _, tx := range Mempool.Txs {
		for _, input := range tx.TxIns {
			if input.TxId == uTxOut.TxID && input.Index == uTxOut.Index {
				exist = true
				break Outer
			}

		}
	}
	return exist
}

func makeCoinbaseTx(address string) (*Tx, error) {
	txIns := []*TxIn{
		{"", -1, "Coinbase"},
	}
	txOuts := []*TxOut{
		{address, minerReward},
	}
	tx := Tx{
		Id:        "",
		Timestamp: int(time.Now().Unix()),
		TxIns:     txIns,
		TxOuts:    txOuts,
	}
	tx.getId()
	tx.sign()
	valid := validate(&tx)
	if !valid {
		return nil, ErrorNotValid
	}
	return &tx, nil
}

func makeTx(from, to string, amount int) (*Tx, error) {
	if BalanceByAddress(from, Blockchain()) < amount {
		return nil, ErrorNoMoney

	}
	var txOuts []*TxOut
	var txIns []*TxIn
	total := 0
	uTxOuts := UTxOutsByAddress(from, Blockchain())
	for _, UTxOut := range uTxOuts {
		if total >= amount {
			break
		}
		txIn := &TxIn{UTxOut.TxID, UTxOut.Index, from}
		txIns = append(txIns, txIn)
		total += UTxOut.Amount
	}

	if change := total - amount; change != 0 {
		changeTxOut := &TxOut{from, change}
		txOuts = append(txOuts, changeTxOut)
	}
	txOut := &TxOut{to, amount}
	txOuts = append(txOuts, txOut)
	tx := &Tx{
		Id:        "",
		Timestamp: int(time.Now().Unix()),
		TxIns:     txIns,
		TxOuts:    txOuts,
	}
	tx.getId()
	return tx, nil

}

func (m *mempool) AddTx(to string, amount int) error {
	tx, err := makeTx(wallet.Wallet().Address, to, amount)
	if err != nil {
		return err
	}
	m.Txs = append(m.Txs, tx)
	return nil
}

func (m *mempool) txToConfirm() []*Tx {
	coinbase := makeCoinbaseTx(wallet.Wallet().Address)
	txs := m.Txs
	txs = append(txs, coinbase)
	m.Txs = nil
	return txs
}
