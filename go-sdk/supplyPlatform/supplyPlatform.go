// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package supplyPlatform

import (
	"math/big"
	"strings"

	"github.com/FISCO-BCOS/go-sdk/abi"
	"github.com/FISCO-BCOS/go-sdk/abi/bind"
	"github.com/FISCO-BCOS/go-sdk/core/types"
	"github.com/FISCO-BCOS/go-sdk/event"
	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// SupplyPlatformABI is the input ABI used to generate the binding from.
const SupplyPlatformABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"From\",\"type\":\"string\"},{\"name\":\"To\",\"type\":\"string\"}],\"name\":\"TransferCount\",\"outputs\":[{\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"Name\",\"type\":\"string\"},{\"name\":\"Money\",\"type\":\"int256\"},{\"name\":\"Type\",\"type\":\"int256\"}],\"name\":\"Register\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"s1\",\"type\":\"string\"},{\"name\":\"s2\",\"type\":\"string\"}],\"name\":\"isEqual\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"From\",\"type\":\"string\"}],\"name\":\"RepayCount\",\"outputs\":[{\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"From\",\"type\":\"string\"}],\"name\":\"Repay\",\"outputs\":[{\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"Name\",\"type\":\"string\"}],\"name\":\"GetCompany\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"int256\"},{\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"From\",\"type\":\"string\"},{\"name\":\"To\",\"type\":\"string\"},{\"name\":\"Money\",\"type\":\"int256\"}],\"name\":\"Transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"From\",\"type\":\"string\"},{\"name\":\"To\",\"type\":\"string\"}],\"name\":\"Raise\",\"outputs\":[{\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"id\",\"type\":\"int256\"}],\"name\":\"GetBill\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"int256\"},{\"name\":\"\",\"type\":\"int256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"From\",\"type\":\"string\"},{\"name\":\"To\",\"type\":\"string\"},{\"name\":\"Money\",\"type\":\"int256\"}],\"name\":\"NewBill\",\"outputs\":[{\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// SupplyPlatformBin is the compiled bytecode used for deploying new contracts.
var SupplyPlatformBin = "0x608060405234801561001057600080fd5b5060008081905550600060018190555061215c8061002f6000396000f3006080604052600436106100a4576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680630c86c6b8146100a9578063146e52091461016c578063465c4105146101e95780634c3eb6ef146102b057806354b720d41461032d5780635aa5af05146103aa578063881696021461049a57806390bee2ee14610567578063acb1dfb21461062a578063f656f55014610758575b600080fd5b3480156100b557600080fd5b50610156600480360381019080803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050610825565b6040518082815260200191505060405180910390f35b34801561017857600080fd5b506101e7600480360381019080803590602001908201803590602001908080601f01602080910402602001604051908101604052809392919081815260200183838082843782019150505050505091929192908035906020019092919080359060200190929190505050610b5e565b005b3480156101f557600080fd5b50610296600480360381019080803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050610c2f565b604051808215151515815260200191505060405180910390f35b3480156102bc57600080fd5b50610317600480360381019080803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050610d73565b6040518082815260200191505060405180910390f35b34801561033957600080fd5b50610394600480360381019080803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050610f5f565b6040518082815260200191505060405180910390f35b3480156103b657600080fd5b50610411600480360381019080803590602001908201803590602001908080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050919291929050505061121f565b6040518080602001848152602001838152602001828103825285818151815260200191508051906020019080838360005b8381101561045d578082015181840152602081019050610442565b50505050905090810190601f16801561048a5780820380516001836020036101000a031916815260200191505b5094505050505060405180910390f35b3480156104a657600080fd5b50610551600480360381019080803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290803590602001908201803590602001908080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050919291929080359060200190929190505050611419565b6040518082815260200191505060405180910390f35b34801561057357600080fd5b50610614600480360381019080803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050611892565b6040518082815260200191505060405180910390f35b34801561063657600080fd5b5061065560048036038101908080359060200190929190505050611c35565b604051808060200180602001878152602001868152602001858152602001848152602001838103835289818151815260200191508051906020019080838360005b838110156106b1578082015181840152602081019050610696565b50505050905090810190601f1680156106de5780820380516001836020036101000a031916815260200191505b50838103825288818151815260200191508051906020019080838360005b838110156107175780820151818401526020810190506106fc565b50505050905090810190601f1680156107445780820380516001836020036101000a031916815260200191505b509850505050505050505060405180910390f35b34801561076457600080fd5b5061080f600480360381019080803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290803590602001908201803590602001908080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050919291929080359060200190929190505050611e13565b6040518082815260200191505060405180910390f35b600080600080915060016002866040518082805190602001908083835b6020831015156108675780518252602082019150602081019050602083039250610842565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060020154148015610919575060016002856040518082805190602001908083835b6020831015156108e057805182526020820191506020810190506020830392506108bb565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060020154145b801561092c575061092a8585610c2f565b155b15610b3257600090505b600154811215610b2a576109f8600360008381526020019081526020016000206001018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156109ed5780601f106109c2576101008083540402835291602001916109ed565b820191906000526020600020905b8154815290600101906020018083116109d057829003601f168201915b505050505086610c2f565b8015610ab95750610ab7600360008381526020019081526020016000206000018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610aac5780601f10610a8157610100808354040283529160200191610aac565b820191906000526020600020905b815481529060010190602001808311610a8f57829003601f168201915b505050505085610c2f565b155b8015610afc5750600160036000838152602001908152602001600020600201541480610afb575060026003600083815260200190815260200160002060020154145b5b15610b1d576003600082815260200190815260200160002060030154820191505b8080600101915050610936565b819250610b56565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff92505b505092915050565b606060405190810160405280848152602001838152602001828152506002846040518082805190602001908083835b602083101515610bb25780518252602082019150602081019050602083039250610b8d565b6001836020036101000a03801982511681845116808217855250505050505090500191505090815260200160405180910390206000820151816000019080519060200190610c0192919061200b565b5060208201518160010155604082015181600201559050506000808154809291906001019190505550505050565b60008082518451141515610c465760009150610d6c565b600090505b8351811015610d67578281815181101515610c6257fe5b9060200101517f010000000000000000000000000000000000000000000000000000000000000090047f0100000000000000000000000000000000000000000000000000000000000000027effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168482815181101515610cdd57fe5b9060200101517f010000000000000000000000000000000000000000000000000000000000000090047f0100000000000000000000000000000000000000000000000000000000000000027effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916141515610d5a5760009150610d6c565b8080600101915050610c4b565b600191505b5092915050565b600080600080915060016002856040518082805190602001908083835b602083101515610db55780518252602082019150602081019050602083039250610d90565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040518091039020600201541415610f3457600090505b600154811215610f2c57610eb9600360008381526020019081526020016000206000018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610eae5780601f10610e8357610100808354040283529160200191610eae565b820191906000526020600020905b815481529060010190602001808311610e9157829003601f168201915b505050505085610c2f565b8015610edc57506000600360008381526020019081526020016000206002015414155b8015610efe575060036000828152602001908152602001600020600501544210155b15610f1f576003600082815260200190815260200160002060030154820191505b8080600101915050610df7565b819250610f58565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff92505b5050919050565b60008060008091507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff610f9185610d73565b1415801561101557506002846040518082805190602001908083835b602083101515610fd25780518252602082019150602081019050602083039250610fad565b6001836020036101000a03801982511681845116808217855250505050505090500191505090815260200160405180910390206001015461101285610d73565b13155b156111f457600090505b600154811215611170576110e1600360008381526020019081526020016000206000018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156110d65780601f106110ab576101008083540402835291602001916110d6565b820191906000526020600020905b8154815290600101906020018083116110b957829003601f168201915b505050505085610c2f565b801561110457506000600360008381526020019081526020016000206002015414155b8015611126575060036000828152602001908152602001600020600501544210155b1561116357600060036000838152602001908152602001600020600201819055506003600082815260200190815260200160002060030154820191505b808060010191505061101f565b816002856040518082805190602001908083835b6020831015156111a95780518252602082019150602081019050602083039250611184565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060010160008282540392505081905550819250611218565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff92505b5050919050565b60606000806002846040518082805190602001908083835b60208310151561125c5780518252602082019150602081019050602083039250611237565b6001836020036101000a03801982511681845116808217855250505050505090500191505090815260200160405180910390206000016002856040518082805190602001908083835b6020831015156112ca57805182526020820191506020810190506020830392506112a5565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040518091039020600101546002866040518082805190602001908083835b6020831015156113395780518252602082019150602081019050602083039250611314565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060020154828054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156114055780601f106113da57610100808354040283529160200191611405565b820191906000526020600020905b8154815290600101906020018083116113e857829003601f168201915b505050505092509250925092509193909250565b60008060008391508361142c8787610825565b12151561186557600090505b60015481121561185c576114fa600360008381526020019081526020016000206001018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156114ef5780601f106114c4576101008083540402835291602001916114ef565b820191906000526020600020905b8154815290600101906020018083116114d257829003601f168201915b505050505087610c2f565b801561153d575060016003600083815260200190815260200160002060020154148061153c575060026003600083815260200190815260200160002060020154145b5b1561184f57816003600083815260200190815260200160002060030154131561171e5781600360008381526020019081526020016000206003016000828254039250508190555060c060405190810160405280600360008481526020019081526020016000206000018054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561163a5780601f1061160f5761010080835404028352916020019161163a565b820191906000526020600020905b81548152906001019060200180831161161d57829003601f168201915b5050505050815260200186815260200160028152602001838152602001600360008481526020019081526020016000206004015481526020016003600084815260200190815260200160002060050154815250600360006001600081548092919060010191905055815260200190815260200160002060008201518160000190805190602001906116cc92919061200b565b5060208201518160010190805190602001906116e992919061200b565b5060408201518160020155606082015181600301556080820151816004015560a0820151816005015590505060009150611840565b6117d6600360008381526020019081526020016000206000018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156117cb5780601f106117a0576101008083540402835291602001916117cb565b820191906000526020600020905b8154815290600101906020018083116117ae57829003601f168201915b505050505086610c2f565b151561183f578460036000838152602001908152602001600020600101908051906020019061180692919061208b565b50600260036000838152602001908152602001600020600201819055506003600082815260200190815260200160002060030154820391505b5b600082141561184e5761185c565b5b8080600101915050611438565b60009250611889565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff92505b50509392505050565b600080600080915060016002866040518082805190602001908083835b6020831015156118d457805182526020820191506020810190506020830392506118af565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060020154148015611986575060006002856040518082805190602001908083835b60208310151561194d5780518252602082019150602081019050602083039250611928565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060020154145b15611c0957600090505b600154811215611b0957611a52600360008381526020019081526020016000206001018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015611a475780601f10611a1c57610100808354040283529160200191611a47565b820191906000526020600020905b815481529060010190602001808311611a2a57829003601f168201915b505050505086610c2f565b8015611a955750600160036000838152602001908152602001600020600201541480611a94575060026003600083815260200190815260200160002060020154145b5b15611afc5783600360008381526020019081526020016000206001019080519060200190611ac492919061208b565b506003806000838152602001908152602001600020600201819055506003600082815260200190815260200160002060030154820191505b8080600101915050611990565b816002866040518082805190602001908083835b602083101515611b425780518252602082019150602081019050602083039250611b1d565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060010160008282540192505081905550816002856040518082805190602001908083835b602083101515611bbe5780518252602082019150602081019050602083039250611b99565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060010160008282540392505081905550819250611c2d565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff92505b505092915050565b6060806000806000806003600088815260200190815260200160002060000160036000898152602001908152602001600020600101600360008a815260200190815260200160002060020154600360008b815260200190815260200160002060030154600360008c815260200190815260200160002060040154600360008d815260200190815260200160002060050154858054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015611d5b5780601f10611d3057610100808354040283529160200191611d5b565b820191906000526020600020905b815481529060010190602001808311611d3e57829003601f168201915b50505050509550848054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015611df75780601f10611dcc57610100808354040283529160200191611df7565b820191906000526020600020905b815481529060010190602001808311611dda57829003601f168201915b5050505050945095509550955095509550955091939550919395565b600060016002856040518082805190602001908083835b602083101515611e4f5780518252602082019150602081019050602083039250611e2a565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060020154148015611f01575060016002846040518082805190602001908083835b602083101515611ec85780518252602082019150602081019050602083039250611ea3565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060020154145b8015611f145750611f128484610c2f565b155b15611fe05760c0604051908101604052808581526020018481526020016001815260200183815260200142815260200161ea60420181525060036000600160008154809291906001019190505581526020019081526020016000206000820151816000019080519060200190611f8b92919061200b565b506020820151816001019080519060200190611fa892919061200b565b5060408201518160020155606082015181600301556080820151816004015560a0820151816005015590505060018054039050612004565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff90505b9392505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061204c57805160ff191683800117855561207a565b8280016001018555821561207a579182015b8281111561207957825182559160200191906001019061205e565b5b509050612087919061210b565b5090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106120cc57805160ff19168380011785556120fa565b828001600101855582156120fa579182015b828111156120f95782518255916020019190600101906120de565b5b509050612107919061210b565b5090565b61212d91905b80821115612129576000816000905550600101612111565b5090565b905600a165627a7a72305820d43ee995d974f3c4df21235fb44fd7c94e260c10af7e36f42a5567d345ea6d0a0029"

// DeploySupplyPlatform deploys a new contract, binding an instance of SupplyPlatform to it.
func DeploySupplyPlatform(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SupplyPlatform, error) {
	parsed, err := abi.JSON(strings.NewReader(SupplyPlatformABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SupplyPlatformBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SupplyPlatform{SupplyPlatformCaller: SupplyPlatformCaller{contract: contract}, SupplyPlatformTransactor: SupplyPlatformTransactor{contract: contract}, SupplyPlatformFilterer: SupplyPlatformFilterer{contract: contract}}, nil
}

func AsyncDeploySupplyPlatform(auth *bind.TransactOpts, handler func(*types.Receipt, error), backend bind.ContractBackend) (*types.Transaction, error) {
	parsed, err := abi.JSON(strings.NewReader(SupplyPlatformABI))
	if err != nil {
		return nil, err
	}

	tx, err := bind.AsyncDeployContract(auth, handler, parsed, common.FromHex(SupplyPlatformBin), backend)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

// SupplyPlatform is an auto generated Go binding around a Solidity contract.
type SupplyPlatform struct {
	SupplyPlatformCaller     // Read-only binding to the contract
	SupplyPlatformTransactor // Write-only binding to the contract
	SupplyPlatformFilterer   // Log filterer for contract events
}

// SupplyPlatformCaller is an auto generated read-only Go binding around a Solidity contract.
type SupplyPlatformCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SupplyPlatformTransactor is an auto generated write-only Go binding around a Solidity contract.
type SupplyPlatformTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SupplyPlatformFilterer is an auto generated log filtering Go binding around a Solidity contract events.
type SupplyPlatformFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SupplyPlatformSession is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type SupplyPlatformSession struct {
	Contract     *SupplyPlatform   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SupplyPlatformCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type SupplyPlatformCallerSession struct {
	Contract *SupplyPlatformCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// SupplyPlatformTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type SupplyPlatformTransactorSession struct {
	Contract     *SupplyPlatformTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// SupplyPlatformRaw is an auto generated low-level Go binding around a Solidity contract.
type SupplyPlatformRaw struct {
	Contract *SupplyPlatform // Generic contract binding to access the raw methods on
}

// SupplyPlatformCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type SupplyPlatformCallerRaw struct {
	Contract *SupplyPlatformCaller // Generic read-only contract binding to access the raw methods on
}

// SupplyPlatformTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type SupplyPlatformTransactorRaw struct {
	Contract *SupplyPlatformTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSupplyPlatform creates a new instance of SupplyPlatform, bound to a specific deployed contract.
func NewSupplyPlatform(address common.Address, backend bind.ContractBackend) (*SupplyPlatform, error) {
	contract, err := bindSupplyPlatform(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SupplyPlatform{SupplyPlatformCaller: SupplyPlatformCaller{contract: contract}, SupplyPlatformTransactor: SupplyPlatformTransactor{contract: contract}, SupplyPlatformFilterer: SupplyPlatformFilterer{contract: contract}}, nil
}

// NewSupplyPlatformCaller creates a new read-only instance of SupplyPlatform, bound to a specific deployed contract.
func NewSupplyPlatformCaller(address common.Address, caller bind.ContractCaller) (*SupplyPlatformCaller, error) {
	contract, err := bindSupplyPlatform(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SupplyPlatformCaller{contract: contract}, nil
}

// NewSupplyPlatformTransactor creates a new write-only instance of SupplyPlatform, bound to a specific deployed contract.
func NewSupplyPlatformTransactor(address common.Address, transactor bind.ContractTransactor) (*SupplyPlatformTransactor, error) {
	contract, err := bindSupplyPlatform(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SupplyPlatformTransactor{contract: contract}, nil
}

// NewSupplyPlatformFilterer creates a new log filterer instance of SupplyPlatform, bound to a specific deployed contract.
func NewSupplyPlatformFilterer(address common.Address, filterer bind.ContractFilterer) (*SupplyPlatformFilterer, error) {
	contract, err := bindSupplyPlatform(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SupplyPlatformFilterer{contract: contract}, nil
}

// bindSupplyPlatform binds a generic wrapper to an already deployed contract.
func bindSupplyPlatform(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SupplyPlatformABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SupplyPlatform *SupplyPlatformRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SupplyPlatform.Contract.SupplyPlatformCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SupplyPlatform *SupplyPlatformRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _SupplyPlatform.Contract.SupplyPlatformTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SupplyPlatform *SupplyPlatformRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _SupplyPlatform.Contract.SupplyPlatformTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SupplyPlatform *SupplyPlatformCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SupplyPlatform.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SupplyPlatform *SupplyPlatformTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _SupplyPlatform.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SupplyPlatform *SupplyPlatformTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _SupplyPlatform.Contract.contract.Transact(opts, method, params...)
}

// GetBill is a free data retrieval call binding the contract method 0xacb1dfb2.
//
// Solidity: function GetBill(int256 id) constant returns(string, string, int256, int256, uint256, uint256)
func (_SupplyPlatform *SupplyPlatformCaller) GetBill(opts *bind.CallOpts, id *big.Int) (string, string, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	var (
		ret0 = new(string)
		ret1 = new(string)
		ret2 = new(*big.Int)
		ret3 = new(*big.Int)
		ret4 = new(*big.Int)
		ret5 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
		ret5,
	}
	err := _SupplyPlatform.contract.Call(opts, out, "GetBill", id)
	return *ret0, *ret1, *ret2, *ret3, *ret4, *ret5, err
}

// GetBill is a free data retrieval call binding the contract method 0xacb1dfb2.
//
// Solidity: function GetBill(int256 id) constant returns(string, string, int256, int256, uint256, uint256)
func (_SupplyPlatform *SupplyPlatformSession) GetBill(id *big.Int) (string, string, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _SupplyPlatform.Contract.GetBill(&_SupplyPlatform.CallOpts, id)
}

// GetBill is a free data retrieval call binding the contract method 0xacb1dfb2.
//
// Solidity: function GetBill(int256 id) constant returns(string, string, int256, int256, uint256, uint256)
func (_SupplyPlatform *SupplyPlatformCallerSession) GetBill(id *big.Int) (string, string, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _SupplyPlatform.Contract.GetBill(&_SupplyPlatform.CallOpts, id)
}

// GetCompany is a free data retrieval call binding the contract method 0x5aa5af05.
//
// Solidity: function GetCompany(string Name) constant returns(string, int256, int256)
func (_SupplyPlatform *SupplyPlatformCaller) GetCompany(opts *bind.CallOpts, Name string) (string, *big.Int, *big.Int, error) {
	var (
		ret0 = new(string)
		ret1 = new(*big.Int)
		ret2 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _SupplyPlatform.contract.Call(opts, out, "GetCompany", Name)
	return *ret0, *ret1, *ret2, err
}

// GetCompany is a free data retrieval call binding the contract method 0x5aa5af05.
//
// Solidity: function GetCompany(string Name) constant returns(string, int256, int256)
func (_SupplyPlatform *SupplyPlatformSession) GetCompany(Name string) (string, *big.Int, *big.Int, error) {
	return _SupplyPlatform.Contract.GetCompany(&_SupplyPlatform.CallOpts, Name)
}

// GetCompany is a free data retrieval call binding the contract method 0x5aa5af05.
//
// Solidity: function GetCompany(string Name) constant returns(string, int256, int256)
func (_SupplyPlatform *SupplyPlatformCallerSession) GetCompany(Name string) (string, *big.Int, *big.Int, error) {
	return _SupplyPlatform.Contract.GetCompany(&_SupplyPlatform.CallOpts, Name)
}

// NewBill is a paid mutator transaction binding the contract method 0xf656f550.
//
// Solidity: function NewBill(string From, string To, int256 Money) returns(int256)
func (_SupplyPlatform *SupplyPlatformTransactor) NewBill(opts *bind.TransactOpts, From string, To string, Money *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _SupplyPlatform.contract.Transact(opts, "NewBill", From, To, Money)
}

func (_SupplyPlatform *SupplyPlatformTransactor) AsyncNewBill(handler func(*types.Receipt, error), opts *bind.TransactOpts, From string, To string, Money *big.Int) (*types.Transaction, error) {
	return _SupplyPlatform.contract.AsyncTransact(opts, handler, "NewBill", From, To, Money)
}

// NewBill is a paid mutator transaction binding the contract method 0xf656f550.
//
// Solidity: function NewBill(string From, string To, int256 Money) returns(int256)
func (_SupplyPlatform *SupplyPlatformSession) NewBill(From string, To string, Money *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _SupplyPlatform.Contract.NewBill(&_SupplyPlatform.TransactOpts, From, To, Money)
}

func (_SupplyPlatform *SupplyPlatformSession) AsyncNewBill(handler func(*types.Receipt, error), From string, To string, Money *big.Int) (*types.Transaction, error) {
	return _SupplyPlatform.Contract.AsyncNewBill(handler, &_SupplyPlatform.TransactOpts, From, To, Money)
}

// NewBill is a paid mutator transaction binding the contract method 0xf656f550.
//
// Solidity: function NewBill(string From, string To, int256 Money) returns(int256)
func (_SupplyPlatform *SupplyPlatformTransactorSession) NewBill(From string, To string, Money *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _SupplyPlatform.Contract.NewBill(&_SupplyPlatform.TransactOpts, From, To, Money)
}

func (_SupplyPlatform *SupplyPlatformTransactorSession) AsyncNewBill(handler func(*types.Receipt, error), From string, To string, Money *big.Int) (*types.Transaction, error) {
	return _SupplyPlatform.Contract.AsyncNewBill(handler, &_SupplyPlatform.TransactOpts, From, To, Money)
}

// Raise is a paid mutator transaction binding the contract method 0x90bee2ee.
//
// Solidity: function Raise(string From, string To) returns(int256)
func (_SupplyPlatform *SupplyPlatformTransactor) Raise(opts *bind.TransactOpts, From string, To string) (*types.Transaction, *types.Receipt, error) {
	return _SupplyPlatform.contract.Transact(opts, "Raise", From, To)
}

func (_SupplyPlatform *SupplyPlatformTransactor) AsyncRaise(handler func(*types.Receipt, error), opts *bind.TransactOpts, From string, To string) (*types.Transaction, error) {
	return _SupplyPlatform.contract.AsyncTransact(opts, handler, "Raise", From, To)
}

// Raise is a paid mutator transaction binding the contract method 0x90bee2ee.
//
// Solidity: function Raise(string From, string To) returns(int256)
func (_SupplyPlatform *SupplyPlatformSession) Raise(From string, To string) (*types.Transaction, *types.Receipt, error) {
	return _SupplyPlatform.Contract.Raise(&_SupplyPlatform.TransactOpts, From, To)
}

func (_SupplyPlatform *SupplyPlatformSession) AsyncRaise(handler func(*types.Receipt, error), From string, To string) (*types.Transaction, error) {
	return _SupplyPlatform.Contract.AsyncRaise(handler, &_SupplyPlatform.TransactOpts, From, To)
}

// Raise is a paid mutator transaction binding the contract method 0x90bee2ee.
//
// Solidity: function Raise(string From, string To) returns(int256)
func (_SupplyPlatform *SupplyPlatformTransactorSession) Raise(From string, To string) (*types.Transaction, *types.Receipt, error) {
	return _SupplyPlatform.Contract.Raise(&_SupplyPlatform.TransactOpts, From, To)
}

func (_SupplyPlatform *SupplyPlatformTransactorSession) AsyncRaise(handler func(*types.Receipt, error), From string, To string) (*types.Transaction, error) {
	return _SupplyPlatform.Contract.AsyncRaise(handler, &_SupplyPlatform.TransactOpts, From, To)
}

// Register is a paid mutator transaction binding the contract method 0x146e5209.
//
// Solidity: function Register(string Name, int256 Money, int256 Type) returns()
func (_SupplyPlatform *SupplyPlatformTransactor) Register(opts *bind.TransactOpts, Name string, Money *big.Int, Type *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _SupplyPlatform.contract.Transact(opts, "Register", Name, Money, Type)
}

func (_SupplyPlatform *SupplyPlatformTransactor) AsyncRegister(handler func(*types.Receipt, error), opts *bind.TransactOpts, Name string, Money *big.Int, Type *big.Int) (*types.Transaction, error) {
	return _SupplyPlatform.contract.AsyncTransact(opts, handler, "Register", Name, Money, Type)
}

// Register is a paid mutator transaction binding the contract method 0x146e5209.
//
// Solidity: function Register(string Name, int256 Money, int256 Type) returns()
func (_SupplyPlatform *SupplyPlatformSession) Register(Name string, Money *big.Int, Type *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _SupplyPlatform.Contract.Register(&_SupplyPlatform.TransactOpts, Name, Money, Type)
}

func (_SupplyPlatform *SupplyPlatformSession) AsyncRegister(handler func(*types.Receipt, error), Name string, Money *big.Int, Type *big.Int) (*types.Transaction, error) {
	return _SupplyPlatform.Contract.AsyncRegister(handler, &_SupplyPlatform.TransactOpts, Name, Money, Type)
}

// Register is a paid mutator transaction binding the contract method 0x146e5209.
//
// Solidity: function Register(string Name, int256 Money, int256 Type) returns()
func (_SupplyPlatform *SupplyPlatformTransactorSession) Register(Name string, Money *big.Int, Type *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _SupplyPlatform.Contract.Register(&_SupplyPlatform.TransactOpts, Name, Money, Type)
}

func (_SupplyPlatform *SupplyPlatformTransactorSession) AsyncRegister(handler func(*types.Receipt, error), Name string, Money *big.Int, Type *big.Int) (*types.Transaction, error) {
	return _SupplyPlatform.Contract.AsyncRegister(handler, &_SupplyPlatform.TransactOpts, Name, Money, Type)
}

// Repay is a paid mutator transaction binding the contract method 0x54b720d4.
//
// Solidity: function Repay(string From) returns(int256)
func (_SupplyPlatform *SupplyPlatformTransactor) Repay(opts *bind.TransactOpts, From string) (*types.Transaction, *types.Receipt, error) {
	return _SupplyPlatform.contract.Transact(opts, "Repay", From)
}

func (_SupplyPlatform *SupplyPlatformTransactor) AsyncRepay(handler func(*types.Receipt, error), opts *bind.TransactOpts, From string) (*types.Transaction, error) {
	return _SupplyPlatform.contract.AsyncTransact(opts, handler, "Repay", From)
}

// Repay is a paid mutator transaction binding the contract method 0x54b720d4.
//
// Solidity: function Repay(string From) returns(int256)
func (_SupplyPlatform *SupplyPlatformSession) Repay(From string) (*types.Transaction, *types.Receipt, error) {
	return _SupplyPlatform.Contract.Repay(&_SupplyPlatform.TransactOpts, From)
}

func (_SupplyPlatform *SupplyPlatformSession) AsyncRepay(handler func(*types.Receipt, error), From string) (*types.Transaction, error) {
	return _SupplyPlatform.Contract.AsyncRepay(handler, &_SupplyPlatform.TransactOpts, From)
}

// Repay is a paid mutator transaction binding the contract method 0x54b720d4.
//
// Solidity: function Repay(string From) returns(int256)
func (_SupplyPlatform *SupplyPlatformTransactorSession) Repay(From string) (*types.Transaction, *types.Receipt, error) {
	return _SupplyPlatform.Contract.Repay(&_SupplyPlatform.TransactOpts, From)
}

func (_SupplyPlatform *SupplyPlatformTransactorSession) AsyncRepay(handler func(*types.Receipt, error), From string) (*types.Transaction, error) {
	return _SupplyPlatform.Contract.AsyncRepay(handler, &_SupplyPlatform.TransactOpts, From)
}

// RepayCount is a paid mutator transaction binding the contract method 0x4c3eb6ef.
//
// Solidity: function RepayCount(string From) returns(int256)
func (_SupplyPlatform *SupplyPlatformTransactor) RepayCount(opts *bind.TransactOpts, From string) (*types.Transaction, *types.Receipt, error) {
	return _SupplyPlatform.contract.Transact(opts, "RepayCount", From)
}

func (_SupplyPlatform *SupplyPlatformTransactor) AsyncRepayCount(handler func(*types.Receipt, error), opts *bind.TransactOpts, From string) (*types.Transaction, error) {
	return _SupplyPlatform.contract.AsyncTransact(opts, handler, "RepayCount", From)
}

// RepayCount is a paid mutator transaction binding the contract method 0x4c3eb6ef.
//
// Solidity: function RepayCount(string From) returns(int256)
func (_SupplyPlatform *SupplyPlatformSession) RepayCount(From string) (*types.Transaction, *types.Receipt, error) {
	return _SupplyPlatform.Contract.RepayCount(&_SupplyPlatform.TransactOpts, From)
}

func (_SupplyPlatform *SupplyPlatformSession) AsyncRepayCount(handler func(*types.Receipt, error), From string) (*types.Transaction, error) {
	return _SupplyPlatform.Contract.AsyncRepayCount(handler, &_SupplyPlatform.TransactOpts, From)
}

// RepayCount is a paid mutator transaction binding the contract method 0x4c3eb6ef.
//
// Solidity: function RepayCount(string From) returns(int256)
func (_SupplyPlatform *SupplyPlatformTransactorSession) RepayCount(From string) (*types.Transaction, *types.Receipt, error) {
	return _SupplyPlatform.Contract.RepayCount(&_SupplyPlatform.TransactOpts, From)
}

func (_SupplyPlatform *SupplyPlatformTransactorSession) AsyncRepayCount(handler func(*types.Receipt, error), From string) (*types.Transaction, error) {
	return _SupplyPlatform.Contract.AsyncRepayCount(handler, &_SupplyPlatform.TransactOpts, From)
}

// Transfer is a paid mutator transaction binding the contract method 0x88169602.
//
// Solidity: function Transfer(string From, string To, int256 Money) returns(int256)
func (_SupplyPlatform *SupplyPlatformTransactor) Transfer(opts *bind.TransactOpts, From string, To string, Money *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _SupplyPlatform.contract.Transact(opts, "Transfer", From, To, Money)
}

func (_SupplyPlatform *SupplyPlatformTransactor) AsyncTransfer(handler func(*types.Receipt, error), opts *bind.TransactOpts, From string, To string, Money *big.Int) (*types.Transaction, error) {
	return _SupplyPlatform.contract.AsyncTransact(opts, handler, "Transfer", From, To, Money)
}

// Transfer is a paid mutator transaction binding the contract method 0x88169602.
//
// Solidity: function Transfer(string From, string To, int256 Money) returns(int256)
func (_SupplyPlatform *SupplyPlatformSession) Transfer(From string, To string, Money *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _SupplyPlatform.Contract.Transfer(&_SupplyPlatform.TransactOpts, From, To, Money)
}

func (_SupplyPlatform *SupplyPlatformSession) AsyncTransfer(handler func(*types.Receipt, error), From string, To string, Money *big.Int) (*types.Transaction, error) {
	return _SupplyPlatform.Contract.AsyncTransfer(handler, &_SupplyPlatform.TransactOpts, From, To, Money)
}

// Transfer is a paid mutator transaction binding the contract method 0x88169602.
//
// Solidity: function Transfer(string From, string To, int256 Money) returns(int256)
func (_SupplyPlatform *SupplyPlatformTransactorSession) Transfer(From string, To string, Money *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _SupplyPlatform.Contract.Transfer(&_SupplyPlatform.TransactOpts, From, To, Money)
}

func (_SupplyPlatform *SupplyPlatformTransactorSession) AsyncTransfer(handler func(*types.Receipt, error), From string, To string, Money *big.Int) (*types.Transaction, error) {
	return _SupplyPlatform.Contract.AsyncTransfer(handler, &_SupplyPlatform.TransactOpts, From, To, Money)
}

// TransferCount is a paid mutator transaction binding the contract method 0x0c86c6b8.
//
// Solidity: function TransferCount(string From, string To) returns(int256)
func (_SupplyPlatform *SupplyPlatformTransactor) TransferCount(opts *bind.TransactOpts, From string, To string) (*types.Transaction, *types.Receipt, error) {
	return _SupplyPlatform.contract.Transact(opts, "TransferCount", From, To)
}

func (_SupplyPlatform *SupplyPlatformTransactor) AsyncTransferCount(handler func(*types.Receipt, error), opts *bind.TransactOpts, From string, To string) (*types.Transaction, error) {
	return _SupplyPlatform.contract.AsyncTransact(opts, handler, "TransferCount", From, To)
}

// TransferCount is a paid mutator transaction binding the contract method 0x0c86c6b8.
//
// Solidity: function TransferCount(string From, string To) returns(int256)
func (_SupplyPlatform *SupplyPlatformSession) TransferCount(From string, To string) (*types.Transaction, *types.Receipt, error) {
	return _SupplyPlatform.Contract.TransferCount(&_SupplyPlatform.TransactOpts, From, To)
}

func (_SupplyPlatform *SupplyPlatformSession) AsyncTransferCount(handler func(*types.Receipt, error), From string, To string) (*types.Transaction, error) {
	return _SupplyPlatform.Contract.AsyncTransferCount(handler, &_SupplyPlatform.TransactOpts, From, To)
}

// TransferCount is a paid mutator transaction binding the contract method 0x0c86c6b8.
//
// Solidity: function TransferCount(string From, string To) returns(int256)
func (_SupplyPlatform *SupplyPlatformTransactorSession) TransferCount(From string, To string) (*types.Transaction, *types.Receipt, error) {
	return _SupplyPlatform.Contract.TransferCount(&_SupplyPlatform.TransactOpts, From, To)
}

func (_SupplyPlatform *SupplyPlatformTransactorSession) AsyncTransferCount(handler func(*types.Receipt, error), From string, To string) (*types.Transaction, error) {
	return _SupplyPlatform.Contract.AsyncTransferCount(handler, &_SupplyPlatform.TransactOpts, From, To)
}

// IsEqual is a paid mutator transaction binding the contract method 0x465c4105.
//
// Solidity: function isEqual(string s1, string s2) returns(bool)
func (_SupplyPlatform *SupplyPlatformTransactor) IsEqual(opts *bind.TransactOpts, s1 string, s2 string) (*types.Transaction, *types.Receipt, error) {
	return _SupplyPlatform.contract.Transact(opts, "isEqual", s1, s2)
}

func (_SupplyPlatform *SupplyPlatformTransactor) AsyncIsEqual(handler func(*types.Receipt, error), opts *bind.TransactOpts, s1 string, s2 string) (*types.Transaction, error) {
	return _SupplyPlatform.contract.AsyncTransact(opts, handler, "isEqual", s1, s2)
}

// IsEqual is a paid mutator transaction binding the contract method 0x465c4105.
//
// Solidity: function isEqual(string s1, string s2) returns(bool)
func (_SupplyPlatform *SupplyPlatformSession) IsEqual(s1 string, s2 string) (*types.Transaction, *types.Receipt, error) {
	return _SupplyPlatform.Contract.IsEqual(&_SupplyPlatform.TransactOpts, s1, s2)
}

func (_SupplyPlatform *SupplyPlatformSession) AsyncIsEqual(handler func(*types.Receipt, error), s1 string, s2 string) (*types.Transaction, error) {
	return _SupplyPlatform.Contract.AsyncIsEqual(handler, &_SupplyPlatform.TransactOpts, s1, s2)
}

// IsEqual is a paid mutator transaction binding the contract method 0x465c4105.
//
// Solidity: function isEqual(string s1, string s2) returns(bool)
func (_SupplyPlatform *SupplyPlatformTransactorSession) IsEqual(s1 string, s2 string) (*types.Transaction, *types.Receipt, error) {
	return _SupplyPlatform.Contract.IsEqual(&_SupplyPlatform.TransactOpts, s1, s2)
}

func (_SupplyPlatform *SupplyPlatformTransactorSession) AsyncIsEqual(handler func(*types.Receipt, error), s1 string, s2 string) (*types.Transaction, error) {
	return _SupplyPlatform.Contract.AsyncIsEqual(handler, &_SupplyPlatform.TransactOpts, s1, s2)
}
