CREATE TABLE verifications (
    id UUID PRIMARY KEY,
    user_id UUID,
    type NUMERIC,
    domain NUMERIC,
    token VARCHAR(16),
    expires_at DATE
);
