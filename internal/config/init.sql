CREATE TABLE payment
(
    id            SERIAL PRIMARY KEY,
    transaction   VARCHAR(50),
    request_id    VARCHAR(50),
    currency      VARCHAR(50),
    provider      VARCHAR(50),
    amount        INT,
    payment_dt    INT,
    bank          VARCHAR(100),
    delivery_cost INT,
    goods_total   INT,
    custom_fee    INT
);

CREATE TABLE delivery
(
    id      SERIAL PRIMARY KEY,
    name    VARCHAR(100),
    phone   VARCHAR(20),
    zip     VARCHAR(20),
    city    VARCHAR(100),
    address VARCHAR(255),
    region  VARCHAR(100),
    email   VARCHAR(100)
);

CREATE TABLE orders
(
    id                 SERIAL PRIMARY KEY,
    order_uid          VARCHAR(50),
    track_number       VARCHAR(50),
    entry              VARCHAR(50),
    locale             VARCHAR(10),
    internal_signature VARCHAR(50),
    customer_id        VARCHAR(50),
    delivery_service   VARCHAR(50),
    shardkey           VARCHAR(50),
    sm_id              INT,
    oof_shard          VARCHAR(50),
    delivery_id        INT NOT NULL,
    payment_id         INT NOT NULL,
    FOREIGN KEY (delivery_id) REFERENCES delivery (id),
    FOREIGN KEY (payment_id) REFERENCES payment (id)

);
CREATE TABLE item
(
    id           SERIAL PRIMARY KEY,
    chrt_id      INT,
    track_number VARCHAR(50),
    price        INT,
    rid          VARCHAR(50),
    name         VARCHAR(50),
    sale         INT,
    size         VARCHAR(50),
    total_price  INT,
    nm_id        INT,
    brand        VARCHAR(50),
    status       INT,
    order_id     INT REFERENCES orders (id)
);
