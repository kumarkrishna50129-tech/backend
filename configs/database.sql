
CREATE DATABASE protypist_db;

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password TEXT NOT NULL,
    is_premium BOOLEAN DEFAULT FALSE,
    subscription_until TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE -- Soft delete ke liye
);

-- Email par index taaki Login fast ho
CREATE INDEX idx_users_email ON users(email);

CREATE TABLE results (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    wpm INTEGER NOT NULL,
    accuracy DECIMAL(5,2) NOT NULL,
    language VARCHAR(20) NOT NULL, -- 'hindi' or 'english'
    mistakes INTEGER DEFAULT 0,
    time_taken INTEGER, -- seconds mein
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE
);

-- UserID par index taaki History load karne mein lag na aaye
CREATE INDEX idx_results_user_id ON results(user_id);


CREATE VIEW global_leaderboard AS
SELECT 
    u.name, 
    MAX(r.wpm) as top_speed, 
    AVG(r.accuracy) as avg_accuracy,
    COUNT(r.id) as tests_taken
FROM users u
JOIN results r ON u.id = r.user_id
GROUP BY u.id
ORDER BY top_speed DESC
LIMIT 100;