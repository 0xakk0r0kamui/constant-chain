package metadata

import (
	"github.com/constant-money/constant-chain/common"
	"github.com/constant-money/constant-chain/privacy"
)

type VoteBoardMetadata struct {
	CandidatePaymentAddress privacy.PaymentAddress
	BoardIndex              uint32
}

func NewVoteBoardMetadata(candidatePaymentAddress privacy.PaymentAddress, boardIndex uint32) *VoteBoardMetadata {
	return &VoteBoardMetadata{CandidatePaymentAddress: candidatePaymentAddress, BoardIndex: boardIndex}
}

func (voteBoardMetadata *VoteBoardMetadata) GetBytes() []byte {
	record := string(voteBoardMetadata.CandidatePaymentAddress.Bytes())
	record += string(common.Uint32ToBytes(voteBoardMetadata.BoardIndex))
	return []byte(record)
}

func (voteGOVBoardMetadata *VoteGOVBoardMetadata) CalculateSize() uint64 {
	return calculateSize(voteGOVBoardMetadata)
}
