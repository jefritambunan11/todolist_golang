-- phpMyAdmin SQL Dump
-- version 5.1.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Apr 04, 2023 at 09:20 AM
-- Server version: 10.4.21-MariaDB
-- PHP Version: 7.4.23

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `todolist`
--

-- --------------------------------------------------------

--
-- Table structure for table `todos`
--

CREATE TABLE `todos` (
  `id` int(255) NOT NULL,
  `todo` varchar(255) DEFAULT NULL,
  `date_time` datetime DEFAULT NULL,
  `user_id` int(255) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `todos`
--

INSERT INTO `todos` (`id`, `todo`, `date_time`, `user_id`, `created_at`, `updated_at`) VALUES
(1, 'Main Game Playstations 3', '2023-04-04 19:00:00', 1, '2023-04-04 12:11:52', '2023-04-04 12:11:52'),
(2, 'Main Game Playstations 3', '2023-04-04 19:00:00', 1, '2023-04-04 13:29:11', '2023-04-04 13:29:11'),
(3, 'Main Game Playstations 3', '2023-04-04 19:00:00', 1, '2023-04-04 13:29:12', '2023-04-04 13:29:12'),
(4, 'Main Game Playstations 3', '2023-04-04 19:00:00', 1, '2023-04-04 13:29:12', '2023-04-04 13:29:12'),
(5, 'Main Game Playstations 3', '2023-04-04 19:00:00', 1, '2023-04-04 13:29:13', '2023-04-04 13:29:13'),
(6, 'Main Game Playstations 3', '2023-04-04 19:00:00', 1, '2023-04-04 13:29:13', '2023-04-04 13:29:13'),
(7, 'Main Game Playstations 3', '2023-04-04 19:00:00', 1, '2023-04-04 13:29:14', '2023-04-04 13:29:14'),
(8, 'Main Game Playstations 3', '2023-04-04 19:00:00', 1, '2023-04-04 13:29:14', '2023-04-04 13:29:14'),
(9, 'Main Game Playstations 3', '2023-04-04 19:00:00', 1, '2023-04-04 13:29:15', '2023-04-04 13:29:15'),
(10, 'Main Game Playstations 3', '2023-04-04 19:00:00', 1, '2023-04-04 13:29:15', '2023-04-04 13:29:15'),
(11, 'Main Game Playstations 3', '2023-04-04 19:00:00', 1, '2023-04-04 13:29:16', '2023-04-04 13:29:16'),
(12, 'Main Game Playstations 3', '2023-04-04 19:00:00', 1, '2023-04-04 13:29:16', '2023-04-04 13:29:16'),
(13, 'Main Game Playstations 3', '2023-04-04 19:00:00', 1, '2023-04-04 13:29:17', '2023-04-04 13:29:17'),
(14, 'Main Game Playstations 3', '2023-04-04 19:00:00', 1, '2023-04-04 13:29:17', '2023-04-04 13:29:17'),
(15, 'Main Game Playstations 3', '2023-04-04 19:00:00', 1, '2023-04-04 13:29:17', '2023-04-04 13:29:17'),
(16, 'Main Game Playstations 3', '2023-04-04 19:00:00', 1, '2023-04-04 13:29:18', '2023-04-04 13:29:18'),
(17, 'Main Game Playstations 3', '2023-04-04 19:00:00', 1, '2023-04-04 13:29:18', '2023-04-04 13:29:18'),
(18, 'Main Game Playstations 3', '2023-04-04 19:00:00', 1, '2023-04-04 13:29:19', '2023-04-04 13:29:19'),
(19, 'Main Game Playstations 3', '2023-04-04 19:00:00', 1, '2023-04-04 13:29:20', '2023-04-04 13:29:20'),
(20, 'Main Game Playstations 3', '2023-04-04 19:00:00', 1, '2023-04-04 13:29:21', '2023-04-04 13:29:21'),
(21, 'Main Game Playstations 3', '2023-04-04 19:00:00', 1, '2023-04-04 13:32:52', '2023-04-04 13:32:52');

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` int(255) NOT NULL,
  `name` varchar(255) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `name`, `email`, `password`, `created_at`, `updated_at`) VALUES
(1, 'Jefri', 'jefri@gmail.com', '$2a$04$2Pp.Fg.2Jn.nGkfJuu.ECe.RmWcKFamiRPhAp0PGDL9PdSsmGU9xy', '2023-04-04 11:29:49', '2023-04-04 11:29:49');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `todos`
--
ALTER TABLE `todos`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `todos`
--
ALTER TABLE `todos`
  MODIFY `id` int(255) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=22;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` int(255) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
