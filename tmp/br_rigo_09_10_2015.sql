-- MySQL dump 10.13  Distrib 5.5.43, for debian-linux-gnu (x86_64)
--
-- Host: localhost    Database: br_rigo
-- ------------------------------------------------------
-- Server version	5.5.43-0ubuntu0.14.04.1

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `articles`
--

DROP TABLE IF EXISTS `articles`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `articles` (
  `ArticleID` int(11) NOT NULL AUTO_INCREMENT,
  `Title` longtext NOT NULL,
  `content` varchar(255) NOT NULL DEFAULT '',
  `Author` longtext NOT NULL,
  `CreatedAt` varchar(19) NOT NULL DEFAULT '',
  `UpdatedAt` varchar(19) NOT NULL DEFAULT '',
  PRIMARY KEY (`ArticleID`)
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `articles`
--

LOCK TABLES `articles` WRITE;
/*!40000 ALTER TABLE `articles` DISABLE KEYS */;
INSERT INTO `articles` VALUES (1,'你只是看起来很努力','<p>读李尚龙的《你只是看起来很努力》，读到三分之一的时候，我抱着书沉思，真的，我只是看起来很努力而已。我平常总是对家人说，我已经很努力了，生活于我也许只能就这样子了，我只能是这样的一个结果吧。可是现在，我门心自问：我真的不够努力。我用努力安慰自己的时候，我用努力敷衍家人的时候，我放慢自己进步的脚步的时候，我拖延自己梦想的时候，我真的真的太不努力了。</p>','张小龙','2015-09-28 08:00:01',''),(2,'成为一个无可替代的人','<p>里行间中，我看到一个倔强并目标坚定的男孩儿，不抛弃不放弃，为了自己前途努力拼搏。我能看见他的坚持，收获了应有的荣誉。我看见了他的努力，拍出了受欢迎的影片。不得不说，我喜欢这样一个坚持并且有梦想的李尚龙。</p>','小雨 ','2015-09-28 08:00:01',''),(3,'从你的全世界路过','<p>我不认识张嘉佳。 我一朋友认识。几年前，他跟我说，咪蒙，我一校友，跟你一样，写些乱七八糟的文章，有点小才情、小趣味以及小猥琐。 反正我就当这是夸我了，从那时候开始，我就下意识地，把张嘉佳列为跟我同类了。 可见我不要脸也不是一天两天了。 然后，我断断续续听到他的二手消息，当了电影编剧，被金马奖提名。找了个女人，分了。开了个酒吧，赔了；然后丫的又穷困潦倒了。</p>','张嘉佳 ','2015-09-28 08:00:01',''),(4,'拿撒泼当个性、拿偷窥当炫耀、拿意淫当真枪、拿颓废当青春','<p>实在点说，这本书有不错的几篇，看得出微博达人的想象力与汉字（语）拼接组合能力。但泼点冷水，做为一本书，而且是号称让所有人心动的短篇小说集来讲，就实在不敢恭维了。看了几篇就看不下去了，故事生编硬造痕迹严重，行文语言句型重复，思维点似跳跃，粗暴插入的微博金句，随处可见为博眼球而看似随意的刻意，我想，能心动的可能是那些精力旺盛的过剩、睡前需要前戏、需要不断刷新微博的微博控吧！就内容与形式而言，我也没看到什么宣传中正能量，而且觉得这个词与这本书也不搭调儿。</p>','sysjc','2015-09-28 08:00:01',''),(5,'如何写一本阅读量过亿的烂书','<p>1.你一定要有丰富的感情经历。不管你是当电灯泡还是当备胎，还是傲娇地拒绝别人，你如果没有上十次，你都不好意思出来见人吧。什么，真的没有？没事，这些可以编啊。如果只有5个，那你可以换个名字，不同体位，都来一遍。这样就什么都有了，看的人也会跟着高潮的</p><p>2.你一定要学会造句。特别是这样的：云和海分居两地，人和人相聚别离。如果你不会，我来教你：</p><p>4.一定要写爱情，最好有小三、出轨和浪子回头</p>','梦里不知身是客','2015-09-28 08:00:01',''),(6,'在所有人事已非的景色里，我最喜欢你','<p>前几天，有人问过我，你觉得大师和一般的写作者有什么不同？我想了很久，打了一个比方，跟大师相比，其他的写作者就像刚刚入门的厨子，总偏好鲍参翅肚这类昂贵食材，用无数花哨的手法来展现高难度的烹饪技艺，选极其精巧奢华的杯盘碟盏来装盘，一招一式都是卖弄；大师呢，经历过这个阶段之后，返璞归真，往往喜欢做个最平凡的家常菜，但是就算是个西红柿炒蛋都能给你感动的涕泪横流，而且，你会觉得他做的菜，是如此亲切，就像小时候坐在厨房里看着妈妈系着围裙忙活出来的菜，普普通通，却是让你一生难忘。同理，只有三流的小说家才喜欢用虐心','女人的勇气常被误认为是疯狂','2015-09-28 08:00:01',''),(7,'突然就走到了西藏','<p>这不只是一本明星写的带有自传色彩的书，更是一个在世间行走的人老老实实的心灵告白，你可以读到他的挣扎，他的茫然和转化。这本书也不仅关乎心灵，更关乎现世人生，关乎如何看待这个世界，看待这个无常世界里的命运变幻。</p>','陈坤','2015-09-28 08:00:01',''),(13,'what','fuck','archer','2015-09-28 08:00:01',''),(14,'test update func','<p>testing update functionality</p>','archer','2015-09-28 08:00:01',''),(15,'谁的青春不迷茫 读书','<p>对于心灵鸡汤，人们说那是正确的废话，人们喜欢在微信等传播媒体上看之类的文章，以为很有启发，事实上这些远不如看书籍。书籍存在了上千年，当PDF等电子文档疯狂流行的时候，人们担心纸质书籍会走向灭亡，可是现在来看，并没有灭亡，我个人还是喜欢使用纸质书籍，看着更有感觉而已。</p><p>今天看了周鸿祎写的自传，里面主要将公司的一些事情，觉得太有道理，要是这样的公司文化，果断去了。</p>','特立独行的猪','2015-09-28 08:00:01','');
/*!40000 ALTER TABLE `articles` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `users` (
  `UserID` int(11) NOT NULL AUTO_INCREMENT,
  `UserName` longtext NOT NULL,
  `Password` longtext NOT NULL,
  `CreatedAt` varchar(19) NOT NULL DEFAULT '',
  `UpdatedAt` varchar(19) NOT NULL DEFAULT '',
  PRIMARY KEY (`UserID`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (2,'archer','archer','2015-09-28 08:00:01','');
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

-- Dump completed on 2015-10-08 23:41:34
