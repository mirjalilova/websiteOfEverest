CREATE TABLE teachers (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    experience_years INT,
    ielts_score NUMERIC(1, 1),
    contact VARCHAR(50) NOT NULL,
    profile_picture_url VARCHAR(255),
    graduated_students INT,
    bio TEXT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW() ON UPDATE CURRENT_TIMESTAMP,
    deleted_at BIGINT DEFAULT 0
);

CREATE TABLE courses (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    duration FLOAT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW() ON UPDATE CURRENT_TIMESTAMP,
    deleted_at BIGINT DEFAULT 0
);

CREATE TABLE course_items (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    course_id UUID REFERENCES courses(id),
    description VARCHAR(100) NOT NULL,
    price FLOAT,
    days_per_week INT,
    lesson_hours FLOAT,
    duration_weeks FLOAT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW() ON UPDATE CURRENT_TIMESTAMP,
    deleted_at BIGINT DEFAULT 0
);

CREATE TABLE dashboard (
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

CREATE TABLE gallery(
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    picture_url VARCHAR(255),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW() ON UPDATE CURRENT_TIMESTAMP,
    deleted_at BIGINT DEFAULT 0
);

CREATE TABLE branches(
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description VARCHAR(255),
    google_url TEXT,
    yandex_url TEXT,
    contact VARCHAR(50) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW() ON UPDATE CURRENT_TIMESTAMP,
    deleted_at BIGINT DEFAULT 0
);

CREATE TABLE certificates(
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    name VARCHAR (50),
    ielts_score NUMERIC(2, 1),
    cefr_level VARCHAR(2),
    description VARCHAR(255),
    certificate_url VARCHAR(255),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW() ON UPDATE CURRENT_TIMESTAMP,
    deleted_at BIGINT DEFAULT 0
);
