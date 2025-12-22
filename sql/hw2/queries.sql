
SELECT full_name, gender, birth_date 
FROM students 
WHERE gender = 'female' 
ORDER BY birth_date DESC;

ALTER TABLE students ADD COLUMN student_gpa VARCHAR(50);


ALTER TABLE students DROP COLUMN student_gpa;