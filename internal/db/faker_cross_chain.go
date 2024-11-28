package db

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"math/rand"
	"os"
	"time"

	"github.com/go-faker/faker/v4"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Assuming these are your existing chain and token details
type ChainToken struct {
	ChainID      int64
	TokenAddress string
}



// ConvertToPgNumeric converts a big.Float to pgtype.Numeric
func ConvertToPgNumeric(f big.Float) pgtype.Numeric {
	value := pgtype.Numeric{
		Valid: true,
		Int:   new(big.Int).SetInt64(0), // Using 18 decimals as common in most tokens
		Exp:   -18,
	}

	value.ScanScientific(f.String())
	return value
}

func ConvertStringToPgNumeric(f string) pgtype.Numeric {
	value := pgtype.Numeric{
		Valid: true,
		Int:   new(big.Int).SetInt64(0), // Using 18 decimals as common in most tokens
		Exp:   -18,
	}

	value.ScanScientific(f)
	return value
}

func GenerateFakeSwaps(ctx context.Context, conn *pgxpool.Pool, numSwaps int) error {
	// Mock data - replace with your actual supported chains and tokens
	supportedTokens := []ChainToken{
		{ChainID: 1, TokenAddress: "0xdAC17F958D2ee523a2206206994597C13D831ec7"},   // USDT on Ethereum
		{ChainID: 56, TokenAddress: "0x55d398326f99059fF775485246999027B3197955"},  // USDT on BSC
		{ChainID: 137, TokenAddress: "0xc2132D05D31c914a87C6611C10748AEb04B58e8F"}, // USDT on Polygon
		// Add more tokens as needed
	}

	for i := 0; i < numSwaps; i++ {
		// Get random source and destination tokens
		sourceToken := supportedTokens[rand.Intn(len(supportedTokens))]
		destToken := supportedTokens[rand.Intn(len(supportedTokens))]

		// Ensure source and destination are different
		for destToken == sourceToken {
			destToken = supportedTokens[rand.Intn(len(supportedTokens))]
		}

		// Generate random amounts (between 1 and 1000 USDT with 6 decimals)
		sourceAmount := new(big.Int).Mul(
			big.NewInt(rand.Int63n(1000)+1),
			big.NewInt(1000000),
		).String()

		destAmount := new(big.Int).Mul(
			big.NewInt(rand.Int63n(1000)+1),
			big.NewInt(1000000),
		).String()

		// Generate random dollar values (same range as amounts)
		dollarValue := big.NewFloat(float64(rand.Intn(1000) + 1))
		feeValue := big.NewFloat(0).Mul(dollarValue, big.NewFloat(0.001)) // 0.1% fee

		// Generate random addresses
		sourceAddr := fmt.Sprintf("0x%x", rand.Int63())
		destAddr := fmt.Sprintf("0x%x", rand.Int63())

		// Generate unique transaction hash
		txHash := fmt.Sprintf("0x%x", faker.UnixTime())

		queries := New(conn)

		createSwapRequest := CreateSwapOrderParams{
			SourceChainID:           sourceToken.ChainID,
			SourceTokenAddress:      sourceToken.TokenAddress,
			SourceAmount:            ConvertStringToPgNumeric(sourceAmount),
			DestinationChainID:      destToken.ChainID,
			DestinationTokenAddress: destToken.TokenAddress,
			DestinationAmount:       ConvertStringToPgNumeric(destAmount),
			DollarValue:             ConvertToPgNumeric(*dollarValue),
			FeeDollarValue:          ConvertToPgNumeric(*feeValue),
			SourceAddress:           sourceAddr,
			DestinationAddress:      destAddr,
			TransactionHash:         txHash,
		}

		err := queries.CreateSwapOrder(
			ctx,
			createSwapRequest,
		)

		if err != nil {
			return fmt.Errorf("failed to insert swap %d: %w", i, err)
		}
	}

	return nil
}

// Usage example:
func RunCrossChainFaker() {
	ctx := context.Background()

	// Initialize random seed
	rand.NewSource(time.Now().UnixNano())

	// Setup connection pool
	// dbURL := "postgres://postgres:postgres@localhost:5432/rampx?pool_max_conns=10"

    pool, err := NewDBPool(os.Getenv("DATABASE_URI"))
    if(err != nil){
        log.Panicln(err.Error())
    }
    defer pool.Close()

	// Generate 100 fake swaps
	err = GenerateFakeSwaps(ctx, pool, 100)
	if err != nil {
		panic(err)
	}
}
