
INSERT INTO faculties (name) VALUES ('Engineering'), ('Humanities');


INSERT INTO student_groups (group_name, faculty_id) VALUES 
('CSCI 231', 1), ('CSCI 235', 1), ('MATH 231', 2), ('MATH 162', 2), ('PHYS 162', 1);

INSERT INTO students (full_name, gender, birth_date, group_id) VALUES 
('Adelya Mamyr', 'female', '2004-05-12', 1), ('Dayana Yergaliyeva', 'female', '2003-11-20', 1),
('Dilnaz Amrasheva', 'female', '2005-01-15', 1), ('Madi Shayakhmetov', 'male', '2004-03-12', 1),
('Muslima Sandybay', 'female', '2003-07-22', 1), ('Sultanali Saken', 'male', '2004-02-10', 1),
('Yernar Duzelbay', 'male', '2002-12-05', 2), ('Zhaina Zakirkhan', 'female', '2003-09-18', 2),
('Aidana Ispayeva', 'female', '2004-06-30', 2), ('Aishabibi Dauytova', 'female', '2002-08-14', 2),
('Akbota Bolat', 'female', '2004-05-12', 2), ('Ayana Bakirova', 'female', '2003-04-20', 2),
('Sara Alen', 'female', '2002-11-30', 3), ('Zhansaya Syzdykova', 'female', '2004-01-10', 3),
('Dias Tagayev', 'male', '2003-04-12', 3), ('Gleb Vassyutinskiy', 'male', '2004-01-10', 3),
('Olzhas Mussalimov', 'male', '2002-11-30', 3), ('Ualikhan Yertayev', 'male', '2003-05-25', 3),
('Yerassyl Auyeskhan', 'male', '2004-08-15', 4), ('Assemay Bakhytzhan', 'female', '2003-12-12', 4),
('Zhassyn Yegnay', 'male', '2004-03-03', 4), ('Alina Alibekova', 'female', '2003-09-09', 4),
('Aliya Bissekeyeva', 'female', '2002-07-07', 4), ('Amina Amirova', 'female', '2004-11-11', 4),
('Dauren Apas', 'male', '2003-02-02', 5), ('Gauhar Assanbay', 'female', '2004-05-05', 5),
('Iliyas Zharylgassin', 'male', '2003-10-10', 5), ('Karina Shilnikova', 'female', '2002-01-01', 5),
('Temirlan Kaskabassov', 'male', '2004-04-04', 5), ('Olzhas Mussalimov', 'male', '2003-06-06', 5);

INSERT INTO schedule (subject_name, time_slot, group_id) VALUES  ('PE', 'Morning (09:00)', 3),('PE', 'Morning (09:00)', 4),
('PE', 'Afternoon (15:00)', 1),('PE', 'Afternoon (15:00)', 2),('PE', 'Afternoon (15:00)', 5);