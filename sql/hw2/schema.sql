CREATE TABLE faculties (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL
);

CREATE TABLE student_groups (
    id SERIAL PRIMARY KEY,
    group_name VARCHAR(50) NOT NULL,
    faculty_id INT REFERENCES faculties(id)
);

CREATE TABLE students (
    id SERIAL PRIMARY KEY,
    full_name VARCHAR(100) NOT NULL,
    gender VARCHAR(10),
    birth_date DATE,
    group_id INT REFERENCES student_groups(id)
);


CREATE TABLE schedule (
    id SERIAL PRIMARY KEY,
    subject_name VARCHAR(100),
    time_slot VARCHAR(50), 
    group_id INT REFERENCES student_groups(id)
);

