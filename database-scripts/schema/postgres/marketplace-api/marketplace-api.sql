CREATE TABLE IF NOT EXISTS orders (
                                      external_id VARCHAR UNIQUE NOT NULL,
                                      client_id VARCHAR NOT NULL,
                                      order_id BIGINT,
                                      status VARCHAR,
                                      created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                      updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                      PRIMARY KEY (client_id, order_id)
);

CREATE INDEX IF NOT EXISTS orders_external_id_idx ON orders(external_id);