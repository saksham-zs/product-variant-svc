package products

const insertIntoProducts = "INSERT INTO `products` (`name`,`brand_name`,`details`,`image_url`) VALUES (?,?,?,?)"
const getProductByID = "SELECT id, name, brand_name, details, image_url FROM `products` WHERE id=?"
const getAllProducts = "SELECT id, name, brand_name, details, image_url FROM `products`"
