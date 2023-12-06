package repo

import (
	"context"
	"database/sql"
	"fmt"
	"sync"

	"github.com/evrone/go-clean-template/internal/blockchain/transport"
	blockchainlogic "github.com/evrone/go-clean-template/pkg/blockchain_logic"
)

type BlockchainRepo struct {
	*sql.DB
	chain             *blockchainlogic.Blockchain
	userGrpcTransport *transport.UserGrpcTransport
}

func NewBlockchainRepo(db *sql.DB, address string, userGrpcTransport *transport.UserGrpcTransport) *BlockchainRepo {
	chain := blockchainlogic.CreateBlockchain(db, address)

	return &BlockchainRepo{db, chain, userGrpcTransport}
}

func (br *BlockchainRepo) GetWallets(_ context.Context) ([]string, error) {
	res := blockchainlogic.ListAddresses()

	return res, nil
}

func (br *BlockchainRepo) GetWallet(ctx context.Context, userID string) (wallet string, err error) {
	address, err := br.userGrpcTransport.GetUserWallet(ctx, userID)
	if err != nil {
		return "", err
	}
	if address.Wallet == "" {
		return "", fmt.Errorf("user does not have a wallet")
	}

	return address.Wallet, nil
}

func (br *BlockchainRepo) GetBalance(ctx context.Context, userID string) (balance float64, err error) {
	address, err := br.GetWallet(ctx, userID)
	if err != nil {
		return 0, err
	}

	res := br.chain.GetBalance(address)

	return res, nil
}

func (br *BlockchainRepo) GetBalanceByAddress(_ context.Context, address string) (balance float64, err error) {
	res := br.chain.GetBalance(address)

	return res, nil
}

func (br *BlockchainRepo) GetBalanceUSD(ctx context.Context, userID string) (balance float64, err error) {
	address, err := br.GetWallet(ctx, userID)
	if err != nil {
		return -1, err
	}
	res, err := br.chain.GetBalanceInUSD(address)
	if err != nil {
		return -1, err
	}

	return res, nil
}

func (br *BlockchainRepo) CreateWallet(ctx context.Context, userID string) (string, error) {
	builder := NewWalletBuilder(br, ctx, userID)
	err := builder.Build()
	if err != nil {
		return "", err
	}

	return builder.address, nil
}

func (br *BlockchainRepo) Send(ctx context.Context, from, to string, amount float64, wg *sync.WaitGroup) error {
	defer wg.Done()
	user, err := br.userGrpcTransport.GetUserByID(ctx, from)
	if err != nil {
		return err
	}
	err = br.chain.Send(user.Wallet, to, amount)
	if err != nil {
		return err
	}

	return nil
}

func (br *BlockchainRepo) TopUp(ctx context.Context, from, to string, amount float64, wg *sync.WaitGroup) error {
	defer wg.Done()
	user, err := br.userGrpcTransport.GetUserByID(ctx, to)
	if err != nil {
		return err
	}
	err = br.chain.Send(from, user.Wallet, amount)
	if err != nil {
		return err
	}

	return nil
}
