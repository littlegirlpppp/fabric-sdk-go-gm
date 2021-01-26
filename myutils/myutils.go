package myutils

import (
	"encoding/pem"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/fabric-protos-go/common"
	"github.com/hyperledger/fabric-protos-go/ledger/rwset"
	"github.com/hyperledger/fabric-protos-go/ledger/rwset/kvrwset"
	"github.com/littlegirlpppp/fabric-sdk-go-gm/internal/github.com/hyperledger/fabric/protoutil"
	gmx509 "github.com/littlegirlpppp/fabric-sdk-go-gm/third_party/github.com/tjfoc/x509"
	"strings"
	"time"
)

type Read struct {
	Key      string `json:"key"`
	BlockNum uint64 `json:"blocknum"`
	TxNum    uint64 `json:"txnum"`
}

type Write struct {
	Key      string `json:"key"`
	IsDelete bool   `json:"isdelete"`
	Value    string `json:"value"`
}

//TransactionInfo 解析后的交易信息
type TransactionInfo struct {
	BlockHeight      uint64 `json:"blockheight"`      //区块高度ttt
	ChannelID        string `json:"channelid"`        //通道名称
	CreateTime       string `json:"createtime"`       //交易创建时间
	ChaincodeID      string `json:"chaincodeid"`      //交易调用链码ID
	ChaincodeVersion string `json:"chaincodeversion"` //交易调用链码版本
	//Nonce            []byte  `json:"channelid"` //随机数
	Mspid string `json:"mspid"` //交易发起者MSPID
	Name  string `json:"name"`  //交易发起者名称
	//OUTypes          string   //交易发起者OU分组
	Args []string `json:"args"` //输入参数
	//Type             int32    `json:"type"`//交易类型
	TxID     string   `json:"txid"` //交易ID
	ReadSet  []*Read  `json:"readset"` //读集
	WriteSet []*Write `json:"writeset"` //写集
}

// 获取当前区块哈数
func UnmarshalBlockHash(b *common.BlockHeader) string  {
	previousBlockHash := protoutil.BlockHeaderHash(b)
	return fmt.Sprintf("%x", previousBlockHash)
}

// 解析区块
func UnmarshalBlock(payloadRaw []byte,height uint64) (*TransactionInfo, error) {
	// 解析block
	block1, err := protoutil.GetEnvelopeFromBlock(payloadRaw)
	if err != nil {
		return nil, err
	}
	payload, err := protoutil.UnmarshalPayload(block1.Payload)
	if err != nil {
		return nil, err
	}

	return setTransaction(payload,height)
}

// UnmarshalTransaction 从某个交易的payload来解析它
func UnmarshalTransaction(payloadRaw []byte) (*TransactionInfo, error) {
	//解析成payload
	payload, err := protoutil.UnmarshalPayload(payloadRaw)
	if err != nil {
		return nil, err
	}

	return setTransaction(payload,0)
}

func setTransaction(payload*common.Payload,height uint64) (*TransactionInfo,error) {
	result := &TransactionInfo{}
	result.BlockHeight=height
	//解析成ChannelHeader（包含通道ID、交易ID及交易创建时间等)
	channelHeader, err := protoutil.UnmarshalChannelHeader(payload.Header.ChannelHeader)
	if err != nil {
		return nil, err
	}
	result.ChannelID = channelHeader.ChannelId
	//解析成SignatureHeader（包含创建者和nonce)
	signHeader, err := protoutil.UnmarshalSignatureHeader(payload.Header.SignatureHeader)
	if err != nil {
		return nil, err
	}
	//解析成SerializedIdentity（包含证书和MSPID）
	identity, err := protoutil.UnmarshalSerializedIdentity(signHeader.GetCreator())
	if err != nil {
		return nil, err
	}
	//下面为解析证书
	block, _ := pem.Decode(identity.GetIdBytes())
	if block == nil {
		return nil, fmt.Errorf("identity could not be decoded from credential")
	}
	cert, err := gmx509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse certificate: %s", err)
	}

	//解析用户名和OU分组
	uname := cert.Subject.CommonName
	//outypes := cert.Subject.OrganizationalUnit
	//解析成transaction
	tx, err := protoutil.UnmarshalTransaction(payload.Data)
	if err != nil {
		return nil, err
	}
	//进一步从transaction中解析成ChaincodeActionPayload
	chaincodeActionPayload, err := protoutil.UnmarshalChaincodeActionPayload(tx.Actions[0].Payload)
	if err != nil {
		return nil, err
	}
	//进一步解析成proposalPayload
	proposalPayload, err := protoutil.UnmarshalChaincodeProposalPayload(chaincodeActionPayload.ChaincodeProposalPayload)
	if err != nil {
		return nil, err
	}

	// 获取链码版本
	proposalPayload1, err := protoutil.UnmarshalProposalResponsePayload(chaincodeActionPayload.Action.ProposalResponsePayload)
	chaincode, err := protoutil.UnmarshalChaincodeAction(proposalPayload1.Extension)
	if err != nil {
		return nil, err
	}

	//获取读写集
	pubSimulationResults := &rwset.TxReadWriteSet{}
	err = proto.Unmarshal(chaincode.Results, pubSimulationResults)
	if err != nil {
		return nil, err
	}

	reads := make([]*Read, 0)
	writes := make([]*Write, 0)

	for _, rs := range pubSimulationResults.NsRwset {
		if rs.Namespace == "_lifecycle" {
			continue
		}
		RWSet := &kvrwset.KVRWSet{}
		err = proto.Unmarshal(rs.Rwset, RWSet)
		if err != nil {
			return nil, err
		}

		for _, r := range RWSet.Reads {
			_read := new(Read)
			if strings.Index(r.Key,"initialized")>0{
				continue
			}
			_read.Key = r.Key
			if r.Version==nil{
				_read.BlockNum=0
				_read.TxNum =0
			}else {
				_read.BlockNum = r.Version.BlockNum
				_read.TxNum = r.Version.TxNum
			}
			reads = append(reads, _read)
		}

		for _, w := range RWSet.Writes {
			_write := new(Write)
			_write.Key = w.Key
			_write.IsDelete = w.IsDelete
			_write.Value = string(w.Value)
			writes = append(writes, _write)
		}

	}
	result.ReadSet = reads
	result.WriteSet = writes
	//pubSimulationResults.NsRwset[0].
	//得到交易调用的链码信息
	chaincodeInvocationSpec, err := protoutil.UnmarshalChaincodeInvocationSpec(proposalPayload.Input)
	if err != nil {
		return nil, err
	}
	//得到调用的链码的ID，版本和PATH（这里PATH省略了）
	//得到输入参数
	var args []string
	if chaincodeInvocationSpec.ChaincodeSpec!=nil{
		result.ChaincodeID = chaincodeInvocationSpec.ChaincodeSpec.ChaincodeId.Name
		result.ChaincodeVersion = chaincode.ChaincodeId.Version
		for _, v := range chaincodeInvocationSpec.ChaincodeSpec.Input.Args {
			args = append(args, string(v))
		}
	}
	result.Args = args
	//result.Nonce = signHeader.GetNonce()
	//result.Type = channelHeader.GetType()
	result.TxID = channelHeader.GetTxId()
	result.Mspid = identity.GetMspid()
	result.Name = uname
	//result.OUTypes = outypes[0]
	result.CreateTime = time.Unix(channelHeader.Timestamp.Seconds, 0).Format("2006-01-02 15:04:05")

	return result, err
}