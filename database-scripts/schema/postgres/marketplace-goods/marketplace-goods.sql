CREATE TABLE IF NOT EXISTS goods (
  id varchar(255) NOT NULL,
  user_id varchar(255) NOT NULL,
  name varchar(255) NOT NULL,
  description varchar(255) NOT NULL,
  price BIGINT NOT NULL,
  status varchar(255) NOT NULL,
  image_url varchar(1024) NOT NULL,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
);

CREATE INDEX IF NOT EXISTS idx_goods_user_id ON goods (user_id);
CREATE INDEX IF NOT EXISTS idx_goods_id ON goods (id);