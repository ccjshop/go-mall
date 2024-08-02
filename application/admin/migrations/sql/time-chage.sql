DELIMITER //
CREATE PROCEDURE ConvertDatetimeToTimestamp(IN tableName VARCHAR(64), IN columnName VARCHAR(64))
BEGIN
    -- 创建一个新列来存储转换后的时间戳
    SET @sql = CONCAT('ALTER TABLE ', tableName, ' ADD ', columnName, '_new INT(10) NULL');
PREPARE stmt FROM @sql;
EXECUTE stmt;
DEALLOCATE PREPARE stmt;

-- 将非 null 的日期时间值转换为 Unix 时间戳（秒级）
SET @sql = CONCAT('UPDATE ', tableName, ' SET ', columnName, '_new = UNIX_TIMESTAMP(', columnName, ') WHERE ', columnName, ' IS NOT NULL');
PREPARE stmt FROM @sql;
EXECUTE stmt;
DEALLOCATE PREPARE stmt;

-- 将 null 值设置为 0
SET @sql = CONCAT('UPDATE ', tableName, ' SET ', columnName, '_new = 0 WHERE ', columnName, ' IS NULL');
PREPARE stmt FROM @sql;
EXECUTE stmt;
DEALLOCATE PREPARE stmt;

-- 删除旧的列
SET @sql = CONCAT('ALTER TABLE ', tableName, ' DROP COLUMN ', columnName);
PREPARE stmt FROM @sql;
EXECUTE stmt;
DEALLOCATE PREPARE stmt;

-- 将新列重命名为旧列名
SET @sql = CONCAT('ALTER TABLE ', tableName, ' CHANGE ', columnName, '_new ', columnName, ' INT(10)');
PREPARE stmt FROM @sql;
EXECUTE stmt;
DEALLOCATE PREPARE stmt;
END //
DELIMITER ;

-- datetime 转成秒级时间戳
CALL ConvertDatetimeToTimestamp('oms_order', 'delivery_time');