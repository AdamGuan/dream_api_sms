-- --------------------------------------------------------
-- 主机:                           127.0.0.1
-- 服务器版本:                        5.5.40-0ubuntu1 - (Ubuntu)
-- 服务器操作系统:                      debian-linux-gnu
-- HeidiSQL 版本:                  9.1.0.4867
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;

-- 导出 dream_api_sms 的数据库结构
DROP DATABASE IF EXISTS `dream_api_sms`;
CREATE DATABASE IF NOT EXISTS `dream_api_sms` /*!40100 DEFAULT CHARACTER SET utf8 */;
USE `dream_api_sms`;


-- 导出  表 dream_api_sms.t_config_pkg 结构
DROP TABLE IF EXISTS `t_config_pkg`;
CREATE TABLE IF NOT EXISTS `t_config_pkg` (
  `F_pkg` varchar(250) NOT NULL COMMENT '包名',
  `F_app_name` varchar(250) NOT NULL COMMENT '包对应的应用名字',
  `F_app_id` varchar(250) NOT NULL COMMENT 'leancloud对应的app id',
  `F_app_key` varchar(250) NOT NULL COMMENT 'leancloud对应的app key',
  `F_app_master_key` varchar(250) NOT NULL COMMENT 'leancloud对应的master key',
  `F_app_msm_template` varchar(250) NOT NULL COMMENT 'leancloud对应的短信模板名',
  UNIQUE KEY `F_pkg` (`F_pkg`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='包相关信息';

-- 正在导出表  dream_api_sms.t_config_pkg 的数据：~2 rows (大约)
DELETE FROM `t_config_pkg`;
/*!40000 ALTER TABLE `t_config_pkg` DISABLE KEYS */;
INSERT INTO `t_config_pkg` (`F_pkg`, `F_app_name`, `F_app_id`, `F_app_key`, `F_app_master_key`, `F_app_msm_template`) VALUES
	('abc', '刷题吧', '1ogxif29tbur554rh6n2m9yefhajgqkjqwspvr4lzu9rczxvn', '2qdmwrqh979waj4emidd0yh07jcu9xm5rz4vuqam1bt4lq0k', '06midcv0qs66lq3w4e8r7s7njngcd18t19wv53huegtga47s', 'template1'),
	('com.readboy.gaokao.debug', '刷题吧', '1ogxif29tbur554rh6n2m9yefhajgqkjqwspvr4lzu9rczxvn', '2qdmwrqh979waj4emidd0yh07jcu9xm5rz4vuqam1bt4lq0k', '06midcv0qs66lq3w4e8r7s7njngcd18t19wv53huegtga47s', 'template1');
/*!40000 ALTER TABLE `t_config_pkg` ENABLE KEYS */;


-- 导出  表 dream_api_sms.t_config_response 结构
DROP TABLE IF EXISTS `t_config_response`;
CREATE TABLE IF NOT EXISTS `t_config_response` (
  `F_response_no` smallint(5) NOT NULL COMMENT '响应code',
  `F_response_msg` varchar(50) NOT NULL COMMENT '响应信息'
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COMMENT='api的响应配置';

-- 正在导出表  dream_api_sms.t_config_response 的数据：10 rows
DELETE FROM `t_config_response`;
/*!40000 ALTER TABLE `t_config_response` DISABLE KEYS */;
INSERT INTO `t_config_response` (`F_response_no`, `F_response_msg`) VALUES
	(0, '成功'),
	(-1, '失败'),
	(-2, '已注册'),
	(-3, '密码不符合规则'),
	(-4, '没有注册'),
	(-5, '用户名或密码错误'),
	(-6, '签名错误'),
	(-7, '包名不存在'),
	(-8, '现有密码错误'),
	(-9, '密码错误');
/*!40000 ALTER TABLE `t_config_response` ENABLE KEYS */;


-- 导出  表 dream_api_sms.t_sms_rate 结构
DROP TABLE IF EXISTS `t_sms_rate`;
CREATE TABLE IF NOT EXISTS `t_sms_rate` (
  `F_action` char(32) NOT NULL COMMENT '动作，由(手机号码，包名，一起md5构成)',
  `F_last_timestamp` datetime NOT NULL COMMENT '时间',
  UNIQUE KEY `F_action` (`F_action`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COMMENT='记录短信发送的频率，用于限制短信的频繁发送，暂时的，会改为redis';

-- 正在导出表  dream_api_sms.t_sms_rate 的数据：0 rows
DELETE FROM `t_sms_rate`;
/*!40000 ALTER TABLE `t_sms_rate` DISABLE KEYS */;
/*!40000 ALTER TABLE `t_sms_rate` ENABLE KEYS */;


-- 导出  表 dream_api_sms.t_user 结构
DROP TABLE IF EXISTS `t_user`;
CREATE TABLE IF NOT EXISTS `t_user` (
  `F_user_name` varchar(50) NOT NULL COMMENT '用户名',
  `F_user_password` char(40) NOT NULL COMMENT '用户密码',
  `F_pkg` varchar(250) NOT NULL DEFAULT 'default' COMMENT '包名',
  UNIQUE KEY `F_user_name` (`F_user_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户表';

-- 正在导出表  dream_api_sms.t_user 的数据：~0 rows (大约)
DELETE FROM `t_user`;
/*!40000 ALTER TABLE `t_user` DISABLE KEYS */;
/*!40000 ALTER TABLE `t_user` ENABLE KEYS */;
/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IF(@OLD_FOREIGN_KEY_CHECKS IS NULL, 1, @OLD_FOREIGN_KEY_CHECKS) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
