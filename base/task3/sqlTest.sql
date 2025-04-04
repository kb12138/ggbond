CREATE TABLE students
(
    id    INT PRIMARY KEY AUTO_INCREMENT, -- MySQL 自增语法（PostgreSQL 用 SERIAL）
    name  VARCHAR(50) NOT NULL,           -- 姓名长度限制
    age   TINYINT UNSIGNED,               -- 年龄合理范围 0-255（学生场景可缩小为 5-25）
    grade VARCHAR(20)                     -- 年级如 "三年级"、"高二"
);

-- 题目1：基本CRUD操作
-- 假设有一个名为 students 的表，包含字段 id （主键，自增）、 name （学生姓名，字符串类型）、 age （学生年龄，整数类型）、 grade （学生年级，字符串类型）。
-- 要求 ：
-- 编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
INSERT INTO students (name, age, grade)
VALUES ('张三', 20, '三年级');

-- 编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
SELECT id, name, age, grade -- 明确指定需要的字段（避免 SELECT *）
FROM students
WHERE age > 18 -- 条件过滤：年龄大于 18 岁
ORDER BY age DESC;

-- 编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
UPDATE students
SET grade = '四年级', -- 支持同时更新多个字段（如顺便修改年龄：age = 21）
    WHERE name = '张三';

-- 编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。
DELETE
FROM students
WHERE age < 15; -- 物理删除年龄 < 15 岁的记录



