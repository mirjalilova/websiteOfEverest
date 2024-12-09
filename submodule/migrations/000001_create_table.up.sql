CREATE TABLE IF NOT EXISTS teachers (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    experience_years VARCHAR(100) NOT NULL,
    ielts_score NUMERIC(2, 1),
    contact VARCHAR(50) NOT NULL,
    profile_picture_url VARCHAR(255),
    graduated_students VARCHAR(100),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at BIGINT DEFAULT 0
);

CREATE TABLE IF NOT EXISTS courses  ( 
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    duration VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at BIGINT DEFAULT 0
);

CREATE TABLE IF NOT EXISTS course_items  (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    course_id UUID REFERENCES courses(id) NOT NULL,
    description VARCHAR(255) NOT NULL,
    price FLOAT,
    days_per_week INT NOT NULL,
    lesson_hours FLOAT NOT NULL,
    week_days VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at BIGINT DEFAULT 0
);

CREATE TABLE IF NOT EXISTS dashboard  (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    average_experience_years FLOAT,
    branch_count INT,
    graduated_student_count INT,
    current_student_count INT,
    employee_count INT,
    ielts_trainers_count INT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at BIGINT DEFAULT 0
);

CREATE TABLE IF NOT EXISTS gallery (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    picture_url VARCHAR(255),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at BIGINT DEFAULT 0
);

CREATE TABLE IF NOT EXISTS branches  (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description VARCHAR(255),
    google_url TEXT,
    yandex_url TEXT,
    contact VARCHAR(50) NOT NULL,
    img_url TEXT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at BIGINT DEFAULT 0
);

CREATE TABLE IF NOT EXISTS certificates  (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    name VARCHAR (50) NOT NULL,
    ielts_score NUMERIC(2, 1) NOT NULL,
    cefr_level VARCHAR(50) NOT NULL,
    certificate_url VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at BIGINT DEFAULT 0
);
