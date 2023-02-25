DROP TABLE item;
DROP TABLE orders;
DROP TABLE delivery;
DROP TABLE payment;

CREATE TABLE payment
(
    id            SERIAL PRIMARY KEY,
    transaction   VARCHAR,
    request_id    VARCHAR,
    currency      VARCHAR,
    provider      VARCHAR,
    amount        INT,
    payment_dt    INT,
    bank          VARCHAR,
    delivery_cost INT,
    goods_total   INT,
    custom_fee    INT
);

CREATE TABLE delivery
(
    id      SERIAL PRIMARY KEY,
    name    VARCHAR,
    phone   VARCHAR,
    zip     VARCHAR,
    city    VARCHAR,
    address VARCHAR,
    region  VARCHAR,
    email   VARCHAR
);

CREATE TABLE orders
(
    id                 SERIAL PRIMARY KEY,
    order_uid          VARCHAR,
    track_number       VARCHAR,
    entry              VARCHAR,
    locale             VARCHAR,
    internal_signature VARCHAR,
    customer_id        VARCHAR,
    delivery_service   VARCHAR,
    shardkey           VARCHAR,
    sm_id              INT,
    oof_shard          VARCHAR,
    delivery_id        INT,
    payment_id         INT,
    FOREIGN KEY (delivery_id) REFERENCES delivery (id),
    FOREIGN KEY (payment_id) REFERENCES payment (id)

);
CREATE TABLE item
(
    id           SERIAL PRIMARY KEY,
    chrt_id      INT,
    track_number VARCHAR(255),
    price        INT,
    rid          VARCHAR(255),
    name         VARCHAR(255),
    sale         INT,
    size         VARCHAR,
    total_price  INT,
    nm_id        INT,
    brand        VARCHAR,
    status       INT,
    order_id     INT REFERENCES orders (id)
);

INSERT INTO orders(order_uid, track_number)
VALUES ('jfvdjn', 'jfdsbgjkdf'),
       ('jkjfdnsv', 'sbdfgvfgd');