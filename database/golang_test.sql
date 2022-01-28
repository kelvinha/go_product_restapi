-- phpMyAdmin SQL Dump
-- version 5.0.2
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Jan 28, 2022 at 09:35 PM
-- Server version: 10.4.13-MariaDB
-- PHP Version: 7.3.19

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `golang_test`
--

-- --------------------------------------------------------

--
-- Table structure for table `table_product`
--

CREATE TABLE `table_product` (
  `product_id` int(11) NOT NULL,
  `product_name` varchar(50) DEFAULT NULL,
  `product_price` int(11) DEFAULT NULL,
  `product_description` text DEFAULT NULL,
  `product_quantity` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `table_product`
--

INSERT INTO `table_product` (`product_id`, `product_name`, `product_price`, `product_description`, `product_quantity`) VALUES
(4, 'Nike Vapor Max 2019', 1530000, 'Nike Vapor Max 2019 adalah sepatu sneakers Unisex model Low Cut dengan ujung sepatu berbentuk Almond. Dibuat dengan material insole dari Polyster dan outsole dari Rubber, sepatu sneakers Nike ini bisa membuat penampilan kamu terlihat semakin keren. Langsung beli dan dapatkan harga Nike Vapor Max 2019 yang termurah mulai dari IDR1449000 hanya di iPrice Indonesia.', 10),
(5, 'Nike Air Jordan 3 Retro SE Original', 1480000, 'Nike Air Jordan 3 Retro SE adalah sepatu sneakers Pria model High Cut dengan ujung sepatu berbentuk Almond. Dibuat dengan material insole dari Kulit dan outsole dari Karet, sepatu sneakers Nike ini bisa membuat penampilan kamu terlihat semakin keren. Langsung beli dan dapatkan harga Nike Air Jordan 3 Retro SE yang termurah mulai dari IDR1500000 hanya di iPrice Indonesia.', 30),
(6, 'Nike Air Huarache Run DNA CH', 1790000, 'Nike Air Huarache Run DNA Ch.1 adalah sepatu sneakers Unisex model Low Cut dengan ujung sepatu berbentuk Runcing. Dibuat dengan material insole dari Cotton dan outsole dari Rubber, sepatu sneakers Nike ini bisa membuat penampilan kamu terlihat semakin keren. Langsung beli dan dapatkan harga Nike Air Huarache Run DNA Ch.1 yang termurah mulai dari IDR1967700 hanya di iPrice Indonesia.', 10),
(7, 'Sepatu Vans Era Authentic', 500000, 'Born in Anaheim, California in 1966, the Authentic is the original Vans heritage style. Originally know as Vans #44 Deck Shoes, the Authentic became an immediate cult icon, and has embodied our “Off The Wall” attitude ever since.Constructed with a simple low top, lace-up profile, this classic shoe also features sturdy canvas uppers, metal eyelets, and signature rubber waffle outsoles.', 20),
(8, 'Sepatu Vans OLD SKOOL', 900000, 'First known as the Vans #36, the Old Skool debuted in 1977 with a unique new addition: a random doodle drawn by founder Paul Van Doren, and originally referred to as the “jazz stripe.” Today, the famous Vans Sidestripe has become the unmistakable—and instantly recognizable—hallmark of the Vans brand.Constructed with durable suede and canvas uppers in a range of fresh colorways, the Old Skool pays homage to our heritage while ensuring that this low top, lace-up shoe remains as iconic as ever. It also features re-enforced toe caps, supportive padded collars, and signature rubber waffle outsoles.', 10);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `table_product`
--
ALTER TABLE `table_product`
  ADD PRIMARY KEY (`product_id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `table_product`
--
ALTER TABLE `table_product`
  MODIFY `product_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=9;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
