package account

type dwRequest struct {
	isDeposit bool
	name      string
	accNumber uint64
	password  string
	currency  string
	amount    uint64
}

func (req *dwRequest) unmarshal(data []byte) {

}

type dwResponse struct {
	balance uint64
}

func (res *dwResponse) marshal() []byte {
	return nil
}

type dwMonitorResponse struct {
	accNumber uint64
	isDeposit bool
	amount    uint64
}

func (res *dwMonitorResponse) marshal() []byte {
	return nil
}

func DepositWithdraw(content []byte) ([]byte, []byte) {

}
