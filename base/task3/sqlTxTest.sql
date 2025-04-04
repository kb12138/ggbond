START TRANSACTION;
-- 开启事务

-- 1. 锁定账户A（防幻读）并检查余额
SELECT balance
INTO @a_balance -- 将账户A余额存入用户变量
FROM accounts
WHERE id = 1001 -- 假设账户A的ID为1001
    FOR UPDATE;
-- 加行级锁（避免其他事务同时修改）

-- 2. 检查余额是否足够（账户存在且余额 ≥ 100）
IF
@a_balance IS NULL THEN  -- 账户A不存在
    SIGNAL SQLSTATE '45000'  -- 自定义错误码
    SET MESSAGE_TEXT = '转出账户不存在';
ELSEIF
@a_balance < 100 THEN  -- 余额不足
    SIGNAL SQLSTATE '45001'
    SET MESSAGE_TEXT = '转出账户余额不足';
END IF;

-- 3. 执行转账操作（原子更新）
UPDATE accounts
SET balance = balance - 100
WHERE id = 1001; -- 扣减账户A
UPDATE accounts
SET balance = balance + 100
WHERE id = 1002;
-- 增加账户B（假设ID为1002）

-- 4. 记录交易日志（含事务ID和时间戳）
INSERT INTO transactions (from_account_id, to_account_id, amount, create_time)
VALUES (1001, 1002, 100, NOW());

-- 5. 提交事务（所有操作成功）
COMMIT;