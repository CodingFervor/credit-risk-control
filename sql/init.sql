-- Credit Risk Control Schema
CREATE TABLE loan_applications (id BIGSERIAL PRIMARY KEY, user_id BIGINT, amount DECIMAL(12,2), term INT, purpose VARCHAR(256), status VARCHAR(20) DEFAULT 'pending', credit_score INT, risk_level VARCHAR(10), approver_id BIGINT, reject_reason TEXT, created_at TIMESTAMPTZ DEFAULT NOW());
CREATE TABLE credit_accounts (id BIGSERIAL PRIMARY KEY, user_id BIGINT UNIQUE, total_limit DECIMAL(12,2), used_limit DECIMAL(12,2) DEFAULT 0, available_limit DECIMAL(12,2), overdue_count INT DEFAULT 0, status VARCHAR(20) DEFAULT 'active');
CREATE TABLE repayment_plans (id BIGSERIAL PRIMARY KEY, loan_id BIGINT, period INT, due_date DATE, principal DECIMAL(12,2), interest DECIMAL(12,2), total DECIMAL(12,2), status VARCHAR(20) DEFAULT 'pending', paid_at TIMESTAMPTZ);
CREATE TABLE risk_rules (id BIGSERIAL PRIMARY KEY, name VARCHAR(100), category VARCHAR(20), condition JSONB, score_weight DECIMAL(5,2), is_enabled BOOLEAN DEFAULT TRUE);
CREATE TABLE blacklist (id BIGSERIAL PRIMARY KEY, user_id BIGINT, phone VARCHAR(20), id_card VARCHAR(128), device_id VARCHAR(128), reason TEXT, source VARCHAR(20) DEFAULT 'manual', created_at TIMESTAMPTZ DEFAULT NOW());
CREATE TABLE collection_tasks (id BIGSERIAL PRIMARY KEY, loan_id BIGINT, user_id BIGINT, overdue_days INT, amount DECIMAL(12,2), assignee_id BIGINT, status VARCHAR(20) DEFAULT 'pending', created_at TIMESTAMPTZ DEFAULT NOW());
CREATE TABLE risk_events (id BIGSERIAL PRIMARY KEY, user_id BIGINT, event_type VARCHAR(50), detail JSONB, risk_score INT, created_at TIMESTAMPTZ DEFAULT NOW());
