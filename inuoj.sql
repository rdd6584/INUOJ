-- MySQL dump 10.13  Distrib 8.0.18, for Win64 (x86_64)
--
-- Host: localhost    Database: inuoj
-- ------------------------------------------------------
-- Server version	8.0.18

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES UTF8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `authtokens`
--

DROP TABLE IF EXISTS `authtokens`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `authtokens` (
  `email` varchar(100) NOT NULL,
  `token` varchar(50) NOT NULL,
  `auth_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `count` int(11) NOT NULL DEFAULT '1'
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `authtokens`
--

LOCK TABLES `authtokens` WRITE;
/*!40000 ALTER TABLE `authtokens` DISABLE KEYS */;
INSERT INTO `authtokens` VALUES ('rdd573@naver.com','GD7396W2498GZ4255XLS5944S106659','2020-01-22 15:49:46',1),('rdd6584@gmail.com','ZC2846507078Z8384YQN75J70M86F87','2020-01-22 15:50:23',1),('201601598@inu.ac.kr','R76JLIH46456JG151156EP14562213','2020-01-22 15:57:45',1);
/*!40000 ALTER TABLE `authtokens` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `judge_q`
--

DROP TABLE IF EXISTS `judge_q`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `judge_q` (
  `subm_no` int(11) NOT NULL,
  `ori_no` int(11) NOT NULL,
  `lang` int(11) NOT NULL,
  PRIMARY KEY (`subm_no`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `judge_q`
--

LOCK TABLES `judge_q` WRITE;
/*!40000 ALTER TABLE `judge_q` DISABLE KEYS */;
/*!40000 ALTER TABLE `judge_q` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `posts`
--

DROP TABLE IF EXISTS `posts`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `posts` (
  `post_no` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(50) NOT NULL,
  `id` varchar(20) NOT NULL,
  `category` tinyint(4) NOT NULL,
  `notice` tinyint(1) NOT NULL DEFAULT '0',
  `prob_no` int(11) DEFAULT '0',
  `cmt_no` int(11) NOT NULL DEFAULT '0',
  `post_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`post_no`)
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `posts`
--

LOCK TABLES `posts` WRITE;
/*!40000 ALTER TABLE `posts` DISABLE KEYS */;
INSERT INTO `posts` VALUES (1,'치킨','orihehe',1,0,1,0,'2010-01-01 23:59:59'),(2,'치킨','orihehe',1,0,3,0,'2010-01-01 23:59:59'),(3,'this is notice','orihehe',1,1,0,0,'2011-01-01 23:59:59'),(4,'당근 훔쳐먹기 왜 되는 걸까요?','orihehe',2,0,2,0,'2020-01-23 20:39:12'),(5,'당근은 왜 훔쳐먹을까요?','rdd6584',3,0,2,39,'2020-01-23 23:30:14'),(6,'왜 안 되 정 렬','rdd6584',2,0,0,0,'2020-01-28 19:57:20'),(7,'상태','rdd6584',2,0,0,0,'2020-01-28 19:58:20'),(8,'ㅠ','rdd6584',2,0,0,0,'2020-01-28 19:59:47'),(9,'z','rdd6584',2,0,0,0,'2020-01-28 20:07:13'),(10,'융','rdd6584',2,0,0,0,'2020-01-28 20:20:30'),(11,'융융','rdd6584',2,0,0,0,'2020-01-28 20:20:40'),(12,'정렬','rdd6584',2,0,0,0,'2020-01-28 22:14:08'),(13,'q','rdd6584',2,0,0,0,'2020-01-29 00:05:08'),(14,'15번 글입니다','rdd6584',2,0,0,0,'2020-01-29 00:57:42'),(15,'16번 글입니다.','rdd6584',2,0,0,0,'2020-01-29 00:57:55'),(16,'17번 글입니다!','rdd6584',2,0,0,0,'2020-01-29 00:58:14'),(17,'??','rdd6584',2,0,0,0,'2020-01-29 00:58:48');
/*!40000 ALTER TABLE `posts` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `prob_auth`
--

DROP TABLE IF EXISTS `prob_auth`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `prob_auth` (
  `id` varchar(20) NOT NULL,
  `ori_no` int(11) NOT NULL,
  PRIMARY KEY (`id`,`ori_no`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `prob_auth`
--

LOCK TABLES `prob_auth` WRITE;
/*!40000 ALTER TABLE `prob_auth` DISABLE KEYS */;
INSERT INTO `prob_auth` VALUES ('ori',1),('ori',2),('ori',1234),('ori',1235),('ori',1236),('ori',1237),('ori',1238),('ori',1239),('ori',1240),('ori',1241),('ori',1242),('ori',1243),('ori',1244),('ori',1245),('ori',1246),('ori',1247),('rdd6584',3),('rdd6584',4),('rdd6584',5),('rdd6584',6),('rdd6584',7),('rdd6584',8),('rdd6584',9),('rdd6584',10),('rdd6584',11),('rdd6584',12),('rdd6584',13),('rdd6584',14),('rdd6584',15),('rdd6584',16),('rdd6584',17),('rdd6584',18),('rdd6584',19),('rdd6584',20),('rdd6584',21);
/*!40000 ALTER TABLE `prob_auth` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `probs`
--

DROP TABLE IF EXISTS `probs`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `probs` (
  `ori_no` int(11) NOT NULL AUTO_INCREMENT,
  `prob_no` int(11) DEFAULT NULL,
  `t_limit` int(11) NOT NULL DEFAULT '1000',
  `m_limit` int(11) NOT NULL DEFAULT '256',
  `attempt` int(11) NOT NULL DEFAULT '0',
  `accept` int(11) NOT NULL DEFAULT '0',
  `owner` varchar(20) DEFAULT '',
  `title` varchar(30) NOT NULL DEFAULT '',
  `stat` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`ori_no`),
  UNIQUE KEY `prob_no` (`prob_no`)
) ENGINE=InnoDB AUTO_INCREMENT=22 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `probs`
--

LOCK TABLES `probs` WRITE;
/*!40000 ALTER TABLE `probs` DISABLE KEYS */;
INSERT INTO `probs` VALUES (1,1,2000,256,2,1,'ori','지금 만나러 갑니다',2),(2,2,1000,256,1,1,'ori','당근 훔쳐먹기',2),(3,NULL,1000,256,0,0,'rdd6584','',0),(4,NULL,1000,256,0,0,'rdd6584','',0),(5,NULL,1000,256,0,0,'rdd6584','',0),(6,NULL,1000,256,0,0,'rdd6584','',0),(7,NULL,1000,256,0,0,'rdd6584','',0),(8,NULL,1000,256,0,0,'rdd6584','',0),(9,NULL,1000,256,0,0,'rdd6584','',0),(10,NULL,1000,256,0,0,'rdd6584','',0),(11,NULL,1000,256,0,0,'rdd6584','',0),(12,NULL,1000,256,0,0,'rdd6584','',0),(13,NULL,1000,256,0,0,'rdd6584','',0),(14,NULL,1000,256,0,0,'rdd6584','',0),(15,NULL,1000,256,0,0,'rdd6584','',0),(16,NULL,1000,256,0,0,'rdd6584','',0),(17,NULL,1000,256,0,0,'rdd6584','',0),(18,NULL,1000,256,0,0,'rdd6584','',0),(19,NULL,1000,256,0,0,'rdd6584','',0),(20,NULL,1000,256,0,0,'rdd6584','',0),(21,NULL,1000,256,0,0,'rdd6584','',0);
/*!40000 ALTER TABLE `probs` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `result_list`
--

DROP TABLE IF EXISTS `result_list`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `result_list` (
  `id` varchar(20) NOT NULL,
  `prob_no` int(11) NOT NULL,
  `result` tinyint(4) DEFAULT '0',
  PRIMARY KEY (`id`,`prob_no`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `result_list`
--

LOCK TABLES `result_list` WRITE;
/*!40000 ALTER TABLE `result_list` DISABLE KEYS */;
INSERT INTO `result_list` VALUES ('ori',1,1),('ori',2,1),('orihehe',1,1),('orihehe',2,1),('rdd6584',1,2);
/*!40000 ALTER TABLE `result_list` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `submits`
--

DROP TABLE IF EXISTS `submits`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `submits` (
  `subm_no` int(11) NOT NULL AUTO_INCREMENT,
  `id` varchar(20) NOT NULL,
  `prob_no` int(11) NOT NULL,
  `result` int(11) DEFAULT '0',
  `lang` int(11) NOT NULL,
  `run_time` int(11) DEFAULT '0',
  `memory` int(11) DEFAULT '0',
  `codelen` int(11) NOT NULL,
  `subm_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`subm_no`)
) ENGINE=InnoDB AUTO_INCREMENT=22 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `submits`
--

LOCK TABLES `submits` WRITE;
/*!40000 ALTER TABLE `submits` DISABLE KEYS */;
INSERT INTO `submits` VALUES (1,'ori',1,1,1,173,2208,751,'2020-01-21 17:20:41'),(2,'ori',1,6,1,0,0,2,'2020-01-21 17:29:53'),(3,'ori',2,6,1,0,0,6,'2020-01-21 17:33:48'),(4,'ori',2,1,1,221,1484,349,'2020-01-21 17:34:24'),(5,'rdd',2,6,2,0,0,3,'2020-01-21 20:25:22'),(6,'rdd',1,6,1,0,0,3,'2020-01-21 20:25:40'),(7,'rdd',1,6,1,0,0,26,'2020-01-21 20:25:54'),(8,'rdd',1,6,2,0,0,8,'2020-01-21 20:26:07'),(9,'rdd',1,6,2,0,0,7,'2020-01-21 20:26:16'),(10,'rdd',1,6,1,0,0,14,'2020-01-21 20:26:27'),(11,'rdd',1,6,1,0,0,16,'2020-01-21 20:26:39'),(12,'rdd',1,6,1,0,0,28,'2020-01-21 20:26:53'),(13,'rdd',1,6,1,0,0,16,'2020-01-21 20:27:17'),(14,'rdd',1,6,1,0,0,7,'2020-01-21 20:27:24'),(15,'rdd',1,6,1,0,0,2,'2020-01-21 20:27:35'),(16,'rdd',1,6,1,0,0,8,'2020-01-21 20:27:40'),(17,'ori',2,6,2,0,0,7,'2020-01-21 20:59:58'),(18,'orihehe',2,1,1,210,1484,349,'2020-01-22 17:13:58'),(19,'orihehe',2,1,1,233,1480,349,'2020-01-22 17:21:23'),(20,'orihehe',1,1,1,64,1476,856,'2020-01-22 17:23:02'),(21,'rdd6584',1,6,2,0,0,2,'2020-01-22 18:15:09');
/*!40000 ALTER TABLE `submits` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_info`
--

DROP TABLE IF EXISTS `user_info`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user_info` (
  `id` varchar(20) NOT NULL,
  `pr` varchar(50) DEFAULT '잘 부탁드립니다~^^*',
  `ac_count` int(11) NOT NULL DEFAULT '0',
  `wa_count` int(11) NOT NULL DEFAULT '0',
  `all_count` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_info`
--

LOCK TABLES `user_info` WRITE;
/*!40000 ALTER TABLE `user_info` DISABLE KEYS */;
INSERT INTO `user_info` VALUES ('orihehe','나는 다영야',2,0,3),('rdd6584','HuDDuck 가즈아~~~!!!',0,1,1);
/*!40000 ALTER TABLE `user_info` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `users` (
  `id` varchar(20) NOT NULL,
  `password` varchar(255) NOT NULL,
  `email` varchar(50) NOT NULL,
  `auth` tinyint(1) NOT NULL DEFAULT '0',
  `admin` tinyint(2) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES ('orihehe','5f1059ff008c294b854f44e80bf29af7794725e7f798f92a30e82d80c4f0cf62','anda8658@inu.ac.kr',1,2),('rdd6584','9a91dcb67349fe3ee7ce9a5b56f212e0a685587fc7654552cce56e3a9cd17264','201501539@inu.ac.kr',1,2);
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2020-01-30  3:31:59
