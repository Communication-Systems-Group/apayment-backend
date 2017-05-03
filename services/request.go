package services

import (
	"github.com/scmo/foodchain-backend/models"
	"github.com/astaxie/beego"
	"github.com/scmo/foodchain-backend/ethereum"
	"github.com/astaxie/beego/orm"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"math/big"
	"github.com/scmo/foodchain-backend/smart-contracts/request"
)

func CreateRequest(r models.Request) error {
	ethereumController := ethereum.GetEthereumController()
	beego.Debug(r.Remark)
	address, tx, _, err := smartcontracts.DeployRequestContract(ethereumController.Auth, ethereumController.Client, r.User.Id, getContributionCodes(&r), r.Remark)
	if err != nil {
		beego.Critical("Failed to deploy new token contract: ", err)
	}
	beego.Info("Contract pending deploy: ", address.String())
	beego.Info("Transaction waiting to be mined: ", tx.Hash().String())

	r.Address = address.String()

	o := orm.NewOrm()
	_, err = o.Insert(&r)
	if err != nil {
		beego.Error("Failed to insert new Request: ", err)
	}
	return err
}

func GetAllRequests() []*models.Request {
	o := orm.NewOrm()
	var requests []*models.Request
	o.QueryTable(new(models.Request)).RelatedSel().All(&requests)
	for _, request := range requests {
		requestContract, err := getRequestByAddress(request.Address)
		if err != nil {
			beego.Error("Failed to instantiate a Token contract: %v", err)
		}
		assignRequest(request, requestContract)
	}
	return requests
}

func GetAllRequestsByUserId(userId int64) []*models.Request {
	o := orm.NewOrm()
	var requests []*models.Request
	o.QueryTable(new(models.Request)).Filter("user", userId).RelatedSel().All(&requests)
	for _, request := range requests {
		requestContract, err := getRequestByAddress(request.Address)
		if err != nil {
			beego.Error("Failed to instantiate a Token contract: %v", err)
		}
		assignRequest(request, requestContract)
	}
	return requests
}

func GetRequestById(requestId int64) *models.Request {
	o := orm.NewOrm()
	var request models.Request
	err := o.QueryTable(new(models.Request)).Filter("Id", requestId).RelatedSel().One(&request)
	if err != nil {
		beego.Error("Error while fetching Request.", err)
	}
	requestContract, err := getRequestByAddress(request.Address)
	if err != nil {
		beego.Error("Failed to instantiate a Token contract: %v", err)
	}
	assignRequest(&request, requestContract)

	// load control categories

	return &request
}

func GetRequestAddressById(requestId int64) string {
	o := orm.NewOrm()
	var request models.Request
	err := o.QueryTable(new(models.Request)).Filter("Id", requestId).RelatedSel().One(&request)
	if err != nil {
		beego.Error("Error while fetching Request.", err)
	}

	return request.Address
}

func GetAllRequestsForInspection() []*models.Request {
	o := orm.NewOrm()
	var requests []*models.Request
	o.QueryTable(new(models.Request)).Filter("inspector__isnull", false).RelatedSel().All(&requests)
	return requests
}

func GetAllRequestsForInspectionByInspectorId(inspectorId int64) []*models.Request {
	o := orm.NewOrm()
	var requests []*models.Request
	o.QueryTable(new(models.Request)).Filter("inspector", inspectorId).RelatedSel().All(&requests)
	return requests
}

func AddInspectorToRequest(request *models.Request) {
	// Add to DB
	o := orm.NewOrm()
	o.Update(request, "Inspector")

	// Add to the SmartContract
	requestContract, err := getRequestByAddress(request.Address)
	if err != nil {
		beego.Error("Error while fetching RequestContract by Address: ", err)
	}
	session := getRequestContractSession(requestContract)
	tx, err := session.SetInspectorId(request.Inspector.Id)
	if err != nil {
		beego.Critical("Failed to update name: ", err)
	}
	beego.Info("Transaction waiting to be mined: ", tx.Hash().String())
}

// Add inspection Lacks to Request
func AddLacksToRequest(inspection *models.Inspection) error {

	requestAddress := GetRequestAddressById(inspection.RequestId)
	// Add to the SmartContract
	requestContract, err := getRequestByAddress(requestAddress)
	if err != nil {
		beego.Error("Error while fetching RequestContract by Address: ", err)
		return err
	}
	session := getRequestContractSession(requestContract)
	for _, lack := range inspection.Lacks {
		tx, err := session.AddLack(lack.ContributionCode, lack.ControlCategoryId, lack.PointGroupId, lack.ControlPointId, lack.LackId)
		if err != nil {
			beego.Critical("Failed to update name: ", err)
			return err
		}
		beego.Info("Transaction waiting to be mined: ", tx.Hash().String())
	}
	return err
}

func getRequestContractSession(requestContract *smartcontracts.RequestContract) (*smartcontracts.RequestContractSession) {
	ethereumController := ethereum.GetEthereumController()
	requestContractSesssion := &smartcontracts.RequestContractSession{
		Contract:requestContract,
		CallOpts:bind.CallOpts{Pending:true},
		TransactOpts:bind.TransactOpts{
			From: ethereumController.Auth.From,
			Signer: ethereumController.Auth.Signer,
			GasLimit: big.NewInt(3141592),
		},
	}
	return requestContractSesssion

}

func getContributionCodes(request *models.Request) []uint16 {
	var codes = make([]uint16, len(request.Contributions))
	for i, contribution := range request.Contributions {
		codes[i] = contribution.Code
	}
	return codes
}

func getRequestByAddress(address string) (*smartcontracts.RequestContract, error) {
	ethereumController := ethereum.GetEthereumController()
	return smartcontracts.NewRequestContract(common.HexToAddress(address), ethereumController.Client)
}

func assignRequest(request *models.Request, requestContract *smartcontracts.RequestContract) {
	session := getRequestContractSession(requestContract)

	userId, err := session.UserId()
	if err != nil {
		beego.Error("Error while reading UserId from Contract: ", err)
	}
	if (userId != request.User.Id) {
		beego.Error("Request User and Contract UserId does not match!")
	}

	setInspector(request, session)
	setContributions(request, session)
	setLacksInspected(request, session)

	request.Remark, err = session.Remark()
	if err != nil {
		beego.Error("Failed to instantiate a Token contract: ", err)
	}

	createdTimestamp, err := session.CreatedTimestamp()
	if err != nil {
		beego.Error("Error while reading createdTimestamp from Contract: ", err)
	}
	request.Created = createdTimestamp
}
func setInspector(request *models.Request, session *smartcontracts.RequestContractSession) {
	if (request.Inspector != nil) {
		inspectorId, err := session.InspectorId()
		if err != nil {
			beego.Error("Error while reading InspectorId from Contract: ", err)
		}
		if (inspectorId != request.Inspector.Id) {
			beego.Debug(inspectorId, request.Inspector.Id)
			beego.Error("Request Inspector Id and Contract InspectorId does not match!")
		}
	}
}

func setContributions(request *models.Request, session *smartcontracts.RequestContractSession) {
	request.Contributions = make([]*models.Contribution, 0)
	next := true
	index := big.NewInt(0)
	for next {
		code, err := session.ContributionCodes(index)
		if err != nil {
			next = false
		}
		if err == nil {
			contribution, err := GetContributionByCode(code)
			if err != nil {
				beego.Error("Error getting Contribution", err)
			}
			request.Contributions = append(request.Contributions, contribution)
		}
		index.Add(index, big.NewInt(1))
	}

}

func setLacksInspected(request *models.Request, session *smartcontracts.RequestContractSession) {
	request.InspectionLacks = make([]*models.InspectionLack, 0)
	numLack, err := session.NumLacks()
	if err != nil {
		beego.Error("Error getting NumLacks", err)
		return
	}
	for i := new(big.Int).Set(big.NewInt(1)); i.Cmp(numLack) < 0; i.Add(i, big.NewInt(1)) {
		lack, err := session.Lacks(i)
		if err != nil {
			beego.Error("Error getting Lack", err)
			return
		}
		inspectionLacks := models.InspectionLack(lack)
		request.InspectionLacks = append(request.InspectionLacks, &inspectionLacks)
	}
}


