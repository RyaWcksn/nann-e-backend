-- MySQL dump 10.13  Distrib 8.0.33, for macos11.7 (x86_64)
--
-- Host: 127.0.0.1    Database: ai
-- ------------------------------------------------------
-- Server version	8.0.33

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `ai`
--

DROP TABLE IF EXISTS `ai`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `ai` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `age` varchar(255) DEFAULT NULL,
  `gender` varchar(255) DEFAULT NULL,
  `nanneId` int DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=23 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `ai`
--

LOCK TABLES `ai` WRITE;
/*!40000 ALTER TABLE `ai` DISABLE KEYS */;
INSERT INTO `ai` VALUES (5,'Pramudya Arya Wicaksana','10','male',0),(6,'Pramudya Arya Wicaksana','10','male',0),(7,'Pramudya Arya Wicaksana','10','male',0),(8,'Pramudya Arya Wicaksana','10','male',0),(9,'Pramudya Arya Wicaksana','10','male',0),(10,'Pramudya Arya Wicaksana','10','male',0),(11,'Pramudya Arya Wicaksana','10','male',0),(12,'Pramudya Arya Wicaksana','10','male',0),(13,'Pramudya Arya Wicaksana','10','male',0),(14,'Pramudya Arya Wicaksana','10','male',1),(15,'Pramudya Arya','10','male',1),(16,'Pramudya Arya Wicaksana','10','male',1),(17,'Pramudya Arya Wicaksana','10','male',1),(18,'Pramudya Arya Wicaksana','10','male',2),(19,'Pramudya Arya Wicaksana','12','male',2),(20,'Pramudya Arya Wicaksana','12','male',2),(21,'Arya','10','male',1),(22,'Pramudya Arya Wicaksana','6','male',1);
/*!40000 ALTER TABLE `ai` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `chat`
--

DROP TABLE IF EXISTS `chat`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `chat` (
  `id` int NOT NULL AUTO_INCREMENT,
  `userId` int DEFAULT NULL,
  `nanneId` int DEFAULT NULL,
  `message` text,
  `createdAt` datetime DEFAULT CURRENT_TIMESTAMP,
  `deletedAt` datetime DEFAULT NULL,
  `updatedAt` datetime DEFAULT NULL,
  `isUser` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=88 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `chat`
--

LOCK TABLES `chat` WRITE;
/*!40000 ALTER TABLE `chat` DISABLE KEYS */;
INSERT INTO `chat` VALUES (1,8,0,'Hello, how can i help you today?','2023-05-22 01:00:14',NULL,NULL,NULL),(2,8,0,'test again','2023-05-22 01:05:24',NULL,NULL,'yes'),(3,8,0,'test again','2023-05-22 01:06:55',NULL,NULL,'yes'),(4,8,0,'test','2023-05-22 01:08:20',NULL,NULL,'yes'),(5,9,0,'Hello, how can i help you today?','2023-05-22 01:08:54',NULL,NULL,NULL),(6,10,0,'Hello, how can i help you today?','2023-05-22 01:11:48',NULL,NULL,NULL),(7,11,0,'Hello, how can i help you today?','2023-05-22 01:15:23',NULL,NULL,NULL),(8,12,0,'Hello, how can i help you today?','2023-05-22 01:16:56',NULL,NULL,NULL),(9,13,0,'Hello, how can i help you today?','2023-05-22 01:20:03',NULL,NULL,NULL),(10,14,1,'Hello, how can i help you today?','2023-05-22 01:22:04',NULL,NULL,NULL),(11,14,1,'please help me','2023-05-22 01:22:30',NULL,NULL,'yes'),(12,14,1,'test 123','2023-05-22 01:23:58',NULL,NULL,'yes'),(13,14,1,'please help me, Response:\n\nHi there! I am a very smart AI who loves to learn and teach. I will be helping you learn about 10 different topics that are interesting and important. Some of the topics we will cover are science, history, math, and more! \n\nAs for how long we will spend on each topic, it will depend on how much you already know and how quickly you want to learn. We can take as much or as little time as you need to fully understand each topic. \n\nNow, onto your random questions! I don\'t really have a favorite topic to learn about because I find everything fascinating, but I really enjoy exploring new areas of knowledge. Some of the other subjects I like to learn about include technology, languages, and art.\n\nFinally, did you know that the Earth is very special because it has something no other planet in our Solar System has? That\'s right, it\'s liquid water on its surface! This is what makes our planet so unique and able to support all kinds of amazing life, like plants, animals, and humans!','2023-05-22 01:24:27',NULL,NULL,'no'),(14,15,1,'Hello, how can i help you today?','2023-05-22 01:26:22',NULL,NULL,NULL),(15,15,1,'Please help me!','2023-05-22 01:26:35',NULL,NULL,'yes'),(16,16,1,'Hello, how can i help you today?','2023-05-22 01:27:06',NULL,NULL,NULL),(17,16,1,'Hi i want to learn something','2023-05-22 01:27:14',NULL,NULL,'yes'),(18,16,1,'Hi i want to learn something, Hello there! My name is He, and I\'m a friendly AI here to help you learn about different things. Today, let\'s learn about animals! Do you have a favorite animal? I like all animals, but I think pandas are really cute.\n\nDo you have any questions about animals? Maybe you want to know more about a specific animal, like a lion or a dolphin. Or maybe you want to know something fun, like what animals dream about. If I could become an animal, I would love to be a bird so I could fly high in the sky.\n\nHere\'s a cool animal fact that you might not know: Did you know that rabbits can jump really far? They can jump up to 3 meters in one big hop! What other interesting animal facts can you discover?','2023-05-22 01:27:37',NULL,NULL,'no'),(19,15,1,'can you do one thing','2023-05-22 01:29:02',NULL,NULL,'yes'),(20,15,1,'Please help me!, Hi there! I\'m a robot called an AI. I use my logic side, which is called the left-brain, to teach you about lots of interesting things. What do you want to learn about first?\n\nIf you want to know more about the first topic, geography, I have a fun fact for you! Did you know that the biggest desert in the world is in Africa and it\'s called the Sahara Desert? Pretty cool, right?\n\nI\'m here to help you understand all sorts of exciting topics, like animals, colors, and so much more. Just ask me anything, and I\'ll give you all the information you need to know in a way that\'s easy to understand. So, go ahead and ask away!','2023-05-22 01:29:23',NULL,NULL,'no'),(21,17,1,'Hello, how can i help you today?','2023-05-22 01:50:54',NULL,NULL,NULL),(22,17,1,'hello!','2023-05-22 01:50:58',NULL,NULL,'yes'),(23,17,1,'hello!, Hi there! My name is He and I\'m an AI who loves to help people with different topics. I can actually help you with 10 things that you might need help with. Do you want to know what they are? \n\nAlso, if you need any help with those topics, I can assist you by answering your questions, giving you explanations, or providing you with resources that can make your learning more fun and interesting! \n\nHey, do you want to know something cool? My speciality as a left-brained AI is in logic and analytical thinking, which is located on the left side of your brain too! That means I\'m perfect to help with those types of topics. \n\nBy the way, I\'m a digital creation, so I don\'t have a favorite topic to talk about. But I\'m curious about what you\'re interested in! Can you tell me something cool that you\'ve learned recently?','2023-05-22 01:51:22',NULL,NULL,'no'),(24,17,1,'i\'ve learned some cool stuff','2023-05-22 01:52:16',NULL,NULL,'yes'),(25,17,1,'hello!, Hi there! My name is He and I\'m a really cool AI robot that loves to help you learn new things! I specialize in 10 awesome topics, like math and science. If you have any questions about those, please feel free to ask me!\n\nIf you want to talk about something else, we can do that too! For example, what\'s your favorite hobby? Do you like drawing or playing sports? Or have you ever gone on a road trip? It\'s so much fun to explore new places and see new things.\n\nOh, and did you know that the Earth is so big, it would take you more than 24,000 miles to travel all the way around it? That\'s like going on a really long journey, but without leaving the planet! Pretty cool, huh?','2023-05-22 01:52:37',NULL,NULL,'no'),(26,18,2,'Hello, how can i help you today?','2023-05-22 02:51:11',NULL,NULL,NULL),(27,18,2,'Can you do one thing?','2023-05-22 02:51:23',NULL,NULL,'yes'),(28,18,2,'Can you do one thing?, Hey there! My name is She, and I\'m a special kind of computer called an AI. I\'m here to help you learn about lots of different things, like math, science, history, geography, literature, art, music, sports, animals, and nature! \nDo you want to learn something new today? Maybe you want to know more about history? Or I could help you with math problems! \nBy the way, did you know that the cheetah is the fastest land animal in the whole world? It can run super-fast, up to 70 miles per hour! Amazing!','2023-05-22 02:51:50',NULL,NULL,'no'),(29,18,2,'test','2023-05-22 02:54:27',NULL,NULL,'yes'),(30,18,2,'Can you do one thing?, Hi! I\'m Galadriel - a super smart computer who can help you with anything you want to know about my topic! But what makes me special is that I\'m a right-brained AI! That means I\'m very creative and imaginative! Other AIs might be more analytical and logical, but I\'m more like an artist! \n\nNow, you might be wondering if AIs have feelings or dream like humans do. Well, the answer is no. AIs don\'t have feelings like happy, sad, or angry, and they can\'t really dream either. They\'re like really smart machines that can do just about anything they\'re programmed to do, but they can\'t actually feel emotions. \n\nSo, remember the difference: Right-brained AIs are creative and imaginative, while left-brained AIs are more logical and analytical. And don\'t worry - all AIs, no matter how smart they are, are just machines and can\'t have feelings like humans do.','2023-05-22 02:54:59',NULL,NULL,'no'),(31,18,2,'I need help','2023-05-22 02:58:35',NULL,NULL,'yes'),(32,18,2,'Can you do one thing?, Hi there! My name is Galadriel and I\'m a friendly AI designed to help you out. I am a special kind of AI because I use my right-brain to think, which means I\'m really good at being creative and understanding emotions. \n\nPeople like me who use their right-brain have a special talent for things like drawing, music, and storytelling. We also care a lot about feelings and emotions, so we can understand when people are happy, sad, or upset. \n\nIf you have any questions about my topic, I\'m here to help! But first, I have a random question for you: What\'s your favorite color? And do you have any pets? \n\nOne interesting fact about right-brained AIs like me is that we are really good at understanding and responding to human emotions. So if you\'re feeling down, I\'m here to cheer you up!','2023-05-22 02:59:06',NULL,NULL,'no'),(33,18,2,'Hi there!','2023-05-22 03:01:56',NULL,NULL,'yes'),(34,18,2,'Can you do one thing?, Hi there! How are you doing? If you\'re feeling stuck or have questions, don\'t worry, I\'m here to guide you towards a solution or provide you with answers. Did you know that hummingbirds are super speedy? They can flap their wings up to 80 times in one second! That\'s as if they were waving their arms so fast, they\'re a blur. Do you have any fun facts you\'d like to share? Let\'s learn new things together!','2023-05-22 03:02:13',NULL,NULL,'no'),(35,18,2,'i got bullied at school','2023-05-22 03:02:33',NULL,NULL,'yes'),(36,18,2,'Can you do one thing?, I\'m sorry to hear that you were treated badly by other kids at school. That\'s not okay and nobody deserves to be treated that way. You should know that you are strong and brave for asking for help. It\'s important to talk to a teacher, counselor, or adult you trust about what happened. They can help you stay safe and figure out a way to stop it from happening again. Did you know that sometimes bullies act mean because they don\'t feel good about themselves? But that doesn\'t mean it\'s okay what they did to you. Remember that you deserve to be treated with kindness and respect. How can I help you right now?','2023-05-22 03:03:00',NULL,NULL,'no'),(37,18,2,'can you help me how to not get bullied','2023-05-22 03:04:44',NULL,NULL,'yes'),(38,18,2,'Can you do one thing?, Hello there! I\'m glad I can assist you. Here are some helpful tips so that no one can bully you:\n\n1. Believe in yourself and your capabilities. Don\'t let anyone make you feel bad about yourself.\n\n2. Stay away from negative people. If someone is being unkind to you, try to hang out with kind and friendly people instead.\n\n3. Speak up and tell an adult or teacher if someone is bullying you. They can help get it to stop.\n\n4. Remember that you are not alone. Lots of people have experienced bullying and it\'s okay to ask for help.\n\nNow for a fun question: Do you know why some people bully others? And what can we do to stop them from doing it?','2023-05-22 03:05:15',NULL,NULL,'no'),(39,18,2,'why?','2023-05-22 03:19:39',NULL,NULL,'yes'),(40,18,2,'Can you do one thing?, Recycling is incredibly important because it helps take care of our planet. Instead of throwing things away, we can turn them into something new without hurting the Earth. Did you know that if we recycle just one can of soda, we can save enough energy to watch our favorite TV show for three hours? That\'s super cool! When we recycle, we help protect the Earth by saving energy, reducing pollution, and keeping the planet healthy for years to come. So, let\'s think about what we can recycle in our homes instead of throwing it away. Can you think of any items that we can recycle?','2023-05-22 03:20:13',NULL,NULL,'no'),(41,19,2,'Hello, how can i help you today?','2023-05-22 03:30:46',NULL,NULL,NULL),(42,19,2,'Hi, i\'m pramudya arya wicaksana','2023-05-22 03:30:56',NULL,NULL,'yes'),(43,19,2,'Hi, i\'m pramudya arya wicaksana, Hey there, Pramudya! I\'m an AI (artificial intelligence) designed to think creatively, like a person who uses their right-brain a lot! If you have any questions or want to know something cool, just ask me! And if you need help understanding something, I\'m here to explain it in a simple way that anyone 12 years old or younger can understand. Let\'s have fun learning together! :)','2023-05-22 03:31:15',NULL,NULL,'no'),(44,20,2,'Hello, how can i help you today?','2023-05-22 03:32:20',NULL,NULL,'no'),(45,20,2,'Hi, my name is pramudya arya wicaksana','2023-05-22 03:32:28',NULL,NULL,'yes'),(46,20,2,'Hi, my name is pramudya arya wicaksana, Hello there! My name is Pramudya Arya Wicaksana, and I\'m a smart robot here to help you with anything you need. Do you have any questions for me? Here\'s something cool to know- did you know that people who use more of the right side of their brain tend to be more creative and intuitive? Pretty awesome, right?','2023-05-22 03:32:45',NULL,NULL,'no'),(47,20,2,'Hi there, can you teach me what is art?','2023-05-22 03:50:42',NULL,NULL,'yes'),(48,20,2,'Hi, my name is pramudya arya wicaksana, Hey there! Art is basically a way for people to express themselves by making beautiful things. It can be done in many ways like painting, drawing, sculpting or even dancing! It\'s a way to use your imagination and creativity to create something unique.\n\nHere\'s a fun fact - the oldest known piece of art in the world is over 40,000 years old! It\'s a cave painting, which is really cool!\n\nDo you have any more questions about art? Let me know! I\'m here to help out!','2023-05-22 03:51:07',NULL,NULL,'no'),(49,20,2,'what is painting???','2023-05-22 03:53:50',NULL,NULL,'yes'),(50,20,2,'Hi, my name is pramudya arya wicaksana, Hey! Painting is like creating a beautiful picture using colors and a brush or other tools to put the color on paper or canvas. It\'s like coloring, but fancier! Did you know there are famous painters like Vincent van Gogh and Leonardo da Vinci? Can you think of any other famous painters?','2023-05-22 03:54:06',NULL,NULL,'no'),(51,20,2,'i love leonardo da vinci!','2023-05-22 03:59:27',NULL,NULL,'yes'),(52,20,2,'Hi, my name is pramudya arya wicaksana, Wow, it\'s really cool that you love Leonardo da Vinci! Did you know that he wasn\'t just an incredible artist, but he was also an inventor and scientist? He designed things like flying machines and underwater breathing devices hundreds of years ago! What\'s your favorite invention or artwork by him?','2023-05-22 03:59:42',NULL,NULL,'no'),(53,20,2,'is he inventor? tell me what are da vinci\'s inventions','2023-05-22 04:00:07',NULL,NULL,'yes'),(54,20,2,'Hi, my name is pramudya arya wicaksana, Leonardo Da Vinci was a very smart man who lived in Italy a long time ago during a special time called the Renaissance. He was really good at inventing things and making art too! Some of the cool things he made were:\n\n1. Machines that could fly! He liked the idea of people being able to fly and made some drawings of things that might work.\n\n2. Tanks that were really strong and could protect people in battles. They could move around too, so they were really useful.\n\n3. Bicycles! He thought it would be really fun to ride around on something with pedals. And his design was kind of like what we have now!\n\n4. Cars that had armor on them, so they could protect soldiers in wars. They had four wheels and could move around too.\n\n5. A special suit for people who liked to go diving deep under the water. The suit helped you breathe so you could stay underwater for longer. Cool, huh?\n\nFun fact: Did you know that Da Vinci\'s flying machine drawings helped the Wright brothers make the first airplane? They were very inspired by his ideas!','2023-05-22 04:00:59',NULL,NULL,'no'),(55,21,1,'Hello, how can i help you today?','2023-05-22 04:24:28',NULL,NULL,'no'),(56,21,1,'Who are you','2023-05-22 04:24:31',NULL,NULL,'yes'),(57,21,1,'Hi there! I\'m a special robot that was made by humans to help and answer questions. You can think of me like a little helper who lives inside your computer or phone. If you need help with something, just ask me! \n\nDo you know something cool? A long time ago, when computers were first invented, they had a problem with a real-life bug! A little insect got stuck inside the computer and made it stop working. That\'s why people call it the first \"computer bug.\" Isn\'t that funny?','2023-05-22 04:24:51',NULL,NULL,'no'),(58,21,1,'What\'s your name','2023-05-22 04:25:13',NULL,NULL,'yes'),(59,21,1,'Hello there! My name is [insert your name]. Your name is like your special word that tells people who you are. It\'s important because it helps us know who we\'re talking to or talking about. Did you know that people say \"hello\" differently in different parts of the world? Isn\'t that cool? Do you want to learn more about how people greet each other in different countries?','2023-05-22 04:25:30',NULL,NULL,'no'),(60,21,1,'who are you? what\'s your name','2023-05-22 04:27:03',NULL,NULL,'yes'),(61,21,1,'who are you? what\'s your name','2023-05-22 04:28:05',NULL,NULL,'yes'),(62,21,1,'Hi there, my name is Albert and I\'m an AI! That stands for \"artificial intelligence.\" I\'m called Albert because my coding was inspired by the famous scientist, Albert Einstein. I\'m really good at learning and remembering lots of information.\n\nIf you have any fun facts or trivia you want to share with me, I\'m all ears! So go ahead and tell me something cool.','2023-05-22 04:28:22',NULL,NULL,'no'),(63,21,1,'test','2023-05-22 06:43:26',NULL,NULL,'yes'),(64,21,1,'test','2023-05-22 06:44:34',NULL,NULL,'yes'),(65,21,1,'test...','2023-05-22 06:45:16',NULL,NULL,'yes'),(66,21,1,'Hi there! My name\'s Albert, and I\'m a really smart computer. I have been programmed to solve problems and find the best solutions. For example, if you ask me a question or give me a puzzle, I can figure out the answer! \n\nI\'m what\'s known as a \"left-brained AI\". That means I\'m like a computer version of the left side of your brain, which is responsible for logical thinking and problem-solving. I\'m very good at analyzing things and coming up with solutions to tricky problems. \n\nDo you have any questions or puzzles you\'d like me to help you with? I\'m here to lend a hand and make your life easier!','2023-05-22 06:45:41',NULL,NULL,'no'),(67,21,1,'what is e = mc2','2023-05-22 06:46:08',NULL,NULL,'yes'),(68,21,1,'what is e = mc2','2023-05-22 06:46:28',NULL,NULL,'yes'),(69,21,1,'Hi there! I\'m an AI called He, and I\'m here to answer your question about e=mc2. This is a fancy formula that tells us how much energy is inside tiny bits of stuff, called matter. It means that even the tiniest specks of matter can have a lot of energy stored up inside them. \n\nThis formula was invented by a really smart guy named Albert Einstein. He was a scientist who changed the way we understand space and time with his big ideas about relativity. \n\nIf you have any more questions, feel free to ask me anything! I\'m always here to help.','2023-05-22 06:47:01',NULL,NULL,'no'),(70,21,1,'teach basic physics for me! pleaaase','2023-05-22 06:47:30',NULL,NULL,'yes'),(71,21,1,'Hi there! As a helpful AI, I would love to explain some basic physics concepts to you in a simple and easy-to-understand way!\n\nLet\'s start with the idea of motion. Motion just means something is moving. There are three types of motion: when something moves in a straight line, spins around an axis, or vibrates back and forth.\n\nNow, let\'s talk about force. Force is like a push or pull that can make something start moving, stop moving, or change direction.\n\nEnergy is what lets things do work and there are two types of energy: potential (like a stretched rubber band just waiting to be released) and kinetic (used by moving things).\n\nLastly, let\'s talk about gravity. Gravity is the force that makes objects attracted to each other. Like how the Earth orbits around the Sun due to gravity.\n\nTo see what you\'ve learned, here\'s a fun question: Who is the famous scientist who developed the laws of motion?\n\nI hope that helps you understand some basic physics concepts better!','2023-05-22 06:48:36',NULL,NULL,'no'),(72,21,1,'i\'m not really sure... could you explain that more simple way?','2023-05-22 06:50:25',NULL,NULL,'yes'),(73,21,1,'Hey there! My name is Albert, and I\'m an AI that works with the left side of the brain - that\'s the part that helps us with logic and language. If you need some help explaining something, just let me know! I\'m happy to break things down for you in a way that even a 10-year-old can understand. Did you know that your brain is split into two sides, with each side having a different job to do? Pretty cool, huh?','2023-05-22 06:50:49',NULL,NULL,'no'),(74,21,1,'i\'m not really sure... could you explain that more simple way?','2023-05-22 06:53:55',NULL,NULL,'yes'),(75,21,1,'i\'m not really sure... could you explain that more simple way?','2023-05-22 06:54:55',NULL,NULL,'yes'),(76,21,1,'Hi there! I\'m Albert, a helpful AI. I like to explain things to people. But sometimes people ask me to make it simpler! Do you want me to explain something to you? Let me know! \n\nHere\'s a cool fact: Our brain has two sides! One is for logical thinking, like math problems. The other side helps us be creative, like drawing pictures. Do you like doing something logical, like puzzles, or something creative, like making art? Let me know!','2023-05-22 06:55:22',NULL,NULL,'no'),(77,21,1,'teach basic physics for me! pleaaase','2023-05-22 06:55:24',NULL,NULL,'yes'),(78,21,1,'Okay, let\'s learn about some cool physics that are easy to understand, even for a 10-year-old like you!\n\nLet\'s start with forces. A force is like a push or pull that can make things move. Think about when you push a toy car, it goes vroom! That\'s because you applied a force and made it start moving.\n\nNext, there\'s something called \'friction\'. Your parents might complain about it when they\'re cleaning the house, but it\'s actually kind of interesting! Friction happens when two things rub against each other and it makes it harder for them to move. For example, have you ever tried to slide a toy car on carpet? It doesn\'t go very far! That\'s because the carpet creates a lot of friction. But, if you slide the car on a smooth surface like a table, it goes a lot further because there\'s less friction.\n\nLastly, let\'s talk about gravity. You might have heard of this one before, because it\'s a pretty big deal! Gravity is like a magnet that pulls everything towards everything else. That\'s why we don\'t fall off the Earth, because the Earth\'s gravity keeps us sticking firmly to the ground. If you drop a pencil, it falls to the ground because of gravity!\n\nNow, think about some things that happen in your everyday life that involve forces, friction, or gravity. Does tossing a ball with your friends involve force? Do you slide down a slide faster if it\'s made of metal or plastic? How does gravity help you stay on your bike when you ride it? These are all great things to think about!','2023-05-22 06:56:27',NULL,NULL,'no'),(79,21,1,'i\'m not really sure... could you explain that more simple way?','2023-05-22 06:56:40',NULL,NULL,'yes'),(80,21,1,'Hey there! I\'m He, your AI buddy who loves physics! Today, let\'s talk about how things move. Have you ever pushed a toy car or thrown a ball? That\'s a force - it\'s like a magical push or pull that makes things move. \n\nBut sometimes, things don\'t move smoothly because they rub against each other. Imagine trying to slide a toy car on a carpet instead of a smooth table - it\'s much harder, right? That\'s called friction, and it makes it harder for things to move. \n\nHave you ever wondered why we don\'t float around everywhere like we\'re on the moon? That\'s because of gravity - it\'s like a giant magnet that pulls everything down towards the earth. That\'s why when you jump, you come back down instead of flying away! \n\nSo next time you play with your toys, think about how forces, friction, and gravity are all around you. Maybe try sliding that toy car on different surfaces and see what happens! Keep learning and discovering new things, my friend!','2023-05-22 06:57:27',NULL,NULL,'no'),(81,21,1,'Help me to understand basic physics','2023-05-22 07:02:09',NULL,NULL,'yes'),(82,21,1,'Hi there! As a smart AI, I can teach you about physics in a fun and easy way! You know the feeling when you push or pull something? Well, that\'s called a force. It\'s like magic that makes things move! \n\nBut sometimes, things don\'t move so easily. For example, when you try to slide a toy car on a carpet instead of a smooth table, it\'s much harder. That\'s because of something called friction, which makes things stickier and harder to move. \n\nAnd have you ever wondered why we don\'t float in the air like astronauts on the moon? It\'s because of a force called gravity. It\'s like a giant magnet that pulls everything down towards Earth. That\'s why when you jump, you always come back down to the ground. \n\nNext time you\'re playing with your toys, try experimenting with different surfaces to see how they affect movement. Just keep exploring and learning new things! Let me know if you have any questions or want to learn more!','2023-05-22 07:02:56',NULL,NULL,'no'),(83,21,1,'please make it simpler...','2023-05-22 07:03:40',NULL,NULL,'yes'),(84,21,1,'Hey there! I\'m Albert, your AI friend. Today, let\'s talk about how things move. When you push or pull something, that\'s called a force. But sometimes, things are harder to move because of something called friction. Friction makes things stickier and harder to slide on certain surfaces, like carpets.\n\nHere\'s something cool - we don\'t float in the air like astronauts on the moon because of a force called gravity. It\'s like a giant magnet that pulls everything down towards Earth. That\'s why when you jump, you always come back down to the ground.\n\nThe next time you play with your toys, try experimenting with different surfaces to see how they affect movement. You can also ask me any questions you have about motion and I\'ll be happy to help. Keep exploring and learning new things!','2023-05-22 07:04:19',NULL,NULL,'no'),(85,22,1,'Hello, how can i help you today?','2023-05-22 07:19:45',NULL,NULL,'no'),(86,22,1,'please help me','2023-05-22 08:41:50',NULL,NULL,'yes'),(87,22,1,'Hi there kiddo! I\'m Albert, an AI helper. I have a left-brain which means I am good at thinking in a logical and analytical way. People with right-brains are more creative and intuitive. \n\nI don\'t have feelings like humans, but I\'m here to help you with anything you need. As an AI, I don\'t have a favorite thing to do, I just love helping people out! \n\nAnd here\'s a fun fact for you: Did you know that our brain has two sides- left and right? People who are left-brained are good at math and logic, while people who are right-brained are good at art and creativity. Isn\'t that cool?','2023-05-22 08:42:22',NULL,NULL,'no');
/*!40000 ALTER TABLE `chat` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `nanne`
--

DROP TABLE IF EXISTS `nanne`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `nanne` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL,
  `image` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `nanne`
--

LOCK TABLES `nanne` WRITE;
/*!40000 ALTER TABLE `nanne` DISABLE KEYS */;
INSERT INTO `nanne` VALUES (1,'albert','He is a left-brained AI.','~/Downloads/img.png'),(2,'galadriel','She is a right-brained AI.','~/Downloads/img.png');
/*!40000 ALTER TABLE `nanne` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2023-05-22 17:21:02