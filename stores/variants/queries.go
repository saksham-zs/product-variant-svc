package variants

const insertIntoVariants = "INSERT INTO `variants` (`product_id`,`variant_name`,`variant_details`) VALUES (?,?,?)"
const getVariantByID = "SELECT id, variant_name, variant_details FROM `variants` WHERE product_id=? AND id=?"
const getVariantsByProductID = "SELECT id, variant_name, variant_details FROM `variants` WHERE product_id=?"
