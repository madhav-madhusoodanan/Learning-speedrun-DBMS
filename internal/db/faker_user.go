package db

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
)

func GenerateEthAddress() (string, error) {
	// Generate 20 random bytes (160 bits) for the address
	b := make([]byte, 20)
	_, err := rand.Read(b)
	if err != nil {
		return "", fmt.Errorf("failed to generate random bytes: %v", err)
	}

	// Convert to hex and add 0x prefix
	address := "0x" + hex.EncodeToString(b)

	// Ensure address is lowercase (Ethereum standard)
	address = strings.ToLower(address)

	return address, nil
}

var firstNames = []string{
	"James", "Mary", "John", "Patricia", "Robert", "Jennifer", "Michael",
	"Linda", "William", "Elizabeth", "David", "Barbara", "Richard", "Susan",
	"Joseph", "Jessica", "Thomas", "Sarah", "Charles", "Karen", "Emma",
	"Olivia", "Ava", "Isabella", "Sophia", "Mia", "Charlotte", "Amelia",
	"Harper", "Evelyn", "Liam", "Noah", "Oliver", "Elijah", "William",
	"James", "Benjamin", "Lucas", "Henry", "Theodore",
}

// Assuming these are your existing chain and token details
func GenerateFakeUserData(ctx context.Context, conn *pgxpool.Pool, numSwaps int) error {

	for i := 0; i < numSwaps; i++ {
		// Get random source and destination tokens
		name := firstNames[0]
		chainId := int64(137)

		userAddress, _ := GenerateEthAddress()

		queries := New(conn)

		createUserRequest := AddUserWalletParams{
			UserName:    name,
			ChainID:     chainId,
			UserAddress: userAddress,
		}

		err := queries.AddUserWallet(
			ctx,
			createUserRequest,
		)

		if err != nil {
			return fmt.Errorf("failed to insert swap %d: %w", i, err)
		}
	}

	return nil
}

// Usage example:
func RunUserFaker() {
	ctx := context.Background()

	// Setup connection pool
	pool, err := NewDBPool(os.Getenv("DATABASE_URI"))
	if err != nil {
		log.Panicln(err.Error())
	}
	defer pool.Close()

	// Generate 100 fake swaps
	err = GenerateFakeUserData(ctx, pool, 2)
	if err != nil {
		log.Println(err)
	}
}
