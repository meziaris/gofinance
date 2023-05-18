CREATE TABLE IF NOT EXISTS transactions (
    id BIGSERIAL PRIMARY KEY NOT NULL,
    user_id BIGINT NOT NULL,
    transaction_category_id BIGINT NOT NULL,
    transaction_type_id BIGINT NOT NULL,
    currency_id BIGINT NOT NULL,
    transaction_date TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    notes text,
    amount INT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ,
    FOREIGN KEY(user_id) REFERENCES users(id),
    FOREIGN KEY(transaction_category_id) REFERENCES transaction_categories(id),
    FOREIGN KEY(transaction_type_id) REFERENCES transaction_types(id),
    FOREIGN KEY(currency_id) REFERENCES currencies(id)
);
