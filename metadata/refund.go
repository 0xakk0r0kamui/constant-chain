package metadata

import (
	"github.com/constant-money/constant-chain/common"
	"github.com/constant-money/constant-chain/database"
)

type Refund struct {
	MetadataBase
	SmallTxID common.Hash
}

func NewRefund(smallTxID common.Hash, metaType int) *Refund {
	metadataBase := MetadataBase{
		Type: metaType,
	}
	return &Refund{
		SmallTxID:    smallTxID,
		MetadataBase: metadataBase,
	}
}

func (rf *Refund) CheckTransactionFee(tr Transaction, minFee uint64) bool {
	// no need to have fee for this tx
	return true
}

func (rf *Refund) ValidateTxWithBlockChain(txr Transaction, bcr BlockchainRetriever, shardID byte, db database.DatabaseInterface) (bool, error) {
	// no need to validate tx with blockchain, just need to validate with requeste tx (via SmallTxID) in current block
	return false, nil
}

func (rf *Refund) ValidateSanityData(bcr BlockchainRetriever, txr Transaction) (bool, bool, error) {
	return false, true, nil
}

func (rf *Refund) ValidateMetadataByItself() bool {
	// The validation just need to check at tx level, so returning true here
	return true
}

func (rf *Refund) Hash() *common.Hash {
	record := rf.SmallTxID.String()
	record += rf.MetadataBase.Hash().String()

	// final hash
	hash := common.DoubleHashH([]byte(record))
	return &hash
}

func (rf *Refund) CalculateSize() uint64 {
	return calculateSize(rf)
}
