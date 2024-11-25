-- name: CreateSwapOrder :exec
SELECT insert_cross_chain_swap(
    source_chain_id := $1,
    source_token_address := $2,
    source_amount := $3,
    destination_chain_id := $4,
    destination_token_address := $5,
    destination_amount := $6,
    dollar_value := $7,
    fee_dollar_value := $8,
    source_address := $9,
    destination_address := $10,
    tx_status := 'COMPLETED',
    transaction_hash := $11
);

-- name: GetDailyVolume :many
SELECT 
    DATE(tx_timestamp) as trade_date,
    COUNT(*) as number_of_trades,
    SUM(dollar_value) as total_daily_volume,
    SUM(fee_dollar_value) as total_daily_fee_volume
FROM rampx_cross_chain_swaps
WHERE 
    tx_timestamp >= CURRENT_DATE - INTERVAL '30 days'
    AND status = 'COMPLETED'
GROUP BY DATE(tx_timestamp)
ORDER BY trade_date DESC;

-- name: ConfirmTransaction :exec
UPDATE rampx_cross_chain_swaps 
SET tx_status = $1
WHERE transaction_hash = $2;