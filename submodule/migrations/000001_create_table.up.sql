-- Teachers Table
CREATE TABLE IF NOT EXISTS teachers (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    name JSONB NOT NULL CHECK (jsonb_typeof(name) = 'object'),
    experience_years VARCHAR(100) NOT NULL,
    ielts_score NUMERIC(2, 1),
    contact VARCHAR(50) NOT NULL,
    profile_picture_url VARCHAR(255),
    graduated_students VARCHAR(100),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at BIGINT DEFAULT 0
);

CREATE INDEX teachers_name_idx ON teachers USING GIN (name jsonb_path_ops);

-- Courses Table
CREATE TABLE IF NOT EXISTS courses (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    name JSONB NOT NULL CHECK (jsonb_typeof(name) = 'object'),
    duration JSONB NOT NULL CHECK (jsonb_typeof(duration) = 'object'),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at BIGINT DEFAULT 0
);

CREATE INDEX courses_name_idx ON courses USING GIN (name jsonb_path_ops);

-- Course Items Table
CREATE TABLE IF NOT EXISTS course_items (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    course_id UUID REFERENCES courses(id) NOT NULL,
    description JSONB NOT NULL CHECK (jsonb_typeof(description) = 'object'),
    price FLOAT,
    days_per_week INT NOT NULL,
    lesson_hours FLOAT NOT NULL,
    week_days JSONB NOT NULL CHECK (jsonb_typeof(week_days) = 'object'),
    picture_url VARCHAR(255),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at BIGINT DEFAULT 0
);

CREATE INDEX course_items_description_idx ON course_items USING GIN (description jsonb_path_ops);

-- Dashboard Table
CREATE TABLE IF NOT EXISTS dashboard (
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

-- Gallery Table
CREATE TABLE IF NOT EXISTS gallery (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    picture_url VARCHAR(255),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at BIGINT DEFAULT 0
);

-- Branches Table
CREATE TABLE IF NOT EXISTS branches (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    name JSONB NOT NULL CHECK (jsonb_typeof(name) = 'object'),
    google_url TEXT,
    yandex_url TEXT,
    contact VARCHAR(50) NOT NULL,
    img_url TEXT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at BIGINT DEFAULT 0
);

CREATE INDEX branches_name_idx ON branches USING GIN (name jsonb_path_ops);

-- Certificates Table
CREATE TABLE IF NOT EXISTS certificates (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    name JSONB NOT NULL CHECK (jsonb_typeof(name) = 'object'),
    ielts_score NUMERIC(2, 1) NOT NULL,
    cefr_level VARCHAR(50) NOT NULL,
    certificate_url VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at BIGINT DEFAULT 0
);

CREATE INDEX certificates_name_idx ON certificates USING GIN (name jsonb_path_ops);

-- Trigger for Default Language Validation
CREATE OR REPLACE FUNCTION validate_default_language()
RETURNS TRIGGER AS $$
BEGIN
    IF NOT (NEW.name ? 'en') THEN
        RAISE EXCEPTION 'Default language (en) is required in name JSONB field.';
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Attach Trigger to Tables with JSONB Fields
CREATE TRIGGER validate_teacher_name
BEFORE INSERT OR UPDATE ON teachers
FOR EACH ROW EXECUTE FUNCTION validate_default_language();

CREATE TRIGGER validate_course_name
BEFORE INSERT OR UPDATE ON courses
FOR EACH ROW EXECUTE FUNCTION validate_default_language();

CREATE TRIGGER validate_branch_name
BEFORE INSERT OR UPDATE ON branches
FOR EACH ROW EXECUTE FUNCTION validate_default_language();

CREATE TRIGGER validate_certificate_name
BEFORE INSERT OR UPDATE ON certificates
FOR EACH ROW EXECUTE FUNCTION validate_default_language();
