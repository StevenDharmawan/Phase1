CREATE DATABASE IF NOT EXISTS student_courses

CREATE TABLE students (
    student_id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255),
    major VARCHAR(255) NOT NULL,
    year_of_study INT NOT NULL
);

CREATE TABLE courses(
	course_id INT PRIMARY KEY AUTO_INCREMENT,
    title VARCHAR(255) NOT NULL,
    instructor VARCHAR(255),
    `schedule` DATETIME NOT NULL,
    credit_hours INT
);

CREATE TABLE student_courses(
    student_id INT,
    course_id INT,
    preferred_schedule DATETIME,
    FOREIGN KEY (student_id) REFERENCES students(student_id),
    FOREIGN KEY (course_id) REFERENCES courses(course_id)
);

INSERT INTO courses(title, instructor, major, year_of_study)
VALUES 	('Benny Jax', 'Benny@gmail.com', 'Major A', 2),
        ('Harry Fen', 'Harry@gmail.com', 'Major C', 1),
        ('Ken Baran', 'Ken@gmail.com', 'Major D', 1),
        ('Merry Marion ', 'Merry@gmail.com', 'Major B', 3),
        ('Hendi Maxwell', 'Hendi@gmail.com', 'Major A', 4)

INSERT INTO courses(title, instructor, `schedule`, credit_hours)
VALUES 	('Title A', 'John Smith', NOW(), 10),
        ('Title E', 'John Smith', NOW(), 9),
        ('Title D', 'Jeremy Jean', NOW(), 11),
        ('Title C', 'Welby Bran', NOW(), 13),
        ('Title B', 'Welby Bran', NOW(), 14)

INSERT INTO student_courses(student_id, course_id, preferred_schedule)
VALUES 	(1, 1, NOW()),
        (2, 3, NOW()),
        (3, 4, NOW()),
        (4, 5, NOW()),
        (5, 1, NOW()),
        (1, 2, NOW()),
        (2, 4, NOW()),
        (3, 3, NOW()),
        (4, 1, NOW()),
        (5, 5, NOW())

