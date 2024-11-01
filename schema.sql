-- تفعيل امتداد uuid-ossp
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- جدول الحسابات
CREATE TABLE IF NOT EXISTS Users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- جدول المرضى
CREATE TABLE IF NOT EXISTS Patients (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    date_of_birth DATE NOT NULL,
    national_id VARCHAR(20) UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- جدول الفحوصات
CREATE TABLE IF NOT EXISTS Examinations (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    patient_id UUID REFERENCES Patients(id) ON DELETE CASCADE,
    examination_data JSONB NOT NULL, -- البيانات مشفرة
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- جدول المعلومات العامة
CREATE TABLE IF NOT EXISTS GeneralInfo (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID REFERENCES Users(id) ON DELETE CASCADE,
    full_name VARCHAR(200),
    date_of_birth DATE,
    address TEXT,
    phone_number VARCHAR(20),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- جدول العيادات
CREATE TABLE IF NOT EXISTS Clinics (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID REFERENCES Users(id) ON DELETE CASCADE,
    clinic_name VARCHAR(200) NOT NULL,
    address TEXT,
    phone_number VARCHAR(20),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


-- جدول الأدوار
CREATE TABLE IF NOT EXISTS Roles (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    role_name VARCHAR(100) NOT NULL,
    description TEXT
);

CREATE TABLE IF NOT EXISTS AccountRoles (
    user_id UUID REFERENCES Users(id) ON DELETE CASCADE,
    role_id UUID REFERENCES Roles(id) ON DELETE CASCADE,
    PRIMARY KEY (user_id, role_id)
);