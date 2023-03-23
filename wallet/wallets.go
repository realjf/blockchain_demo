package wallet

const (
	walletFile = "./tmp/wallets.dat"
)

type Wallets struct {
	Wallets map[string]*Wallet `json:"wallets"`
}

func NewWallets() (*Wallets, error) {
	wallets := &Wallets{}
	wallets.Wallets = make(map[string]*Wallet)
	err := wallets.LoadFile()
	return wallets, err
}

func (ws *Wallets) AddWallet() string {
	wallet := MakeWallet()
	if wallet == nil {
		panic("make wallet failed")
	}
	address := string(wallet.Address()[:])

	ws.Wallets[address] = wallet
	return address
}

func (ws *Wallets) GetAllAddresses() []string {
	var addresses []string

	for address := range ws.Wallets {
		addresses = append(addresses, address)
	}

	return addresses
}

func (ws *Wallets) GetWallet(address string) Wallet {
	return *ws.Wallets[address]
}

func (ws *Wallets) LoadFile() error {
	wallets, err := DeserializeWallets()
	if err != nil {
		return err
	}
	ws.Wallets = wallets.Wallets
	return nil
}

func (ws *Wallets) SaveFile() {
	SerializeWallets(ws)
}
