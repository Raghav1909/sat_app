-- name: CreateStudent :exec
INSERT INTO students (name, address, city, country, pincode, sat_score, passed)
VALUES (?, ?, ?, ?, ?, ?, ?);

-- View all data from the students table
-- name: GetAllStudents :many
SELECT 
    name, 
    address, 
    city, 
    country, 
    pincode, 
    sat_score, 
    CASE passed WHEN 1 THEN true ELSE false END AS passed
FROM students;

-- name: GetStudentByName :one
SELECT name, address, city, country, pincode, sat_score, passed
FROM students
WHERE name = ? LIMIT 1;

-- name: GetStudentRank :one
SELECT COUNT(*) + 1 AS student_rank
FROM students
WHERE students.sat_score > (SELECT students.sat_score FROM students WHERE students.name = ?);

-- name: UpdateStudentScore :exec
UPDATE students
SET sat_score = ?, passed = ?
WHERE students.name = ?;

-- name: DeleteStudent :exec
DELETE FROM students WHERE students.name = ?;

