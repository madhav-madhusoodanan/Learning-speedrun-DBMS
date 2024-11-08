-- Create table for cross-chain swaps
CREATE TYPE TX_STATUS AS ENUM ('INITIATED', 'PROCESSING', 'COMPLETED', 'FAILED');

CREATE TABLE rampx_cross_chain_swaps (
    swap_id BIGSERIAL PRIMARY KEY,
    source_chain INTEGER NOT NULL,
    destination_chain INTEGER NOT NULL,
    source_token VARCHAR(70) NOT NULL,
    destination_token VARCHAR(70) NOT NULL,
    source_amount NUMERIC(78) NOT NULL,
    destination_amount NUMERIC(78) NOT NULL,
    dollar_value DECIMAL(20,2) NOT NULL,
    source_address VARCHAR(70) NOT NULL,
    destination_address VARCHAR(70) NOT NULL,
    tx_timestamp TIMESTAMP DEFAULT(now() at time zone 'utc'),
    tx_status TX_STATUS NOT NULL,
    transaction_hash VARCHAR(70) NOT NULL
);

-- Create statistics table to track daily volumes
CREATE TABLE rampx_swap_statistics (
    stat_id BIGSERIAL PRIMARY KEY,
    stat_date DATE NOT NULL,
    total_volume DECIMAL(20,2) NOT NULL DEFAULT 0,
    number_of_swaps INTEGER NOT NULL DEFAULT 0,
    created_at TIMESTAMP DEFAULT(now() at time zone 'utc'),
    updated_at TIMESTAMP DEFAULT(now() at time zone 'utc')
);

-- Create unique index on date to ensure we have one record per day
CREATE UNIQUE INDEX idx_stat_date ON rampx_swap_statistics(stat_date);

-- Create index on timestamp for better query performance
CREATE INDEX idx_timestamp ON rampx_cross_chain_swaps(tx_timestamp);

-- Create function to handle the trigger
CREATE OR REPLACE FUNCTION update_swap_statistics()
RETURNS TRIGGER AS $$
BEGIN
    -- Only proceed if status is being changed to 'COMPLETED'
    IF (TG_OP = 'INSERT' AND NEW.tx_status = 'COMPLETED') OR
       (TG_OP = 'UPDATE' AND NEW.tx_status = 'COMPLETED' AND OLD.tx_status != 'COMPLETED') THEN
        
        -- Insert or update statistics for the day
        INSERT INTO rampx_swap_statistics (
            stat_date,
            total_volume,
            number_of_swaps,
            updated_at
        )
        VALUES (
            DATE(NEW.tx_timestamp),
            NEW.dollar_value,
            1,
            now() at time zone 'utc'
        )
        ON CONFLICT (stat_date) DO UPDATE
        SET
            total_volume = rampx_swap_statistics.total_volume + NEW.dollar_value,
            number_of_swaps = rampx_swap_statistics.number_of_swaps + 1,
            updated_at = now() at time zone 'utc';
    END IF;
    
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Create trigger that fires after insert or update on the swaps table
CREATE TRIGGER trigger_update_swap_statistics
    AFTER INSERT OR UPDATE ON rampx_cross_chain_swaps
    FOR EACH ROW
    EXECUTE FUNCTION update_swap_statistics();

