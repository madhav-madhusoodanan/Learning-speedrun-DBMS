-- Create table for cross-chain swaps
CREATE TYPE TX_STATUS AS ENUM ('INITIATED', 'PROCESSING', 'COMPLETED', 'FAILED');

CREATE TABLE rampx_chain_details (
    chain_id BIGINT PRIMARY KEY,
    chain_name VARCHAR(30) NOT NULL,
    chain_logo_uri VARCHAR NOT NULL
);

CREATE TABLE rampx_token_details (
    token_id BIGSERIAL PRIMARY KEY,
    chain_id BIGINT NOT NULL,
    token_address VARCHAR(80) NOT NULL,
    FOREIGN KEY (chain_id) REFERENCES rampx_chain_details(chain_id)
);

CREATE TABLE rampx_cross_chain_swaps (
    swap_id BIGSERIAL PRIMARY KEY,
    source_token BIGINT NOT NULL,
    source_amount NUMERIC(78) NOT NULL,
    destination_token BIGINT NOT NULL,
    destination_amount NUMERIC(78) NOT NULL,
    dollar_value DECIMAL(20,2) NOT NULL,
    fee_dollar_value DECIMAL(20,2) NOT NULL,
    source_address VARCHAR(80) NOT NULL,
    destination_address VARCHAR(80) NOT NULL,
    tx_timestamp TIMESTAMP DEFAULT(now() at time zone 'utc'),
    tx_status TX_STATUS NOT NULL,
    transaction_hash VARCHAR(100) NOT NULL,
    FOREIGN KEY (source_token) REFERENCES rampx_token_details(token_id),
    FOREIGN KEY (destination_token) REFERENCES rampx_token_details(token_id)
);

-- Create statistics table to track daily volumes
CREATE TABLE rampx_statistics (
    stat_id BIGSERIAL PRIMARY KEY,
    stat_date DATE NOT NULL,
    total_dollar_volume DECIMAL(20,2) NOT NULL DEFAULT 0,
    fee_dollar_volume DECIMAL(20,2) NOT NULL DEFAULT 0,
    created_at TIMESTAMP DEFAULT(now() at time zone 'utc'),
    updated_at TIMESTAMP DEFAULT(now() at time zone 'utc')
);

-- Create unique index on date to ensure we have one record per day
CREATE UNIQUE INDEX idx_stat_date ON rampx_statistics(stat_date);

-- Create index on timestamp for better query performance
CREATE INDEX idx_timestamp ON rampx_cross_chain_swaps(tx_timestamp);

-- Create function to handle inserts into rampx_cross_chain_swaps
CREATE OR REPLACE FUNCTION insert_cross_chain_swap(
    source_chain_id BIGINT,
    source_token_address VARCHAR(80),
    source_amount NUMERIC(78),
    destination_chain_id BIGINT,
    destination_token_address VARCHAR(80),
    destination_amount NUMERIC(78),
    dollar_value DECIMAL(20,2),
    fee_dollar_value DECIMAL(20,2),
    source_address VARCHAR(80),
    destination_address VARCHAR(80),
    tx_status TX_STATUS,
    transaction_hash VARCHAR(100)
) RETURNS BIGINT AS $$
DECLARE
    source_token_id BIGINT;
    destination_token_id BIGINT;
    new_swap_id BIGINT;
BEGIN
    -- Get source token ID
    SELECT token_id INTO source_token_id
    FROM rampx_token_details
    WHERE chain_id = source_chain_id 
    AND token_address = source_token_address;

    IF source_token_id IS NULL THEN
        RAISE EXCEPTION 'Source token not found for chain_id % and address %', 
            source_chain_id, source_token_address;
    END IF;

    -- Get destination token ID
    SELECT token_id INTO destination_token_id
    FROM rampx_token_details
    WHERE chain_id = destination_chain_id 
    AND token_address = destination_token_address;

    IF destination_token_id IS NULL THEN
        RAISE EXCEPTION 'Destination token not found for chain_id % and address %', 
            destination_chain_id, destination_token_address;
    END IF;

    -- Insert the swap record
    INSERT INTO rampx_cross_chain_swaps (
        source_token,
        source_amount,
        destination_token,
        destination_amount,
        dollar_value,
        fee_dollar_value,
        source_address,
        destination_address,
        tx_status,
        transaction_hash
    ) VALUES (
        source_token_id,
        source_amount,
        destination_token_id,
        destination_amount,
        dollar_value,
        fee_dollar_value,
        source_address,
        destination_address,
        tx_status,
        transaction_hash
    ) RETURNING swap_id INTO new_swap_id;
    RETURN new_swap_id;
END;
$$ LANGUAGE plpgsql;

-- Create function to handle the trigger
CREATE OR REPLACE FUNCTION update_swap_statistics()
RETURNS TRIGGER AS $$
BEGIN
    -- Only proceed if status is being changed to 'COMPLETED'
    IF (TG_OP = 'INSERT' AND NEW.tx_status = 'COMPLETED') OR
       (TG_OP = 'UPDATE' AND NEW.tx_status = 'COMPLETED' AND OLD.tx_status != 'COMPLETED') THEN
        
        -- Insert or update statistics for the day
        INSERT INTO rampx_statistics (
            stat_date,
            total_dollar_volume,
            fee_dollar_volume,
            updated_at
        )
        VALUES (
            DATE(NEW.tx_timestamp),
            NEW.dollar_value,
            NEW.fee_dollar_value,
            now() at time zone 'utc'
        )
        ON CONFLICT (stat_date) DO UPDATE
        SET
            total_dollar_volume = rampx_statistics.total_dollar_volume + NEW.dollar_value,
            fee_dollar_volume = rampx_statistics.fee_dollar_volume + NEW.fee_dollar_value,
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

