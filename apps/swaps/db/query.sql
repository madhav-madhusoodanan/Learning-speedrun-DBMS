-- name: CreateSwapOrder :exec
INSERT INTO rampx_cross_chain_swaps (
  source_chain, destination_chain, source_token, destination_token, source_amount, destination_amount, dollar_value, source_address, destination_address, tx_status, transaction_hash
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11
);

-- name: GetDailyVolume :many
SELECT 
    DATE(tx_timestamp) as trade_date,
    COUNT(*) as number_of_trades,
    SUM(dollar_value) as total_daily_volume
FROM rampx_cross_chain_swaps
WHERE 
    tx_timestamp >= CURRENT_DATE - INTERVAL '30 days'
    AND status = 'COMPLETED'
GROUP BY DATE(tx_timestamp)
ORDER BY trade_date DESC;

-- name: GetPendingTransactions :many
SELECT 
    transaction_hash,
    source_chain,
    destination_chain,
    tx_timestamp
FROM rampx_cross_chain_swaps
WHERE 
    tx_status = 'PROCESSING'
    AND tx_timestamp >= CURRENT_DATE - INTERVAL '1 hour'
ORDER BY tx_timestamp DESC 
LIMIT 190;

-- name: ConfirmTransaction :exec
UPDATE rampx_cross_chain_swaps 
SET tx_status = $1
WHERE transaction_hash = $2;