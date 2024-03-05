CREATE DATABASE IF NOT EXISTS course_enrollment

CREATE TABLE student(
    student_id INT PRIMARY KEY AUTO_INCREMENT,
    student_name VARCHAR(100) NOT NULL
);

CREATE TABLE professor(
    professor_id INT PRIMARY KEY AUTO_INCREMENT,
    professor_name VARCHAR(100) NOT NULL
);

CREATE TABLE course(
    course_id INT PRIMARY KEY AUTO_INCREMENT,
    course_title VARCHAR(200) NOT NULL,
    max_capacity INT NOT NULL
);

CREATE TABLE enrollment(
    enrollment_id INT PRIMARY KEY AUTO_INCREMENT,
    student_id INT NOT NULL,
    course_id INT NOT NULL,
    enrollment_date DATE NOT NULL,
    FOREIGN KEY (student_id) REFERENCES student(student_id),
    FOREIGN KEY (course_id) REFERENCES course(course_id)
);

CREATE TABLE teaching_assignment(
    assignment_id INT PRIMARY KEY AUTO_INCREMENT,
    professor_id INT NOT NULL,
    course_id INT NOT NULL,
    start_date DATE NOT NULL,
    FOREIGN KEY (professor_id) REFERENCES professor(professor_id),
    FOREIGN KEY (course_id) REFERENCES course(course_id)
);

INSERT INTO student (student_name) 
VALUES  ('John Doe'),
        ('Alice Smith'),
        ('Bob Johnson'),
        ('Emily Davis'),
        ('Michael Brown');

INSERT INTO professor (professor_name) 
VALUES  ('Dr. Smith'),
        ('Prof. Johnson'),
        ('Dr. Anderson'),
        ('Prof. Williams'),
        ('Dr. Garcia');

INSERT INTO course (course_title, max_capacity) 
VALUES  ('Mathematics', 30),
        ('History', 25),
        ('Physics', 20),
        ('Computer Science', 35),
        ('Literature', 25);

INSERT INTO enrollment (student_id, course_id, enrollment_date) 
VALUES  (1, 1, '2024-02-29'),
        (2, 3, '2024-02-29'),
        (3, 2, '2024-02-29'),
        (4, 5, '2024-02-29'),
        (5, 4, '2024-02-29');

INSERT INTO teaching_assignment (professor_id, course_id, start_date) 
VALUES  (1, 1, '2024-02-29'),
        (2, 3, '2024-02-29'),
        (3, 2, '2024-02-29'),
        (4, 5, '2024-02-29'),
        (5, 4, '2024-02-29');
