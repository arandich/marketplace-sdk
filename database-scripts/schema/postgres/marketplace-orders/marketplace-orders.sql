CREATE TABLE IF NOT EXISTS orders (
  id varchar(255) NOT NULL,
  user_id varchar(255) NOT NULL,
  good_ids varchar(255) [] NOT NULL,
  status varchar(255) NOT NULL,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
);

CREATE INDEX IF NOT EXISTS idx_orders_user_id ON orders (user_id);
CREATE INDEX IF NOT EXISTS idx_orders_id ON orders (id);

