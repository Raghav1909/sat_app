-- name: CreateStudent :exec
INSERT INTO students (name, address, city, country, pincode, sat_score, passed)
VALUES (?, ?, ?, ?, ?, ?, ?);

-- View all data from the students table
-- name: GetAllStudents :many
SELECT * FROM students;

-- Get the rank of a student by SAT score (higher SAT score = better rank)
-- name: GetStudentRank :one
SELECT COUNT(*) + 1 AS student_rank
FROM students
WHERE students.sat_score > (SELECT students.sat_score FROM students WHERE students.name = ?);

-- Update the SAT score and passed status for a student by name
-- name: UpdateStudentScore :exec
UPDATE students
SET sat_score = ?, passed = ?
WHERE students.name = ?;

-- Delete a student record by name
-- name: DeleteStudent :exec
DELETE FROM students WHERE students.name = ?;

